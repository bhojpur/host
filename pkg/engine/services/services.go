package services

// Copyright (c) 2018 Bhojpur Consulting Private Limited, India. All rights reserved.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

import (
	"context"
	"fmt"

	"github.com/bhojpur/host/pkg/container/log"
	"github.com/bhojpur/host/pkg/engine/docker"
	"github.com/bhojpur/host/pkg/engine/hosts"
	v3 "github.com/bhojpur/host/pkg/engine/types"
	"github.com/bhojpur/host/pkg/engine/util"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/go-connections/nat"
	"github.com/sirupsen/logrus"
)

const (
	ETCDRole    = "etcd"
	ControlRole = "controlplane"
	WorkerRole  = "worker"

	SidekickServiceName   = "sidekick"
	RBACAuthorizationMode = "rbac"

	KubeAPIContainerName                        = "kube-apiserver"
	KubeletContainerName                        = "kubelet"
	KubeproxyContainerName                      = "kube-proxy"
	KubeControllerContainerName                 = "kube-controller-manager"
	SchedulerContainerName                      = "kube-scheduler"
	EtcdContainerName                           = "etcd"
	EtcdSnapshotContainerName                   = "etcd-rolling-snapshots"
	EtcdSnapshotOnceContainerName               = "etcd-snapshot-once"
	EtcdSnapshotRemoveContainerName             = "etcd-remove-snapshot"
	EtcdRestoreContainerName                    = "etcd-restore"
	EtcdDownloadBackupContainerName             = "etcd-download-backup"
	EtcdServeBackupContainerName                = "etcd-Serve-backup"
	EtcdChecksumContainerName                   = "etcd-checksum-checker"
	EtcdStateFileContainerName                  = "etcd-extract-statefile"
	ControlPlaneConfigMapStateFileContainerName = "extract-statefile-configmap"
	NginxProxyContainerName                     = "nginx-proxy"
	SidekickContainerName                       = "service-sidekick"
	LogLinkContainerName                        = "bke-log-linker"
	LogCleanerContainerName                     = "bke-log-cleaner"

	KubeAPIPort               = 6443
	SchedulerPortMaxV122      = 10251
	SchedulerPort             = 10259
	KubeControllerPortMaxV121 = 10252
	KubeControllerPort        = 10257
	KubeletPort               = 10248
	KubeproxyPort             = 10256

	WorkerThreads = util.WorkerThreads

	ContainerNameLabel = "net.bhojpur.bke.container.name"
	MCSLabel           = "label=level:s0:c1000,c1001"
	SELinuxLabel       = "label=type:bke_container_t"
)

type RestartFunc func(context.Context, *hosts.Host) error

func runSidekick(ctx context.Context, host *hosts.Host, prsMap map[string]v3.PrivateRegistry, sidecarProcess v3.Process, k8sVersion string) error {
	isRunning, err := docker.IsContainerRunning(ctx, host.DClient, host.Address, SidekickContainerName, true)
	if err != nil {
		return err
	}
	imageCfg, hostCfg, _ := GetProcessConfig(sidecarProcess, host, k8sVersion)
	isUpgradable := false
	if isRunning {
		isUpgradable, err = docker.IsContainerUpgradable(ctx, host.DClient, imageCfg, hostCfg, SidekickContainerName, host.Address, SidekickServiceName)
		if err != nil {
			return err
		}

		if !isUpgradable {
			log.Infof(ctx, "[%s] Sidekick container already created on host [%s]", SidekickServiceName, host.Address)
			return nil
		}
	}

	if err := docker.UseLocalOrPull(ctx, host.DClient, host.Address, sidecarProcess.Image, SidekickServiceName, prsMap); err != nil {
		return err
	}
	if isUpgradable {
		if err := docker.DoRemoveContainer(ctx, host.DClient, SidekickContainerName, host.Address); err != nil {
			return err
		}
	}
	if _, err := docker.CreateContainer(ctx, host.DClient, host.Address, SidekickContainerName, imageCfg, hostCfg); err != nil {
		return err
	}
	if host.DockerInfo.OSType == "windows" {
		// windows dockerfile VOLUME declaration must to satisfy one of them:
		//  - a non-existing or empty directory
		//  - a drive other than C:
		// so we could use a script to **start** the container to put expected resources into the "shared" directory,
		// like the action of `/usr/bin/sidecar.ps1` for windows bke-tools container
		return docker.StartContainer(ctx, host.DClient, host.Address, SidekickContainerName)
	}
	return nil
}

func removeSidekick(ctx context.Context, host *hosts.Host) error {
	return docker.DoRemoveContainer(ctx, host.DClient, SidekickContainerName, host.Address)
}

func GetProcessConfig(process v3.Process, host *hosts.Host, k8sVersion string) (*container.Config, *container.HostConfig, string) {
	imageCfg := &container.Config{
		Entrypoint: process.Command,
		Cmd:        process.Args,
		Env:        process.Env,
		Image:      process.Image,
		Labels:     process.Labels,
		User:       process.User,
	}
	// var pidMode container.PidMode
	// pidMode = process.PidMode
	_, portBindings, _ := nat.ParsePortSpecs(process.Publish)
	hostCfg := &container.HostConfig{
		VolumesFrom:  process.VolumesFrom,
		Binds:        process.Binds,
		NetworkMode:  container.NetworkMode(process.NetworkMode),
		PidMode:      container.PidMode(process.PidMode),
		Privileged:   process.Privileged,
		PortBindings: portBindings,
	}
	if len(process.RestartPolicy) > 0 {
		hostCfg.RestartPolicy = container.RestartPolicy{Name: process.RestartPolicy}
	}
	// The MCS label only needs to be applied when container is not running privileged, and running privileged negates need for applying the label
	// If Docker is configured with selinux-enabled:true, we need to specify MCS label to allow files from service-sidekick to be shared between containers
	if !process.Privileged && hosts.IsDockerSELinuxEnabled(host) {
		logrus.Debugf("Found selinux in DockerInfo.SecurityOptions on host [%s]", host.Address)
		// Check for containers having the sidekick container
		for _, volumeFrom := range hostCfg.VolumesFrom {
			if volumeFrom == SidekickContainerName {
				logrus.Debugf("Found [%s] in VolumesFrom on host [%s], applying MCSLabel [%s]", SidekickContainerName, host.Address, MCSLabel)
				hostCfg.SecurityOpt = []string{MCSLabel}
			}
		}
		// Check for sidekick container itself
		if value, ok := imageCfg.Labels[ContainerNameLabel]; ok {
			if value == SidekickContainerName {
				logrus.Debugf("Found [%s=%s] in Labels on host [%s], applying MCSLabel [%s]", ContainerNameLabel, SidekickContainerName, host.Address, MCSLabel)
				hostCfg.SecurityOpt = []string{MCSLabel}
			}
		}
		// We apply the label because we do not rewrite SELinux labels anymore on volume mounts (no :z)
		// Limited to Kubernetes 1.22 and higher
		matchedRange, _ := util.SemVerMatchRange(k8sVersion, util.SemVerK8sVersion122OrHigher)

		if matchedRange {
			logrus.Debugf("Applying security opt label [%s] for etcd container on host [%s]", SELinuxLabel, host.Address)
			hostCfg.SecurityOpt = append(hostCfg.SecurityOpt, SELinuxLabel)
		}

	}
	return imageCfg, hostCfg, process.HealthCheck.URL
}

func GetHealthCheckURL(useTLS bool, port int) string {
	if useTLS {
		return fmt.Sprintf("%s%s:%d%s", HTTPSProtoPrefix, HealthzAddress, port, HealthzEndpoint)
	}
	return fmt.Sprintf("%s%s:%d%s", HTTPProtoPrefix, HealthzAddress, port, HealthzEndpoint)
}

func createLogLink(ctx context.Context, host *hosts.Host, containerName, plane, image string, prsMap map[string]v3.PrivateRegistry) error {
	logrus.Debugf("[%s] Creating log link for Container [%s] on host [%s]", plane, containerName, host.Address)
	containerInspect, err := docker.InspectContainer(ctx, host.DClient, host.Address, containerName)
	if err != nil {
		return err
	}
	containerID := containerInspect.ID
	containerLogPath := containerInspect.LogPath
	containerLogLink := fmt.Sprintf("%s/%s_%s.log", hosts.BKELogsPath, containerName, containerID)
	imageCfg := &container.Config{
		Image: image,
		Tty:   true,
		Cmd: []string{
			"sh",
			"-c",
			fmt.Sprintf("mkdir -p %s ; ln -s %s %s", hosts.BKELogsPath, containerLogPath, containerLogLink),
		},
	}
	hostCfg := &container.HostConfig{
		Binds: []string{
			"/var/lib:/var/lib",
		},
		Privileged: true,
	}
	if err := docker.DoRemoveContainer(ctx, host.DClient, LogLinkContainerName, host.Address); err != nil {
		return err
	}
	if err := docker.DoRunContainer(ctx, host.DClient, imageCfg, hostCfg, LogLinkContainerName, host.Address, plane, prsMap); err != nil {
		return err
	}
	if err := docker.DoRemoveContainer(ctx, host.DClient, LogLinkContainerName, host.Address); err != nil {
		return err
	}
	logrus.Debugf("[%s] Successfully created log link for Container [%s] on host [%s]", plane, containerName, host.Address)
	return nil
}
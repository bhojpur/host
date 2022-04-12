package dind

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

	"github.com/bhojpur/host/pkg/engine/docker"
	"github.com/bhojpur/host/pkg/engine/util"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/sirupsen/logrus"
)

const (
	DINDImage           = "docker:19.03.12-dind"
	DINDContainerPrefix = "bke-dind"
	DINDPlane           = "dind"
	DINDNetwork         = "dind-network"
	DINDSubnet          = "172.18.0.0/16"
)

func StartUpDindContainer(ctx context.Context, dindAddress, dindNetwork, dindStorageDriver, dindDNS string) (string, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return "", err
	}
	// its recommended to use host's storage driver
	dockerInfo, err := cli.Info(ctx)
	if err != nil {
		return "", err
	}
	storageDriver := dindStorageDriver
	if len(storageDriver) == 0 {
		storageDriver = dockerInfo.Driver
	}

	// Get dind container name
	containerName := fmt.Sprintf("%s-%s", DINDContainerPrefix, dindAddress)
	_, err = cli.ContainerInspect(ctx, containerName)
	if err != nil {
		if !client.IsErrNotFound(err) {
			return "", err
		}
		if err := docker.UseLocalOrPull(ctx, cli, cli.DaemonHost(), DINDImage, DINDPlane, nil); err != nil {
			return "", err
		}
		binds := []string{
			fmt.Sprintf("/var/lib/kubelet-%s:/var/lib/kubelet:shared", containerName),
			"/etc/machine-id:/etc/machine-id:ro",
		}
		isLink, err := util.IsSymlink("/etc/resolv.conf")
		if err != nil {
			return "", err
		}
		if isLink {
			logrus.Infof("[%s] symlinked [/etc/resolv.conf] file detected. Using [%s] as DNS server.", DINDPlane, dindDNS)
		} else {
			binds = append(binds, "/etc/resolv.conf:/etc/resolv.conf")
		}
		imageCfg := &container.Config{
			Image: DINDImage,
			Entrypoint: []string{
				"sh",
				"-c",
				"mount --make-shared / && " +
					"mount --make-shared /sys && " +
					"mount --make-shared /var/lib/docker && " +
					"dockerd-entrypoint.sh --storage-driver=" + storageDriver,
			},
			Hostname: dindAddress,
			Env:      []string{"DOCKER_TLS_CERTDIR="},
		}
		hostCfg := &container.HostConfig{
			Privileged: true,
			Binds:      binds,
			// this gets ignored if resolv.conf is bind mounted. So it's ok to have it anyway.
			DNS: []string{dindDNS},
			// Calico needs this
			Sysctls: map[string]string{
				"net.ipv4.conf.all.rp_filter": "1",
			},
		}
		resp, err := cli.ContainerCreate(ctx, imageCfg, hostCfg, nil, containerName)
		if err != nil {
			return "", fmt.Errorf("Failed to create [%s] container on host [%s]: %v", containerName, cli.DaemonHost(), err)
		}

		if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
			return "", fmt.Errorf("Failed to start [%s] container on host [%s]: %v", containerName, cli.DaemonHost(), err)
		}
		logrus.Infof("[%s] Successfully started [%s] container on host [%s]", DINDPlane, containerName, cli.DaemonHost())
		dindContainer, err := cli.ContainerInspect(ctx, containerName)
		if err != nil {
			return "", fmt.Errorf("Failed to get the address of container [%s] on host [%s]: %v", containerName, cli.DaemonHost(), err)
		}
		dindIPAddress := dindContainer.NetworkSettings.IPAddress

		return dindIPAddress, nil
	}
	dindContainer, err := cli.ContainerInspect(ctx, containerName)
	if err != nil {
		return "", fmt.Errorf("Failed to get the address of container [%s] on host [%s]: %v", containerName, cli.DaemonHost(), err)
	}
	dindIPAddress := dindContainer.NetworkSettings.IPAddress
	logrus.Infof("[%s] container [%s] is already running on host[%s]", DINDPlane, containerName, cli.DaemonHost())
	return dindIPAddress, nil
}

func RmoveDindContainer(ctx context.Context, dindAddress string) error {
	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}
	containerName := fmt.Sprintf("%s-%s", DINDContainerPrefix, dindAddress)
	logrus.Infof("[%s] Removing dind container [%s] on host [%s]", DINDPlane, containerName, cli.DaemonHost())
	_, err = cli.ContainerInspect(ctx, containerName)
	if err != nil {
		if !client.IsErrNotFound(err) {
			return nil
		}
	}
	if err := cli.ContainerRemove(ctx, containerName, types.ContainerRemoveOptions{
		Force:         true,
		RemoveVolumes: true}); err != nil {
		if client.IsErrNotFound(err) {
			logrus.Debugf("[remove/%s] Container doesn't exist on host [%s]", containerName, cli.DaemonHost())
			return nil
		}
		return fmt.Errorf("Failed to remove dind container [%s] on host [%s]: %v", containerName, cli.DaemonHost(), err)
	}
	logrus.Infof("[%s] Successfully Removed dind container [%s] on host [%s]", DINDPlane, containerName, cli.DaemonHost())
	return nil
}

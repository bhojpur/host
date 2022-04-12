package cmd

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
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/bhojpur/host/pkg/container/log"
	"github.com/bhojpur/host/pkg/engine/cluster"
	"github.com/bhojpur/host/pkg/engine/dind"
	"github.com/bhojpur/host/pkg/engine/hosts"
	"github.com/bhojpur/host/pkg/engine/pki"
	"github.com/bhojpur/host/pkg/engine/pki/cert"
	v3 "github.com/bhojpur/host/pkg/engine/types"
	"github.com/urfave/cli"
)

const DINDWaitTime = 3

func UpCommand() cli.Command {
	upFlags := []cli.Flag{
		cli.StringFlag{
			Name:   "config",
			Usage:  "Specify an alternate cluster YAML file",
			Value:  pki.ClusterConfig,
			EnvVar: "BKE_CONFIG",
		},
		cli.BoolFlag{
			Name:  "local",
			Usage: "Deploy Kubernetes cluster locally",
		},
		cli.BoolFlag{
			Name:  "dind",
			Usage: "Deploy Kubernetes cluster in docker containers (experimental)",
		},
		cli.StringFlag{
			Name:  "dind-storage-driver",
			Usage: "Storage driver for the docker in docker containers (experimental)",
		},
		cli.StringFlag{
			Name:  "dind-dns-server",
			Usage: "DNS resolver to be used by docker in docker container. Useful if host is running systemd-resovld",
			Value: "8.8.8.8",
		},
		cli.BoolFlag{
			Name:  "update-only",
			Usage: "Skip idempotent deployment of control and etcd plane",
		},
		cli.BoolFlag{
			Name:  "disable-port-check",
			Usage: "Disable port check validation between nodes",
		},
		cli.BoolFlag{
			Name:  "init",
			Usage: "Initiate BKE cluster",
		},
		cli.StringFlag{
			Name:  "cert-dir",
			Usage: "Specify a certificate dir path",
		},
		cli.BoolFlag{
			Name:  "custom-certs",
			Usage: "Use custom certificates from a cert dir",
		},
	}

	upFlags = append(upFlags, commonFlags...)

	return cli.Command{
		Name:   "up",
		Usage:  "Bring the cluster up",
		Action: clusterUpFromCli,
		Flags:  upFlags,
	}
}

func ClusterUp(ctx context.Context, dialersOptions hosts.DialersOptions, flags cluster.ExternalFlags, data map[string]interface{}) (string, string, string, string, map[string]pki.CertificatePKI, error) {
	var APIURL, caCrt, clientCert, clientKey string
	var reconcileCluster, restore bool

	clusterState, err := cluster.ReadStateFile(ctx, cluster.GetStateFilePath(flags.ClusterFilePath, flags.ConfigDir))
	if err != nil {
		return APIURL, caCrt, clientCert, clientKey, nil, err
	}

	// We generate the first encryption config in ClusterInit, to store it ASAP. It's written to the DesiredState
	stateEncryptionConfig := clusterState.DesiredState.EncryptionConfig
	// if CurrentState has EncryptionConfig, it means this is NOT the first time we enable encryption, we should use the _latest_ applied value from the current cluster
	if clusterState.CurrentState.EncryptionConfig != "" {
		stateEncryptionConfig = clusterState.CurrentState.EncryptionConfig
	}

	kubeCluster, err := cluster.InitClusterObject(ctx, clusterState.DesiredState.BhojpurKubernetesEngineConfig.DeepCopy(), flags, stateEncryptionConfig)
	if err != nil {
		return APIURL, caCrt, clientCert, clientKey, nil, err
	}
	svcOptionsData := cluster.GetServiceOptionData(data)
	// check if rotate certificates is triggered
	if kubeCluster.BhojpurKubernetesEngineConfig.RotateCertificates != nil {
		return rebuildClusterWithRotatedCertificates(ctx, dialersOptions, flags, svcOptionsData)
	}
	// if we need to rotate the encryption key, do so and then return
	// note that we rotate the encryption key only when updating an existing cluster that has secret encryption enabled
	// all other cases will be handled later by reconciling the encryption provider config
	if kubeCluster.BhojpurKubernetesEngineConfig.RotateEncryptionKey {
		appliedConfig := clusterState.CurrentState.BhojpurKubernetesEngineConfig
		if appliedConfig != nil && appliedConfig.Services.KubeAPI.SecretsEncryptionConfig != nil && appliedConfig.Services.KubeAPI.SecretsEncryptionConfig.Enabled {
			return RotateEncryptionKey(ctx, clusterState.CurrentState.BhojpurKubernetesEngineConfig.DeepCopy(), dialersOptions, flags)
		}
	}

	log.Infof(ctx, "Building Kubernetes cluster")
	err = kubeCluster.SetupDialers(ctx, dialersOptions)
	if err != nil {
		return APIURL, caCrt, clientCert, clientKey, nil, err
	}
	err = kubeCluster.TunnelHosts(ctx, flags)
	if err != nil {
		return APIURL, caCrt, clientCert, clientKey, nil, err
	}
	currentCluster, err := kubeCluster.GetClusterState(ctx, clusterState)
	if err != nil {
		return APIURL, caCrt, clientCert, clientKey, nil, err
	}
	if !flags.DisablePortCheck {
		if err = kubeCluster.CheckClusterPorts(ctx, currentCluster); err != nil {
			return APIURL, caCrt, clientCert, clientKey, nil, err
		}
	}

	if err = kubeCluster.RunSELinuxCheck(ctx); err != nil {
		return APIURL, caCrt, clientCert, clientKey, nil, err
	}

	err = cluster.SetUpAuthentication(ctx, kubeCluster, currentCluster, clusterState)
	if err != nil {
		return APIURL, caCrt, clientCert, clientKey, nil, err
	}
	if len(kubeCluster.ControlPlaneHosts) > 0 {
		APIURL = fmt.Sprintf("https://%s:6443", kubeCluster.ControlPlaneHosts[0].Address)
	}
	clientCert = string(cert.EncodeCertPEM(kubeCluster.Certificates[pki.KubeAdminCertName].Certificate))
	clientKey = string(cert.EncodePrivateKeyPEM(kubeCluster.Certificates[pki.KubeAdminCertName].Key))
	caCrt = string(cert.EncodeCertPEM(kubeCluster.Certificates[pki.CACertName].Certificate))

	// moved deploying certs before reconcile to remove all unneeded certs generation from reconcile
	err = kubeCluster.SetUpHosts(ctx, flags)
	if err != nil {
		return APIURL, caCrt, clientCert, clientKey, nil, err
	}

	err = cluster.ReconcileCluster(ctx, kubeCluster, currentCluster, flags, svcOptionsData)
	if err != nil {
		return APIURL, caCrt, clientCert, clientKey, nil, err
	}

	/* reconcileCluster flag decides whether zero downtime upgrade logic is used or not.
	Zero-downtime upgrades should happen only when upgrading existing clusters. Not for new clusters or during etcd snapshot restore.
	currentCluster != nil indicates this is an existing cluster. Restore flag on DesiredState.BhojpurKubernetesEngineConfig indicates if it's a snapshot restore or not.
	reconcileCluster flag should be set to true only if currentCluster is not nil and restore is set to false
	*/
	if clusterState.DesiredState.BhojpurKubernetesEngineConfig != nil {
		restore = clusterState.DesiredState.BhojpurKubernetesEngineConfig.Restore.Restore
	}
	if currentCluster != nil && !restore {
		// reconcile this cluster, to check if upgrade is needed, or new nodes are getting added/removed
		/*This is to separate newly added nodes, so we don't try to check their status/cordon them before upgrade.
		This will also cover nodes that were considered inactive first time cluster was provisioned, but are now active during upgrade*/
		currentClusterNodes := make(map[string]bool)
		for _, node := range clusterState.CurrentState.BhojpurKubernetesEngineConfig.Nodes {
			currentClusterNodes[node.HostnameOverride] = true
		}

		newNodes := make(map[string]bool)
		for _, node := range clusterState.DesiredState.BhojpurKubernetesEngineConfig.Nodes {
			if !currentClusterNodes[node.HostnameOverride] {
				newNodes[node.HostnameOverride] = true
			}
		}
		kubeCluster.NewHosts = newNodes
		reconcileCluster = true

		maxUnavailableWorker, maxUnavailableControl, err := kubeCluster.CalculateMaxUnavailable()
		if err != nil {
			return APIURL, caCrt, clientCert, clientKey, nil, err
		}
		logrus.Infof("Setting maxUnavailable for worker nodes to: %v", maxUnavailableWorker)
		logrus.Infof("Setting maxUnavailable for controlplane nodes to: %v", maxUnavailableControl)
		kubeCluster.MaxUnavailableForWorkerNodes, kubeCluster.MaxUnavailableForControlNodes = maxUnavailableWorker, maxUnavailableControl
	}

	// update APIURL after reconcile
	if len(kubeCluster.ControlPlaneHosts) > 0 {
		APIURL = fmt.Sprintf("https://%s:6443", kubeCluster.ControlPlaneHosts[0].Address)
	}
	if err = cluster.ReconcileEncryptionProviderConfig(ctx, kubeCluster, currentCluster); err != nil {
		return APIURL, caCrt, clientCert, clientKey, nil, err
	}

	if err := kubeCluster.PrePullK8sImages(ctx); err != nil {
		return APIURL, caCrt, clientCert, clientKey, nil, err
	}

	errMsgMaxUnavailableNotFailedCtrl, err := kubeCluster.DeployControlPlane(ctx, svcOptionsData, reconcileCluster)
	if err != nil {
		return APIURL, caCrt, clientCert, clientKey, nil, err
	}
	// Apply Authz configuration after deploying controlplane
	err = cluster.ApplyAuthzResources(ctx, kubeCluster.BhojpurKubernetesEngineConfig, flags, dialersOptions)
	if err != nil {
		return APIURL, caCrt, clientCert, clientKey, nil, err
	}

	err = kubeCluster.UpdateClusterCurrentState(ctx, clusterState)
	if err != nil {
		return APIURL, caCrt, clientCert, clientKey, nil, err
	}

	err = cluster.SaveFullStateToKubernetes(ctx, kubeCluster, clusterState)
	if err != nil {
		return APIURL, caCrt, clientCert, clientKey, nil, err
	}

	errMsgMaxUnavailableNotFailedWrkr, err := kubeCluster.DeployWorkerPlane(ctx, svcOptionsData, reconcileCluster)
	if err != nil {
		return APIURL, caCrt, clientCert, clientKey, nil, err
	}

	if err = kubeCluster.CleanDeadLogs(ctx); err != nil {
		return APIURL, caCrt, clientCert, clientKey, nil, err
	}

	err = kubeCluster.SyncLabelsAndTaints(ctx, currentCluster)
	if err != nil {
		return APIURL, caCrt, clientCert, clientKey, nil, err
	}

	err = cluster.ConfigureCluster(ctx, kubeCluster.BhojpurKubernetesEngineConfig, kubeCluster.Certificates, flags, dialersOptions, data, false)
	if err != nil {
		return APIURL, caCrt, clientCert, clientKey, nil, err
	}
	if kubeCluster.EncryptionConfig.RewriteSecrets {
		if err = kubeCluster.RewriteSecrets(ctx); err != nil {
			return APIURL, caCrt, clientCert, clientKey, nil, err
		}
	}

	if err := checkAllIncluded(kubeCluster); err != nil {
		return APIURL, caCrt, clientCert, clientKey, nil, err
	}

	if errMsgMaxUnavailableNotFailedCtrl != "" || errMsgMaxUnavailableNotFailedWrkr != "" {
		return APIURL, caCrt, clientCert, clientKey, nil, fmt.Errorf(errMsgMaxUnavailableNotFailedCtrl + errMsgMaxUnavailableNotFailedWrkr)
	}
	log.Infof(ctx, "Finished building Kubernetes cluster successfully")
	return APIURL, caCrt, clientCert, clientKey, kubeCluster.Certificates, nil
}

func checkAllIncluded(cluster *cluster.Cluster) error {
	if len(cluster.InactiveHosts) == 0 {
		return nil
	}

	var names []string
	for _, host := range cluster.InactiveHosts {
		names = append(names, host.Address)
	}

	if len(names) > 0 {
		return fmt.Errorf("Provisioning incomplete, host(s) [%s] skipped because they could not be contacted", strings.Join(names, ","))
	}
	return nil
}

func clusterUpFromCli(ctx *cli.Context) error {
	logrus.Infof("Running BKE version: %v", ctx.App.Version)
	if ctx.Bool("local") {
		return clusterUpLocal(ctx)
	}
	if ctx.Bool("dind") {
		return clusterUpDind(ctx)
	}
	clusterFile, filePath, err := resolveClusterFile(ctx)
	if err != nil {
		return fmt.Errorf("Failed to resolve cluster file: %v", err)
	}

	bkeConfig, err := cluster.ParseConfig(clusterFile)
	if err != nil {
		return fmt.Errorf("Failed to parse cluster file: %v", err)
	}

	bkeConfig, err = setOptionsFromCLI(ctx, bkeConfig)
	if err != nil {
		return err
	}
	updateOnly := ctx.Bool("update-only")
	disablePortCheck := ctx.Bool("disable-port-check")
	// setting up the flags
	flags := cluster.GetExternalFlags(false, updateOnly, disablePortCheck, false, "", filePath)
	// Custom certificates and certificate dir flags
	flags.CertificateDir = ctx.String("cert-dir")
	flags.CustomCerts = ctx.Bool("custom-certs")
	if ctx.Bool("init") {
		return ClusterInit(context.Background(), bkeConfig, hosts.DialersOptions{}, flags)
	}
	if err := ClusterInit(context.Background(), bkeConfig, hosts.DialersOptions{}, flags); err != nil {
		return err
	}

	_, _, _, _, _, err = ClusterUp(context.Background(), hosts.DialersOptions{}, flags, map[string]interface{}{})
	return err
}

func clusterUpLocal(ctx *cli.Context) error {
	var bkeConfig *v3.BhojpurKubernetesEngineConfig
	clusterFile, filePath, err := resolveClusterFile(ctx)
	if err != nil {
		log.Infof(context.Background(), "Failed to resolve cluster file, using default cluster instead")
		bkeConfig = cluster.GetLocalBKEConfig()
	} else {
		bkeConfig, err = cluster.ParseConfig(clusterFile)
		if err != nil {
			return fmt.Errorf("Failed to parse cluster file: %v", err)
		}
		bkeConfig.Nodes = []v3.BKEConfigNode{*cluster.GetLocalBKENodeConfig()}
	}

	ignoreDockerVersion := ctx.Bool("ignore-docker-version")
	bkeConfig.IgnoreDockerVersion = &ignoreDockerVersion

	// setting up the dialers
	dialers := hosts.GetDialerOptions(nil, hosts.LocalHealthcheckFactory, nil)
	// setting up the flags
	flags := cluster.GetExternalFlags(true, false, false, false, "", filePath)

	if ctx.Bool("init") {
		return ClusterInit(context.Background(), bkeConfig, dialers, flags)
	}
	if err := ClusterInit(context.Background(), bkeConfig, dialers, flags); err != nil {
		return err
	}
	_, _, _, _, _, err = ClusterUp(context.Background(), dialers, flags, map[string]interface{}{})
	return err
}

func clusterUpDind(ctx *cli.Context) error {
	// get dind config
	bkeConfig, disablePortCheck, dindStorageDriver, filePath, dindDNS, err := getDindConfig(ctx)
	if err != nil {
		return err
	}
	// setup dind environment
	if err = createDINDEnv(context.Background(), bkeConfig, dindStorageDriver, dindDNS); err != nil {
		return err
	}

	// setting up the dialers
	dialers := hosts.GetDialerOptions(hosts.DindConnFactory, hosts.DindHealthcheckConnFactory, nil)
	// setting up flags
	flags := cluster.GetExternalFlags(false, false, disablePortCheck, false, "", filePath)
	flags.DinD = true

	if ctx.Bool("init") {
		return ClusterInit(context.Background(), bkeConfig, dialers, flags)
	}
	if err := ClusterInit(context.Background(), bkeConfig, dialers, flags); err != nil {
		return err
	}
	// start cluster
	_, _, _, _, _, err = ClusterUp(context.Background(), dialers, flags, map[string]interface{}{})
	return err
}

func getDindConfig(ctx *cli.Context) (*v3.BhojpurKubernetesEngineConfig, bool, string, string, string, error) {
	disablePortCheck := ctx.Bool("disable-port-check")
	dindStorageDriver := ctx.String("dind-storage-driver")
	dindDNS := ctx.String("dind-dns-server")

	clusterFile, filePath, err := resolveClusterFile(ctx)
	if err != nil {
		return nil, disablePortCheck, "", "", "", fmt.Errorf("Failed to resolve cluster file: %v", err)
	}

	bkeConfig, err := cluster.ParseConfig(clusterFile)
	if err != nil {
		return nil, disablePortCheck, "", "", "", fmt.Errorf("Failed to parse cluster file: %v", err)
	}

	bkeConfig, err = setOptionsFromCLI(ctx, bkeConfig)
	if err != nil {
		return nil, disablePortCheck, "", "", "", err
	}
	// Setting conntrack max for kubeproxy to 0
	if bkeConfig.Services.Kubeproxy.ExtraArgs == nil {
		bkeConfig.Services.Kubeproxy.ExtraArgs = make(map[string]string)
	}
	bkeConfig.Services.Kubeproxy.ExtraArgs["conntrack-max-per-core"] = "0"

	return bkeConfig, disablePortCheck, dindStorageDriver, filePath, dindDNS, nil
}

func createDINDEnv(ctx context.Context, bkeConfig *v3.BhojpurKubernetesEngineConfig, dindStorageDriver, dindDNS string) error {
	for i := range bkeConfig.Nodes {
		address, err := dind.StartUpDindContainer(ctx, bkeConfig.Nodes[i].Address, dind.DINDNetwork, dindStorageDriver, dindDNS)
		if err != nil {
			return err
		}
		if bkeConfig.Nodes[i].HostnameOverride == "" {
			bkeConfig.Nodes[i].HostnameOverride = bkeConfig.Nodes[i].Address
		}
		bkeConfig.Nodes[i].Address = address
	}
	time.Sleep(DINDWaitTime * time.Second)
	return nil
}

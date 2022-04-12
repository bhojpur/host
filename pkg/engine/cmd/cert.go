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
	"crypto/x509"
	"fmt"
	"time"

	"github.com/bhojpur/host/pkg/container/log"
	"github.com/bhojpur/host/pkg/engine/cluster"
	"github.com/bhojpur/host/pkg/engine/hosts"
	"github.com/bhojpur/host/pkg/engine/pki"
	"github.com/bhojpur/host/pkg/engine/pki/cert"
	"github.com/bhojpur/host/pkg/engine/services"
	v3 "github.com/bhojpur/host/pkg/engine/types"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func CertificateCommand() cli.Command {
	rotateFlags := []cli.Flag{
		cli.StringFlag{
			Name:   "config",
			Usage:  "Specify an alternate cluster YAML file",
			Value:  pki.ClusterConfig,
			EnvVar: "BKE_CONFIG",
		},
		cli.StringSliceFlag{
			Name: "service",
			Usage: fmt.Sprintf("Specify a k8s service to rotate certs, (allowed values: %s, %s, %s, %s, %s, %s)",
				services.KubeAPIContainerName,
				services.KubeControllerContainerName,
				services.SchedulerContainerName,
				services.KubeletContainerName,
				services.KubeproxyContainerName,
				services.EtcdContainerName,
			),
		},
		cli.BoolFlag{
			Name:  "rotate-ca",
			Usage: "Rotate all certificates including CA certs",
		},
	}
	rotateFlags = append(rotateFlags, commonFlags...)
	return cli.Command{
		Name:  "cert",
		Usage: "Certificates management for BKE cluster",
		Subcommands: cli.Commands{
			cli.Command{
				Name:   "rotate",
				Usage:  "Rotate BKE cluster certificates",
				Action: rotateBKECertificatesFromCli,
				Flags:  rotateFlags,
			},
			cli.Command{
				Name:   "generate-csr",
				Usage:  "Generate certificate sign requests for k8s components",
				Action: generateCSRFromCli,
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:   "config",
						Usage:  "Specify an alternate cluster YAML file",
						Value:  pki.ClusterConfig,
						EnvVar: "BKE_CONFIG",
					},
					cli.StringFlag{
						Name:  "cert-dir",
						Usage: "Specify a certificate dir path",
					},
				},
			},
		},
	}
}

func rotateBKECertificatesFromCli(ctx *cli.Context) error {
	logrus.Infof("Running BKE version: %v", ctx.App.Version)
	k8sComponents := ctx.StringSlice("service")
	rotateCACerts := ctx.Bool("rotate-ca")
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
	// setting up the flags
	externalFlags := cluster.GetExternalFlags(false, false, false, false, "", filePath)
	// setting up rotate flags
	bkeConfig.RotateCertificates = &v3.RotateCertificates{
		CACertificates: rotateCACerts,
		Services:       k8sComponents,
	}
	if err := ClusterInit(context.Background(), bkeConfig, hosts.DialersOptions{}, externalFlags); err != nil {
		return err
	}
	_, _, _, _, _, err = ClusterUp(context.Background(), hosts.DialersOptions{}, externalFlags, map[string]interface{}{})
	return err
}

func generateCSRFromCli(ctx *cli.Context) error {
	logrus.Infof("Running BKE version: %v", ctx.App.Version)
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
	// setting up the flags
	externalFlags := cluster.GetExternalFlags(false, false, false, false, "", filePath)
	externalFlags.CertificateDir = ctx.String("cert-dir")
	externalFlags.CustomCerts = ctx.Bool("custom-certs")

	return GenerateBKECSRs(context.Background(), bkeConfig, externalFlags)
}

func rebuildClusterWithRotatedCertificates(ctx context.Context,
	dialersOptions hosts.DialersOptions,
	flags cluster.ExternalFlags, svcOptionData map[string]*v3.KubernetesServicesOptions) (string, string, string, string, map[string]pki.CertificatePKI, error) {
	var APIURL, caCrt, clientCert, clientKey string
	log.Infof(ctx, "Rebuilding Kubernetes cluster with rotated certificates")
	clusterState, err := cluster.ReadStateFile(ctx, cluster.GetStateFilePath(flags.ClusterFilePath, flags.ConfigDir))
	if err != nil {
		return APIURL, caCrt, clientCert, clientKey, nil, err
	}

	kubeCluster, err := cluster.InitClusterObject(ctx, clusterState.DesiredState.BhojpurKubernetesEngineConfig.DeepCopy(), flags, clusterState.DesiredState.EncryptionConfig)
	if err != nil {
		return APIURL, caCrt, clientCert, clientKey, nil, err
	}
	if err := kubeCluster.SetupDialers(ctx, dialersOptions); err != nil {
		return APIURL, caCrt, clientCert, clientKey, nil, err
	}

	if err := kubeCluster.TunnelHosts(ctx, flags); err != nil {
		return APIURL, caCrt, clientCert, clientKey, nil, err
	}

	if err := cluster.SetUpAuthentication(ctx, kubeCluster, nil, clusterState); err != nil {
		return APIURL, caCrt, clientCert, clientKey, nil, err
	}
	if len(kubeCluster.ControlPlaneHosts) > 0 {
		APIURL = fmt.Sprintf("https://%s:6443", kubeCluster.ControlPlaneHosts[0].Address)
	}
	clientCert = string(cert.EncodeCertPEM(kubeCluster.Certificates[pki.KubeAdminCertName].Certificate))
	clientKey = string(cert.EncodePrivateKeyPEM(kubeCluster.Certificates[pki.KubeAdminCertName].Key))
	caCrt = string(cert.EncodeCertPEM(kubeCluster.Certificates[pki.CACertName].Certificate))

	if err := kubeCluster.SetUpHosts(ctx, flags); err != nil {
		return APIURL, caCrt, clientCert, clientKey, nil, err
	}

	// Save new State
	if err := saveClusterState(ctx, kubeCluster, clusterState); err != nil {
		return APIURL, caCrt, clientCert, clientKey, nil, err
	}

	// Restarting Kubernetes components
	servicesMap := make(map[string]bool)
	for _, component := range kubeCluster.RotateCertificates.Services {
		servicesMap[component] = true
	}

	if len(kubeCluster.RotateCertificates.Services) == 0 || kubeCluster.RotateCertificates.CACertificates || servicesMap[services.EtcdContainerName] {
		if err := services.RestartEtcdPlane(ctx, kubeCluster.EtcdHosts); err != nil {
			return APIURL, caCrt, clientCert, clientKey, nil, err
		}
	}
	isLegacyKubeAPI, err := cluster.IsLegacyKubeAPI(ctx, kubeCluster)
	if err != nil {
		return APIURL, caCrt, clientCert, clientKey, nil, err
	}
	if isLegacyKubeAPI {
		log.Infof(ctx, "[controlplane] Redeploying controlplane to update kubeapi parameters")
		if _, err := kubeCluster.DeployControlPlane(ctx, svcOptionData, true); err != nil {
			return APIURL, caCrt, clientCert, clientKey, nil, err
		}
	}
	if err := services.RestartControlPlane(ctx, kubeCluster.ControlPlaneHosts); err != nil {
		return APIURL, caCrt, clientCert, clientKey, nil, err
	}

	allHosts := hosts.GetUniqueHostList(kubeCluster.EtcdHosts, kubeCluster.ControlPlaneHosts, kubeCluster.WorkerHosts)
	if err := services.RestartWorkerPlane(ctx, allHosts); err != nil {
		return APIURL, caCrt, clientCert, clientKey, nil, err
	}

	if kubeCluster.RotateCertificates.CACertificates {
		if err := cluster.RestartClusterPods(ctx, kubeCluster); err != nil {
			return APIURL, caCrt, clientCert, clientKey, nil, err
		}
	}
	return APIURL, caCrt, clientCert, clientKey, kubeCluster.Certificates, nil
}

func saveClusterState(ctx context.Context, kubeCluster *cluster.Cluster, clusterState *cluster.FullState) error {
	var err error
	if err = kubeCluster.UpdateClusterCurrentState(ctx, clusterState); err != nil {
		return err
	}
	// Attempt to store cluster full state to Kubernetes
	for i := 1; i <= 3; i++ {
		err = cluster.SaveFullStateToKubernetes(ctx, kubeCluster, clusterState)
		if err != nil {
			time.Sleep(time.Second * time.Duration(2))
			continue
		}
		break
	}
	if err != nil {
		logrus.Warnf("Failed to save full cluster state to Kubernetes")
	}
	return nil
}

func rotateBKECertificates(ctx context.Context, kubeCluster *cluster.Cluster, flags cluster.ExternalFlags, bkeFullState *cluster.FullState) (*cluster.FullState, error) {
	log.Infof(ctx, "Rotating Kubernetes cluster certificates")
	currentCluster, err := kubeCluster.GetClusterState(ctx, bkeFullState)
	if err != nil {
		return nil, err
	}
	if currentCluster == nil {
		return nil, fmt.Errorf("Failed to rotate certificates: can't find old certificates")
	}
	currentCluster.RotateCertificates = kubeCluster.RotateCertificates
	if !kubeCluster.RotateCertificates.CACertificates {
		caCertPKI, ok := bkeFullState.CurrentState.CertificatesBundle[pki.CACertName]
		if !ok {
			return nil, fmt.Errorf("Failed to rotate certificates: can't find CA certificate")
		}
		caCert := caCertPKI.Certificate
		if caCert == nil {
			return nil, fmt.Errorf("Failed to rotate certificates: CA certificate is nil")
		}
		certPool := x509.NewCertPool()
		certPool.AddCert(caCert)
		if _, err := caCert.Verify(x509.VerifyOptions{Roots: certPool, KeyUsages: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth}}); err != nil {
			return nil, fmt.Errorf("Failed to rotate certificates: CA certificate is invalid, please use the --rotate-ca flag to rotate CA certificate, error: %v", err)
		}
	}
	if err := cluster.RotateBKECertificates(ctx, currentCluster, flags, bkeFullState); err != nil {
		return nil, err
	}
	bkeFullState.DesiredState.BhojpurKubernetesEngineConfig = &kubeCluster.BhojpurKubernetesEngineConfig
	return bkeFullState, nil
}

func GenerateBKECSRs(ctx context.Context, bkeConfig *v3.BhojpurKubernetesEngineConfig, flags cluster.ExternalFlags) error {
	log.Infof(ctx, "Generating Kubernetes cluster CSR certificates")
	if len(flags.CertificateDir) == 0 {
		flags.CertificateDir = cluster.GetCertificateDirPath(flags.ClusterFilePath, flags.ConfigDir)
	}

	certBundle, err := pki.ReadCSRsAndKeysFromDir(flags.CertificateDir)
	if err != nil {
		return err
	}

	// initialze the cluster object from the config file
	kubeCluster, err := cluster.InitClusterObject(ctx, bkeConfig, flags, "")
	if err != nil {
		return err
	}

	// Generating csrs for kubernetes components
	if err := pki.GenerateBKEServicesCSRs(ctx, certBundle, kubeCluster.BhojpurKubernetesEngineConfig); err != nil {
		return err
	}
	return pki.WriteCertificates(kubeCluster.CertificateDir, certBundle)
}

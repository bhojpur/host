package commands

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
	"github.com/bhojpur/host/pkg/engine/cluster"
	"github.com/bhojpur/host/pkg/engine/hosts"
	"github.com/bhojpur/host/pkg/engine/pki"
	"github.com/bhojpur/host/pkg/engine/pki/cert"
	v3 "github.com/bhojpur/host/pkg/engine/types"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func EncryptionCommand() cli.Command {
	encryptFlags := []cli.Flag{
		cli.StringFlag{
			Name:   "config",
			Usage:  "Specify an alternate cluster YAML file",
			Value:  pki.ClusterConfig,
			EnvVar: "BKE_CONFIG",
		},
	}
	encryptFlags = append(encryptFlags, commonFlags...)
	return cli.Command{
		Name:  "encrypt",
		Usage: "Manage cluster encryption provider keys",
		Subcommands: cli.Commands{
			cli.Command{
				Name:   "rotate-key",
				Usage:  "Rotate cluster encryption provider key",
				Action: rotateEncryptionKeyFromCli,
				Flags:  encryptFlags,
			},
		},
	}
}

func rotateEncryptionKeyFromCli(ctx *cli.Context) error {
	logrus.Infof("Running Bhojpur Kubernetes Engine version: %v", ctx.App.Version)
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
	flags := cluster.GetExternalFlags(false, false, false, false, "", filePath)

	_, _, _, _, _, err = RotateEncryptionKey(context.Background(), bkeConfig, hosts.DialersOptions{}, flags)
	return err
}

func RotateEncryptionKey(
	ctx context.Context,
	bkeConfig *v3.BhojpurKubernetesEngineConfig,
	dialersOptions hosts.DialersOptions,
	flags cluster.ExternalFlags,
) (string, string, string, string, map[string]pki.CertificatePKI, error) {
	log.Infof(ctx, "Rotating cluster secrets encryption key")

	var APIURL, caCrt, clientCert, clientKey string

	stateFilePath := cluster.GetStateFilePath(flags.ClusterFilePath, flags.ConfigDir)
	bkeFullState, _ := cluster.ReadStateFile(ctx, stateFilePath)

	// We generate the first encryption config in ClusterInit, to store it ASAP. It's written to the DesiredState
	stateEncryptionConfig := bkeFullState.DesiredState.EncryptionConfig
	// if CurrentState has EncryptionConfig, it means this is NOT the first time we enable encryption, we should use the _latest_ applied value from the current cluster
	if bkeFullState.CurrentState.EncryptionConfig != "" {
		stateEncryptionConfig = bkeFullState.CurrentState.EncryptionConfig
	}

	kubeCluster, err := cluster.InitClusterObject(ctx, bkeConfig, flags, stateEncryptionConfig)
	if err != nil {
		return APIURL, caCrt, clientCert, clientKey, nil, err
	}

	if kubeCluster.IsEncryptionCustomConfig() {
		return APIURL, caCrt, clientCert, clientKey, nil, fmt.Errorf("can't rotate encryption keys: Key Rotation is not supported with custom configuration")
	}
	if !kubeCluster.IsEncryptionEnabled() {
		return APIURL, caCrt, clientCert, clientKey, nil, fmt.Errorf("can't rotate encryption keys: Encryption Configuration is disabled. Please disable rotate_encryption_key and run hostops up again")
	}

	kubeCluster.Certificates = bkeFullState.DesiredState.CertificatesBundle
	if err := kubeCluster.SetupDialers(ctx, dialersOptions); err != nil {
		return APIURL, caCrt, clientCert, clientKey, nil, err
	}
	if err := kubeCluster.TunnelHosts(ctx, flags); err != nil {
		return APIURL, caCrt, clientCert, clientKey, nil, err
	}
	if len(kubeCluster.ControlPlaneHosts) > 0 {
		APIURL = fmt.Sprintf("https://%s:6443", kubeCluster.ControlPlaneHosts[0].Address)
	}
	clientCert = string(cert.EncodeCertPEM(kubeCluster.Certificates[pki.KubeAdminCertName].Certificate))
	clientKey = string(cert.EncodePrivateKeyPEM(kubeCluster.Certificates[pki.KubeAdminCertName].Key))
	caCrt = string(cert.EncodeCertPEM(kubeCluster.Certificates[pki.CACertName].Certificate))

	err = kubeCluster.RotateEncryptionKey(ctx, bkeFullState)
	if err != nil {
		return APIURL, caCrt, clientCert, clientKey, nil, err
	}

	// make sure we have the latest state
	bkeFullState, _ = cluster.ReadStateFile(ctx, stateFilePath)

	log.Infof(ctx, "Reconciling cluster state")
	if err := kubeCluster.ReconcileDesiredStateEncryptionConfig(ctx, bkeFullState); err != nil {
		return APIURL, caCrt, clientCert, clientKey, nil, err
	}

	log.Infof(ctx, "Cluster secrets encryption key rotated successfully")
	return APIURL, caCrt, clientCert, clientKey, kubeCluster.Certificates, nil
}

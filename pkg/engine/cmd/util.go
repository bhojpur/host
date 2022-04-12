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
	"encoding/json"
	"fmt"

	"github.com/bhojpur/host/pkg/engine/cluster"
	"github.com/bhojpur/host/pkg/engine/hosts"
	"github.com/bhojpur/host/pkg/engine/k8s"
	"github.com/bhojpur/host/pkg/engine/pki"
	v3 "github.com/bhojpur/host/pkg/engine/types"
	"github.com/bhojpur/host/pkg/engine/util"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func UtilCommand() cli.Command {
	utilCfgFlags := []cli.Flag{
		cli.StringFlag{
			Name:   "config",
			Usage:  "Specify an alternate cluster YAML file",
			Value:  pki.ClusterConfig,
			EnvVar: "BKE_CONFIG",
		},
	}
	utilFlags := append(utilCfgFlags, commonFlags...)

	return cli.Command{
		Name:  "util",
		Usage: "Various utilities to retrieve cluster related files and troubleshoot",
		Subcommands: cli.Commands{
			cli.Command{
				Name:   "get-state-file",
				Usage:  "Retrieve state file from cluster",
				Action: getStateFile,
				Flags:  utilFlags,
			},
			cli.Command{
				Name:   "get-kubeconfig",
				Usage:  "Retrieve kubeconfig file from cluster state",
				Action: getKubeconfigFile,
				Flags:  utilFlags,
			},
		},
	}
}

func getKubeconfigFile(ctx *cli.Context) error {
	logrus.Infof("Creating new kubeconfig file")
	// Check if we can successfully connect to the cluster using the existing kubeconfig file
	clusterFile, clusterFilePath, err := resolveClusterFile(ctx)
	if err != nil {
		return fmt.Errorf("failed to resolve cluster file: %v", err)
	}

	// setting up the flags
	flags := cluster.GetExternalFlags(false, false, false, false, "", clusterFilePath)
	bkeConfig, err := cluster.ParseConfig(clusterFile)
	if err != nil {
		return fmt.Errorf("failed to parse cluster file: %v", err)
	}

	bkeConfig, err = setOptionsFromCLI(ctx, bkeConfig)
	if err != nil {
		return err
	}

	clusterState, err := cluster.ReadStateFile(context.Background(), cluster.GetStateFilePath(flags.ClusterFilePath, flags.ConfigDir))
	if err != nil {
		return err
	}

	// Creating temp cluster to check if snapshot archive contains state file and retrieve it
	tempCluster, err := cluster.InitClusterObject(context.Background(), bkeConfig, flags, "")
	if err != nil {
		return err
	}

	// Move current kubeconfig file
	err = util.CopyFileWithPrefix(tempCluster.LocalKubeConfigPath, "kube_config")
	if err != nil {
		return err
	}
	kubeCluster, _ := tempCluster.GetClusterState(context.Background(), clusterState)

	return cluster.RebuildKubeconfig(context.Background(), kubeCluster)
}

func getStateFile(ctx *cli.Context) error {
	logrus.Infof("Retrieving state file from cluster")
	// Check if we can successfully connect to the cluster using the existing kubeconfig file
	localKubeConfig := pki.GetLocalKubeConfig(ctx.String("config"), "")
	clusterFile, clusterFilePath, err := resolveClusterFile(ctx)
	if err != nil {
		return fmt.Errorf("failed to resolve cluster file: %v", err)
	}
	// setting up the flags
	flags := cluster.GetExternalFlags(false, false, false, false, "", clusterFilePath)

	// not going to use a k8s dialer here.. this is a CLI command
	serverVersion, err := cluster.GetK8sVersion(localKubeConfig, nil)
	if err != nil {
		logrus.Infof("Unable to connect to server using kubeconfig, trying to get state from Control Plane node(s), error: %v", err)
		// We need to retrieve the state file using Docker on the node(s)

		bkeConfig, err := cluster.ParseConfig(clusterFile)
		if err != nil {
			return fmt.Errorf("failed to parse cluster file: %v", err)
		}

		bkeConfig, err = setOptionsFromCLI(ctx, bkeConfig)
		if err != nil {
			return err
		}

		_, _, _, _, _, err = RetrieveClusterStateConfigMap(context.Background(), bkeConfig, hosts.DialersOptions{}, flags, map[string]interface{}{})
		if err != nil {
			return err
		}

		return nil
	}
	logrus.Infof("Successfully connected to server using kubeconfig, retrieved server version [%s]", serverVersion)
	// Retrieve full-cluster-state configmap
	k8sClient, err := k8s.NewClient(localKubeConfig, nil)
	if err != nil {
		return err
	}
	cfgMap, err := k8s.GetConfigMap(k8sClient, cluster.FullStateConfigMapName)
	if err != nil {
		return err
	}
	clusterData := cfgMap.Data[cluster.FullStateConfigMapName]
	bkeFullState := &cluster.FullState{}
	if err = json.Unmarshal([]byte(clusterData), bkeFullState); err != nil {
		return err
	}

	// Move current state file
	stateFilePath := cluster.GetStateFilePath(flags.ClusterFilePath, flags.ConfigDir)
	err = util.ReplaceFileWithBackup(stateFilePath, "bkestate")
	if err != nil {
		return err
	}

	// Write new state file
	err = bkeFullState.WriteStateFile(context.Background(), stateFilePath)
	if err != nil {
		return err
	}

	return nil
}

func RetrieveClusterStateConfigMap(
	ctx context.Context,
	bkeConfig *v3.BhojpurKubernetesEngineConfig,
	dialersOptions hosts.DialersOptions,
	flags cluster.ExternalFlags,
	data map[string]interface{}) (string, string, string, string, map[string]pki.CertificatePKI, error) {
	var APIURL, caCrt, clientCert, clientKey string

	bkeFullState := &cluster.FullState{}

	// Creating temp cluster to check if snapshot archive contains state file and retrieve it
	tempCluster, err := cluster.InitClusterObject(ctx, bkeConfig, flags, "")
	if err != nil {
		return APIURL, caCrt, clientCert, clientKey, nil, err
	}
	if err := tempCluster.SetupDialers(ctx, dialersOptions); err != nil {
		return APIURL, caCrt, clientCert, clientKey, nil, err
	}
	if err := tempCluster.TunnelHosts(ctx, flags); err != nil {
		return APIURL, caCrt, clientCert, clientKey, nil, err
	}
	// Get ConfigMap containing cluster state from Control Plane Hosts
	stateFile, err := tempCluster.GetStateFileFromConfigMap(ctx)

	if err != nil {
		return APIURL, caCrt, clientCert, clientKey, nil, err
	}
	bkeFullState, err = cluster.StringToFullState(ctx, stateFile)

	if err != nil {
		return APIURL, caCrt, clientCert, clientKey, nil, err
	}

	// Move current state file
	stateFilePath := cluster.GetStateFilePath(flags.ClusterFilePath, flags.ConfigDir)
	err = util.ReplaceFileWithBackup(stateFilePath, "bkestate")
	if err != nil {
		return APIURL, caCrt, clientCert, clientKey, nil, err
	}

	err = bkeFullState.WriteStateFile(context.Background(), stateFilePath)
	if err != nil {
		return APIURL, caCrt, clientCert, clientKey, nil, err
	}

	// Move current kubeconfig file
	err = util.CopyFileWithPrefix(tempCluster.LocalKubeConfigPath, "kube_config")
	if err != nil {
		return APIURL, caCrt, clientCert, clientKey, nil, err
	}
	kubeCluster, _ := tempCluster.GetClusterState(ctx, bkeFullState)

	if err := cluster.RebuildKubeconfig(ctx, kubeCluster); err != nil {
		return APIURL, caCrt, clientCert, clientKey, nil, nil
	}

	return APIURL, caCrt, clientCert, clientKey, nil, nil
}

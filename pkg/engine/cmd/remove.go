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
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/bhojpur/host/pkg/container/log"
	"github.com/bhojpur/host/pkg/engine/cluster"
	"github.com/bhojpur/host/pkg/engine/dind"
	"github.com/bhojpur/host/pkg/engine/hosts"
	"github.com/bhojpur/host/pkg/engine/pki"
	v3 "github.com/bhojpur/host/pkg/engine/types"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func RemoveCommand() cli.Command {
	removeFlags := []cli.Flag{
		cli.StringFlag{
			Name:   "config",
			Usage:  "Specify an alternate cluster YAML file",
			Value:  pki.ClusterConfig,
			EnvVar: "BKE_CONFIG",
		},
		cli.BoolFlag{
			Name:  "force",
			Usage: "Force removal of the cluster",
		},
		cli.BoolFlag{
			Name:  "local",
			Usage: "Remove Kubernetes cluster locally",
		},
		cli.BoolFlag{
			Name:  "dind",
			Usage: "Remove Kubernetes cluster deployed in dind mode",
		},
	}

	removeFlags = append(removeFlags, commonFlags...)

	return cli.Command{
		Name:   "remove",
		Usage:  "Teardown the cluster and clean cluster nodes",
		Action: clusterRemoveFromCli,
		Flags:  removeFlags,
	}
}

func ClusterRemove(
	ctx context.Context,
	bkeConfig *v3.BhojpurKubernetesEngineConfig,
	dialersOptions hosts.DialersOptions,
	flags cluster.ExternalFlags) error {

	log.Infof(ctx, "Tearing down Kubernetes cluster")

	kubeCluster, err := cluster.InitClusterObject(ctx, bkeConfig, flags, "")
	if err != nil {
		return err
	}
	if err := kubeCluster.SetupDialers(ctx, dialersOptions); err != nil {
		return err
	}

	err = kubeCluster.TunnelHosts(ctx, flags)
	if err != nil {
		return err
	}

	logrus.Debugf("Starting Cluster removal")
	err = kubeCluster.ClusterRemove(ctx)
	if err != nil {
		return err
	}

	log.Infof(ctx, "Cluster removed successfully")
	return nil
}

func clusterRemoveFromCli(ctx *cli.Context) error {
	logrus.Infof("Running BKE version: %v", ctx.App.Version)
	if ctx.Bool("local") {
		return clusterRemoveLocal(ctx)
	}
	clusterFile, filePath, err := resolveClusterFile(ctx)
	if err != nil {
		return fmt.Errorf("Failed to resolve cluster file: %v", err)
	}
	force := ctx.Bool("force")
	if !force {
		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Are you sure you want to remove Kubernetes cluster [y/n]: ")
		input, err := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if err != nil {
			return err
		}
		if input != "y" && input != "Y" {
			return nil
		}
	}
	if ctx.Bool("dind") {
		return clusterRemoveDind(ctx)
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

	return ClusterRemove(context.Background(), bkeConfig, hosts.DialersOptions{}, flags)
}

func clusterRemoveLocal(ctx *cli.Context) error {
	var bkeConfig *v3.BhojpurKubernetesEngineConfig
	clusterFile, filePath, err := resolveClusterFile(ctx)
	if err != nil {
		log.Warnf(context.Background(), "Failed to resolve cluster file, using default cluster instead")
		bkeConfig = cluster.GetLocalBKEConfig()
	} else {
		bkeConfig, err = cluster.ParseConfig(clusterFile)
		if err != nil {
			return fmt.Errorf("Failed to parse cluster file: %v", err)
		}
		bkeConfig.Nodes = []v3.BKEConfigNode{*cluster.GetLocalBKENodeConfig()}
	}

	bkeConfig, err = setOptionsFromCLI(ctx, bkeConfig)
	if err != nil {
		return err
	}
	// setting up the flags
	flags := cluster.GetExternalFlags(true, false, false, false, "", filePath)

	return ClusterRemove(context.Background(), bkeConfig, hosts.DialersOptions{}, flags)
}

func clusterRemoveDind(ctx *cli.Context) error {
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

	for _, node := range bkeConfig.Nodes {
		if err = dind.RmoveDindContainer(context.Background(), node.Address); err != nil {
			return err
		}
	}
	// remove the kube config file
	localKubeConfigPath := pki.GetLocalKubeConfig(filePath, "")
	pki.RemoveAdminConfig(context.Background(), localKubeConfigPath)

	// remove cluster state file
	stateFilePath := cluster.GetStateFilePath(filePath, "")
	cluster.RemoveStateFile(context.Background(), stateFilePath)
	return err
}

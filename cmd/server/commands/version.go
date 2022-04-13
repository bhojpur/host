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
	"fmt"

	"github.com/sirupsen/logrus"

	"github.com/bhojpur/host/pkg/engine/cluster"
	"github.com/bhojpur/host/pkg/engine/pki"
	"github.com/urfave/cli"
)

func VersionCommand() cli.Command {
	versionFlags := []cli.Flag{
		cli.StringFlag{
			Name:   "config",
			Usage:  "Specify an alternate cluster YAML file",
			Value:  pki.ClusterConfig,
			EnvVar: "BKE_CONFIG",
		},
	}
	return cli.Command{
		Name:   "version",
		Usage:  "Show cluster Kubernetes version",
		Action: getClusterVersion,
		Flags:  versionFlags,
	}
}

func getClusterVersion(ctx *cli.Context) error {
	logrus.Infof("Running Bhojpur Kubernetes Engine version: %v", ctx.App.Version)
	localKubeConfig := pki.GetLocalKubeConfig(ctx.String("config"), "")
	// not going to use a k8s dialer here.. this is a CLI command
	serverVersion, err := cluster.GetK8sVersion(localKubeConfig, nil)
	if err != nil {
		return err
	}
	fmt.Printf("Server version: %s\n", serverVersion)
	return nil
}

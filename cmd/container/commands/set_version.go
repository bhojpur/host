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

	"github.com/bhojpur/host/pkg/container/store"
	"github.com/bhojpur/host/pkg/container/types"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func SetVersionCommand() cli.Command {
	return cli.Command{
		Name:      "set-version",
		ShortName: "sv",
		Usage:     "Set the version of Kubernetes",
		Action:    setVersion,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "version",
				Usage: "The version to upgade/downgrade kubernetes to",
			},
		},
	}
}

func setVersion(ctx *cli.Context) error {
	debug := lookUpDebugFlag()
	if debug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	clusters, err := store.GetAllClusterFromStore()

	if err != nil {
		return err
	}

	for _, name := range ctx.Args() {
		if name == "" || name == "--help" {
			return cli.ShowCommandHelp(ctx, "set-version")
		}

		cluster, ok := clusters[name]

		if !ok {
			err = fmt.Errorf("could not find cluster: %v", err)
			logrus.Error(err.Error())

			return err
		}

		rpcClient, _, err := runRPCDriver(cluster.DriverName)

		if err != nil {
			return err
		}

		configGetter := cliConfigGetter{
			name: name,
			ctx:  ctx,
		}

		cluster.ConfigGetter = configGetter
		cluster.PersistStore = store.CLIPersistStore{}
		cluster.Driver = rpcClient

		cap, err := cluster.GetCapabilities(context.Background())
		if err != nil {
			return fmt.Errorf("error getting capabilities: %v", err)
		}

		if cap.HasSetVersionCapability() {
			err := cluster.SetVersion(context.Background(), &types.KubernetesVersion{Version: ctx.String("version")})

			if err != nil {
				return err
			}

			fmt.Printf("%v updated to %v\n", name, ctx.String("version"))
		} else {
			return fmt.Errorf("no set-version capability available")
		}
	}

	return nil
}

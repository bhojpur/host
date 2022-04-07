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
	"github.com/urfave/cli"
)

// RmCommand defines the remove command
func RmCommand() cli.Command {
	return cli.Command{
		Name:      "remove",
		ShortName: "rm",
		Usage:     "Remove kubernetes clusters",
		Action:    rmCluster,
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "force,f",
				Usage: "force to remove a cluster",
			},
		},
	}
}

func rmCluster(ctx *cli.Context) error {
	var lastErr error

	for _, name := range ctx.Args() {
		if name == "" || name == "--help" {
			return cli.ShowCommandHelp(ctx, "remove")
		}
		clusters, err := store.GetAllClusterFromStore()
		if err != nil {
			lastErr = err
			continue
		}
		cluster, ok := clusters[name]
		if !ok {
			lastErr = fmt.Errorf("cluster %v can't be found", name)
			continue
		}
		rpcClient, _, err := runRPCDriver(cluster.DriverName)
		if err != nil {
			lastErr = err
			continue
		}
		configGetter := cliConfigGetter{
			name: name,
			ctx:  ctx,
		}
		cluster.ConfigGetter = configGetter
		cluster.PersistStore = store.CLIPersistStore{}
		cluster.Driver = rpcClient
		if err := cluster.Remove(context.Background(), true); err != nil {
			if ctx.Bool("force") {
				cluster.PersistStore.Remove(name)
			} else {
				lastErr = err
				continue
			}
		}

		fmt.Println(cluster.Name)
	}

	return lastErr
}

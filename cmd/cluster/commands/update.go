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
	"errors"
	"fmt"
	"os"

	"github.com/bhojpur/host/pkg/cluster/store"
	"github.com/bhojpur/host/pkg/cluster/types"
	"github.com/urfave/cli"
)

var updateHelpTmeplate = `{{.Usage}}
{{if .Description}}{{.Description}}{{end}}
Usage: hostfarm [global option] {{.Name}} {{if .Flags}}[OPTIONS] {{end}}{{if ne "None" .ArgsUsage}}{{if ne "" .ArgsUsage}}{{.ArgsUsage}}{{else}}[cluster-name]{{end}}{{end}}

{{if .Flags}}Options:{{range .Flags}}
	 {{.}}{{end}}{{end}}
`

// UpdateCommand defines the update command
func UpdateCommand() cli.Command {
	return cli.Command{
		Name:               "update",
		Usage:              "update Kubernetes clusters",
		Action:             updateWrapper,
		SkipFlagParsing:    true,
		CustomHelpTemplate: updateHelpTmeplate,
	}
}

func updateWrapper(ctx *cli.Context) error {
	name := ctx.Args().Get(len(ctx.Args()) - 1)
	if name == "--help" {
		if len(ctx.Args())-2 >= 0 {
			name = ctx.Args().Get(len(ctx.Args()) - 2)
		} else {
			return cli.ShowCommandHelp(ctx, "update")
		}
	}
	clusters, err := store.GetAllClusterFromStore()
	if err != nil {
		return err
	}
	cluster, ok := clusters[name]
	if !ok {
		return fmt.Errorf("cluster %v can't be found", name)
	}
	rpcClient, addr, err := runRPCDriver(cluster.DriverName)
	if err != nil {
		return err
	}

	driverFlags, err := rpcClient.GetDriverUpdateOptions(context.Background())
	if err != nil {
		return err
	}
	flags := getDriverFlags(driverFlags)
	for i, command := range ctx.App.Commands {
		if command.Name == "update" {
			updateCmd := &ctx.App.Commands[i]
			updateCmd.SkipFlagParsing = false
			updateCmd.Flags = append(updateCmd.Flags, flags...)
			updateCmd.Action = updateCluster
		}
	}
	if len(os.Args) > 1 && addr != "" {
		args := []string{os.Args[0], "--plugin-listen-addr", addr}
		args = append(args, os.Args[1:len(os.Args)]...)
		return ctx.App.Run(args)
	}
	return ctx.App.Run(os.Args)
}

func updateCluster(ctx *cli.Context) error {
	name := ctx.Args().Get(0)
	if name == "" {
		return errors.New("name is required when inspecting cluster")
	} else if name == "--help" {
		// in case of `./hostfarm update cluster1 --help`
		return cli.ShowCommandHelp(ctx, "update")
	}
	clusters, err := store.GetAllClusterFromStore()
	if err != nil {
		return err
	}
	cluster, ok := clusters[name]
	if !ok {
		return fmt.Errorf("cluster %v can't be found", name)
	}
	addr := ctx.GlobalString("plugin-listen-addr")
	rpcClient, err := types.NewClient(cluster.DriverName, addr)
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
	return cluster.Update(context.Background())
}

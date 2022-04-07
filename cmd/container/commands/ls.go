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
	"github.com/bhojpur/host/pkg/container/store"
	"github.com/bhojpur/host/pkg/container/utils"
	"github.com/urfave/cli"
)

// LsCommand defines the ls command
func LsCommand() cli.Command {
	return cli.Command{
		Name:      "list",
		ShortName: "ls",
		Usage:     "list kubernetes clusters",
		Action:    lsCluster,
		Flags:     []cli.Flag{},
	}
}

func lsCluster(ctx *cli.Context) error {
	// todo: add filter support
	clusters, err := store.GetAllClusterFromStore()
	if err != nil {
		return err
	}

	writer := utils.NewTableWriter([][]string{
		{"NAME", "Name"},
		{"DRIVER", "DriverName"},
		{"VERISON", "Version"},
		{"ENDPOINT", "Endpoint"},
		{"NODE_COUNT", "NodeCount"},
		{"STATUS", "Status"},
	}, ctx)
	defer writer.Close()
	for _, cluster := range clusters {
		writer.Write(cluster)
	}
	return writer.Err()
}

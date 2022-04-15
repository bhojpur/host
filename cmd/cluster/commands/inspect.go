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
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/bhojpur/host/pkg/cluster/store"
	"github.com/urfave/cli"
)

var inspectHelpTemplate = `{{.Usage}}
{{if .Description}}{{.Description}}{{end}}
Usage: hostfarm [global options] {{.Name}} {{if .Flags}}[OPTIONS] {{end}}{{if ne "None" .ArgsUsage}}{{if ne "" .ArgsUsage}}{{.ArgsUsage}}{{else}}[cluster-name]{{end}}{{end}}

{{if .Flags}}Options:{{range .Flags}}
	 {{.}}{{end}}{{end}}
`

// InspectCommand defines the inspect command
func InspectCommand() cli.Command {
	return cli.Command{
		Name:               "inspect",
		Usage:              "inspect Kubernetes clusters",
		Action:             inspectCluster,
		Flags:              []cli.Flag{},
		CustomHelpTemplate: inspectHelpTemplate,
	}
}

func inspectCluster(ctx *cli.Context) error {
	name := ctx.Args().Get(0)
	if name == "" {
		return errors.New("name is required when inspecting cluster")
	}
	clusters, err := store.GetAllClusterFromStore()
	if err != nil {
		return err
	}
	cluster, ok := clusters[name]
	if !ok {
		return fmt.Errorf("cluster %v can't be found", name)
	}
	cluster.ClientKey = "Redacted"
	cluster.ClientCertificate = "Redacted"
	cluster.RootCACert = "Redacted"
	data, err := json.MarshalIndent(cluster, "", "\t")
	if err != nil {
		return err
	}
	_, err = os.Stdout.Write(data)
	return err
}

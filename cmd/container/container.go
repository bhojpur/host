package main

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
	"os"
	"path/filepath"

	cmd "github.com/bhojpur/host/cmd/container/commands"
	"github.com/bhojpur/host/pkg/machine/log"
	"github.com/bhojpur/host/pkg/version"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

var appHelpTemplate = `{{.Usage}}

Usage: {{.Name}} {{if .Flags}}[GLOBAL_OPTIONS] {{end}}COMMAND [arg...]

Version: {{.Version}}{{if or .Author .Email}}
{{if .Flags}}
Options:
  {{range .Flags}}{{if .Hidden}}{{else}}{{.}}
  {{end}}{{end}}{{end}}
Commands:
  {{range .Commands}}{{.Name}}{{with .ShortName}}, {{.}}{{end}}{{ "\t" }}{{.Usage}}
  {{end}}
Run '{{.Name}} COMMAND --help' for more information on a command.
`

var commandHelpTemplate = `{{.Usage}}
{{if .Description}}{{.Description}}{{end}}
Usage: hostfarm [global options] {{.Name}} {{if .Flags}}[OPTIONS] {{end}}{{if ne "None" .ArgsUsage}}{{if ne "" .ArgsUsage}}{{.ArgsUsage}}{{else}}[arg...]{{end}}{{end}}

{{if .Flags}}Options:{{range .Flags}}
	 {{.}}{{end}}{{end}}
`

func main() {
	cli.AppHelpTemplate = appHelpTemplate
	cli.CommandHelpTemplate = commandHelpTemplate

	app := cli.NewApp()
	app.Name = filepath.Base(os.Args[0])
	app.Usage = "Bhojpur CLI tool for creating and managing Kubernetes clusters"
	app.Version = version.FullVersion()
	app.Before = func(ctx *cli.Context) error {
		if ctx.GlobalBool("debug") {
			logrus.SetLevel(logrus.DebugLevel)
		}
		logrus.Debugf("Bhojpur Host version: %v", app.Version)
		return nil
	}
	app.Commands = []cli.Command{
		cmd.CreateCommand(),
		cmd.UpdateCommand(),
		cmd.InspectCommand(),
		cmd.LsCommand(),
		cmd.RmCommand(),
		cmd.EnvCommand(),
		cmd.GetVersionCommand(),
		cmd.SetVersionCommand(),
		cmd.GetClusterSizeCommand(),
		cmd.SetClusterSizeCommand(),
	}
	app.CommandNotFound = cmdNotFound
	app.Author = "Bhojpur Consulting Private Limited, India"
	app.Email = "https://www.bhojpur-consulting.com"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug",
			Usage: "Enable verbose logging",
		},
		cli.StringFlag{
			Name:  "plugin-listen-addr",
			Usage: "The listening address for RPC plugin server",
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func cmdNotFound(c *cli.Context, command string) {
	log.Errorf(
		"%s: '%s' is not a %s command. See '%s --help'.",
		c.App.Name,
		command,
		c.App.Name,
		os.Args[0],
	)
	os.Exit(1)
}

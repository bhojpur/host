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
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	cmd "github.com/bhojpur/host/cmd/cluster/commands"
	"github.com/bhojpur/host/pkg/machine/log"
	"github.com/bhojpur/host/pkg/version"
	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

var released = regexp.MustCompile(`^v[0-9]+\.[0-9]+\.[0-9]+$`)

var appHelpTemplate = `Usage: {{.Name}} {{if .Flags}}[OPTIONS] {{end}}COMMAND [arg...]

{{.Usage}}

Version: {{.Version}}{{if or .Author .Email}}

Author:{{if .Author}}
  {{.Author}}{{if .Email}} - <{{.Email}}>{{end}}{{else}}
  {{.Email}}{{end}}{{end}}
{{if .Flags}}
Options:
  {{range .Flags}}{{.}}
  {{end}}{{end}}
Commands:
  {{range .Commands}}{{.Name}}{{with .ShortName}}, {{.}}{{end}}{{ "\t" }}{{.Usage}}
  {{end}}
Run '{{.Name}} COMMAND --help' for more information on a command.
`

var commandHelpTemplate = `Usage: hostfarm {{.Name}}{{if .Flags}} [OPTIONS]{{end}} [arg...]

{{.Usage}}{{if .Description}}

Description:
   {{.Description}}{{end}}{{if .Flags}}

Options:
   {{range .Flags}}
   {{.}}{{end}}{{ end }}
`

func setDebugOutputLevel() {
	// check -D, --debug and -debug, if set force debug and env var
	for _, f := range os.Args {
		if f == "-D" || f == "--debug" || f == "-debug" {
			os.Setenv("BHOJPUR_HOST_CLUSTER_DEBUG", "1")
			log.SetDebug(true)
			return
		}
	}

	// check env
	debugEnv := os.Getenv("BHOJPUR_HOST_CLUSTER_DEBUG")
	if debugEnv != "" {
		showDebug, err := strconv.ParseBool(debugEnv)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing boolean value from BHOJPUR_HOST_CLUSTER_DEBUG: %s\n", err)
			os.Exit(1)
		}
		log.SetDebug(showDebug)
	}
}

func main() {
	cli.AppHelpTemplate = appHelpTemplate
	cli.CommandHelpTemplate = commandHelpTemplate

	logrus.SetOutput(colorable.NewColorableStdout())
	setDebugOutputLevel()

	if err := mainErr(); err != nil {
		logrus.Fatal(err)
	}
}

func mainErr() error {
	app := cli.NewApp()
	app.Name = filepath.Base(os.Args[0])
	app.Author = "Bhojpur Consulting Private Limited, India"
	app.Email = "https://www.bhojpur-consulting.com"

	app.Usage = "Bhojpur Host CLI tool for creating and managing Kubernetes cluster in a Data Center"
	app.Version = version.FullVersion()

	app.Before = func(ctx *cli.Context) error {
		if ctx.GlobalBool("debug") {
			logrus.SetLevel(logrus.DebugLevel)
		}
		logrus.Debugf("Bhojpur Host cluster version: %v", app.Version)
		logrus.Debugf("This is not an officially supported version (%s) of\nthe HostFarm Provision Engine. Please download latest official release from\n\thttps://github.com/bhojpur/host/releases", app.Version)
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
	return nil
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

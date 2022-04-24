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

	"github.com/bhojpur/host/cmd/machine/commands"
	mdirs "github.com/bhojpur/host/cmd/machine/commands/dirs"
	"github.com/bhojpur/host/pkg/drivers/amazonec2"
	"github.com/bhojpur/host/pkg/drivers/azure"
	"github.com/bhojpur/host/pkg/drivers/digitalocean"
	"github.com/bhojpur/host/pkg/drivers/exoscale"
	"github.com/bhojpur/host/pkg/drivers/generic"
	"github.com/bhojpur/host/pkg/drivers/google"
	"github.com/bhojpur/host/pkg/drivers/hyperv"
	"github.com/bhojpur/host/pkg/drivers/none"
	"github.com/bhojpur/host/pkg/drivers/openstack"
	"github.com/bhojpur/host/pkg/drivers/pod"
	"github.com/bhojpur/host/pkg/drivers/rackspace"
	"github.com/bhojpur/host/pkg/drivers/softlayer"
	"github.com/bhojpur/host/pkg/drivers/virtualbox"
	"github.com/bhojpur/host/pkg/drivers/vmwarefusion"
	"github.com/bhojpur/host/pkg/drivers/vmwarevcloudair"
	"github.com/bhojpur/host/pkg/drivers/vmwarevsphere"
	"github.com/bhojpur/host/pkg/machine/drivers/plugin"
	"github.com/bhojpur/host/pkg/machine/drivers/plugin/localbinary"
	"github.com/bhojpur/host/pkg/machine/log"
	"github.com/bhojpur/host/pkg/version"
	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
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

var commandHelpTemplate = `Usage: hostutl {{.Name}}{{if .Flags}} [OPTIONS]{{end}} [arg...]

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
			os.Setenv("BHOJPUR_HOST_MACHINE_DEBUG", "1")
			log.SetDebug(true)
			return
		}
	}

	// check env
	debugEnv := os.Getenv("BHOJPUR_HOST_MACHINE_DEBUG")
	if debugEnv != "" {
		showDebug, err := strconv.ParseBool(debugEnv)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing boolean value from BHOJPUR_HOST_MACHINE_DEBUG: %s\n", err)
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
	if os.Getenv(localbinary.PluginEnvKey) == localbinary.PluginEnvVal {
		driverName := os.Getenv(localbinary.PluginEnvDriverName)
		runDriver(driverName)
		return nil
	}

	localbinary.CurrentBinaryIsBhojpurMachine = true

	app := cli.NewApp()
	app.Name = filepath.Base(os.Args[0])
	app.Author = "Bhojpur Consulting Private Limited, India"
	app.Email = "https://www.bhojpur-consulting.com"

	app.Usage = "Bhojpur Host CLI tool for creating and managing machine instances in a Data Center"
	app.Version = version.FullVersion()

	app.Before = func(ctx *cli.Context) error {
		if ctx.GlobalBool("debug") {
			logrus.SetLevel(logrus.DebugLevel)
		}
		logrus.Debugf("Bhojpur Host machine version: %v", app.Version)
		logrus.Debugf("This is not an officially supported version (%s) of\nthe Machine Provision Engine. Please download latest official release from\n\thttps://github.com/bhojpur/host/releases", app.Version)
		return nil
	}
	app.Commands = commands.Commands
	app.CommandNotFound = cmdNotFound

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, D",
			Usage: "Enable debug mode",
		},
		cli.StringFlag{
			EnvVar: "MACHINE_STORAGE_PATH",
			Name:   "storage-path, s",
			Value:  mdirs.GetBaseDir(),
			Usage:  "Configures storage path",
		},
		cli.StringFlag{
			EnvVar: "MACHINE_TLS_CA_CERT",
			Name:   "tls-ca-cert",
			Usage:  "CA to verify remotes against",
			Value:  "",
		},
		cli.StringFlag{
			EnvVar: "MACHINE_TLS_CA_KEY",
			Name:   "tls-ca-key",
			Usage:  "Private key to generate certificates",
			Value:  "",
		},
		cli.StringFlag{
			EnvVar: "MACHINE_TLS_CLIENT_CERT",
			Name:   "tls-client-cert",
			Usage:  "Client certificate to use for TLS",
			Value:  "",
		},
		cli.StringFlag{
			EnvVar: "MACHINE_TLS_CLIENT_KEY",
			Name:   "tls-client-key",
			Usage:  "Private key used in client TLS auth",
			Value:  "",
		},
		cli.StringFlag{
			EnvVar: "MACHINE_GITHUB_API_TOKEN",
			Name:   "github-api-token",
			Usage:  "Token to use for requests to the GitHub API",
			Value:  "",
		},
		cli.BoolFlag{
			EnvVar: "MACHINE_NATIVE_SSH",
			Name:   "native-ssh",
			Usage:  "Use the native (Go-based) SSH implementation.",
		},
		cli.StringFlag{
			EnvVar: "MACHINE_BUGSNAG_API_TOKEN",
			Name:   "bugsnag-api-token",
			Usage:  "BugSnag API token for crash reporting",
			Value:  "",
		},
		cli.StringFlag{
			EnvVar: "K8S_SECRET_NAME",
			Name:   "secret-name",
			Usage:  "The name of a Kubernetes secret to pull and save machine config",
			Value:  "",
		},
		cli.StringFlag{
			EnvVar: "K8S_SECRET_NAMESPACE",
			Name:   "secret-namespace",
			Usage:  "The namespace of a Kubernetes secret to pull and save machine config",
			Value:  "default",
		},
		cli.StringFlag{
			EnvVar: "KUBECONFIG",
			Name:   "kubeconfig",
			Usage:  "The path to the kubeconfig needed for secrets management",
			Value:  "",
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Error(err)
	}
	return nil
}

func runDriver(driverName string) {
	switch driverName {
	case "amazonec2":
		plugin.RegisterDriver(amazonec2.NewDriver("", ""))
	case "azure":
		plugin.RegisterDriver(azure.NewDriver("", ""))
	case "digitalocean":
		plugin.RegisterDriver(digitalocean.NewDriver("", ""))
	case "exoscale":
		plugin.RegisterDriver(exoscale.NewDriver("", ""))
	case "generic":
		plugin.RegisterDriver(generic.NewDriver("", ""))
	case "google":
		plugin.RegisterDriver(google.NewDriver("", ""))
	case "hyperv":
		plugin.RegisterDriver(hyperv.NewDriver("", ""))
	case "none":
		plugin.RegisterDriver(none.NewDriver("", ""))
	case "openstack":
		plugin.RegisterDriver(openstack.NewDriver("", ""))
	case "rackspace":
		plugin.RegisterDriver(rackspace.NewDriver("", ""))
	case "softlayer":
		plugin.RegisterDriver(softlayer.NewDriver("", ""))
	case "virtualbox":
		plugin.RegisterDriver(virtualbox.NewDriver("", ""))
	case "vmwarefusion":
		plugin.RegisterDriver(vmwarefusion.NewDriver("", ""))
	case "vmwarevcloudair":
		plugin.RegisterDriver(vmwarevcloudair.NewDriver("", ""))
	case "vmwarevsphere":
		plugin.RegisterDriver(vmwarevsphere.NewDriver("", ""))
	case "pod":
		plugin.RegisterDriver(pod.NewDriver("", ""))
	default:
		fmt.Fprintf(os.Stderr, "Unsupported Bhojpur Host machine driver: %s\n", driverName)
		os.Exit(1)
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

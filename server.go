//go:build !server
// +build !server

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

package main

import (
	"io/ioutil"
	"os"
	"regexp"

	cmd "github.com/bhojpur/host/cmd/engine/commands"
	"github.com/bhojpur/host/pkg/engine/metadata"
	"github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

// VERSION gets overridden at build time using -X main.VERSION=$VERSION
var VERSION = "dev"
var released = regexp.MustCompile(`^v[0-9]+\.[0-9]+\.[0-9]+$`)

func main() {
	logrus.SetOutput(colorable.NewColorableStdout())

	if err := mainErr(); err != nil {
		logrus.Fatal(err)
	}
}

func mainErr() error {
	app := cli.NewApp()
	app.Name = "hostsvr"
	app.Version = VERSION
	app.Usage = "Bhojpur CLI tool for installing fast Kubernetes Engine that works everywhere"
	app.Before = func(ctx *cli.Context) error {
		if ctx.GlobalBool("quiet") {
			logrus.SetOutput(ioutil.Discard)
		} else {
			if ctx.GlobalBool("debug") {
				logrus.SetLevel(logrus.DebugLevel)
				logrus.Debugf("Loglevel set to [%v]", logrus.DebugLevel)
			}
			if ctx.GlobalBool("trace") {
				logrus.SetLevel(logrus.TraceLevel)
				logrus.Tracef("Loglevel set to [%v]", logrus.TraceLevel)
			}
		}
		if released.MatchString(app.Version) {
			metadata.BKEVersion = app.Version
			return nil
		}
		logrus.Warnf("This is not an officially supported version (%s) of Bhojpur Kubernetes Engine. Please download the latest official release at https://github.com/bhojpur/host/releases", app.Version)
		return nil
	}
	app.Author = "Bhojpur Consulting Private Limited, India."
	app.Email = "https://www.bhojpur-consulting.com"
	app.Commands = []cli.Command{
		cmd.UpCommand(),
		cmd.RemoveCommand(),
		cmd.VersionCommand(),
		cmd.ConfigCommand(),
		cmd.EtcdCommand(),
		cmd.CertificateCommand(),
		cmd.EncryptionCommand(),
		cmd.UtilCommand(),
	}
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug,d",
			Usage: "Debug logging",
		},
		cli.BoolFlag{
			Name:  "quiet,q",
			Usage: "Quiet mode, disables logging and only critical output will be printed",
		},
		cli.BoolFlag{
			Name:  "trace",
			Usage: "Trace logging",
		},
	}
	return app.Run(os.Args)
}

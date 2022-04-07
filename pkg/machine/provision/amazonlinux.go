package provision

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
	"bytes"
	"fmt"
	"text/template"

	"github.com/bhojpur/host/pkg/machine/auth"
	"github.com/bhojpur/host/pkg/machine/drivers"
	"github.com/bhojpur/host/pkg/machine/engine"
	"github.com/bhojpur/host/pkg/machine/log"
	"github.com/bhojpur/host/pkg/machine/provision/pkgaction"
	"github.com/bhojpur/host/pkg/machine/provision/serviceaction"
	"github.com/bhojpur/host/pkg/machine/swarm"
	mutils "github.com/bhojpur/host/pkg/machine/utils"
)

func init() {
	Register("amzn", &RegisteredProvisioner{
		New: NewAmazonLinuxProvisioner,
	})
}

func NewAmazonLinuxProvisioner(d drivers.Driver) Provisioner {
	return &AmazonLinuxProvisioner{
		NewSystemdProvisioner("amzn", d),
	}
}

type AmazonLinuxProvisioner struct {
	SystemdProvisioner
}

func (provisioner *AmazonLinuxProvisioner) String() string {
	return "amzn"
}

func (provisioner *AmazonLinuxProvisioner) Package(name string, action pkgaction.PackageAction) error {
	var packageAction string

	switch action {
	case pkgaction.Install:
		packageAction = "install"
	case pkgaction.Remove:
		packageAction = "remove"
	case pkgaction.Purge:
		packageAction = "remove"
	case pkgaction.Upgrade:
		packageAction = "upgrade"
	}

	command := fmt.Sprintf("sudo -E yum %s -y %s", packageAction, name)

	if _, err := provisioner.SSHCommand(command); err != nil {
		return err
	}

	return nil
}

func (provisioner *AmazonLinuxProvisioner) bhojpurDaemonResponding() bool {
	log.Debug("checking Bhojpur Host daemon")

	if out, err := provisioner.SSHCommand("sudo hostutl version"); err != nil {
		log.Warnf("Error getting SSH command to check if the daemon is up: %s", err)
		log.Debugf("'sudo hostutl version' output:\n%s", out)
		return false
	}

	// The daemon is up if the command worked.  Carry on.
	return true
}

func (provisioner *AmazonLinuxProvisioner) Provision(swarmOptions swarm.Options, authOptions auth.Options, engineOptions engine.Options) error {
	provisioner.SwarmOptions = swarmOptions
	provisioner.AuthOptions = authOptions
	provisioner.EngineOptions = engineOptions
	swarmOptions.Env = engineOptions.Env

	// set default storage driver for redhat
	storageDriver, err := decideStorageDriver(provisioner, DefaultStorageDriver, engineOptions.StorageDriver)
	if err != nil {
		return err
	}
	provisioner.EngineOptions.StorageDriver = storageDriver

	if err := provisioner.SetHostname(provisioner.Driver.GetMachineName()); err != nil {
		return err
	}

	for _, pkg := range provisioner.Packages {
		log.Debugf("installing base package: name=%s", pkg)
		if err := provisioner.Package(pkg, pkgaction.Install); err != nil {
			return err
		}
	}

	log.Debug("Installing Bhojpur Host")
	if err := provisioner.Package("bhojpur", pkgaction.Install); err != nil {
		return err
	}

	log.Debug("Starting systemd hostutl service")
	if err := provisioner.Service("bhojpur", serviceaction.Start); err != nil {
		return err
	}

	log.Debug("Waiting for Bhojpur Host daemon")
	if err := mutils.WaitFor(provisioner.bhojpurDaemonResponding); err != nil {
		return err
	}

	if err := makeBhojpurOptionsDir(provisioner); err != nil {
		return err
	}

	provisioner.AuthOptions = setRemoteAuthOptions(provisioner)

	if err := ConfigureAuth(provisioner); err != nil {
		return err
	}

	err = configureSwarm(provisioner, swarmOptions, provisioner.AuthOptions)
	return err
}

func (provisioner *AmazonLinuxProvisioner) GenerateBhojpurOptions(bhojpurPort int) (*BhojpurOptions, error) {
	var (
		engineCfg  bytes.Buffer
		configPath = provisioner.DaemonOptionsFile
	)

	driverNameLabel := fmt.Sprintf("provider=%s", provisioner.Driver.DriverName())
	provisioner.EngineOptions.Labels = append(provisioner.EngineOptions.Labels, driverNameLabel)

	// systemd / redhat will not load options if they are on newlines
	// instead, it just continues with a different set of options; yeah...
	t, err := template.New("engineConfig").Parse(engineConfigTemplate)
	if err != nil {
		return nil, err
	}

	engineConfigContext := EngineConfigContext{
		BhojpurPort:       bhojpurPort,
		AuthOptions:       provisioner.AuthOptions,
		EngineOptions:     provisioner.EngineOptions,
		BhojpurOptionsDir: provisioner.BhojpurOptionsDir,
	}

	t.Execute(&engineCfg, engineConfigContext)

	daemonOptsDir := configPath
	return &BhojpurOptions{
		EngineOptions:     engineCfg.String(),
		EngineOptionsPath: daemonOptsDir,
	}, nil
}

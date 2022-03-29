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
	"fmt"

	"github.com/bhojpur/host/pkg/core/auth"
	"github.com/bhojpur/host/pkg/core/drivers"
	"github.com/bhojpur/host/pkg/core/engine"
	"github.com/bhojpur/host/pkg/core/log"
	"github.com/bhojpur/host/pkg/core/provision/pkgaction"
	"github.com/bhojpur/host/pkg/core/provision/serviceaction"
	"github.com/bhojpur/host/pkg/core/swarm"
	mutils "github.com/bhojpur/host/pkg/core/utils"
)

func init() {
	Register("Debian", &RegisteredProvisioner{
		New: NewDebianProvisioner,
	})
}

func NewDebianProvisioner(d drivers.Driver) Provisioner {
	return &DebianProvisioner{
		NewSystemdProvisioner("debian", d),
	}
}

type DebianProvisioner struct {
	SystemdProvisioner
}

func (provisioner *DebianProvisioner) String() string {
	return "debian"
}

func (provisioner *DebianProvisioner) Package(name string, action pkgaction.PackageAction) error {
	var packageAction string

	updateMetadata := true

	switch action {
	case pkgaction.Install, pkgaction.Upgrade:
		packageAction = "install"
	case pkgaction.Remove:
		packageAction = "remove"
		updateMetadata = false
	case pkgaction.Purge:
		packageAction = "purge"
		updateMetadata = false
	}

	switch name {
	case "bhojpur":
		name = "bhojpur-engine"
	}

	if updateMetadata {
		if err := waitForLockAptGetUpdate(provisioner); err != nil {
			return err
		}
	}

	command := fmt.Sprintf("DEBIAN_FRONTEND=noninteractive sudo -E apt-get %s -y  %s", packageAction, name)

	log.Debugf("package: action=%s name=%s", action.String(), name)

	return waitForLock(provisioner, command)
}

func (provisioner *DebianProvisioner) bhojpurDaemonResponding() bool {
	log.Debug("checking Bhojpur Host daemon")

	if out, err := provisioner.SSHCommand("sudo hostutl version"); err != nil {
		log.Warnf("Error getting SSH command to check if the daemon is up: %s", err)
		log.Debugf("'sudo hostutl version' output:\n%s", out)
		return false
	}

	// The daemon is up if the command worked.  Carry on.
	return true
}

func (provisioner *DebianProvisioner) Provision(swarmOptions swarm.Options, authOptions auth.Options, engineOptions engine.Options) error {
	provisioner.SwarmOptions = swarmOptions
	provisioner.AuthOptions = authOptions
	provisioner.EngineOptions = engineOptions
	swarmOptions.Env = engineOptions.Env

	storageDriver, err := decideStorageDriver(provisioner, DefaultStorageDriver, engineOptions.StorageDriver)
	if err != nil {
		return err
	}
	provisioner.EngineOptions.StorageDriver = storageDriver

	// HACK: since debian does not come with sudo by default we install
	log.Debug("installing sudo")
	if _, err := provisioner.SSHCommand("if ! type sudo; then apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y sudo; fi"); err != nil {
		return err
	}

	log.Debug("setting hostname")
	if err := provisioner.SetHostname(provisioner.Driver.GetMachineName()); err != nil {
		return err
	}

	log.Debug("installing base packages")
	for _, pkg := range provisioner.Packages {
		if err := provisioner.Package(pkg, pkgaction.Install); err != nil {
			return err
		}
	}

	if err := installBhojpurGeneric(provisioner, provisioner.EngineOptions.InstallURL); err != nil {
		return err
	}

	log.Debug("waiting for Bhojpur Host daemon")
	if err := mutils.WaitFor(provisioner.bhojpurDaemonResponding); err != nil {
		return err
	}

	provisioner.AuthOptions = setRemoteAuthOptions(provisioner)

	log.Debug("configuring auth")
	if err := ConfigureAuth(provisioner); err != nil {
		return err
	}

	log.Debug("configuring swarm")
	if err := configureSwarm(provisioner, swarmOptions, provisioner.AuthOptions); err != nil {
		return err
	}

	// enable in systemd
	log.Debug("enabling Bhojpur Host in systemd")
	err = provisioner.Service("bhojpur", serviceaction.Enable)
	return err
}

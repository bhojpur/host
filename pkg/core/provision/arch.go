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
	Register("Arch", &RegisteredProvisioner{
		New: NewArchProvisioner,
	})
}

func NewArchProvisioner(d drivers.Driver) Provisioner {
	return &ArchProvisioner{
		NewSystemdProvisioner("arch", d),
	}
}

type ArchProvisioner struct {
	SystemdProvisioner
}

func (provisioner *ArchProvisioner) String() string {
	return "arch"
}

func (provisioner *ArchProvisioner) CompatibleWithHost() bool {
	return provisioner.OsReleaseInfo.ID == provisioner.OsReleaseID || provisioner.OsReleaseInfo.IDLike == provisioner.OsReleaseID
}

func (provisioner *ArchProvisioner) Package(name string, action pkgaction.PackageAction) error {
	var packageAction string

	updateMetadata := true

	switch action {
	case pkgaction.Install, pkgaction.Upgrade:
		packageAction = "S"
	case pkgaction.Remove:
		packageAction = "R"
		updateMetadata = false
	}

	switch name {
	case "bhojpur-engine":
		name = "bhojpur"
	case "bhojpur":
		name = "bhojpur"
	}

	pacmanOpts := "-" + packageAction
	if updateMetadata {
		pacmanOpts = pacmanOpts + "y"
	}

	pacmanOpts = pacmanOpts + " --noconfirm --noprogressbar"

	command := fmt.Sprintf("sudo -E pacman %s %s", pacmanOpts, name)

	log.Debugf("package: action=%s name=%s", action.String(), name)

	if _, err := provisioner.SSHCommand(command); err != nil {
		return err
	}

	return nil
}

func (provisioner *ArchProvisioner) bhojpurDaemonResponding() bool {
	log.Debug("checking Bhojpur Host daemon")

	if out, err := provisioner.SSHCommand("sudo hostutl version"); err != nil {
		log.Warnf("Error getting SSH command to check if the daemon is up: %s", err)
		log.Debugf("'sudo hostutl version' output:\n%s", out)
		return false
	}

	// The daemon is up if the command worked.  Carry on.
	return true
}

func (provisioner *ArchProvisioner) Provision(swarmOptions swarm.Options, authOptions auth.Options, engineOptions engine.Options) error {
	provisioner.SwarmOptions = swarmOptions
	provisioner.AuthOptions = authOptions
	provisioner.EngineOptions = engineOptions
	swarmOptions.Env = engineOptions.Env

	storageDriver, err := decideStorageDriver(provisioner, "overlay", engineOptions.StorageDriver)
	if err != nil {
		return err
	}
	provisioner.EngineOptions.StorageDriver = storageDriver

	// HACK: since Arch does not come with sudo by default we install
	log.Debug("Installing sudo")
	if _, err := provisioner.SSHCommand("if ! type sudo; then pacman -Sy --noconfirm --noprogressbar sudo; fi"); err != nil {
		return err
	}

	log.Debug("Setting hostname")
	if err := provisioner.SetHostname(provisioner.Driver.GetMachineName()); err != nil {
		return err
	}

	log.Debug("Installing base packages")
	for _, pkg := range provisioner.Packages {
		if err := provisioner.Package(pkg, pkgaction.Install); err != nil {
			return err
		}
	}

	log.Debug("Installing Bhojpur Host")
	if err := provisioner.Package("bhojpur", pkgaction.Install); err != nil {
		return err
	}

	log.Debug("Starting systemd bhojpur service")
	if err := provisioner.Service("bhojpur", serviceaction.Start); err != nil {
		return err
	}

	log.Debug("Waiting for Bhojpur Host daemon")
	if err := mutils.WaitFor(provisioner.bhojpurDaemonResponding); err != nil {
		return err
	}

	provisioner.AuthOptions = setRemoteAuthOptions(provisioner)

	log.Debug("Configuring auth")
	if err := ConfigureAuth(provisioner); err != nil {
		return err
	}

	log.Debug("Configuring swarm")
	if err := configureSwarm(provisioner, swarmOptions, provisioner.AuthOptions); err != nil {
		return err
	}

	// enable in systemd
	log.Debug("Enabling Bhojpur Host in systemd")
	err = provisioner.Service("bhojpur", serviceaction.Enable)
	return err
}

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
	"bufio"
	"fmt"
	"net/http"
	"strings"

	mdirs "github.com/bhojpur/host/cmd/machine/commands/dirs"
	"github.com/bhojpur/host/pkg/core/auth"
	"github.com/bhojpur/host/pkg/core/drivers"
	"github.com/bhojpur/host/pkg/core/engine"
	"github.com/bhojpur/host/pkg/core/log"
	"github.com/bhojpur/host/pkg/core/provision/pkgaction"
	"github.com/bhojpur/host/pkg/core/provision/serviceaction"
	"github.com/bhojpur/host/pkg/core/state"
	"github.com/bhojpur/host/pkg/core/swarm"
	mutils "github.com/bhojpur/host/pkg/core/utils"
)

const (
	versionsURL  = "https://releases.rancher.com/os/releases.yml"
	isoURL       = "https://releases.rancher.com/os/%s/rancheros.iso"
	hostnameTmpl = `sudo mkdir -p /var/lib/rancher/conf/cloud-config.d/
sudo tee /var/lib/rancher/conf/cloud-config.d/machine-hostname.yml << EOF
#cloud-config

hostname: %s
EOF
`
)

func init() {
	Register("RancherOS", &RegisteredProvisioner{
		New: NewRancherProvisioner,
	})
}

func NewRancherProvisioner(d drivers.Driver) Provisioner {
	return &RancherProvisioner{
		GenericProvisioner{
			SSHCommander:      GenericSSHCommander{Driver: d},
			BhojpurOptionsDir: "/var/lib/rancher/conf",
			DaemonOptionsFile: "/var/lib/rancher/conf/bhojpur",
			OsReleaseID:       "rancheros",
			Driver:            d,
		},
	}
}

type RancherProvisioner struct {
	GenericProvisioner
}

func (provisioner *RancherProvisioner) String() string {
	return "rancheros"
}

func (provisioner *RancherProvisioner) Service(name string, action serviceaction.ServiceAction) error {
	command := fmt.Sprintf("sudo system-bhojpur %s %s", action.String(), name)

	if _, err := provisioner.SSHCommand(command); err != nil {
		return err
	}

	return nil
}

func (provisioner *RancherProvisioner) Package(name string, action pkgaction.PackageAction) error {
	var packageAction string

	if name == "bhojpur" && action == pkgaction.Upgrade {
		return provisioner.upgrade()
	}

	switch action {
	case pkgaction.Install:
		packageAction = "enable"
	case pkgaction.Remove:
		packageAction = "disable"
	case pkgaction.Upgrade:
		// TODO: support upgrade
		packageAction = "upgrade"
	}

	command := fmt.Sprintf("sudo ros service %s %s", packageAction, name)

	if _, err := provisioner.SSHCommand(command); err != nil {
		return err
	}

	return nil
}

func (provisioner *RancherProvisioner) Provision(swarmOptions swarm.Options, authOptions auth.Options, engineOptions engine.Options) error {
	log.Debugf("Running RancherOS provisioner on %s", provisioner.Driver.GetMachineName())

	provisioner.SwarmOptions = swarmOptions
	provisioner.AuthOptions = authOptions
	provisioner.EngineOptions = engineOptions
	swarmOptions.Env = engineOptions.Env

	if provisioner.EngineOptions.StorageDriver == "" {
		provisioner.EngineOptions.StorageDriver = DefaultStorageDriver
	} else if provisioner.EngineOptions.StorageDriver != "overlay" && provisioner.EngineOptions.StorageDriver != "overlay2" {
		return fmt.Errorf("Unsupported storage driver: %s", provisioner.EngineOptions.StorageDriver)
	}

	log.Debugf("Setting hostname %s", provisioner.Driver.GetMachineName())
	if err := provisioner.SetHostname(provisioner.Driver.GetMachineName()); err != nil {
		return err
	}

	for _, pkg := range provisioner.Packages {
		log.Debugf("Installing package %s", pkg)
		if err := provisioner.Package(pkg, pkgaction.Install); err != nil {
			return err
		}
	}

	if engineOptions.InstallURL == drivers.DefaultEngineInstallURL {
		log.Debugf("Skipping Bhojpur Host engine default: %s", engineOptions.InstallURL)
	} else {
		log.Debugf("Selecting Bhojpur Host engine: %s", engineOptions.InstallURL)
		if err := selectBhojpurHost(provisioner, engineOptions.InstallURL); err != nil {
			return err
		}
	}

	log.Debugf("Preparing certificates")
	provisioner.AuthOptions = setRemoteAuthOptions(provisioner)

	log.Debugf("Setting up certificates")
	if err := ConfigureAuth(provisioner); err != nil {
		return err
	}

	log.Debugf("Configuring swarm")
	err := configureSwarm(provisioner, swarmOptions, provisioner.AuthOptions)
	return err
}

func (provisioner *RancherProvisioner) SetHostname(hostname string) error {
	// /etc/hosts is bind mounted from Bhojpur Host, this is hack to that the generic provisioner doesn't try to mv /etc/hosts
	if _, err := provisioner.SSHCommand("sed /127.0.1.1/d /etc/hosts > /tmp/hosts && cat /tmp/hosts | sudo tee /etc/hosts"); err != nil {
		return err
	}

	if err := provisioner.GenericProvisioner.SetHostname(hostname); err != nil {
		return err
	}

	if _, err := provisioner.SSHCommand(fmt.Sprintf(hostnameTmpl, hostname)); err != nil {
		return err
	}

	return nil
}

func (provisioner *RancherProvisioner) upgrade() error {
	switch provisioner.Driver.DriverName() {
	case "virtualbox":
		return provisioner.upgradeIso()
	default:
		log.Infof("Running upgrade")
		if _, err := provisioner.SSHCommand("sudo ros os upgrade -f --no-reboot"); err != nil {
			return err
		}

		log.Infof("Upgrade succeeded, rebooting")
		// ignore errors here because the SSH connection will close
		provisioner.SSHCommand("sudo reboot")

		return nil
	}
}

func (provisioner *RancherProvisioner) upgradeIso() error {
	// Largely copied from Boot2Docker provisioner, we should find a way to share this code
	log.Info("Stopping machine to do the upgrade...")

	if err := provisioner.Driver.Stop(); err != nil {
		return err
	}

	if err := mutils.WaitFor(drivers.MachineInState(provisioner.Driver, state.Stopped)); err != nil {
		return err
	}

	machineName := provisioner.GetDriver().GetMachineName()

	log.Infof("Upgrading machine %s...", machineName)

	// TODO: Ideally, we should not read from mcndirs directory at all.
	// The driver should be able to communicate how and where to place the
	// relevant files.
	b2dutils := mutils.NewB2dUtils(mdirs.GetBaseDir())

	url, err := provisioner.getLatestISOURL()
	if err != nil {
		return err
	}

	//if err := b2dutils.DownloadISOFromURL(url); err != nil {
	//	return err
	//}

	// Copy the latest version of rancheros ISO to the machine's directory
	if err := b2dutils.CopyIsoToMachineDir(url, machineName); err != nil {
		return err
	}

	log.Infof("Starting machine back up...")

	if err := provisioner.Driver.Start(); err != nil {
		return err
	}

	return mutils.WaitFor(drivers.MachineInState(provisioner.Driver, state.Running))
}

func (provisioner *RancherProvisioner) getLatestISOURL() (string, error) {
	log.Debugf("Reading %s", versionsURL)
	resp, err := http.Get(versionsURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Don't want to pull in yaml parser, we'll do this manually
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "current: ") {
			log.Debugf("Found %s", line)
			return fmt.Sprintf(isoURL, strings.Split(line, ":")[2]), err
		}
	}

	return "", fmt.Errorf("Failed to find current version")
}

func selectBhojpurHost(p Provisioner, baseURL string) error {
	// TODO: detect if its a cloud-init, or a ros setting - and use that..
	if output, err := p.SSHCommand(fmt.Sprintf("wget -O- %s | sh -", baseURL)); err != nil {
		return fmt.Errorf("error selecting bhojpur host: (%s) %s", err, output)
	}

	return nil
}

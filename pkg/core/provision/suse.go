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
	"strings"

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
	Register("openSUSE", &RegisteredProvisioner{
		New: NewOpenSUSEProvisioner,
	})
	Register("SUSE Linux Enterprise Desktop", &RegisteredProvisioner{
		New: NewSLEDProvisioner,
	})
	Register("SUSE Linux Enterprise Server", &RegisteredProvisioner{
		New: NewSLESProvisioner,
	})
}

func NewSLEDProvisioner(d drivers.Driver) Provisioner {
	return &SUSEProvisioner{
		NewSystemdProvisioner("sled", d),
	}
}

func NewSLESProvisioner(d drivers.Driver) Provisioner {
	return &SUSEProvisioner{
		NewSystemdProvisioner("sles", d),
	}
}

func NewOpenSUSEProvisioner(d drivers.Driver) Provisioner {
	return &SUSEProvisioner{
		NewSystemdProvisioner("opensuse", d),
	}
}

type SUSEProvisioner struct {
	SystemdProvisioner
}

func (provisioner *SUSEProvisioner) CompatibleWithHost() bool {
	return strings.ToLower(provisioner.OsReleaseInfo.ID) == strings.ToLower(provisioner.OsReleaseID) || strings.Contains(provisioner.OsReleaseInfo.IDLike, "opensuse")
}

func (provisioner *SUSEProvisioner) String() string {
	return "openSUSE"
}

func (provisioner *SUSEProvisioner) Package(name string, action pkgaction.PackageAction) error {
	var packageAction string

	switch action {
	case pkgaction.Install:
		packageAction = "in"
		// This is an optimization that reduces the provisioning time of certain
		// systems in a significant way.
		// The invocation of "zypper in <pkg>" causes the download of the metadata
		// of all the repositories that have never been refreshed or that have
		// automatic refresh toggled and have not been refreshed recently.
		// Refreshing the repository metadata can take quite some time and can cause
		// longer provisioning times for machines that have been pre-optimized for
		// Bhojpur Host by including all the needed packages.
		if _, err := provisioner.SSHCommand(fmt.Sprintf("rpm -q %s", name)); err == nil {
			log.Debugf("%s is already installed, skipping operation", name)
			return nil
		}
	case pkgaction.Remove:
		packageAction = "rm"
	case pkgaction.Upgrade:
		packageAction = "up"
	}

	command := fmt.Sprintf("sudo -E zypper -n %s %s", packageAction, name)

	log.Debugf("zypper: action=%s name=%s", action.String(), name)

	if _, err := provisioner.SSHCommand(command); err != nil {
		return err
	}

	return nil
}

func (provisioner *SUSEProvisioner) bhojpurDaemonResponding() bool {
	log.Debug("checking Bhojpur Host daemon")

	if out, err := provisioner.SSHCommand("sudo hostutl version"); err != nil {
		log.Warnf("Error getting SSH command to check if the daemon is up: %s", err)
		log.Debugf("'sudo hostutl version' output:\n%s", out)
		return false
	}

	// The daemon is up if the command worked. Carry on.
	return true
}

func (provisioner *SUSEProvisioner) Provision(swarmOptions swarm.Options, authOptions auth.Options, engineOptions engine.Options) error {
	provisioner.SwarmOptions = swarmOptions
	provisioner.AuthOptions = authOptions
	provisioner.EngineOptions = engineOptions
	swarmOptions.Env = engineOptions.Env

	// figure out the filesystem used by /var/lib/hostutl
	fs, err := provisioner.SSHCommand("stat -f -c %T /var/lib/hostutl")
	if err != nil {
		// figure out the filesystem used by /var/lib
		fs, err = provisioner.SSHCommand("stat -f -c %T /var/lib/")
		if err != nil {
			return err
		}
	}
	graphDriver := "overlay"
	if strings.Contains(fs, "btrfs") {
		graphDriver = "btrfs"
	}

	storageDriver, err := decideStorageDriver(provisioner, graphDriver, engineOptions.StorageDriver)
	if err != nil {
		return err
	}
	provisioner.EngineOptions.StorageDriver = storageDriver

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

	if err := installBhojpurGeneric(provisioner, provisioner.EngineOptions.InstallURL); err != nil {
		return err
	}

	// Is yast2 firewall installed?
	if _, installed := provisioner.SSHCommand("rpm -q yast2-firewall"); installed == nil {
		log.Debug("Configuring SUSE firewall")
		if err := provisioner.configureFirewall(); err != nil {
			return err
		}
	}

	log.Debug("Starting systemd Bhojpur Host service")
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

func (provisioner *SUSEProvisioner) configureFirewall() error {
	tcpPorts := "22 80 443 2376 2379 2380 6443 9099 9796 10250 10254 30000:32767"
	udpPorts := "8472 30000:32767"
	var cmds []string
	// SLES15 has dropped SuSEfirewall2: https://www.suse.com/releasenotes/x86_64/SUSE-SLES/15/#fate-320794
	// Change to logic to allow handling the same //
	if _, installed := provisioner.SSHCommand("rpm -q SuSEfirewall2"); installed == nil {
		cmds = []string{
			fmt.Sprintf(`sudo sed -i 's/FW_SERVICES_EXT_TCP=".*"/FW_SERVICES_EXT_TCP="%s"/' /etc/sysconfig/SuSEfirewall2`, tcpPorts),
			fmt.Sprintf(`sudo sed -i 's/FW_SERVICES_EXT_UDP=".*"/FW_SERVICES_EXT_UDP="%s"/' /etc/sysconfig/SuSEfirewall2`, udpPorts),
			"sudo /sbin/SuSEfirewall2",
		}
	}

	// firewall is an optional package unlike SuSEfirewall2
	if _, installed := provisioner.SSHCommand("rpm -q firewalld"); installed == nil {
		// check if firewalld is running //
		if _, running := provisioner.SSHCommand("sudo firewall-cmd --stat"); running == nil {
			// edge case where someone may have installed both, we default to the running firewalld
			if len(cmds) != 0 {
				cmds = nil
			}
			tcpPorts := strings.ReplaceAll(tcpPorts, ":", "-")
			udpPorts := strings.ReplaceAll(udpPorts, ":", "-")
			for _, port := range strings.Split(tcpPorts, " ") {
				cmds = append(cmds, fmt.Sprintf("sudo firewall-cmd --permanent --add-port=%s/tcp", port))
			}
			for _, port := range strings.Split(udpPorts, " ") {
				cmds = append(cmds, fmt.Sprintf("sudo firewall-cmd --permanent --add-port=%s/udp", port))
			}
			cmds = append(cmds, "sudo firewall-cmd --reload")
		}

	}

	for _, cmd := range cmds {
		if _, err := provisioner.SSHCommand(cmd); err != nil {
			return err
		}
	}
	return nil
}

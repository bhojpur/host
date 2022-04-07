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

	"github.com/bhojpur/host/pkg/machine/auth"
	"github.com/bhojpur/host/pkg/machine/drivers"
	"github.com/bhojpur/host/pkg/machine/engine"
	"github.com/bhojpur/host/pkg/machine/log"
	"github.com/bhojpur/host/pkg/machine/provision/pkgaction"
	"github.com/bhojpur/host/pkg/machine/provision/serviceaction"
	"github.com/bhojpur/host/pkg/machine/swarm"
)

var (
	provisioners          = make(map[string]*RegisteredProvisioner)
	detector     Detector = &StandardDetector{}
)

const (
	LastReleaseBeforeCEVersioning = "1.13.1"
	DefaultStorageDriver          = "overlay2"
)

type SSHCommander interface {
	// Short-hand for accessing an SSH command from the driver.
	SSHCommand(args string) (string, error)
}

type Detector interface {
	DetectProvisioner(d drivers.Driver) (Provisioner, error)
}

type StandardDetector struct{}

func SetDetector(newDetector Detector) {
	detector = newDetector
}

// Provisioner defines distribution specific actions
type Provisioner interface {
	fmt.Stringer
	SSHCommander

	// Create the files for the daemon to consume configuration settings (return struct of content and path)
	GenerateBhojpurOptions(bhojpurPort int) (*BhojpurOptions, error)

	// Get the directory where the settings files for Bhojpur Host are to be found
	GetBhojpurOptionsDir() string

	// Return the auth options used to configure remote connection for the daemon.
	GetAuthOptions() auth.Options

	// Get the swarm options associated with this host.
	GetSwarmOptions() swarm.Options

	// Run a package action e.g. install
	Package(name string, action pkgaction.PackageAction) error

	// Get the pakcages to install
	GetPackages() []string

	// Get Hostname
	Hostname() (string, error)

	// Set hostname
	SetHostname(hostname string) error

	// Figure out if this is the right provisioner to use based on /etc/os-release info
	CompatibleWithHost() bool

	// Do the actual provisioning piece:
	//     1. Set the hostname on the instance.
	//     2. Install Bhojpur Host if it is not present.
	//     3. Configure the daemon to accept connections over TLS.
	//     4. Copy the needed certificates to the server and local config dir.
	//     5. Configure / activate swarm if applicable.
	Provision(swarmOptions swarm.Options, authOptions auth.Options, engineOptions engine.Options) error

	// Perform action on a named service e.g. stop
	Service(name string, action serviceaction.ServiceAction) error

	// Get the driver which is contained in the provisioner.
	GetDriver() drivers.Driver

	// Set the OS Release info depending on how it's represented
	// internally
	SetOsReleaseInfo(info *OsRelease)

	// Get the OS Release info for the current provisioner
	GetOsReleaseInfo() (*OsRelease, error)
}

// RegisteredProvisioner creates a new provisioner
type RegisteredProvisioner struct {
	New func(d drivers.Driver) Provisioner
}

func Register(name string, p *RegisteredProvisioner) {
	provisioners[name] = p
}

func DetectProvisioner(d drivers.Driver) (Provisioner, error) {
	return detector.DetectProvisioner(d)
}

func (detector StandardDetector) DetectProvisioner(d drivers.Driver) (Provisioner, error) {
	log.Info("Waiting for SSH to be available...")
	if err := drivers.WaitForSSH(d); err != nil {
		return nil, err
	}

	log.Info("Detecting the Bhojpur Host provisioner...")

	osReleaseOut, err := drivers.RunSSHCommandFromDriver(d, "cat /etc/os-release")
	if err != nil {
		return nil, fmt.Errorf("Error getting SSH command: %s", err)
	}

	osReleaseInfo, err := NewOsRelease([]byte(osReleaseOut))
	if err != nil {
		return nil, fmt.Errorf("Error parsing /etc/os-release file: %s", err)
	}

	for _, p := range provisioners {
		provisioner := p.New(d)
		provisioner.SetOsReleaseInfo(osReleaseInfo)

		if provisioner.CompatibleWithHost() {
			log.Debugf("found compatible Bhojpur Host: %s", osReleaseInfo.ID)
			return provisioner, nil
		}
	}

	return nil, ErrDetectionFailed
}

package host

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
	"regexp"

	"github.com/bhojpur/host/pkg/machine/auth"
	"github.com/bhojpur/host/pkg/machine/cert"
	cengine "github.com/bhojpur/host/pkg/machine/client"
	"github.com/bhojpur/host/pkg/machine/drivers"
	"github.com/bhojpur/host/pkg/machine/engine"
	merrors "github.com/bhojpur/host/pkg/machine/errors"
	"github.com/bhojpur/host/pkg/machine/log"
	"github.com/bhojpur/host/pkg/machine/provision"
	"github.com/bhojpur/host/pkg/machine/provision/pkgaction"
	"github.com/bhojpur/host/pkg/machine/provision/serviceaction"
	"github.com/bhojpur/host/pkg/machine/ssh"
	"github.com/bhojpur/host/pkg/machine/state"
	"github.com/bhojpur/host/pkg/machine/swarm"
	mutils "github.com/bhojpur/host/pkg/machine/utils"
	"github.com/bhojpur/host/pkg/machine/versioncmp"
)

const noBhojpurError = "Bhojpur Host was not provisioned on machine %s, %s"

var (
	validHostNamePattern                  = regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9\-\.]*$`)
	stdSSHClientCreator  SSHClientCreator = &StandardSSHClientCreator{}
)

type SSHClientCreator interface {
	CreateSSHClient(d drivers.Driver) (ssh.Client, error)
}

type StandardSSHClientCreator struct {
	drivers.Driver
}

func SetSSHClientCreator(creator SSHClientCreator) {
	stdSSHClientCreator = creator
}

type Host struct {
	ConfigVersion int
	Driver        drivers.Driver
	DriverName    string
	HostOptions   *Options
	Name          string
	RawDriver     []byte `json:"-"`
}

type Options struct {
	Driver              string
	Memory              int
	Disk                int
	CustomInstallScript string
	MachineOS           string
	EngineOptions       *engine.Options
	SwarmOptions        *swarm.Options
	AuthOptions         *auth.Options
}

type Metadata struct {
	ConfigVersion int
	DriverName    string
	HostOptions   Options
}

func ValidateHostName(name string) bool {
	return validHostNamePattern.MatchString(name)
}

func (h *Host) RunSSHCommand(command string) (string, error) {
	return drivers.RunSSHCommandFromDriver(h.Driver, command)
}

func (h *Host) CreateSSHClient() (ssh.Client, error) {
	return stdSSHClientCreator.CreateSSHClient(h.Driver)
}

func (creator *StandardSSHClientCreator) CreateSSHClient(d drivers.Driver) (ssh.Client, error) {
	addr, err := d.GetSSHHostname()
	if err != nil {
		return &ssh.ExternalClient{}, err
	}

	port, err := d.GetSSHPort()
	if err != nil {
		return &ssh.ExternalClient{}, err
	}

	auth := &ssh.Auth{}
	if d.GetSSHKeyPath() != "" {
		auth.Keys = []string{d.GetSSHKeyPath()}
	}

	return ssh.NewClient(d.GetSSHUsername(), addr, port, auth)
}

func (h *Host) runActionForState(action func() error, desiredState state.State) error {
	if drivers.MachineInState(h.Driver, desiredState)() {
		return merrors.ErrHostAlreadyInState{
			Name:  h.Name,
			State: desiredState,
		}
	}

	if err := action(); err != nil {
		return err
	}

	return mutils.WaitFor(drivers.MachineInState(h.Driver, desiredState))
}

func (h *Host) WaitForBhojpurHost() error {
	provisioner, err := provision.DetectProvisioner(h.Driver)
	if err != nil {
		return err
	}

	return provision.WaitForBhojpur(provisioner, engine.DefaultPort)
}

func (h *Host) Start() error {
	log.Infof("Starting %q...", h.Name)
	if err := h.runActionForState(h.Driver.Start, state.Running); err != nil {
		return err
	}

	log.Infof("Bhojpur Host machine %q was started.", h.Name)

	return h.WaitForBhojpurHost()
}

func (h *Host) Stop() error {
	log.Infof("Stopping %q...", h.Name)
	if err := h.runActionForState(h.Driver.Stop, state.Stopped); err != nil {
		return err
	}

	log.Infof("Bhojpur Host machine %q was stopped.", h.Name)
	return nil
}

func (h *Host) Kill() error {
	log.Infof("Killing %q...", h.Name)
	if err := h.runActionForState(h.Driver.Kill, state.Stopped); err != nil {
		return err
	}

	log.Infof("Bhojpur Host machine %q was killed.", h.Name)
	return nil
}

func (h *Host) Restart() error {
	log.Infof("Restarting %q...", h.Name)
	if drivers.MachineInState(h.Driver, state.Stopped)() {
		if err := h.Start(); err != nil {
			return err
		}
	} else if drivers.MachineInState(h.Driver, state.Running)() {
		if err := h.Driver.Restart(); err != nil {
			return err
		}
		if err := mutils.WaitFor(drivers.MachineInState(h.Driver, state.Running)); err != nil {
			return err
		}
	}

	return h.WaitForBhojpurHost()
}

func (h *Host) BhojpurVersion() (string, error) {
	url, err := h.Driver.GetURL()
	if err != nil {
		return "", err
	}

	bhojpurHost := &cengine.RemoteBhojpur{
		HostURL:    url,
		AuthOption: h.AuthOptions(),
	}
	bhojpurVersion, err := cengine.BhojpurVersion(bhojpurHost)
	if err != nil {
		return "", err
	}

	return bhojpurVersion, nil
}

func (h *Host) Upgrade() error {
	if h.HostOptions.AuthOptions == nil {
		log.Warnf(noBhojpurError, h.Name, "cannot upgrade Bhojpur Host")
		return nil
	}

	machineState, err := h.Driver.GetState()
	if err != nil {
		return err
	}

	if machineState != state.Running {
		log.Info("Starting Bhojpur Host machine so that machine could be upgraded...")
		if err := h.Start(); err != nil {
			return err
		}
	}

	provisioner, err := provision.DetectProvisioner(h.Driver)
	if err != nil {
		return err
	}

	bhojpurVersion, err := h.BhojpurVersion()
	if err != nil {
		return err
	}

	// If we're upgrading from a pre-CE (e.g., 1.13.1) release to a CE
	// release (e.g., 17.03.0-ce), we should simply uninstall and
	// re-install from scratch, since the official package names will
	// change from 'bhojpur-engine' to 'bhojpur-ce'.
	if versioncmp.LessThanOrEqualTo(bhojpurVersion, provision.LastReleaseBeforeCEVersioning) &&
		// BhojpurOS and boot2docker, being 'static ISO builds', have
		// an upgrade process which simply grabs the latest if it's
		// different, and so do not need to jump through this hoop to
		// upgrade safely.
		provisioner.String() != "bhojpuros" &&
		provisioner.String() != "boot2docker" {

		// Name of package 'bhojpur-engine' will fall through in this
		// case, so that we execute, e.g.,
		//
		// 'sudo apt-get purge -y bhojpur-engine'
		if err := provisioner.Package("bhojpur-engine", pkgaction.Purge); err != nil {
			return err
		}

		// Then we kick off the normal provisioning process which will
		// go off and install Bhojpur Host (get.docker.com script should work
		// fine to install Bhojpur Host from scratch after removing the old
		// packages, and images/containers etc. should be preserved in
		// /var/lib/bhojpur)
		return h.Provision()
	}

	log.Info("Upgrading Bhojpur Host...")
	if err := provisioner.Package("bhojpur", pkgaction.Upgrade); err != nil {
		return err
	}

	log.Info("Restarting Bhojpur Host...")
	return provisioner.Service("bhojpur", serviceaction.Restart)
}

func (h *Host) URL() (string, error) {
	return h.Driver.GetURL()
}

func (h *Host) AuthOptions() *auth.Options {
	if h.HostOptions == nil {
		return nil
	}
	return h.HostOptions.AuthOptions
}

func (h *Host) ConfigureAuth() error {
	if h.HostOptions.AuthOptions == nil {
		log.Warnf(noBhojpurError, h.Name, "cannot configure auth")
		return nil
	}

	provisioner, err := provision.DetectProvisioner(h.Driver)
	if err != nil {
		return err
	}

	// TODO: This is kind of a hack (or is it?  I'm not really sure until
	// we have more clearly defined outlook on what the responsibilities
	// and modularity of the provisioners should be).
	//
	// Call provision to re-provision the certs properly.
	return provisioner.Provision(swarm.Options{}, *h.HostOptions.AuthOptions, *h.HostOptions.EngineOptions)
}

func (h *Host) ConfigureAllAuth() error {
	if h.HostOptions.AuthOptions == nil {
		log.Warnf(noBhojpurError, h.Name, "cannot configure auth")
		return nil
	}

	log.Info("Regenerating local certificates")
	if err := cert.BootstrapCertificates(h.AuthOptions()); err != nil {
		return err
	}
	return h.ConfigureAuth()
}

func (h *Host) Provision() error {
	provisioner, err := provision.DetectProvisioner(h.Driver)
	if err != nil {
		return err
	}

	if h.HostOptions.CustomInstallScript != "" {
		log.Infof("Bhojpur Host machine %s was provisioned with a custom install script, using this script for provisioning", h.Name)
		return provision.WithCustomScript(provisioner, h.HostOptions.CustomInstallScript)
	}

	return provisioner.Provision(*h.HostOptions.SwarmOptions, *h.HostOptions.AuthOptions, *h.HostOptions.EngineOptions)
}

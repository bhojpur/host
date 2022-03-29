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
	"errors"
	"fmt"
	"regexp"
	"text/template"

	"github.com/bhojpur/host/pkg/core/auth"
	"github.com/bhojpur/host/pkg/core/drivers"
	"github.com/bhojpur/host/pkg/core/engine"
	"github.com/bhojpur/host/pkg/core/log"
	"github.com/bhojpur/host/pkg/core/provision/pkgaction"
	"github.com/bhojpur/host/pkg/core/provision/serviceaction"
	"github.com/bhojpur/host/pkg/core/swarm"
	mutils "github.com/bhojpur/host/pkg/core/utils"
)

var (
	ErrUnknownYumOsRelease = errors.New("unknown OS for Yum repository")
	engineConfigTemplate   = `[Service]
ExecStart=
ExecStart=/usr/bin/hostd -H tcp://0.0.0.0:{{.BhojpurPort}} -H unix:///var/run/bhojpur.sock --storage-driver {{.EngineOptions.StorageDriver}} --tlsverify --tlscacert {{.AuthOptions.CaCertRemotePath}} --tlscert {{.AuthOptions.ServerCertRemotePath}} --tlskey {{.AuthOptions.ServerKeyRemotePath}} {{ range .EngineOptions.Labels }}--label {{.}} {{ end }}{{ range .EngineOptions.InsecureRegistry }}--insecure-registry {{.}} {{ end }}{{ range .EngineOptions.RegistryMirror }}--registry-mirror {{.}} {{ end }}{{ range .EngineOptions.ArbitraryFlags }}--{{.}} {{ end }}
Environment={{range .EngineOptions.Env}}{{ printf "%q" . }} {{end}}
`
	majorVersionRE = regexp.MustCompile(`^(\d+)(\..*)?`)
)

type PackageListInfo struct {
	OsRelease        string
	OsReleaseVersion string
}

func init() {
	Register("RedHat", &RegisteredProvisioner{
		New: func(d drivers.Driver) Provisioner {
			return NewRedHatProvisioner("rhel", d)
		},
	})
}

func NewRedHatProvisioner(osReleaseID string, d drivers.Driver) *RedHatProvisioner {
	systemdProvisioner := NewSystemdProvisioner(osReleaseID, d)
	systemdProvisioner.SSHCommander = RedHatSSHCommander{Driver: d}
	return &RedHatProvisioner{
		systemdProvisioner,
	}
}

type RedHatProvisioner struct {
	SystemdProvisioner
}

func (provisioner *RedHatProvisioner) String() string {
	return "redhat"
}

func (provisioner *RedHatProvisioner) SetHostname(hostname string) error {
	// we have to have SetHostname here as well to use the RedHat provisioner
	// SSHCommand to add the tty allocation
	if _, err := provisioner.SSHCommand(fmt.Sprintf(
		"sudo hostname %s && echo %q | sudo tee /etc/hostname",
		hostname,
		hostname,
	)); err != nil {
		return err
	}

	if _, err := provisioner.SSHCommand(fmt.Sprintf(
		"if grep -xq 127.0.1.1.* /etc/hosts; then sudo sed -i 's/^127.0.1.1.*/127.0.1.1 %s/g' /etc/hosts; else echo '127.0.1.1 %s' | sudo tee -a /etc/hosts; fi",
		hostname,
		hostname,
	)); err != nil {
		return err
	}

	return nil
}

func (provisioner *RedHatProvisioner) Package(name string, action pkgaction.PackageAction) error {
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

func (provisioner *RedHatProvisioner) bhojpurDaemonResponding() bool {
	log.Debug("checking Bhojpur Host daemon")

	if out, err := provisioner.SSHCommand("sudo hostutl version"); err != nil {
		log.Warnf("Error getting SSH command to check if the daemon is up: %s", err)
		log.Debugf("'sudo hostutl version' output:\n%s", out)
		return false
	}

	// The daemon is up if the command worked.  Carry on.
	return true
}

func (provisioner *RedHatProvisioner) Provision(swarmOptions swarm.Options, authOptions auth.Options, engineOptions engine.Options) error {
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

	if err := installBhojpurGeneric(provisioner, provisioner.EngineOptions.InstallURL); err != nil {
		return err
	} else if err == nil {
		if err := provisioner.Service("bhojpur", serviceaction.Restart); err != nil {
			return err
		}
		if err := provisioner.Service("bhojpur", serviceaction.Enable); err != nil {
			return err
		}
	}

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

func (provisioner *RedHatProvisioner) GenerateBhojpurOptions(bhojpurPort int) (*BhojpurOptions, error) {
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

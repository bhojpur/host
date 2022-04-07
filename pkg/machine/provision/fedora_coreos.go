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
	"github.com/bhojpur/host/pkg/machine/swarm"
)

func init() {
	Register("Fedora-CoreOS", &RegisteredProvisioner{
		New: NewFedoraCoreOSProvisioner,
	})
}

// NewFedoraCoreOSProvisioner creates a new provisioner for a driver
func NewFedoraCoreOSProvisioner(d drivers.Driver) Provisioner {
	return &FedoraCoreOSProvisioner{
		NewSystemdProvisioner("fedora", d),
	}
}

// FedoraCoreOSProvisioner is a provisioner based on the CoreOS provisioner
type FedoraCoreOSProvisioner struct {
	SystemdProvisioner
}

// String returns the name of the provisioner
func (provisioner *FedoraCoreOSProvisioner) String() string {
	return "Fedora CoreOS"
}

// SetHostname sets the hostname of the remote machine
func (provisioner *FedoraCoreOSProvisioner) SetHostname(hostname string) error {
	log.Debugf("SetHostname: %s", hostname)

	command := fmt.Sprintf("sudo hostnamectl set-hostname %s", hostname)
	if _, err := provisioner.SSHCommand(command); err != nil {
		return err
	}

	return nil
}

// GenerateBhojpurOptions formats a systemd drop-in unit which adds support for
// Bhojpur Host machine
func (provisioner *FedoraCoreOSProvisioner) GenerateBhojpurOptions(bhojpurPort int) (*BhojpurOptions, error) {
	var (
		engineCfg bytes.Buffer
	)

	driverNameLabel := fmt.Sprintf("provider=%s", provisioner.Driver.DriverName())
	provisioner.EngineOptions.Labels = append(provisioner.EngineOptions.Labels, driverNameLabel)

	engineConfigTmpl := `[Service]
ExecStart=
ExecStart=/usr/bin/hostd \\
          --host=fd:// \\
          --exec-opt native.cgroupdriver=systemd \\
          --host=tcp://0.0.0.0:{{.BhojpurPort}} \\
          --tlsverify \\
          --tlscacert {{.AuthOptions.CaCertRemotePath}} \\
          --tlscert {{.AuthOptions.ServerCertRemotePath}} \\
          --tlskey {{.AuthOptions.ServerKeyRemotePath}}{{ range .EngineOptions.Labels }} \\
          --label {{.}}{{ end }}{{ range .EngineOptions.InsecureRegistry }} \\
          --insecure-registry {{.}}{{ end }}{{ range .EngineOptions.RegistryMirror }} \\
          --registry-mirror {{.}}{{ end }}{{ range .EngineOptions.ArbitraryFlags }} \\
          -{{.}}{{ end }} \\
          \$OPTIONS
Environment={{range .EngineOptions.Env}}{{ printf "%q" . }} {{end}}
`

	t, err := template.New("engineConfig").Parse(engineConfigTmpl)
	if err != nil {
		return nil, err
	}

	engineConfigContext := EngineConfigContext{
		BhojpurPort:   bhojpurPort,
		AuthOptions:   provisioner.AuthOptions,
		EngineOptions: provisioner.EngineOptions,
	}

	t.Execute(&engineCfg, engineConfigContext)

	return &BhojpurOptions{
		EngineOptions:     engineCfg.String(),
		EngineOptionsPath: provisioner.DaemonOptionsFile,
	}, nil
}

// CompatibleWithHost returns whether or not this provisoner is compatible
// with the target host
func (provisioner *FedoraCoreOSProvisioner) CompatibleWithHost() bool {
	isFedora := provisioner.OsReleaseInfo.ID == "fedora"
	isCoreOS := provisioner.OsReleaseInfo.VariantID == "coreos"
	return isFedora && isCoreOS
}

// Package installs a package on the remote host. The Fedora CoreOS provisioner
// does not support (or need) any package installation
func (provisioner *FedoraCoreOSProvisioner) Package(name string, action pkgaction.PackageAction) error {
	return nil
}

// Provision provisions the machine
func (provisioner *FedoraCoreOSProvisioner) Provision(swarmOptions swarm.Options, authOptions auth.Options, engineOptions engine.Options) error {
	provisioner.SwarmOptions = swarmOptions
	provisioner.AuthOptions = authOptions
	provisioner.EngineOptions = engineOptions

	if err := provisioner.SetHostname(provisioner.Driver.GetMachineName()); err != nil {
		return err
	}

	if err := makeBhojpurOptionsDir(provisioner); err != nil {
		return err
	}

	log.Debugf("Preparing certificates")
	provisioner.AuthOptions = setRemoteAuthOptions(provisioner)

	log.Debugf("Setting up certificates")
	if err := ConfigureAuth(provisioner); err != nil {
		return err
	}

	log.Debug("Configuring swarm")
	err := configureSwarm(provisioner, swarmOptions, provisioner.AuthOptions)
	return err
}

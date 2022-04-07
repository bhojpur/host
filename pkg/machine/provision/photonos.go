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
	"strings"
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
	Register("PhotonOS", &RegisteredProvisioner{
		New: NewPhotonOSProvisioner,
	})
}

// NewPhotonOSProvisioner creates a new provisioner for a driver
func NewPhotonOSProvisioner(d drivers.Driver) Provisioner {
	return &PhotonOSProvisioner{
		NewSystemdProvisioner("photon", d),
	}
}

// PhotonOSProvisioner is a provisioner based on the SystemdProvisioner provisioner
type PhotonOSProvisioner struct {
	SystemdProvisioner
}

// String returns the name of the provisioner
func (provisioner *PhotonOSProvisioner) String() string {
	return "Photon OS"
}

// SetHostname sets the hostname of the remote machine
func (provisioner *PhotonOSProvisioner) SetHostname(hostname string) error {
	log.Debugf("SetHostname: %s", hostname)

	command := fmt.Sprintf("sudo hostnamectl set-hostname %s", hostname)
	if _, err := provisioner.SSHCommand(command); err != nil {
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

// GenerateBhojpurOptions formats a systemd drop-in unit which adds support for
// Bhojpur Host machine
func (provisioner *PhotonOSProvisioner) GenerateBhojpurOptions(bhojpurPort int) (*BhojpurOptions, error) {
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
func (provisioner *PhotonOSProvisioner) CompatibleWithHost() bool {
	return provisioner.OsReleaseInfo.ID == provisioner.OsReleaseID
}

// Package installs a package on the remote host. The Photon OS provisioner
// does not support (or need) any package installation
func (provisioner *PhotonOSProvisioner) Package(name string, action pkgaction.PackageAction) error {
	return nil
}

func (provisioner *PhotonOSProvisioner) bhojpurDaemonResponding() bool {
	log.Debug("checking Bhojpur Host daemon")

	if out, err := provisioner.SSHCommand("sudo hostutl version"); err != nil {
		log.Warnf("Error getting SSH command to check if the daemon is up: %s", err)
		log.Debugf("'sudo hostutl version' output:\n%s", out)
		return false
	}

	// The daemon is up if the command worked.  Carry on.
	return true
}

// Provision provisions the machine
func (provisioner *PhotonOSProvisioner) Provision(swarmOptions swarm.Options, authOptions auth.Options, engineOptions engine.Options) error {
	provisioner.SwarmOptions = swarmOptions
	provisioner.AuthOptions = authOptions
	provisioner.EngineOptions = engineOptions

	if err := provisioner.SetHostname(provisioner.Driver.GetMachineName()); err != nil {
		return err
	}

	log.Debug("Configuring Photon OS firewall")
	if err := provisioner.configureFirewall(); err != nil {
		return err
	}

	if err := makeBhojpurOptionsDir(provisioner); err != nil {
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

	log.Debugf("Preparing certificates")
	provisioner.AuthOptions = setRemoteAuthOptions(provisioner)

	log.Debugf("Setting up certificates")
	if err := ConfigureAuth(provisioner); err != nil {
		return err
	}

	log.Debug("Configuring swarm")
	err := configureSwarm(provisioner, swarmOptions, provisioner.AuthOptions)

	// enable in systemd
	log.Debug("Enabling Bhojpur Host in systemd")
	err = provisioner.Service("bhojpur", serviceaction.Enable)
	return err
}

// configureFirewall sets up proper iptable rules
func (provisioner *PhotonOSProvisioner) configureFirewall() error {
	tcpPorts := "22 80 443 2376 2379 2380 6443 9099 9796 10250 10254 30000:32767"
	udpPorts := "8472 30000:32767"
	var cmds []string

	for _, port := range strings.Split(tcpPorts, " ") {
		cmds = append(cmds, fmt.Sprintf("sudo iptables -A INPUT -p tcp --dport %s -j ACCEPT", port))
	}
	for _, port := range strings.Split(udpPorts, " ") {
		cmds = append(cmds, fmt.Sprintf("sudo iptables -A INPUT -p udp --dport %s -j ACCEPT", port))
	}
	cmds = append(cmds, "sudo sh -c 'iptables-save > /etc/systemd/scripts/ip4save'")

	for _, cmd := range cmds {
		if _, err := provisioner.SSHCommand(cmd); err != nil {
			return err
		}
	}
	return nil
}

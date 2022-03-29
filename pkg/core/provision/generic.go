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

	"github.com/bhojpur/host/pkg/core/auth"
	"github.com/bhojpur/host/pkg/core/drivers"
	"github.com/bhojpur/host/pkg/core/engine"
	"github.com/bhojpur/host/pkg/core/swarm"
)

type GenericProvisioner struct {
	SSHCommander
	OsReleaseID       string
	BhojpurOptionsDir string
	DaemonOptionsFile string
	Packages          []string
	OsReleaseInfo     *OsRelease
	Driver            drivers.Driver
	AuthOptions       auth.Options
	EngineOptions     engine.Options
	SwarmOptions      swarm.Options
}

type GenericSSHCommander struct {
	Driver drivers.Driver
}

func (sshCmder GenericSSHCommander) SSHCommand(args string) (string, error) {
	return drivers.RunSSHCommandFromDriver(sshCmder.Driver, args)
}

func (provisioner *GenericProvisioner) Hostname() (string, error) {
	return provisioner.SSHCommand("hostname")
}

func (provisioner *GenericProvisioner) SetHostname(hostname string) error {
	if _, err := provisioner.SSHCommand(fmt.Sprintf(
		"sudo hostname %s && echo %q | sudo tee /etc/hostname",
		hostname,
		hostname,
	)); err != nil {
		return err
	}

	// ubuntu/debian use 127.0.1.1 for non "localhost" loopback hostnames: https://www.debian.org/doc/manuals/debian-reference/ch05.en.html#_the_hostname_resolution
	if _, err := provisioner.SSHCommand(fmt.Sprintf(`
		if ! grep -xq '.*\s%s' /etc/hosts; then
			if grep -xq '127.0.1.1\s.*' /etc/hosts; then
				sudo sed -i 's/^127.0.1.1\s.*/127.0.1.1 %s/g' /etc/hosts;
			else 
				echo '127.0.1.1 %s' | sudo tee -a /etc/hosts; 
			fi
		fi`,
		hostname,
		hostname,
		hostname,
	)); err != nil {
		return err
	}

	return nil
}

func (provisioner *GenericProvisioner) GetBhojpurOptionsDir() string {
	return provisioner.BhojpurOptionsDir
}

func (provisioner *GenericProvisioner) CompatibleWithHost() bool {
	return provisioner.OsReleaseInfo.ID == provisioner.OsReleaseID
}

func (provisioner *GenericProvisioner) GetAuthOptions() auth.Options {
	return provisioner.AuthOptions
}

func (provisioner *GenericProvisioner) GetSwarmOptions() swarm.Options {
	return provisioner.SwarmOptions
}

func (provisioner *GenericProvisioner) SetOsReleaseInfo(info *OsRelease) {
	provisioner.OsReleaseInfo = info
}

func (provisioner *GenericProvisioner) GetOsReleaseInfo() (*OsRelease, error) {
	return provisioner.OsReleaseInfo, nil
}

func (provisioner *GenericProvisioner) GenerateBhojpurOptions(bhojpurPort int) (*BhojpurOptions, error) {
	var (
		engineCfg bytes.Buffer
	)

	driverNameLabel := fmt.Sprintf("provider=%s", provisioner.Driver.DriverName())
	provisioner.EngineOptions.Labels = append(provisioner.EngineOptions.Labels, driverNameLabel)

	engineConfigTmpl := `
BHOJPUR_OPTS='
-H tcp://0.0.0.0:{{.BhojpurPort}}
-H unix:///var/run/bhojpur.sock
--storage-driver {{.EngineOptions.StorageDriver}}
--tlsverify
--tlscacert {{.AuthOptions.CaCertRemotePath}}
--tlscert {{.AuthOptions.ServerCertRemotePath}}
--tlskey {{.AuthOptions.ServerKeyRemotePath}}
{{ range .EngineOptions.Labels }}--label {{.}}
{{ end }}{{ range .EngineOptions.InsecureRegistry }}--insecure-registry {{.}}
{{ end }}{{ range .EngineOptions.RegistryMirror }}--registry-mirror {{.}}
{{ end }}{{ range .EngineOptions.ArbitraryFlags }}--{{.}}
{{ end }}
'
{{range .EngineOptions.Env}}export \"{{ printf "%q" . }}\"
{{end}}
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

func (provisioner *GenericProvisioner) GetDriver() drivers.Driver {
	return provisioner.Driver
}

func (provisioner *GenericProvisioner) GetPackages() []string {
	return provisioner.Packages
}

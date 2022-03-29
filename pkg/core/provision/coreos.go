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
	"github.com/bhojpur/host/pkg/core/log"
	"github.com/bhojpur/host/pkg/core/provision/pkgaction"
	"github.com/bhojpur/host/pkg/core/provision/serviceaction"
	"github.com/bhojpur/host/pkg/core/swarm"
	"github.com/bhojpur/host/pkg/core/versioncmp"
)

const (
	hostTmpl = `sudo tee /var/tmp/hostname.yml << EOF
#cloud-config

hostname: %s
EOF
`
)

func init() {
	Register("CoreOS", &RegisteredProvisioner{
		New: NewCoreOSProvisioner,
	})
}

func NewCoreOSProvisioner(d drivers.Driver) Provisioner {
	return &CoreOSProvisioner{
		NewSystemdProvisioner("coreos", d),
	}
}

type CoreOSProvisioner struct {
	SystemdProvisioner
}

func (provisioner *CoreOSProvisioner) String() string {
	return "coreOS"
}

func (provisioner *CoreOSProvisioner) CompatibleWithHost() bool {
	return provisioner.OsReleaseInfo.ID == provisioner.OsReleaseID || provisioner.OsReleaseInfo.IDLike == provisioner.OsReleaseID
}

func (provisioner *CoreOSProvisioner) SetHostname(hostname string) error {
	log.Debugf("SetHostname: %s", hostname)

	if _, err := provisioner.SSHCommand(fmt.Sprintf(hostTmpl, hostname)); err != nil {
		return err
	}

	if _, err := provisioner.SSHCommand("sudo systemctl start system-cloudinit@var-tmp-hostname.yml.service"); err != nil {
		return err
	}

	return nil
}

func (provisioner *CoreOSProvisioner) GenerateBhojpurOptions(bhojpurPort int) (*BhojpurOptions, error) {
	var (
		engineCfg bytes.Buffer
	)

	driverNameLabel := fmt.Sprintf("provider=%s", provisioner.Driver.DriverName())
	provisioner.EngineOptions.Labels = append(provisioner.EngineOptions.Labels, driverNameLabel)

	bhojpurVersion, err := BhojpurClientVersion(provisioner)
	if err != nil {
		return nil, err
	}

	arg := "daemon"
	if versioncmp.GreaterThanOrEqualTo(bhojpurVersion, "1.12.0") {
		arg = ""
	}

	engineConfigTmpl := `[Service]
Environment=TMPDIR=/var/tmp
ExecStart=
ExecStart=/usr/lib/coreos/hostd ` + arg + ` --host=unix:///var/run/bhojpur.sock --host=tcp://0.0.0.0:{{.BhojpurPort}} --tlsverify --tlscacert {{.AuthOptions.CaCertRemotePath}} --tlscert {{.AuthOptions.ServerCertRemotePath}} --tlskey {{.AuthOptions.ServerKeyRemotePath}}{{ range .EngineOptions.Labels }} --label {{.}}{{ end }}{{ range .EngineOptions.InsecureRegistry }} --insecure-registry {{.}}{{ end }}{{ range .EngineOptions.RegistryMirror }} --registry-mirror {{.}}{{ end }}{{ range .EngineOptions.ArbitraryFlags }} --{{.}}{{ end }} \$BHOJPUR_OPTS \$BHOJPUR_OPT_BIP \$BHOJPUR_OPT_MTU \$BHOJPUR_OPT_IPMASQ
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

func (provisioner *CoreOSProvisioner) Package(name string, action pkgaction.PackageAction) error {
	return nil
}

func (provisioner *CoreOSProvisioner) Provision(swarmOptions swarm.Options, authOptions auth.Options, engineOptions engine.Options) error {
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
	if err := configureSwarm(provisioner, swarmOptions, provisioner.AuthOptions); err != nil {
		return err
	}

	// enable in systemd
	log.Debug("enabling Bhojpur Host in systemd")
	err := provisioner.Service("bhojpur", serviceaction.Enable)
	return err
}

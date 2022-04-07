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

	"github.com/bhojpur/host/pkg/machine/drivers"
	"github.com/bhojpur/host/pkg/machine/provision/serviceaction"
	"github.com/bhojpur/host/pkg/machine/versioncmp"
)

type SystemdProvisioner struct {
	GenericProvisioner
}

func (p *SystemdProvisioner) String() string {
	return "redhat"
}

func NewSystemdProvisioner(osReleaseID string, d drivers.Driver) SystemdProvisioner {
	return SystemdProvisioner{
		GenericProvisioner{
			SSHCommander:      GenericSSHCommander{Driver: d},
			BhojpurOptionsDir: "/etc/bhojpur",
			DaemonOptionsFile: "/etc/systemd/system/bhojpur.service.d/10-machine.conf",
			OsReleaseID:       osReleaseID,
			Packages: []string{
				"curl",
			},
			Driver: d,
		},
	}
}

func (p *SystemdProvisioner) GenerateBhojpurOptions(bhojpurPort int) (*BhojpurOptions, error) {
	var (
		engineCfg bytes.Buffer
	)

	driverNameLabel := fmt.Sprintf("provider=%s", p.Driver.DriverName())
	p.EngineOptions.Labels = append(p.EngineOptions.Labels, driverNameLabel)

	bhojpurVersion, err := BhojpurClientVersion(p)
	if err != nil {
		return nil, err
	}

	arg := "hostd"
	if versioncmp.LessThan(bhojpurVersion, "1.12.0") {
		arg = "hostd daemon"
	}

	engineConfigTmpl := `[Service]
ExecStart=
ExecStart=/usr/bin/` + arg + ` -H tcp://0.0.0.0:{{.BhojpurPort}} -H unix:///var/run/bhojpur.sock --storage-driver {{.EngineOptions.StorageDriver}} --tlsverify --tlscacert {{.AuthOptions.CaCertRemotePath}} --tlscert {{.AuthOptions.ServerCertRemotePath}} --tlskey {{.AuthOptions.ServerKeyRemotePath}} {{ range .EngineOptions.Labels }}--label {{.}} {{ end }}{{ range .EngineOptions.InsecureRegistry }}--insecure-registry {{.}} {{ end }}{{ range .EngineOptions.RegistryMirror }}--registry-mirror {{.}} {{ end }}{{ range .EngineOptions.ArbitraryFlags }}--{{.}} {{ end }}
Environment={{range .EngineOptions.Env}}{{ printf "%q" . }} {{end}}
`
	t, err := template.New("engineConfig").Parse(engineConfigTmpl)
	if err != nil {
		return nil, err
	}

	engineConfigContext := EngineConfigContext{
		BhojpurPort:   bhojpurPort,
		AuthOptions:   p.AuthOptions,
		EngineOptions: p.EngineOptions,
	}

	t.Execute(&engineCfg, engineConfigContext)

	return &BhojpurOptions{
		EngineOptions:     engineCfg.String(),
		EngineOptionsPath: p.DaemonOptionsFile,
	}, nil
}

func (p *SystemdProvisioner) Service(name string, action serviceaction.ServiceAction) error {
	reloadDaemon := false
	switch action {
	case serviceaction.Start, serviceaction.Restart:
		reloadDaemon = true
	}

	// systemd needs reloaded when config changes on disk; we cannot
	// be sure exactly when it changes from the provisioner so
	// we call a reload on every restart to be safe
	if reloadDaemon {
		if _, err := p.SSHCommand("sudo systemctl daemon-reload"); err != nil {
			return err
		}
	}

	command := fmt.Sprintf("sudo systemctl -f %s %s", action.String(), name)

	if _, err := p.SSHCommand(command); err != nil {
		return err
	}

	return nil
}

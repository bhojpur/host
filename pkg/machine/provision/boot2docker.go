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
	"encoding/json"
	"fmt"
	"net"
	"path"
	"text/template"
	"time"

	mdirs "github.com/bhojpur/host/cmd/machine/commands/dirs"
	"github.com/bhojpur/host/pkg/machine/auth"
	"github.com/bhojpur/host/pkg/machine/drivers"
	"github.com/bhojpur/host/pkg/machine/engine"
	"github.com/bhojpur/host/pkg/machine/log"
	"github.com/bhojpur/host/pkg/machine/provision/pkgaction"
	"github.com/bhojpur/host/pkg/machine/provision/serviceaction"
	"github.com/bhojpur/host/pkg/machine/state"
	"github.com/bhojpur/host/pkg/machine/swarm"
	mutils "github.com/bhojpur/host/pkg/machine/utils"
)

func init() {
	Register("boot2docker", &RegisteredProvisioner{
		New: NewBoot2DockerProvisioner,
	})
}

func NewBoot2DockerProvisioner(d drivers.Driver) Provisioner {
	return &Boot2DockerProvisioner{
		Driver: d,
	}
}

type Boot2DockerProvisioner struct {
	OsReleaseInfo *OsRelease
	Driver        drivers.Driver
	AuthOptions   auth.Options
	EngineOptions engine.Options
	SwarmOptions  swarm.Options
}

func (provisioner *Boot2DockerProvisioner) String() string {
	return "boot2docker"
}

func (provisioner *Boot2DockerProvisioner) Service(name string, action serviceaction.ServiceAction) error {
	_, err := provisioner.SSHCommand(fmt.Sprintf("sudo /etc/init.d/%s %s", name, action.String()))
	return err
}

func (provisioner *Boot2DockerProvisioner) upgradeIso() error {
	// TODO: Ideally, we should not read from mcndirs directory at all.
	// The driver should be able to communicate how and where to place the
	// relevant files.
	b2dutils := mutils.NewB2dUtils(mdirs.GetBaseDir())

	// Check if the driver has specified a custom b2d url
	jsonDriver, err := json.Marshal(provisioner.GetDriver())
	if err != nil {
		return err
	}
	var d struct {
		Boot2DockerURL string
	}
	json.Unmarshal(jsonDriver, &d)

	log.Info("Stopping Bhojpur Host machine to do the upgrade...")

	if err := provisioner.Driver.Stop(); err != nil {
		return err
	}

	if err := mutils.WaitFor(drivers.MachineInState(provisioner.Driver, state.Stopped)); err != nil {
		return err
	}

	machineName := provisioner.GetDriver().GetMachineName()

	log.Infof("Upgrading machine %q...", machineName)

	// Either download the latest version of the b2d url that was explicitly
	// specified when creating the VM or copy the (updated) default ISO
	if err := b2dutils.CopyIsoToMachineDir(d.Boot2DockerURL, machineName); err != nil {
		return err
	}

	log.Infof("Starting Bhojpur Host machine back up...")

	if err := provisioner.Driver.Start(); err != nil {
		return err
	}

	return mutils.WaitFor(drivers.MachineInState(provisioner.Driver, state.Running))
}

func (provisioner *Boot2DockerProvisioner) Package(name string, action pkgaction.PackageAction) error {
	if name == "docker" && action == pkgaction.Upgrade {
		if err := provisioner.upgradeIso(); err != nil {
			return err
		}
	}
	return nil
}

func (provisioner *Boot2DockerProvisioner) Hostname() (string, error) {
	return provisioner.SSHCommand("hostname")
}

func (provisioner *Boot2DockerProvisioner) SetHostname(hostname string) error {
	if _, err := provisioner.SSHCommand(fmt.Sprintf(
		"sudo /usr/bin/sethostname %s && echo %q | sudo tee /var/lib/boot2docker/etc/hostname",
		hostname,
		hostname,
	)); err != nil {
		return err
	}

	return nil
}

func (provisioner *Boot2DockerProvisioner) GetBhojpurOptionsDir() string {
	return "/var/lib/boot2docker"
}

func (provisioner *Boot2DockerProvisioner) GetAuthOptions() auth.Options {
	return provisioner.AuthOptions
}

func (provisioner *Boot2DockerProvisioner) GetSwarmOptions() swarm.Options {
	return provisioner.SwarmOptions
}

func (provisioner *Boot2DockerProvisioner) GenerateBhojpurOptions(bhojpurPort int) (*BhojpurOptions, error) {
	var (
		engineCfg bytes.Buffer
	)

	driverNameLabel := fmt.Sprintf("provider=%s", provisioner.Driver.DriverName())
	provisioner.EngineOptions.Labels = append(provisioner.EngineOptions.Labels, driverNameLabel)

	engineConfigTmpl := `
EXTRA_ARGS='
{{ range .EngineOptions.Labels }}--label {{.}}
{{ end }}{{ range .EngineOptions.InsecureRegistry }}--insecure-registry {{.}}
{{ end }}{{ range .EngineOptions.RegistryMirror }}--registry-mirror {{.}}
{{ end }}{{ range .EngineOptions.ArbitraryFlags }}--{{.}}
{{ end }}
'
CACERT={{.AuthOptions.CaCertRemotePath}}
BHOJPUR_HOST='-H tcp://0.0.0.0:{{.BhojpurPort}}'
BHOJPUR_STORAGE={{.EngineOptions.StorageDriver}}
BHOJPUR_TLS=auto
SERVERKEY={{.AuthOptions.ServerKeyRemotePath}}
SERVERCERT={{.AuthOptions.ServerCertRemotePath}}

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

	daemonOptsDir := path.Join(provisioner.GetBhojpurOptionsDir(), "profile")
	return &BhojpurOptions{
		EngineOptions:     engineCfg.String(),
		EngineOptionsPath: daemonOptsDir,
	}, nil
}

func (provisioner *Boot2DockerProvisioner) CompatibleWithHost() bool {
	return provisioner.OsReleaseInfo.ID == "boot2docker"
}

func (provisioner *Boot2DockerProvisioner) SetOsReleaseInfo(info *OsRelease) {
	provisioner.OsReleaseInfo = info
}

func (provisioner *Boot2DockerProvisioner) GetOsReleaseInfo() (*OsRelease, error) {
	return provisioner.OsReleaseInfo, nil
}

func (provisioner *Boot2DockerProvisioner) AttemptIPContact(bhojpurPort int) {
	ip, err := provisioner.Driver.GetIP()
	if err != nil {
		log.Warnf("Could not get IP address for created machine: %s", err)
		return
	}

	if conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ip, bhojpurPort), 5*time.Second); err != nil {
		log.Warnf(`
This machine has been allocated an IP address, but Bhojpur Host machine could not
reach it successfully.

SSH for the machine should still work, but connecting to exposed ports, such as
the Bhojpur Host daemon port (usually <ip>:%d), may not work properly.

You may need to add the route manually, or use another related workaround.

This could be due to a VPN, proxy, or host file configuration issue.

You also might want to clear any VirtualBox host only interfaces you are not using.`, engine.DefaultPort)
	} else {
		conn.Close()
	}
}

func (provisioner *Boot2DockerProvisioner) Provision(swarmOptions swarm.Options, authOptions auth.Options, engineOptions engine.Options) error {
	var (
		err error
	)

	defer func() {
		if err == nil {
			provisioner.AttemptIPContact(engine.DefaultPort)
		}
	}()

	provisioner.SwarmOptions = swarmOptions
	provisioner.AuthOptions = authOptions
	provisioner.EngineOptions = engineOptions
	swarmOptions.Env = engineOptions.Env

	if provisioner.EngineOptions.StorageDriver == "" {
		provisioner.EngineOptions.StorageDriver = DefaultStorageDriver
	}

	if err = provisioner.SetHostname(provisioner.Driver.GetMachineName()); err != nil {
		return err
	}

	// b2d hosts need to wait for the daemon to be up
	// before continuing with provisioning
	if err = WaitForBhojpur(provisioner, engine.DefaultPort); err != nil {
		return err
	}

	if err = makeBhojpurOptionsDir(provisioner); err != nil {
		return err
	}

	provisioner.AuthOptions = setRemoteAuthOptions(provisioner)

	if err = ConfigureAuth(provisioner); err != nil {
		return err
	}

	err = configureSwarm(provisioner, swarmOptions, provisioner.AuthOptions)
	return err
}

func (provisioner *Boot2DockerProvisioner) SSHCommand(args string) (string, error) {
	return drivers.RunSSHCommandFromDriver(provisioner.Driver, args)
}

func (provisioner *Boot2DockerProvisioner) GetDriver() drivers.Driver {
	return provisioner.Driver
}

func (provisioner *Boot2DockerProvisioner) GetPackages() []string {
	return nil
}

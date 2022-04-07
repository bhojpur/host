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
	"io/ioutil"
	"net/url"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/bhojpur/host/pkg/machine/auth"
	"github.com/bhojpur/host/pkg/machine/cert"
	"github.com/bhojpur/host/pkg/machine/engine"
	"github.com/bhojpur/host/pkg/machine/log"
	"github.com/bhojpur/host/pkg/machine/provision/serviceaction"
	mutils "github.com/bhojpur/host/pkg/machine/utils"
)

type BhojpurOptions struct {
	EngineOptions     string
	EngineOptionsPath string
}

func installBhojpurGeneric(p Provisioner, baseURL string) error {
	if strings.EqualFold(baseURL, "none") {
		log.Info("Skipping Bhojpur Host installation")
		return nil
	}
	// install Bhojpur Host - until cloudinit we use ubuntu everywhere so we
	// just install it using the Bhojpur Host repos
	log.Infof("Installing Bhojpur Host from: %s", baseURL)
	if output, err := p.SSHCommand(fmt.Sprintf("if ! type hostutl; then curl -sSL %s | sh -; fi", baseURL)); err != nil {
		return fmt.Errorf("Error installing Bhojpur Host: %s", output)
	}

	return nil
}

func makeBhojpurOptionsDir(p Provisioner) error {
	bhojpurDir := p.GetBhojpurOptionsDir()
	if _, err := p.SSHCommand(fmt.Sprintf("sudo mkdir -p %s", bhojpurDir)); err != nil {
		return err
	}

	return nil
}

func setRemoteAuthOptions(p Provisioner) auth.Options {
	bhojpurDir := p.GetBhojpurOptionsDir()
	authOptions := p.GetAuthOptions()

	// due to windows clients, we cannot use filepath.Join as the paths
	// will be mucked on the linux hosts
	authOptions.CaCertRemotePath = path.Join(bhojpurDir, "ca.pem")
	authOptions.ServerCertRemotePath = path.Join(bhojpurDir, "server.pem")
	authOptions.ServerKeyRemotePath = path.Join(bhojpurDir, "server-key.pem")

	return authOptions
}

func ConfigureAuth(p Provisioner) error {
	var (
		err error
	)

	driver := p.GetDriver()
	machineName := driver.GetMachineName()
	authOptions := p.GetAuthOptions()
	swarmOptions := p.GetSwarmOptions()
	org := mutils.GetUsername() + "." + machineName
	bits := 2048

	ip, err := driver.GetIP()
	if err != nil {
		return err
	}

	log.Info("Copying certs to the local machine directory...")

	if err := mutils.CopyFile(authOptions.CaCertPath, filepath.Join(authOptions.StorePath, "ca.pem")); err != nil {
		return fmt.Errorf("Copying ca.pem to machine dir failed: %s", err)
	}

	if err := mutils.CopyFile(authOptions.ClientCertPath, filepath.Join(authOptions.StorePath, "cert.pem")); err != nil {
		return fmt.Errorf("Copying cert.pem to machine dir failed: %s", err)
	}

	if err := mutils.CopyFile(authOptions.ClientKeyPath, filepath.Join(authOptions.StorePath, "key.pem")); err != nil {
		return fmt.Errorf("Copying key.pem to machine dir failed: %s", err)
	}

	// The Host IP is always added to the certificate's SANs list
	hosts := append(authOptions.ServerCertSANs, ip, "localhost")
	log.Debugf("generating server cert: %s ca-key=%s private-key=%s org=%s san=%s",
		authOptions.ServerCertPath,
		authOptions.CaCertPath,
		authOptions.CaPrivateKeyPath,
		org,
		hosts,
	)

	// TODO: Switch to passing just authOptions to this func
	// instead of all these individual fields
	err = cert.GenerateCert(&cert.Options{
		Hosts:       hosts,
		CertFile:    authOptions.ServerCertPath,
		KeyFile:     authOptions.ServerKeyPath,
		CAFile:      authOptions.CaCertPath,
		CAKeyFile:   authOptions.CaPrivateKeyPath,
		Org:         org,
		Bits:        bits,
		SwarmMaster: swarmOptions.Master,
	})

	if err != nil {
		return fmt.Errorf("error generating server cert: %s", err)
	}

	if err := p.Service("bhojpur", serviceaction.Stop); err != nil {
		return err
	}

	if _, err := p.SSHCommand(`if [ ! -z "$(ip link show bhojpur0)" ]; then sudo ip link delete bhojpur0; fi`); err != nil {
		return err
	}

	// upload certs and configure TLS auth
	caCert, err := ioutil.ReadFile(authOptions.CaCertPath)
	if err != nil {
		return err
	}

	serverCert, err := ioutil.ReadFile(authOptions.ServerCertPath)
	if err != nil {
		return err
	}
	serverKey, err := ioutil.ReadFile(authOptions.ServerKeyPath)
	if err != nil {
		return err
	}

	log.Info("Copying certs to the remote machine...")

	// printf will choke if we don't pass a format string because of the
	// dashes, so that's the reason for the '%%s'
	certTransferCmdFmt := "printf '%%s' '%s' | sudo tee %s"

	// These ones are for Jessie and Mike <3 <3 <3
	if _, err := p.SSHCommand(fmt.Sprintf(certTransferCmdFmt, string(caCert), authOptions.CaCertRemotePath)); err != nil {
		return err
	}

	if _, err := p.SSHCommand(fmt.Sprintf(certTransferCmdFmt, string(serverCert), authOptions.ServerCertRemotePath)); err != nil {
		return err
	}

	if _, err := p.SSHCommand(fmt.Sprintf(certTransferCmdFmt, string(serverKey), authOptions.ServerKeyRemotePath)); err != nil {
		return err
	}

	bhojpurURL, err := driver.GetURL()
	if err != nil {
		return err
	}
	u, err := url.Parse(bhojpurURL)
	if err != nil {
		return err
	}
	bhojpurPort := engine.DefaultPort
	parts := strings.Split(u.Host, ":")
	if len(parts) == 2 {
		dPort, err := strconv.Atoi(parts[1])
		if err != nil {
			return err
		}
		bhojpurPort = dPort
	}

	dkrcfg, err := p.GenerateBhojpurOptions(bhojpurPort)
	if err != nil {
		return err
	}

	log.Info("Setting Bhojpur Host configuration on the remote daemon...")

	if _, err = p.SSHCommand(fmt.Sprintf("sudo mkdir -p %s && printf %%s \"%s\" | sudo tee %s", path.Dir(dkrcfg.EngineOptionsPath), dkrcfg.EngineOptions, dkrcfg.EngineOptionsPath)); err != nil {
		return err
	}

	if err := p.Service("bhojpur", serviceaction.Restart); err != nil {
		return err
	}

	return WaitForBhojpur(p, bhojpurPort)
}

func matchNetstatOut(reDaemonListening, netstatOut string) bool {
	// TODO: I would really prefer this be a Scanner directly on
	// the STDOUT of the executed command than to do all the string
	// manipulation hokey-pokey.
	//
	// TODO: Unit test this matching.
	for _, line := range strings.Split(netstatOut, "\n") {
		match, err := regexp.MatchString(reDaemonListening, line)
		if err != nil {
			log.Warnf("Regex warning: %s", err)
		}
		if match && line != "" {
			return true
		}
	}

	return false
}

func decideStorageDriver(p Provisioner, defaultDriver, suppliedDriver string) (string, error) {
	if suppliedDriver != "" {
		return suppliedDriver, nil
	}
	bestSuitedDriver := ""

	defer func() {
		if bestSuitedDriver != "" {
			log.Debugf("No storagedriver specified, using %s\n", bestSuitedDriver)
		}
	}()

	if defaultDriver != DefaultStorageDriver {
		bestSuitedDriver = defaultDriver
	} else {
		remoteFilesystemType, err := getFilesystemType(p, "/var/lib")
		if err != nil {
			return "", err
		}
		if remoteFilesystemType == "btrfs" {
			bestSuitedDriver = "btrfs"
		} else {
			bestSuitedDriver = DefaultStorageDriver
		}
	}
	return bestSuitedDriver, nil

}

func getFilesystemType(p Provisioner, directory string) (string, error) {
	statCommandOutput, err := p.SSHCommand("stat -f -c %T " + directory)
	if err != nil {
		err = fmt.Errorf("Error looking up filesystem type: %s", err)
		return "", err
	}

	fstype := strings.TrimSpace(statCommandOutput)
	return fstype, nil
}

func checkDaemonUp(p Provisioner, bhojpurPort int) func() bool {
	reDaemonListening := fmt.Sprintf(":%d\\s+.*:.*", bhojpurPort)
	return func() bool {
		// HACK: Check netstat's output to see if anyone's listening on the Bhojpur Host API port.
		netstatOut, err := p.SSHCommand("if ! type netstat 1>/dev/null; then ss -tln; else netstat -tln; fi")
		if err != nil {
			log.Warnf("Error running SSH command: %s", err)
			return false
		}

		return matchNetstatOut(reDaemonListening, netstatOut)
	}
}

func WaitForBhojpur(p Provisioner, bhojpurPort int) error {
	if err := mutils.WaitForSpecific(checkDaemonUp(p, bhojpurPort), 10, 3*time.Second); err != nil {
		return NewErrDaemonAvailable(err)
	}

	return nil
}

// BhojpurClientVersion returns the version of the Bhojpur Host client on the host
// that ssh is connected to, e.g. "1.12.1".
func BhojpurClientVersion(ssh SSHCommander) (string, error) {
	// `hostutl version --format {{.Client.Version}}` would be preferable, but
	// that fails if the server isn't running yet.
	//
	// output is expected to be something like
	//
	//     Bhojpur Host version 1.12.1, build 7a86f89
	output, err := ssh.SSHCommand("hostutl --version")
	if err != nil {
		return "", err
	}

	words := strings.Fields(output)
	if len(words) < 4 || words[0] != "Bhojpur" || words[2] != "version" {
		return "", fmt.Errorf("Bhojpur Host client version: cannot parse version string from %q", output)
	}

	return strings.TrimRight(words[3], ","), nil
}

func waitForLockAptGetUpdate(ssh SSHCommander) error {
	return waitForLock(ssh, "sudo apt-get update")
}

func waitForLock(ssh SSHCommander, cmd string) error {
	var sshErr error
	err := mutils.WaitFor(func() bool {
		_, sshErr = ssh.SSHCommand(cmd)
		if sshErr != nil {
			if strings.Contains(sshErr.Error(), "Could not get lock") {
				sshErr = nil
				return false
			}
			return true
		}
		return true
	})
	if sshErr != nil {
		return fmt.Errorf("Error running %q: %s", cmd, sshErr)
	}
	if err != nil {
		return fmt.Errorf("Failed to obtain lock: %s", err)
	}
	return nil
}

package generic

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
	"errors"
	"fmt"
	"net"
	"os"
	"path"
	"strconv"
	"time"

	"github.com/bhojpur/host/pkg/machine/drivers"
	"github.com/bhojpur/host/pkg/machine/engine"
	mflag "github.com/bhojpur/host/pkg/machine/flag"
	"github.com/bhojpur/host/pkg/machine/log"
	"github.com/bhojpur/host/pkg/machine/state"
	mutils "github.com/bhojpur/host/pkg/machine/utils"
)

type Driver struct {
	*drivers.BaseDriver
	EnginePort int
	SSHKey     string
}

const (
	defaultTimeout = 15 * time.Second
)

// GetCreateFlags registers the flags this driver adds to
// "hostutl hosts create"
func (d *Driver) GetCreateFlags() []mflag.Flag {
	return []mflag.Flag{
		mflag.IntFlag{
			Name:   "generic-engine-port",
			Usage:  "Bhojpur Host engine port",
			Value:  engine.DefaultPort,
			EnvVar: "GENERIC_ENGINE_PORT",
		},
		mflag.StringFlag{
			Name:   "generic-ip-address",
			Usage:  "IP Address of machine",
			EnvVar: "GENERIC_IP_ADDRESS",
		},
		mflag.StringFlag{
			Name:   "generic-ssh-user",
			Usage:  "SSH user",
			Value:  drivers.DefaultSSHUser,
			EnvVar: "GENERIC_SSH_USER",
		},
		mflag.StringFlag{
			Name:   "generic-ssh-key",
			Usage:  "SSH private key path (if not provided, default SSH key will be used)",
			Value:  "",
			EnvVar: "GENERIC_SSH_KEY",
		},
		mflag.IntFlag{
			Name:   "generic-ssh-port",
			Usage:  "SSH port",
			Value:  drivers.DefaultSSHPort,
			EnvVar: "GENERIC_SSH_PORT",
		},
	}
}

// NewDriver creates and returns a new instance of the driver
func NewDriver(hostName, storePath string) drivers.Driver {
	return &Driver{
		EnginePort: engine.DefaultPort,
		BaseDriver: &drivers.BaseDriver{
			MachineName: hostName,
			StorePath:   storePath,
		},
	}
}

// DriverName returns the name of the driver
func (d *Driver) DriverName() string {
	return "generic"
}

func (d *Driver) GetSSHHostname() (string, error) {
	return d.GetIP()
}

func (d *Driver) GetSSHUsername() string {
	return d.SSHUser
}

func (d *Driver) GetSSHKeyPath() string {
	return d.SSHKeyPath
}

func (d *Driver) SetConfigFromFlags(flags drivers.DriverOptions) error {
	d.EnginePort = flags.Int("generic-engine-port")
	d.IPAddress = flags.String("generic-ip-address")
	d.SSHUser = flags.String("generic-ssh-user")
	d.SSHKey = flags.String("generic-ssh-key")
	d.SSHPort = flags.Int("generic-ssh-port")

	if d.IPAddress == "" {
		return errors.New("generic driver requires the --generic-ip-address option")
	}

	return nil
}

func (d *Driver) PreCreateCheck() error {
	if d.SSHKey != "" {
		if _, err := os.Stat(d.SSHKey); os.IsNotExist(err) {
			return fmt.Errorf("SSH key does not exist: %q", d.SSHKey)
		}

		// TODO: validate the key is a valid key
	}

	return nil
}

func (d *Driver) Create() error {
	if d.SSHKey == "" {
		log.Info("No SSH key specified. Assuming an existing key at the default location.")
	} else {
		log.Info("Importing SSH key...")

		d.SSHKeyPath = d.ResolveStorePath(path.Base(d.SSHKey))
		if err := copySSHKey(d.SSHKey, d.SSHKeyPath); err != nil {
			return err
		}

		if err := copySSHKey(d.SSHKey+".pub", d.SSHKeyPath+".pub"); err != nil {
			log.Infof("Couldn't copy SSH public key : %s", err)
		}
	}

	log.Debugf("IP: %s", d.IPAddress)

	return nil
}

func (d *Driver) GetURL() (string, error) {
	if err := drivers.MustBeRunning(d); err != nil {
		return "", err
	}

	ip, err := d.GetIP()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("tcp://%s", net.JoinHostPort(ip, strconv.Itoa(d.EnginePort))), nil
}

func (d *Driver) GetState() (state.State, error) {
	address := net.JoinHostPort(d.IPAddress, strconv.Itoa(d.SSHPort))

	_, err := net.DialTimeout("tcp", address, defaultTimeout)
	if err != nil {
		return state.Stopped, nil
	}

	return state.Running, nil
}

func (d *Driver) Start() error {
	return errors.New("generic driver does not support start")
}

func (d *Driver) Stop() error {
	return errors.New("generic driver does not support stop")
}

func (d *Driver) Restart() error {
	_, err := drivers.RunSSHCommandFromDriver(d, "sudo shutdown -r now")
	return err
}

func (d *Driver) Kill() error {
	return errors.New("generic driver does not support kill")
}

func (d *Driver) Remove() error {
	return nil
}

func copySSHKey(src, dst string) error {
	if err := mutils.CopyFile(src, dst); err != nil {
		return fmt.Errorf("unable to copy ssh key: %s", err)
	}

	if err := os.Chmod(dst, 0600); err != nil {
		return fmt.Errorf("unable to set permissions on the ssh key: %s", err)
	}

	return nil
}

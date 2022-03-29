package google

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
	"io/ioutil"
	"net"
	"strings"

	"github.com/bhojpur/host/pkg/core/drivers"
	mflag "github.com/bhojpur/host/pkg/core/flag"
	"github.com/bhojpur/host/pkg/core/log"
	"github.com/bhojpur/host/pkg/core/ssh"
	"github.com/bhojpur/host/pkg/core/state"
)

// Driver is a struct compatible with the bhojpur.hosts.drivers.Driver interface.
type Driver struct {
	*drivers.BaseDriver
	Auth              string
	Zone              string
	MachineType       string
	MachineImage      string
	DiskType          string
	Address           string
	Network           string
	Subnetwork        string
	Preemptible       bool
	UseInternalIP     bool
	UseInternalIPOnly bool
	Scopes            string
	DiskSize          int
	Project           string
	Tags              string
	UseExisting       bool
	OpenPorts         []string
	Userdata          string
}

const (
	defaultZone        = "us-central1-a"
	defaultUser        = "bhojpur-user"
	defaultMachineType = "n1-standard-1"
	defaultImageName   = "ubuntu-os-cloud/global/images/ubuntu-1604-xenial-v20170721"
	defaultScopes      = "https://www.googleapis.com/auth/devstorage.read_only,https://www.googleapis.com/auth/logging.write,https://www.googleapis.com/auth/monitoring.write"
	defaultDiskType    = "pd-standard"
	defaultDiskSize    = 10
	defaultNetwork     = "default"
	defaultSubnetwork  = ""
)

// GetCreateFlags registers the flags this driver adds to
// "hostutl hosts create"
func (d *Driver) GetCreateFlags() []mflag.Flag {
	return []mflag.Flag{
		mflag.StringFlag{
			Name:   "google-zone",
			Usage:  "GCE Zone",
			Value:  defaultZone,
			EnvVar: "GOOGLE_ZONE",
		},
		mflag.StringFlag{
			Name:   "google-machine-type",
			Usage:  "GCE Machine Type",
			Value:  defaultMachineType,
			EnvVar: "GOOGLE_MACHINE_TYPE",
		},
		mflag.StringFlag{
			Name:   "google-machine-image",
			Usage:  "GCE Machine Image Absolute URL",
			Value:  defaultImageName,
			EnvVar: "GOOGLE_MACHINE_IMAGE",
		},
		mflag.StringFlag{
			Name:   "google-username",
			Usage:  "GCE User Name",
			Value:  defaultUser,
			EnvVar: "GOOGLE_USERNAME",
		},
		mflag.StringFlag{
			Name:   "google-auth-encoded-json",
			Usage:  "Base64 encoded GCE auth json",
			EnvVar: "GOOGLE_APPLICATION_CREDENTIALS_ENCODED_JSON",
		},
		mflag.StringFlag{
			Name:   "google-project",
			Usage:  "GCE Project",
			EnvVar: "GOOGLE_PROJECT",
		},
		mflag.StringFlag{
			Name:   "google-scopes",
			Usage:  "GCE Scopes (comma-separated if multiple scopes)",
			Value:  defaultScopes,
			EnvVar: "GOOGLE_SCOPES",
		},
		mflag.IntFlag{
			Name:   "google-disk-size",
			Usage:  "GCE Instance Disk Size (in GB)",
			Value:  defaultDiskSize,
			EnvVar: "GOOGLE_DISK_SIZE",
		},
		mflag.StringFlag{
			Name:   "google-disk-type",
			Usage:  "GCE Instance Disk type",
			Value:  defaultDiskType,
			EnvVar: "GOOGLE_DISK_TYPE",
		},
		mflag.StringFlag{
			Name:   "google-network",
			Usage:  "Specify network in which to provision vm",
			Value:  defaultNetwork,
			EnvVar: "GOOGLE_NETWORK",
		},
		mflag.StringFlag{
			Name:   "google-subnetwork",
			Usage:  "Specify subnetwork in which to provision vm",
			Value:  defaultSubnetwork,
			EnvVar: "GOOGLE_SUBNETWORK",
		},
		mflag.StringFlag{
			Name:   "google-address",
			Usage:  "GCE Instance External IP",
			EnvVar: "GOOGLE_ADDRESS",
		},
		mflag.BoolFlag{
			Name:   "google-preemptible",
			Usage:  "GCE Instance Preemptibility",
			EnvVar: "GOOGLE_PREEMPTIBLE",
		},
		mflag.StringFlag{
			Name:   "google-tags",
			Usage:  "GCE Instance Tags (comma-separated)",
			EnvVar: "GOOGLE_TAGS",
			Value:  "",
		},
		mflag.BoolFlag{
			Name:   "google-use-internal-ip",
			Usage:  "Use internal GCE Instance IP rather than public one",
			EnvVar: "GOOGLE_USE_INTERNAL_IP",
		},
		mflag.BoolFlag{
			Name:   "google-use-internal-ip-only",
			Usage:  "Configure GCE instance to not have an external IP address",
			EnvVar: "GOOGLE_USE_INTERNAL_IP_ONLY",
		},
		mflag.BoolFlag{
			Name:   "google-use-existing",
			Usage:  "Don't create a new VM, use an existing one",
			EnvVar: "GOOGLE_USE_EXISTING",
		},
		mflag.StringSliceFlag{
			Name:  "google-open-port",
			Usage: "Make the specified port number accessible from the Internet, e.g, 8080/tcp",
		},
		mflag.StringFlag{
			Name:   "google-userdata",
			Usage:  "A user-data file to be passed to cloud-init",
			EnvVar: "GOOGLE_USERDATA",
			Value:  "",
		},
	}
}

// NewDriver creates a Driver with the specified storePath.
func NewDriver(machineName string, storePath string) *Driver {
	return &Driver{
		Zone:         defaultZone,
		DiskType:     defaultDiskType,
		DiskSize:     defaultDiskSize,
		MachineType:  defaultMachineType,
		MachineImage: defaultImageName,
		Network:      defaultNetwork,
		Subnetwork:   defaultSubnetwork,
		Scopes:       defaultScopes,
		BaseDriver: &drivers.BaseDriver{
			SSHUser:     defaultUser,
			MachineName: machineName,
			StorePath:   storePath,
		},
	}
}

// GetSSHHostname returns hostname for use with ssh
func (d *Driver) GetSSHHostname() (string, error) {
	return d.GetIP()
}

// GetSSHUsername returns username for use with ssh
func (d *Driver) GetSSHUsername() string {
	if d.SSHUser == "" {
		d.SSHUser = "bhojpur-user"
	}
	return d.SSHUser
}

// DriverName returns the name of the driver
func (d *Driver) DriverName() string {
	return "google"
}

// SetConfigFromFlags initializes the driver based on the command line flags.
func (d *Driver) SetConfigFromFlags(flags drivers.DriverOptions) error {
	d.Project = flags.String("google-project")
	if d.Project == "" {
		return errors.New("no Google Cloud Project name specified (--google-project)")
	}
	d.Auth = flags.String("google-auth-encoded-json")

	d.Zone = flags.String("google-zone")
	d.UseExisting = flags.Bool("google-use-existing")
	if !d.UseExisting {
		d.MachineType = flags.String("google-machine-type")
		d.MachineImage = flags.String("google-machine-image")
		d.MachineImage = strings.TrimPrefix(d.MachineImage, "https://www.googleapis.com/compute/v1/projects/")
		d.DiskSize = flags.Int("google-disk-size")
		d.DiskType = flags.String("google-disk-type")
		d.Address = flags.String("google-address")
		d.Network = flags.String("google-network")
		d.Subnetwork = flags.String("google-subnetwork")
		d.Preemptible = flags.Bool("google-preemptible")
		d.UseInternalIP = flags.Bool("google-use-internal-ip") || flags.Bool("google-use-internal-ip-only")
		d.UseInternalIPOnly = flags.Bool("google-use-internal-ip-only")
		d.Scopes = flags.String("google-scopes")
		d.Tags = flags.String("google-tags")
		d.OpenPorts = flags.StringSlice("google-open-port")
	}
	d.SSHUser = flags.String("google-username")
	d.SSHPort = 22
	d.Userdata = flags.String("google-userdata")
	d.SetSwarmConfigFromFlags(flags)

	return nil
}

// PreCreateCheck is called to enforce pre-creation steps
func (d *Driver) PreCreateCheck() error {
	c, err := newComputeUtil(d)
	if err != nil {
		return err
	}

	// Check that the project exists. It will also check the credentials
	// at the same time.
	log.Infof("Check that the project exists")

	if _, err = c.service.Projects.Get(d.Project).Do(); err != nil {
		return fmt.Errorf("Project with ID %q not found. %v", d.Project, err)
	}

	// Check if the instance already exists. There will be an error if the instance
	// doesn't exist, so just check instance for nil.
	log.Infof("Check if the instance already exists")

	instance, _ := c.instance()
	if d.UseExisting {
		if instance == nil {
			return fmt.Errorf("unable to find instance %q in zone %q", d.MachineName, d.Zone)
		}
	} else {
		if instance != nil {
			return fmt.Errorf("instance %q already exists in zone %q", d.MachineName, d.Zone)
		}
	}

	if d.Userdata != "" {
		file, err := ioutil.ReadFile(d.Userdata)
		if err != nil {
			return fmt.Errorf("cannot read userdata file %v: %v", d.Userdata, err)
		}
		d.Userdata = string(file)
	}

	return nil
}

// Create creates a GCE VM instance acting as a Bhojpur Host.
func (d *Driver) Create() error {
	log.Infof("Generating SSH Key")

	if err := ssh.GenerateSSHKey(d.GetSSHKeyPath()); err != nil {
		return err
	}

	log.Infof("Creating host...")

	c, err := newComputeUtil(d)
	if err != nil {
		return err
	}

	if err := c.openFirewallPorts(d); err != nil {
		return err
	}

	if d.UseExisting {
		return c.configureInstance(d)
	}
	return c.createInstance(d)
}

// GetURL returns the URL of the remote Bhojpur Host daemon.
func (d *Driver) GetURL() (string, error) {
	ip, err := d.GetIP()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("tcp://%s", net.JoinHostPort(ip, "2376")), nil
}

// GetIP returns the IP address of the GCE instance.
func (d *Driver) GetIP() (string, error) {
	c, err := newComputeUtil(d)
	if err != nil {
		return "", err
	}

	ip, err := c.ip()
	if err != nil {
		return "", err
	}
	if ip == "" {
		return "", drivers.ErrHostIsNotRunning
	}

	return ip, nil
}

// GetState returns a bhojpur.hosts.state.State value representing the current state of the host.
func (d *Driver) GetState() (state.State, error) {
	c, err := newComputeUtil(d)
	if err != nil {
		return state.None, err
	}

	// All we care about is whether the disk exists, so we just check disk for a nil value.
	// There will be no error if disk is not nil.
	instance, err := c.instance()
	if instance == nil {
		if err != nil && strings.Contains(err.Error(), "not found") {
			return state.NotFound, nil
		}
		disk, _ := c.disk()
		if disk == nil {
			return state.None, nil
		}
		return state.Stopped, nil
	}

	switch instance.Status {
	case "PROVISIONING", "STAGING":
		return state.Starting, nil
	case "RUNNING":
		return state.Running, nil
	case "STOPPING", "STOPPED", "TERMINATED":
		return state.Stopped, nil
	}
	return state.None, nil
}

// Start starts an existing GCE instance or create an instance with an existing disk.
func (d *Driver) Start() error {
	c, err := newComputeUtil(d)
	if err != nil {
		return err
	}

	instance, err := c.instance()
	if err != nil {
		if !isNotFound(err) {
			return err
		}
	}

	if instance == nil {
		if err = c.createInstance(d); err != nil {
			return err
		}
	} else {
		if err := c.startInstance(); err != nil {
			return err
		}
	}

	d.IPAddress, err = d.GetIP()
	return err
}

// Stop stops an existing GCE instance.
func (d *Driver) Stop() error {
	c, err := newComputeUtil(d)
	if err != nil {
		return err
	}

	if err := c.stopInstance(); err != nil {
		return err
	}

	d.IPAddress = ""
	return nil
}

// Restart restarts a machine which is known to be running.
func (d *Driver) Restart() error {
	if err := d.Stop(); err != nil {
		return err
	}

	return d.Start()
}

// Kill stops an existing GCE instance.
func (d *Driver) Kill() error {
	return d.Stop()
}

// Remove deletes the GCE instance and the disk.
func (d *Driver) Remove() error {
	c, err := newComputeUtil(d)
	if err != nil {
		return err
	}

	if err := c.deleteInstance(); err != nil {
		if isNotFound(err) {
			log.Warn("Remote instance does not exist, proceeding with removing local reference")
		} else {
			return err
		}
	}

	if err := c.deleteDisk(); err != nil {
		if isNotFound(err) {
			log.Warn("Remote disk does not exist, proceeding")
		} else {
			return err
		}
	}

	return nil
}

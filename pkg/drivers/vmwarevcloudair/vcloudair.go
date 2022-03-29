package vmwarevcloudair

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
	"net"
	"strconv"
	"strings"

	"github.com/vmware/govcloudair"

	"github.com/bhojpur/host/pkg/core/drivers"
	mflag "github.com/bhojpur/host/pkg/core/flag"
	"github.com/bhojpur/host/pkg/core/log"
	"github.com/bhojpur/host/pkg/core/ssh"
	"github.com/bhojpur/host/pkg/core/state"
	mutils "github.com/bhojpur/host/pkg/core/utils"
)

type Driver struct {
	*drivers.BaseDriver
	UserName     string
	UserPassword string
	ComputeID    string
	VDCID        string
	OrgVDCNet    string
	EdgeGateway  string
	PublicIP     string
	Catalog      string
	CatalogItem  string
	BhojpurPort  int
	CPUCount     int
	MemorySize   int
	VAppID       string
}

const (
	defaultCatalog     = "Public Catalog"
	defaultCatalogItem = "Ubuntu Server 12.04 LTS (amd64 20150127)"
	defaultCpus        = 1
	defaultMemory      = 2048
	defaultSSHPort     = 22
	defaultBhojpurPort = 2376
)

// GetCreateFlags registers the flags this driver adds to
// "hostutl hosts create"
func (d *Driver) GetCreateFlags() []mflag.Flag {
	return []mflag.Flag{
		mflag.StringFlag{
			EnvVar: "VCLOUDAIR_USERNAME",
			Name:   "vmwarevcloudair-username",
			Usage:  "vCloud Air username",
		},
		mflag.StringFlag{
			EnvVar: "VCLOUDAIR_PASSWORD",
			Name:   "vmwarevcloudair-password",
			Usage:  "vCloud Air password",
		},
		mflag.StringFlag{
			EnvVar: "VCLOUDAIR_COMPUTEID",
			Name:   "vmwarevcloudair-computeid",
			Usage:  "vCloud Air Compute ID (if using Dedicated Cloud)",
		},
		mflag.StringFlag{
			EnvVar: "VCLOUDAIR_VDCID",
			Name:   "vmwarevcloudair-vdcid",
			Usage:  "vCloud Air VDC ID",
		},
		mflag.StringFlag{
			EnvVar: "VCLOUDAIR_ORGVDCNETWORK",
			Name:   "vmwarevcloudair-orgvdcnetwork",
			Usage:  "vCloud Air Org VDC Network (Default is <vdcid>-default-routed)",
		},
		mflag.StringFlag{
			EnvVar: "VCLOUDAIR_EDGEGATEWAY",
			Name:   "vmwarevcloudair-edgegateway",
			Usage:  "vCloud Air Org Edge Gateway (Default is <vdcid>)",
		},
		mflag.StringFlag{
			EnvVar: "VCLOUDAIR_PUBLICIP",
			Name:   "vmwarevcloudair-publicip",
			Usage:  "vCloud Air Org Public IP to use",
		},
		mflag.StringFlag{
			EnvVar: "VCLOUDAIR_CATALOG",
			Name:   "vmwarevcloudair-catalog",
			Usage:  "vCloud Air Catalog (default is Public Catalog)",
			Value:  defaultCatalog,
		},
		mflag.StringFlag{
			EnvVar: "VCLOUDAIR_CATALOGITEM",
			Name:   "vmwarevcloudair-catalogitem",
			Usage:  "vCloud Air Catalog Item (default is Ubuntu Precise)",
			Value:  defaultCatalogItem,
		},
		mflag.IntFlag{
			EnvVar: "VCLOUDAIR_CPU_COUNT",
			Name:   "vmwarevcloudair-cpu-count",
			Usage:  "vCloud Air VM Cpu Count (default 1)",
			Value:  defaultCpus,
		},
		mflag.IntFlag{
			EnvVar: "VCLOUDAIR_MEMORY_SIZE",
			Name:   "vmwarevcloudair-memory-size",
			Usage:  "vCloud Air VM Memory Size in MB (default 2048)",
			Value:  defaultMemory,
		},
		mflag.IntFlag{
			EnvVar: "VCLOUDAIR_SSH_PORT",
			Name:   "vmwarevcloudair-ssh-port",
			Usage:  "vCloud Air SSH port",
			Value:  defaultSSHPort,
		},
		mflag.IntFlag{
			EnvVar: "VCLOUDAIR_BHOJPUR_PORT",
			Name:   "vmwarevcloudair-bhojpur-port",
			Usage:  "vCloud Air Bhojpur port",
			Value:  defaultBhojpurPort,
		},
	}
}

func NewDriver(hostName, storePath string) drivers.Driver {
	return &Driver{
		Catalog:     defaultCatalog,
		CatalogItem: defaultCatalogItem,
		CPUCount:    defaultCpus,
		MemorySize:  defaultMemory,
		BhojpurPort: defaultBhojpurPort,
		BaseDriver: &drivers.BaseDriver{
			SSHPort:     defaultSSHPort,
			MachineName: hostName,
			StorePath:   storePath,
		},
	}
}

func (d *Driver) GetSSHHostname() (string, error) {
	return d.GetIP()
}

// DriverName returns the name of the driver
func (d *Driver) DriverName() string {
	return "vmwarevcloudair"
}

func (d *Driver) SetConfigFromFlags(flags drivers.DriverOptions) error {

	d.UserName = flags.String("vmwarevcloudair-username")
	d.UserPassword = flags.String("vmwarevcloudair-password")
	d.VDCID = flags.String("vmwarevcloudair-vdcid")
	d.PublicIP = flags.String("vmwarevcloudair-publicip")
	d.SetSwarmConfigFromFlags(flags)

	// Check for required Params
	if d.UserName == "" || d.UserPassword == "" || d.VDCID == "" || d.PublicIP == "" {
		return fmt.Errorf("Please specify vcloudair mandatory params using options: -vmwarevcloudair-username -vmwarevcloudair-password -vmwarevcloudair-vdcid and -vmwarevcloudair-publicip")
	}

	// If ComputeID is not set we're using a VPC, hence setting ComputeID = VDCID
	if flags.String("vmwarevcloudair-computeid") == "" {
		d.ComputeID = flags.String("vmwarevcloudair-vdcid")
	} else {
		d.ComputeID = flags.String("vmwarevcloudair-computeid")
	}

	// If the Org VDC Network is empty, set it to the default routed network.
	if flags.String("vmwarevcloudair-orgvdcnetwork") == "" {
		d.OrgVDCNet = flags.String("vmwarevcloudair-vdcid") + "-default-routed"
	} else {
		d.OrgVDCNet = flags.String("vmwarevcloudair-orgvdcnetwork")
	}

	// If the Edge Gateway is empty, just set it to the default edge gateway.
	if flags.String("vmwarevcloudair-edgegateway") == "" {
		d.EdgeGateway = flags.String("vmwarevcloudair-vdcid")
	} else {
		d.EdgeGateway = flags.String("vmwarevcloudair-edgegateway")
	}

	d.Catalog = flags.String("vmwarevcloudair-catalog")
	d.CatalogItem = flags.String("vmwarevcloudair-catalogitem")

	d.BhojpurPort = flags.Int("vmwarevcloudair-bhojpur-port")
	d.SSHUser = "root"
	d.SSHPort = flags.Int("vmwarevcloudair-ssh-port")
	d.CPUCount = flags.Int("vmwarevcloudair-cpu-count")
	d.MemorySize = flags.Int("vmwarevcloudair-memory-size")

	return nil
}

func (d *Driver) GetURL() (string, error) {
	if err := drivers.MustBeRunning(d); err != nil {
		return "", err
	}

	return fmt.Sprintf("tcp://%s", net.JoinHostPort(d.PublicIP, strconv.Itoa(d.BhojpurPort))), nil
}

func (d *Driver) GetIP() (string, error) {
	return d.PublicIP, nil
}

func (d *Driver) GetState() (state.State, error) {
	p, err := govcloudair.NewClient()
	if err != nil {
		return state.Error, err
	}

	log.Debug("Connecting to vCloud Air to fetch vApp Status...")
	// Authenticate to vCloud Air
	v, err := p.Authenticate(d.UserName, d.UserPassword, d.ComputeID, d.VDCID)
	if err != nil {
		return state.Error, err
	}

	vapp, err := v.FindVAppByID(d.VAppID)
	if err != nil {
		return state.Error, err
	}

	status, err := vapp.GetStatus()
	if err != nil {
		return state.Error, err
	}

	if err = p.Disconnect(); err != nil {
		return state.Error, err
	}

	switch status {
	case "POWERED_ON":
		return state.Running, nil
	case "POWERED_OFF":
		return state.Stopped, nil
	}
	return state.None, nil
}

func (d *Driver) Create() error {
	key, err := d.createSSHKey()
	if err != nil {
		return err
	}

	p, err := govcloudair.NewClient()
	if err != nil {
		return err
	}

	log.Infof("Connecting to vCloud Air...")
	// Authenticate to vCloud Air
	v, err := p.Authenticate(d.UserName, d.UserPassword, d.ComputeID, d.VDCID)
	if err != nil {
		return err
	}

	// Find VDC Network
	net, err := v.FindVDCNetwork(d.OrgVDCNet)
	if err != nil {
		return err
	}

	// Find our Edge Gateway
	edge, err := v.FindEdgeGateway(d.EdgeGateway)
	if err != nil {
		return err
	}

	// Get the Org our VDC belongs to
	org, err := v.GetVDCOrg()
	if err != nil {
		return err
	}

	// Find our Catalog
	cat, err := org.FindCatalog(d.Catalog)
	if err != nil {
		return err
	}

	// Find our Catalog Item
	cati, err := cat.FindCatalogItem(d.CatalogItem)
	if err != nil {
		return err
	}

	// Fetch the vApp Template in the Catalog Item
	vapptemplate, err := cati.GetVAppTemplate()
	if err != nil {
		return err
	}

	// Create a new empty vApp
	vapp := govcloudair.NewVApp(p)

	log.Infof("Creating a new vApp: %s...", d.MachineName)
	// Compose the vApp with ComposeVApp
	task, err := vapp.ComposeVApp(net, vapptemplate, d.MachineName, "Container Host created with Bhojpur Host")
	if err != nil {
		return err
	}

	// Wait for the creation to be completed
	if err = task.WaitTaskCompletion(); err != nil {
		return err
	}

	task, err = vapp.ChangeCPUcount(d.CPUCount)
	if err != nil {
		return err
	}

	if err = task.WaitTaskCompletion(); err != nil {
		return err
	}

	task, err = vapp.ChangeMemorySize(d.MemorySize)
	if err != nil {
		return err
	}

	if err = task.WaitTaskCompletion(); err != nil {
		return err
	}

	sshCustomScript := "echo \"" + strings.TrimSpace(key) + "\" > /root/.ssh/authorized_keys"

	task, err = vapp.RunCustomizationScript(d.MachineName, sshCustomScript)
	if err != nil {
		return err
	}

	if err = task.WaitTaskCompletion(); err != nil {
		return err
	}

	task, err = vapp.PowerOn()
	if err != nil {
		return err
	}

	log.Infof("Waiting for the VM to power on and run the customization script...")

	if err = task.WaitTaskCompletion(); err != nil {
		return err
	}

	log.Infof("Creating NAT and Firewall Rules on %s...", d.EdgeGateway)
	task, err = edge.Create1to1Mapping(vapp.VApp.Children.VM[0].NetworkConnectionSection.NetworkConnection.IPAddress, d.PublicIP, d.MachineName)
	if err != nil {
		return err
	}

	if err = task.WaitTaskCompletion(); err != nil {
		return err
	}

	log.Debugf("Disconnecting from vCloud Air...")

	if err = p.Disconnect(); err != nil {
		return err
	}

	// Set VAppID with ID of the created VApp
	d.VAppID = vapp.VApp.ID

	d.IPAddress, err = d.GetIP()
	return err
}

func (d *Driver) Remove() error {
	p, err := govcloudair.NewClient()
	if err != nil {
		return err
	}

	log.Infof("Connecting to vCloud Air...")
	// Authenticate to vCloud Air
	v, err := p.Authenticate(d.UserName, d.UserPassword, d.ComputeID, d.VDCID)
	if err != nil {
		return err
	}

	// Find our Edge Gateway
	edge, err := v.FindEdgeGateway(d.EdgeGateway)
	if err != nil {
		return err
	}

	vapp, err := v.FindVAppByID(d.VAppID)
	if err != nil {
		log.Infof("Can't find the vApp, assuming it was deleted already...")
		return nil
	}

	status, err := vapp.GetStatus()
	if err != nil {
		return err
	}

	log.Infof("Removing NAT and Firewall Rules on %s...", d.EdgeGateway)
	task, err := edge.Remove1to1Mapping(vapp.VApp.Children.VM[0].NetworkConnectionSection.NetworkConnection.IPAddress, d.PublicIP)
	if err != nil {
		return err
	}
	if err = task.WaitTaskCompletion(); err != nil {
		return err
	}

	if status == "POWERED_ON" {
		// If it's powered on, power it off before deleting
		log.Infof("Powering Off %s...", d.MachineName)
		task, err = vapp.PowerOff()
		if err != nil {
			return err
		}
		if err = task.WaitTaskCompletion(); err != nil {
			return err
		}

	}

	log.Debugf("Undeploying %s...", d.MachineName)
	task, err = vapp.Undeploy()
	if err != nil {
		return err
	}
	if err = task.WaitTaskCompletion(); err != nil {
		return err
	}

	log.Infof("Deleting %s...", d.MachineName)
	task, err = vapp.Delete()
	if err != nil {
		return err
	}
	if err = task.WaitTaskCompletion(); err != nil {
		return err
	}

	err = p.Disconnect()
	return err
}

func (d *Driver) Start() error {
	p, err := govcloudair.NewClient()
	if err != nil {
		return err
	}

	log.Infof("Connecting to vCloud Air...")
	// Authenticate to vCloud Air
	v, err := p.Authenticate(d.UserName, d.UserPassword, d.ComputeID, d.VDCID)
	if err != nil {
		return err
	}

	vapp, err := v.FindVAppByID(d.VAppID)
	if err != nil {
		return err
	}

	status, err := vapp.GetStatus()
	if err != nil {
		return err
	}

	if status == "POWERED_OFF" {
		log.Infof("Starting %s...", d.MachineName)
		task, err := vapp.PowerOn()
		if err != nil {
			return err
		}
		if err = task.WaitTaskCompletion(); err != nil {
			return err
		}

	}

	if err = p.Disconnect(); err != nil {
		return err
	}

	d.IPAddress, err = d.GetIP()
	return err
}

func (d *Driver) Stop() error {
	p, err := govcloudair.NewClient()
	if err != nil {
		return err
	}

	log.Infof("Connecting to vCloud Air...")
	// Authenticate to vCloud Air
	v, err := p.Authenticate(d.UserName, d.UserPassword, d.ComputeID, d.VDCID)
	if err != nil {
		return err
	}

	vapp, err := v.FindVAppByID(d.VAppID)
	if err != nil {
		return err
	}

	task, err := vapp.Shutdown()
	if err != nil {
		return err
	}
	if err = task.WaitTaskCompletion(); err != nil {
		return err
	}

	if err = p.Disconnect(); err != nil {
		return err
	}

	d.IPAddress = ""

	return nil
}

func (d *Driver) Restart() error {
	p, err := govcloudair.NewClient()
	if err != nil {
		return err
	}

	log.Infof("Connecting to vCloud Air...")
	// Authenticate to vCloud Air
	v, err := p.Authenticate(d.UserName, d.UserPassword, d.ComputeID, d.VDCID)
	if err != nil {
		return err
	}

	vapp, err := v.FindVAppByID(d.VAppID)
	if err != nil {
		return err
	}

	task, err := vapp.Reset()
	if err != nil {
		return err
	}
	if err = task.WaitTaskCompletion(); err != nil {
		return err
	}

	if err = p.Disconnect(); err != nil {
		return err
	}

	d.IPAddress, err = d.GetIP()
	return err
}

func (d *Driver) Kill() error {
	p, err := govcloudair.NewClient()
	if err != nil {
		return err
	}

	log.Infof("Connecting to vCloud Air...")
	// Authenticate to vCloud Air
	v, err := p.Authenticate(d.UserName, d.UserPassword, d.ComputeID, d.VDCID)
	if err != nil {
		return err
	}

	vapp, err := v.FindVAppByID(d.VAppID)
	if err != nil {
		return err
	}

	task, err := vapp.PowerOff()
	if err != nil {
		return err
	}
	if err = task.WaitTaskCompletion(); err != nil {
		return err
	}

	if err = p.Disconnect(); err != nil {
		return err
	}

	d.IPAddress = ""

	return nil
}

// Helpers

func generateVMName() string {
	randomID := mutils.TruncateID(mutils.GenerateRandomID())
	return fmt.Sprintf("bhojpur-host-%s", randomID)
}

func (d *Driver) createSSHKey() (string, error) {
	if err := ssh.GenerateSSHKey(d.GetSSHKeyPath()); err != nil {
		return "", err
	}

	publicKey, err := ioutil.ReadFile(d.publicSSHKeyPath())
	if err != nil {
		return "", err
	}

	return string(publicKey), nil
}

func (d *Driver) publicSSHKeyPath() string {
	return d.GetSSHKeyPath() + ".pub"
}
package vmwarevsphere

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
	"os"
	"strings"

	"github.com/bhojpur/host/pkg/core/drivers"
	mflag "github.com/bhojpur/host/pkg/core/flag"
	"github.com/vmware/govmomi/vim25/types"
)

var (
	supportedMachineOS = map[string]struct{}{
		"windows": {},
		"linux":   {},
	}
	supportedCreationTypes = map[string]struct{}{
		creationTypeVM:      {},
		creationTypeTmpl:    {},
		creationTypeLibrary: {},
		creationTypeLegacy:  {},
	}
)

func (d *Driver) GetCreateFlags() []mflag.Flag {
	return []mflag.Flag{
		mflag.IntFlag{
			EnvVar: "VSPHERE_CPU_COUNT",
			Name:   "vmwarevsphere-cpu-count",
			Usage:  "vSphere CPU number for Bhojpur Host VM",
			Value:  defaultCpus,
		},
		mflag.IntFlag{
			EnvVar: "VSPHERE_MEMORY_SIZE",
			Name:   "vmwarevsphere-memory-size",
			Usage:  "vSphere size of memory for Bhojpur Host VM (in MB)",
			Value:  defaultMemory,
		},
		mflag.IntFlag{
			EnvVar: "VSPHERE_DISK_SIZE",
			Name:   "vmwarevsphere-disk-size",
			Usage:  "vSphere size of disk for Bhojpur Host VM (in MB)",
			Value:  defaultDiskSize,
		},
		mflag.StringFlag{
			EnvVar: "VSPHERE_BOOT2DOCKER_URL",
			Name:   "vmwarevsphere-boot2docker-url",
			Usage:  "vSphere URL for boot2docker image",
		},
		mflag.StringFlag{
			EnvVar: "VSPHERE_VCENTER",
			Name:   "vmwarevsphere-vcenter",
			Usage:  "vSphere IP/hostname for vCenter",
		},
		mflag.IntFlag{
			EnvVar: "VSPHERE_VCENTER_PORT",
			Name:   "vmwarevsphere-vcenter-port",
			Usage:  "vSphere Port for vCenter",
			Value:  defaultSDKPort,
		},
		mflag.StringFlag{
			EnvVar: "VSPHERE_USERNAME",
			Name:   "vmwarevsphere-username",
			Usage:  "vSphere username",
		},
		mflag.StringFlag{
			EnvVar: "VSPHERE_PASSWORD",
			Name:   "vmwarevsphere-password",
			Usage:  "vSphere password",
		},
		mflag.StringSliceFlag{
			EnvVar: "VSPHERE_NETWORK",
			Name:   "vmwarevsphere-network",
			Usage:  "vSphere network where the virtual machine will be attached",
		},
		mflag.StringFlag{
			EnvVar: "VSPHERE_DATASTORE",
			Name:   "vmwarevsphere-datastore",
			Usage:  "vSphere datastore for virtual machine",
		},
		mflag.StringFlag{
			EnvVar: "VSPHERE_DATASTORE_CLUSTER",
			Name:   "vmwarevsphere-datastore-cluster",
			Usage:  "vSphere datastore cluster for virtual machine",
		},
		mflag.StringFlag{
			EnvVar: "VSPHERE_DATACENTER",
			Name:   "vmwarevsphere-datacenter",
			Usage:  "vSphere datacenter for virtual machine",
		},
		mflag.StringFlag{
			EnvVar: "VSPHERE_FOLDER",
			Name:   "vmwarevsphere-folder",
			Usage:  "vSphere folder for the Bhojpur Host VM. This folder must already exist in the datacenter",
		},
		mflag.StringFlag{
			EnvVar: "VSPHERE_POOL",
			Name:   "vmwarevsphere-pool",
			Usage:  "vSphere resource pool for Bhojpur Host VM",
		},
		mflag.StringFlag{
			EnvVar: "VSPHERE_HOSTSYSTEM",
			Name:   "vmwarevsphere-hostsystem",
			Usage:  "vSphere compute resource where the Bhojpur Host VM will be instantiated. This can be omitted if using a cluster with DRS",
		},
		mflag.StringSliceFlag{
			EnvVar: "VSPHERE_CFGPARAM",
			Name:   "vmwarevsphere-cfgparam",
			Usage:  "vSphere vm configuration parameters (used for guestinfo)",
		},
		mflag.StringFlag{
			EnvVar: "VSPHERE_CLOUDINIT",
			Name:   "vmwarevsphere-cloudinit",
			Usage:  "vSphere cloud-init filepath or url to add to guestinfo, filepath will be read and base64 encoded before adding",
		},
		mflag.StringFlag{
			EnvVar: "VSPHERE_CLOUD_CONFIG",
			Name:   "vmwarevsphere-cloud-config",
			Usage:  "Filepath to a cloud-config yaml file to put into the ISO user-data",
		},
		mflag.StringFlag{
			EnvVar: "VSPHERE_VAPP_IPPROTOCOL",
			Name:   "vmwarevsphere-vapp-ipprotocol",
			Usage:  "vSphere vApp IP protocol for this deployment. Supported values are: IPv4 and IPv6",
		},
		mflag.StringFlag{
			EnvVar: "VSPHERE_VAPP_IPALLOCATIONPOLICY",
			Name:   "vmwarevsphere-vapp-ipallocationpolicy",
			Usage:  "vSphere vApp IP allocation policy. Supported values are: dhcp, fixed, transient and fixedAllocated",
		},
		mflag.StringFlag{
			EnvVar: "VSPHERE_VAPP_TRANSPORT",
			Name:   "vmwarevsphere-vapp-transport",
			Usage:  "vSphere OVF environment transports to use for properties. Supported values are: iso and com.vmware.guestInfo",
		},
		mflag.StringSliceFlag{
			EnvVar: "VSPHERE_VAPP_PROPERTY",
			Name:   "vmwarevsphere-vapp-property",
			Usage:  "vSphere vApp properties",
		},
		mflag.StringFlag{
			EnvVar: "VSPHERE_CREATION_TYPE",
			Name:   "vmwarevsphere-creation-type",
			Usage:  "Creation type when creating a new virtual machine. Supported values: vm, template, library, legacy",
			Value:  creationTypeLegacy,
		},
		mflag.StringFlag{
			EnvVar: "VSPHERE_CLONE_FROM",
			Name:   "vmwarevsphere-clone-from",
			Usage:  "If you choose creation type clone a name of what you want to clone is required",
		},
		mflag.StringFlag{
			EnvVar: "VSPHERE_CONTENT_LIBRARY",
			Name:   "vmwarevsphere-content-library",
			Usage:  "If you choose to clone from a content library template specify the name of the library",
		},
		mflag.StringFlag{
			EnvVar: "VSPHERE_SSH_USER",
			Name:   "vmwarevsphere-ssh-user",
			Usage:  "If using a non-B2D image you can specify the ssh user",
			Value:  defaultSSHUser,
		},
		mflag.StringFlag{
			EnvVar: "VSPHERE_SSH_PASSWORD",
			Name:   "vmwarevsphere-ssh-password",
			Usage:  "If using a non-B2D image you can specify the ssh password",
			Value:  defaultSSHPass,
		},
		mflag.IntFlag{
			EnvVar: "VSPHERE_SSH_PORT",
			Name:   "vmwarevsphere-ssh-port",
			Usage:  "If using a non-B2D image you can specify the ssh port",
			Value:  drivers.DefaultSSHPort,
		},
		mflag.StringFlag{
			EnvVar: "VSPHERE_SSH_USER_GROUP",
			Name:   "vmwarevsphere-ssh-user-group",
			Usage:  "If using a non-B2D image the uploaded keys will need chown'ed, defaults to staff e.g. bhojpur:staff",
			Value:  defaultSSHUserGroup,
		},
		mflag.StringFlag{
			EnvVar: "",
			Name:   "vmwarevsphere-os",
			Usage:  "If using a non-B2D image you can specify the desired machine OS",
			Value:  defaultMachineOS,
		},
		mflag.StringSliceFlag{
			EnvVar: "",
			Name:   "vmwarevsphere-tag",
			Usage:  "vSphere tag id e.g. urn:xxx",
		},
		mflag.StringSliceFlag{
			EnvVar: "",
			Name:   "vmwarevsphere-custom-attribute",
			Usage:  "vSphere custom attribute, format key/value e.g. '200=my custom value'",
		},
	}
}

func (d *Driver) SetConfigFromFlags(flags drivers.DriverOptions) error {
	d.SSHUser = flags.String("vmwarevsphere-ssh-user")
	d.SSHPassword = flags.String("vmwarevsphere-ssh-password")
	d.SSHPort = flags.Int("vmwarevsphere-ssh-port")
	d.SSHUserGroup = flags.String("vmwarevsphere-ssh-user-group")
	d.CPU = flags.Int("vmwarevsphere-cpu-count")
	d.Memory = flags.Int("vmwarevsphere-memory-size")
	d.DiskSize = flags.Int("vmwarevsphere-disk-size")
	d.Boot2DockerURL = flags.String("vmwarevsphere-boot2docker-url")
	d.IP = flags.String("vmwarevsphere-vcenter")
	d.Port = flags.Int("vmwarevsphere-vcenter-port")
	d.Username = flags.String("vmwarevsphere-username")
	d.Password = flags.String("vmwarevsphere-password")
	d.Networks = flags.StringSlice("vmwarevsphere-network")
	d.Tags = flags.StringSlice("vmwarevsphere-tag")
	d.CustomAttributes = flags.StringSlice("vmwarevsphere-custom-attribute")
	d.Datastore = flags.String("vmwarevsphere-datastore")
	d.DatastoreCluster = flags.String("vmwarevsphere-datastore-cluster")
	d.Datacenter = flags.String("vmwarevsphere-datacenter")
	// Sanitize input on ingress.
	d.Folder = flags.String("vmwarevsphere-folder")
	d.Pool = flags.String("vmwarevsphere-pool")
	d.HostSystem = flags.String("vmwarevsphere-hostsystem")
	d.CfgParams = flags.StringSlice("vmwarevsphere-cfgparam")
	d.CloudInit = flags.String("vmwarevsphere-cloudinit")
	d.CloudConfig = flags.String("vmwarevsphere-cloud-config")
	if d.CloudConfig != "" {
		if _, err := os.Stat(d.CloudConfig); err != nil {
			return err
		}
		ud, err := ioutil.ReadFile(d.CloudConfig)
		if err != nil {
			return err
		}
		d.CloudConfig = string(ud)
	}

	d.VAppIpProtocol = flags.String("vmwarevsphere-vapp-ipprotocol")
	d.VAppIpAllocationPolicy = flags.String("vmwarevsphere-vapp-ipallocationpolicy")
	d.VAppTransport = flags.String("vmwarevsphere-vapp-transport")
	d.VAppProperties = flags.StringSlice("vmwarevsphere-vapp-property")
	d.SetSwarmConfigFromFlags(flags)
	d.ISO = d.ResolveStorePath(isoFilename)

	d.CreationType = flags.String("vmwarevsphere-creation-type")
	if _, ok := supportedCreationTypes[d.CreationType]; !ok {
		return fmt.Errorf("creation type %s not supported", d.CreationType)
	}
	err := d.SetMachineOSFromFlags(flags)
	if err != nil {
		return err
	}

	d.ContentLibrary = flags.String("vmwarevsphere-content-library")
	if d.CreationType != "legacy" {
		d.CloneFrom = flags.String("vmwarevsphere-clone-from")
		if d.CloneFrom == "" {
			return fmt.Errorf("creation type clone needs a VM name to clone from, use --vmwarevsphere-clone-from")
		}
	}

	return nil
}

type AuthFlag struct {
	auth types.NamePasswordAuthentication
}

func NewAuthFlag(u, p string) *AuthFlag {
	return &AuthFlag{
		auth: types.NamePasswordAuthentication{
			Username: u,
			Password: p,
		},
	}
}

func (f *AuthFlag) Auth() types.BaseGuestAuthentication {
	return &f.auth
}

type FileAttrFlag struct {
	types.GuestPosixFileAttributes
}

func (f *FileAttrFlag) SetPerms(owner, group, perms int) {
	owner32 := int32(owner)
	group32 := int32(group)
	f.OwnerId = &owner32
	f.GroupId = &group32
	f.Permissions = int64(perms)
}

func (f *FileAttrFlag) Attr() types.BaseGuestFileAttributes {
	return &f.GuestPosixFileAttributes
}

func (d *Driver) SetMachineOSFromFlags(flags drivers.DriverOptions) error {
	d.OS = strings.ToLower(flags.String("vmwarevsphere-os"))
	if d.OS == "" {
		d.OS = defaultMachineOS
		return nil
	}
	err := d.CheckMachineOS(d.OS)
	if !err {
		return fmt.Errorf("[SetMachineOSFromFlags] Machine [%s] has an unsupported MachineOS [%s]\n", d.MachineName, d.OS)
	}
	return nil
}

func (d *Driver) CheckMachineOS(os string) bool {
	_, ok := supportedMachineOS[os]
	return ok
}

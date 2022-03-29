package rackspace

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

	"github.com/bhojpur/host/pkg/core/drivers"
	mflag "github.com/bhojpur/host/pkg/core/flag"
	"github.com/bhojpur/host/pkg/core/log"
	"github.com/bhojpur/host/pkg/drivers/openstack"
)

// Driver is a machine driver for Rackspace. It's a specialization of the generic OpenStack one.
type Driver struct {
	*openstack.Driver

	APIKey string
}

const (
	defaultRegionName     = "IAD"
	defaultEndpointType   = "publicURL"
	defaultFlavorID       = "general1-1"
	defaultSSHUser        = "root"
	defaultSSHPort        = 22
	defaultBhojpurInstall = "true"
	defaultActiveTimeout  = 300
)

// GetCreateFlags registers the "machine create" flags recognized by this driver, including
// their help text and defaults.
func (d *Driver) GetCreateFlags() []mflag.Flag {
	return []mflag.Flag{
		mflag.StringFlag{
			EnvVar: "OS_USERNAME",
			Name:   "rackspace-username",
			Usage:  "Rackspace account username",
			Value:  "",
		},
		mflag.StringFlag{
			EnvVar: "OS_API_KEY",
			Name:   "rackspace-api-key",
			Usage:  "Rackspace API key",
			Value:  "",
		},
		mflag.StringFlag{
			EnvVar: "OS_REGION_NAME",
			Name:   "rackspace-region",
			Usage:  "Rackspace region name",
			Value:  defaultRegionName,
		},
		mflag.StringFlag{
			EnvVar: "OS_ENDPOINT_TYPE",
			Name:   "rackspace-endpoint-type",
			Usage:  "Rackspace endpoint type (adminURL, internalURL or the default publicURL)",
			Value:  defaultEndpointType,
		},
		mflag.StringFlag{
			EnvVar: "OS_IMAGE_ID",
			Name:   "rackspace-image-id",
			Usage:  "Rackspace image ID. Default: Ubuntu 16.04 LTS (Xenial Xerus) (PVHVM)",
		},
		mflag.StringFlag{
			EnvVar: "OS_FLAVOR_ID",
			Name:   "rackspace-flavor-id",
			Usage:  "Rackspace flavor ID. Default: General Purpose 1GB",
			Value:  defaultFlavorID,
		},
		mflag.StringFlag{
			EnvVar: "OS_SSH_USER",
			Name:   "rackspace-ssh-user",
			Usage:  "SSH user for the newly booted machine. Set to root by default",
			Value:  defaultSSHUser,
		},
		mflag.IntFlag{
			EnvVar: "OS_SSH_PORT",
			Name:   "rackspace-ssh-port",
			Usage:  "SSH port for the newly booted machine. Set to 22 by default",
			Value:  defaultSSHPort,
		},
		mflag.StringFlag{
			Name:  "rackspace-docker-install",
			Usage: "Set if Bhojpur Host have to be installed on the machine",
			Value: defaultBhojpurInstall,
		},
		mflag.IntFlag{
			EnvVar: "OS_ACTIVE_TIMEOUT",
			Name:   "rackspace-active-timeout",
			Usage:  "Rackspace active timeout",
			Value:  defaultActiveTimeout,
		},
	}
}

// NewDriver instantiates a Rackspace driver.
func NewDriver(machineName, storePath string) drivers.Driver {
	log.Debug("Instantiating Rackspace driver.", map[string]string{"machineName": machineName})

	inner := openstack.NewDerivedDriver(machineName, storePath)
	driver := &Driver{
		Driver: inner,
	}
	inner.SetClient(&Client{
		driver: driver,
	})

	return driver
}

// DriverName returns the name of the driver
func (d *Driver) DriverName() string {
	return "rackspace"
}

func missingEnvOrOption(setting, envVar, opt string) error {
	return fmt.Errorf(
		"%s must be specified either using the environment variable %s or the CLI option %s",
		setting,
		envVar,
		opt,
	)
}

// SetConfigFromFlags assigns and verifies the command-line arguments presented to the driver.
func (d *Driver) SetConfigFromFlags(flags drivers.DriverOptions) error {
	d.ActiveTimeout = flags.Int("rackspace-active-timeout")
	d.Username = flags.String("rackspace-username")
	d.APIKey = flags.String("rackspace-api-key")
	d.Region = flags.String("rackspace-region")
	d.EndpointType = flags.String("rackspace-endpoint-type")
	d.ImageId = flags.String("rackspace-image-id")
	d.FlavorId = flags.String("rackspace-flavor-id")
	d.SSHUser = flags.String("rackspace-ssh-user")
	d.SSHPort = flags.Int("rackspace-ssh-port")
	d.SetSwarmConfigFromFlags(flags)

	if d.Region == "" {
		return missingEnvOrOption("Region", "OS_REGION_NAME", "--rackspace-region")
	}
	if d.Username == "" {
		return missingEnvOrOption("Username", "OS_USERNAME", "--rackspace-username")
	}
	if d.APIKey == "" {
		return missingEnvOrOption("API key", "OS_API_KEY", "--rackspace-api-key")
	}

	if d.ImageId == "" {
		// Default to the Ubuntu 16.04 image.
		// This is done here, rather than in the option registration, to keep the default value
		// from making "machine create --help" ugly.
		d.ImageId = "821ba5f4-712d-4ec8-9c65-a3fa4bc500f9"
	}

	if d.EndpointType != "publicURL" && d.EndpointType != "adminURL" && d.EndpointType != "internalURL" {
		return fmt.Errorf("invalid endpoint type %q (must be publicURL, adminURL or internalURL)", d.EndpointType)
	}

	return nil
}

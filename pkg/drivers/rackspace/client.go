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

	"github.com/bhojpur/host/pkg/machine/log"
	// "github.com/bhojpur/host/pkg/machine/version"
	"github.com/bhojpur/host/pkg/drivers/openstack"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/rackspace"
)

func unsupportedOpErr(operation string) error {
	return fmt.Errorf("Rackspace does not currently support the %s operation", operation)
}

// Client is a Rackspace specialization of the generic OpenStack driver.
type Client struct {
	openstack.GenericClient
	Provider *gophercloud.ProviderClient
	driver   *Driver
}

// Authenticate creates a Rackspace-specific Gophercloud client.
func (c *Client) Authenticate(d *openstack.Driver) error {
	if c.Provider != nil {
		return nil
	}

	log.Debug("Authenticating to Rackspace.", map[string]string{
		"Username": d.Username,
	})

	apiKey := c.driver.APIKey
	opts := gophercloud.AuthOptions{
		Username: d.Username,
		APIKey:   apiKey,
	}

	provider, err := rackspace.NewClient(rackspace.RackspaceUSIdentity)
	if err != nil {
		return err
	}

	//provider.UserAgent.Prepend(fmt.Sprintf("bhojpur-machine/v%d", version.APIVersion))

	err = rackspace.Authenticate(provider, opts)
	if err != nil {
		return err
	}

	c.Provider = provider

	return nil
}

// StartInstance is unfortunately not supported on Rackspace at this time.
func (c *Client) StartInstance(d *openstack.Driver) error {
	return unsupportedOpErr("start")
}

// StopInstance is unfortunately not support on Rackspace at this time.
func (c *Client) StopInstance(d *openstack.Driver) error {
	return unsupportedOpErr("stop")
}

// GetInstanceIPAddresses can be short-circuited with the server's AccessIPv4Addr on Rackspace.
func (c *Client) GetInstanceIPAddresses(d *openstack.Driver) ([]openstack.IPAddress, error) {
	server, err := c.GetServerDetail(d)
	if err != nil {
		return nil, err
	}
	return []openstack.IPAddress{
		{
			Network:     "public",
			Address:     server.AccessIPv4,
			AddressType: openstack.Fixed,
		},
	}, nil
}

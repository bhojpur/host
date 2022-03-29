package none

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
	neturl "net/url"

	"github.com/bhojpur/host/pkg/core/drivers"
	mflag "github.com/bhojpur/host/pkg/core/flag"
	"github.com/bhojpur/host/pkg/core/state"
)

const driverName = "none"

// Driver is the driver used when no driver is selected. It is used to
// connect to existing Bhojpur Hosts by specifying the URL of the host as
// an option.
type Driver struct {
	*drivers.BaseDriver
	URL string
}

func NewDriver(hostName, storePath string) *Driver {
	return &Driver{
		BaseDriver: &drivers.BaseDriver{
			MachineName: hostName,
			StorePath:   storePath,
		},
	}
}

func (d *Driver) GetCreateFlags() []mflag.Flag {
	return []mflag.Flag{
		mflag.StringFlag{
			Name:  "url",
			Usage: "URL of host when no driver is selected",
			Value: "",
		},
	}
}

func (d *Driver) Create() error {
	return nil
}

// DriverName returns the name of the driver
func (d *Driver) DriverName() string {
	return driverName
}

func (d *Driver) GetIP() (string, error) {
	return d.IPAddress, nil
}

func (d *Driver) GetSSHHostname() (string, error) {
	return "", nil
}

func (d *Driver) GetSSHKeyPath() string {
	return ""
}

func (d *Driver) GetSSHPort() (int, error) {
	return 0, nil
}

func (d *Driver) GetSSHUsername() string {
	return ""
}

func (d *Driver) GetURL() (string, error) {
	return d.URL, nil
}

func (d *Driver) GetState() (state.State, error) {
	return state.Running, nil
}

func (d *Driver) Kill() error {
	return fmt.Errorf("hosts without a driver cannot be killed")
}

func (d *Driver) Remove() error {
	return nil
}

func (d *Driver) Restart() error {
	return fmt.Errorf("hosts without a driver cannot be restarted")
}

func (d *Driver) SetConfigFromFlags(flags drivers.DriverOptions) error {
	url := flags.String("url")

	if url == "" {
		return fmt.Errorf("--url option is required when no driver is selected")
	}

	d.URL = url
	u, err := neturl.Parse(url)
	if err != nil {
		return err
	}

	d.IPAddress = u.Host
	return nil
}

func (d *Driver) Start() error {
	return fmt.Errorf("hosts without a driver cannot be started")
}

func (d *Driver) Stop() error {
	return fmt.Errorf("hosts without a driver cannot be stopped")
}

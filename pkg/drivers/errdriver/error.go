package errdriver

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

	"github.com/bhojpur/host/pkg/machine/drivers"
	mflag "github.com/bhojpur/host/pkg/machine/flag"
	"github.com/bhojpur/host/pkg/machine/state"
)

type Driver struct {
	Name string
}

type NotLoadable struct {
	Name string
}

func (e NotLoadable) Error() string {
	return fmt.Sprintf("Driver %q not found. Do you have the plugin binary accessible in your PATH?", e.Name)
}

func NewDriver(Name string) drivers.Driver {
	return &Driver{
		Name: Name,
	}
}

// DriverName returns the name of the driver
func (d *Driver) DriverName() string {
	return "not-found"
}

func (d *Driver) PreCreateCheck() error {
	return NotLoadable{d.Name}
}

func (d *Driver) GetCreateFlags() []mflag.Flag {
	return nil
}

func (d *Driver) SetConfigFromFlags(flags drivers.DriverOptions) error {
	return NotLoadable{d.Name}
}

func (d *Driver) GetURL() (string, error) {
	return "", NotLoadable{d.Name}
}

func (d *Driver) GetMachineName() string {
	return d.Name
}

func (d *Driver) GetIP() (string, error) {
	return "1.2.3.4", NotLoadable{d.Name}
}

func (d *Driver) GetSSHHostname() (string, error) {
	return "", NotLoadable{d.Name}
}

func (d *Driver) GetSSHKeyPath() string {
	return ""
}

func (d *Driver) GetSSHPort() (int, error) {
	return 0, NotLoadable{d.Name}
}

func (d *Driver) GetSSHUsername() string {
	return ""
}

func (d *Driver) GetState() (state.State, error) {
	return state.Error, NotLoadable{d.Name}
}

func (d *Driver) Create() error {
	return NotLoadable{d.Name}
}

func (d *Driver) Remove() error {
	return NotLoadable{d.Name}
}

func (d *Driver) Start() error {
	return NotLoadable{d.Name}
}

func (d *Driver) Stop() error {
	return NotLoadable{d.Name}
}

func (d *Driver) Restart() error {
	return NotLoadable{d.Name}
}

func (d *Driver) Kill() error {
	return NotLoadable{d.Name}
}

func (d *Driver) Upgrade() error {
	return NotLoadable{d.Name}
}

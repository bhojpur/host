package drivers

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

	mflag "github.com/bhojpur/host/pkg/core/flag"
	"github.com/bhojpur/host/pkg/core/state"
)

type DriverNotSupported struct {
	*BaseDriver
	Name string
}

type NotSupported struct {
	DriverName string
}

func (e NotSupported) Error() string {
	return fmt.Sprintf("Driver %q not supported on this platform.", e.DriverName)
}

// NewDriverNotSupported creates a placeholder Driver that replaces
// a driver that is not supported on a given platform. eg fusion on linux.
func NewDriverNotSupported(driverName, hostName, storePath string) Driver {
	return &DriverNotSupported{
		BaseDriver: &BaseDriver{
			MachineName: hostName,
			StorePath:   storePath,
		},
		Name: driverName,
	}
}

func (d *DriverNotSupported) DriverName() string {
	return d.Name
}

func (d *DriverNotSupported) PreCreateCheck() error {
	return NotSupported{d.DriverName()}
}

func (d *DriverNotSupported) GetCreateFlags() []mflag.Flag {
	return nil
}

func (d *DriverNotSupported) SetConfigFromFlags(flags DriverOptions) error {
	return NotSupported{d.DriverName()}
}

func (d *DriverNotSupported) GetURL() (string, error) {
	return "", NotSupported{d.DriverName()}
}

func (d *DriverNotSupported) GetSSHHostname() (string, error) {
	return "", NotSupported{d.DriverName()}
}

func (d *DriverNotSupported) GetState() (state.State, error) {
	return state.Error, NotSupported{d.DriverName()}
}

func (d *DriverNotSupported) Create() error {
	return NotSupported{d.DriverName()}
}

func (d *DriverNotSupported) Remove() error {
	return NotSupported{d.DriverName()}
}

func (d *DriverNotSupported) Start() error {
	return NotSupported{d.DriverName()}
}

func (d *DriverNotSupported) Stop() error {
	return NotSupported{d.DriverName()}
}

func (d *DriverNotSupported) Restart() error {
	return NotSupported{d.DriverName()}
}

func (d *DriverNotSupported) Kill() error {
	return NotSupported{d.DriverName()}
}

func (d *DriverNotSupported) Upgrade() error {
	return NotSupported{d.DriverName()}
}

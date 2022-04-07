package apitest

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
	core "github.com/bhojpur/host/pkg/machine"
	"github.com/bhojpur/host/pkg/machine/drivers"
	merrors "github.com/bhojpur/host/pkg/machine/errors"
	"github.com/bhojpur/host/pkg/machine/host"
	"github.com/bhojpur/host/pkg/machine/state"
)

type FakeAPI struct {
	Hosts []*host.Host
}

func (api *FakeAPI) NewPluginDriver(string, []byte) (drivers.Driver, error) {
	return nil, nil
}

func (api *FakeAPI) Close() error {
	return nil
}

func (api *FakeAPI) NewHost(driverName string, rawDriver []byte) (*host.Host, error) {
	return nil, nil
}

func (api *FakeAPI) Create(h *host.Host) error {
	return nil
}

func (api *FakeAPI) Exists(name string) (bool, error) {
	for _, host := range api.Hosts {
		if name == host.Name {
			return true, nil
		}
	}

	return false, nil
}

func (api *FakeAPI) List() ([]string, error) {
	return []string{}, nil
}

func (api *FakeAPI) Load(name string) (*host.Host, error) {
	for _, host := range api.Hosts {
		if name == host.Name {
			return host, nil
		}
	}

	return nil, merrors.ErrHostDoesNotExist{
		Name: name,
	}
}

func (api *FakeAPI) Remove(name string) error {
	newHosts := []*host.Host{}

	for _, host := range api.Hosts {
		if name != host.Name {
			newHosts = append(newHosts, host)
		}
	}

	api.Hosts = newHosts

	return nil
}

func (api *FakeAPI) Save(host *host.Host) error {
	return nil
}

func (api FakeAPI) GetMachinesDir() string {
	return ""
}

func State(api core.API, name string) state.State {
	host, _ := api.Load(name)
	machineState, _ := host.Driver.GetState()
	return machineState
}

func Exists(api core.API, name string) bool {
	exists, _ := api.Exists(name)
	return exists
}

package commands

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
	"testing"

	"errors"

	ctest "github.com/bhojpur/host/cmd/machine/commands/test"
	"github.com/bhojpur/host/pkg/drivers/fakedriver"
	atest "github.com/bhojpur/host/pkg/machine/apitest"
	"github.com/bhojpur/host/pkg/machine/host"
	"github.com/stretchr/testify/assert"
)

func TestCmdRmMissingMachineName(t *testing.T) {
	commandLine := &ctest.FakeCommandLine{}
	api := &atest.FakeAPI{}

	err := cmdRm(commandLine, api)

	assert.Equal(t, ErrNoMachineSpecified, err)
	assert.True(t, commandLine.HelpShown)
}

func TestCmdRm(t *testing.T) {
	commandLine := &ctest.FakeCommandLine{
		CliArgs: []string{"machineToRemove1", "machineToRemove2"},
		LocalFlags: &ctest.FakeFlagger{
			Data: map[string]interface{}{
				"y": true,
			},
		},
	}
	api := &atest.FakeAPI{
		Hosts: []*host.Host{
			{
				Name:   "machineToRemove1",
				Driver: &fakedriver.Driver{},
			},
			{
				Name:   "machineToRemove2",
				Driver: &fakedriver.Driver{},
			},
			{
				Name:   "machine",
				Driver: &fakedriver.Driver{},
			},
		},
	}

	err := cmdRm(commandLine, api)
	assert.NoError(t, err)

	assert.False(t, atest.Exists(api, "machineToRemove1"))
	assert.False(t, atest.Exists(api, "machineToRemove2"))
	assert.True(t, atest.Exists(api, "machine"))
}

func TestCmdRmforcefully(t *testing.T) {
	commandLine := &ctest.FakeCommandLine{
		CliArgs: []string{"machineToRemove1", "machineToRemove2"},
		LocalFlags: &ctest.FakeFlagger{
			Data: map[string]interface{}{
				"force": true,
			},
		},
	}
	api := &atest.FakeAPI{
		Hosts: []*host.Host{
			{
				Name:   "machineToRemove1",
				Driver: &fakedriver.Driver{},
			},
			{
				Name:   "machineToRemove2",
				Driver: &fakedriver.Driver{},
			},
		},
	}

	err := cmdRm(commandLine, api)
	assert.NoError(t, err)

	assert.False(t, atest.Exists(api, "machineToRemove1"))
	assert.False(t, atest.Exists(api, "machineToRemove2"))
}

func TestCmdRmforceDoesAutoConfirm(t *testing.T) {
	commandLine := &ctest.FakeCommandLine{
		CliArgs: []string{"machineToRemove1", "machineToRemove2"},
		LocalFlags: &ctest.FakeFlagger{
			Data: map[string]interface{}{
				"y":     false,
				"force": true,
			},
		},
	}
	api := &atest.FakeAPI{
		Hosts: []*host.Host{
			{
				Name:   "machineToRemove1",
				Driver: &fakedriver.Driver{},
			},
			{
				Name:   "machineToRemove2",
				Driver: &fakedriver.Driver{},
			},
		},
	}

	err := cmdRm(commandLine, api)
	assert.NoError(t, err)

	assert.False(t, atest.Exists(api, "machineToRemove1"))
	assert.False(t, atest.Exists(api, "machineToRemove2"))
}

func TestCmdRmforceConfirmUnset(t *testing.T) {
	commandLine := &ctest.FakeCommandLine{
		CliArgs: []string{"machineToRemove1"},
		LocalFlags: &ctest.FakeFlagger{
			Data: map[string]interface{}{
				"y":     false,
				"force": false,
			},
		},
	}
	api := &atest.FakeAPI{
		Hosts: []*host.Host{
			{
				Name:   "machineToRemove1",
				Driver: &fakedriver.Driver{},
			},
		},
	}

	err := cmdRm(commandLine, api)
	assert.NoError(t, err)

	assert.True(t, atest.Exists(api, "machineToRemove1"))
}

type DriverWithRemoveWhichFail struct {
	fakedriver.Driver
}

func (d *DriverWithRemoveWhichFail) Remove() error {
	return errors.New("unknown error")
}

func TestDontStopWhenADriverRemovalFails(t *testing.T) {
	commandLine := &ctest.FakeCommandLine{
		CliArgs: []string{"machineToRemove1", "machineToRemove2", "machineToRemove3"},
		LocalFlags: &ctest.FakeFlagger{
			Data: map[string]interface{}{
				"y": true,
			},
		},
	}
	api := &atest.FakeAPI{
		Hosts: []*host.Host{
			{
				Name:   "machineToRemove1",
				Driver: &fakedriver.Driver{},
			},
			{
				Name:   "machineToRemove2",
				Driver: &DriverWithRemoveWhichFail{},
			},
			{
				Name:   "machineToRemove3",
				Driver: &fakedriver.Driver{},
			},
		},
	}

	err := cmdRm(commandLine, api)
	assert.EqualError(t, err, "Error removing host \"machineToRemove2\": unknown error")

	assert.False(t, atest.Exists(api, "machineToRemove1"))
	assert.True(t, atest.Exists(api, "machineToRemove2"))
	assert.False(t, atest.Exists(api, "machineToRemove3"))
}

func TestForceRemoveEvenWhenItFails(t *testing.T) {
	commandLine := &ctest.FakeCommandLine{
		CliArgs: []string{"machineToRemove1"},
		LocalFlags: &ctest.FakeFlagger{
			Data: map[string]interface{}{
				"y":     true,
				"force": true,
			},
		},
	}
	api := &atest.FakeAPI{
		Hosts: []*host.Host{
			{
				Name:   "machineToRemove1",
				Driver: &DriverWithRemoveWhichFail{},
			},
		},
	}

	err := cmdRm(commandLine, api)
	assert.NoError(t, err)

	assert.False(t, atest.Exists(api, "machineToRemove1"))
}

func TestDontRemoveMachineIsRemovalFailsAndNotForced(t *testing.T) {
	commandLine := &ctest.FakeCommandLine{
		CliArgs: []string{"machineToRemove1"},
		LocalFlags: &ctest.FakeFlagger{
			Data: map[string]interface{}{
				"y":     true,
				"force": false,
			},
		},
	}
	api := &atest.FakeAPI{
		Hosts: []*host.Host{
			{
				Name:   "machineToRemove1",
				Driver: &DriverWithRemoveWhichFail{},
			},
		},
	}

	err := cmdRm(commandLine, api)
	assert.EqualError(t, err, "Error removing host \"machineToRemove1\": unknown error")

	assert.True(t, atest.Exists(api, "machineToRemove1"))
}

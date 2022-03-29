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

	ctest "github.com/bhojpur/host/cmd/machine/commands/test"
	atest "github.com/bhojpur/host/pkg/core/apitest"
	"github.com/bhojpur/host/pkg/core/host"
	"github.com/bhojpur/host/pkg/core/state"
	"github.com/bhojpur/host/pkg/drivers/fakedriver"
	"github.com/stretchr/testify/assert"
)

func TestCmdURLMissingMachineName(t *testing.T) {
	commandLine := &ctest.FakeCommandLine{}
	api := &atest.FakeAPI{}

	err := cmdURL(commandLine, api)

	assert.Equal(t, ErrNoDefault, err)
}

func TestCmdURLTooManyNames(t *testing.T) {
	commandLine := &ctest.FakeCommandLine{
		CliArgs: []string{"machineToRemove1", "machineToRemove2"},
	}
	api := &atest.FakeAPI{}

	err := cmdURL(commandLine, api)

	assert.EqualError(t, err, "Error: Expected one machine name as an argument")
}

func TestCmdURL(t *testing.T) {
	commandLine := &ctest.FakeCommandLine{
		CliArgs: []string{"machine"},
	}
	api := &atest.FakeAPI{
		Hosts: []*host.Host{
			{
				Name: "machine",
				Driver: &fakedriver.Driver{
					MockState: state.Running,
					MockIP:    "120.0.0.1",
				},
			},
		},
	}

	stdoutGetter := ctest.NewStdoutGetter()
	defer stdoutGetter.Stop()

	err := cmdURL(commandLine, api)

	assert.NoError(t, err)
	assert.Equal(t, "tcp://120.0.0.1:2376\n", stdoutGetter.Output())
}

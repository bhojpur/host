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
	"errors"
	"testing"

	"bytes"

	ctest "github.com/bhojpur/host/cmd/machine/commands/test"
	atest "github.com/bhojpur/host/pkg/machine/apitest"
	cengine "github.com/bhojpur/host/pkg/machine/client"
	"github.com/bhojpur/host/pkg/machine/host"
	"github.com/stretchr/testify/assert"
)

func TestCmdVersion(t *testing.T) {
	commandLine := &ctest.FakeCommandLine{}
	api := &atest.FakeAPI{}

	err := cmdVersion(commandLine, api)

	assert.True(t, commandLine.VersionShown)
	assert.NoError(t, err)
}

func TestCmdVersionTooManyNames(t *testing.T) {
	commandLine := &ctest.FakeCommandLine{
		CliArgs: []string{"machine1", "machine2"},
	}
	api := &atest.FakeAPI{}

	err := cmdVersion(commandLine, api)

	assert.EqualError(t, err, "Error: Expected one machine name as an argument")
}

func TestCmdVersionNotFound(t *testing.T) {
	commandLine := &ctest.FakeCommandLine{
		CliArgs: []string{"unknown"},
	}
	api := &atest.FakeAPI{}

	err := cmdVersion(commandLine, api)

	assert.EqualError(t, err, `Bhojpur Host machine "unknown" does not exist. Use "hostutl ls" to list machines. Use "hostutl create" to add a new one.`)
}

func TestCmdVersionOnHost(t *testing.T) {
	defer func(versioner cengine.BhojpurVersioner) { cengine.CurrentBhojpurVersioner = versioner }(cengine.CurrentBhojpurVersioner)
	cengine.CurrentBhojpurVersioner = &cengine.FakeBhojpurVersioner{Version: "1.9.1"}

	commandLine := &ctest.FakeCommandLine{
		CliArgs: []string{"machine"},
	}
	api := &atest.FakeAPI{
		Hosts: []*host.Host{
			{
				Name: "machine",
			},
		},
	}

	out := &bytes.Buffer{}
	err := printVersion(commandLine, api, out)

	assert.NoError(t, err)
	assert.Equal(t, "1.9.1\n", out.String())
}

func TestCmdVersionFailure(t *testing.T) {
	defer func(versioner cengine.BhojpurVersioner) { cengine.CurrentBhojpurVersioner = versioner }(cengine.CurrentBhojpurVersioner)
	cengine.CurrentBhojpurVersioner = &cengine.FakeBhojpurVersioner{Err: errors.New("connection failure")}

	commandLine := &ctest.FakeCommandLine{
		CliArgs: []string{"machine"},
	}
	api := &atest.FakeAPI{
		Hosts: []*host.Host{
			{
				Name: "machine",
			},
		},
	}

	out := &bytes.Buffer{}
	err := printVersion(commandLine, api, out)

	assert.EqualError(t, err, "connection failure")
}

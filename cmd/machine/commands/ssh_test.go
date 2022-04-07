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
	"github.com/bhojpur/host/pkg/drivers/fakedriver"
	core "github.com/bhojpur/host/pkg/machine"
	atest "github.com/bhojpur/host/pkg/machine/apitest"
	"github.com/bhojpur/host/pkg/machine/drivers"
	"github.com/bhojpur/host/pkg/machine/host"
	"github.com/bhojpur/host/pkg/machine/ssh"
	"github.com/bhojpur/host/pkg/machine/ssh/sshtest"
	"github.com/bhojpur/host/pkg/machine/state"
	"github.com/stretchr/testify/assert"
)

type FakeSSHClientCreator struct {
	client ssh.Client
}

func (fsc *FakeSSHClientCreator) CreateSSHClient(d drivers.Driver) (ssh.Client, error) {
	if fsc.client == nil {
		fsc.client = &sshtest.FakeClient{}
	}
	return fsc.client, nil
}

func TestCmdSSH(t *testing.T) {
	testCases := []struct {
		commandLine   CommandLine
		api           core.API
		expectedErr   error
		helpShown     bool
		clientCreator host.SSHClientCreator
		expectedShell []string
	}{
		{
			commandLine: &ctest.FakeCommandLine{
				CliArgs: []string{"-h"},
			},
			api:         &atest.FakeAPI{},
			expectedErr: nil,
			helpShown:   true,
		},
		{
			commandLine: &ctest.FakeCommandLine{
				CliArgs: []string{"--help"},
			},
			api:         &atest.FakeAPI{},
			expectedErr: nil,
			helpShown:   true,
		},
		{
			commandLine: &ctest.FakeCommandLine{
				CliArgs: []string{},
			},
			api:         &atest.FakeAPI{},
			expectedErr: ErrNoDefault,
		},
		{
			commandLine: &ctest.FakeCommandLine{
				CliArgs: []string{"default", "df", "-h"},
			},
			api: &atest.FakeAPI{
				Hosts: []*host.Host{
					{
						Name: "default",
						Driver: &fakedriver.Driver{
							MockState: state.Running,
						},
					},
				},
			},
			expectedErr:   nil,
			clientCreator: &FakeSSHClientCreator{},
			expectedShell: []string{"df", "-h"},
		},
		{
			commandLine: &ctest.FakeCommandLine{
				CliArgs: []string{"default"},
			},
			api: &atest.FakeAPI{
				Hosts: []*host.Host{
					{
						Name: "default",
						Driver: &fakedriver.Driver{
							MockState: state.Stopped,
						},
					},
				},
			},
			expectedErr: errStateInvalidForSSH{"default"},
		},
	}

	for _, tc := range testCases {
		host.SetSSHClientCreator(tc.clientCreator)

		err := cmdSSH(tc.commandLine, tc.api)
		assert.Equal(t, err, tc.expectedErr)

		if fcl, ok := tc.commandLine.(*ctest.FakeCommandLine); ok {
			assert.Equal(t, tc.helpShown, fcl.HelpShown)
		}

		if fcc, ok := tc.clientCreator.(*FakeSSHClientCreator); ok {
			assert.Equal(t, tc.expectedShell, fcc.client.(*sshtest.FakeClient).ActivatedShell)
		}
	}
}

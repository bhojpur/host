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
	"github.com/bhojpur/host/pkg/core"
	atest "github.com/bhojpur/host/pkg/core/apitest"
	"github.com/bhojpur/host/pkg/core/host"
	"github.com/bhojpur/host/pkg/core/state"
	"github.com/bhojpur/host/pkg/drivers/fakedriver"
	"github.com/stretchr/testify/assert"
)

func TestCmdIPMissingMachineName(t *testing.T) {
	commandLine := &ctest.FakeCommandLine{}
	api := &atest.FakeAPI{}

	err := cmdURL(commandLine, api)

	assert.Equal(t, err, ErrNoDefault)
}

func TestCmdIP(t *testing.T) {
	testCases := []struct {
		commandLine CommandLine
		api         core.API
		expectedErr error
		expectedOut string
	}{
		{
			commandLine: &ctest.FakeCommandLine{
				CliArgs: []string{"machine"},
			},
			api: &atest.FakeAPI{
				Hosts: []*host.Host{
					{
						Name: "machine",
						Driver: &fakedriver.Driver{
							MockState: state.Running,
							MockIP:    "1.2.3.4",
						},
					},
				},
			},
			expectedErr: nil,
			expectedOut: "1.2.3.4\n",
		},
		{
			commandLine: &ctest.FakeCommandLine{
				CliArgs: []string{},
			},
			api: &atest.FakeAPI{
				Hosts: []*host.Host{
					{
						Name: "default",
						Driver: &fakedriver.Driver{
							MockState: state.Running,
							MockIP:    "1.2.3.4",
						},
					},
				},
			},
			expectedErr: nil,
			expectedOut: "1.2.3.4\n",
		},
	}

	for _, tc := range testCases {
		stdoutGetter := ctest.NewStdoutGetter()

		err := cmdIP(tc.commandLine, tc.api)

		assert.Equal(t, tc.expectedErr, err)
		assert.Equal(t, tc.expectedOut, stdoutGetter.Output())

		stdoutGetter.Stop()
	}
}

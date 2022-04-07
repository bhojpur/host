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
	"github.com/bhojpur/host/pkg/machine/host"
	"github.com/bhojpur/host/pkg/machine/state"
	"github.com/stretchr/testify/assert"
)

func TestCmdStop(t *testing.T) {
	testCases := []struct {
		commandLine    CommandLine
		api            core.API
		expectedErr    error
		expectedStates map[string]state.State
	}{
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
						},
					},
				},
			},
			expectedErr: nil,
			expectedStates: map[string]state.State{
				"default": state.Stopped,
			},
		},
		{
			commandLine: &ctest.FakeCommandLine{
				CliArgs: []string{},
			},
			api: &atest.FakeAPI{
				Hosts: []*host.Host{
					{
						Name: "foobar",
						Driver: &fakedriver.Driver{
							MockState: state.Running,
						},
					},
				},
			},
			expectedErr: ErrNoDefault,
			expectedStates: map[string]state.State{
				"foobar": state.Running,
			},
		},
		{
			commandLine: &ctest.FakeCommandLine{
				CliArgs: []string{"machineToStop1", "machineToStop2"},
			},
			api: &atest.FakeAPI{
				Hosts: []*host.Host{
					{
						Name: "machineToStop1",
						Driver: &fakedriver.Driver{
							MockState: state.Running,
						},
					},
					{
						Name: "machineToStop2",
						Driver: &fakedriver.Driver{
							MockState: state.Running,
						},
					},
					{
						Name: "machine",
						Driver: &fakedriver.Driver{
							MockState: state.Running,
						},
					},
				},
			},
			expectedErr: nil,
			expectedStates: map[string]state.State{
				"machineToStop1": state.Stopped,
				"machineToStop2": state.Stopped,
				"machine":        state.Running,
			},
		},
	}

	for _, tc := range testCases {
		err := cmdStop(tc.commandLine, tc.api)
		assert.Equal(t, tc.expectedErr, err)

		for hostName, expectedState := range tc.expectedStates {
			assert.Equal(t, expectedState, atest.State(tc.api, hostName))
		}
	}
}

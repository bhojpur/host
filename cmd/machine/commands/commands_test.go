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
	"flag"
	"testing"

	ctest "github.com/bhojpur/host/cmd/machine/commands/test"
	"github.com/bhojpur/host/pkg/drivers/fakedriver"
	core "github.com/bhojpur/host/pkg/machine"
	"github.com/bhojpur/host/pkg/machine/crashreport"
	merrors "github.com/bhojpur/host/pkg/machine/errors"
	"github.com/bhojpur/host/pkg/machine/host"
	"github.com/bhojpur/host/pkg/machine/hosttest"
	"github.com/bhojpur/host/pkg/machine/provision"
	"github.com/bhojpur/host/pkg/machine/state"
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli"
)

func TestRunActionForeachMachine(t *testing.T) {
	defer provision.SetDetector(&provision.StandardDetector{})
	provision.SetDetector(&provision.FakeDetector{
		Provisioner: provision.NewNetstatProvisioner(),
	})

	// Assume a bunch of machines in randomly started or
	// stopped states.
	machines := []*host.Host{
		{
			Name:       "foo",
			DriverName: "fakedriver",
			Driver: &fakedriver.Driver{
				MockState: state.Running,
			},
		},
		{
			Name:       "bar",
			DriverName: "fakedriver",
			Driver: &fakedriver.Driver{
				MockState: state.Stopped,
			},
		},
		{
			Name: "baz",
			// Ssh, don't tell anyone but this
			// driver only _thinks_ it's named
			// virtualbox...  (to test serial actions)
			// It's actually FakeDriver!
			DriverName: "virtualbox",
			Driver: &fakedriver.Driver{
				MockState: state.Stopped,
			},
		},
		{
			Name:       "spam",
			DriverName: "virtualbox",
			Driver: &fakedriver.Driver{
				MockState: state.Running,
			},
		},
		{
			Name:       "eggs",
			DriverName: "fakedriver",
			Driver: &fakedriver.Driver{
				MockState: state.Stopped,
			},
		},
		{
			Name:       "ham",
			DriverName: "fakedriver",
			Driver: &fakedriver.Driver{
				MockState: state.Running,
			},
		},
	}

	runActionForeachMachine("start", machines)

	for _, machine := range machines {
		machineState, _ := machine.Driver.GetState()

		assert.Equal(t, state.Running, machineState)
	}

	runActionForeachMachine("stop", machines)

	for _, machine := range machines {
		machineState, _ := machine.Driver.GetState()

		assert.Equal(t, state.Stopped, machineState)
	}
}

func TestPrintIPEmptyGivenLocalEngine(t *testing.T) {
	stdoutGetter := ctest.NewStdoutGetter()
	defer stdoutGetter.Stop()

	host, _ := hosttest.GetDefaultTestHost()
	err := printIP(host)()

	assert.NoError(t, err)
	assert.Equal(t, "\n", stdoutGetter.Output())
}

func TestPrintIPPrintsGivenRemoteEngine(t *testing.T) {
	stdoutGetter := ctest.NewStdoutGetter()
	defer stdoutGetter.Stop()

	host, _ := hosttest.GetDefaultTestHost()
	host.Driver = &fakedriver.Driver{
		MockState: state.Running,
		MockIP:    "1.2.3.4",
	}
	err := printIP(host)()

	assert.NoError(t, err)
	assert.Equal(t, "1.2.3.4\n", stdoutGetter.Output())
}

func TestConsolidateError(t *testing.T) {
	cases := []struct {
		inputErrs   []error
		expectedErr error
	}{
		{
			inputErrs: []error{
				errors.New("Couldn't remove host 'bar'"),
			},
			expectedErr: errors.New("Couldn't remove host 'bar'"),
		},
		{
			inputErrs: []error{
				errors.New("Couldn't remove host 'bar'"),
				errors.New("Couldn't remove host 'foo'"),
			},
			expectedErr: errors.New("Couldn't remove host 'bar'\nCouldn't remove host 'foo'"),
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.expectedErr, consolidateErrs(c.inputErrs))
	}
}

type MockCrashReporter struct {
	sent bool
}

func (m *MockCrashReporter) Send(err crashreport.CrashError) error {
	m.sent = true
	return nil
}

func TestSendCrashReport(t *testing.T) {
	defer func(fnOsExit func(code int)) { osExit = fnOsExit }(osExit)
	osExit = func(code int) {}

	defer func(factory func(baseDir string, apiKey string) crashreport.CrashReporter) {
		crashreport.NewCrashReporter = factory
	}(crashreport.NewCrashReporter)

	tests := []struct {
		description string
		err         error
		sent        bool
	}{
		{
			description: "Should send crash error",
			err: crashreport.CrashError{
				Cause:      errors.New("BUG"),
				Command:    "command",
				Context:    "context",
				DriverName: "virtualbox",
			},
			sent: true,
		},
		{
			description: "Should not send standard error",
			err:         errors.New("BUG"),
			sent:        false,
		},
	}

	for _, test := range tests {
		mockCrashReporter := &MockCrashReporter{}
		crashreport.NewCrashReporter = func(baseDir string, apiKey string) crashreport.CrashReporter {
			return mockCrashReporter
		}

		command := func(commandLine CommandLine, api core.API) error {
			return test.err
		}

		context := cli.NewContext(cli.NewApp(), &flag.FlagSet{}, nil)
		runCommand(command)(context)

		assert.Equal(t, test.sent, mockCrashReporter.sent, test.description)
	}
}

func TestReturnExitCode1onError(t *testing.T) {
	command := func(commandLine CommandLine, api core.API) error {
		return errors.New("foo is not bar")
	}

	exitCode := checkErrorCodeForCommand(command)

	assert.Equal(t, 1, exitCode)
}

func TestReturnExitCode3onErrorDuringPreCreate(t *testing.T) {
	command := func(commandLine CommandLine, api core.API) error {
		return crashreport.CrashError{
			Cause: merrors.ErrDuringPreCreate{
				Cause: errors.New("foo is not bar"),
			},
		}
	}

	exitCode := checkErrorCodeForCommand(command)

	assert.Equal(t, 3, exitCode)
}

func checkErrorCodeForCommand(command func(commandLine CommandLine, api core.API) error) int {
	var setExitCode int

	originalOSExit := osExit

	defer func() {
		osExit = originalOSExit
	}()

	osExit = func(code int) {
		setExitCode = code
	}

	context := cli.NewContext(cli.NewApp(), &flag.FlagSet{}, nil)
	runCommand(command)(context)

	return setExitCode
}

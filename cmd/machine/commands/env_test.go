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
	"os"
	"path/filepath"
	"testing"

	mdirs "github.com/bhojpur/host/cmd/machine/commands/dirs"
	ctest "github.com/bhojpur/host/cmd/machine/commands/test"
	"github.com/bhojpur/host/pkg/drivers/fakedriver"
	core "github.com/bhojpur/host/pkg/machine"
	atest "github.com/bhojpur/host/pkg/machine/apitest"
	"github.com/bhojpur/host/pkg/machine/auth"
	"github.com/bhojpur/host/pkg/machine/check"
	"github.com/bhojpur/host/pkg/machine/host"
	"github.com/bhojpur/host/pkg/machine/state"
	"github.com/stretchr/testify/assert"
)

type FakeConnChecker struct {
	BhojpurHost string
	AuthOptions *auth.Options
	Err         error
}

func (fcc *FakeConnChecker) Check(_ *host.Host, _ bool) (string, *auth.Options, error) {
	return fcc.BhojpurHost, fcc.AuthOptions, fcc.Err
}

type SimpleUsageHintGenerator struct {
	Hint string
}

func (suhg *SimpleUsageHintGenerator) GenerateUsageHint(_ string, _ []string) string {
	return suhg.Hint
}

func TestHints(t *testing.T) {
	var tests = []struct {
		userShell     string
		commandLine   []string
		expectedHints string
	}{
		{"", []string{"hostutl", "env", "default"}, "# Run this command to configure your shell: \n# eval $(hostutl env default)\n"},
		{"", []string{"hostutl", "env", "--no-proxy", "default"}, "# Run this command to configure your shell: \n# eval $(hostutl env --no-proxy default)\n"},
		{"", []string{"hostutl", "env", "--swarm", "default"}, "# Run this command to configure your shell: \n# eval $(hostutl env --swarm default)\n"},
		{"", []string{"hostutl", "env", "--no-proxy", "--swarm", "default"}, "# Run this command to configure your shell: \n# eval $(hostutl env --no-proxy --swarm default)\n"},
		{"", []string{"hostutl", "env", "--unset"}, "# Run this command to configure your shell: \n# eval $(hostutl env --unset)\n"},
		{"", []string{`C:\\Program Files\bhojpur-machine.exe`, "env", "default"}, "# Run this command to configure your shell: \n# eval $(\"C:\\\\Program Files\\bhojpur-machine.exe\" env default)\n"},
		{"", []string{`C:\\Me\bhojpur-machine.exe`, "env", "default"}, "# Run this command to configure your shell: \n# eval $(\"C:\\\\Me\\bhojpur-machine.exe\" env default)\n"},

		{"fish", []string{"./hostutl", "env", "--shell=fish", "default"}, "# Run this command to configure your shell: \n# eval (./hostutl env --shell=fish default)\n"},
		{"fish", []string{"./hostutl", "env", "--shell=fish", "--no-proxy", "default"}, "# Run this command to configure your shell: \n# eval (./hostutl env --shell=fish --no-proxy default)\n"},
		{"fish", []string{"./hostutl", "env", "--shell=fish", "--swarm", "default"}, "# Run this command to configure your shell: \n# eval (./hostutl env --shell=fish --swarm default)\n"},
		{"fish", []string{"./hostutl", "env", "--shell=fish", "--no-proxy", "--swarm", "default"}, "# Run this command to configure your shell: \n# eval (./hostutl env --shell=fish --no-proxy --swarm default)\n"},
		{"fish", []string{"./hostutl", "env", "--shell=fish", "--unset"}, "# Run this command to configure your shell: \n# eval (./hostutl env --shell=fish --unset)\n"},

		{"powershell", []string{"./hostutl", "env", "--shell=powershell", "default"}, "# Run this command to configure your shell: \n# & ./hostutl env --shell=powershell default | Invoke-Expression\n"},
		{"powershell", []string{"./hostutl", "env", "--shell=powershell", "--no-proxy", "default"}, "# Run this command to configure your shell: \n# & ./hostutl env --shell=powershell --no-proxy default | Invoke-Expression\n"},
		{"powershell", []string{"./hostutl", "env", "--shell=powershell", "--swarm", "default"}, "# Run this command to configure your shell: \n# & ./hostutl env --shell=powershell --swarm default | Invoke-Expression\n"},
		{"powershell", []string{"./hostutl", "env", "--shell=powershell", "--no-proxy", "--swarm", "default"}, "# Run this command to configure your shell: \n# & ./hostutl env --shell=powershell --no-proxy --swarm default | Invoke-Expression\n"},
		{"powershell", []string{"./hostutl", "env", "--shell=powershell", "--unset"}, "# Run this command to configure your shell: \n# & ./hostutl env --shell=powershell --unset | Invoke-Expression\n"},
		{"powershell", []string{"./hostutl", "env", "--shell=powershell", "--unset"}, "# Run this command to configure your shell: \n# & ./hostutl env --shell=powershell --unset | Invoke-Expression\n"},
		{"powershell", []string{`C:\\Program Files\bhojpur-machine.exe`, "env", "--shell=powershell", "default"}, "# Run this command to configure your shell: \n# & \"C:\\\\Program Files\\bhojpur-machine.exe\" env --shell=powershell default | Invoke-Expression\n"},
		{"powershell", []string{`C:\\Me\bhojpur-machine.exe`, "env", "--shell=powershell", "default"}, "# Run this command to configure your shell: \n# & \"C:\\\\Me\\bhojpur-machine.exe\" env --shell=powershell default | Invoke-Expression\n"},

		{"cmd", []string{"./hostutl", "env", "--shell=cmd", "default"}, "REM Run this command to configure your shell: \nREM \t@FOR /f \"tokens=*\" %i IN ('./hostutl env --shell=cmd default') DO @%i\n"},
		{"cmd", []string{"./hostutl", "env", "--shell=cmd", "--no-proxy", "default"}, "REM Run this command to configure your shell: \nREM \t@FOR /f \"tokens=*\" %i IN ('./hostutl env --shell=cmd --no-proxy default') DO @%i\n"},
		{"cmd", []string{"./hostutl", "env", "--shell=cmd", "--swarm", "default"}, "REM Run this command to configure your shell: \nREM \t@FOR /f \"tokens=*\" %i IN ('./hostutl env --shell=cmd --swarm default') DO @%i\n"},
		{"cmd", []string{"./hostutl", "env", "--shell=cmd", "--no-proxy", "--swarm", "default"}, "REM Run this command to configure your shell: \nREM \t@FOR /f \"tokens=*\" %i IN ('./hostutl env --shell=cmd --no-proxy --swarm default') DO @%i\n"},
		{"cmd", []string{"./hostutl", "env", "--shell=cmd", "--unset"}, "REM Run this command to configure your shell: \nREM \t@FOR /f \"tokens=*\" %i IN ('./hostutl env --shell=cmd --unset') DO @%i\n"},
		{"cmd", []string{`C:\\Program Files\bhojpur-machine.exe`, "env", "--shell=cmd", "default"}, "REM Run this command to configure your shell: \nREM \t@FOR /f \"tokens=*\" %i IN ('\"C:\\\\Program Files\\bhojpur-machine.exe\" env --shell=cmd default') DO @%i\n"},
		{"cmd", []string{`C:\\Me\bhojpur-machine.exe`, "env", "--shell=cmd", "default"}, "REM Run this command to configure your shell: \nREM \t@FOR /f \"tokens=*\" %i IN ('\"C:\\\\Me\\bhojpur-machine.exe\" env --shell=cmd default') DO @%i\n"},

		{"emacs", []string{"./hostutl", "env", "--shell=emacs", "default"}, ";; Run this command to configure your shell: \n;; (with-temp-buffer (shell-command \"./hostutl env --shell=emacs default\" (current-buffer)) (eval-buffer))\n"},
		{"emacs", []string{"./hostutl", "env", "--shell=emacs", "--no-proxy", "default"}, ";; Run this command to configure your shell: \n;; (with-temp-buffer (shell-command \"./hostutl env --shell=emacs --no-proxy default\" (current-buffer)) (eval-buffer))\n"},
		{"emacs", []string{"./hostutl", "env", "--shell=emacs", "--swarm", "default"}, ";; Run this command to configure your shell: \n;; (with-temp-buffer (shell-command \"./hostutl env --shell=emacs --swarm default\" (current-buffer)) (eval-buffer))\n"},
		{"emacs", []string{"./hostutl", "env", "--shell=emacs", "--no-proxy", "--swarm", "default"}, ";; Run this command to configure your shell: \n;; (with-temp-buffer (shell-command \"./hostutl env --shell=emacs --no-proxy --swarm default\" (current-buffer)) (eval-buffer))\n"},
		{"emacs", []string{"./hostutl", "env", "--shell=emacs", "--unset"}, ";; Run this command to configure your shell: \n;; (with-temp-buffer (shell-command \"./hostutl env --shell=emacs --unset\" (current-buffer)) (eval-buffer))\n"},

		{"tcsh", []string{"./hostutl", "env", "--shell=tcsh", "default"}, ": Run this command to configure your shell: \n: eval `./hostutl env --shell=tcsh default`\n"},
		{"tcsh", []string{"./hostutl", "env", "--shell=tcsh", "--no-proxy", "default"}, ": Run this command to configure your shell: \n: eval `./hostutl env --shell=tcsh --no-proxy default`\n"},
		{"tcsh", []string{"./hostutl", "env", "--shell=tcsh", "--swarm", "default"}, ": Run this command to configure your shell: \n: eval `./hostutl env --shell=tcsh --swarm default`\n"},
		{"tcsh", []string{"./hostutl", "env", "--shell=tcsh", "--no-proxy", "--swarm", "default"}, ": Run this command to configure your shell: \n: eval `./hostutl env --shell=tcsh --no-proxy --swarm default`\n"},
		{"tcsh", []string{"./hostutl", "env", "--shell=tcsh", "--unset"}, ": Run this command to configure your shell: \n: eval `./hostutl env --shell=tcsh --unset`\n"},
	}

	for _, test := range tests {
		hints := defaultUsageHinter.GenerateUsageHint(test.userShell, test.commandLine)
		assert.Equal(t, test.expectedHints, hints)
	}
}

func revertUsageHinter(uhg UsageHintGenerator) {
	defaultUsageHinter = uhg
}

func TestShellCfgSet(t *testing.T) {
	const (
		usageHint = "This is a usage hint"
	)

	// TODO: This should be embedded in some kind of wrapper struct for all
	// these `env` operations.
	defer revertUsageHinter(defaultUsageHinter)
	defaultUsageHinter = &SimpleUsageHintGenerator{usageHint}
	isRuntimeWindows := runtimeOS() == "windows"

	var tests = []struct {
		description      string
		commandLine      CommandLine
		api              core.API
		connChecker      check.ConnChecker
		noProxyVar       string
		noProxyValue     string
		expectedShellCfg *ShellConfig
		expectedErr      error
	}{
		{
			description: "no host name specified",
			api: &atest.FakeAPI{
				Hosts: []*host.Host{},
			},
			commandLine: &ctest.FakeCommandLine{
				CliArgs: nil,
			},
			expectedShellCfg: nil,
			expectedErr:      ErrNoDefault,
		},
		{
			description: "bash shell set happy path without any flags set",
			commandLine: &ctest.FakeCommandLine{
				CliArgs: []string{"quux"},
				LocalFlags: &ctest.FakeFlagger{
					Data: map[string]interface{}{
						"shell":    "bash",
						"swarm":    false,
						"no-proxy": false,
					},
				},
			},
			api: &atest.FakeAPI{
				Hosts: []*host.Host{
					{
						Name: "quux",
					},
				},
			},
			connChecker: &FakeConnChecker{
				BhojpurHost: "tcp://1.2.3.4:2376",
				AuthOptions: nil,
				Err:         nil,
			},
			expectedShellCfg: &ShellConfig{
				Prefix:           "export ",
				Delimiter:        "=\"",
				Suffix:           "\"\n",
				BhojpurCertPath:  filepath.Join(mdirs.GetMachineDir(), "quux"),
				BhojpurHost:      "tcp://1.2.3.4:2376",
				BhojpurTLSVerify: "1",
				UsageHint:        usageHint,
				MachineName:      "quux",
				ComposePathsVar:  isRuntimeWindows,
			},
			expectedErr: nil,
		},
		{
			description: "bash shell set happy path with 'default' vm",
			commandLine: &ctest.FakeCommandLine{
				CliArgs: []string{},
				LocalFlags: &ctest.FakeFlagger{
					Data: map[string]interface{}{
						"shell":    "bash",
						"swarm":    false,
						"no-proxy": false,
					},
				},
			},
			api: &atest.FakeAPI{
				Hosts: []*host.Host{
					{
						Name: defaultMachineName,
					},
				},
			},
			connChecker: &FakeConnChecker{
				BhojpurHost: "tcp://1.2.3.4:2376",
				AuthOptions: nil,
				Err:         nil,
			},
			expectedShellCfg: &ShellConfig{
				Prefix:           "export ",
				Delimiter:        "=\"",
				Suffix:           "\"\n",
				BhojpurCertPath:  filepath.Join(mdirs.GetMachineDir(), defaultMachineName),
				BhojpurHost:      "tcp://1.2.3.4:2376",
				BhojpurTLSVerify: "1",
				UsageHint:        usageHint,
				MachineName:      defaultMachineName,
				ComposePathsVar:  isRuntimeWindows,
			},
			expectedErr: nil,
		},
		{
			description: "fish shell set happy path",
			commandLine: &ctest.FakeCommandLine{
				CliArgs: []string{"quux"},
				LocalFlags: &ctest.FakeFlagger{
					Data: map[string]interface{}{
						"shell":    "fish",
						"swarm":    false,
						"no-proxy": false,
					},
				},
			},
			api: &atest.FakeAPI{
				Hosts: []*host.Host{
					{
						Name: "quux",
					},
				},
			},
			connChecker: &FakeConnChecker{
				BhojpurHost: "tcp://1.2.3.4:2376",
				AuthOptions: nil,
				Err:         nil,
			},
			expectedShellCfg: &ShellConfig{
				Prefix:           "set -gx ",
				Suffix:           "\";\n",
				Delimiter:        " \"",
				BhojpurCertPath:  filepath.Join(mdirs.GetMachineDir(), "quux"),
				BhojpurHost:      "tcp://1.2.3.4:2376",
				BhojpurTLSVerify: "1",
				UsageHint:        usageHint,
				MachineName:      "quux",
				ComposePathsVar:  isRuntimeWindows,
			},
			expectedErr: nil,
		},
		{
			description: "powershell set happy path",
			commandLine: &ctest.FakeCommandLine{
				CliArgs: []string{"quux"},
				LocalFlags: &ctest.FakeFlagger{
					Data: map[string]interface{}{
						"shell":    "powershell",
						"swarm":    false,
						"no-proxy": false,
					},
				},
			},
			api: &atest.FakeAPI{
				Hosts: []*host.Host{
					{
						Name: "quux",
					},
				},
			},
			connChecker: &FakeConnChecker{
				BhojpurHost: "tcp://1.2.3.4:2376",
				AuthOptions: nil,
				Err:         nil,
			},
			expectedShellCfg: &ShellConfig{
				Prefix:           "$Env:",
				Suffix:           "\"\n",
				Delimiter:        " = \"",
				BhojpurCertPath:  filepath.Join(mdirs.GetMachineDir(), "quux"),
				BhojpurHost:      "tcp://1.2.3.4:2376",
				BhojpurTLSVerify: "1",
				UsageHint:        usageHint,
				MachineName:      "quux",
				ComposePathsVar:  isRuntimeWindows,
			},
			expectedErr: nil,
		},
		{
			description: "emacs set happy path",
			commandLine: &ctest.FakeCommandLine{
				CliArgs: []string{"quux"},
				LocalFlags: &ctest.FakeFlagger{
					Data: map[string]interface{}{
						"shell":    "emacs",
						"swarm":    false,
						"no-proxy": false,
					},
				},
			},
			api: &atest.FakeAPI{
				Hosts: []*host.Host{
					{
						Name: "quux",
					},
				},
			},
			connChecker: &FakeConnChecker{
				BhojpurHost: "tcp://1.2.3.4:2376",
				AuthOptions: nil,
				Err:         nil,
			},
			expectedShellCfg: &ShellConfig{
				Prefix:           "(setenv \"",
				Suffix:           "\")\n",
				Delimiter:        "\" \"",
				BhojpurCertPath:  filepath.Join(mdirs.GetMachineDir(), "quux"),
				BhojpurHost:      "tcp://1.2.3.4:2376",
				BhojpurTLSVerify: "1",
				UsageHint:        usageHint,
				MachineName:      "quux",
				ComposePathsVar:  isRuntimeWindows,
			},
			expectedErr: nil,
		},
		{
			description: "cmd.exe happy path",
			commandLine: &ctest.FakeCommandLine{
				CliArgs: []string{"quux"},
				LocalFlags: &ctest.FakeFlagger{
					Data: map[string]interface{}{
						"shell":    "cmd",
						"swarm":    false,
						"no-proxy": false,
					},
				},
			},
			api: &atest.FakeAPI{
				Hosts: []*host.Host{
					{
						Name: "quux",
					},
				},
			},
			connChecker: &FakeConnChecker{
				BhojpurHost: "tcp://1.2.3.4:2376",
				AuthOptions: nil,
				Err:         nil,
			},
			expectedShellCfg: &ShellConfig{
				Prefix:           "SET ",
				Suffix:           "\n",
				Delimiter:        "=",
				BhojpurCertPath:  filepath.Join(mdirs.GetMachineDir(), "quux"),
				BhojpurHost:      "tcp://1.2.3.4:2376",
				BhojpurTLSVerify: "1",
				UsageHint:        usageHint,
				MachineName:      "quux",
				ComposePathsVar:  isRuntimeWindows,
			},
			expectedErr: nil,
		},
		{
			description: "bash shell set happy path with --no-proxy flag; no existing environment variable set",
			commandLine: &ctest.FakeCommandLine{
				CliArgs: []string{"quux"},
				LocalFlags: &ctest.FakeFlagger{
					Data: map[string]interface{}{
						"shell":    "bash",
						"swarm":    false,
						"no-proxy": true,
					},
				},
			},
			api: &atest.FakeAPI{
				Hosts: []*host.Host{
					{
						Name: "quux",
						Driver: &fakedriver.Driver{
							MockState: state.Running,
							MockIP:    "1.2.3.4",
						},
					},
				},
			},
			connChecker: &FakeConnChecker{
				BhojpurHost: "tcp://1.2.3.4:2376",
				AuthOptions: nil,
				Err:         nil,
			},
			expectedShellCfg: &ShellConfig{
				Prefix:           "export ",
				Delimiter:        "=\"",
				Suffix:           "\"\n",
				BhojpurCertPath:  filepath.Join(mdirs.GetMachineDir(), "quux"),
				BhojpurHost:      "tcp://1.2.3.4:2376",
				BhojpurTLSVerify: "1",
				UsageHint:        usageHint,
				NoProxyVar:       "NO_PROXY",
				NoProxyValue:     "1.2.3.4", // From FakeDriver
				MachineName:      "quux",
				ComposePathsVar:  isRuntimeWindows,
			},
			noProxyVar:   "NO_PROXY",
			noProxyValue: "",
			expectedErr:  nil,
		},
		{
			description: "bash shell set happy path with --no-proxy flag; existing environment variable _is_ set",
			commandLine: &ctest.FakeCommandLine{
				CliArgs: []string{"quux"},
				LocalFlags: &ctest.FakeFlagger{
					Data: map[string]interface{}{
						"shell":    "bash",
						"swarm":    false,
						"no-proxy": true,
					},
				},
			},
			api: &atest.FakeAPI{
				Hosts: []*host.Host{
					{
						Name: "quux",
						Driver: &fakedriver.Driver{
							MockState: state.Running,
							MockIP:    "1.2.3.4",
						},
					},
				},
			},
			connChecker: &FakeConnChecker{
				BhojpurHost: "tcp://1.2.3.4:2376",
				AuthOptions: nil,
				Err:         nil,
			},
			expectedShellCfg: &ShellConfig{
				Prefix:           "export ",
				Delimiter:        "=\"",
				Suffix:           "\"\n",
				BhojpurCertPath:  filepath.Join(mdirs.GetMachineDir(), "quux"),
				BhojpurHost:      "tcp://1.2.3.4:2376",
				BhojpurTLSVerify: "1",
				UsageHint:        usageHint,
				NoProxyVar:       "no_proxy",
				NoProxyValue:     "192.168.59.1,1.2.3.4", // From FakeDriver
				MachineName:      "quux",
				ComposePathsVar:  isRuntimeWindows,
			},
			noProxyVar:   "no_proxy",
			noProxyValue: "192.168.59.1",
			expectedErr:  nil,
		},
	}

	for _, test := range tests {
		// TODO: Ideally this should not hit the environment at all but
		// rather should go through an interface.
		os.Setenv(test.noProxyVar, test.noProxyValue)

		t.Log(test.description)

		check.DefaultConnChecker = test.connChecker
		shellCfg, err := shellCfgSet(test.commandLine, test.api)
		assert.Equal(t, test.expectedShellCfg, shellCfg)
		assert.Equal(t, test.expectedErr, err)

		os.Unsetenv(test.noProxyVar)
	}
}

func TestShellCfgSetWindowsRuntime(t *testing.T) {
	const (
		usageHint = "This is a usage hint"
	)

	// TODO: This should be embedded in some kind of wrapper struct for all
	// these `env` operations.
	defer revertUsageHinter(defaultUsageHinter)
	defaultUsageHinter = &SimpleUsageHintGenerator{usageHint}

	var tests = []struct {
		description      string
		commandLine      CommandLine
		api              core.API
		connChecker      check.ConnChecker
		noProxyVar       string
		noProxyValue     string
		expectedShellCfg *ShellConfig
		expectedErr      error
	}{
		{
			description: "powershell set happy path",
			commandLine: &ctest.FakeCommandLine{
				CliArgs: []string{"quux"},
				LocalFlags: &ctest.FakeFlagger{
					Data: map[string]interface{}{
						"shell":    "powershell",
						"swarm":    false,
						"no-proxy": false,
					},
				},
			},
			api: &atest.FakeAPI{
				Hosts: []*host.Host{
					{
						Name: "quux",
					},
				},
			},
			connChecker: &FakeConnChecker{
				BhojpurHost: "tcp://1.2.3.4:2376",
				AuthOptions: nil,
				Err:         nil,
			},
			expectedShellCfg: &ShellConfig{
				Prefix:           "$Env:",
				Suffix:           "\"\n",
				Delimiter:        " = \"",
				BhojpurCertPath:  filepath.Join(mdirs.GetMachineDir(), "quux"),
				BhojpurHost:      "tcp://1.2.3.4:2376",
				BhojpurTLSVerify: "1",
				UsageHint:        usageHint,
				MachineName:      "quux",
				ComposePathsVar:  true,
			},
			expectedErr: nil,
		},
	}

	actualRuntimeOS := runtimeOS
	runtimeOS = func() string { return "windows" }
	defer func() { runtimeOS = actualRuntimeOS }()

	for _, test := range tests {
		// TODO: Ideally this should not hit the environment at all but
		// rather should go through an interface.
		os.Setenv(test.noProxyVar, test.noProxyValue)

		t.Log(test.description)

		check.DefaultConnChecker = test.connChecker
		shellCfg, err := shellCfgSet(test.commandLine, test.api)
		assert.Equal(t, test.expectedShellCfg, shellCfg)
		assert.Equal(t, test.expectedErr, err)

		os.Unsetenv(test.noProxyVar)
	}
}

func TestShellCfgUnset(t *testing.T) {
	const (
		usageHint = "This is the unset usage hint"
	)

	defer revertUsageHinter(defaultUsageHinter)
	defaultUsageHinter = &SimpleUsageHintGenerator{usageHint}

	var tests = []struct {
		description      string
		commandLine      CommandLine
		api              core.API
		connChecker      check.ConnChecker
		noProxyVar       string
		noProxyValue     string
		expectedShellCfg *ShellConfig
		expectedErr      error
	}{
		{
			description: "more than expected args passed in",
			commandLine: &ctest.FakeCommandLine{
				CliArgs: []string{"foo", "bar"},
			},
			expectedShellCfg: nil,
			expectedErr:      errImproperUnsetEnvArgs,
		},
		{
			description: "bash shell unset happy path without any flags set",
			commandLine: &ctest.FakeCommandLine{
				CliArgs: nil,
				LocalFlags: &ctest.FakeFlagger{
					Data: map[string]interface{}{
						"shell":    "bash",
						"swarm":    false,
						"no-proxy": false,
					},
				},
			},
			api: &atest.FakeAPI{},
			connChecker: &FakeConnChecker{
				BhojpurHost: "tcp://1.2.3.4:2376",
				AuthOptions: nil,
				Err:         nil,
			},
			expectedShellCfg: &ShellConfig{
				Prefix:    "unset ",
				Suffix:    "\n",
				Delimiter: "",
				UsageHint: usageHint,
			},
			expectedErr: nil,
		},
		{
			description: "fish shell unset happy path",
			commandLine: &ctest.FakeCommandLine{
				CliArgs: nil,
				LocalFlags: &ctest.FakeFlagger{
					Data: map[string]interface{}{
						"shell":    "fish",
						"swarm":    false,
						"no-proxy": false,
					},
				},
			},
			api: &atest.FakeAPI{},
			connChecker: &FakeConnChecker{
				BhojpurHost: "tcp://1.2.3.4:2376",
				AuthOptions: nil,
				Err:         nil,
			},
			expectedShellCfg: &ShellConfig{
				Prefix:    "set -e ",
				Suffix:    ";\n",
				Delimiter: "",
				UsageHint: usageHint,
			},
			expectedErr: nil,
		},
		{
			description: "powershell unset happy path",
			commandLine: &ctest.FakeCommandLine{
				CliArgs: nil,
				LocalFlags: &ctest.FakeFlagger{
					Data: map[string]interface{}{
						"shell":    "powershell",
						"swarm":    false,
						"no-proxy": false,
					},
				},
			},
			api: &atest.FakeAPI{},
			connChecker: &FakeConnChecker{
				BhojpurHost: "tcp://1.2.3.4:2376",
				AuthOptions: nil,
				Err:         nil,
			},
			expectedShellCfg: &ShellConfig{
				Prefix:    `Remove-Item Env:\\`,
				Suffix:    "\n",
				Delimiter: "",
				UsageHint: usageHint,
			},
			expectedErr: nil,
		},
		{
			description: "cmd.exe unset happy path",
			commandLine: &ctest.FakeCommandLine{
				CliArgs: nil,
				LocalFlags: &ctest.FakeFlagger{
					Data: map[string]interface{}{
						"shell":    "cmd",
						"swarm":    false,
						"no-proxy": false,
					},
				},
			},
			api: &atest.FakeAPI{},
			connChecker: &FakeConnChecker{
				BhojpurHost: "tcp://1.2.3.4:2376",
				AuthOptions: nil,
				Err:         nil,
			},
			expectedShellCfg: &ShellConfig{
				Prefix:    "SET ",
				Suffix:    "\n",
				Delimiter: "=",
				UsageHint: usageHint,
			},
			expectedErr: nil,
		},
		// TODO: There is kind of a funny bug (feature?) I discovered
		// reasoning about unset() where if there was a NO_PROXY value
		// set _before_ the original bhojpur-machine env, it won't be
		// restored (NO_PROXY won't be unset at all, it will stay the
		// same).  We should define expected behavior in this case.
	}

	for _, test := range tests {
		os.Setenv(test.noProxyVar, test.noProxyValue)

		t.Log(test.description)

		check.DefaultConnChecker = test.connChecker
		shellCfg, err := shellCfgUnset(test.commandLine, test.api)
		assert.Equal(t, test.expectedShellCfg, shellCfg)
		assert.Equal(t, test.expectedErr, err)

		os.Setenv(test.noProxyVar, "")
	}
}

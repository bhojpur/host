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
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"text/template"

	mdirs "github.com/bhojpur/host/cmd/machine/commands/dirs"
	"github.com/bhojpur/host/pkg/core"
	"github.com/bhojpur/host/pkg/core/check"
	"github.com/bhojpur/host/pkg/core/log"
	"github.com/bhojpur/host/pkg/core/shell"
)

const (
	envTmpl = `{{ .Prefix }}DOCKER_TLS_VERIFY{{ .Delimiter }}{{ .BhojpurTLSVerify }}{{ .Suffix }}{{ .Prefix }}DOCKER_HOST{{ .Delimiter }}{{ .BhojpurHost }}{{ .Suffix }}{{ .Prefix }}DOCKER_CERT_PATH{{ .Delimiter }}{{ .BhojpurCertPath }}{{ .Suffix }}{{ .Prefix }}DOCKER_MACHINE_NAME{{ .Delimiter }}{{ .MachineName }}{{ .Suffix }}{{ if .ComposePathsVar }}{{ .Prefix }}COMPOSE_CONVERT_WINDOWS_PATHS{{ .Delimiter }}true{{ .Suffix }}{{end}}{{ if .NoProxyVar }}{{ .Prefix }}{{ .NoProxyVar }}{{ .Delimiter }}{{ .NoProxyValue }}{{ .Suffix }}{{end}}{{ .UsageHint }}`
)

var (
	errImproperUnsetEnvArgs = errors.New("Error: Expected no Bhojpur Host machine name when the -u flag is present")
	defaultUsageHinter      UsageHintGenerator
	runtimeOS               = func() string { return runtime.GOOS }
)

func init() {
	defaultUsageHinter = &EnvUsageHintGenerator{}
}

type ShellConfig struct {
	Prefix           string
	Delimiter        string
	Suffix           string
	BhojpurCertPath  string
	BhojpurHost      string
	BhojpurTLSVerify string
	UsageHint        string
	MachineName      string
	NoProxyVar       string
	NoProxyValue     string
	ComposePathsVar  bool
}

func cmdEnv(c CommandLine, api core.API) error {
	var (
		err      error
		shellCfg *ShellConfig
	)

	// Ensure that log messages always go to stderr when this command is
	// being run (it is intended to be run in a subshell)
	log.SetOutWriter(os.Stderr)

	if c.Bool("unset") {
		shellCfg, err = shellCfgUnset(c, api)
		if err != nil {
			return err
		}
	} else {
		shellCfg, err = shellCfgSet(c, api)
		if err != nil {
			return err
		}
	}

	return executeTemplateStdout(shellCfg)
}

func shellCfgSet(c CommandLine, api core.API) (*ShellConfig, error) {
	if len(c.Args()) > 1 {
		return nil, ErrExpectedOneMachine
	}

	target, err := targetHost(c, api)
	if err != nil {
		return nil, err
	}

	host, err := api.Load(target)
	if err != nil {
		return nil, err
	}

	bhojpurHost, _, err := check.DefaultConnChecker.Check(host, c.Bool("swarm"))
	if err != nil {
		return nil, fmt.Errorf("Error checking TLS connection: %s", err)
	}

	userShell, err := getShell(c.String("shell"))
	if err != nil {
		return nil, err
	}

	shellCfg := &ShellConfig{
		BhojpurCertPath:  filepath.Join(mdirs.GetMachineDir(), host.Name),
		BhojpurHost:      bhojpurHost,
		BhojpurTLSVerify: "1",
		UsageHint:        defaultUsageHinter.GenerateUsageHint(userShell, os.Args),
		MachineName:      host.Name,
	}

	if c.Bool("no-proxy") {
		ip, err := host.Driver.GetIP()
		if err != nil {
			return nil, fmt.Errorf("Error getting host IP: %s", err)
		}

		noProxyVar, noProxyValue := findNoProxyFromEnv()

		// add the Bhojpur Host to the no_proxy list idempotently
		switch {
		case noProxyValue == "":
			noProxyValue = ip
		case strings.Contains(noProxyValue, ip):
		//ip already in no_proxy list, nothing to do
		default:
			noProxyValue = fmt.Sprintf("%s,%s", noProxyValue, ip)
		}

		shellCfg.NoProxyVar = noProxyVar
		shellCfg.NoProxyValue = noProxyValue
	}

	if runtimeOS() == "windows" {
		shellCfg.ComposePathsVar = true
	}

	switch userShell {
	case "fish":
		shellCfg.Prefix = "set -gx "
		shellCfg.Suffix = "\";\n"
		shellCfg.Delimiter = " \""
	case "powershell":
		shellCfg.Prefix = "$Env:"
		shellCfg.Suffix = "\"\n"
		shellCfg.Delimiter = " = \""
	case "cmd":
		shellCfg.Prefix = "SET "
		shellCfg.Suffix = "\n"
		shellCfg.Delimiter = "="
	case "tcsh":
		shellCfg.Prefix = "setenv "
		shellCfg.Suffix = "\";\n"
		shellCfg.Delimiter = " \""
	case "emacs":
		shellCfg.Prefix = "(setenv \""
		shellCfg.Suffix = "\")\n"
		shellCfg.Delimiter = "\" \""
	default:
		shellCfg.Prefix = "export "
		shellCfg.Suffix = "\"\n"
		shellCfg.Delimiter = "=\""
	}

	return shellCfg, nil
}

func shellCfgUnset(c CommandLine, api core.API) (*ShellConfig, error) {
	if len(c.Args()) != 0 {
		return nil, errImproperUnsetEnvArgs
	}

	userShell, err := getShell(c.String("shell"))
	if err != nil {
		return nil, err
	}

	shellCfg := &ShellConfig{
		UsageHint: defaultUsageHinter.GenerateUsageHint(userShell, os.Args),
	}

	if c.Bool("no-proxy") {
		shellCfg.NoProxyVar, shellCfg.NoProxyValue = findNoProxyFromEnv()
	}

	switch userShell {
	case "fish":
		shellCfg.Prefix = "set -e "
		shellCfg.Suffix = ";\n"
		shellCfg.Delimiter = ""
	case "powershell":
		shellCfg.Prefix = `Remove-Item Env:\\`
		shellCfg.Suffix = "\n"
		shellCfg.Delimiter = ""
	case "cmd":
		shellCfg.Prefix = "SET "
		shellCfg.Suffix = "\n"
		shellCfg.Delimiter = "="
	case "emacs":
		shellCfg.Prefix = "(setenv \""
		shellCfg.Suffix = ")\n"
		shellCfg.Delimiter = "\" nil"
	case "tcsh":
		shellCfg.Prefix = "unsetenv "
		shellCfg.Suffix = ";\n"
		shellCfg.Delimiter = ""
	default:
		shellCfg.Prefix = "unset "
		shellCfg.Suffix = "\n"
		shellCfg.Delimiter = ""
	}

	return shellCfg, nil
}

func executeTemplateStdout(shellCfg *ShellConfig) error {
	t := template.New("envConfig")
	tmpl, err := t.Parse(envTmpl)
	if err != nil {
		return err
	}

	return tmpl.Execute(os.Stdout, shellCfg)
}

func getShell(userShell string) (string, error) {
	if userShell != "" {
		return userShell, nil
	}
	return shell.Detect()
}

func findNoProxyFromEnv() (string, string) {
	// first check for an existing lower case no_proxy var
	noProxyVar := "no_proxy"
	noProxyValue := os.Getenv("no_proxy")

	// otherwise default to allcaps HTTP_PROXY
	if noProxyValue == "" {
		noProxyVar = "NO_PROXY"
		noProxyValue = os.Getenv("NO_PROXY")
	}
	return noProxyVar, noProxyValue
}

type UsageHintGenerator interface {
	GenerateUsageHint(string, []string) string
}

type EnvUsageHintGenerator struct{}

func (g *EnvUsageHintGenerator) GenerateUsageHint(userShell string, args []string) string {
	cmd := ""
	comment := "#"

	bhojpurMachinePath := args[0]
	if strings.Contains(bhojpurMachinePath, " ") || strings.Contains(bhojpurMachinePath, `\`) {
		args[0] = fmt.Sprintf("\"%s\"", bhojpurMachinePath)
	}

	commandLine := strings.Join(args, " ")

	switch userShell {
	case "fish":
		cmd = fmt.Sprintf("eval (%s)", commandLine)
	case "powershell":
		cmd = fmt.Sprintf("& %s | Invoke-Expression", commandLine)
	case "cmd":
		cmd = fmt.Sprintf("\t@FOR /f \"tokens=*\" %%i IN ('%s') DO @%%i", commandLine)
		comment = "REM"
	case "emacs":
		cmd = fmt.Sprintf("(with-temp-buffer (shell-command \"%s\" (current-buffer)) (eval-buffer))", commandLine)
		comment = ";;"
	case "tcsh":
		cmd = fmt.Sprintf("eval `%s`", commandLine)
		comment = ":"
	default:
		cmd = fmt.Sprintf("eval $(%s)", commandLine)
	}

	return fmt.Sprintf("%s Run this command to configure your shell: \n%s %s\n", comment, comment, cmd)
}

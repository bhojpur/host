package hyperv

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
	"bufio"
	"bytes"
	"errors"
	"os/exec"
	"strings"

	"fmt"

	"github.com/bhojpur/host/pkg/machine/log"
)

var powershell string

var (
	ErrPowerShellNotFound = errors.New("Powershell was not found in the path")
	ErrNotAdministrator   = errors.New("Hyper-v commands have to be run as an Administrator")
	ErrNotInstalled       = errors.New("Hyper-V PowerShell Module is not available")
)

func init() {
	powershell, _ = exec.LookPath("powershell.exe")
}

func cmdOut(args ...string) (string, error) {
	args = append([]string{"-NoProfile", "-NonInteractive"}, args...)
	cmd := exec.Command(powershell, args...)
	log.Debugf("[executing ==>] : %v %v", powershell, strings.Join(args, " "))
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	log.Debugf("[stdout =====>] : %s", stdout.String())
	log.Debugf("[stderr =====>] : %s", stderr.String())
	return stdout.String(), err
}

func cmd(args ...string) error {
	_, err := cmdOut(args...)
	return err
}

func parseLines(stdout string) []string {
	resp := []string{}

	s := bufio.NewScanner(strings.NewReader(stdout))
	for s.Scan() {
		resp = append(resp, s.Text())
	}

	return resp
}

func hypervAvailable() error {
	stdout, err := cmdOut("@(Get-Module -ListAvailable hyper-v).Name | Get-Unique")
	if err != nil {
		return err
	}

	resp := parseLines(stdout)
	if resp[0] != "Hyper-V" {
		return ErrNotInstalled
	}

	return nil
}

func isAdministrator() (bool, error) {
	hypervAdmin := isHypervAdministrator()

	if hypervAdmin {
		return true, nil
	}

	windowsAdmin, err := isWindowsAdministrator()

	if err != nil {
		return false, err
	}

	return windowsAdmin, nil
}

func isHypervAdministrator() bool {
	stdout, err := cmdOut(`@([Security.Principal.WindowsPrincipal][Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole("S-1-5-32-578")`)
	if err != nil {
		log.Debug(err)
		return false
	}

	resp := parseLines(stdout)
	return resp[0] == "True"
}

func isWindowsAdministrator() (bool, error) {
	stdout, err := cmdOut(`@([Security.Principal.WindowsPrincipal][Security.Principal.WindowsIdentity]::GetCurrent()).IsInRole([Security.Principal.WindowsBuiltInRole] "Administrator")`)
	if err != nil {
		return false, err
	}

	resp := parseLines(stdout)
	return resp[0] == "True", nil
}

func quote(text string) string {
	return fmt.Sprintf("'%s'", text)
}

func toMb(value int) string {
	return fmt.Sprintf("%dMB", value)
}

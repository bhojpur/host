package provision

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
	"fmt"

	"github.com/bhojpur/host/pkg/core/drivers"
	"github.com/bhojpur/host/pkg/core/log"
	"github.com/bhojpur/host/pkg/core/ssh"
)

type RedHatSSHCommander struct {
	Driver drivers.Driver
}

func (sshCmder RedHatSSHCommander) SSHCommand(args string) (string, error) {
	client, err := drivers.GetSSHClientFromDriver(sshCmder.Driver)
	if err != nil {
		return "", err
	}

	log.Debugf("About to run SSH command:\n%s", args)

	// redhat needs "-t" for tty allocation on ssh therefore we check for the
	// external client and add as needed.
	// Note: CentOS 7.0 needs multiple "-tt" to force tty allocation when ssh has
	// no local tty.
	var output string
	switch c := client.(type) {
	case *ssh.ExternalClient:
		c.BaseArgs = append(c.BaseArgs, "-tt")
		output, err = c.Output(args)
	case *ssh.NativeClient:
		output, err = c.OutputWithPty(args)
	}

	log.Debugf("SSH cmd err, output: %v: %s", err, output)
	if err != nil {
		return "", fmt.Errorf(`RHEL ssh command error: command: %s err: %v output: %s`, args, err, output)
	}

	return output, nil
}

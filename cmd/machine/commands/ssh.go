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
	"fmt"

	core "github.com/bhojpur/host/pkg/machine"
	"github.com/bhojpur/host/pkg/machine/state"
)

type errStateInvalidForSSH struct {
	HostName string
}

func (e errStateInvalidForSSH) Error() string {
	return fmt.Sprintf("Error: Cannot run SSH command: Host %q is not running", e.HostName)
}

func cmdSSH(c CommandLine, api core.API) error {
	// Check for help flag -- Needed due to SkipFlagParsing
	firstArg := c.Args().First()
	if firstArg == "-help" || firstArg == "--help" || firstArg == "-h" {
		c.ShowHelp()
		return nil
	}

	target, err := targetHost(c, api)
	if err != nil {
		return err
	}

	host, err := api.Load(target)
	if err != nil {
		return err
	}

	currentState, err := host.Driver.GetState()
	if err != nil {
		return err
	}

	if currentState != state.Running {
		return errStateInvalidForSSH{host.Name}
	}

	client, err := host.CreateSSHClient()
	if err != nil {
		return err
	}

	return client.Shell(c.Args().Tail()...)
}

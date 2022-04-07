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
	"strings"

	core "github.com/bhojpur/host/pkg/machine"
	merrors "github.com/bhojpur/host/pkg/machine/errors"
	"github.com/bhojpur/host/pkg/machine/log"
)

func cmdRm(c CommandLine, api core.API) error {
	if len(c.Args()) == 0 {
		c.ShowHelp()
		return ErrNoMachineSpecified
	}

	log.Info(fmt.Sprintf("About to remove Bhojpur Host: %s", strings.Join(c.Args(), ", ")))
	log.Warn("WARNING: This action will delete both local reference and remote instance.")

	force := c.Bool("force")
	confirm := c.Bool("y")
	var errorOccurred []string

	if !userConfirm(confirm, force) {
		return nil
	}

	for _, hostName := range c.Args() {
		err := removeRemoteMachine(hostName, api)
		if err != nil {
			if _, ok := err.(merrors.ErrHostDoesNotExist); !ok {
				errorOccurred = collectError(fmt.Sprintf("Error removing Bhojpur Host %q: %s", hostName, err), force, errorOccurred)
			} else {
				log.Infof("Bhojpur Host machine config for %s does not exists, so nothing to do...", hostName)
			}
		}

		if err == nil || force {
			removeErr := removeLocalMachine(hostName, api)
			if removeErr != nil {
				errorOccurred = collectError(fmt.Sprintf("Can't remove \"%s\"", hostName), force, errorOccurred)
			} else {
				log.Infof("Successfully removed Bhojpur Host %s", hostName)
			}
		}
	}

	if len(errorOccurred) > 0 && !force {
		return errors.New(strings.Join(errorOccurred, "\n"))
	}

	return nil
}

func userConfirm(confirm bool, force bool) bool {
	if confirm || force {
		return true
	}

	sure, err := confirmInput(fmt.Sprintf("Are you sure?"))
	if err != nil {
		return false
	}

	return sure
}

func removeRemoteMachine(hostName string, api core.API) error {
	currentHost, loaderr := api.Load(hostName)
	if loaderr != nil {
		return loaderr
	}

	err := currentHost.Driver.Remove()
	if err != nil && !strings.Contains(strings.ToLower(err.Error()), "not found") {
		return err
	}

	return nil
}

func removeLocalMachine(hostName string, api core.API) error {
	exist, _ := api.Exists(hostName)
	if !exist {
		return errors.New(hostName + " does not exist.")
	}
	return api.Remove(hostName)
}

func collectError(message string, force bool, errorOccurred []string) []string {
	if force {
		log.Error(message)
	}
	return append(errorOccurred, message)
}

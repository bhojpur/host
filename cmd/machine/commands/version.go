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
	"io"
	"os"

	core "github.com/bhojpur/host/pkg/machine"
	cengine "github.com/bhojpur/host/pkg/machine/client"
)

func cmdVersion(c CommandLine, api core.API) error {
	return printVersion(c, api, os.Stdout)
}

func printVersion(c CommandLine, api core.API, out io.Writer) error {
	if len(c.Args()) == 0 {
		c.ShowVersion()
		return nil
	}

	if len(c.Args()) != 1 {
		return ErrExpectedOneMachine
	}

	host, err := api.Load(c.Args().First())
	if err != nil {
		return err
	}

	if host.HostOptions.AuthOptions != nil {
		version, err := cengine.BhojpurVersion(host)
		if err != nil {
			return err
		}

		fmt.Fprintln(out, version)
	} else {
		fmt.Fprintln(out, "Bhojpur Host was not installed on machine")
	}

	return nil
}

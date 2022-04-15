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

	"time"

	core "github.com/bhojpur/host/pkg/machine"
	"github.com/bhojpur/host/pkg/machine/persist"
	"github.com/bhojpur/host/pkg/machine/state"
)

const (
	activeDefaultTimeout = 10
)

var (
	errNoActiveHost  = errors.New("No active Bhojpur Host found")
	errActiveTimeout = errors.New("Error getting active Bhojpur Host: timeout")
)

func cmdActive(c CommandLine, api core.API) error {
	if len(c.Args()) > 0 {
		return ErrTooManyArguments
	}

	hosts, hostsInError, err := persist.LoadAllHosts(api)
	if err != nil {
		return fmt.Errorf("Error getting active Bhojpur Host: %s", err)
	}

	timeout := time.Duration(c.Int("timeout")) * time.Second
	items := getHostListItems(hosts, hostsInError, timeout)

	active, err := activeHost(items)

	if err != nil {
		return err
	}

	fmt.Println(active.Name)
	return nil
}

func activeHost(items []HostListItem) (HostListItem, error) {
	timeout := false
	for _, item := range items {
		if item.ActiveHost || item.ActiveSwarm {
			return item, nil
		}
		if item.State == state.Timeout {
			timeout = true
		}
	}
	if timeout {
		return HostListItem{}, errActiveTimeout
	}
	return HostListItem{}, errNoActiveHost
}

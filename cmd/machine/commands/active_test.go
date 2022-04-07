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

	"github.com/bhojpur/host/pkg/machine/state"
	"github.com/stretchr/testify/assert"
)

func TestCmdActiveNone(t *testing.T) {
	hostListItems := []HostListItem{
		{
			Name:        "host1",
			ActiveHost:  false,
			ActiveSwarm: false,
			State:       state.Running,
		},
		{
			Name:        "host2",
			ActiveHost:  false,
			ActiveSwarm: false,
			State:       state.Running,
		},
		{
			Name:        "host3",
			ActiveHost:  false,
			ActiveSwarm: false,
			State:       state.Running,
		},
	}
	_, err := activeHost(hostListItems)
	assert.Equal(t, err, errNoActiveHost)
}

func TestCmdActiveHost(t *testing.T) {
	hostListItems := []HostListItem{
		{
			Name:        "host1",
			ActiveHost:  false,
			ActiveSwarm: false,
			State:       state.Timeout,
		},
		{
			Name:        "host2",
			ActiveHost:  true,
			ActiveSwarm: false,
			State:       state.Running,
		},
		{
			Name:        "host3",
			ActiveHost:  false,
			ActiveSwarm: false,
			State:       state.Running,
		},
	}
	active, err := activeHost(hostListItems)
	assert.Equal(t, err, nil)
	assert.Equal(t, active.Name, "host2")
}

func TestCmdActiveSwarm(t *testing.T) {
	hostListItems := []HostListItem{
		{
			Name:        "host1",
			ActiveHost:  false,
			ActiveSwarm: false,
			State:       state.Running,
		},
		{
			Name:        "host2",
			ActiveHost:  false,
			ActiveSwarm: false,
			State:       state.Running,
		},
		{
			Name:        "host3",
			ActiveHost:  false,
			ActiveSwarm: true,
			State:       state.Running,
		},
	}
	active, err := activeHost(hostListItems)
	assert.Equal(t, err, nil)
	assert.Equal(t, active.Name, "host3")
}

func TestCmdActiveTimeout(t *testing.T) {
	hostListItems := []HostListItem{
		{
			Name:        "host1",
			ActiveHost:  false,
			ActiveSwarm: false,
			State:       state.Running,
		},
		{
			Name:        "host2",
			ActiveHost:  false,
			ActiveSwarm: false,
			State:       state.Running,
		},
		{
			Name:        "host3",
			ActiveHost:  false,
			ActiveSwarm: false,
			State:       state.Timeout,
		},
	}
	_, err := activeHost(hostListItems)
	assert.Equal(t, err, errActiveTimeout)
}

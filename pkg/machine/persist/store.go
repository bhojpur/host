package persist

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
	"github.com/bhojpur/host/pkg/machine/host"
)

type Store interface {
	// Exists returns whether a machine exists or not
	Exists(name string) (bool, error)

	// List returns a list of all hosts in the store
	List() ([]string, error)

	// Load loads a host by name
	Load(name string) (*host.Host, error)

	// Remove removes a machine from the store
	Remove(name string) error

	// Save persists a machine in the store
	Save(host *host.Host) error

	// GetMachinesDir gets the location on disk where the machine configs will be stored
	GetMachinesDir() string
}

func LoadHosts(s Store, hostNames []string) ([]*host.Host, map[string]error) {
	loadedHosts := []*host.Host{}
	errors := map[string]error{}

	for _, hostName := range hostNames {
		h, err := s.Load(hostName)
		if err != nil {
			errors[hostName] = err
		} else {
			loadedHosts = append(loadedHosts, h)
		}
	}

	return loadedHosts, errors
}

func LoadAllHosts(s Store) ([]*host.Host, map[string]error, error) {
	hostNames, err := s.List()
	if err != nil {
		return nil, nil, err
	}
	loadedHosts, hostInError := LoadHosts(s, hostNames)
	return loadedHosts, hostInError, nil
}

package persisttest

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

import "github.com/bhojpur/host/pkg/machine/host"

type FakeStore struct {
	Hosts                                           []*host.Host
	ExistsErr, ListErr, LoadErr, RemoveErr, SaveErr error
}

func (fs *FakeStore) Exists(name string) (bool, error) {
	if fs.ExistsErr != nil {
		return false, fs.ExistsErr
	}
	for _, h := range fs.Hosts {
		if h.Name == name {
			return true, nil
		}
	}

	return false, nil
}

func (fs *FakeStore) List() ([]string, error) {
	names := []string{}
	for _, h := range fs.Hosts {
		names = append(names, h.Name)
	}
	return names, fs.ListErr
}

func (fs *FakeStore) Load(name string) (*host.Host, error) {
	if fs.LoadErr != nil {
		return nil, fs.LoadErr
	}
	for _, h := range fs.Hosts {
		if h.Name == name {
			return h, nil
		}
	}

	return nil, nil
}

func (fs *FakeStore) Remove(name string) error {
	if fs.RemoveErr != nil {
		return fs.RemoveErr
	}
	for i, h := range fs.Hosts {
		if h.Name == name {
			fs.Hosts = append(fs.Hosts[:i], fs.Hosts[i+1:]...)
			return nil
		}
	}
	return nil
}

func (fs *FakeStore) Save(host *host.Host) error {
	if fs.SaveErr == nil {
		fs.Hosts = append(fs.Hosts, host)
		return nil
	}
	return fs.SaveErr
}

func (fs *FakeStore) GetMachinesDir() string {
	return ""
}

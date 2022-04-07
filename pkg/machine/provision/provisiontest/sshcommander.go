package provisiontest

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

// It provides utilities for testing provisioners

import "errors"

//FakeSSHCommanderOptions is intended to create a FakeSSHCommander without actually knowing the underlying sshcommands by passing it to NewSSHCommander
type FakeSSHCommanderOptions struct {
	//Result of the ssh command to look up the FilesystemType
	FilesystemType string
}

//FakeSSHCommander is an implementation of provision.SSHCommander to provide predictable responses set by testing code
//Extend it when needed
type FakeSSHCommander struct {
	Responses map[string]string
}

//NewFakeSSHCommander creates a FakeSSHCommander without actually knowing the underlying sshcommands
func NewFakeSSHCommander(options FakeSSHCommanderOptions) *FakeSSHCommander {
	if options.FilesystemType == "" {
		options.FilesystemType = "ext4"
	}
	sshCmder := &FakeSSHCommander{
		Responses: map[string]string{
			"stat -f -c %T /var/lib": options.FilesystemType + "\n",
		},
	}

	return sshCmder
}

//SSHCommand is an implementation of provision.SSHCommander.SSHCommand to provide predictable responses set by testing code
func (sshCmder *FakeSSHCommander) SSHCommand(args string) (string, error) {
	response, commandRegistered := sshCmder.Responses[args]
	if !commandRegistered {
		return "", errors.New("Command not registered in FakeSSHCommander")
	}
	return response, nil
}

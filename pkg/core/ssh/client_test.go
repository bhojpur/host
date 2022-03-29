package ssh

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
	"io/ioutil"
	"os"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSSHCmdArgs(t *testing.T) {
	cases := []struct {
		binaryPath   string
		args         []string
		expectedArgs []string
	}{
		{
			binaryPath: "/usr/local/bin/ssh",
			args: []string{
				"bhojpur@localhost",
				"apt-get install -y htop",
			},
			expectedArgs: []string{
				"/usr/local/bin/ssh",
				"bhojpur@localhost",
				"apt-get install -y htop",
			},
		},
		{
			binaryPath: "C:\\Program Files\\Git\\bin\\ssh.exe",
			args: []string{
				"bhojpur@localhost",
				"sudo /usr/bin/sethostname foobar && echo 'foobar' | sudo tee /var/lib/boot2docker/etc/hostname",
			},
			expectedArgs: []string{
				"C:\\Program Files\\Git\\bin\\ssh.exe",
				"bhojpur@localhost",
				"sudo /usr/bin/sethostname foobar && echo 'foobar' | sudo tee /var/lib/boot2docker/etc/hostname",
			},
		},
	}

	for _, c := range cases {
		cmd := getSSHCmd(c.binaryPath, c.args...)
		assert.Equal(t, cmd.Args, c.expectedArgs)
	}
}

func TestNewExternalClient(t *testing.T) {
	keyFile, err := ioutil.TempFile("", "bhojpur-machine-tests-dummy-private-key")
	if err != nil {
		t.Fatal(err)
	}
	defer keyFile.Close()

	keyFilename := keyFile.Name()
	defer os.Remove(keyFilename)

	cases := []struct {
		sshBinaryPath string
		user          string
		host          string
		port          int
		auth          *Auth
		perm          os.FileMode
		expectedError string
		skipOS        string
	}{
		{
			auth:          &Auth{Keys: []string{"/tmp/private-key-not-exist"}},
			expectedError: "stat /tmp/private-key-not-exist: no such file or directory",
			skipOS:        "none",
		},
		{
			auth:   &Auth{Keys: []string{keyFilename}},
			perm:   0400,
			skipOS: "windows",
		},
		{
			auth:          &Auth{Keys: []string{keyFilename}},
			perm:          0100,
			expectedError: fmt.Sprintf("'%s' is not readable", keyFilename),
			skipOS:        "windows",
		},
		{
			auth:          &Auth{Keys: []string{keyFilename}},
			perm:          0644,
			expectedError: fmt.Sprintf("permissions 0644 for '%s' are too open", keyFilename),
			skipOS:        "windows",
		},
	}

	for _, c := range cases {
		if runtime.GOOS != c.skipOS {
			keyFile.Chmod(c.perm)
			_, err := NewExternalClient(c.sshBinaryPath, c.user, c.host, c.port, c.auth)
			if c.expectedError != "" {
				assert.EqualError(t, err, c.expectedError)
			} else {
				assert.Equal(t, err, nil)
			}
		}
	}
}

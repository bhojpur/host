package localbinary

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
	"bufio"
	"fmt"
	"io"
	"testing"
	"time"

	"os"

	"github.com/bhojpur/host/pkg/machine/log"
	"github.com/stretchr/testify/assert"
)

type FakeExecutor struct {
	stdout, stderr io.ReadCloser
	closed         bool
}

func (fe *FakeExecutor) Start() (*bufio.Scanner, *bufio.Scanner, error) {
	return bufio.NewScanner(fe.stdout), bufio.NewScanner(fe.stderr), nil
}

func (fe *FakeExecutor) Close() error {
	fe.closed = true
	return nil
}

func TestLocalBinaryPluginAddress(t *testing.T) {
	lbp := &Plugin{}
	expectedAddr := "127.0.0.1:12345"

	lbp.addrCh = make(chan string, 1)
	lbp.addrCh <- expectedAddr

	// Call the first time to read from the channel
	addr, err := lbp.Address()
	if err != nil {
		t.Fatalf("Expected no error, instead got %s", err)
	}
	if addr != expectedAddr {
		t.Fatal("Expected did not match actual address")
	}

	// Call the second time to read the "cached" address value
	addr, err = lbp.Address()
	if err != nil {
		t.Fatalf("Expected no error, instead got %s", err)
	}
	if addr != expectedAddr {
		t.Fatal("Expected did not match actual address")
	}
}

func TestLocalBinaryPluginAddressTimeout(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping timeout test")
	}

	lbp := &Plugin{
		addrCh:  make(chan string, 1),
		timeout: 1 * time.Second,
	}

	addr, err := lbp.Address()

	assert.Empty(t, addr)
	assert.EqualError(t, err, "Failed to dial the plugin server in 1s")
}

func TestLocalBinaryPluginClose(t *testing.T) {
	lbp := &Plugin{}
	lbp.stopCh = make(chan bool, 1)
	go lbp.Close()
	stopped := <-lbp.stopCh
	if !stopped {
		t.Fatal("Close did not send a stop message on the proper channel")
	}
}

func TestExecServer(t *testing.T) {
	logOutReader, logOutWriter := io.Pipe()
	logErrReader, logErrWriter := io.Pipe()

	log.SetDebug(true)
	log.SetOutWriter(logOutWriter)
	log.SetErrWriter(logErrWriter)

	defer func() {
		log.SetDebug(false)
		log.SetOutWriter(os.Stdout)
		log.SetErrWriter(os.Stderr)
	}()

	stdoutReader, stdoutWriter := io.Pipe()
	stderrReader, stderrWriter := io.Pipe()

	fe := &FakeExecutor{
		stdout: stdoutReader,
		stderr: stderrReader,
	}

	machineName := "test"
	lbp := &Plugin{
		MachineName: machineName,
		Executor:    fe,
		addrCh:      make(chan string, 1),
		stopCh:      make(chan bool, 1),
	}

	finalErr := make(chan error)

	// Start the bhojpur-machine-foo plugin server
	go func() {
		finalErr <- lbp.execServer()
	}()

	logOutScanner := bufio.NewScanner(logOutReader)
	logErrScanner := bufio.NewScanner(logErrReader)

	// Write the ip address
	expectedAddr := "127.0.0.1:12345"
	if _, err := io.WriteString(stdoutWriter, expectedAddr+"\n"); err != nil {
		t.Fatalf("Error attempting to write plugin address: %s", err)
	}

	if addr := <-lbp.addrCh; addr != expectedAddr {
		t.Fatalf("Expected to read the expected address properly in server but did not")
	}

	// Write a log in stdout
	expectedPluginOut := "Doing some fun plugin stuff..."
	if _, err := io.WriteString(stdoutWriter, expectedPluginOut+"\n"); err != nil {
		t.Fatalf("Error attempting to write to out in plugin: %s", err)
	}

	expectedOut := fmt.Sprintf(pluginOut, machineName, expectedPluginOut)
	if logOutScanner.Scan(); logOutScanner.Text() != expectedOut {
		t.Fatalf("Output written to log was not what we expected\nexpected: %s\nactual:   %s", expectedOut, logOutScanner.Text())
	}

	// Write a log in stderr
	expectedPluginErr := "Uh oh, something in plugin went wrong..."
	if _, err := io.WriteString(stderrWriter, expectedPluginErr+"\n"); err != nil {
		t.Fatalf("Error attempting to write to err in plugin: %s", err)
	}

	expectedErr := fmt.Sprintf(pluginErr, machineName, expectedPluginErr)
	if logErrScanner.Scan(); logErrScanner.Text() != expectedErr {
		t.Fatalf("Error written to log was not what we expected\nexpected: %s\nactual:   %s", expectedErr, logErrScanner.Text())
	}

	lbp.Close()

	if err := <-finalErr; err != nil {
		t.Fatalf("Error serving: %s", err)
	}
}

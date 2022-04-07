package log

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
	"io"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func captureOutput(testLogger MachineLogger, lambda func()) string {
	pipeReader, pipeWriter := io.Pipe()
	scanner := bufio.NewScanner(pipeReader)
	testLogger.SetOutWriter(pipeWriter)
	go lambda()
	scanner.Scan()
	return scanner.Text()
}

func captureError(testLogger MachineLogger, lambda func()) string {
	pipeReader, pipeWriter := io.Pipe()
	scanner := bufio.NewScanner(pipeReader)
	testLogger.SetErrWriter(pipeWriter)
	go lambda()
	scanner.Scan()
	return scanner.Text()
}

func TestSetDebugToTrue(t *testing.T) {
	testLogger := NewFmtMachineLogger().(*FmtMachineLogger)
	testLogger.SetDebug(true)
	assert.Equal(t, true, testLogger.debug)
}

func TestSetDebugToFalse(t *testing.T) {
	testLogger := NewFmtMachineLogger().(*FmtMachineLogger)
	testLogger.SetDebug(true)
	testLogger.SetDebug(false)
	assert.Equal(t, false, testLogger.debug)
}

func TestSetOut(t *testing.T) {
	testLogger := NewFmtMachineLogger().(*FmtMachineLogger)
	testLogger.SetOutWriter(ioutil.Discard)
	assert.Equal(t, ioutil.Discard, testLogger.outWriter)
}

func TestSetErr(t *testing.T) {
	testLogger := NewFmtMachineLogger().(*FmtMachineLogger)
	testLogger.SetErrWriter(ioutil.Discard)
	assert.Equal(t, ioutil.Discard, testLogger.errWriter)
}

func TestDebug(t *testing.T) {
	testLogger := NewFmtMachineLogger()
	testLogger.SetDebug(true)

	result := captureError(testLogger, func() { testLogger.Debug("debug") })

	assert.Equal(t, result, "debug")
}

func TestInfo(t *testing.T) {
	testLogger := NewFmtMachineLogger()

	result := captureOutput(testLogger, func() { testLogger.Info("info") })

	assert.Equal(t, result, "info")
}

func TestWarn(t *testing.T) {
	testLogger := NewFmtMachineLogger()

	result := captureOutput(testLogger, func() { testLogger.Warn("warn") })

	assert.Equal(t, result, "warn")
}

func TestError(t *testing.T) {
	testLogger := NewFmtMachineLogger()

	result := captureError(testLogger, func() { testLogger.Error("error") })

	assert.Equal(t, result, "error")
}

func TestEntriesAreCollected(t *testing.T) {
	testLogger := NewFmtMachineLogger()
	testLogger.Debug("debug")
	testLogger.Info("info")
	testLogger.Error("error")
	assert.Equal(t, 3, len(testLogger.History()))
	assert.Equal(t, "debug", testLogger.History()[0])
	assert.Equal(t, "info", testLogger.History()[1])
	assert.Equal(t, "error", testLogger.History()[2])
}

//go:build darwin || linux
// +build darwin linux

package signal

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
	"os"
	"syscall"
	"testing"
	"time"
)

func TestCatchAll(t *testing.T) {
	sigs := make(chan os.Signal, 1)
	CatchAll(sigs)
	defer StopCatch(sigs)

	listOfSignals := map[string]string{
		"CONT": syscall.SIGCONT.String(),
		"HUP":  syscall.SIGHUP.String(),
		"CHLD": syscall.SIGCHLD.String(),
		"ILL":  syscall.SIGILL.String(),
		"FPE":  syscall.SIGFPE.String(),
		"CLD":  syscall.SIGCLD.String(),
	}

	for sigStr := range listOfSignals {
		if signal, ok := SignalMap[sigStr]; ok {
			_ = syscall.Kill(syscall.Getpid(), signal)
			s := <-sigs
			if s.String() != signal.String() {
				t.Errorf("expected: %q, got: %q", signal, s)
			}
		}
	}
}

func TestCatchAllIgnoreSigUrg(t *testing.T) {
	sigs := make(chan os.Signal, 1)
	CatchAll(sigs)
	defer StopCatch(sigs)

	err := syscall.Kill(syscall.Getpid(), syscall.SIGURG)
	if err != nil {
		t.Fatal(err)
	}
	timer := time.NewTimer(1 * time.Second)
	defer timer.Stop()
	select {
	case <-timer.C:
	case s := <-sigs:
		t.Fatalf("expected no signals to be handled, but received %q", s.String())
	}
}

func TestStopCatch(t *testing.T) {
	signal := SignalMap["HUP"]
	channel := make(chan os.Signal, 1)
	CatchAll(channel)
	_ = syscall.Kill(syscall.Getpid(), signal)
	signalString := <-channel
	if signalString.String() != signal.String() {
		t.Errorf("expected: %q, got: %q", signal, signalString)
	}

	StopCatch(channel)
	_, ok := <-channel
	if ok {
		t.Error("expected: !ok, got: ok")
	}
}

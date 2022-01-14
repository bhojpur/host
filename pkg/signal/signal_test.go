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
	"syscall"
	"testing"
)

func TestParseSignal(t *testing.T) {
	_, err := ParseSignal("0")
	expectedErr := "invalid signal: 0"
	if err == nil || err.Error() != expectedErr {
		t.Errorf("expected  %q, but got %v", expectedErr, err)
	}

	_, err = ParseSignal("SIG")
	expectedErr = "invalid signal: SIG"
	if err == nil || err.Error() != expectedErr {
		t.Errorf("expected  %q, but got %v", expectedErr, err)
	}

	for sigStr := range SignalMap {
		responseSignal, err := ParseSignal(sigStr)
		if err != nil {
			t.Error(err)
		}
		signal := SignalMap[sigStr]
		if responseSignal != signal {
			t.Errorf("expected: %q, got: %q", signal, responseSignal)
		}
	}
}

func TestValidSignalForPlatform(t *testing.T) {
	isValidSignal := ValidSignalForPlatform(syscall.Signal(0))
	if isValidSignal {
		t.Error("expected !isValidSignal")
	}

	for _, sigN := range SignalMap {
		isValidSignal = ValidSignalForPlatform(sigN)
		if !isValidSignal {
			t.Error("expected isValidSignal")
		}
	}
}

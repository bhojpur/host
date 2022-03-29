package virtualbox

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

	"errors"

	"github.com/stretchr/testify/assert"
)

type MockLogsReader struct {
	content []string
	err     error
}

func (r *MockLogsReader) Read(path string) ([]string, error) {
	return r.content, r.err
}

func TestIsVTXEnabledInTheVM(t *testing.T) {
	driver := NewDriver("default", "path")

	var tests = []struct {
		description string
		content     []string
		err         error
	}{
		{"Empty log", []string{}, nil},
		{"Raw mode", []string{"Falling back to raw-mode: VT-x is disabled in the BIOS for all CPU modes"}, nil},
		{"Raw mode", []string{"HM: HMR3Init: Falling back to raw-mode: VT-x is not available"}, nil},
	}

	for _, test := range tests {
		driver.logsReader = &MockLogsReader{
			content: test.content,
			err:     test.err,
		}

		disabled, err := driver.IsVTXDisabledInTheVM()

		assert.False(t, disabled, test.description)
		assert.Equal(t, test.err, err)
	}
}

func TestIsVTXDisabledInTheVM(t *testing.T) {
	driver := NewDriver("default", "path")

	var tests = []struct {
		description string
		content     []string
		err         error
	}{
		{"VT-x Disabled", []string{"VT-x is disabled"}, nil},
		{"No HW virtualization", []string{"the host CPU does NOT support HW virtualization"}, nil},
		{"Unable to start VM", []string{"VERR_VMX_UNABLE_TO_START_VM"}, nil},
		{"Power up failed", []string{"00:00:00.318604 Power up failed (vrc=VERR_VMX_NO_VMX, rc=NS_ERROR_FAILURE (0X80004005))"}, nil},
		{"Unable to read log", nil, errors.New("Unable to read log")},
	}

	for _, test := range tests {
		driver.logsReader = &MockLogsReader{
			content: test.content,
			err:     test.err,
		}

		disabled, err := driver.IsVTXDisabledInTheVM()

		assert.True(t, disabled, test.description)
		assert.Equal(t, test.err, err)
	}
}

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

import "strings"

// IsVTXDisabledInTheVM checks if VT-X is disabled in the started vm.
func (d *Driver) IsVTXDisabledInTheVM() (bool, error) {
	lines, err := d.readVBoxLog()
	if err != nil {
		return true, err
	}

	for _, line := range lines {
		if strings.Contains(line, "VT-x is disabled") && !strings.Contains(line, "Falling back to raw-mode: VT-x is disabled in the BIOS for all CPU modes") {
			return true, nil
		}
		if strings.Contains(line, "the host CPU does NOT support HW virtualization") {
			return true, nil
		}
		if strings.Contains(line, "VERR_VMX_UNABLE_TO_START_VM") {
			return true, nil
		}
		if strings.Contains(line, "Power up failed") && strings.Contains(line, "VERR_VMX_NO_VMX") {
			return true, nil
		}
	}

	return false, nil
}

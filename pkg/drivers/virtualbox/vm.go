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

import "strconv"

type VM struct {
	CPUs   int
	Memory int
}

func getVMInfo(name string, vbox VBoxManager) (*VM, error) {
	out, err := vbox.vbmOut("showvminfo", name, "--machinereadable")
	if err != nil {
		return nil, err
	}

	vm := &VM{}

	err = parseKeyValues(out, reEqualLine, func(key, val string) error {
		switch key {
		case "cpus":
			v, err := strconv.Atoi(val)
			if err != nil {
				return err
			}
			vm.CPUs = v
		case "memory":
			v, err := strconv.Atoi(val)
			if err != nil {
				return err
			}
			vm.Memory = v
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return vm, nil
}

package drivers

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

import mflag "github.com/bhojpur/host/pkg/machine/flag"

// CheckDriverOptions implements DriverOptions and is used to validate flag parsing
type CheckDriverOptions struct {
	FlagsValues  map[string]interface{}
	CreateFlags  []mflag.Flag
	InvalidFlags []string
}

func (o *CheckDriverOptions) String(key string) string {
	for _, flag := range o.CreateFlags {
		if flag.String() == key {
			f, ok := flag.(mflag.StringFlag)
			if !ok {
				o.InvalidFlags = append(o.InvalidFlags, flag.String())
			}

			value, present := o.FlagsValues[key].(string)
			if present {
				return value
			}
			return f.Value
		}
	}

	return ""
}

func (o *CheckDriverOptions) StringSlice(key string) []string {
	for _, flag := range o.CreateFlags {
		if flag.String() == key {
			f, ok := flag.(mflag.StringSliceFlag)
			if !ok {
				o.InvalidFlags = append(o.InvalidFlags, flag.String())
			}

			value, present := o.FlagsValues[key].([]string)
			if present {
				return value
			}
			return f.Value
		}
	}

	return nil
}

func (o *CheckDriverOptions) Int(key string) int {
	for _, flag := range o.CreateFlags {
		if flag.String() == key {
			f, ok := flag.(mflag.IntFlag)
			if !ok {
				o.InvalidFlags = append(o.InvalidFlags, flag.String())
			}

			value, present := o.FlagsValues[key].(int)
			if present {
				return value
			}
			return f.Value
		}
	}

	return 0
}

func (o *CheckDriverOptions) Bool(key string) bool {
	for _, flag := range o.CreateFlags {
		if flag.String() == key {
			_, ok := flag.(mflag.BoolFlag)
			if !ok {
				o.InvalidFlags = append(o.InvalidFlags, flag.String())
			}
		}
	}

	value, present := o.FlagsValues[key].(bool)
	if present {
		return value
	}
	return false
}

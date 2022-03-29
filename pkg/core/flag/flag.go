package mcnflag

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

import "fmt"

type Flag interface {
	fmt.Stringer
	Default() interface{}
}

type StringFlag struct {
	Name   string
	Usage  string
	EnvVar string
	Value  string
}

// TODO: Could this be done more succinctly using embedding?
func (f StringFlag) String() string {
	return f.Name
}

func (f StringFlag) Default() interface{} {
	return f.Value
}

type StringSliceFlag struct {
	Name   string
	Usage  string
	EnvVar string
	Value  []string
}

// TODO: Could this be done more succinctly using embedding?
func (f StringSliceFlag) String() string {
	return f.Name
}

func (f StringSliceFlag) Default() interface{} {
	return f.Value
}

type IntFlag struct {
	Name   string
	Usage  string
	EnvVar string
	Value  int
}

// TODO: Could this be done more succinctly using embedding?
func (f IntFlag) String() string {
	return f.Name
}

func (f IntFlag) Default() interface{} {
	return f.Value
}

type BoolFlag struct {
	Name   string
	Usage  string
	EnvVar string
}

// TODO: Could this be done more succinctly using embedding?
func (f BoolFlag) String() string {
	return f.Name
}

func (f BoolFlag) Default() interface{} {
	return false
}

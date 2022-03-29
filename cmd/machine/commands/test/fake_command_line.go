package test

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
	"github.com/urfave/cli"
)

type FakeFlagger struct {
	Data map[string]interface{}
}

type FakeCommandLine struct {
	LocalFlags, GlobalFlags *FakeFlagger
	HelpShown, VersionShown bool
	CliArgs                 []string
}

func (ff FakeFlagger) String(key string) string {
	if value, ok := ff.Data[key]; ok {
		return value.(string)
	}
	return ""
}

func (ff FakeFlagger) StringSlice(key string) []string {
	if value, ok := ff.Data[key]; ok {
		return value.([]string)
	}
	return []string{}
}

func (ff FakeFlagger) Int(key string) int {
	if value, ok := ff.Data[key]; ok {
		return value.(int)
	}
	return 0
}

func (ff FakeFlagger) Bool(key string) bool {
	if value, ok := ff.Data[key]; ok {
		return value.(bool)
	}
	return false
}

func (fcli *FakeCommandLine) IsSet(key string) bool {
	_, ok := fcli.LocalFlags.Data[key]
	return ok
}

func (fcli *FakeCommandLine) String(key string) string {
	return fcli.LocalFlags.String(key)
}

func (fcli *FakeCommandLine) StringSlice(key string) []string {
	return fcli.LocalFlags.StringSlice(key)
}

func (fcli *FakeCommandLine) Int(key string) int {
	return fcli.LocalFlags.Int(key)
}

func (fcli *FakeCommandLine) Bool(key string) bool {
	if fcli.LocalFlags == nil {
		return false
	}
	return fcli.LocalFlags.Bool(key)
}

func (fcli *FakeCommandLine) GlobalString(key string) string {
	return fcli.GlobalFlags.String(key)
}

func (fcli *FakeCommandLine) Generic(name string) interface{} {
	return fcli.LocalFlags.Data[name]
}

func (fcli *FakeCommandLine) FlagNames() []string {
	flagNames := []string{}
	for key := range fcli.LocalFlags.Data {
		flagNames = append(flagNames, key)
	}

	return flagNames
}

func (fcli *FakeCommandLine) ShowHelp() {
	fcli.HelpShown = true
}

func (fcli *FakeCommandLine) Application() *cli.App {
	return cli.NewApp()
}

func (fcli *FakeCommandLine) Args() cli.Args {
	return fcli.CliArgs
}

func (fcli *FakeCommandLine) ShowVersion() {
	fcli.VersionShown = true
}

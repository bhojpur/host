package commands

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

	"flag"

	ctest "github.com/bhojpur/host/cmd/machine/commands/test"
	mflag "github.com/bhojpur/host/pkg/core/flag"
	"github.com/stretchr/testify/assert"
)

func TestValidateSwarmDiscoveryErrorsGivenInvalidURL(t *testing.T) {
	err := validateSwarmDiscovery("foo")
	assert.Error(t, err)
}

func TestValidateSwarmDiscoveryAcceptsEmptyString(t *testing.T) {
	err := validateSwarmDiscovery("")
	assert.NoError(t, err)
}

func TestValidateSwarmDiscoveryAcceptsValidFormat(t *testing.T) {
	err := validateSwarmDiscovery("token://deadbeefcafe")
	assert.NoError(t, err)
}

type fakeFlagGetter struct {
	flag.Value
	value interface{}
}

func (ff fakeFlagGetter) Get() interface{} {
	return ff.value
}

var nilStringSlice []string

var getDriverOptsFlags = []mflag.Flag{
	mflag.BoolFlag{
		Name: "bool",
	},
	mflag.IntFlag{
		Name: "int",
	},
	mflag.IntFlag{
		Name:  "int_defaulted",
		Value: 42,
	},
	mflag.StringFlag{
		Name: "string",
	},
	mflag.StringFlag{
		Name:  "string_defaulted",
		Value: "bob",
	},
	mflag.StringSliceFlag{
		Name: "stringslice",
	},
	mflag.StringSliceFlag{
		Name:  "stringslice_defaulted",
		Value: []string{"joe"},
	},
}

var getDriverOptsTests = []struct {
	data     map[string]interface{}
	expected map[string]interface{}
}{
	{
		expected: map[string]interface{}{
			"bool":                  false,
			"int":                   0,
			"int_defaulted":         42,
			"string":                "",
			"string_defaulted":      "bob",
			"stringslice":           nilStringSlice,
			"stringslice_defaulted": []string{"joe"},
		},
	},
	{
		data: map[string]interface{}{
			"bool":             fakeFlagGetter{value: true},
			"int":              fakeFlagGetter{value: 42},
			"int_defaulted":    fakeFlagGetter{value: 37},
			"string":           fakeFlagGetter{value: "jake"},
			"string_defaulted": fakeFlagGetter{value: "george"},
			// NB: StringSlices are not flag.Getters.
			"stringslice":           []string{"ford"},
			"stringslice_defaulted": []string{"zaphod", "arthur"},
		},
		expected: map[string]interface{}{
			"bool":                  true,
			"int":                   42,
			"int_defaulted":         37,
			"string":                "jake",
			"string_defaulted":      "george",
			"stringslice":           []string{"ford"},
			"stringslice_defaulted": []string{"zaphod", "arthur"},
		},
	},
}

func TestGetDriverOpts(t *testing.T) {
	for _, tt := range getDriverOptsTests {
		commandLine := &ctest.FakeCommandLine{
			LocalFlags: &ctest.FakeFlagger{
				Data: tt.data,
			},
		}
		driverOpts := getDriverOpts(commandLine, getDriverOptsFlags)
		assert.Equal(t, tt.expected["bool"], driverOpts.Bool("bool"))
		assert.Equal(t, tt.expected["int"], driverOpts.Int("int"))
		assert.Equal(t, tt.expected["int_defaulted"], driverOpts.Int("int_defaulted"))
		assert.Equal(t, tt.expected["string"], driverOpts.String("string"))
		assert.Equal(t, tt.expected["string_defaulted"], driverOpts.String("string_defaulted"))
		assert.Equal(t, tt.expected["stringslice"], driverOpts.StringSlice("stringslice"))
		assert.Equal(t, tt.expected["stringslice_defaulted"], driverOpts.StringSlice("stringslice_defaulted"))
	}
}

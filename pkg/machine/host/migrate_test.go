package host

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

	"github.com/bhojpur/host/pkg/drivers/none"
	"github.com/bhojpur/host/pkg/machine/auth"
	"github.com/stretchr/testify/assert"
)

func TestMigrateHost(t *testing.T) {
	testCases := []struct {
		description                string
		hostBefore                 *Host
		rawData                    []byte
		expectedHostAfter          *Host
		expectedMigrationPerformed bool
		expectedMigrationError     error
	}{
		{
			// Point of this test is largely that no matter what was in RawDriver
			// before, it should load into the Host struct based on what is actually
			// in the Driver field.
			//
			// Note that we don't check for the presence of RawDriver's literal "on
			// disk" here.  It's intentional.
			description: "Config version 3 load with existing RawDriver on disk",
			hostBefore: &Host{
				Name: "default",
			},
			rawData: []byte(`{
    "ConfigVersion": 3,
    "Driver": {"MachineName": "default"},
    "DriverName": "virtualbox",
    "HostOptions": {
        "Driver": "",
        "Memory": 0,
        "Disk": 0,
        "AuthOptions": {
            "StorePath": "/Users/shashi.rai/.bhojpur/machine/machines/default"
        }
    },
    "Name": "default",
    "RawDriver": "eyJWQm94TWFuYWdlciI6e30sIklQQWRkcmVzcyI6IjE5Mi4xNjguOTkuMTAwIiwiTWFjaGluZU5hbWUiOiJkZWZhdWx0IiwiU1NIVXNlciI6ImRvY2tlciIsIlNTSFBvcnQiOjU4MTQ1LCJTU0hLZXlQYXRoIjoiL1VzZXJzL25hdGhhbmxlY2xhaXJlLy5kb2NrZXIvbWFjaGluZS9tYWNoaW5lcy9kZWZhdWx0L2lkX3JzYSIsIlN0b3JlUGF0aCI6Ii9Vc2Vycy9uYXRoYW5sZWNsYWlyZS8uZG9ja2VyL21hY2hpbmUiLCJTd2FybU1hc3RlciI6ZmFsc2UsIlN3YXJtSG9zdCI6InRjcDovLzAuMC4wLjA6MzM3NiIsIlN3YXJtRGlzY292ZXJ5IjoiIiwiQ1BVIjoxLCJNZW1vcnkiOjEwMjQsIkRpc2tTaXplIjoyMDAwMCwiQm9vdDJEb2NrZXJVUkwiOiIiLCJCb290MkRvY2tlckltcG9ydFZNIjoiIiwiSG9zdE9ubHlDSURSIjoiMTkyLjE2OC45OS4xLzI0IiwiSG9zdE9ubHlOaWNUeXBlIjoiODI1NDBFTSIsIkhvc3RPbmx5UHJvbWlzY01vZGUiOiJkZW55IiwiTm9TaGFyZSI6ZmFsc2V9"
}`),
			expectedHostAfter: &Host{
				ConfigVersion: 3,
				HostOptions: &Options{
					AuthOptions: &auth.Options{
						StorePath: "/Users/shashi.rai/.bhojpur/machine/machines/default",
					},
				},
				Name:       "default",
				DriverName: "virtualbox",
				RawDriver:  []byte(`{"MachineName": "default"}`),
				Driver: &RawDataDriver{
					Data: []byte(`{"MachineName": "default"}`),

					// TODO: The "." argument here is a already existing
					// bug (or at least likely to cause them in the future) and most
					// likely should be "/Users/shashi.rai/.bhojpur/machine"
					//
					// These default StorePath settings get over-written when we
					// instantiate the plugin driver, but this seems entirely incidental.
					Driver: none.NewDriver("default", "."),
				},
			},
			expectedMigrationPerformed: false,
			expectedMigrationError:     nil,
		},
		{
			description: "Config version 4 (from the FUTURE) on disk",
			hostBefore: &Host{
				Name: "default",
			},
			rawData: []byte(`{
    "ConfigVersion": 4,
    "Driver": {"MachineName": "default"},
    "DriverName": "virtualbox",
    "HostOptions": {
        "Driver": "",
        "Memory": 0,
        "Disk": 0,
        "AuthOptions": {
            "StorePath": "/Users/shashi.rai/.bhojpur/machine/machines/default"
        }
    },
    "Name": "default"
}`),
			expectedHostAfter:          nil,
			expectedMigrationPerformed: false,
			expectedMigrationError:     errConfigFromFuture,
		},
		{
			description: "Config version 3 load WITHOUT any existing RawDriver field on disk",
			hostBefore: &Host{
				Name: "default",
			},
			rawData: []byte(`{
    "ConfigVersion": 3,
    "Driver": {"MachineName": "default"},
    "DriverName": "virtualbox",
    "HostOptions": {
        "Driver": "",
        "Memory": 0,
        "Disk": 0,
        "AuthOptions": {
            "StorePath": "/Users/shashi.rai/.bhojpur/machine/machines/default"
        }
    },
    "Name": "default"
}`),
			expectedHostAfter: &Host{
				ConfigVersion: 3,
				HostOptions: &Options{
					AuthOptions: &auth.Options{
						StorePath: "/Users/shashi.rai/.bhojpur/machine/machines/default",
					},
				},
				Name:       "default",
				DriverName: "virtualbox",
				RawDriver:  []byte(`{"MachineName": "default"}`),
				Driver: &RawDataDriver{
					Data: []byte(`{"MachineName": "default"}`),

					// TODO: See note above.
					Driver: none.NewDriver("default", "."),
				},
			},
			expectedMigrationPerformed: false,
			expectedMigrationError:     nil,
		},
		{
			description: "Config version 2 load and migrate.  Ensure StorePath gets set properly.",
			hostBefore: &Host{
				Name: "default",
			},
			rawData: []byte(`{
    "ConfigVersion": 2,
    "Driver": {"MachineName": "default"},
    "DriverName": "virtualbox",
    "HostOptions": {
        "Driver": "",
        "Memory": 0,
        "Disk": 0,
        "AuthOptions": {
            "StorePath": "/Users/shashi.rai/.bhojpur/machine/machines/default"
        }
    },
    "StorePath": "/Users/shashi.rai/.bhojpur/machine/machines/default",
    "Name": "default"
}`),
			expectedHostAfter: &Host{
				ConfigVersion: 3,
				HostOptions: &Options{
					AuthOptions: &auth.Options{
						StorePath: "/Users/shashi.rai/.bhojpur/machine/machines/default",
					},
				},
				Name:       "default",
				DriverName: "virtualbox",
				RawDriver:  []byte(`{"MachineName":"default","StorePath":"/Users/shashi.rai/.bhojpur/machine"}`),
				Driver: &RawDataDriver{
					Data:   []byte(`{"MachineName":"default","StorePath":"/Users/shashi.rai/.bhojpur/machine"}`),
					Driver: none.NewDriver("default", "/Users/shashi.rai/.bhojpur/machine"),
				},
			},
			expectedMigrationPerformed: true,
			expectedMigrationError:     nil,
		},
	}

	for _, tc := range testCases {
		actualHostAfter, actualMigrationPerformed, actualMigrationError := MigrateHost(tc.hostBefore, tc.rawData)

		assert.Equal(t, tc.expectedHostAfter, actualHostAfter)
		assert.Equal(t, tc.expectedMigrationPerformed, actualMigrationPerformed)
		assert.Equal(t, tc.expectedMigrationError, actualMigrationError)
	}
}

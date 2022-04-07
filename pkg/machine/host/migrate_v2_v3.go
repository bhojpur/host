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
	"bytes"
	"encoding/json"

	"github.com/bhojpur/host/pkg/machine/log"
)

type RawHost struct {
	Driver *json.RawMessage
}

func MigrateHostV2ToHostV3(hostV2 *V2, data []byte, storePath string) *Host {
	// Migrate to include RawDriver so that driver plugin will work
	// smoothly.
	rawHost := &RawHost{}
	if err := json.Unmarshal(data, &rawHost); err != nil {
		log.Warnf("Could not unmarshal raw host for RawDriver information: %s", err)
	}

	m := make(map[string]interface{})

	// Must migrate to include store path in driver since it was not
	// previously stored in drivers directly
	d := json.NewDecoder(bytes.NewReader(*rawHost.Driver))
	d.UseNumber()
	if err := d.Decode(&m); err != nil {
		log.Warnf("Could not unmarshal raw host into map[string]interface{}: %s", err)
	}

	m["StorePath"] = storePath

	// Now back to []byte
	rawDriver, err := json.Marshal(m)
	if err != nil {
		log.Warnf("Could not re-marshal raw driver: %s", err)
	}

	h := &Host{
		ConfigVersion: 2,
		DriverName:    hostV2.DriverName,
		Name:          hostV2.Name,
		HostOptions:   hostV2.HostOptions,
		RawDriver:     rawDriver,
	}

	return h
}

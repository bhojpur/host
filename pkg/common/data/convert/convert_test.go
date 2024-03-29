package convert

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
)

type data struct {
	TTLMillis int `json:"ttl"`
}

func TestJSON(t *testing.T) {
	d := &data{
		TTLMillis: 57600000,
	}

	m, err := EncodeToMap(d)
	if err != nil {
		t.Fatal(err)
	}

	i, _ := ToNumber(m["ttl"])
	if i != 57600000 {
		t.Fatal("not", 57600000, "got", m["ttl"])
	}
}

func TestArgKey(t *testing.T) {
	data := []struct {
		input  string
		output string
	}{
		{
			input:  "disableOpenAPIValidation",
			output: "--disable-open-api-validation",
		},
		{
			input:  "skipCRDs",
			output: "--skip-crds",
		},
	}

	for _, data := range data {
		if ToArgKey(data.input) != data.output {
			t.Errorf("expected %s, got %s", data.output, ToArgKey(data.input))
		}
	}
}

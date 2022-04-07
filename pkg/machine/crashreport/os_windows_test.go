package crashreport

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

	"github.com/stretchr/testify/assert"
)

func TestParseVerOutput(t *testing.T) {
	output := `

Microsoft Windows [version 6.3.9600]

`

	assert.Equal(t, "Microsoft Windows [version 6.3.9600]", parseVerOutput(output))
}

func TestParseSystemInfoOutput(t *testing.T) {
	output := `
Host Name:                 DESKTOP-3A5PULA
OS Name:                   Microsoft Windows 10 Enterprise
OS Version:                10.0.10240 N/A Build 10240
OS Manufacturer:           Microsoft Corporation
OS Configuration:          Standalone Workstation
OS Build Type:             Multiprocessor Free
Registered Owner:          Windows User
`

	assert.Equal(t, "10.0.10240 N/A Build 10240", parseSystemInfoOutput(output))
}

func TestParseNonEnglishSystemInfoOutput(t *testing.T) {
	output := `
Ignored:                 ...
Ignored:                 ...
Version du Système:      10.0.10350
`

	assert.Equal(t, "10.0.10350", parseSystemInfoOutput(output))
}

func TestParseInvalidSystemInfoOutput(t *testing.T) {
	output := "Invalid"

	assert.Empty(t, parseSystemInfoOutput(output))
}

package memory

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
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"

	units "github.com/bhojpur/units/pkg/uom"
)

// ReadMemInfo retrieves memory statistics of the host system and returns a
// MemInfo type.
func ReadMemInfo() (*MemInfo, error) {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return parseMemInfo(file)
}

// parseMemInfo parses the /proc/meminfo file into
// a MemInfo object given an io.Reader to the file.
// Throws error if there are problems reading from the file
func parseMemInfo(reader io.Reader) (*MemInfo, error) {
	meminfo := &MemInfo{}
	scanner := bufio.NewScanner(reader)
	memAvailable := int64(-1)
	for scanner.Scan() {
		// Expected format: ["MemTotal:", "1234", "kB"]
		parts := strings.Fields(scanner.Text())

		// Sanity checks: Skip malformed entries.
		if len(parts) < 3 || parts[2] != "kB" {
			continue
		}

		// Convert to bytes.
		size, err := strconv.Atoi(parts[1])
		if err != nil {
			continue
		}
		bytes := int64(size) * units.KiB

		switch parts[0] {
		case "MemTotal:":
			meminfo.MemTotal = bytes
		case "MemFree:":
			meminfo.MemFree = bytes
		case "MemAvailable:":
			memAvailable = bytes
		case "SwapTotal:":
			meminfo.SwapTotal = bytes
		case "SwapFree:":
			meminfo.SwapFree = bytes
		}

	}
	if memAvailable != -1 {
		meminfo.MemFree = memAvailable
	}

	// Handle errors that may have occurred during the reading of the file.
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return meminfo, nil
}

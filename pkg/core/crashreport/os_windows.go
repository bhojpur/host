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
	"os/exec"
	"strings"
)

func localOSVersion() string {
	command := exec.Command("ver")
	output, err := command.Output()
	if err == nil {
		return parseVerOutput(string(output))
	}

	command = exec.Command("systeminfo")
	output, err = command.Output()
	if err == nil {
		return parseSystemInfoOutput(string(output))
	}

	return ""
}

func parseSystemInfoOutput(output string) string {
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "OS Version:") {
			return strings.TrimSpace(line[len("OS Version:"):])
		}
	}

	// If we couldn't find the version, maybe the output is not in English
	// Let's parse the fourth line since it seems to be the one always used
	// for the version.
	if len(lines) >= 4 {
		parts := strings.Split(lines[3], ":")
		if len(parts) == 2 {
			return strings.TrimSpace(parts[1])
		}
	}

	return ""
}

func parseVerOutput(output string) string {
	return strings.TrimSpace(output)
}

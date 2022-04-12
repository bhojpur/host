package kv

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

import "strings"

// Like split but if there is only one item return "", item
func RSplit(s, sep string) (string, string) {
	parts := strings.SplitN(s, sep, 2)
	if len(parts) == 1 {
		return "", strings.TrimSpace(parts[0])
	}
	return strings.TrimSpace(parts[0]), strings.TrimSpace(safeIndex(parts, 1))
}

func Split(s, sep string) (string, string) {
	parts := strings.SplitN(s, sep, 2)
	return strings.TrimSpace(parts[0]), strings.TrimSpace(safeIndex(parts, 1))
}

func SplitLast(s, sep string) (string, string) {
	idx := strings.LastIndex(s, sep)
	if idx > -1 {
		return strings.TrimSpace(s[:idx]), strings.TrimSpace(s[idx+1:])
	}
	return s, ""
}

func SplitMap(s, sep string) map[string]string {
	return SplitMapFromSlice(strings.Split(s, sep))
}

func SplitMapFromSlice(parts []string) map[string]string {
	result := map[string]string{}
	for _, part := range parts {
		k, v := Split(part, "=")
		result[k] = v
	}
	return result
}

func safeIndex(parts []string, idx int) string {
	if len(parts) <= idx {
		return ""
	}
	return parts[idx]
}

//go:build !windows
// +build !windows

package statistics

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
	"os"
	"syscall"
)

// StatT type contains status of a file. It contains metadata
// like permission, owner, group, size, etc about a file.
type StatT struct {
	mode uint32
	uid  uint32
	gid  uint32
	rdev uint64
	size int64
	mtim syscall.Timespec
}

// Mode returns file's permission mode.
func (s StatT) Mode() uint32 {
	return s.mode
}

// UID returns file's user id of owner.
func (s StatT) UID() uint32 {
	return s.uid
}

// GID returns file's group id of owner.
func (s StatT) GID() uint32 {
	return s.gid
}

// Rdev returns file's device ID (if it's special file).
func (s StatT) Rdev() uint64 {
	return s.rdev
}

// Size returns file's size.
func (s StatT) Size() int64 {
	return s.size
}

// Mtim returns file's last modification time.
func (s StatT) Mtim() syscall.Timespec {
	return s.mtim
}

// IsDir reports whether s describes a directory.
func (s StatT) IsDir() bool {
	return s.mode&syscall.S_IFDIR != 0
}

// Stat takes a path to a file and returns
// a system.StatT type pertaining to that file.
//
// Throws an error if the file does not exist
func Stat(path string) (*StatT, error) {
	s := &syscall.Stat_t{}
	if err := syscall.Stat(path, s); err != nil {
		return nil, &os.PathError{Op: "Stat", Path: path, Err: err}
	}
	return fromStatT(s)
}

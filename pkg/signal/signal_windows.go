package signal

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
	"syscall"

	"golang.org/x/sys/windows"
)

// Signals used in cli/command (no windows equivalent, use
// invalid signals so they don't get handled)
const (
	SIGCHLD  = syscall.Signal(0xff)
	SIGWINCH = syscall.Signal(0xff)
	SIGPIPE  = syscall.Signal(0xff)
)

// SignalMap is a map of "supported" signals. As per the comment in GOLang's
// ztypes_windows.go: "More invented values for signals". Windows doesn't
// really support signals in any way, shape or form that Unix does.
var SignalMap = map[string]syscall.Signal{
	"ABRT": syscall.Signal(windows.SIGABRT),
	"ALRM": syscall.Signal(windows.SIGALRM),
	"BUS":  syscall.Signal(windows.SIGBUS),
	"FPE":  syscall.Signal(windows.SIGFPE),
	"HUP":  syscall.Signal(windows.SIGHUP),
	"ILL":  syscall.Signal(windows.SIGILL),
	"INT":  syscall.Signal(windows.SIGINT),
	"KILL": syscall.Signal(windows.SIGKILL),
	"PIPE": syscall.Signal(windows.SIGPIPE),
	"QUIT": syscall.Signal(windows.SIGQUIT),
	"SEGV": syscall.Signal(windows.SIGSEGV),
	"TERM": syscall.Signal(windows.SIGTERM),
	"TRAP": syscall.Signal(windows.SIGTRAP),
}

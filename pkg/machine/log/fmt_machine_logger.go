package log

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
	"fmt"
	"io"
	"os"
)

type FmtMachineLogger struct {
	outWriter io.Writer
	errWriter io.Writer
	debug     bool
	history   *HistoryRecorder
}

// NewFmtMachineLogger creates a MachineLogger implementation used by the drivers
func NewFmtMachineLogger() MachineLogger {
	return &FmtMachineLogger{
		outWriter: os.Stdout,
		errWriter: os.Stderr,
		debug:     false,
		history:   NewHistoryRecorder(),
	}
}

func (ml *FmtMachineLogger) SetDebug(debug bool) {
	ml.debug = debug
}

func (ml *FmtMachineLogger) SetOutWriter(out io.Writer) {
	ml.outWriter = out
}

func (ml *FmtMachineLogger) SetErrWriter(err io.Writer) {
	ml.errWriter = err
}

func (ml *FmtMachineLogger) Debug(args ...interface{}) {
	ml.history.Record(args...)
	if ml.debug {
		fmt.Fprintln(ml.outWriter, args...)
	}
}

func (ml *FmtMachineLogger) Debugf(fmtString string, args ...interface{}) {
	ml.history.Recordf(fmtString, args...)
	if ml.debug {
		fmt.Fprintf(ml.outWriter, fmtString+"\n", args...)
	}
}

func (ml *FmtMachineLogger) Error(args ...interface{}) {
	ml.history.Record(args...)
	fmt.Fprintln(ml.errWriter, args...)
}

func (ml *FmtMachineLogger) Errorf(fmtString string, args ...interface{}) {
	ml.history.Recordf(fmtString, args...)
	fmt.Fprintf(ml.errWriter, fmtString+"\n", args...)
}

func (ml *FmtMachineLogger) Info(args ...interface{}) {
	ml.history.Record(args...)
	fmt.Fprintln(ml.outWriter, args...)
}

func (ml *FmtMachineLogger) Infof(fmtString string, args ...interface{}) {
	ml.history.Recordf(fmtString, args...)
	fmt.Fprintf(ml.outWriter, fmtString+"\n", args...)
}

func (ml *FmtMachineLogger) Warn(args ...interface{}) {
	ml.history.Record(args...)
	fmt.Fprintln(ml.outWriter, args...)
}

func (ml *FmtMachineLogger) Warnf(fmtString string, args ...interface{}) {
	ml.history.Recordf(fmtString, args...)
	fmt.Fprintf(ml.outWriter, fmtString+"\n", args...)
}

func (ml *FmtMachineLogger) History() []string {
	return ml.history.records
}

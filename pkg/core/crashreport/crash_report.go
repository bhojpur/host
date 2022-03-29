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
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/bhojpur/host/pkg/core/log"
	"github.com/bhojpur/host/pkg/core/shell"
	"github.com/bhojpur/host/pkg/version"
	"github.com/bugsnag/bugsnag-go"
)

const (
	defaultAPIKey  = "a9697f9a010c33ee218a65e5b1f3b0c1"
	noreportAPIKey = "no-report"
)

type CrashReporter interface {
	Send(err CrashError) error
}

// CrashError describes an error that should be reported to bugsnag
type CrashError struct {
	Cause       error
	Command     string
	Context     string
	DriverName  string
	LogFilePath string
}

func (e CrashError) Error() string {
	return e.Cause.Error()
}

type BugsnagCrashReporter struct {
	baseDir string
	apiKey  string
}

// NewCrashReporter creates a new bugsnag based CrashReporter. Needs an apiKey.
var NewCrashReporter = func(baseDir string, apiKey string) CrashReporter {
	if apiKey == "" {
		apiKey = defaultAPIKey
	}

	return &BugsnagCrashReporter{
		baseDir: baseDir,
		apiKey:  apiKey,
	}
}

// Send sends a crash report to bugsnag via an http call.
func (r *BugsnagCrashReporter) Send(err CrashError) error {
	if r.noReportFileExist() || r.apiKey == noreportAPIKey {
		log.Debug("Opting out of crash reporting.")
		return nil
	}

	if r.apiKey == "" {
		return errors.New("no api key has been set")
	}

	bugsnag.Configure(bugsnag.Configuration{
		APIKey: r.apiKey,
		// XXX we need to abuse bugsnag metrics to get the OS/ARCH information as a usable filter
		// Can do that with either "stage" or "hostname"
		ReleaseStage:    fmt.Sprintf("%s (%s)", runtime.GOOS, runtime.GOARCH),
		ProjectPackages: []string{"github.com/bhojpur/host/[^v]*"},
		AppVersion:      version.FullVersion(),
		Synchronous:     true,
		PanicHandler:    func() {},
		Logger:          new(logger),
	})

	metaData := bugsnag.MetaData{}

	metaData.Add("app", "compiler", fmt.Sprintf("%s (%s)", runtime.Compiler, runtime.Version()))
	metaData.Add("device", "os", runtime.GOOS)
	metaData.Add("device", "arch", runtime.GOARCH)

	detectRunningShell(&metaData)
	detectUname(&metaData)
	detectOSVersion(&metaData)
	addFile(err.LogFilePath, &metaData)

	var buffer bytes.Buffer
	for _, message := range log.History() {
		buffer.WriteString(message + "\n")
	}
	metaData.Add("history", "trace", buffer.String())

	return bugsnag.Notify(err.Cause, metaData, bugsnag.SeverityError, bugsnag.Context{String: err.Context}, bugsnag.ErrorClass{Name: fmt.Sprintf("%s/%s", err.DriverName, err.Command)})
}

func (r *BugsnagCrashReporter) noReportFileExist() bool {
	optOutFilePath := filepath.Join(r.baseDir, "no-error-report")
	if _, err := os.Stat(optOutFilePath); os.IsNotExist(err) {
		return false
	}
	return true
}

func addFile(path string, metaData *bugsnag.MetaData) {
	if path == "" {
		return
	}
	file, err := os.Open(path)
	if err != nil {
		log.Debug(err)
		return
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Debug(err)
		return
	}
	metaData.Add("logfile", filepath.Base(path), string(data))
}

func detectRunningShell(metaData *bugsnag.MetaData) {
	shell, err := shell.Detect()
	if err == nil {
		metaData.Add("device", "shell", shell)
	}
}

func detectUname(metaData *bugsnag.MetaData) {
	cmd := exec.Command("uname", "-s")
	output, err := cmd.Output()
	if err != nil {
		return
	}
	metaData.Add("device", "uname", string(output))
}

func detectOSVersion(metaData *bugsnag.MetaData) {
	metaData.Add("device", "os version", localOSVersion())
}

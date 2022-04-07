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
	"io"
	"regexp"
)

const redactedText = "<REDACTED>"

var (
	logger = NewFmtMachineLogger()

	// (?s) enables '.' to match '\n' -- see https://golang.org/pkg/regexp/syntax/
	certRegex = regexp.MustCompile("(?s)-----BEGIN CERTIFICATE-----.*-----END CERTIFICATE-----")
	keyRegex  = regexp.MustCompile("(?s)-----BEGIN RSA PRIVATE KEY-----.*-----END RSA PRIVATE KEY-----")
)

func stripSecrets(original []string) []string {
	stripped := []string{}
	for _, line := range original {
		line = certRegex.ReplaceAllString(line, redactedText)
		line = keyRegex.ReplaceAllString(line, redactedText)
		stripped = append(stripped, line)
	}
	return stripped
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Debugf(fmtString string, args ...interface{}) {
	logger.Debugf(fmtString, args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Errorf(fmtString string, args ...interface{}) {
	logger.Errorf(fmtString, args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Infof(fmtString string, args ...interface{}) {
	logger.Infof(fmtString, args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Warnf(fmtString string, args ...interface{}) {
	logger.Warnf(fmtString, args...)
}

func SetDebug(debug bool) {
	logger.SetDebug(debug)
}

func SetOutWriter(out io.Writer) {
	logger.SetOutWriter(out)
}

func SetErrWriter(err io.Writer) {
	logger.SetErrWriter(err)
}

func History() []string {
	return stripSecrets(logger.History())
}

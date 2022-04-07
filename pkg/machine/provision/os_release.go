package provision

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
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/bhojpur/host/pkg/machine/log"
)

// The /etc/os-release file contains operating system identification data
// See http://www.freedesktop.org/software/systemd/man/os-release.html for more details

// OsRelease reflects values in /etc/os-release
// Values in this struct must always be string
// or the reflection will not work properly.
type OsRelease struct {
	AnsiColor    string `osr:"ANSI_COLOR"`
	Name         string `osr:"NAME"`
	Version      string `osr:"VERSION"`
	Variant      string `osr:"VARIANT"`
	VariantID    string `osr:"VARIANT_ID"`
	ID           string `osr:"ID"`
	IDLike       string `osr:"ID_LIKE"`
	PrettyName   string `osr:"PRETTY_NAME"`
	VersionID    string `osr:"VERSION_ID"`
	HomeURL      string `osr:"HOME_URL"`
	SupportURL   string `osr:"SUPPORT_URL"`
	BugReportURL string `osr:"BUG_REPORT_URL"`
}

func stripQuotes(val string) string {
	if len(val) > 0 && val[0] == '"' {
		return val[1 : len(val)-1]
	}
	return val
}

func (osr *OsRelease) setIfPossible(key, val string) error {
	v := reflect.ValueOf(osr).Elem()
	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i)
		fieldType := v.Type().Field(i)
		originalName := fieldType.Tag.Get("osr")
		if key == originalName && fieldValue.Kind() == reflect.String {
			fieldValue.SetString(val)
			return nil
		}
	}
	return fmt.Errorf("Couldn't set key %s, no corresponding struct field found", key)
}

func parseLine(osrLine string) (string, string, error) {
	if osrLine == "" {
		return "", "", nil
	}

	vals := strings.Split(osrLine, "=")
	if len(vals) != 2 {
		return "", "", fmt.Errorf("Expected %s to split by '=' char into two strings, instead got %d strings", osrLine, len(vals))
	}
	key := vals[0]
	val := stripQuotes(vals[1])
	return key, val, nil
}

func (osr *OsRelease) ParseOsRelease(osReleaseContents []byte) error {
	r := bytes.NewReader(osReleaseContents)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		key, val, err := parseLine(scanner.Text())
		if err != nil {
			log.Warnf("Warning: got an invalid line error parsing /etc/os-release: %s", err)
			continue
		}
		if err := osr.setIfPossible(key, val); err != nil {
			log.Debug(err)
		}
	}
	return nil
}

func NewOsRelease(contents []byte) (*OsRelease, error) {
	osr := &OsRelease{}
	if err := osr.ParseOsRelease(contents); err != nil {
		return nil, err
	}
	return osr, nil
}

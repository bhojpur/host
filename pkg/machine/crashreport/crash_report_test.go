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
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/bugsnag/bugsnag-go"
	"github.com/stretchr/testify/assert"
)

func TestFileIsNotReadWhenNotExisting(t *testing.T) {
	metaData := bugsnag.MetaData{}
	addFile("not existing", &metaData)
	assert.Empty(t, metaData)
}

func TestRead(t *testing.T) {
	metaData := bugsnag.MetaData{}
	content := "foo\nbar\nqix\n"
	fileName := createTempFile(t, content)
	defer os.Remove(fileName)
	addFile(fileName, &metaData)
	assert.Equal(t, "foo\nbar\nqix\n", metaData["logfile"][filepath.Base(fileName)])
}

func createTempFile(t *testing.T, content string) string {
	file, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	if err := ioutil.WriteFile(file.Name(), []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
	return file.Name()
}

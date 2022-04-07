package cert

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
)

func TestGenerateCACertificate(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "machine-test-")
	if err != nil {
		t.Fatal(err)
	}
	// cleanup
	defer os.RemoveAll(tmpDir)

	caCertPath := filepath.Join(tmpDir, "ca.pem")
	caKeyPath := filepath.Join(tmpDir, "key.pem")
	testOrg := "test-org"
	bits := 2048
	if err := GenerateCACertificate(caCertPath, caKeyPath, testOrg, bits); err != nil {
		t.Fatal(err)
	}

	if _, err := os.Stat(caCertPath); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(caKeyPath); err != nil {
		t.Fatal(err)
	}
}

func TestGenerateCert(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "machine-test-")
	if err != nil {
		t.Fatal(err)
	}
	// cleanup
	defer os.RemoveAll(tmpDir)

	caCertPath := filepath.Join(tmpDir, "ca.pem")
	caKeyPath := filepath.Join(tmpDir, "key.pem")
	certPath := filepath.Join(tmpDir, "cert.pem")
	keyPath := filepath.Join(tmpDir, "cert-key.pem")
	testOrg := "test-org"
	bits := 2048
	if err := GenerateCACertificate(caCertPath, caKeyPath, testOrg, bits); err != nil {
		t.Fatal(err)
	}

	if _, err := os.Stat(caCertPath); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(caKeyPath); err != nil {
		t.Fatal(err)
	}

	opts := &Options{
		Hosts:       []string{},
		CertFile:    certPath,
		CAKeyFile:   caKeyPath,
		CAFile:      caCertPath,
		KeyFile:     keyPath,
		Org:         testOrg,
		Bits:        bits,
		SwarmMaster: false,
	}

	if err := GenerateCert(opts); err != nil {
		t.Fatal(err)
	}

	if _, err := os.Stat(certPath); err != nil {
		t.Fatalf("certificate not created at %s", certPath)
	}

	if _, err := os.Stat(keyPath); err != nil {
		t.Fatalf("key not created at %s", keyPath)
	}
}

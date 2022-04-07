package rkecerts

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
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"io"

	"github.com/rancher/rke/pki"
	"github.com/rancher/rke/pki/cert"
)

type savedCertificatePKI struct {
	pki.CertificatePKI
	CertPEM string
	KeyPEM  string
}

func LoadString(input string) (map[string]pki.CertificatePKI, error) {
	return Load(bytes.NewBufferString(input))
}

func Load(f io.Reader) (map[string]pki.CertificatePKI, error) {
	saved := map[string]savedCertificatePKI{}
	if err := json.NewDecoder(f).Decode(&saved); err != nil {
		return nil, err
	}

	certs := map[string]pki.CertificatePKI{}

	for name, savedCert := range saved {
		if savedCert.CertPEM != "" {
			certs, err := cert.ParseCertsPEM([]byte(savedCert.CertPEM))
			if err != nil {
				return nil, err
			}

			if len(certs) == 0 {
				return nil, fmt.Errorf("failed to parse certs, 0 found")
			}

			savedCert.Certificate = certs[0]
		}

		if savedCert.KeyPEM != "" {
			key, err := cert.ParsePrivateKeyPEM([]byte(savedCert.KeyPEM))
			if err != nil {
				return nil, err
			}
			savedCert.Key = key.(*rsa.PrivateKey)
		}

		certs[name] = savedCert.CertificatePKI
	}

	return certs, nil
}

func ToString(certs map[string]pki.CertificatePKI) (string, error) {
	output := &bytes.Buffer{}
	err := Save(certs, output)
	return output.String(), err
}

func Save(certs map[string]pki.CertificatePKI, w io.Writer) error {
	toSave := map[string]savedCertificatePKI{}

	for name, bundleCert := range certs {
		toSaveCert := savedCertificatePKI{
			CertificatePKI: bundleCert,
		}

		if toSaveCert.Certificate != nil {
			toSaveCert.CertPEM = string(cert.EncodeCertPEM(toSaveCert.Certificate))
		}

		if toSaveCert.Key != nil {
			toSaveCert.KeyPEM = string(cert.EncodePrivateKeyPEM(toSaveCert.Key))
		}

		toSaveCert.Certificate = nil
		toSaveCert.Key = nil

		toSave[name] = toSaveCert
	}

	return json.NewEncoder(w).Encode(toSave)
}

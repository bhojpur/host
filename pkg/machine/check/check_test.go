package check

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
	"errors"
	"testing"

	"crypto/tls"

	"github.com/bhojpur/host/pkg/machine/auth"
	"github.com/bhojpur/host/pkg/machine/cert"
	"github.com/stretchr/testify/assert"
)

type FakeValidateCertificate struct {
	IsValid bool
	Err     error
}

type FakeCertGenerator struct {
	fakeValidateCertificate *FakeValidateCertificate
}

func (fcg FakeCertGenerator) GenerateCACertificate(certFile, keyFile, org string, bits int) error {
	return nil
}

func (fcg FakeCertGenerator) GenerateCert(opts *cert.Options) error {
	return nil
}

func (fcg FakeCertGenerator) ValidateCertificate(addr string, authOptions *auth.Options) (bool, error) {
	return fcg.fakeValidateCertificate.IsValid, fcg.fakeValidateCertificate.Err
}

func (fcg FakeCertGenerator) ReadTLSConfig(addr string, authOptions *auth.Options) (*tls.Config, error) {
	return nil, nil
}

func TestCheckCert(t *testing.T) {
	errCertsExpired := errors.New("Certs have expired")

	cases := []struct {
		hostURL     string
		authOptions *auth.Options
		valid       bool
		checkErr    error
		expectedErr error
	}{
		{"192.168.99.100:2376", &auth.Options{}, true, nil, nil},
		{"192.168.99.100:2376", &auth.Options{}, false, nil, ErrCertInvalid{wrappedErr: nil, hostURL: "192.168.99.100:2376"}},
		{"192.168.99.100:2376", &auth.Options{}, false, errCertsExpired, ErrCertInvalid{wrappedErr: errCertsExpired, hostURL: "192.168.99.100:2376"}},
	}

	for _, c := range cases {
		fcg := FakeCertGenerator{fakeValidateCertificate: &FakeValidateCertificate{c.valid, c.checkErr}}
		cert.SetCertGenerator(fcg)
		err := checkCert(c.hostURL, c.authOptions)
		assert.Equal(t, c.expectedErr, err)
	}
}

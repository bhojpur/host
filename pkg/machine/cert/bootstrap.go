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
	"errors"
	"fmt"
	"os"

	"github.com/bhojpur/host/pkg/machine/auth"
	"github.com/bhojpur/host/pkg/machine/log"
	mutils "github.com/bhojpur/host/pkg/machine/utils"
)

func createCACert(authOptions *auth.Options, caOrg string, bits int) error {
	caCertPath := authOptions.CaCertPath
	caPrivateKeyPath := authOptions.CaPrivateKeyPath

	log.Infof("Creating CA: %s", caCertPath)

	// check if the key path exists; if so, error
	if _, err := os.Stat(caPrivateKeyPath); err == nil {
		return errors.New("certificate authority key already exists")
	}

	if err := GenerateCACertificate(caCertPath, caPrivateKeyPath, caOrg, bits); err != nil {
		return fmt.Errorf("generating CA certificate failed: %s", err)
	}

	return nil
}

func createCert(authOptions *auth.Options, org string, bits int) error {
	certDir := authOptions.CertDir
	caCertPath := authOptions.CaCertPath
	caPrivateKeyPath := authOptions.CaPrivateKeyPath
	clientCertPath := authOptions.ClientCertPath
	clientKeyPath := authOptions.ClientKeyPath

	log.Infof("Creating client certificate: %s", clientCertPath)

	if _, err := os.Stat(certDir); err != nil {
		if os.IsNotExist(err) {
			if err := os.Mkdir(certDir, 0700); err != nil {
				return fmt.Errorf("failure creating machine client cert dir: %s", err)
			}
		} else {
			return err
		}
	}

	// check if the key path exists; if so, error
	if _, err := os.Stat(clientKeyPath); err == nil {
		return errors.New("client key already exists")
	}

	// Used to generate the client certificate.
	certOptions := &Options{
		Hosts:       []string{""},
		CertFile:    clientCertPath,
		KeyFile:     clientKeyPath,
		CAFile:      caCertPath,
		CAKeyFile:   caPrivateKeyPath,
		Org:         org,
		Bits:        bits,
		SwarmMaster: false,
	}

	if err := GenerateCert(certOptions); err != nil {
		return fmt.Errorf("failure generating client certificate: %s", err)
	}

	return nil
}

func BootstrapCertificates(authOptions *auth.Options) error {
	certDir := authOptions.CertDir
	caCertPath := authOptions.CaCertPath
	clientCertPath := authOptions.ClientCertPath
	clientKeyPath := authOptions.ClientKeyPath
	caPrivateKeyPath := authOptions.CaPrivateKeyPath

	// TODO: I'm not super happy about this use of "org", the user should
	// have to specify it explicitly instead of implicitly basing it on
	// $USER.
	caOrg := mutils.GetUsername()
	org := caOrg + ".<bootstrap>"

	bits := 2048

	if _, err := os.Stat(certDir); err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(certDir, 0700); err != nil {
				return fmt.Errorf("creating machine certificate dir failed: %s", err)
			}
		} else {
			return err
		}
	}

	if _, err := os.Stat(caCertPath); os.IsNotExist(err) {
		if err := createCACert(authOptions, caOrg, bits); err != nil {
			return err
		}
	} else {
		current, err := CheckCertificateDate(caCertPath)
		if err != nil {
			return err
		}
		if !current {
			log.Info("CA certificate is outdated and needs to be regenerated")
			os.Remove(caPrivateKeyPath)
			if err := createCACert(authOptions, caOrg, bits); err != nil {
				return err
			}
		}
	}

	if _, err := os.Stat(clientCertPath); os.IsNotExist(err) {
		if err := createCert(authOptions, org, bits); err != nil {
			return err
		}
	} else {
		current, err := CheckCertificateDate(clientCertPath)
		if err != nil {
			return err
		}
		if !current {
			log.Info("Client certificate is outdated and needs to be regenerated")
			os.Remove(clientKeyPath)
			if err := createCert(authOptions, org, bits); err != nil {
				return err
			}
		}
	}

	return nil
}

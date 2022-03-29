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
	"fmt"
	"net/url"
	"strings"

	"github.com/bhojpur/host/pkg/core/auth"
	"github.com/bhojpur/host/pkg/core/cert"
	"github.com/bhojpur/host/pkg/core/host"
)

var (
	DefaultConnChecker ConnChecker
	ErrSwarmNotStarted = errors.New("Connection to Swarm cannot be checked but the certs are valid. Maybe swarm is not started")
)

func init() {
	DefaultConnChecker = &MachineConnChecker{}
}

// ErrCertInvalid for when the cert is computed to be invalid.
type ErrCertInvalid struct {
	wrappedErr error
	hostURL    string
}

func (e ErrCertInvalid) Error() string {
	return fmt.Sprintf(`There was an error validating certificates for Bhojpur Host %q: %s
You can attempt to regenerate them using 'hostutl regenerate-certs [name]'.
Be advised that this will trigger a Bhojpur Host daemon restart which might stop running containers.
`, e.hostURL, e.wrappedErr)
}

type ConnChecker interface {
	Check(*host.Host, bool) (bhojpurHost string, authOptions *auth.Options, err error)
}

type MachineConnChecker struct{}

func (mcc *MachineConnChecker) Check(h *host.Host, swarm bool) (string, *auth.Options, error) {
	bhojpurHost, err := h.Driver.GetURL()
	if err != nil {
		return "", &auth.Options{}, err
	}

	bhojpurURL := bhojpurHost
	if swarm {
		bhojpurURL, err = parseSwarm(bhojpurHost, h)
		if err != nil {
			return "", &auth.Options{}, err
		}
	}

	u, err := url.Parse(bhojpurURL)
	if err != nil {
		return "", &auth.Options{}, fmt.Errorf("Error parsing Bhojpur Host URL: %s", err)
	}

	authOptions := h.AuthOptions()

	if err := checkCert(u.Host, authOptions); err != nil {
		if swarm {
			// Connection to the swarm port cannot be checked. Maybe it's just the swarm containers that are down
			// TODO: check the containers and restart them
			// Let's check the non-swarm connection to give a better error message to the user.
			if _, _, err := mcc.Check(h, false); err == nil {
				return "", &auth.Options{}, ErrSwarmNotStarted
			}
		}

		return "", &auth.Options{}, fmt.Errorf("Error checking and/or regenerating the certs: %s", err)
	}

	return bhojpurURL, authOptions, nil
}

func checkCert(hostURL string, authOptions *auth.Options) error {
	valid, err := cert.ValidateCertificate(hostURL, authOptions)
	if !valid || err != nil {
		return ErrCertInvalid{
			wrappedErr: err,
			hostURL:    hostURL,
		}
	}

	return nil
}

// TODO: This could use a unit test.
func parseSwarm(hostURL string, h *host.Host) (string, error) {
	swarmOptions := h.HostOptions.SwarmOptions

	if !swarmOptions.Master {
		return "", fmt.Errorf("%q is not a Bhojpur Host swarm master. The --swarm flag is intended for use with swarm masters", h.Name)
	}

	u, err := url.Parse(swarmOptions.Host)
	if err != nil {
		return "", fmt.Errorf("There was an error parsing the url: %s", err)
	}
	parts := strings.Split(u.Host, ":")
	swarmPort := parts[1]

	// get IP of machine to replace in case swarm host is 0.0.0.0
	mURL, err := url.Parse(hostURL)
	if err != nil {
		return "", fmt.Errorf("There was an error parsing the url: %s", err)
	}

	mParts := strings.Split(mURL.Host, ":")
	machineIP := mParts[0]

	hostURL = fmt.Sprintf("tcp://%s:%s", machineIP, swarmPort)

	return hostURL, nil
}

package client

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

	"github.com/bhojpur/host/pkg/machine/auth"
)

type URLer interface {
	// URL returns the Bhojpur Host URL
	URL() (string, error)
}

type AuthOptionser interface {
	// AuthOptions returns the authOptions
	AuthOptions() *auth.Options
}

type BhojpurHost interface {
	URLer
	AuthOptionser
}

type RemoteBhojpur struct {
	HostURL    string
	AuthOption *auth.Options
}

// URL returns the Bhojpur Host URL
func (rd *RemoteBhojpur) URL() (string, error) {
	if rd == nil || rd.HostURL == "" {
		return "", fmt.Errorf("Bhojpur Host URL not set")
	}

	return rd.HostURL, nil
}

// AuthOptions returns the authOptions
func (rd *RemoteBhojpur) AuthOptions() *auth.Options {
	if rd == nil {
		return nil
	}
	return rd.AuthOption
}

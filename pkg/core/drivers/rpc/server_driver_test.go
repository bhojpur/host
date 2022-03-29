package rpcdriver

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

	"github.com/bhojpur/host/pkg/drivers/fakedriver"
	"github.com/stretchr/testify/assert"
)

type panicDriver struct {
	*fakedriver.Driver
	panicErr  error
	returnErr error
}

type FakeStacker struct {
	trace []byte
}

func (fs *FakeStacker) Stack() []byte {
	return fs.trace
}

func (p *panicDriver) Create() error {
	if p.panicErr != nil {
		panic(p.panicErr)
	}
	return p.returnErr
}

func TestRPCServerDriverCreate(t *testing.T) {
	testCases := []struct {
		description  string
		expectedErr  error
		serverDriver *RPCServerDriver
		stacker      Stacker
	}{
		{
			description: "Happy path",
			expectedErr: nil,
			serverDriver: &RPCServerDriver{
				ActualDriver: &panicDriver{
					returnErr: nil,
				},
			},
		},
		{
			description: "Normal error, no panic",
			expectedErr: errors.New("API not available"),
			serverDriver: &RPCServerDriver{
				ActualDriver: &panicDriver{
					returnErr: errors.New("API not available"),
				},
			},
		},
		{
			description: "Panic happened during create",
			expectedErr: errors.New("Panic in the driver: index out of range\nSTACK TRACE"),
			serverDriver: &RPCServerDriver{
				ActualDriver: &panicDriver{
					panicErr: errors.New("index out of range"),
				},
			},
			stacker: &FakeStacker{
				trace: []byte("STACK TRACE"),
			},
		},
	}

	for _, tc := range testCases {
		stdStacker = tc.stacker
		assert.Equal(t, tc.expectedErr, tc.serverDriver.Create(nil, nil))
	}
}

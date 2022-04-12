package service

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
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) {
	check.TestingT(t)
}

type StubTestSuite struct {
}

type testCase struct {
	input       string
	expectedVal string
	expectedErr error
}

var _ = check.Suite(&StubTestSuite{})

func (s *StubTestSuite) SetUpSuite(c *check.C) {
}

func newTestCase(input string, expectedVal string, isErr bool) testCase {
	var expectedErr error

	if isErr {
		expectedErr = fmt.Errorf("failed to parse port from address [%s]", input)
	}
	return testCase{
		input:       input,
		expectedVal: expectedVal,
		expectedErr: expectedErr,
	}
}

func TestPortOnly(t *testing.T) {
	assert := assert.New(t)

	testCases := []testCase{
		// strings should be of the form "string:port"
		newTestCase("asdf", "", true),
		newTestCase("a:asdf", "", true),
		newTestCase("3000", "", true),
		newTestCase("300:asdf", "", true),
		newTestCase("300!asdf", "", true),
		newTestCase("a:as:300", "", true),
		newTestCase(":300:", "", true),
		newTestCase(":::", "", true),
		newTestCase("asdf.asdf:99999999", "", true),
		newTestCase("asdf.asdf:-99999999", "", true),
		newTestCase("300.com:3000", "3000", false),
		newTestCase("a:200", "200", false),
		newTestCase("a.com:3000", "3000", false),
	}

	for _, test := range testCases {
		port, err := portOnly(test.input)

		assert.Equal(test.expectedVal, port)
		if test.expectedErr != nil {
			assert.Contains(err.Error(), test.expectedErr.Error())
		} else {
			assert.Nil(err)
		}
	}
}

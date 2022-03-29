package amazonec2

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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccessKeyIsMandatoryWhenSystemCredentialsAreNotPresent(t *testing.T) {
	awsCreds := NewAWSCredentials("", "", "")
	awsCreds.fallbackProvider = nil

	_, err := awsCreds.Credentials().Get()
	assert.Error(t, err)
}

func TestAccessKeyIsMandatoryEvenIfSecretKeyIsPassedWhenSystemCredentialsAreNotPresent(t *testing.T) {
	awsCreds := NewAWSCredentials("", "secret", "")
	awsCreds.fallbackProvider = nil

	_, err := awsCreds.Credentials().Get()
	assert.Error(t, err)
}

func TestSecretKeyIsMandatoryWhenSystemCredentialsAreNotPresent(t *testing.T) {
	awsCreds := NewAWSCredentials("access", "", "")
	awsCreds.fallbackProvider = nil

	_, err := awsCreds.Credentials().Get()
	assert.Error(t, err)
}

func TestFallbackCredentialsAreLoadedWhenAccessKeyAndSecretKeyAreMissing(t *testing.T) {
	awsCreds := NewAWSCredentials("", "", "")
	awsCreds.fallbackProvider = &fallbackCredentials{}

	creds, err := awsCreds.Credentials().Get()

	assert.NoError(t, err)
	assert.Equal(t, "fallback_access", creds.AccessKeyID)
	assert.Equal(t, "fallback_secret", creds.SecretAccessKey)
	assert.Equal(t, "fallback_token", creds.SessionToken)
}

func TestFallbackCredentialsAreLoadedWhenAccessKeyIsMissing(t *testing.T) {
	awsCreds := NewAWSCredentials("", "secret", "")
	awsCreds.fallbackProvider = &fallbackCredentials{}

	creds, err := awsCreds.Credentials().Get()

	assert.NoError(t, err)
	assert.Equal(t, "fallback_access", creds.AccessKeyID)
	assert.Equal(t, "fallback_secret", creds.SecretAccessKey)
	assert.Equal(t, "fallback_token", creds.SessionToken)
}

func TestFallbackCredentialsAreLoadedWhenSecretKeyIsMissing(t *testing.T) {
	awsCreds := NewAWSCredentials("access", "", "")
	awsCreds.fallbackProvider = &fallbackCredentials{}

	creds, err := awsCreds.Credentials().Get()

	assert.NoError(t, err)
	assert.Equal(t, "fallback_access", creds.AccessKeyID)
	assert.Equal(t, "fallback_secret", creds.SecretAccessKey)
	assert.Equal(t, "fallback_token", creds.SessionToken)
}

func TestOptionCredentialsAreLoadedWhenAccessKeyAndSecretKeyAreProvided(t *testing.T) {
	awsCreds := NewAWSCredentials("access", "secret", "")
	awsCreds.fallbackProvider = &fallbackCredentials{}

	creds, err := awsCreds.Credentials().Get()

	assert.NoError(t, err)
	assert.Equal(t, "access", creds.AccessKeyID)
	assert.Equal(t, "secret", creds.SecretAccessKey)
	assert.Equal(t, "", creds.SessionToken)
}

func TestFallbackCredentialsAreLoadedIfStaticCredentialsGenerateError(t *testing.T) {
	awsCreds := NewAWSCredentials("access", "secret", "token")
	awsCreds.fallbackProvider = &fallbackCredentials{}
	awsCreds.providerFactory = &errorCredentialsProvider{}

	creds, err := awsCreds.Credentials().Get()

	assert.NoError(t, err)
	assert.Equal(t, "fallback_access", creds.AccessKeyID)
	assert.Equal(t, "fallback_secret", creds.SecretAccessKey)
	assert.Equal(t, "fallback_token", creds.SessionToken)
}

func TestErrorGeneratedWhenAllProvidersGenerateErrors(t *testing.T) {
	awsCreds := NewAWSCredentials("access", "secret", "token")
	awsCreds.fallbackProvider = &errorFallbackCredentials{}
	awsCreds.providerFactory = &errorCredentialsProvider{}

	_, err := awsCreds.Credentials().Get()
	assert.Error(t, err)
}

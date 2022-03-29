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
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

type awsCredentials interface {
	Credentials() *credentials.Credentials
}

type ProviderFactory interface {
	NewStaticProvider(id, secret, token string) credentials.Provider
}

type defaultAWSCredentials struct {
	AccessKey        string
	SecretKey        string
	SessionToken     string
	providerFactory  ProviderFactory
	fallbackProvider awsCredentials
}

func NewAWSCredentials(id, secret, token string) *defaultAWSCredentials {
	creds := defaultAWSCredentials{
		AccessKey:        id,
		SecretKey:        secret,
		SessionToken:     token,
		fallbackProvider: &AwsDefaultCredentialsProvider{},
		providerFactory:  &defaultProviderFactory{},
	}
	return &creds
}

func (c *defaultAWSCredentials) Credentials() *credentials.Credentials {
	providers := []credentials.Provider{}
	if c.AccessKey != "" && c.SecretKey != "" {
		providers = append(providers, c.providerFactory.NewStaticProvider(c.AccessKey, c.SecretKey, c.SessionToken))
	}
	if c.fallbackProvider != nil {
		fallbackCreds, err := c.fallbackProvider.Credentials().Get()
		if err == nil {
			providers = append(providers, &credentials.StaticProvider{Value: fallbackCreds})
		}
	}
	return credentials.NewChainCredentials(providers)
}

type AwsDefaultCredentialsProvider struct{}

func (c *AwsDefaultCredentialsProvider) Credentials() *credentials.Credentials {
	return session.New().Config.Credentials
}

type defaultProviderFactory struct{}

func (c *defaultProviderFactory) NewStaticProvider(id, secret, token string) credentials.Provider {
	return &credentials.StaticProvider{Value: credentials.Value{
		AccessKeyID:     id,
		SecretAccessKey: secret,
		SessionToken:    token,
	}}
}

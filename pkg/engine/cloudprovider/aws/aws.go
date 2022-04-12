package aws

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
	"fmt"

	"github.com/go-ini/ini"

	v3 "github.com/bhojpur/host/pkg/engine/types"
)

const (
	AWSCloudProviderName = "aws"
	AWSConfig            = "AWSConfig"
)

type CloudProvider struct {
	Config *v3.AWSCloudProvider
	Name   string
}

func GetInstance() *CloudProvider {
	return &CloudProvider{}
}

func (p *CloudProvider) Init(cloudProviderConfig v3.CloudProvider) error {
	p.Name = AWSCloudProviderName
	if cloudProviderConfig.AWSCloudProvider == nil {
		return nil
	}
	p.Config = cloudProviderConfig.AWSCloudProvider

	return nil
}
func (p *CloudProvider) GetName() string {
	return p.Name
}

func (p *CloudProvider) GenerateCloudConfigFile() (string, error) {
	if p.Config == nil {
		return "", nil
	}
	// Generate INI style configuration
	buf := new(bytes.Buffer)
	cloudConfig, _ := ini.LoadSources(ini.LoadOptions{IgnoreInlineComment: true}, []byte(""))
	if err := ini.ReflectFrom(cloudConfig, p.Config); err != nil {
		return "", fmt.Errorf("Failed to parse AWS cloud config")
	}
	if _, err := cloudConfig.WriteTo(buf); err != nil {
		return "", err
	}
	return buf.String(), nil
}

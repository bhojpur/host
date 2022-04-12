package openstack

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

	v3 "github.com/bhojpur/host/pkg/engine/types"
	"github.com/go-ini/ini"
)

const (
	OpenstackCloudProviderName = "openstack"
)

type CloudProvider struct {
	Config *v3.OpenstackCloudProvider
	Name   string
}

func GetInstance() *CloudProvider {
	return &CloudProvider{}
}

func (p *CloudProvider) Init(cloudProviderConfig v3.CloudProvider) error {
	if cloudProviderConfig.OpenstackCloudProvider == nil {
		return fmt.Errorf("Openstack Cloud Provider Config is empty")
	}
	p.Name = OpenstackCloudProviderName
	if cloudProviderConfig.Name != "" {
		p.Name = cloudProviderConfig.Name
	}
	p.Config = cloudProviderConfig.OpenstackCloudProvider
	return nil
}

func (p *CloudProvider) GetName() string {
	return p.Name
}

func (p *CloudProvider) GenerateCloudConfigFile() (string, error) {
	// Generate INI style configuration
	buf := new(bytes.Buffer)
	cloudConfig, _ := ini.LoadSources(ini.LoadOptions{IgnoreInlineComment: true}, []byte(""))
	if err := ini.ReflectFrom(cloudConfig, p.Config); err != nil {
		return "", fmt.Errorf("Failed to parse Openstack cloud config")
	}
	if _, err := cloudConfig.WriteTo(buf); err != nil {
		return "", err
	}
	return buf.String(), nil
}

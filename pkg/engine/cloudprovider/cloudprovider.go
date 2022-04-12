package cloudprovider

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
	"github.com/bhojpur/host/pkg/engine/cloudprovider/aws"
	"github.com/bhojpur/host/pkg/engine/cloudprovider/azure"
	"github.com/bhojpur/host/pkg/engine/cloudprovider/custom"
	"github.com/bhojpur/host/pkg/engine/cloudprovider/harvester"
	"github.com/bhojpur/host/pkg/engine/cloudprovider/openstack"
	"github.com/bhojpur/host/pkg/engine/cloudprovider/vsphere"
	v3 "github.com/bhojpur/host/pkg/engine/types"
)

type CloudProvider interface {
	Init(cloudProviderConfig v3.CloudProvider) error
	GenerateCloudConfigFile() (string, error)
	GetName() string
}

func InitCloudProvider(cloudProviderConfig v3.CloudProvider) (CloudProvider, error) {
	var p CloudProvider
	if cloudProviderConfig.AWSCloudProvider != nil || cloudProviderConfig.Name == aws.AWSCloudProviderName {
		p = aws.GetInstance()
	}
	if cloudProviderConfig.AzureCloudProvider != nil || cloudProviderConfig.Name == azure.AzureCloudProviderName {
		p = azure.GetInstance()
	}
	if cloudProviderConfig.OpenstackCloudProvider != nil || cloudProviderConfig.Name == openstack.OpenstackCloudProviderName {
		p = openstack.GetInstance()
	}
	if cloudProviderConfig.VsphereCloudProvider != nil || cloudProviderConfig.Name == vsphere.VsphereCloudProviderName {
		p = vsphere.GetInstance()
	}
	if cloudProviderConfig.HarvesterCloudProvider != nil || cloudProviderConfig.Name == harvester.HarvesterCloudProviderName {
		p = harvester.GetInstance()
	}
	if cloudProviderConfig.CustomCloudProvider != "" {
		p = custom.GetInstance()
	}

	if p != nil {
		if err := p.Init(cloudProviderConfig); err != nil {
			return nil, err
		}
	}
	return p, nil
}

package v3

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
	"github.com/bhojpur/host/pkg/core/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GlobalDns struct {
	types.Namespaced

	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GlobalDNSSpec   `json:"spec,omitempty"`
	Status GlobalDNSStatus `json:"status,omitempty"`
}

type GlobalDNSSpec struct {
	FQDN                string   `json:"fqdn,omitempty" bhojpur:"type=hostname,required"`
	TTL                 int64    `json:"ttl,omitempty" bhojpur:"default=300"`
	ProjectNames        []string `json:"projectNames" bhojpur:"type=array[reference[project]],noupdate"`
	MultiClusterAppName string   `json:"multiClusterAppName,omitempty" bhojpur:"type=reference[multiClusterApp]"`
	ProviderName        string   `json:"providerName,omitempty" bhojpur:"type=reference[globalDnsProvider],required"`
	Members             []Member `json:"members,omitempty"`
}

type GlobalDNSStatus struct {
	Endpoints        []string            `json:"endpoints,omitempty"`
	ClusterEndpoints map[string][]string `json:"clusterEndpoints,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GlobalDnsProvider struct {
	types.Namespaced

	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	//ObjectMeta.Name = GlobalDNSProviderID
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec GlobalDNSProviderSpec `json:"spec,omitempty"`
}

type GlobalDNSProviderSpec struct {
	Route53ProviderConfig    *Route53ProviderConfig    `json:"route53ProviderConfig,omitempty"`
	CloudflareProviderConfig *CloudflareProviderConfig `json:"cloudflareProviderConfig,omitempty"`
	AlidnsProviderConfig     *AlidnsProviderConfig     `json:"alidnsProviderConfig,omitempty"`
	Members                  []Member                  `json:"members,omitempty"`
	RootDomain               string                    `json:"rootDomain"`
}

type Route53ProviderConfig struct {
	AccessKey         string            `json:"accessKey" bhojpur:"notnullable,required,minLength=1"`
	SecretKey         string            `json:"secretKey" bhojpur:"notnullable,required,minLength=1,type=password"`
	CredentialsPath   string            `json:"credentialsPath" bhojpur:"default=/.aws"`
	RoleArn           string            `json:"roleArn,omitempty"`
	Region            string            `json:"region" bhojpur:"default=us-east-1"`
	ZoneType          string            `json:"zoneType" bhojpur:"default=public"`
	AdditionalOptions map[string]string `json:"additionalOptions,omitempty"`
}

type CloudflareProviderConfig struct {
	APIKey            string            `json:"apiKey" bhojpur:"notnullable,required,minLength=1,type=password"`
	APIEmail          string            `json:"apiEmail" bhojpur:"notnullable,required,minLength=1"`
	ProxySetting      *bool             `json:"proxySetting" bhojpur:"default=true"`
	AdditionalOptions map[string]string `json:"additionalOptions,omitempty"`
}

type UpdateGlobalDNSTargetsInput struct {
	ProjectNames []string `json:"projectNames" bhojpur:"type=array[reference[project]]"`
}

type AlidnsProviderConfig struct {
	AccessKey         string            `json:"accessKey" bhojpur:"notnullable,required,minLength=1"`
	SecretKey         string            `json:"secretKey" bhojpur:"notnullable,required,minLength=1,type=password"`
	AdditionalOptions map[string]string `json:"additionalOptions,omitempty"`
}

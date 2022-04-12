package pki

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
	"context"
	"reflect"
	"testing"

	"github.com/bhojpur/host/pkg/engine/hosts"
	v3 "github.com/bhojpur/host/pkg/engine/types"
	"github.com/stretchr/testify/assert"
)

func TestDeleteUnusedCerts(t *testing.T) {
	tests := []struct {
		ctx             context.Context
		name            string
		certs           map[string]CertificatePKI
		certName        string
		hosts           []*hosts.Host
		expectLeftCerts map[string]CertificatePKI
	}{
		{
			ctx:  context.Background(),
			name: "Keep valid etcd certs",
			certs: map[string]CertificatePKI{
				"kube-etcd-172-17-0-3":    CertificatePKI{},
				"kube-etcd-172-17-0-4":    CertificatePKI{},
				"kube-node":               CertificatePKI{},
				"kube-kubelet-172-17-0-4": CertificatePKI{},
				"kube-apiserver":          CertificatePKI{},
				"kube-proxy":              CertificatePKI{},
			},
			certName: EtcdCertName,
			hosts: []*hosts.Host{
				{BKEConfigNode: v3.BKEConfigNode{
					Address: "172.17.0.3",
				}},
				{BKEConfigNode: v3.BKEConfigNode{
					Address: "172.17.0.4",
				}},
			},
			expectLeftCerts: map[string]CertificatePKI{
				"kube-etcd-172-17-0-3":    CertificatePKI{},
				"kube-etcd-172-17-0-4":    CertificatePKI{},
				"kube-node":               CertificatePKI{},
				"kube-kubelet-172-17-0-4": CertificatePKI{},
				"kube-apiserver":          CertificatePKI{},
				"kube-proxy":              CertificatePKI{},
			},
		},
		{
			ctx:  context.Background(),
			name: "Keep valid kubelet certs",
			certs: map[string]CertificatePKI{
				"kube-kubelet-172-17-0-5": CertificatePKI{},
				"kube-kubelet-172-17-0-6": CertificatePKI{},
				"kube-node":               CertificatePKI{},
				"kube-apiserver":          CertificatePKI{},
				"kube-proxy":              CertificatePKI{},
				"kube-etcd-172-17-0-6":    CertificatePKI{},
			},
			certName: KubeletCertName,
			hosts: []*hosts.Host{
				{BKEConfigNode: v3.BKEConfigNode{
					Address: "172.17.0.5",
				}},
				{BKEConfigNode: v3.BKEConfigNode{
					Address: "172.17.0.6",
				}},
			},
			expectLeftCerts: map[string]CertificatePKI{
				"kube-kubelet-172-17-0-5": CertificatePKI{},
				"kube-kubelet-172-17-0-6": CertificatePKI{},
				"kube-node":               CertificatePKI{},
				"kube-apiserver":          CertificatePKI{},
				"kube-proxy":              CertificatePKI{},
				"kube-etcd-172-17-0-6":    CertificatePKI{},
			},
		},
		{
			ctx:  context.Background(),
			name: "Remove unused etcd certs",
			certs: map[string]CertificatePKI{
				"kube-etcd-172-17-0-11":    CertificatePKI{},
				"kube-etcd-172-17-0-10":    CertificatePKI{},
				"kube-kubelet-172-17-0-11": CertificatePKI{},
				"kube-node":                CertificatePKI{},
				"kube-apiserver":           CertificatePKI{},
				"kube-proxy":               CertificatePKI{},
			},
			certName: EtcdCertName,
			hosts: []*hosts.Host{
				{BKEConfigNode: v3.BKEConfigNode{
					Address: "172.17.0.11",
				}},
				{BKEConfigNode: v3.BKEConfigNode{
					Address: "172.17.0.12",
				}},
			},
			expectLeftCerts: map[string]CertificatePKI{
				"kube-etcd-172-17-0-11":    CertificatePKI{},
				"kube-kubelet-172-17-0-11": CertificatePKI{},
				"kube-node":                CertificatePKI{},
				"kube-apiserver":           CertificatePKI{},
				"kube-proxy":               CertificatePKI{},
			},
		},
		{
			ctx:  context.Background(),
			name: "Remove unused kubelet certs",
			certs: map[string]CertificatePKI{
				"kube-kubelet-172-17-0-11": CertificatePKI{},
				"kube-kubelet-172-17-0-10": CertificatePKI{},
				"kube-etcd-172-17-0-10":    CertificatePKI{},
				"kube-node":                CertificatePKI{},
				"kube-apiserver":           CertificatePKI{},
				"kube-proxy":               CertificatePKI{},
			},
			certName: KubeletCertName,
			hosts: []*hosts.Host{
				{BKEConfigNode: v3.BKEConfigNode{
					Address: "172.17.0.11",
				}},
				{BKEConfigNode: v3.BKEConfigNode{
					Address: "172.17.0.12",
				}},
			},
			expectLeftCerts: map[string]CertificatePKI{
				"kube-kubelet-172-17-0-11": CertificatePKI{},
				"kube-etcd-172-17-0-10":    CertificatePKI{},
				"kube-node":                CertificatePKI{},
				"kube-apiserver":           CertificatePKI{},
				"kube-proxy":               CertificatePKI{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleteUnusedCerts(tt.ctx, tt.certs, tt.certName, tt.hosts)
			assert.Equal(t, true, reflect.DeepEqual(tt.certs, tt.expectLeftCerts))
		})
	}
}

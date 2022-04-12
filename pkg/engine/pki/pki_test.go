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
	"crypto/x509"
	"fmt"
	"net"
	"testing"

	v3 "github.com/bhojpur/host/pkg/engine/types"
)

const (
	FakeClusterDomain = "cluster.test"
	FakeClusterCidr   = "10.0.0.1/24"
)

func TestPKI(t *testing.T) {
	bkeConfig := v3.BhojpurKubernetesEngineConfig{
		Nodes: []v3.BKEConfigNode{
			v3.BKEConfigNode{
				Address:          "1.1.1.1",
				InternalAddress:  "192.168.1.5",
				Role:             []string{"controlplane"},
				HostnameOverride: "server1",
			},
		},
		Services: v3.BKEConfigServices{
			KubeAPI: v3.KubeAPIService{
				ServiceClusterIPRange: FakeClusterCidr,
			},
			Kubelet: v3.KubeletService{
				ClusterDomain: FakeClusterDomain,
			},
		},
	}
	certificateMap, err := GenerateBKECerts(context.Background(), bkeConfig, "", "")
	if err != nil {
		t.Fatalf("Failed To generate certificates: %v", err)
	}
	assertEqual(t, certificateMap[CACertName].Certificate.IsCA, true, "")
	roots := x509.NewCertPool()
	roots.AddCert(certificateMap[CACertName].Certificate)

	certificatesToVerify := []string{
		KubeAPICertName,
		KubeNodeCertName,
		KubeProxyCertName,
		KubeControllerCertName,
		KubeSchedulerCertName,
		KubeAdminCertName,
	}
	opts := x509.VerifyOptions{
		Roots:     roots,
		KeyUsages: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
	}
	for _, cert := range certificatesToVerify {
		if _, err := certificateMap[cert].Certificate.Verify(opts); err != nil {
			t.Fatalf("Failed to verify certificate %s: %v", cert, err)
		}
	}
	// Test DNS ALT names
	kubeAPIDNSNames := []string{
		"localhost",
		"kubernetes",
		"kubernetes.default",
		"kubernetes.default.svc",
		"kubernetes.default.svc." + FakeClusterDomain,
	}
	for _, testDNS := range kubeAPIDNSNames {
		assertEqual(
			t,
			isStringInSlice(
				testDNS,
				certificateMap[KubeAPICertName].Certificate.DNSNames),
			true,
			fmt.Sprintf("DNS %s is not found in ALT names of Kube API certificate", testDNS))
	}

	kubernetesServiceIP, err := GetKubernetesServiceIP(FakeClusterCidr)
	if err != nil {
		t.Fatalf("Failed to get kubernetes service ip for service cidr: %v", err)
	}
	// Test ALT IPs
	kubeAPIAltIPs := []net.IP{
		net.ParseIP("127.0.0.1"),
		net.ParseIP(bkeConfig.Nodes[0].InternalAddress),
		net.ParseIP(bkeConfig.Nodes[0].Address),
	}
	kubeAPIAltIPs = append(kubeAPIAltIPs, kubernetesServiceIP...)

	for _, testIP := range kubeAPIAltIPs {
		found := false
		for _, altIP := range certificateMap[KubeAPICertName].Certificate.IPAddresses {
			if testIP.Equal(altIP) {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("IP Address %v is not found in ALT Ips of kube API", testIP)
		}
	}
}

func isStringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

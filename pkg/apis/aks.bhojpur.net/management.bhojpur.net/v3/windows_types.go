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

type WindowsSystemImages struct {
	// Windows nginx-proxy image
	NginxProxy string `yaml:"nginx_proxy" json:"nginxProxy,omitempty"`
	// Kubernetes binaries image
	KubernetesBinaries string `yaml:"kubernetes_binaries" json:"kubernetesBinaries,omitempty"`
	// Kubelet pause image
	KubeletPause string `yaml:"kubelet_pause" json:"kubeletPause,omitempty"`
	// Flannel CNI binaries image
	FlannelCNIBinaries string `yaml:"flannel_cni_binaries" json:"flannelCniBinaries,omitempty"`
	// Calico CNI binaries image
	CalicoCNIBinaries string `yaml:"calico_cni_binaries" json:"calicoCniBinaries,omitempty"`
	// Canal CNI binaries image
	CanalCNIBinaries string `yaml:"canal_cni_binaries" json:"canalCniBinaries,omitempty"`
}

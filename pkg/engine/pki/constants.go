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

import "time"

const (
	K8sBaseDir              = "/etc/kubernetes/"
	CertPathPrefix          = K8sBaseDir + "ssl/"
	CertificatesServiceName = "certificates"
	CrtDownloaderContainer  = "cert-deployer"
	CertFetcherContainer    = "cert-fetcher"
	CertificatesSecretName  = "k8s-certs"
	TempCertPath            = "/etc/kubernetes/.tmp/"
	ClusterConfig           = "cluster.yml"
	ClusterStateFile        = "cluster-state.yml"
	ClusterStateExt         = ".bkestate"
	ClusterStateEnv         = "CLUSTER_STATE"
	BundleCertPath          = "/backup/pki.bundle.tar.gz"

	CACertName                 = "kube-ca"
	RequestHeaderCACertName    = "kube-apiserver-requestheader-ca"
	KubeAPICertName            = "kube-apiserver"
	KubeControllerCertName     = "kube-controller-manager"
	KubeSchedulerCertName      = "kube-scheduler"
	KubeProxyCertName          = "kube-proxy"
	KubeNodeCertName           = "kube-node"
	KubeletCertName            = "kube-kubelet"
	EtcdCertName               = "kube-etcd"
	EtcdClientCACertName       = "kube-etcd-client-ca"
	EtcdClientCertName         = "kube-etcd-client"
	APIProxyClientCertName     = "kube-apiserver-proxy-client"
	ServiceAccountTokenKeyName = "kube-service-account-token"

	KubeNodeCommonName       = "system:node"
	KubeNodeOrganizationName = "system:nodes"

	KubeAdminCertName         = "kube-admin"
	KubeAdminOrganizationName = "system:masters"
	KubeAdminConfigPrefix     = "kube_config_"
	duration365d              = time.Hour * 24 * 365
)

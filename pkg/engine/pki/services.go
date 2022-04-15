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
	"crypto/rsa"
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/bhojpur/host/pkg/cluster/log"
	"github.com/bhojpur/host/pkg/engine/hosts"
	"github.com/bhojpur/host/pkg/engine/pki/cert"
	v3 "github.com/bhojpur/host/pkg/engine/types"
	"github.com/sirupsen/logrus"
)

func GenerateKubeAPICertificate(ctx context.Context, certs map[string]CertificatePKI, bkeConfig v3.BhojpurKubernetesEngineConfig, configPath, configDir string, rotate bool) error {
	// generate API certificate and key
	caCrt := certs[CACertName].Certificate
	caKey := certs[CACertName].Key
	if caCrt == nil || caKey == nil {
		return fmt.Errorf("CA Certificate or Key is empty")
	}
	kubernetesServiceIP, err := GetKubernetesServiceIP(bkeConfig.Services.KubeAPI.ServiceClusterIPRange)
	if err != nil {
		return fmt.Errorf("Failed to get Kubernetes Service IP: %v", err)
	}
	clusterDomain := bkeConfig.Services.Kubelet.ClusterDomain
	cpHosts := hosts.NodesToHosts(bkeConfig.Nodes, controlRole)
	kubeAPIAltNames := GetAltNames(cpHosts, clusterDomain, kubernetesServiceIP, bkeConfig.Authentication.SANs)
	kubeAPICert := certs[KubeAPICertName].Certificate
	if kubeAPICert != nil &&
		reflect.DeepEqual(kubeAPIAltNames.DNSNames, kubeAPICert.DNSNames) &&
		DeepEqualIPsAltNames(kubeAPIAltNames.IPs, kubeAPICert.IPAddresses) && !rotate {
		return nil
	}
	logrus.Info("[certificates] Generating Kubernetes API server certificates")
	var serviceKey *rsa.PrivateKey
	if !rotate {
		serviceKey = certs[KubeAPICertName].Key
	}
	kubeAPICrt, kubeAPIKey, err := GenerateSignedCertAndKey(caCrt, caKey, true, KubeAPICertName, kubeAPIAltNames, serviceKey, nil)
	if err != nil {
		return err
	}
	certs[KubeAPICertName] = ToCertObject(KubeAPICertName, "", "", kubeAPICrt, kubeAPIKey, nil)
	// handle service account tokens in old clusters
	apiCert := certs[KubeAPICertName]
	if certs[ServiceAccountTokenKeyName].Key == nil {
		logrus.Info("[certificates] Generating Service account token key")
		certs[ServiceAccountTokenKeyName] = ToCertObject(ServiceAccountTokenKeyName, ServiceAccountTokenKeyName, "", apiCert.Certificate, apiCert.Key, nil)
	}
	return nil
}

func GenerateKubeAPICSR(ctx context.Context, certs map[string]CertificatePKI, bkeConfig v3.BhojpurKubernetesEngineConfig) error {
	// generate API csr and key
	kubernetesServiceIP, err := GetKubernetesServiceIP(bkeConfig.Services.KubeAPI.ServiceClusterIPRange)
	if err != nil {
		return fmt.Errorf("Failed to get Kubernetes Service IP: %v", err)
	}
	clusterDomain := bkeConfig.Services.Kubelet.ClusterDomain
	cpHosts := hosts.NodesToHosts(bkeConfig.Nodes, controlRole)
	kubeAPIAltNames := GetAltNames(cpHosts, clusterDomain, kubernetesServiceIP, bkeConfig.Authentication.SANs)
	kubeAPICert := certs[KubeAPICertName].Certificate
	oldKubeAPICSR := certs[KubeAPICertName].CSR
	if oldKubeAPICSR != nil &&
		reflect.DeepEqual(kubeAPIAltNames.DNSNames, oldKubeAPICSR.DNSNames) &&
		DeepEqualIPsAltNames(kubeAPIAltNames.IPs, oldKubeAPICSR.IPAddresses) {
		return nil
	}
	logrus.Info("[certificates] Generating Kubernetes API server csr")
	kubeAPICSR, kubeAPIKey, err := GenerateCertSigningRequestAndKey(true, KubeAPICertName, kubeAPIAltNames, certs[KubeAPICertName].Key, nil)
	if err != nil {
		return err
	}
	certs[KubeAPICertName] = ToCertObject(KubeAPICertName, "", "", kubeAPICert, kubeAPIKey, kubeAPICSR)
	return nil
}

func GenerateKubeControllerCertificate(ctx context.Context, certs map[string]CertificatePKI, bkeConfig v3.BhojpurKubernetesEngineConfig, configPath, configDir string, rotate bool) error {
	// generate Kube controller-manager certificate and key
	caCrt := certs[CACertName].Certificate
	caKey := certs[CACertName].Key
	if caCrt == nil || caKey == nil {
		return fmt.Errorf("CA Certificate or Key is empty")
	}
	if certs[KubeControllerCertName].Certificate != nil && !rotate {
		return nil
	}
	logrus.Info("[certificates] Generating Kube Controller certificates")
	var serviceKey *rsa.PrivateKey
	if !rotate {
		serviceKey = certs[KubeControllerCertName].Key
	}
	kubeControllerCrt, kubeControllerKey, err := GenerateSignedCertAndKey(caCrt, caKey, false, getDefaultCN(KubeControllerCertName), nil, serviceKey, nil)
	if err != nil {
		return err
	}
	certs[KubeControllerCertName] = ToCertObject(KubeControllerCertName, "", "", kubeControllerCrt, kubeControllerKey, nil)
	return nil
}

func GenerateKubeControllerCSR(ctx context.Context, certs map[string]CertificatePKI, bkeConfig v3.BhojpurKubernetesEngineConfig) error {
	// generate Kube controller-manager csr and key
	kubeControllerCrt := certs[KubeControllerCertName].Certificate
	kubeControllerCSRPEM := certs[KubeControllerCertName].CSRPEM
	if kubeControllerCSRPEM != "" {
		return nil
	}
	logrus.Info("[certificates] Generating Kube Controller csr")
	kubeControllerCSR, kubeControllerKey, err := GenerateCertSigningRequestAndKey(false, getDefaultCN(KubeControllerCertName), nil, certs[KubeControllerCertName].Key, nil)
	if err != nil {
		return err
	}
	certs[KubeControllerCertName] = ToCertObject(KubeControllerCertName, "", "", kubeControllerCrt, kubeControllerKey, kubeControllerCSR)
	return nil
}

func GenerateKubeSchedulerCertificate(ctx context.Context, certs map[string]CertificatePKI, bkeConfig v3.BhojpurKubernetesEngineConfig, configPath, configDir string, rotate bool) error {
	// generate Kube scheduler certificate and key
	caCrt := certs[CACertName].Certificate
	caKey := certs[CACertName].Key
	if caCrt == nil || caKey == nil {
		return fmt.Errorf("CA Certificate or Key is empty")
	}
	if certs[KubeSchedulerCertName].Certificate != nil && !rotate {
		return nil
	}
	logrus.Info("[certificates] Generating Kube Scheduler certificates")
	var serviceKey *rsa.PrivateKey
	if !rotate {
		serviceKey = certs[KubeSchedulerCertName].Key
	}
	kubeSchedulerCrt, kubeSchedulerKey, err := GenerateSignedCertAndKey(caCrt, caKey, false, getDefaultCN(KubeSchedulerCertName), nil, serviceKey, nil)
	if err != nil {
		return err
	}
	certs[KubeSchedulerCertName] = ToCertObject(KubeSchedulerCertName, "", "", kubeSchedulerCrt, kubeSchedulerKey, nil)
	return nil
}

func GenerateKubeSchedulerCSR(ctx context.Context, certs map[string]CertificatePKI, bkeConfig v3.BhojpurKubernetesEngineConfig) error {
	// generate Kube scheduler csr and key
	kubeSchedulerCrt := certs[KubeSchedulerCertName].Certificate
	kubeSchedulerCSRPEM := certs[KubeSchedulerCertName].CSRPEM
	if kubeSchedulerCSRPEM != "" {
		return nil
	}
	logrus.Info("[certificates] Generating Kube Scheduler csr")
	kubeSchedulerCSR, kubeSchedulerKey, err := GenerateCertSigningRequestAndKey(false, getDefaultCN(KubeSchedulerCertName), nil, certs[KubeSchedulerCertName].Key, nil)
	if err != nil {
		return err
	}
	certs[KubeSchedulerCertName] = ToCertObject(KubeSchedulerCertName, "", "", kubeSchedulerCrt, kubeSchedulerKey, kubeSchedulerCSR)
	return nil
}

func GenerateKubeProxyCertificate(ctx context.Context, certs map[string]CertificatePKI, bkeConfig v3.BhojpurKubernetesEngineConfig, configPath, configDir string, rotate bool) error {
	// generate Kube Proxy certificate and key
	caCrt := certs[CACertName].Certificate
	caKey := certs[CACertName].Key
	if caCrt == nil || caKey == nil {
		return fmt.Errorf("CA Certificate or Key is empty")
	}
	if certs[KubeProxyCertName].Certificate != nil && !rotate {
		return nil
	}
	logrus.Info("[certificates] Generating Kube Proxy certificates")
	var serviceKey *rsa.PrivateKey
	if !rotate {
		serviceKey = certs[KubeProxyCertName].Key
	}
	kubeProxyCrt, kubeProxyKey, err := GenerateSignedCertAndKey(caCrt, caKey, false, getDefaultCN(KubeProxyCertName), nil, serviceKey, nil)
	if err != nil {
		return err
	}
	certs[KubeProxyCertName] = ToCertObject(KubeProxyCertName, "", "", kubeProxyCrt, kubeProxyKey, nil)
	return nil
}

func GenerateKubeProxyCSR(ctx context.Context, certs map[string]CertificatePKI, bkeConfig v3.BhojpurKubernetesEngineConfig) error {
	// generate Kube Proxy csr and key
	kubeProxyCrt := certs[KubeProxyCertName].Certificate
	kubeProxyCSRPEM := certs[KubeProxyCertName].CSRPEM
	if kubeProxyCSRPEM != "" {
		return nil
	}
	logrus.Info("[certificates] Generating Kube Proxy csr")
	kubeProxyCSR, kubeProxyKey, err := GenerateCertSigningRequestAndKey(false, getDefaultCN(KubeProxyCertName), nil, certs[KubeProxyCertName].Key, nil)
	if err != nil {
		return err
	}
	certs[KubeProxyCertName] = ToCertObject(KubeProxyCertName, "", "", kubeProxyCrt, kubeProxyKey, kubeProxyCSR)
	return nil
}

func GenerateKubeNodeCertificate(ctx context.Context, certs map[string]CertificatePKI, bkeConfig v3.BhojpurKubernetesEngineConfig, configPath, configDir string, rotate bool) error {
	// generate kubelet certificate
	caCrt := certs[CACertName].Certificate
	caKey := certs[CACertName].Key
	if caCrt == nil || caKey == nil {
		return fmt.Errorf("CA Certificate or Key is empty")
	}
	if certs[KubeNodeCertName].Certificate != nil && !rotate {
		return nil
	}
	logrus.Info("[certificates] Generating Node certificate")
	var serviceKey *rsa.PrivateKey
	if !rotate {
		serviceKey = certs[KubeProxyCertName].Key
	}
	nodeCrt, nodeKey, err := GenerateSignedCertAndKey(caCrt, caKey, false, KubeNodeCommonName, nil, serviceKey, []string{KubeNodeOrganizationName})
	if err != nil {
		return err
	}
	certs[KubeNodeCertName] = ToCertObject(KubeNodeCertName, KubeNodeCommonName, KubeNodeOrganizationName, nodeCrt, nodeKey, nil)
	return nil
}

func GenerateKubeNodeCSR(ctx context.Context, certs map[string]CertificatePKI, bkeConfig v3.BhojpurKubernetesEngineConfig) error {
	// generate kubelet csr and key
	nodeCrt := certs[KubeNodeCertName].Certificate
	nodeCSRPEM := certs[KubeNodeCertName].CSRPEM
	if nodeCSRPEM != "" {
		return nil
	}
	logrus.Info("[certificates] Generating Node csr and key")
	nodeCSR, nodeKey, err := GenerateCertSigningRequestAndKey(false, KubeNodeCommonName, nil, certs[KubeNodeCertName].Key, []string{KubeNodeOrganizationName})
	if err != nil {
		return err
	}
	certs[KubeNodeCertName] = ToCertObject(KubeNodeCertName, KubeNodeCommonName, KubeNodeOrganizationName, nodeCrt, nodeKey, nodeCSR)
	return nil
}

func GenerateKubeAdminCertificate(ctx context.Context, certs map[string]CertificatePKI, bkeConfig v3.BhojpurKubernetesEngineConfig, configPath, configDir string, rotate bool) error {
	// generate Admin certificate and key
	logrus.Info("[certificates] Generating admin certificates and kubeconfig")
	caCrt := certs[CACertName].Certificate
	caKey := certs[CACertName].Key
	if caCrt == nil || caKey == nil {
		return fmt.Errorf("CA Certificate or Key is empty")
	}
	cpHosts := hosts.NodesToHosts(bkeConfig.Nodes, controlRole)
	if len(configPath) == 0 {
		configPath = ClusterConfig
	}
	localKubeConfigPath := GetLocalKubeConfig(configPath, configDir)
	var serviceKey *rsa.PrivateKey
	if !rotate {
		serviceKey = certs[KubeAdminCertName].Key
	}
	kubeAdminCrt, kubeAdminKey, err := GenerateSignedCertAndKey(caCrt, caKey, false, KubeAdminCertName, nil, serviceKey, []string{KubeAdminOrganizationName})
	if err != nil {
		return err
	}
	kubeAdminCertObj := ToCertObject(KubeAdminCertName, KubeAdminCertName, KubeAdminOrganizationName, kubeAdminCrt, kubeAdminKey, nil)
	if len(cpHosts) > 0 {
		kubeAdminConfig := GetKubeConfigX509WithData(
			"https://"+cpHosts[0].Address+":6443",
			bkeConfig.ClusterName,
			KubeAdminCertName,
			string(cert.EncodeCertPEM(caCrt)),
			string(cert.EncodeCertPEM(kubeAdminCrt)),
			string(cert.EncodePrivateKeyPEM(kubeAdminKey)))
		kubeAdminCertObj.Config = kubeAdminConfig
		kubeAdminCertObj.ConfigPath = localKubeConfigPath
	} else {
		kubeAdminCertObj.Config = ""
	}
	certs[KubeAdminCertName] = kubeAdminCertObj
	return nil
}

func GenerateKubeAdminCSR(ctx context.Context, certs map[string]CertificatePKI, bkeConfig v3.BhojpurKubernetesEngineConfig) error {
	// generate Admin certificate and key
	kubeAdminCrt := certs[KubeAdminCertName].Certificate
	kubeAdminCSRPEM := certs[KubeAdminCertName].CSRPEM
	if kubeAdminCSRPEM != "" {
		return nil
	}
	kubeAdminCSR, kubeAdminKey, err := GenerateCertSigningRequestAndKey(false, KubeAdminCertName, nil, certs[KubeAdminCertName].Key, []string{KubeAdminOrganizationName})
	if err != nil {
		return err
	}
	logrus.Info("[certificates] Generating admin csr and kubeconfig")
	kubeAdminCertObj := ToCertObject(KubeAdminCertName, KubeAdminCertName, KubeAdminOrganizationName, kubeAdminCrt, kubeAdminKey, kubeAdminCSR)
	certs[KubeAdminCertName] = kubeAdminCertObj
	return nil
}

func GenerateAPIProxyClientCertificate(ctx context.Context, certs map[string]CertificatePKI, bkeConfig v3.BhojpurKubernetesEngineConfig, configPath, configDir string, rotate bool) error {
	//generate API server proxy client key and certs
	caCrt := certs[RequestHeaderCACertName].Certificate
	caKey := certs[RequestHeaderCACertName].Key
	if caCrt == nil || caKey == nil {
		return fmt.Errorf("Request Header CA Certificate or Key is empty")
	}
	if certs[APIProxyClientCertName].Certificate != nil && !rotate {
		return nil
	}
	logrus.Info("[certificates] Generating Kubernetes API server proxy client certificates")
	var serviceKey *rsa.PrivateKey
	if !rotate {
		serviceKey = certs[APIProxyClientCertName].Key
	}
	apiserverProxyClientCrt, apiserverProxyClientKey, err := GenerateSignedCertAndKey(caCrt, caKey, true, APIProxyClientCertName, nil, serviceKey, nil)
	if err != nil {
		return err
	}
	certs[APIProxyClientCertName] = ToCertObject(APIProxyClientCertName, "", "", apiserverProxyClientCrt, apiserverProxyClientKey, nil)
	return nil
}

func GenerateAPIProxyClientCSR(ctx context.Context, certs map[string]CertificatePKI, bkeConfig v3.BhojpurKubernetesEngineConfig) error {
	//generate API server proxy client key and certs
	apiserverProxyClientCrt := certs[APIProxyClientCertName].Certificate
	apiserverProxyClientCSRPEM := certs[APIProxyClientCertName].CSRPEM
	if apiserverProxyClientCSRPEM != "" {
		return nil
	}
	logrus.Info("[certificates] Generating Kubernetes API server proxy client csr")
	apiserverProxyClientCSR, apiserverProxyClientKey, err := GenerateCertSigningRequestAndKey(true, APIProxyClientCertName, nil, certs[APIProxyClientCertName].Key, nil)
	if err != nil {
		return err
	}
	certs[APIProxyClientCertName] = ToCertObject(APIProxyClientCertName, "", "", apiserverProxyClientCrt, apiserverProxyClientKey, apiserverProxyClientCSR)
	return nil
}

func GenerateExternalEtcdCertificates(ctx context.Context, certs map[string]CertificatePKI, bkeConfig v3.BhojpurKubernetesEngineConfig, configPath, configDir string, rotate bool) error {
	clientCert, err := cert.ParseCertsPEM([]byte(bkeConfig.Services.Etcd.Cert))
	if err != nil {
		return err
	}
	clientKey, err := cert.ParsePrivateKeyPEM([]byte(bkeConfig.Services.Etcd.Key))
	if err != nil {
		return err
	}
	certs[EtcdClientCertName] = ToCertObject(EtcdClientCertName, "", "", clientCert[0], clientKey.(*rsa.PrivateKey), nil)

	caCert, err := cert.ParseCertsPEM([]byte(bkeConfig.Services.Etcd.CACert))
	if err != nil {
		return err
	}
	certs[EtcdClientCACertName] = ToCertObject(EtcdClientCACertName, "", "", caCert[0], nil, nil)
	return nil
}

func GenerateEtcdCertificates(ctx context.Context, certs map[string]CertificatePKI, bkeConfig v3.BhojpurKubernetesEngineConfig, configPath, configDir string, rotate bool) error {
	caCrt := certs[CACertName].Certificate
	caKey := certs[CACertName].Key
	if caCrt == nil || caKey == nil {
		return fmt.Errorf("CA Certificate or Key is empty")
	}
	kubernetesServiceIP, err := GetKubernetesServiceIP(bkeConfig.Services.KubeAPI.ServiceClusterIPRange)
	if err != nil {
		return fmt.Errorf("Failed to get Kubernetes Service IP: %v", err)
	}
	clusterDomain := bkeConfig.Services.Kubelet.ClusterDomain
	etcdHosts := hosts.NodesToHosts(bkeConfig.Nodes, etcdRole)
	etcdAltNames := GetAltNames(etcdHosts, clusterDomain, kubernetesServiceIP, []string{})
	var (
		dnsNames = make([]string, len(etcdAltNames.DNSNames))
		ips      = []string{}
	)
	copy(dnsNames, etcdAltNames.DNSNames)
	sort.Strings(dnsNames)
	for _, ip := range etcdAltNames.IPs {
		ips = append(ips, ip.String())
	}
	sort.Strings(ips)
	for _, host := range etcdHosts {
		etcdName := GetCrtNameForHost(host, EtcdCertName)
		if _, ok := certs[etcdName]; ok && certs[etcdName].CertificatePEM != "" && !rotate {
			cert := certs[etcdName].Certificate
			if cert != nil && len(dnsNames) == len(cert.DNSNames) && len(ips) == len(cert.IPAddresses) {
				var (
					certDNSNames = make([]string, len(cert.DNSNames))
					certIPs      = []string{}
				)
				copy(certDNSNames, cert.DNSNames)
				sort.Strings(certDNSNames)
				for _, ip := range cert.IPAddresses {
					certIPs = append(certIPs, ip.String())
				}
				sort.Strings(certIPs)

				if reflect.DeepEqual(dnsNames, certDNSNames) && reflect.DeepEqual(ips, certIPs) {
					continue
				}
			}
		}
		var serviceKey *rsa.PrivateKey
		if !rotate {
			serviceKey = certs[etcdName].Key
		}
		logrus.Infof("[certificates] Generating %s certificate and key", etcdName)
		etcdCrt, etcdKey, err := GenerateSignedCertAndKey(caCrt, caKey, true, EtcdCertName, etcdAltNames, serviceKey, nil)
		if err != nil {
			return err
		}
		certs[etcdName] = ToCertObject(etcdName, "", "", etcdCrt, etcdKey, nil)
	}
	deleteUnusedCerts(ctx, certs, EtcdCertName, etcdHosts)
	return nil
}

func GenerateEtcdCSRs(ctx context.Context, certs map[string]CertificatePKI, bkeConfig v3.BhojpurKubernetesEngineConfig) error {
	kubernetesServiceIP, err := GetKubernetesServiceIP(bkeConfig.Services.KubeAPI.ServiceClusterIPRange)
	if err != nil {
		return fmt.Errorf("Failed to get Kubernetes Service IP: %v", err)
	}
	clusterDomain := bkeConfig.Services.Kubelet.ClusterDomain
	etcdHosts := hosts.NodesToHosts(bkeConfig.Nodes, etcdRole)
	etcdAltNames := GetAltNames(etcdHosts, clusterDomain, kubernetesServiceIP, []string{})
	for _, host := range etcdHosts {
		etcdName := GetCrtNameForHost(host, EtcdCertName)
		etcdCrt := certs[etcdName].Certificate
		etcdCsr := certs[etcdName].CSR
		if etcdCsr != nil {
			if reflect.DeepEqual(etcdAltNames.DNSNames, etcdCsr.DNSNames) &&
				DeepEqualIPsAltNames(etcdAltNames.IPs, etcdCsr.IPAddresses) {
				continue
			}
		}
		logrus.Infof("[certificates] Generating etcd-%s csr and key", host.InternalAddress)
		etcdCSR, etcdKey, err := GenerateCertSigningRequestAndKey(true, EtcdCertName, etcdAltNames, certs[etcdName].Key, nil)
		if err != nil {
			return err
		}
		certs[etcdName] = ToCertObject(etcdName, "", "", etcdCrt, etcdKey, etcdCSR)
	}
	return nil
}

func GenerateServiceTokenKey(ctx context.Context, certs map[string]CertificatePKI, bkeConfig v3.BhojpurKubernetesEngineConfig, configPath, configDir string, rotate bool) error {
	// generate service account token key
	privateAPIKey := certs[ServiceAccountTokenKeyName].Key
	caCrt := certs[CACertName].Certificate
	caKey := certs[CACertName].Key
	if caCrt == nil || caKey == nil {
		return fmt.Errorf("CA Certificate or Key is empty")
	}
	if certs[ServiceAccountTokenKeyName].Certificate != nil {
		return nil
	}
	// handle rotation on old clusters
	if certs[ServiceAccountTokenKeyName].Key == nil {
		privateAPIKey = certs[KubeAPICertName].Key
	}
	tokenCrt, tokenKey, err := GenerateSignedCertAndKey(caCrt, caKey, false, ServiceAccountTokenKeyName, nil, privateAPIKey, nil)
	if err != nil {
		return fmt.Errorf("Failed to generate private key for service account token: %v", err)
	}
	certs[ServiceAccountTokenKeyName] = ToCertObject(ServiceAccountTokenKeyName, ServiceAccountTokenKeyName, "", tokenCrt, tokenKey, nil)
	return nil
}

func GenerateBKECACerts(ctx context.Context, certs map[string]CertificatePKI, configPath, configDir string) error {
	if err := GenerateBKEMasterCACert(ctx, certs, configPath, configDir); err != nil {
		return err
	}
	return GenerateBKERequestHeaderCACert(ctx, certs, configPath, configDir)
}

func GenerateBKEMasterCACert(ctx context.Context, certs map[string]CertificatePKI, configPath, configDir string) error {
	// generate kubernetes CA certificate and key
	logrus.Info("[certificates] Generating CA kubernetes certificates")

	caCrt, caKey, err := GenerateCACertAndKey(CACertName, nil)
	if err != nil {
		return err
	}
	certs[CACertName] = ToCertObject(CACertName, "", "", caCrt, caKey, nil)
	return nil
}

func GenerateBKERequestHeaderCACert(ctx context.Context, certs map[string]CertificatePKI, configPath, configDir string) error {
	// generate request header client CA certificate and key
	logrus.Info("[certificates] Generating Kubernetes API server aggregation layer requestheader client CA certificates")
	requestHeaderCACrt, requestHeaderCAKey, err := GenerateCACertAndKey(RequestHeaderCACertName, nil)
	if err != nil {
		return err
	}
	certs[RequestHeaderCACertName] = ToCertObject(RequestHeaderCACertName, "", "", requestHeaderCACrt, requestHeaderCAKey, nil)
	return nil
}

func GenerateKubeletCertificate(ctx context.Context, certs map[string]CertificatePKI, bkeConfig v3.BhojpurKubernetesEngineConfig, configPath, configDir string, rotate bool) error {
	// generate kubelet certificate and key
	caCrt := certs[CACertName].Certificate
	caKey := certs[CACertName].Key
	if caCrt == nil || caKey == nil {
		return fmt.Errorf("CA Certificate or Key is empty")
	}
	log.Debugf(ctx, "[certificates] Generating Kubernetes Kubelet certificates")
	allHosts := hosts.NodesToHosts(bkeConfig.Nodes, "")
	for _, host := range allHosts {
		kubeletName := GetCrtNameForHost(host, KubeletCertName)
		kubeletCert := certs[kubeletName].Certificate
		if kubeletCert != nil && !rotate {
			continue
		}
		kubeletAltNames := GetIPHostAltnamesForHost(host)
		if kubeletCert != nil &&
			reflect.DeepEqual(kubeletAltNames.DNSNames, kubeletCert.DNSNames) &&
			DeepEqualIPsAltNames(kubeletAltNames.IPs, kubeletCert.IPAddresses) && !rotate {
			continue
		}
		var serviceKey *rsa.PrivateKey
		if !rotate {
			serviceKey = certs[kubeletName].Key
		}
		log.Debugf(ctx, "[certificates] Generating %s certificate and key", kubeletName)
		kubeletCrt, kubeletKey, err := GenerateSignedCertAndKey(caCrt, caKey, true, kubeletName, kubeletAltNames, serviceKey, nil)
		if err != nil {
			return err
		}
		certs[kubeletName] = ToCertObject(kubeletName, "", "", kubeletCrt, kubeletKey, nil)
	}
	deleteUnusedCerts(ctx, certs, KubeletCertName, allHosts)
	return nil
}

func GenerateKubeletCSR(ctx context.Context, certs map[string]CertificatePKI, bkeConfig v3.BhojpurKubernetesEngineConfig) error {
	allHosts := hosts.NodesToHosts(bkeConfig.Nodes, "")
	for _, host := range allHosts {
		kubeletName := GetCrtNameForHost(host, KubeletCertName)
		kubeletCert := certs[kubeletName].Certificate
		oldKubeletCSR := certs[kubeletName].CSR
		kubeletAltNames := GetIPHostAltnamesForHost(host)
		if oldKubeletCSR != nil &&
			reflect.DeepEqual(kubeletAltNames.DNSNames, oldKubeletCSR.DNSNames) &&
			DeepEqualIPsAltNames(kubeletAltNames.IPs, oldKubeletCSR.IPAddresses) {
			continue
		}
		logrus.Infof("[certificates] Generating %s Kubernetes Kubelet csr", kubeletName)
		kubeletCSR, kubeletKey, err := GenerateCertSigningRequestAndKey(true, kubeletName, kubeletAltNames, certs[kubeletName].Key, nil)
		if err != nil {
			return err
		}
		certs[kubeletName] = ToCertObject(kubeletName, "", "", kubeletCert, kubeletKey, kubeletCSR)
	}
	return nil
}

func GenerateBKEServicesCerts(ctx context.Context, certs map[string]CertificatePKI, bkeConfig v3.BhojpurKubernetesEngineConfig, configPath, configDir string, rotate bool) error {
	BKECerts := []GenFunc{
		GenerateKubeAPICertificate,
		GenerateServiceTokenKey,
		GenerateKubeControllerCertificate,
		GenerateKubeSchedulerCertificate,
		GenerateKubeProxyCertificate,
		GenerateKubeNodeCertificate,
		GenerateKubeAdminCertificate,
		GenerateAPIProxyClientCertificate,
		GenerateEtcdCertificates,
	}
	if IsKubeletGenerateServingCertificateEnabledinConfig(&bkeConfig) {
		BKECerts = append(BKECerts, GenerateKubeletCertificate)
	} else {
		//Clean up kubelet certs when GenerateServingCertificate is disabled
		logrus.Info("[certificates] GenerateServingCertificate is disabled, checking if there are unused kubelet certificates")
		for k := range certs {
			if strings.HasPrefix(k, KubeletCertName) {
				logrus.Infof("[certificates] Deleting unused kubelet certificate: %s", k)
				delete(certs, k)
			}
		}
	}
	for _, gen := range BKECerts {
		if err := gen(ctx, certs, bkeConfig, configPath, configDir, rotate); err != nil {
			return err
		}
	}
	if len(bkeConfig.Services.Etcd.ExternalURLs) > 0 {
		return GenerateExternalEtcdCertificates(ctx, certs, bkeConfig, configPath, configDir, false)
	}
	return nil
}

func GenerateBKEServicesCSRs(ctx context.Context, certs map[string]CertificatePKI, bkeConfig v3.BhojpurKubernetesEngineConfig) error {
	BKECerts := []CSRFunc{
		GenerateKubeAPICSR,
		GenerateKubeControllerCSR,
		GenerateKubeSchedulerCSR,
		GenerateKubeProxyCSR,
		GenerateKubeNodeCSR,
		GenerateKubeAdminCSR,
		GenerateAPIProxyClientCSR,
		GenerateEtcdCSRs,
	}
	if IsKubeletGenerateServingCertificateEnabledinConfig(&bkeConfig) {
		BKECerts = append(BKECerts, GenerateKubeletCSR)
	}
	for _, csr := range BKECerts {
		if err := csr(ctx, certs, bkeConfig); err != nil {
			return err
		}
	}
	return nil
}

func deleteUnusedCerts(ctx context.Context, certs map[string]CertificatePKI, certName string, hostList []*hosts.Host) {
	hostAddresses := hosts.GetInternalAddressForHosts(hostList)
	logrus.Tracef("Checking and deleting unused certificates with prefix [%s] for the following [%d] node(s): %s", certName, len(hostAddresses), strings.Join(hostAddresses, ","))
	unusedCerts := make(map[string]bool)
	for k := range certs {
		if strings.HasPrefix(k, certName) {
			unusedCerts[k] = true
		}
	}
	for _, host := range hostList {
		Name := GetCrtNameForHost(host, certName)
		delete(unusedCerts, Name)
	}
	for k := range unusedCerts {
		logrus.Infof("[certificates] Deleting unused certificate: %s", k)
		delete(certs, k)
	}
}

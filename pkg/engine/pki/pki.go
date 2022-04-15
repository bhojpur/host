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
	"crypto/x509"
	"fmt"
	"net"
	"path"

	"github.com/bhojpur/host/pkg/cluster/log"
	"github.com/bhojpur/host/pkg/engine/docker"
	"github.com/bhojpur/host/pkg/engine/hosts"
	v3 "github.com/bhojpur/host/pkg/engine/types"
	"github.com/bhojpur/host/pkg/engine/util"
	"github.com/docker/docker/api/types/container"
)

type CertificatePKI struct {
	Certificate    *x509.Certificate        `json:"-"`
	Key            *rsa.PrivateKey          `json:"-"`
	CSR            *x509.CertificateRequest `json:"-"`
	CertificatePEM string                   `json:"certificatePEM"`
	KeyPEM         string                   `json:"keyPEM"`
	CSRPEM         string                   `json:"-"`
	Config         string                   `json:"config"`
	Name           string                   `json:"name"`
	CommonName     string                   `json:"commonName"`
	OUName         string                   `json:"ouName"`
	EnvName        string                   `json:"envName"`
	Path           string                   `json:"path"`
	KeyEnvName     string                   `json:"keyEnvName"`
	KeyPath        string                   `json:"keyPath"`
	ConfigEnvName  string                   `json:"configEnvName"`
	ConfigPath     string                   `json:"configPath"`
}

type GenFunc func(context.Context, map[string]CertificatePKI, v3.BhojpurKubernetesEngineConfig, string, string, bool) error
type CSRFunc func(context.Context, map[string]CertificatePKI, v3.BhojpurKubernetesEngineConfig) error

const (
	etcdRole            = "etcd"
	controlRole         = "controlplane"
	workerRole          = "worker"
	BundleCertContainer = "bke-bundle-cert"
)

func GenerateBKECerts(ctx context.Context, bkeConfig v3.BhojpurKubernetesEngineConfig, configPath, configDir string) (map[string]CertificatePKI, error) {
	certs := make(map[string]CertificatePKI)
	// generate BKE CA certificates
	if err := GenerateBKECACerts(ctx, certs, configPath, configDir); err != nil {
		return certs, err
	}
	// Generating certificates for Kubernetes components
	if err := GenerateBKEServicesCerts(ctx, certs, bkeConfig, configPath, configDir, false); err != nil {
		return certs, err
	}
	return certs, nil
}

func GenerateBKENodeCerts(ctx context.Context, bkeConfig v3.BhojpurKubernetesEngineConfig, nodeAddress string, certBundle map[string]CertificatePKI) map[string]CertificatePKI {
	crtMap := make(map[string]CertificatePKI)
	crtKeys := []string{}
	for _, node := range bkeConfig.Nodes {
		if node.Address == nodeAddress {
			for _, role := range node.Role {
				keys := getCertKeys(bkeConfig.Nodes, role, &bkeConfig)
				crtKeys = append(crtKeys, keys...)
			}
			break
		}
	}
	for _, key := range crtKeys {
		crtMap[key] = certBundle[key]
	}
	return crtMap
}

func RegenerateEtcdCertificate(
	ctx context.Context,
	crtMap map[string]CertificatePKI,
	etcdHost *hosts.Host,
	etcdHosts []*hosts.Host,
	clusterDomain string,
	KubernetesServiceIP []net.IP) (map[string]CertificatePKI, error) {

	etcdName := GetCrtNameForHost(etcdHost, EtcdCertName)
	log.Infof(ctx, "[certificates] Regenerating new %s certificate and key", etcdName)
	caCrt := crtMap[CACertName].Certificate
	caKey := crtMap[CACertName].Key
	etcdAltNames := GetAltNames(etcdHosts, clusterDomain, KubernetesServiceIP, []string{})

	etcdCrt, etcdKey, err := GenerateSignedCertAndKey(caCrt, caKey, true, EtcdCertName, etcdAltNames, nil, nil)
	if err != nil {
		return nil, err
	}
	crtMap[etcdName] = ToCertObject(etcdName, "", "", etcdCrt, etcdKey, nil)
	log.Infof(ctx, "[certificates] Successfully generated new %s certificate and key", etcdName)
	return crtMap, nil
}

func SaveBackupBundleOnHost(ctx context.Context, host *hosts.Host, alpineSystemImage, etcdSnapshotPath string, prsMap map[string]v3.PrivateRegistry, k8sVersion string) error {
	imageCfg := &container.Config{
		Cmd: []string{
			"sh",
			"-c",
			fmt.Sprintf("if [ -d %s ] && [ \"$(ls -A %s)\" ]; then tar czvf %s %s;fi", TempCertPath, TempCertPath, BundleCertPath, TempCertPath),
		},
		Image: alpineSystemImage,
	}
	hostCfg := &container.HostConfig{
		Privileged: true,
	}

	binds := []string{
		fmt.Sprintf("%s:/etc/kubernetes:z", path.Join(host.PrefixPath, "/etc/kubernetes")),
		fmt.Sprintf("%s:/backup:z", etcdSnapshotPath),
	}

	matchedRange, err := util.SemVerMatchRange(k8sVersion, util.SemVerK8sVersion122OrHigher)
	if err != nil {
		return err
	}

	if matchedRange {
		binds = util.RemoveZFromBinds(binds)
	}

	hostCfg.Binds = binds

	if err := docker.DoRunContainer(ctx, host.DClient, imageCfg, hostCfg, BundleCertContainer, host.Address, "certificates", prsMap); err != nil {
		return err
	}
	status, err := docker.WaitForContainer(ctx, host.DClient, host.Address, BundleCertContainer)
	if err != nil {
		return err
	}
	if status != 0 {
		return fmt.Errorf("Failed to run certificate bundle compress, exit status is: %d", status)
	}
	log.Infof(ctx, "[certificates] successfully saved certificate bundle [%s/pki.bundle.tar.gz] on host [%s]", etcdSnapshotPath, host.Address)
	return docker.RemoveContainer(ctx, host.DClient, host.Address, BundleCertContainer)
}

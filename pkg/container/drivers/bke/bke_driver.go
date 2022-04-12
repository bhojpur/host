package bke

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
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	v3 "github.com/bhojpur/host/pkg/apis/management.bhojpur.net/v3"
	"github.com/bhojpur/host/pkg/container/drivers/bke/bkecerts"
	"github.com/bhojpur/host/pkg/container/drivers/util"
	"github.com/bhojpur/host/pkg/container/log"
	"github.com/bhojpur/host/pkg/container/types"
	"github.com/bhojpur/host/pkg/core/types/slice"
	"github.com/bhojpur/host/pkg/engine/cluster"
	"github.com/bhojpur/host/pkg/engine/cmd"
	"github.com/bhojpur/host/pkg/engine/hosts"
	"github.com/bhojpur/host/pkg/engine/pki"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/transport"
)

const (
	kubeConfigFile   = "kube_config_cluster.yml"
	bhojpurPath      = "./management-state/bke/"
	clusterStateFile = "cluster.bkestate"
)

// Driver is the struct of BKE driver

type WrapTransportFactory func(config *v3.BhojpurKubernetesEngineConfig) transport.WrapperFunc

type Driver struct {
	DockerDialer         hosts.DialerFactory
	LocalDialer          hosts.DialerFactory
	DataStore            Store
	WrapTransportFactory WrapTransportFactory
	driverCapabilities   types.Capabilities

	types.UnimplementedVersionAccess
	types.UnimplementedClusterSizeAccess
}

type Store interface {
	// Add methods here to get BKE driver specific data
	GetAddonTemplates(k8sVersion string) (map[string]interface{}, error)
	// GetServiceOptions returns the combined result of service options:
	// - `k8s-service-options` corresponds to linux options
	// - `k8s-windows-service-options` corresponds to windows options
	GetServiceOptions(k8sVersion string) (map[string]*v3.KubernetesServicesOptions, error)
}

func NewDriver() types.Driver {
	d := &Driver{
		driverCapabilities: types.Capabilities{
			Capabilities: make(map[int64]bool),
		},
	}

	d.driverCapabilities.AddCapability(types.GetVersionCapability)
	d.driverCapabilities.AddCapability(types.SetVersionCapability)

	d.driverCapabilities.AddCapability(types.GetClusterSizeCapability)
	d.driverCapabilities.AddCapability(types.EtcdBackupCapability)

	return d
}

func (d *Driver) wrapTransport(config *v3.BhojpurKubernetesEngineConfig) transport.WrapperFunc {
	if d.WrapTransportFactory == nil {
		return nil
	}

	return func(rt http.RoundTripper) http.RoundTripper {
		fn := d.WrapTransportFactory(config)
		if fn == nil {
			return rt
		}
		return fn(rt)
	}

}

func (d *Driver) GetCapabilities(ctx context.Context) (*types.Capabilities, error) {
	return &d.driverCapabilities, nil
}

func (d *Driver) GetK8SCapabilities(ctx context.Context, _ *types.DriverOptions) (*types.K8SCapabilities, error) {
	return &types.K8SCapabilities{}, nil
}

// GetDriverCreateOptions returns create flags for BKE driver
func (d *Driver) GetDriverCreateOptions(ctx context.Context) (*types.DriverFlags, error) {
	driverFlag := types.DriverFlags{
		Options: make(map[string]*types.Flag),
	}
	driverFlag.Options["config-file-path"] = &types.Flag{
		Type:  types.StringType,
		Usage: "the path to the config file",
	}
	return &driverFlag, nil
}

// GetDriverUpdateOptions returns update flags for BKE driver
func (d *Driver) GetDriverUpdateOptions(ctx context.Context) (*types.DriverFlags, error) {
	driverFlag := types.DriverFlags{
		Options: make(map[string]*types.Flag),
	}
	driverFlag.Options["config-file-path"] = &types.Flag{
		Type:  types.StringType,
		Usage: "the path to the config file",
	}
	return &driverFlag, nil
}

// SetDriverOptions sets the drivers options to BKE driver
func getYAML(driverOptions *types.DriverOptions) (string, error) {
	// first look up the file path then look up raw bkeConfig
	if path, ok := driverOptions.StringOptions["config-file-path"]; ok {
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return "", err
		}
		return string(data), nil
	}
	return driverOptions.StringOptions["bkeConfig"], nil
}

// Create creates the BKE cluster
func (d *Driver) Create(ctx context.Context, opts *types.DriverOptions, info *types.ClusterInfo) (*types.ClusterInfo, error) {
	yaml, err := getYAML(opts)
	if err != nil {
		return nil, err
	}

	bkeConfig, err := util.ConvertToBkeConfig(yaml)
	if err != nil {
		return nil, err
	}

	stateDir, err := d.restore(info)
	if err != nil {
		return nil, err
	}
	defer d.cleanup(stateDir)

	data, err := getData(d.DataStore, bkeConfig.Version)
	if err != nil {
		return nil, err
	}

	certsStr := ""
	dialers, externalFlags := d.getFlags(bkeConfig, stateDir)
	APIURL, caCrt, clientCert, clientKey, certs, err := clusterUp(ctx, &bkeConfig, dialers, externalFlags, data)
	if len(certs) > 0 {
		certsStr, err = bkecerts.ToString(certs)
	}
	if err != nil && certsStr == "" {
		return d.save(&types.ClusterInfo{
			Metadata: map[string]string{
				"Config": yaml,
			},
		}, stateDir), err
	}

	return d.save(&types.ClusterInfo{
		Metadata: map[string]string{
			"Endpoint":   APIURL,
			"RootCA":     base64.StdEncoding.EncodeToString([]byte(caCrt)),
			"ClientCert": base64.StdEncoding.EncodeToString([]byte(clientCert)),
			"ClientKey":  base64.StdEncoding.EncodeToString([]byte(clientKey)),
			"Config":     yaml,
			"Certs":      certsStr,
		},
	}, stateDir), err
}

func getData(s Store, k8sVersion string) (map[string]interface{}, error) {
	data, err := s.GetAddonTemplates(k8sVersion)
	if err != nil {
		return data, err
	}
	serviceOptions, err := s.GetServiceOptions(k8sVersion)
	if err != nil {
		return data, err
	}
	for key, opts := range serviceOptions {
		if opts != nil {
			data[key] = opts
		}
	}
	return data, nil
}

// Update updates the BKE cluster
func (d *Driver) Update(ctx context.Context, clusterInfo *types.ClusterInfo, opts *types.DriverOptions) (*types.ClusterInfo, error) {
	yaml, err := getYAML(opts)
	if err != nil {
		return nil, err
	}

	bkeConfig, err := util.ConvertToBkeConfig(yaml)
	if err != nil {
		return nil, err
	}

	stateDir, err := d.restore(clusterInfo)
	if err != nil {
		return nil, err
	}
	defer d.cleanup(stateDir)

	data, err := getData(d.DataStore, bkeConfig.Version)
	if err != nil {
		return nil, err
	}

	dialers, externalFlags := d.getFlags(bkeConfig, stateDir)
	/* removed temporarily
	if err := cmd.ClusterInit(ctx, &bkeConfig, dialers, externalFlags); err != nil {
		return nil, err
	}
	*/
	APIURL, caCrt, clientCert, clientKey, certs, err := cmd.ClusterUp(ctx, dialers, externalFlags, data)
	if err != nil {
		return d.save(&types.ClusterInfo{
			Metadata: map[string]string{
				"Config": yaml,
			},
		}, stateDir), err
	}
	metadata, err := updateMetadata(APIURL, caCrt, clientCert, clientKey, yaml, certs)

	clusterInfo.Metadata = metadata
	return d.save(clusterInfo, stateDir), err
}

func (d *Driver) getClientset(info *types.ClusterInfo) (*kubernetes.Clientset, error) {
	yaml := info.Metadata["Config"]

	bkeConfig, err := util.ConvertToBkeConfig(yaml)
	if err != nil {
		return nil, err
	}

	info.Endpoint = info.Metadata["Endpoint"]
	info.ClientCertificate = info.Metadata["ClientCert"]
	info.ClientKey = info.Metadata["ClientKey"]
	info.RootCaCertificate = info.Metadata["RootCA"]

	certBytes, err := base64.StdEncoding.DecodeString(info.ClientCertificate)
	if err != nil {
		return nil, err
	}
	keyBytes, err := base64.StdEncoding.DecodeString(info.ClientKey)
	if err != nil {
		return nil, err
	}
	rootBytes, err := base64.StdEncoding.DecodeString(info.RootCaCertificate)
	if err != nil {
		return nil, err
	}

	host := info.Endpoint
	if !strings.HasPrefix(host, "https://") {
		host = fmt.Sprintf("https://%s", host)
	}
	config := &rest.Config{
		Host: host,
		TLSClientConfig: rest.TLSClientConfig{
			CAData:   rootBytes,
			CertData: certBytes,
			KeyData:  keyBytes,
		},
		WrapTransport: d.WrapTransportFactory(&bkeConfig),
	}

	return kubernetes.NewForConfig(config)
}

// PostCheck does post action
func (d *Driver) PostCheck(ctx context.Context, info *types.ClusterInfo) (*types.ClusterInfo, error) {
	clientset, err := d.getClientset(info)
	if err != nil {
		return nil, err
	}

	var lastErr error
	for i := 0; i < 3; i++ {
		serverVersion, err := clientset.DiscoveryClient.ServerVersion()
		if err != nil {
			lastErr = fmt.Errorf("failed to get Kubernetes server version: %v", err)
			time.Sleep(2 * time.Second)
			continue
		}

		token, err := util.GenerateServiceAccountToken(clientset)
		if err != nil {
			lastErr = err
			time.Sleep(2 * time.Second)
			continue
		}

		info.Version = serverVersion.GitVersion
		info.ServiceAccountToken = token

		info.NodeCount, err = nodeCount(info)
		if err != nil {
			lastErr = err
			time.Sleep(2 * time.Second)
			continue
		}

		return info, err
	}

	return nil, lastErr
}

func (d *Driver) GetVersion(ctx context.Context, info *types.ClusterInfo) (*types.KubernetesVersion, error) {
	clientset, err := d.getClientset(info)
	if err != nil {
		return nil, fmt.Errorf("failed to create clientset: %v", err)
	}

	serviceVersion, err := clientset.DiscoveryClient.ServerVersion()
	if err != nil {
		return nil, fmt.Errorf("failed to get server version: %v", err)
	}

	return &types.KubernetesVersion{Version: serviceVersion.String()}, nil
}

func (d *Driver) SetVersion(ctx context.Context, info *types.ClusterInfo, version *types.KubernetesVersion) error {
	config, err := util.ConvertToBkeConfig(info.Metadata["Config"])

	if err != nil {
		return err
	}

	config.Version = version.Version

	stateDir, err := d.restore(info)
	if err != nil {
		return err
	}
	defer d.cleanup(stateDir)
	dialers, externalFlags := d.getFlags(config, stateDir)
	/* removed temporarily
	if err := cmd.ClusterInit(ctx, &config, dialers, externalFlags); err != nil {
		return err
	}*/

	data, err := getData(d.DataStore, config.Version)
	if err != nil {
		return err
	}

	_, _, _, _, _, err = cmd.ClusterUp(ctx, dialers, externalFlags, data)

	if err != nil {
		return err
	}

	d.save(info, stateDir)

	return nil
}

func (d *Driver) GetClusterSize(ctx context.Context, info *types.ClusterInfo) (*types.NodeCount, error) {
	clientset, err := d.getClientset(info)
	if err != nil {
		return nil, fmt.Errorf("failed to create clientset: %v", err)
	}

	nodeList, err := clientset.CoreV1().Nodes().List(ctx, v1.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get server version: %v", err)
	}

	return &types.NodeCount{Count: int64(len(nodeList.Items))}, nil
}

func nodeCount(info *types.ClusterInfo) (int64, error) {
	yaml, ok := info.Metadata["Config"]
	if !ok {
		return 0, nil
	}

	bkeConfig, err := util.ConvertToBkeConfig(yaml)
	if err != nil {
		return 0, err
	}

	count := int64(0)
	for _, node := range bkeConfig.Nodes {
		if slice.ContainsString(node.Role, "worker") {
			count++
		}
	}

	return count, nil
}

// Remove removes the cluster
func (d *Driver) Remove(ctx context.Context, clusterInfo *types.ClusterInfo) error {
	bkeConfig, err := util.ConvertToBkeConfig(clusterInfo.Metadata["Config"])
	if err != nil {
		return err
	}
	stateDir, _ := d.restore(clusterInfo)
	defer d.save(nil, stateDir)
	dialers, externalFlags := d.getFlags(bkeConfig, stateDir)
	return cmd.ClusterRemove(ctx, &bkeConfig, dialers, externalFlags)
	return nil
}

func (d *Driver) RemoveLegacyServiceAccount(ctx context.Context, info *types.ClusterInfo) error {
	clientset, err := d.getClientset(info)
	if err != nil {
		return err
	}

	return util.DeleteLegacyServiceAccountAndRoleBinding(clientset)
}

func (d *Driver) restore(info *types.ClusterInfo) (string, error) {
	os.MkdirAll(bhojpurPath, 0700)
	dir, err := ioutil.TempDir(bhojpurPath, "bke-")
	if err != nil {
		return "", err
	}

	if info != nil {
		state := info.Metadata["state"]
		if state != "" {
			ioutil.WriteFile(kubeConfig(dir), []byte(state), 0600)
		}
		fullState := info.Metadata["fullState"]
		if fullState != "" {
			ioutil.WriteFile(clusterState(dir), []byte(fullState), 0600)
		}
	}

	return filepath.Join(dir, "cluster.yml"), nil
}

func (d *Driver) save(info *types.ClusterInfo, stateDir string) *types.ClusterInfo {
	if info != nil {
		b, err := ioutil.ReadFile(kubeConfig(stateDir))
		if err == nil {
			if info.Metadata == nil {
				info.Metadata = map[string]string{}
			}
			info.Metadata["state"] = string(b)
		}
		s, err := ioutil.ReadFile(clusterState(stateDir))
		if err == nil {
			info.Metadata["fullState"] = string(s)
		}
	}

	d.cleanup(stateDir)

	return info
}

func (d *Driver) cleanup(stateDir string) {
	if strings.HasSuffix(stateDir, "/cluster.yml") && !strings.Contains(stateDir, "..") {
		os.Remove(stateDir)
		os.Remove(kubeConfig(stateDir))
		os.Remove(filepath.Dir(stateDir))
	}
}

func (d *Driver) getFlags(bkeConfig v3.BhojpurKubernetesEngineConfig, stateDir string) (hosts.DialersOptions, cluster.ExternalFlags) {
	dialers := hosts.GetDialerOptions(d.DockerDialer, d.LocalDialer, d.wrapTransport(&bkeConfig))
	externalFlags := cluster.GetExternalFlags(false, false, false, true, stateDir, "")
	return dialers, externalFlags
}

func kubeConfig(stateDir string) string {
	if strings.HasSuffix(stateDir, "/cluster.yml") {
		return filepath.Join(filepath.Dir(stateDir), kubeConfigFile)
	}
	return filepath.Join(stateDir, kubeConfigFile)
}

func clusterState(stateDir string) string {
	if strings.HasSuffix(stateDir, "/cluster.yml") {
		return filepath.Join(filepath.Dir(stateDir), clusterStateFile)
	}
	return filepath.Join(stateDir, clusterStateFile)
}

func clusterUp(
	ctx context.Context,
	bkeConfig *v3.BhojpurKubernetesEngineConfig,
	dialers hosts.DialersOptions,
	externalFlags cluster.ExternalFlags, data map[string]interface{}) (string, string, string, string, map[string]pki.CertificatePKI, error) {

	if err := cmd.ClusterInit(ctx, bkeConfig, dialers, externalFlags); err != nil {
		log.Warnf(ctx, "%v", err)
	}
	APIURL, caCrt, clientCert, clientKey, certs, err := cmd.ClusterUp(ctx, dialers, externalFlags, data)
	if err != nil {
		log.Warnf(ctx, "%v", err)
	}
	return APIURL, caCrt, clientCert, clientKey, certs, err
}

func (d *Driver) ETCDSave(ctx context.Context, clusterInfo *types.ClusterInfo, opts *types.DriverOptions, snapshotName string) error {
	bkeConfig, err := util.ConvertToBkeConfig(clusterInfo.Metadata["Config"])
	if err != nil {
		return err
	}
	stateDir, err := d.restore(clusterInfo)
	if err != nil {
		return err
	}
	defer d.cleanup(stateDir)

	dialers, externalFlags := d.getFlags(bkeConfig, stateDir)

	return cmd.SnapshotSaveEtcdHosts(ctx, &bkeConfig, dialers, externalFlags, snapshotName)

	return nil
}

func (d *Driver) ETCDRestore(ctx context.Context, clusterInfo *types.ClusterInfo, opts *types.DriverOptions, snapshotName string) (*types.ClusterInfo, error) {
	yaml, err := getYAML(opts)
	if err != nil {
		return nil, err
	}

	bkeConfig, err := util.ConvertToBkeConfig(yaml)
	if err != nil {
		return nil, err
	}
	stateDir, err := d.restore(clusterInfo)
	if err != nil {
		return nil, err
	}
	defer d.cleanup(stateDir)

	data, err := getData(d.DataStore, bkeConfig.Version)
	if err != nil {
		return nil, err
	}

	dialers, externalFlags := d.getFlags(bkeConfig, stateDir)
	APIURL, caCrt, clientCert, clientKey, certs, err := cmd.RestoreEtcdSnapshot(ctx, &bkeConfig, dialers, externalFlags, data, snapshotName)

	if err != nil {
		return d.save(&types.ClusterInfo{
			Metadata: map[string]string{
				"Config": yaml,
			},
		}, stateDir), err
	}

	metadata, err := updateMetadata(APIURL, caCrt, clientCert, clientKey, yaml, certs)
	clusterInfo.Metadata = metadata
	return d.save(clusterInfo, stateDir), err

	return nil, nil
}

func (d *Driver) ETCDRemoveSnapshot(ctx context.Context, clusterInfo *types.ClusterInfo, opts *types.DriverOptions, snapshotName string) error {
	bkeConfig, err := util.ConvertToBkeConfig(clusterInfo.Metadata["Config"])
	if err != nil {
		return err
	}
	stateDir, err := d.restore(clusterInfo)
	if err != nil {
		return err
	}
	defer d.cleanup(stateDir)

	dialers, externalFlags := d.getFlags(bkeConfig, stateDir)

	return cmd.SnapshotRemoveFromEtcdHosts(ctx, &bkeConfig, dialers, externalFlags, snapshotName)

	return nil
}

func updateMetadata(APIURL, caCrt, clientCert, clientKey, yaml string, certs map[string]pki.CertificatePKI) (map[string]string, error) {
	m := map[string]string{}
	certStr := ""
	certStr, err := bkecerts.ToString(certs)
	if err != nil {
		m["Config"] = yaml
		return m, err
	}
	m["Endpoint"] = APIURL
	m["RootCA"] = base64.StdEncoding.EncodeToString([]byte(caCrt))
	m["ClientCert"] = base64.StdEncoding.EncodeToString([]byte(clientCert))
	m["ClientKey"] = base64.StdEncoding.EncodeToString([]byte(clientKey))
	m["Config"] = yaml
	m["Certs"] = certStr
	return m, nil
}

package cluster

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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"k8s.io/client-go/transport"

	"github.com/bhojpur/host/pkg/container/log"
	"github.com/bhojpur/host/pkg/engine/hosts"
	"github.com/bhojpur/host/pkg/engine/k8s"
	"github.com/bhojpur/host/pkg/engine/pki"
	"github.com/bhojpur/host/pkg/engine/services"
	v3 "github.com/bhojpur/host/pkg/engine/types"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	v1 "k8s.io/api/core/v1"
)

const (
	stateFileExt = ".bkestate"
	certDirExt   = "_certs"
)

type FullState struct {
	DesiredState State `json:"desiredState,omitempty"`
	CurrentState State `json:"currentState,omitempty"`
}

type State struct {
	BhojpurKubernetesEngineConfig *v3.BhojpurKubernetesEngineConfig `json:"bkeConfig,omitempty"`
	CertificatesBundle            map[string]pki.CertificatePKI     `json:"certificatesBundle,omitempty"`
	EncryptionConfig              string                            `json:"encryptionConfig,omitempty"`
}

func (c *Cluster) UpdateClusterCurrentState(ctx context.Context, fullState *FullState) error {
	fullState.CurrentState.BhojpurKubernetesEngineConfig = c.BhojpurKubernetesEngineConfig.DeepCopy()
	fullState.CurrentState.CertificatesBundle = c.Certificates
	fullState.CurrentState.EncryptionConfig = c.EncryptionConfig.EncryptionProviderFile
	return fullState.WriteStateFile(ctx, c.StateFilePath)
}

func (c *Cluster) GetClusterState(ctx context.Context, fullState *FullState) (*Cluster, error) {
	var err error
	if fullState.CurrentState.BhojpurKubernetesEngineConfig == nil {
		return nil, nil
	}

	// resetup external flags
	flags := GetExternalFlags(false, false, false, false, c.ConfigDir, c.ConfigPath)
	currentCluster, err := InitClusterObject(ctx, fullState.CurrentState.BhojpurKubernetesEngineConfig, flags, fullState.CurrentState.EncryptionConfig)
	if err != nil {
		return nil, err
	}
	currentCluster.Certificates = fullState.CurrentState.CertificatesBundle
	currentCluster.EncryptionConfig.EncryptionProviderFile = fullState.CurrentState.EncryptionConfig
	// resetup dialers
	dialerOptions := hosts.GetDialerOptions(c.DockerDialerFactory, c.LocalConnDialerFactory, c.K8sWrapTransport)
	if err := currentCluster.SetupDialers(ctx, dialerOptions); err != nil {
		return nil, err
	}
	return currentCluster, nil
}

func (c *Cluster) GetStateFileFromConfigMap(ctx context.Context) (string, error) {
	kubeletImage := c.Services.Kubelet.Image
	for _, host := range c.ControlPlaneHosts {
		stateFile, err := services.RunGetStateFileFromConfigMap(ctx, host, c.PrivateRegistriesMap, kubeletImage, c.Version)
		if err != nil || stateFile == "" {
			logrus.Infof("Could not get ConfigMap with cluster state from host [%s]", host.Address)
			continue
		}
		return stateFile, nil
	}
	return "", fmt.Errorf("Unable to get ConfigMap with cluster state from any Control Plane host")
}

func SaveFullStateToKubernetes(ctx context.Context, kubeCluster *Cluster, fullState *FullState) error {
	k8sClient, err := k8s.NewClient(kubeCluster.LocalKubeConfigPath, kubeCluster.K8sWrapTransport)
	if err != nil {
		return fmt.Errorf("Failed to create Kubernetes Client: %v", err)
	}
	log.Infof(ctx, "[state] Saving full cluster state to Kubernetes")
	stateFile, err := json.Marshal(*fullState)
	if err != nil {
		return err
	}
	timeout := make(chan bool, 1)
	go func() {
		for {
			_, err := k8s.UpdateConfigMap(k8sClient, stateFile, FullStateConfigMapName)
			if err != nil {
				time.Sleep(time.Second * 5)
				continue
			}
			log.Infof(ctx, "[state] Successfully Saved full cluster state to Kubernetes ConfigMap: %s", FullStateConfigMapName)
			timeout <- true
			break
		}
	}()
	select {
	case <-timeout:
		return nil
	case <-time.After(time.Second * UpdateStateTimeout):
		return fmt.Errorf("[state] Timeout waiting for kubernetes to be ready")
	}
}

func GetStateFromKubernetes(ctx context.Context, kubeCluster *Cluster) (*Cluster, error) {
	log.Infof(ctx, "[state] Fetching cluster state from Kubernetes")
	k8sClient, err := k8s.NewClient(kubeCluster.LocalKubeConfigPath, kubeCluster.K8sWrapTransport)
	if err != nil {
		return nil, fmt.Errorf("Failed to create Kubernetes Client: %v", err)
	}
	var cfgMap *v1.ConfigMap
	var currentCluster Cluster
	timeout := make(chan bool, 1)
	go func() {
		for {
			cfgMap, err = k8s.GetConfigMap(k8sClient, StateConfigMapName)
			if err != nil {
				time.Sleep(time.Second * 5)
				continue
			}
			log.Infof(ctx, "[state] Successfully Fetched cluster state to Kubernetes ConfigMap: %s", StateConfigMapName)
			timeout <- true
			break
		}
	}()
	select {
	case <-timeout:
		clusterData := cfgMap.Data[StateConfigMapName]
		err := yaml.Unmarshal([]byte(clusterData), &currentCluster)
		if err != nil {
			return nil, fmt.Errorf("Failed to unmarshal cluster data")
		}
		return &currentCluster, nil
	case <-time.After(time.Second * GetStateTimeout):
		log.Infof(ctx, "Timed out waiting for kubernetes cluster to get state")
		return nil, fmt.Errorf("Timeout waiting for kubernetes cluster to get state")
	}
}

func GetK8sVersion(localConfigPath string, k8sWrapTransport transport.WrapperFunc) (string, error) {
	logrus.Debugf("[version] Using %s to connect to Kubernetes cluster..", localConfigPath)
	k8sClient, err := k8s.NewClient(localConfigPath, k8sWrapTransport)
	if err != nil {
		return "", fmt.Errorf("Failed to create Kubernetes Client: %v", err)
	}
	discoveryClient := k8sClient.DiscoveryClient
	logrus.Debugf("[version] Getting Kubernetes server version..")
	serverVersion, err := discoveryClient.ServerVersion()
	if err != nil {
		return "", fmt.Errorf("Failed to get Kubernetes server version: %v", err)
	}
	return fmt.Sprintf("%#v", *serverVersion), nil
}

func RebuildState(ctx context.Context, kubeCluster *Cluster, oldState *FullState, flags ExternalFlags) (*FullState, error) {
	bkeConfig := &kubeCluster.BhojpurKubernetesEngineConfig
	newState := &FullState{
		DesiredState: State{
			BhojpurKubernetesEngineConfig: bkeConfig.DeepCopy(),
		},
	}

	if flags.CustomCerts {
		certBundle, err := pki.ReadCertsAndKeysFromDir(flags.CertificateDir)
		if err != nil {
			return nil, fmt.Errorf("Failed to read certificates from dir [%s]: %v", flags.CertificateDir, err)
		}
		// make sure all custom certs are included
		if err := pki.ValidateBundleContent(bkeConfig, certBundle, flags.ClusterFilePath, flags.ConfigDir); err != nil {
			return nil, fmt.Errorf("Failed to validates certificates from dir [%s]: %v", flags.CertificateDir, err)
		}
		newState.DesiredState.CertificatesBundle = certBundle
		newState.CurrentState = oldState.CurrentState

		err = updateEncryptionConfig(kubeCluster, oldState, newState)
		if err != nil {
			return nil, err
		}
		return newState, nil
	}

	// Rebuilding the certificates of the desired state
	if oldState.DesiredState.CertificatesBundle == nil { // this is a fresh cluster
		if err := buildFreshState(ctx, kubeCluster, newState); err != nil {
			return nil, err
		}
	} else { // This is an existing cluster with an old DesiredState
		if err := rebuildExistingState(ctx, kubeCluster, oldState, newState, flags); err != nil {
			return nil, err
		}
	}
	newState.CurrentState = oldState.CurrentState
	return newState, nil
}

func (s *FullState) WriteStateFile(ctx context.Context, statePath string) error {
	stateFile, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return fmt.Errorf("Failed to Marshal state object: %v", err)
	}
	logrus.Tracef("Writing state file: %s", stateFile)
	if err := ioutil.WriteFile(statePath, stateFile, 0600); err != nil {
		return fmt.Errorf("Failed to write state file: %v", err)
	}
	log.Infof(ctx, "Successfully Deployed state file at [%s]", statePath)
	return nil
}

func GetStateFilePath(configPath, configDir string) string {
	if configPath == "" {
		configPath = pki.ClusterConfig
	}
	baseDir := filepath.Dir(configPath)
	if len(configDir) > 0 {
		baseDir = filepath.Dir(configDir)
	}
	fileName := filepath.Base(configPath)
	baseDir += "/"
	fullPath := fmt.Sprintf("%s%s", baseDir, fileName)
	trimmedName := strings.TrimSuffix(fullPath, filepath.Ext(fullPath))
	return trimmedName + stateFileExt
}

func GetCertificateDirPath(configPath, configDir string) string {
	if configPath == "" {
		configPath = pki.ClusterConfig
	}
	baseDir := filepath.Dir(configPath)
	if len(configDir) > 0 {
		baseDir = filepath.Dir(configDir)
	}
	fileName := filepath.Base(configPath)
	baseDir += "/"
	fullPath := fmt.Sprintf("%s%s", baseDir, fileName)
	trimmedName := strings.TrimSuffix(fullPath, filepath.Ext(fullPath))
	return trimmedName + certDirExt
}

func StringToFullState(ctx context.Context, stateFileContent string) (*FullState, error) {
	bkeFullState := &FullState{}
	logrus.Tracef("stateFileContent: %s", stateFileContent)
	if err := json.Unmarshal([]byte(stateFileContent), bkeFullState); err != nil {
		return bkeFullState, err
	}
	bkeFullState.DesiredState.CertificatesBundle = pki.TransformPEMToObject(bkeFullState.DesiredState.CertificatesBundle)
	bkeFullState.CurrentState.CertificatesBundle = pki.TransformPEMToObject(bkeFullState.CurrentState.CertificatesBundle)
	logrus.Tracef("bkeFullState: %+v", bkeFullState)

	return bkeFullState, nil
}

func ReadStateFile(ctx context.Context, statePath string) (*FullState, error) {
	bkeFullState := &FullState{}
	fp, err := filepath.Abs(statePath)
	if err != nil {
		return bkeFullState, fmt.Errorf("failed to lookup current directory name: %v", err)
	}
	file, err := os.Open(fp)
	if err != nil {
		return bkeFullState, fmt.Errorf("Can not find BKE state file: %v", err)
	}
	defer file.Close()
	buf, err := ioutil.ReadAll(file)
	if err != nil {
		return bkeFullState, fmt.Errorf("failed to read state file: %v", err)
	}
	if err := json.Unmarshal(buf, bkeFullState); err != nil {
		return bkeFullState, fmt.Errorf("failed to unmarshal the state file: %v", err)
	}
	bkeFullState.DesiredState.CertificatesBundle = pki.TransformPEMToObject(bkeFullState.DesiredState.CertificatesBundle)
	bkeFullState.CurrentState.CertificatesBundle = pki.TransformPEMToObject(bkeFullState.CurrentState.CertificatesBundle)
	return bkeFullState, nil
}

func RemoveStateFile(ctx context.Context, statePath string) {
	log.Infof(ctx, "Removing state file: %s", statePath)
	if err := os.Remove(statePath); err != nil {
		logrus.Warningf("Failed to remove state file: %v", err)
		return
	}
	log.Infof(ctx, "State file removed successfully")
}

func GetStateFromNodes(ctx context.Context, kubeCluster *Cluster) *Cluster {
	var currentCluster Cluster
	var clusterFile string
	var err error

	uniqueHosts := hosts.GetUniqueHostList(kubeCluster.EtcdHosts, kubeCluster.ControlPlaneHosts, kubeCluster.WorkerHosts)
	for _, host := range uniqueHosts {
		filePath := path.Join(pki.TempCertPath, pki.ClusterStateFile)
		clusterFile, err = pki.FetchFileFromHost(ctx, filePath, kubeCluster.SystemImages.Alpine, host, kubeCluster.PrivateRegistriesMap, pki.StateDeployerContainerName, "state", kubeCluster.Version)
		if err == nil {
			break
		}
	}
	if len(clusterFile) == 0 {
		return nil
	}
	err = yaml.Unmarshal([]byte(clusterFile), &currentCluster)
	if err != nil {
		logrus.Debugf("[state] Failed to unmarshal the cluster file fetched from nodes: %v", err)
		return nil
	}
	log.Infof(ctx, "[state] Successfully fetched cluster state from Nodes")
	return &currentCluster
}

func buildFreshState(ctx context.Context, kubeCluster *Cluster, newState *FullState) error {
	bkeConfig := &kubeCluster.BhojpurKubernetesEngineConfig
	// Get the certificate Bundle
	certBundle, err := pki.GenerateBKECerts(ctx, *bkeConfig, "", "")
	if err != nil {
		return fmt.Errorf("Failed to generate certificate bundle: %v", err)
	}
	newState.DesiredState.CertificatesBundle = certBundle
	if isEncryptionEnabled(bkeConfig) {
		if newState.DesiredState.EncryptionConfig, err = kubeCluster.getEncryptionProviderFile(); err != nil {
			return err
		}
	}
	return nil
}

func rebuildExistingState(ctx context.Context, kubeCluster *Cluster, oldState, newState *FullState, flags ExternalFlags) error {
	bkeConfig := &kubeCluster.BhojpurKubernetesEngineConfig
	pkiCertBundle := oldState.DesiredState.CertificatesBundle
	// check for legacy clusters prior to requestheaderca
	if pkiCertBundle[pki.RequestHeaderCACertName].Certificate == nil {
		if err := pki.GenerateBKERequestHeaderCACert(ctx, pkiCertBundle, flags.ClusterFilePath, flags.ConfigDir); err != nil {
			return err
		}
	}
	if err := pki.GenerateBKEServicesCerts(ctx, pkiCertBundle, *bkeConfig, flags.ClusterFilePath, flags.ConfigDir, false); err != nil {
		return err
	}
	newState.DesiredState.CertificatesBundle = pkiCertBundle
	err := updateEncryptionConfig(kubeCluster, oldState, newState)
	return err
}

func updateEncryptionConfig(kubeCluster *Cluster, oldState *FullState, newState *FullState) error {
	if isEncryptionEnabled(&kubeCluster.BhojpurKubernetesEngineConfig) {
		if oldState.DesiredState.EncryptionConfig != "" {
			newState.DesiredState.EncryptionConfig = oldState.DesiredState.EncryptionConfig
		} else {
			var err error
			if newState.DesiredState.EncryptionConfig, err = kubeCluster.getEncryptionProviderFile(); err != nil {
				return err
			}
		}
	}
	return nil
}

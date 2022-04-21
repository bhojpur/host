package hosts

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
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"path/filepath"

	"github.com/bhojpur/host/pkg/cluster/log"
	"github.com/bhojpur/host/pkg/engine/docker"
	"github.com/bhojpur/host/pkg/engine/metadata"
	"github.com/bhojpur/host/pkg/engine/util"
	"github.com/docker/docker/client"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

func (h *Host) TunnelUp(ctx context.Context, dialerFactory DialerFactory, clusterPrefixPath string, clusterVersion string) error {
	if h.DClient != nil {
		return nil
	}
	log.Infof(ctx, "[dialer] Setup tunnel for Bhojpur Host [%s]", h.Address)
	httpClient, err := h.newHTTPClient(dialerFactory)
	if err != nil {
		return fmt.Errorf("Can't establish dialer connection: %v", err)
	}
	// set Docker client
	logrus.Debugf("Connecting to Docker API for Bhojpur Host [%s]", h.Address)
	h.DClient, err = client.NewClientWithOpts(
		client.WithAPIVersionNegotiation(),
		client.WithHTTPClient(httpClient))
	if err != nil {
		return fmt.Errorf("Can't initiate NewClient: %v", err)
	}
	if err := checkDockerVersion(ctx, h, clusterVersion); err != nil {
		return err
	}
	h.SetPrefixPath(clusterPrefixPath)
	return nil
}

func (h *Host) TunnelUpLocal(ctx context.Context, clusterVersion string) error {
	var err error
	if h.DClient != nil {
		return nil
	}
	// set Docker client
	logrus.Debugf("Connecting to Docker API for Bhojpur Host [%s]", h.Address)
	h.DClient, err = client.NewEnvClient()
	if err != nil {
		return fmt.Errorf("Can't initiate NewClient: %v", err)
	}
	return checkDockerVersion(ctx, h, clusterVersion)
}

func checkDockerVersion(ctx context.Context, h *Host, clusterVersion string) error {
	info, err := h.DClient.Info(ctx)
	if err != nil {
		return fmt.Errorf("Can't retrieve Docker Info: %v", err)
	}
	logrus.Debugf("Docker Info found for Bhojpur Host [%s]: %#v", h.Address, info)
	h.DockerInfo = info
	if h.IgnoreDockerVersion {
		return nil
	}
	K8sSemVer, err := util.StrToSemVer(clusterVersion)
	if err != nil {
		return fmt.Errorf("Error while parsing cluster version [%s]: %v", clusterVersion, err)
	}
	K8sVersion := fmt.Sprintf("%d.%d", K8sSemVer.Major, K8sSemVer.Minor)
	isvalid, err := docker.IsSupportedDockerVersion(info, K8sVersion)
	if err != nil {
		return fmt.Errorf("Error while determining supported Docker version [%s]: %v", info.ServerVersion, err)
	}

	if !isvalid {
		return fmt.Errorf("Unsupported Docker version found [%s] on Bhojpur Host [%s], supported versions are %v", info.ServerVersion, h.Address, metadata.K8sVersionToDockerVersions[K8sVersion])
	}
	return nil
}

func parsePrivateKey(keyBuff string) (ssh.Signer, error) {
	return ssh.ParsePrivateKey([]byte(keyBuff))
}

func getSSHConfig(username, sshPrivateKeyString string, sshCertificateString string, useAgentAuth bool) (*ssh.ClientConfig, error) {
	config := &ssh.ClientConfig{
		User:            username,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// Kind of a double check now.
	if useAgentAuth {
		if sshAgentSock := os.Getenv("SSH_AUTH_SOCK"); sshAgentSock != "" {
			sshAgent, err := net.Dial("unix", sshAgentSock)
			if err != nil {
				return config, fmt.Errorf("Cannot connect to SSH Auth socket %q: %s", sshAgentSock, err)
			}

			config.Auth = append(config.Auth, ssh.PublicKeysCallback(agent.NewClient(sshAgent).Signers))

			logrus.Debugf("using %q SSH_AUTH_SOCK", sshAgentSock)
			return config, nil
		}
	}

	signer, err := parsePrivateKey(sshPrivateKeyString)
	if err != nil {
		return config, err
	}

	if len(sshCertificateString) > 0 {
		key, _, _, _, err := ssh.ParseAuthorizedKey([]byte(sshCertificateString))
		if err != nil {
			return config, fmt.Errorf("Unable to parse SSH certificate: %v", err)
		}

		if _, ok := key.(*ssh.Certificate); !ok {
			return config, fmt.Errorf("Unable to cast public key to SSH Certificate")
		}
		signer, err = ssh.NewCertSigner(key.(*ssh.Certificate), signer)
		if err != nil {
			return config, err
		}
	}

	config.Auth = append(config.Auth, ssh.PublicKeys(signer))

	return config, nil
}

func privateKeyPath(sshKeyPath string) (string, error) {
	if sshKeyPath[:2] == "~/" {
		sshKeyPath = filepath.Join(userHome(), sshKeyPath[2:])
	}
	buff, err := ioutil.ReadFile(sshKeyPath)
	if err != nil {
		return "", fmt.Errorf("Error while reading SSH key file: %v", err)
	}
	return string(buff), nil
}

func certificatePath(sshCertPath string) (string, error) {
	if sshCertPath[:2] == "~/" {
		sshCertPath = filepath.Join(userHome(), sshCertPath[2:])
	}
	buff, err := ioutil.ReadFile(sshCertPath)
	if err != nil {
		return "", fmt.Errorf("Error while reading SSH certificate file: %v", err)
	}
	return string(buff), nil
}

func userHome() string {
	if home := os.Getenv("HOME"); home != "" {
		return home
	}
	homeDrive := os.Getenv("HOMEDRIVE")
	homePath := os.Getenv("HOMEPATH")
	if homeDrive != "" && homePath != "" {
		return homeDrive + homePath
	}
	return os.Getenv("USERPROFILE")
}

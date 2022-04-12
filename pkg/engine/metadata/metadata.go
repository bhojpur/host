package metadata

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
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/bhojpur/host/pkg/data"
	v3 "github.com/bhojpur/host/pkg/engine/types"
	"github.com/bhojpur/host/pkg/engine/types/kdm"
	mVersion "github.com/mcuadros/go-version"
)

const (
	BhojpurMetadataURLEnv = "BHOJPUR_METADATA_URL"
)

var (
	BKEVersion                  string
	DefaultK8sVersion           string
	K8sVersionToTemplates       map[string]map[string]string
	K8sVersionToBKESystemImages map[string]v3.BKESystemImages
	K8sVersionToServiceOptions  map[string]v3.KubernetesServicesOptions
	K8sVersionToDockerVersions  map[string][]string
	K8sVersionsCurrent          []string
	K8sBadVersions              = map[string]bool{}

	K8sVersionToWindowsServiceOptions map[string]v3.KubernetesServicesOptions

	c = http.Client{
		Timeout: time.Second * 30,
	}
	kdmMutex = sync.Mutex{}
)

func InitMetadata(ctx context.Context) error {
	kdmMutex.Lock()
	defer kdmMutex.Unlock()
	data, err := loadData()
	if err != nil {
		return fmt.Errorf("failed to load data.json, error: %v", err)
	}
	initK8sBKESystemImages(data)
	initAddonTemplates(data)
	initServiceOptions(data)
	initDockerOptions(data)
	return nil
}

// this method loads metadata, if BHOJPUR_METADATA_URL is provided then load data from specified location. Otherwise load data from bindata.
func loadData() (kdm.Data, error) {
	var b []byte
	var err error
	u := os.Getenv(BhojpurMetadataURLEnv)
	if u != "" {
		logrus.Debugf("Loading data.json from %s", u)
		b, err = readFile(u)
		if err != nil {
			return kdm.Data{}, err
		}
	} else {
		logrus.Debug("Loading data.json from local source")
		b, err = data.Asset("data/data.json")
		if err != nil {
			return kdm.Data{}, err
		}
	}
	logrus.Debugf("data.json SHA256 checksum: %x", sha256.Sum256(b))
	logrus.Tracef("data.json content: %v", string(b))
	return kdm.FromData(b)
}

func readFile(file string) ([]byte, error) {
	if strings.HasPrefix(file, "http") {
		resp, err := c.Get(file)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		return ioutil.ReadAll(resp.Body)
	}
	return ioutil.ReadFile(file)
}

const BKEVersionDev = "v1.3.99"

func initAddonTemplates(data kdm.Data) {
	K8sVersionToTemplates = data.K8sVersionedTemplates
}

func initServiceOptions(data kdm.Data) {
	K8sVersionToServiceOptions = interface{}(data.K8sVersionServiceOptions).(map[string]v3.KubernetesServicesOptions)
	K8sVersionToWindowsServiceOptions = data.K8sVersionWindowsServiceOptions
}

func initDockerOptions(data kdm.Data) {
	K8sVersionToDockerVersions = data.K8sVersionDockerInfo
}

func initK8sBKESystemImages(data kdm.Data) {
	K8sVersionToBKESystemImages = map[string]v3.BKESystemImages{}
	bkeData := data
	// non released versions
	if BKEVersion == "" {
		BKEVersion = BKEVersionDev
	}
	DefaultK8sVersion = bkeData.BKEDefaultK8sVersions["default"]
	if defaultK8sVersion, ok := bkeData.BKEDefaultK8sVersions[BKEVersion[1:]]; ok {
		DefaultK8sVersion = defaultK8sVersion
	}
	maxVersionForMajorK8sVersion := map[string]string{}
	for k8sVersion, systemImages := range bkeData.K8sVersionBKESystemImages {
		bkeVersionInfo, ok := bkeData.K8sVersionInfo[k8sVersion]
		if ok {
			// BKEVersion = 0.2.4, DeprecateBKEVersion = 0.2.2
			if bkeVersionInfo.DeprecateBKEVersion != "" && mVersion.Compare(BKEVersion, bkeVersionInfo.DeprecateBKEVersion, ">=") {
				K8sBadVersions[k8sVersion] = true
				continue
			}
			// BKEVersion = 0.2.4, MinVersion = 0.2.5, don't store
			lowerThanMin := bkeVersionInfo.MinBKEVersion != "" && mVersion.Compare(BKEVersion, bkeVersionInfo.MinBKEVersion, "<")
			if lowerThanMin {
				continue
			}
		}
		// store all for upgrades
		K8sVersionToBKESystemImages[k8sVersion] = interface{}(systemImages).(v3.BKESystemImages)

		majorVersion := getTagMajorVersion(k8sVersion)
		maxVersionInfo, ok := bkeData.K8sVersionInfo[majorVersion]
		if ok {
			// BKEVersion = 0.2.4, MaxVersion = 0.2.3, don't use in current
			greaterThanMax := maxVersionInfo.MaxBKEVersion != "" && mVersion.Compare(BKEVersion, maxVersionInfo.MaxBKEVersion, ">")
			if greaterThanMax {
				continue
			}
		}
		if curr, ok := maxVersionForMajorK8sVersion[majorVersion]; !ok || mVersion.Compare(k8sVersion, curr, ">") {
			maxVersionForMajorK8sVersion[majorVersion] = k8sVersion
		}
	}
	for _, k8sVersion := range maxVersionForMajorK8sVersion {
		K8sVersionsCurrent = append(K8sVersionsCurrent, k8sVersion)
	}
}

func getTagMajorVersion(tag string) string {
	splitTag := strings.Split(tag, ".")
	if len(splitTag) < 2 {
		return ""
	}
	return strings.Join(splitTag[:2], ".")
}

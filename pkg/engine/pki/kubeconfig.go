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

import "encoding/base64"

func getKubeConfigX509(kubernetesURL string, clusterName string, componentName string, caPath string, crtPath string, keyPath string) string {
	return `apiVersion: v1
kind: Config
clusters:
- cluster:
    api-version: v1
    certificate-authority: ` + caPath + `
    server: "` + kubernetesURL + `"
  name: "` + clusterName + `"
contexts:
- context:
    cluster: "` + clusterName + `"
    user: "` + componentName + `-` + clusterName + `"
  name: "` + clusterName + `"
current-context: "` + clusterName + `"
users:
- name: "` + componentName + `-` + clusterName + `"
  user:
    client-certificate: ` + crtPath + `
    client-key: ` + keyPath + ``
}

func GetKubeConfigX509WithData(kubernetesURL string, clusterName string, componentName string, cacrt string, crt string, key string) string {
	return `apiVersion: v1
kind: Config
clusters:
- cluster:
    api-version: v1
    certificate-authority-data: ` + base64.StdEncoding.EncodeToString([]byte(cacrt)) + `
    server: "` + kubernetesURL + `"
  name: "` + clusterName + `"
contexts:
- context:
    cluster: "` + clusterName + `"
    user: "` + componentName + `-` + clusterName + `"
  name: "` + clusterName + `"
current-context: "` + clusterName + `"
users:
- name: "` + componentName + `-` + clusterName + `"
  user:
    client-certificate-data: ` + base64.StdEncoding.EncodeToString([]byte(crt)) + `
    client-key-data: ` + base64.StdEncoding.EncodeToString([]byte(key)) + ``
}

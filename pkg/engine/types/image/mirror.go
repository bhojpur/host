package image

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

import "strings"

var Mirrors = map[string]string{}

func Mirror(image string) string {
	orig := image
	if strings.HasPrefix(image, "weaveworks") || strings.HasPrefix(image, "noiro") {
		return image
	}

	image = strings.Replace(image, "gcr.io/google_containers", "bhojpur", 1)
	image = strings.Replace(image, "quay.io/coreos/", "bhojpur/coreos-", 1)
	image = strings.Replace(image, "quay.io/calico/", "bhojpur/calico-", 1)
	image = strings.Replace(image, "plugins/docker", "bhojpur/plugins-docker", 1)
	image = strings.Replace(image, "k8s.gcr.io/defaultbackend", "bhojpur/nginx-ingress-controller-defaultbackend", 1)
	image = strings.Replace(image, "k8s.gcr.io/k8s-dns-node-cache", "bhojpur/k8s-dns-node-cache", 1)
	image = strings.Replace(image, "plugins/docker", "bhojpur/plugins-docker", 1)
	image = strings.Replace(image, "kibana", "bhojpur/kibana", 1)
	image = strings.Replace(image, "jenkins/", "bhojpur/jenkins-", 1)
	image = strings.Replace(image, "alpine/git", "bhojpur/alpine-git", 1)
	image = strings.Replace(image, "prom/", "bhojpur/prom-", 1)
	image = strings.Replace(image, "quay.io/pires", "bhojpur", 1)
	image = strings.Replace(image, "coredns/", "bhojpur/coredns-", 1)
	image = strings.Replace(image, "minio/", "bhojpur/minio-", 1)

	Mirrors[image] = orig
	return image
}
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

import (
	projectv3 "github.com/bhojpur/host/pkg/apis/project.bhojpur.net/v3"
)

var (
	ToolsSystemImages = struct {
		PipelineSystemImages projectv3.PipelineSystemImages
		AuthSystemImages     AuthSystemImages
	}{
		PipelineSystemImages: projectv3.PipelineSystemImages{
			Jenkins:       "bhojpur/pipeline-jenkins-server:v0.1.4",
			JenkinsJnlp:   "bhojpur/mirrored-jenkins-jnlp-slave:4.7-1",
			AlpineGit:     "bhojpur/pipeline-tools:v0.1.16",
			PluginsDocker: "bhojpur/mirrored-plugins-docker:19.03.8",
			Minio:         "bhojpur/mirrored-minio-minio:RELEASE.2020-07-13T18-09-56Z",
			Registry:      "registry:2",
			RegistryProxy: "bhojpur/pipeline-tools:v0.1.16",
			KubeApply:     "bhojpur/pipeline-tools:v0.1.16",
		},
		AuthSystemImages: AuthSystemImages{
			KubeAPIAuth: "bhojpur/kube-api-auth:v0.1.8",
		},
	}
)

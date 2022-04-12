package addons

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
	"bytes"
	"fmt"
	"testing"

	v1 "k8s.io/api/batch/v1"
	yamlutil "k8s.io/apimachinery/pkg/util/yaml"
)

const (
	AddonSuffix    = "-deploy-job"
	FakeAddonName  = "example-addon"
	FakeNodeName   = "node1"
	FakeAddonImage = "example/example:latest"
	FakeK8sVersion = "v1.21.1-bhojpur1-1"
)

func TestJobManifest(t *testing.T) {
	jobYaml, err := GetAddonsExecuteJob(FakeAddonName, FakeNodeName, FakeAddonImage, FakeK8sVersion)
	if err != nil {
		t.Fatalf("Failed to get addon execute job: %v", err)
	}
	job := v1.Job{}
	decoder := yamlutil.NewYAMLToJSONDecoder(bytes.NewReader([]byte(jobYaml)))
	err = decoder.Decode(&job)
	if err != nil {
		t.Fatalf("Failed To decode Job yaml: %v", err)
	}
	assertEqual(t, job.Name, FakeAddonName+AddonSuffix,
		fmt.Sprintf("Failed to verify job name [%s]", FakeAddonName+AddonSuffix))
	assertEqual(t, job.Spec.Template.Spec.NodeName, FakeNodeName,
		fmt.Sprintf("Failed to verify node name [%s] in the job", FakeNodeName))
	assertEqual(t, job.Spec.Template.Spec.Containers[0].Image, FakeAddonImage,
		fmt.Sprintf("Failed to verify container image [%s] in the job", FakeAddonImage))
}

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}

package mapper

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
	"strings"

	"github.com/bhojpur/host/pkg/core/types"
	"github.com/bhojpur/host/pkg/core/types/convert"
	"github.com/bhojpur/host/pkg/core/types/values"
	"k8s.io/apimachinery/pkg/api/resource"
)

type OSInfo struct {
}

func (o OSInfo) FromInternal(data map[string]interface{}) {
	if data == nil {
		return
	}
	cpuInfo := map[string]interface{}{}
	cpuNum, err := resource.ParseQuantity(convert.ToString(values.GetValueN(data, "capacity", "cpu")))
	if err == nil {
		cpuInfo["count"] = cpuNum.Value()
	}

	memoryInfo := map[string]interface{}{}
	kibNum, err := resource.ParseQuantity(convert.ToString(values.GetValueN(data, "capacity", "memory")))
	if err == nil {
		memoryInfo["memTotalKiB"] = kibNum.Value() / 1024
	}

	osInfo := map[string]interface{}{
		"dockerVersion":   strings.TrimPrefix(convert.ToString(values.GetValueN(data, "nodeInfo", "containerRuntimeVersion")), "docker://"),
		"kernelVersion":   values.GetValueN(data, "nodeInfo", "kernelVersion"),
		"operatingSystem": values.GetValueN(data, "nodeInfo", "osImage"),
	}

	data["info"] = map[string]interface{}{
		"cpu":    cpuInfo,
		"memory": memoryInfo,
		"os":     osInfo,
		"kubernetes": map[string]interface{}{
			"kubeletVersion":   values.GetValueN(data, "nodeInfo", "kubeletVersion"),
			"kubeProxyVersion": values.GetValueN(data, "nodeInfo", "kubeletVersion"),
		},
	}
}

func (o OSInfo) ToInternal(data map[string]interface{}) error {
	return nil
}

func (o OSInfo) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	return nil
}

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
	"fmt"
	"strings"

	"github.com/bhojpur/host/pkg/core/types/values"

	"github.com/bhojpur/host/pkg/core/types"
	"github.com/bhojpur/host/pkg/core/types/convert"
)

var (
	kindMap = map[string]string{
		"deployment":            "Deployment",
		"replicationcontroller": "ReplicationController",
		"statefulset":           "StatefulSet",
		"daemonset":             "DaemonSet",
		"job":                   "Job",
		"cronjob":               "CronJob",
		"replicaset":            "ReplicaSet",
	}
	groupVersionMap = map[string]string{
		"deployment":            "apps/v1beta2",
		"replicationcontroller": "core/v1",
		"statefulset":           "apps/v1beta2",
		"daemonset":             "apps/v1beta2",
		"job":                   "batch/v1",
		"cronjob":               "batch/v1beta1",
		"replicaset":            "apps/v1beta2",
	}
)

type CrossVersionObjectToWorkload struct {
	Field string
}

func (c CrossVersionObjectToWorkload) ToInternal(data map[string]interface{}) error {
	obj, ok := values.GetValue(data, strings.Split(c.Field, "/")...)
	if !ok {
		return nil
	}
	workloadID := convert.ToString(obj)
	parts := strings.SplitN(workloadID, ":", 3)
	newObj := map[string]interface{}{
		"kind":       getKind(parts[0]),
		"apiVersion": groupVersionMap[parts[0]],
		"name":       parts[2],
	}
	values.PutValue(data, newObj, strings.Split(c.Field, "/")...)
	return nil
}

func (c CrossVersionObjectToWorkload) FromInternal(data map[string]interface{}) {
	obj, ok := values.GetValue(data, strings.Split(c.Field, "/")...)
	if !ok {
		return
	}
	cvo := convert.ToMapInterface(obj)
	ns := convert.ToString(data["namespaceId"])
	values.PutValue(data,
		fmt.Sprintf("%s:%s:%s",
			strings.ToLower(convert.ToString(cvo["kind"])),
			ns,
			convert.ToString(cvo["name"])),
		strings.Split(c.Field, "/")...,
	)
}

func (c CrossVersionObjectToWorkload) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	return nil
}

func getKind(i string) string {
	return kindMap[i]
}

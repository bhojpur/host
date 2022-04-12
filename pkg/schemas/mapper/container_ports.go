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
	"github.com/bhojpur/host/pkg/core/types/mapper"
	"github.com/bhojpur/host/pkg/core/types/values"
	"github.com/sirupsen/logrus"
)

type ContainerPorts struct {
}

func (n ContainerPorts) FromInternal(data map[string]interface{}) {
	field := mapper.AnnotationField{
		Field: "ports",
		List:  true,
	}
	field.FromInternal(data)

	containers := convert.ToInterfaceSlice(data["containers"])
	annotationPorts := convert.ToInterfaceSlice(data["ports"])
	annotationsPortsMap := map[string]map[string]interface{}{}

	// process fields defined via annotations
	for i := 0; i < len(annotationPorts) && i < len(containers); i++ {
		container := convert.ToMapInterface(containers[i])
		if container != nil {
			portsSlice := convert.ToInterfaceSlice(annotationPorts[i])
			portMap := map[string]interface{}{}
			for _, port := range portsSlice {
				asMap, err := convert.EncodeToMap(port)
				if err != nil {
					logrus.Warnf("Failed to convert container port to map %v", err)
					continue
				}
				asMap["type"] = "/v3/project/schemas/containerPort"
				name, _ := values.GetValue(asMap, "name")
				portMap[convert.ToString(name)] = asMap
			}
			containerName, _ := values.GetValue(container, "name")
			annotationsPortsMap[convert.ToString(containerName)] = portMap
		}
	}

	for _, container := range containers {
		// iterate over container ports and see if some of them are not defined via annotation
		// set kind to hostport if source port is set, and clusterip if it is not
		containerMap := convert.ToMapInterface(container)
		containerName, _ := values.GetValue(containerMap, "name")
		portMap := annotationsPortsMap[convert.ToString(containerName)]
		if portMap == nil {
			portMap = map[string]interface{}{}
		}
		var containerPorts []interface{}
		containerPortSlice := convert.ToInterfaceSlice(containerMap["ports"])
		for _, port := range containerPortSlice {
			asMap, err := convert.EncodeToMap(port)
			if err != nil {
				logrus.Warnf("Failed to convert container port to map %v", err)
				continue
			}
			portName, _ := values.GetValue(asMap, "name")
			if annotationPort, ok := portMap[convert.ToString(portName)]; ok {
				containerPorts = append(containerPorts, annotationPort)
			} else {
				hostPort, _ := values.GetValue(asMap, "hostPort")
				if hostPort == nil {
					asMap["kind"] = "ClusterIP"
				} else {
					asMap["sourcePort"] = hostPort
					asMap["kind"] = "HostPort"
				}
				containerPorts = append(containerPorts, asMap)
			}
		}
		containerMap["ports"] = containerPorts
	}

}

func (n ContainerPorts) ToInternal(data map[string]interface{}) error {
	field := mapper.AnnotationField{
		Field: "ports",
		List:  true,
	}

	var ports []interface{}
	path := []string{"containers", "{ARRAY}", "ports"}
	convert.Transform(data, path, func(obj interface{}) interface{} {
		if l, ok := obj.([]interface{}); ok {
			for _, p := range l {
				mapped, err := convert.EncodeToMap(p)
				if err != nil {
					logrus.Warnf("Failed to encode port: %v", err)
					return obj
				}
				if strings.EqualFold(convert.ToString(mapped["kind"]), "HostPort") {
					mapped["hostPort"] = mapped["sourcePort"]
				}
			}
			ports = append(ports, l)
		}
		return obj
	})

	if len(ports) != 0 {
		data["ports"] = ports
		return field.ToInternal(data)
	}

	return nil
}

func (n ContainerPorts) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	return nil
}

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
	"github.com/bhojpur/host/pkg/core/types"
	"github.com/bhojpur/host/pkg/core/types/values"
)

// These mappers copy data from the networking.k8s.io/v1-style Ingress fields
// to extensions/v1beta1-style Ingress fields so that when the proxy store
// serializes them into kubernetes resources, there is no loss of data and the
// objects are in the expected format for the API in use. The end result is
// that some data is duplicated on the object, but extraneous fields will be
// ignored by the proxy store.

// IngressSpec mapper copies defaultBackend (type k8s.io/api/networking/v1.IngressBackend)
// to backend (type k8s.io/api/extensions/v1beta1.IngressBackend) on Spec.
type IngressSpec struct{}

func (i IngressSpec) FromInternal(data map[string]interface{}) {
	if _, ok := data["backend"]; ok && data["defaultBackend"] == nil {
		data["defaultBackend"] = map[string]interface{}{
			"targetPort": values.GetValueN(data, "backend", "servicePort"),
			"serviceId":  values.GetValueN(data, "backend", "serviceName"),
		}
		delete(data, "backend")
	}
	return
}

func (i IngressSpec) ToInternal(data map[string]interface{}) error {
	if backend, ok := data["backend"]; (!ok || backend == nil) && data["defaultBackend"] != nil {
		data["backend"] = map[string]interface{}{
			"servicePort": values.GetValueN(data, "defaultBackend", "targetPort"),
			"serviceName": values.GetValueN(data, "defaultBackend", "serviceId"),
		}
	}
	return nil
}

func (i IngressSpec) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	return nil
}

// IngressBackend mapper copies service fields within
// k8s.io/api/networking/v1.IngressBackend to the equivalents in
// k8s.io/api/extensions/v1beta1.IngressBackend.
// Applies to the Spec.DefaultBackend and Spec.Rules[*].Backend.
type IngressBackend struct{}

func (i IngressBackend) FromInternal(data map[string]interface{}) {
	if data == nil {
		return
	}
	if serviceID, ok := data["serviceId"]; !ok || serviceID == nil {
		data["serviceId"] = values.GetValueN(data, "serviceName")
	}
	if targetPort, ok := data["targetPort"]; !ok || targetPort == nil {
		if port := values.GetValueN(data, "servicePort"); port != nil {
			data["targetPort"] = port
		} else if port := values.GetValueN(data, "service", "port", "number"); port != nil {
			data["targetPort"] = port
		} else if port := values.GetValueN(data, "service", "port", "name"); port != nil {
			data["targetPort"] = port
		}
	}
}

func (i IngressBackend) ToInternal(data map[string]interface{}) error {
	if data != nil {
		serviceID := values.GetValueN(data, "serviceId")
		values.PutValue(data, serviceID, "serviceName")
		values.PutValue(data, serviceID, "service", "name")
		port := values.GetValueN(data, "targetPort")
		data["servicePort"] = port
		switch port.(type) {
		case int64:
			values.PutValue(data, port, "service", "port", "number")
		case string:
			values.PutValue(data, port, "service", "port", "name")
		}
	}
	return nil
}

func (i IngressBackend) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	return nil
}

// IngressPath mapper applies the required pathType field to
// k8s.io/api/networking/v1.HTTPIngressPath.
type IngressPath struct{}

func (i IngressPath) FromInternal(data map[string]interface{}) {
	return
}

func (i IngressPath) ToInternal(data map[string]interface{}) error {
	if values.GetValueN(data, "pathType") == nil {
		values.PutValue(data, "ImplementationSpecific", "pathType")
	}
	return nil
}

func (i IngressPath) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	return nil
}

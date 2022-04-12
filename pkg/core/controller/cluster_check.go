package controller

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
	"os"
	"reflect"
	"strings"

	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
)

type ObjectClusterName interface {
	ObjClusterName() string
}

func ObjectInCluster(cluster string, obj interface{}) bool {
	// Check if the object implements the interface, this is best case and
	// what objects should strive to be
	if o, ok := obj.(ObjectClusterName); ok {
		return o.ObjClusterName() == cluster
	}

	// For types outside of Bhojpur Host, attempt to check the anno, then use the namespace
	// This is much better than using the reflect hole below
	switch v := obj.(type) {
	case *corev1.Secret:
		if c, ok := v.Annotations["field.bhojpur.net/projectId"]; ok {
			if parts := strings.SplitN(c, ":", 2); len(parts) == 2 {
				return cluster == parts[0]
			}
		}
		return v.Namespace == cluster
	case *corev1.Namespace:
		if c, ok := v.Annotations["field.bhojpur.net/projectId"]; ok {
			if parts := strings.SplitN(c, ":", 2); len(parts) == 2 {
				return cluster == parts[0]
			}
		}
		return v.Namespace == cluster
	case *corev1.Node:
		if c, ok := v.Annotations["field.bhojpur.net/projectId"]; ok {
			if parts := strings.SplitN(c, ":", 2); len(parts) == 2 {
				return cluster == parts[0]
			}
		}
		return v.Namespace == cluster
	}

	// Seeing this message means something needs to be done with the type, see comments above
	if dm := os.Getenv("BHOJPUR_DEV_MODE"); dm != "" {
		logrus.Errorf("ObjectClusterName not implemented by type %T", obj)
	}

	var clusterName string

	if c := getValue(obj, "ClusterName"); c.IsValid() {
		clusterName = c.String()
	}
	if clusterName == "" {
		if c := getValue(obj, "Spec", "ClusterName"); c.IsValid() {
			clusterName = c.String()
		}

	}
	if clusterName == "" {
		if c := getValue(obj, "ProjectName"); c.IsValid() {
			if parts := strings.SplitN(c.String(), ":", 2); len(parts) == 2 {
				clusterName = parts[0]
			}
		}
	}
	if clusterName == "" {
		if c := getValue(obj, "Spec", "ProjectName"); c.IsValid() {
			if parts := strings.SplitN(c.String(), ":", 2); len(parts) == 2 {
				clusterName = parts[0]
			}
		}
	}
	if clusterName == "" {
		if a := getValue(obj, "Annotations"); a.IsValid() {
			if c := a.MapIndex(reflect.ValueOf("field.bhojpur.net/projectId")); c.IsValid() {
				if parts := strings.SplitN(c.String(), ":", 2); len(parts) == 2 {
					clusterName = parts[0]
				}
			}
		}
	}
	if clusterName == "" {
		if c := getValue(obj, "Namespace"); c.IsValid() {
			clusterName = c.String()
		}
	}

	return clusterName == cluster
}

func getValue(obj interface{}, name ...string) reflect.Value {
	v := reflect.ValueOf(obj)
	t := v.Type()
	if t.Kind() == reflect.Ptr {
		v = v.Elem()
		t = v.Type()
	}

	field := v.FieldByName(name[0])
	if !field.IsValid() || len(name) == 1 {
		return field
	}

	return getFieldValue(field, name[1:]...)
}

func getFieldValue(v reflect.Value, name ...string) reflect.Value {
	field := v.FieldByName(name[0])
	if len(name) == 1 {
		return field
	}
	return getFieldValue(field, name[1:]...)
}

package apply

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
	"encoding/json"
	"fmt"
	"reflect"

	appsv1 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	defaultReconcilers = map[schema.GroupVersionKind]Reconciler{
		v1.SchemeGroupVersion.WithKind("Secret"):         reconcileSecret,
		v1.SchemeGroupVersion.WithKind("Service"):        reconcileService,
		batchv1.SchemeGroupVersion.WithKind("Job"):       reconcileJob,
		appsv1.SchemeGroupVersion.WithKind("Deployment"): reconcileDeployment,
		appsv1.SchemeGroupVersion.WithKind("DaemonSet"):  reconcileDaemonSet,
	}
)

func reconcileDaemonSet(oldObj, newObj runtime.Object) (bool, error) {
	oldSvc, ok := oldObj.(*appsv1.DaemonSet)
	if !ok {
		oldSvc = &appsv1.DaemonSet{}
		if err := convertObj(oldObj, oldSvc); err != nil {
			return false, err
		}
	}
	newSvc, ok := newObj.(*appsv1.DaemonSet)
	if !ok {
		newSvc = &appsv1.DaemonSet{}
		if err := convertObj(newObj, newSvc); err != nil {
			return false, err
		}
	}

	if !equality.Semantic.DeepEqual(oldSvc.Spec.Selector, newSvc.Spec.Selector) {
		return false, ErrReplace
	}

	return false, nil
}

func reconcileDeployment(oldObj, newObj runtime.Object) (bool, error) {
	oldSvc, ok := oldObj.(*appsv1.Deployment)
	if !ok {
		oldSvc = &appsv1.Deployment{}
		if err := convertObj(oldObj, oldSvc); err != nil {
			return false, err
		}
	}
	newSvc, ok := newObj.(*appsv1.Deployment)
	if !ok {
		newSvc = &appsv1.Deployment{}
		if err := convertObj(newObj, newSvc); err != nil {
			return false, err
		}
	}

	if !equality.Semantic.DeepEqual(oldSvc.Spec.Selector, newSvc.Spec.Selector) {
		return false, ErrReplace
	}

	return false, nil
}

func reconcileSecret(oldObj, newObj runtime.Object) (bool, error) {
	oldSvc, ok := oldObj.(*v1.Secret)
	if !ok {
		oldSvc = &v1.Secret{}
		if err := convertObj(oldObj, oldSvc); err != nil {
			return false, err
		}
	}
	newSvc, ok := newObj.(*v1.Secret)
	if !ok {
		newSvc = &v1.Secret{}
		if err := convertObj(newObj, newSvc); err != nil {
			return false, err
		}
	}

	if newSvc.Type != "" && oldSvc.Type != newSvc.Type {
		return false, ErrReplace
	}

	return false, nil
}

func reconcileService(oldObj, newObj runtime.Object) (bool, error) {
	oldSvc, ok := oldObj.(*v1.Service)
	if !ok {
		oldSvc = &v1.Service{}
		if err := convertObj(oldObj, oldSvc); err != nil {
			return false, err
		}
	}
	newSvc, ok := newObj.(*v1.Service)
	if !ok {
		newSvc = &v1.Service{}
		if err := convertObj(newObj, newSvc); err != nil {
			return false, err
		}
	}

	if newSvc.Spec.Type != "" && oldSvc.Spec.Type != newSvc.Spec.Type {
		return false, ErrReplace
	}

	return false, nil
}

func reconcileJob(oldObj, newObj runtime.Object) (bool, error) {
	oldSvc, ok := oldObj.(*batchv1.Job)
	if !ok {
		oldSvc = &batchv1.Job{}
		if err := convertObj(oldObj, oldSvc); err != nil {
			return false, err
		}
	}

	newSvc, ok := newObj.(*batchv1.Job)
	if !ok {
		newSvc = &batchv1.Job{}
		if err := convertObj(newObj, newSvc); err != nil {
			return false, err
		}
	}

	if !equality.Semantic.DeepEqual(oldSvc.Spec.Template, newSvc.Spec.Template) {
		return false, ErrReplace
	}

	return false, nil
}

func convertObj(src interface{}, obj interface{}) error {
	uObj, ok := src.(*unstructured.Unstructured)
	if !ok {
		return fmt.Errorf("expected unstructured but got %v", reflect.TypeOf(src))
	}

	bytes, err := uObj.MarshalJSON()
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, obj)
}

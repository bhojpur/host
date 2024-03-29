package lifecycle

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
	"reflect"

	"github.com/bhojpur/host/pkg/core/objectclient"
	"github.com/bhojpur/host/pkg/core/types/slice"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

var (
	created            = "lifecycle.bhojpur.net/create"
	finalizerKey       = "controller.bhojpur.net/"
	ScopedFinalizerKey = "clusterscoped.controller.bhojpur.net/"
)

type ObjectLifecycle interface {
	Create(obj runtime.Object) (runtime.Object, error)
	Finalize(obj runtime.Object) (runtime.Object, error)
	Updated(obj runtime.Object) (runtime.Object, error)
}

type ObjectLifecycleCondition interface {
	HasCreate() bool
	HasFinalize() bool
}

type objectLifecycleAdapter struct {
	name          string
	clusterScoped bool
	lifecycle     ObjectLifecycle
	objectClient  *objectclient.ObjectClient
}

func NewObjectLifecycleAdapter(name string, clusterScoped bool, lifecycle ObjectLifecycle, objectClient *objectclient.ObjectClient) func(key string, obj interface{}) (interface{}, error) {
	o := objectLifecycleAdapter{
		name:          name,
		clusterScoped: clusterScoped,
		lifecycle:     lifecycle,
		objectClient:  objectClient,
	}
	return o.sync
}

func (o *objectLifecycleAdapter) sync(key string, in interface{}) (interface{}, error) {
	if in == nil || reflect.ValueOf(in).IsNil() {
		return nil, nil
	}

	obj, ok := in.(runtime.Object)
	if !ok {
		return nil, nil
	}

	if newObj, cont, err := o.finalize(obj); err != nil || !cont {
		return nil, err
	} else if newObj != nil {
		obj = newObj
	}

	if newObj, cont, err := o.create(obj); err != nil || !cont {
		return nil, err
	} else if newObj != nil {
		obj = newObj
	}

	return o.record(obj, o.lifecycle.Updated)
}

func (o *objectLifecycleAdapter) update(name string, orig, obj runtime.Object) (runtime.Object, error) {
	if obj != nil && orig != nil && !reflect.DeepEqual(orig, obj) {
		newObj, err := o.objectClient.Update(name, obj)
		if newObj != nil {
			return newObj, err
		}
		return obj, err
	}
	if obj == nil {
		return orig, nil
	}
	return obj, nil
}

func (o *objectLifecycleAdapter) finalize(obj runtime.Object) (runtime.Object, bool, error) {
	if !o.hasFinalize() {
		return obj, true, nil
	}

	metadata, err := meta.Accessor(obj)
	if err != nil {
		return obj, false, err
	}

	// Check finalize
	if metadata.GetDeletionTimestamp() == nil {
		return nil, true, nil
	}

	if !slice.ContainsString(metadata.GetFinalizers(), o.constructFinalizerKey()) {
		return nil, false, nil
	}

	newObj, err := o.record(obj, o.lifecycle.Finalize)
	if err != nil {
		return obj, false, err
	}

	obj, err = o.removeFinalizer(o.constructFinalizerKey(), maybeDeepCopy(obj, newObj))
	return obj, false, err
}

func maybeDeepCopy(old, newObj runtime.Object) runtime.Object {
	if old == newObj {
		return old.DeepCopyObject()
	}
	return newObj
}

func (o *objectLifecycleAdapter) removeFinalizer(name string, obj runtime.Object) (runtime.Object, error) {
	for i := 0; i < 3; i++ {
		metadata, err := meta.Accessor(obj)
		if err != nil {
			return nil, err
		}

		var finalizers []string
		for _, finalizer := range metadata.GetFinalizers() {
			if finalizer == name {
				continue
			}
			finalizers = append(finalizers, finalizer)
		}
		metadata.SetFinalizers(finalizers)

		newObj, err := o.objectClient.Update(metadata.GetName(), obj)
		if err == nil {
			return newObj, nil
		}

		obj, err = o.objectClient.GetNamespaced(metadata.GetNamespace(), metadata.GetName(), metav1.GetOptions{})
		if err != nil {
			return nil, err
		}
	}

	return nil, fmt.Errorf("failed to remove finalizer on %s", name)
}

func (o *objectLifecycleAdapter) createKey() string {
	return created + "." + o.name
}

func (o *objectLifecycleAdapter) constructFinalizerKey() string {
	if o.clusterScoped {
		return ScopedFinalizerKey + o.name
	}
	return finalizerKey + o.name
}

func (o *objectLifecycleAdapter) hasFinalize() bool {
	cond, ok := o.lifecycle.(ObjectLifecycleCondition)
	return !ok || cond.HasFinalize()
}

func (o *objectLifecycleAdapter) hasCreate() bool {
	cond, ok := o.lifecycle.(ObjectLifecycleCondition)
	return !ok || cond.HasCreate()
}

func (o *objectLifecycleAdapter) record(obj runtime.Object, f func(runtime.Object) (runtime.Object, error)) (runtime.Object, error) {
	metadata, err := meta.Accessor(obj)
	if err != nil {
		return obj, err
	}

	origObj := obj
	obj = origObj.DeepCopyObject()
	if newObj, err := checkNil(obj, f); err != nil {
		newObj, _ = o.update(metadata.GetName(), origObj, newObj)
		return newObj, err
	} else if newObj != nil {
		return o.update(metadata.GetName(), origObj, newObj)
	}
	return obj, nil
}

func checkNil(obj runtime.Object, f func(runtime.Object) (runtime.Object, error)) (runtime.Object, error) {
	obj, err := f(obj)
	if obj == nil || reflect.ValueOf(obj).IsNil() {
		return nil, err
	}
	return obj, err
}

func (o *objectLifecycleAdapter) create(obj runtime.Object) (runtime.Object, bool, error) {
	metadata, err := meta.Accessor(obj)
	if err != nil {
		return obj, false, err
	}

	if o.isInitialized(metadata) {
		return nil, true, nil
	}

	if o.hasFinalize() {
		obj, err = o.addFinalizer(obj)
		if err != nil {
			return obj, false, err
		}
	}

	if !o.hasCreate() {
		return obj, true, err
	}

	obj, err = o.record(obj, o.lifecycle.Create)
	if err != nil {
		return obj, false, err
	}

	obj, err = o.setInitialized(obj)
	return obj, false, err
}

func (o *objectLifecycleAdapter) isInitialized(metadata metav1.Object) bool {
	initialized := o.createKey()
	return metadata.GetAnnotations()[initialized] == "true"
}

func (o *objectLifecycleAdapter) setInitialized(obj runtime.Object) (runtime.Object, error) {
	metadata, err := meta.Accessor(obj)
	if err != nil {
		return nil, err
	}

	initialized := o.createKey()

	if metadata.GetAnnotations() == nil {
		metadata.SetAnnotations(map[string]string{})
	}
	metadata.GetAnnotations()[initialized] = "true"

	return o.objectClient.Update(metadata.GetName(), obj)
}

func (o *objectLifecycleAdapter) addFinalizer(obj runtime.Object) (runtime.Object, error) {
	metadata, err := meta.Accessor(obj)
	if err != nil {
		return nil, err
	}

	if slice.ContainsString(metadata.GetFinalizers(), o.constructFinalizerKey()) {
		return obj, nil
	}

	obj = obj.DeepCopyObject()
	metadata, err = meta.Accessor(obj)
	if err != nil {
		return nil, err
	}

	metadata.SetFinalizers(append(metadata.GetFinalizers(), o.constructFinalizerKey()))
	return o.objectClient.Update(metadata.GetName(), obj)
}

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

// Code generated by main. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	"github.com/bhojpur/host/pkg/labni/client"
	"github.com/bhojpur/host/pkg/labni/controller"
	"github.com/bhojpur/host/pkg/common/generic"
	v1 "k8s.io/api/storage/v1"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

type StorageClassHandler func(string, *v1.StorageClass) (*v1.StorageClass, error)

type StorageClassController interface {
	generic.ControllerMeta
	StorageClassClient

	OnChange(ctx context.Context, name string, sync StorageClassHandler)
	OnRemove(ctx context.Context, name string, sync StorageClassHandler)
	Enqueue(name string)
	EnqueueAfter(name string, duration time.Duration)

	Cache() StorageClassCache
}

type StorageClassClient interface {
	Create(*v1.StorageClass) (*v1.StorageClass, error)
	Update(*v1.StorageClass) (*v1.StorageClass, error)

	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*v1.StorageClass, error)
	List(opts metav1.ListOptions) (*v1.StorageClassList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.StorageClass, err error)
}

type StorageClassCache interface {
	Get(name string) (*v1.StorageClass, error)
	List(selector labels.Selector) ([]*v1.StorageClass, error)

	AddIndexer(indexName string, indexer StorageClassIndexer)
	GetByIndex(indexName, key string) ([]*v1.StorageClass, error)
}

type StorageClassIndexer func(obj *v1.StorageClass) ([]string, error)

type storageClassController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewStorageClassController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) StorageClassController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &storageClassController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromStorageClassHandlerToHandler(sync StorageClassHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v1.StorageClass
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v1.StorageClass))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *storageClassController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v1.StorageClass))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateStorageClassDeepCopyOnChange(client StorageClassClient, obj *v1.StorageClass, handler func(obj *v1.StorageClass) (*v1.StorageClass, error)) (*v1.StorageClass, error) {
	if obj == nil {
		return obj, nil
	}

	copyObj := obj.DeepCopy()
	newObj, err := handler(copyObj)
	if newObj != nil {
		copyObj = newObj
	}
	if obj.ResourceVersion == copyObj.ResourceVersion && !equality.Semantic.DeepEqual(obj, copyObj) {
		return client.Update(copyObj)
	}

	return copyObj, err
}

func (c *storageClassController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *storageClassController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *storageClassController) OnChange(ctx context.Context, name string, sync StorageClassHandler) {
	c.AddGenericHandler(ctx, name, FromStorageClassHandlerToHandler(sync))
}

func (c *storageClassController) OnRemove(ctx context.Context, name string, sync StorageClassHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromStorageClassHandlerToHandler(sync)))
}

func (c *storageClassController) Enqueue(name string) {
	c.controller.Enqueue("", name)
}

func (c *storageClassController) EnqueueAfter(name string, duration time.Duration) {
	c.controller.EnqueueAfter("", name, duration)
}

func (c *storageClassController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *storageClassController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *storageClassController) Cache() StorageClassCache {
	return &storageClassCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *storageClassController) Create(obj *v1.StorageClass) (*v1.StorageClass, error) {
	result := &v1.StorageClass{}
	return result, c.client.Create(context.TODO(), "", obj, result, metav1.CreateOptions{})
}

func (c *storageClassController) Update(obj *v1.StorageClass) (*v1.StorageClass, error) {
	result := &v1.StorageClass{}
	return result, c.client.Update(context.TODO(), "", obj, result, metav1.UpdateOptions{})
}

func (c *storageClassController) Delete(name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), "", name, *options)
}

func (c *storageClassController) Get(name string, options metav1.GetOptions) (*v1.StorageClass, error) {
	result := &v1.StorageClass{}
	return result, c.client.Get(context.TODO(), "", name, result, options)
}

func (c *storageClassController) List(opts metav1.ListOptions) (*v1.StorageClassList, error) {
	result := &v1.StorageClassList{}
	return result, c.client.List(context.TODO(), "", result, opts)
}

func (c *storageClassController) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), "", opts)
}

func (c *storageClassController) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (*v1.StorageClass, error) {
	result := &v1.StorageClass{}
	return result, c.client.Patch(context.TODO(), "", name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type storageClassCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *storageClassCache) Get(name string) (*v1.StorageClass, error) {
	obj, exists, err := c.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v1.StorageClass), nil
}

func (c *storageClassCache) List(selector labels.Selector) (ret []*v1.StorageClass, err error) {

	err = cache.ListAll(c.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.StorageClass))
	})

	return ret, err
}

func (c *storageClassCache) AddIndexer(indexName string, indexer StorageClassIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v1.StorageClass))
		},
	}))
}

func (c *storageClassCache) GetByIndex(indexName, key string) (result []*v1.StorageClass, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v1.StorageClass, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v1.StorageClass))
	}
	return result, nil
}
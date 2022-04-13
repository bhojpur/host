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

package v3

import (
	"context"
	"time"

	v3 "github.com/bhojpur/host/pkg/apis/project.bhojpur.net/v3"
	"github.com/bhojpur/host/pkg/common/generic"
	"github.com/bhojpur/host/pkg/labni/client"
	"github.com/bhojpur/host/pkg/labni/controller"
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

type NamespacedBasicAuthHandler func(string, *v3.NamespacedBasicAuth) (*v3.NamespacedBasicAuth, error)

type NamespacedBasicAuthController interface {
	generic.ControllerMeta
	NamespacedBasicAuthClient

	OnChange(ctx context.Context, name string, sync NamespacedBasicAuthHandler)
	OnRemove(ctx context.Context, name string, sync NamespacedBasicAuthHandler)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, duration time.Duration)

	Cache() NamespacedBasicAuthCache
}

type NamespacedBasicAuthClient interface {
	Create(*v3.NamespacedBasicAuth) (*v3.NamespacedBasicAuth, error)
	Update(*v3.NamespacedBasicAuth) (*v3.NamespacedBasicAuth, error)

	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v3.NamespacedBasicAuth, error)
	List(namespace string, opts metav1.ListOptions) (*v3.NamespacedBasicAuthList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v3.NamespacedBasicAuth, err error)
}

type NamespacedBasicAuthCache interface {
	Get(namespace, name string) (*v3.NamespacedBasicAuth, error)
	List(namespace string, selector labels.Selector) ([]*v3.NamespacedBasicAuth, error)

	AddIndexer(indexName string, indexer NamespacedBasicAuthIndexer)
	GetByIndex(indexName, key string) ([]*v3.NamespacedBasicAuth, error)
}

type NamespacedBasicAuthIndexer func(obj *v3.NamespacedBasicAuth) ([]string, error)

type namespacedBasicAuthController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewNamespacedBasicAuthController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) NamespacedBasicAuthController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &namespacedBasicAuthController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromNamespacedBasicAuthHandlerToHandler(sync NamespacedBasicAuthHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v3.NamespacedBasicAuth
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v3.NamespacedBasicAuth))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *namespacedBasicAuthController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v3.NamespacedBasicAuth))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateNamespacedBasicAuthDeepCopyOnChange(client NamespacedBasicAuthClient, obj *v3.NamespacedBasicAuth, handler func(obj *v3.NamespacedBasicAuth) (*v3.NamespacedBasicAuth, error)) (*v3.NamespacedBasicAuth, error) {
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

func (c *namespacedBasicAuthController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *namespacedBasicAuthController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *namespacedBasicAuthController) OnChange(ctx context.Context, name string, sync NamespacedBasicAuthHandler) {
	c.AddGenericHandler(ctx, name, FromNamespacedBasicAuthHandlerToHandler(sync))
}

func (c *namespacedBasicAuthController) OnRemove(ctx context.Context, name string, sync NamespacedBasicAuthHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromNamespacedBasicAuthHandlerToHandler(sync)))
}

func (c *namespacedBasicAuthController) Enqueue(namespace, name string) {
	c.controller.Enqueue(namespace, name)
}

func (c *namespacedBasicAuthController) EnqueueAfter(namespace, name string, duration time.Duration) {
	c.controller.EnqueueAfter(namespace, name, duration)
}

func (c *namespacedBasicAuthController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *namespacedBasicAuthController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *namespacedBasicAuthController) Cache() NamespacedBasicAuthCache {
	return &namespacedBasicAuthCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *namespacedBasicAuthController) Create(obj *v3.NamespacedBasicAuth) (*v3.NamespacedBasicAuth, error) {
	result := &v3.NamespacedBasicAuth{}
	return result, c.client.Create(context.TODO(), obj.Namespace, obj, result, metav1.CreateOptions{})
}

func (c *namespacedBasicAuthController) Update(obj *v3.NamespacedBasicAuth) (*v3.NamespacedBasicAuth, error) {
	result := &v3.NamespacedBasicAuth{}
	return result, c.client.Update(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *namespacedBasicAuthController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), namespace, name, *options)
}

func (c *namespacedBasicAuthController) Get(namespace, name string, options metav1.GetOptions) (*v3.NamespacedBasicAuth, error) {
	result := &v3.NamespacedBasicAuth{}
	return result, c.client.Get(context.TODO(), namespace, name, result, options)
}

func (c *namespacedBasicAuthController) List(namespace string, opts metav1.ListOptions) (*v3.NamespacedBasicAuthList, error) {
	result := &v3.NamespacedBasicAuthList{}
	return result, c.client.List(context.TODO(), namespace, result, opts)
}

func (c *namespacedBasicAuthController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), namespace, opts)
}

func (c *namespacedBasicAuthController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (*v3.NamespacedBasicAuth, error) {
	result := &v3.NamespacedBasicAuth{}
	return result, c.client.Patch(context.TODO(), namespace, name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type namespacedBasicAuthCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *namespacedBasicAuthCache) Get(namespace, name string) (*v3.NamespacedBasicAuth, error) {
	obj, exists, err := c.indexer.GetByKey(namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v3.NamespacedBasicAuth), nil
}

func (c *namespacedBasicAuthCache) List(namespace string, selector labels.Selector) (ret []*v3.NamespacedBasicAuth, err error) {

	err = cache.ListAllByNamespace(c.indexer, namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.NamespacedBasicAuth))
	})

	return ret, err
}

func (c *namespacedBasicAuthCache) AddIndexer(indexName string, indexer NamespacedBasicAuthIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v3.NamespacedBasicAuth))
		},
	}))
}

func (c *namespacedBasicAuthCache) GetByIndex(indexName, key string) (result []*v3.NamespacedBasicAuth, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v3.NamespacedBasicAuth, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v3.NamespacedBasicAuth))
	}
	return result, nil
}

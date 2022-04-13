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

type BasicAuthHandler func(string, *v3.BasicAuth) (*v3.BasicAuth, error)

type BasicAuthController interface {
	generic.ControllerMeta
	BasicAuthClient

	OnChange(ctx context.Context, name string, sync BasicAuthHandler)
	OnRemove(ctx context.Context, name string, sync BasicAuthHandler)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, duration time.Duration)

	Cache() BasicAuthCache
}

type BasicAuthClient interface {
	Create(*v3.BasicAuth) (*v3.BasicAuth, error)
	Update(*v3.BasicAuth) (*v3.BasicAuth, error)

	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v3.BasicAuth, error)
	List(namespace string, opts metav1.ListOptions) (*v3.BasicAuthList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v3.BasicAuth, err error)
}

type BasicAuthCache interface {
	Get(namespace, name string) (*v3.BasicAuth, error)
	List(namespace string, selector labels.Selector) ([]*v3.BasicAuth, error)

	AddIndexer(indexName string, indexer BasicAuthIndexer)
	GetByIndex(indexName, key string) ([]*v3.BasicAuth, error)
}

type BasicAuthIndexer func(obj *v3.BasicAuth) ([]string, error)

type basicAuthController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewBasicAuthController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) BasicAuthController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &basicAuthController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromBasicAuthHandlerToHandler(sync BasicAuthHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v3.BasicAuth
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v3.BasicAuth))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *basicAuthController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v3.BasicAuth))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateBasicAuthDeepCopyOnChange(client BasicAuthClient, obj *v3.BasicAuth, handler func(obj *v3.BasicAuth) (*v3.BasicAuth, error)) (*v3.BasicAuth, error) {
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

func (c *basicAuthController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *basicAuthController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *basicAuthController) OnChange(ctx context.Context, name string, sync BasicAuthHandler) {
	c.AddGenericHandler(ctx, name, FromBasicAuthHandlerToHandler(sync))
}

func (c *basicAuthController) OnRemove(ctx context.Context, name string, sync BasicAuthHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromBasicAuthHandlerToHandler(sync)))
}

func (c *basicAuthController) Enqueue(namespace, name string) {
	c.controller.Enqueue(namespace, name)
}

func (c *basicAuthController) EnqueueAfter(namespace, name string, duration time.Duration) {
	c.controller.EnqueueAfter(namespace, name, duration)
}

func (c *basicAuthController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *basicAuthController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *basicAuthController) Cache() BasicAuthCache {
	return &basicAuthCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *basicAuthController) Create(obj *v3.BasicAuth) (*v3.BasicAuth, error) {
	result := &v3.BasicAuth{}
	return result, c.client.Create(context.TODO(), obj.Namespace, obj, result, metav1.CreateOptions{})
}

func (c *basicAuthController) Update(obj *v3.BasicAuth) (*v3.BasicAuth, error) {
	result := &v3.BasicAuth{}
	return result, c.client.Update(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *basicAuthController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), namespace, name, *options)
}

func (c *basicAuthController) Get(namespace, name string, options metav1.GetOptions) (*v3.BasicAuth, error) {
	result := &v3.BasicAuth{}
	return result, c.client.Get(context.TODO(), namespace, name, result, options)
}

func (c *basicAuthController) List(namespace string, opts metav1.ListOptions) (*v3.BasicAuthList, error) {
	result := &v3.BasicAuthList{}
	return result, c.client.List(context.TODO(), namespace, result, opts)
}

func (c *basicAuthController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), namespace, opts)
}

func (c *basicAuthController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (*v3.BasicAuth, error) {
	result := &v3.BasicAuth{}
	return result, c.client.Patch(context.TODO(), namespace, name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type basicAuthCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *basicAuthCache) Get(namespace, name string) (*v3.BasicAuth, error) {
	obj, exists, err := c.indexer.GetByKey(namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v3.BasicAuth), nil
}

func (c *basicAuthCache) List(namespace string, selector labels.Selector) (ret []*v3.BasicAuth, err error) {

	err = cache.ListAllByNamespace(c.indexer, namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.BasicAuth))
	})

	return ret, err
}

func (c *basicAuthCache) AddIndexer(indexName string, indexer BasicAuthIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v3.BasicAuth))
		},
	}))
}

func (c *basicAuthCache) GetByIndex(indexName, key string) (result []*v3.BasicAuth, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v3.BasicAuth, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v3.BasicAuth))
	}
	return result, nil
}

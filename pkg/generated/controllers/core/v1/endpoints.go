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

	"github.com/bhojpur/host/pkg/common/generic"
	"github.com/bhojpur/host/pkg/labni/client"
	"github.com/bhojpur/host/pkg/labni/controller"
	v1 "k8s.io/api/core/v1"
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

type EndpointsHandler func(string, *v1.Endpoints) (*v1.Endpoints, error)

type EndpointsController interface {
	generic.ControllerMeta
	EndpointsClient

	OnChange(ctx context.Context, name string, sync EndpointsHandler)
	OnRemove(ctx context.Context, name string, sync EndpointsHandler)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, duration time.Duration)

	Cache() EndpointsCache
}

type EndpointsClient interface {
	Create(*v1.Endpoints) (*v1.Endpoints, error)
	Update(*v1.Endpoints) (*v1.Endpoints, error)

	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v1.Endpoints, error)
	List(namespace string, opts metav1.ListOptions) (*v1.EndpointsList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Endpoints, err error)
}

type EndpointsCache interface {
	Get(namespace, name string) (*v1.Endpoints, error)
	List(namespace string, selector labels.Selector) ([]*v1.Endpoints, error)

	AddIndexer(indexName string, indexer EndpointsIndexer)
	GetByIndex(indexName, key string) ([]*v1.Endpoints, error)
}

type EndpointsIndexer func(obj *v1.Endpoints) ([]string, error)

type endpointsController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewEndpointsController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) EndpointsController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &endpointsController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromEndpointsHandlerToHandler(sync EndpointsHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v1.Endpoints
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v1.Endpoints))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *endpointsController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v1.Endpoints))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateEndpointsDeepCopyOnChange(client EndpointsClient, obj *v1.Endpoints, handler func(obj *v1.Endpoints) (*v1.Endpoints, error)) (*v1.Endpoints, error) {
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

func (c *endpointsController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *endpointsController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *endpointsController) OnChange(ctx context.Context, name string, sync EndpointsHandler) {
	c.AddGenericHandler(ctx, name, FromEndpointsHandlerToHandler(sync))
}

func (c *endpointsController) OnRemove(ctx context.Context, name string, sync EndpointsHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromEndpointsHandlerToHandler(sync)))
}

func (c *endpointsController) Enqueue(namespace, name string) {
	c.controller.Enqueue(namespace, name)
}

func (c *endpointsController) EnqueueAfter(namespace, name string, duration time.Duration) {
	c.controller.EnqueueAfter(namespace, name, duration)
}

func (c *endpointsController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *endpointsController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *endpointsController) Cache() EndpointsCache {
	return &endpointsCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *endpointsController) Create(obj *v1.Endpoints) (*v1.Endpoints, error) {
	result := &v1.Endpoints{}
	return result, c.client.Create(context.TODO(), obj.Namespace, obj, result, metav1.CreateOptions{})
}

func (c *endpointsController) Update(obj *v1.Endpoints) (*v1.Endpoints, error) {
	result := &v1.Endpoints{}
	return result, c.client.Update(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *endpointsController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), namespace, name, *options)
}

func (c *endpointsController) Get(namespace, name string, options metav1.GetOptions) (*v1.Endpoints, error) {
	result := &v1.Endpoints{}
	return result, c.client.Get(context.TODO(), namespace, name, result, options)
}

func (c *endpointsController) List(namespace string, opts metav1.ListOptions) (*v1.EndpointsList, error) {
	result := &v1.EndpointsList{}
	return result, c.client.List(context.TODO(), namespace, result, opts)
}

func (c *endpointsController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), namespace, opts)
}

func (c *endpointsController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (*v1.Endpoints, error) {
	result := &v1.Endpoints{}
	return result, c.client.Patch(context.TODO(), namespace, name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type endpointsCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *endpointsCache) Get(namespace, name string) (*v1.Endpoints, error) {
	obj, exists, err := c.indexer.GetByKey(namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v1.Endpoints), nil
}

func (c *endpointsCache) List(namespace string, selector labels.Selector) (ret []*v1.Endpoints, err error) {

	err = cache.ListAllByNamespace(c.indexer, namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Endpoints))
	})

	return ret, err
}

func (c *endpointsCache) AddIndexer(indexName string, indexer EndpointsIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v1.Endpoints))
		},
	}))
}

func (c *endpointsCache) GetByIndex(indexName, key string) (result []*v1.Endpoints, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v1.Endpoints, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v1.Endpoints))
	}
	return result, nil
}

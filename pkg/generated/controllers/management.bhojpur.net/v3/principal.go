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

	v3 "github.com/bhojpur/host/pkg/apis/management.bhojpur.net/v3"
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

type PrincipalHandler func(string, *v3.Principal) (*v3.Principal, error)

type PrincipalController interface {
	generic.ControllerMeta
	PrincipalClient

	OnChange(ctx context.Context, name string, sync PrincipalHandler)
	OnRemove(ctx context.Context, name string, sync PrincipalHandler)
	Enqueue(name string)
	EnqueueAfter(name string, duration time.Duration)

	Cache() PrincipalCache
}

type PrincipalClient interface {
	Create(*v3.Principal) (*v3.Principal, error)
	Update(*v3.Principal) (*v3.Principal, error)

	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*v3.Principal, error)
	List(opts metav1.ListOptions) (*v3.PrincipalList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v3.Principal, err error)
}

type PrincipalCache interface {
	Get(name string) (*v3.Principal, error)
	List(selector labels.Selector) ([]*v3.Principal, error)

	AddIndexer(indexName string, indexer PrincipalIndexer)
	GetByIndex(indexName, key string) ([]*v3.Principal, error)
}

type PrincipalIndexer func(obj *v3.Principal) ([]string, error)

type principalController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewPrincipalController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) PrincipalController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &principalController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromPrincipalHandlerToHandler(sync PrincipalHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v3.Principal
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v3.Principal))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *principalController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v3.Principal))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdatePrincipalDeepCopyOnChange(client PrincipalClient, obj *v3.Principal, handler func(obj *v3.Principal) (*v3.Principal, error)) (*v3.Principal, error) {
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

func (c *principalController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *principalController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *principalController) OnChange(ctx context.Context, name string, sync PrincipalHandler) {
	c.AddGenericHandler(ctx, name, FromPrincipalHandlerToHandler(sync))
}

func (c *principalController) OnRemove(ctx context.Context, name string, sync PrincipalHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromPrincipalHandlerToHandler(sync)))
}

func (c *principalController) Enqueue(name string) {
	c.controller.Enqueue("", name)
}

func (c *principalController) EnqueueAfter(name string, duration time.Duration) {
	c.controller.EnqueueAfter("", name, duration)
}

func (c *principalController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *principalController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *principalController) Cache() PrincipalCache {
	return &principalCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *principalController) Create(obj *v3.Principal) (*v3.Principal, error) {
	result := &v3.Principal{}
	return result, c.client.Create(context.TODO(), "", obj, result, metav1.CreateOptions{})
}

func (c *principalController) Update(obj *v3.Principal) (*v3.Principal, error) {
	result := &v3.Principal{}
	return result, c.client.Update(context.TODO(), "", obj, result, metav1.UpdateOptions{})
}

func (c *principalController) Delete(name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), "", name, *options)
}

func (c *principalController) Get(name string, options metav1.GetOptions) (*v3.Principal, error) {
	result := &v3.Principal{}
	return result, c.client.Get(context.TODO(), "", name, result, options)
}

func (c *principalController) List(opts metav1.ListOptions) (*v3.PrincipalList, error) {
	result := &v3.PrincipalList{}
	return result, c.client.List(context.TODO(), "", result, opts)
}

func (c *principalController) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), "", opts)
}

func (c *principalController) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (*v3.Principal, error) {
	result := &v3.Principal{}
	return result, c.client.Patch(context.TODO(), "", name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type principalCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *principalCache) Get(name string) (*v3.Principal, error) {
	obj, exists, err := c.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v3.Principal), nil
}

func (c *principalCache) List(selector labels.Selector) (ret []*v3.Principal, err error) {

	err = cache.ListAll(c.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.Principal))
	})

	return ret, err
}

func (c *principalCache) AddIndexer(indexName string, indexer PrincipalIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v3.Principal))
		},
	}))
}

func (c *principalCache) GetByIndex(indexName, key string) (result []*v3.Principal, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v3.Principal, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v3.Principal))
	}
	return result, nil
}

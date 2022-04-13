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

type TemplateContentHandler func(string, *v3.TemplateContent) (*v3.TemplateContent, error)

type TemplateContentController interface {
	generic.ControllerMeta
	TemplateContentClient

	OnChange(ctx context.Context, name string, sync TemplateContentHandler)
	OnRemove(ctx context.Context, name string, sync TemplateContentHandler)
	Enqueue(name string)
	EnqueueAfter(name string, duration time.Duration)

	Cache() TemplateContentCache
}

type TemplateContentClient interface {
	Create(*v3.TemplateContent) (*v3.TemplateContent, error)
	Update(*v3.TemplateContent) (*v3.TemplateContent, error)

	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*v3.TemplateContent, error)
	List(opts metav1.ListOptions) (*v3.TemplateContentList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v3.TemplateContent, err error)
}

type TemplateContentCache interface {
	Get(name string) (*v3.TemplateContent, error)
	List(selector labels.Selector) ([]*v3.TemplateContent, error)

	AddIndexer(indexName string, indexer TemplateContentIndexer)
	GetByIndex(indexName, key string) ([]*v3.TemplateContent, error)
}

type TemplateContentIndexer func(obj *v3.TemplateContent) ([]string, error)

type templateContentController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewTemplateContentController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) TemplateContentController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &templateContentController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromTemplateContentHandlerToHandler(sync TemplateContentHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v3.TemplateContent
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v3.TemplateContent))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *templateContentController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v3.TemplateContent))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateTemplateContentDeepCopyOnChange(client TemplateContentClient, obj *v3.TemplateContent, handler func(obj *v3.TemplateContent) (*v3.TemplateContent, error)) (*v3.TemplateContent, error) {
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

func (c *templateContentController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *templateContentController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *templateContentController) OnChange(ctx context.Context, name string, sync TemplateContentHandler) {
	c.AddGenericHandler(ctx, name, FromTemplateContentHandlerToHandler(sync))
}

func (c *templateContentController) OnRemove(ctx context.Context, name string, sync TemplateContentHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromTemplateContentHandlerToHandler(sync)))
}

func (c *templateContentController) Enqueue(name string) {
	c.controller.Enqueue("", name)
}

func (c *templateContentController) EnqueueAfter(name string, duration time.Duration) {
	c.controller.EnqueueAfter("", name, duration)
}

func (c *templateContentController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *templateContentController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *templateContentController) Cache() TemplateContentCache {
	return &templateContentCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *templateContentController) Create(obj *v3.TemplateContent) (*v3.TemplateContent, error) {
	result := &v3.TemplateContent{}
	return result, c.client.Create(context.TODO(), "", obj, result, metav1.CreateOptions{})
}

func (c *templateContentController) Update(obj *v3.TemplateContent) (*v3.TemplateContent, error) {
	result := &v3.TemplateContent{}
	return result, c.client.Update(context.TODO(), "", obj, result, metav1.UpdateOptions{})
}

func (c *templateContentController) Delete(name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), "", name, *options)
}

func (c *templateContentController) Get(name string, options metav1.GetOptions) (*v3.TemplateContent, error) {
	result := &v3.TemplateContent{}
	return result, c.client.Get(context.TODO(), "", name, result, options)
}

func (c *templateContentController) List(opts metav1.ListOptions) (*v3.TemplateContentList, error) {
	result := &v3.TemplateContentList{}
	return result, c.client.List(context.TODO(), "", result, opts)
}

func (c *templateContentController) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), "", opts)
}

func (c *templateContentController) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (*v3.TemplateContent, error) {
	result := &v3.TemplateContent{}
	return result, c.client.Patch(context.TODO(), "", name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type templateContentCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *templateContentCache) Get(name string) (*v3.TemplateContent, error) {
	obj, exists, err := c.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v3.TemplateContent), nil
}

func (c *templateContentCache) List(selector labels.Selector) (ret []*v3.TemplateContent, err error) {

	err = cache.ListAll(c.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.TemplateContent))
	})

	return ret, err
}

func (c *templateContentCache) AddIndexer(indexName string, indexer TemplateContentIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v3.TemplateContent))
		},
	}))
}

func (c *templateContentCache) GetByIndex(indexName, key string) (result []*v3.TemplateContent, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v3.TemplateContent, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v3.TemplateContent))
	}
	return result, nil
}

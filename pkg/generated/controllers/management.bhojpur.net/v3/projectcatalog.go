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

type ProjectCatalogHandler func(string, *v3.ProjectCatalog) (*v3.ProjectCatalog, error)

type ProjectCatalogController interface {
	generic.ControllerMeta
	ProjectCatalogClient

	OnChange(ctx context.Context, name string, sync ProjectCatalogHandler)
	OnRemove(ctx context.Context, name string, sync ProjectCatalogHandler)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, duration time.Duration)

	Cache() ProjectCatalogCache
}

type ProjectCatalogClient interface {
	Create(*v3.ProjectCatalog) (*v3.ProjectCatalog, error)
	Update(*v3.ProjectCatalog) (*v3.ProjectCatalog, error)

	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v3.ProjectCatalog, error)
	List(namespace string, opts metav1.ListOptions) (*v3.ProjectCatalogList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v3.ProjectCatalog, err error)
}

type ProjectCatalogCache interface {
	Get(namespace, name string) (*v3.ProjectCatalog, error)
	List(namespace string, selector labels.Selector) ([]*v3.ProjectCatalog, error)

	AddIndexer(indexName string, indexer ProjectCatalogIndexer)
	GetByIndex(indexName, key string) ([]*v3.ProjectCatalog, error)
}

type ProjectCatalogIndexer func(obj *v3.ProjectCatalog) ([]string, error)

type projectCatalogController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewProjectCatalogController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) ProjectCatalogController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &projectCatalogController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromProjectCatalogHandlerToHandler(sync ProjectCatalogHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v3.ProjectCatalog
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v3.ProjectCatalog))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *projectCatalogController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v3.ProjectCatalog))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateProjectCatalogDeepCopyOnChange(client ProjectCatalogClient, obj *v3.ProjectCatalog, handler func(obj *v3.ProjectCatalog) (*v3.ProjectCatalog, error)) (*v3.ProjectCatalog, error) {
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

func (c *projectCatalogController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *projectCatalogController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *projectCatalogController) OnChange(ctx context.Context, name string, sync ProjectCatalogHandler) {
	c.AddGenericHandler(ctx, name, FromProjectCatalogHandlerToHandler(sync))
}

func (c *projectCatalogController) OnRemove(ctx context.Context, name string, sync ProjectCatalogHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromProjectCatalogHandlerToHandler(sync)))
}

func (c *projectCatalogController) Enqueue(namespace, name string) {
	c.controller.Enqueue(namespace, name)
}

func (c *projectCatalogController) EnqueueAfter(namespace, name string, duration time.Duration) {
	c.controller.EnqueueAfter(namespace, name, duration)
}

func (c *projectCatalogController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *projectCatalogController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *projectCatalogController) Cache() ProjectCatalogCache {
	return &projectCatalogCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *projectCatalogController) Create(obj *v3.ProjectCatalog) (*v3.ProjectCatalog, error) {
	result := &v3.ProjectCatalog{}
	return result, c.client.Create(context.TODO(), obj.Namespace, obj, result, metav1.CreateOptions{})
}

func (c *projectCatalogController) Update(obj *v3.ProjectCatalog) (*v3.ProjectCatalog, error) {
	result := &v3.ProjectCatalog{}
	return result, c.client.Update(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *projectCatalogController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), namespace, name, *options)
}

func (c *projectCatalogController) Get(namespace, name string, options metav1.GetOptions) (*v3.ProjectCatalog, error) {
	result := &v3.ProjectCatalog{}
	return result, c.client.Get(context.TODO(), namespace, name, result, options)
}

func (c *projectCatalogController) List(namespace string, opts metav1.ListOptions) (*v3.ProjectCatalogList, error) {
	result := &v3.ProjectCatalogList{}
	return result, c.client.List(context.TODO(), namespace, result, opts)
}

func (c *projectCatalogController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), namespace, opts)
}

func (c *projectCatalogController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (*v3.ProjectCatalog, error) {
	result := &v3.ProjectCatalog{}
	return result, c.client.Patch(context.TODO(), namespace, name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type projectCatalogCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *projectCatalogCache) Get(namespace, name string) (*v3.ProjectCatalog, error) {
	obj, exists, err := c.indexer.GetByKey(namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v3.ProjectCatalog), nil
}

func (c *projectCatalogCache) List(namespace string, selector labels.Selector) (ret []*v3.ProjectCatalog, err error) {

	err = cache.ListAllByNamespace(c.indexer, namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.ProjectCatalog))
	})

	return ret, err
}

func (c *projectCatalogCache) AddIndexer(indexName string, indexer ProjectCatalogIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v3.ProjectCatalog))
		},
	}))
}

func (c *projectCatalogCache) GetByIndex(indexName, key string) (result []*v3.ProjectCatalog, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v3.ProjectCatalog, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v3.ProjectCatalog))
	}
	return result, nil
}

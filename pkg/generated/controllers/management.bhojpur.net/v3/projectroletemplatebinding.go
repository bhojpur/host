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

type ProjectRoleTemplateBindingHandler func(string, *v3.ProjectRoleTemplateBinding) (*v3.ProjectRoleTemplateBinding, error)

type ProjectRoleTemplateBindingController interface {
	generic.ControllerMeta
	ProjectRoleTemplateBindingClient

	OnChange(ctx context.Context, name string, sync ProjectRoleTemplateBindingHandler)
	OnRemove(ctx context.Context, name string, sync ProjectRoleTemplateBindingHandler)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, duration time.Duration)

	Cache() ProjectRoleTemplateBindingCache
}

type ProjectRoleTemplateBindingClient interface {
	Create(*v3.ProjectRoleTemplateBinding) (*v3.ProjectRoleTemplateBinding, error)
	Update(*v3.ProjectRoleTemplateBinding) (*v3.ProjectRoleTemplateBinding, error)

	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v3.ProjectRoleTemplateBinding, error)
	List(namespace string, opts metav1.ListOptions) (*v3.ProjectRoleTemplateBindingList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v3.ProjectRoleTemplateBinding, err error)
}

type ProjectRoleTemplateBindingCache interface {
	Get(namespace, name string) (*v3.ProjectRoleTemplateBinding, error)
	List(namespace string, selector labels.Selector) ([]*v3.ProjectRoleTemplateBinding, error)

	AddIndexer(indexName string, indexer ProjectRoleTemplateBindingIndexer)
	GetByIndex(indexName, key string) ([]*v3.ProjectRoleTemplateBinding, error)
}

type ProjectRoleTemplateBindingIndexer func(obj *v3.ProjectRoleTemplateBinding) ([]string, error)

type projectRoleTemplateBindingController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewProjectRoleTemplateBindingController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) ProjectRoleTemplateBindingController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &projectRoleTemplateBindingController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromProjectRoleTemplateBindingHandlerToHandler(sync ProjectRoleTemplateBindingHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v3.ProjectRoleTemplateBinding
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v3.ProjectRoleTemplateBinding))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *projectRoleTemplateBindingController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v3.ProjectRoleTemplateBinding))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateProjectRoleTemplateBindingDeepCopyOnChange(client ProjectRoleTemplateBindingClient, obj *v3.ProjectRoleTemplateBinding, handler func(obj *v3.ProjectRoleTemplateBinding) (*v3.ProjectRoleTemplateBinding, error)) (*v3.ProjectRoleTemplateBinding, error) {
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

func (c *projectRoleTemplateBindingController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *projectRoleTemplateBindingController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *projectRoleTemplateBindingController) OnChange(ctx context.Context, name string, sync ProjectRoleTemplateBindingHandler) {
	c.AddGenericHandler(ctx, name, FromProjectRoleTemplateBindingHandlerToHandler(sync))
}

func (c *projectRoleTemplateBindingController) OnRemove(ctx context.Context, name string, sync ProjectRoleTemplateBindingHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromProjectRoleTemplateBindingHandlerToHandler(sync)))
}

func (c *projectRoleTemplateBindingController) Enqueue(namespace, name string) {
	c.controller.Enqueue(namespace, name)
}

func (c *projectRoleTemplateBindingController) EnqueueAfter(namespace, name string, duration time.Duration) {
	c.controller.EnqueueAfter(namespace, name, duration)
}

func (c *projectRoleTemplateBindingController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *projectRoleTemplateBindingController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *projectRoleTemplateBindingController) Cache() ProjectRoleTemplateBindingCache {
	return &projectRoleTemplateBindingCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *projectRoleTemplateBindingController) Create(obj *v3.ProjectRoleTemplateBinding) (*v3.ProjectRoleTemplateBinding, error) {
	result := &v3.ProjectRoleTemplateBinding{}
	return result, c.client.Create(context.TODO(), obj.Namespace, obj, result, metav1.CreateOptions{})
}

func (c *projectRoleTemplateBindingController) Update(obj *v3.ProjectRoleTemplateBinding) (*v3.ProjectRoleTemplateBinding, error) {
	result := &v3.ProjectRoleTemplateBinding{}
	return result, c.client.Update(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *projectRoleTemplateBindingController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), namespace, name, *options)
}

func (c *projectRoleTemplateBindingController) Get(namespace, name string, options metav1.GetOptions) (*v3.ProjectRoleTemplateBinding, error) {
	result := &v3.ProjectRoleTemplateBinding{}
	return result, c.client.Get(context.TODO(), namespace, name, result, options)
}

func (c *projectRoleTemplateBindingController) List(namespace string, opts metav1.ListOptions) (*v3.ProjectRoleTemplateBindingList, error) {
	result := &v3.ProjectRoleTemplateBindingList{}
	return result, c.client.List(context.TODO(), namespace, result, opts)
}

func (c *projectRoleTemplateBindingController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), namespace, opts)
}

func (c *projectRoleTemplateBindingController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (*v3.ProjectRoleTemplateBinding, error) {
	result := &v3.ProjectRoleTemplateBinding{}
	return result, c.client.Patch(context.TODO(), namespace, name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type projectRoleTemplateBindingCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *projectRoleTemplateBindingCache) Get(namespace, name string) (*v3.ProjectRoleTemplateBinding, error) {
	obj, exists, err := c.indexer.GetByKey(namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v3.ProjectRoleTemplateBinding), nil
}

func (c *projectRoleTemplateBindingCache) List(namespace string, selector labels.Selector) (ret []*v3.ProjectRoleTemplateBinding, err error) {

	err = cache.ListAllByNamespace(c.indexer, namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.ProjectRoleTemplateBinding))
	})

	return ret, err
}

func (c *projectRoleTemplateBindingCache) AddIndexer(indexName string, indexer ProjectRoleTemplateBindingIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v3.ProjectRoleTemplateBinding))
		},
	}))
}

func (c *projectRoleTemplateBindingCache) GetByIndex(indexName, key string) (result []*v3.ProjectRoleTemplateBinding, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v3.ProjectRoleTemplateBinding, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v3.ProjectRoleTemplateBinding))
	}
	return result, nil
}

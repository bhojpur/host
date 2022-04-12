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
	v1 "k8s.io/api/rbac/v1"
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

type RoleBindingHandler func(string, *v1.RoleBinding) (*v1.RoleBinding, error)

type RoleBindingController interface {
	generic.ControllerMeta
	RoleBindingClient

	OnChange(ctx context.Context, name string, sync RoleBindingHandler)
	OnRemove(ctx context.Context, name string, sync RoleBindingHandler)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, duration time.Duration)

	Cache() RoleBindingCache
}

type RoleBindingClient interface {
	Create(*v1.RoleBinding) (*v1.RoleBinding, error)
	Update(*v1.RoleBinding) (*v1.RoleBinding, error)

	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v1.RoleBinding, error)
	List(namespace string, opts metav1.ListOptions) (*v1.RoleBindingList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.RoleBinding, err error)
}

type RoleBindingCache interface {
	Get(namespace, name string) (*v1.RoleBinding, error)
	List(namespace string, selector labels.Selector) ([]*v1.RoleBinding, error)

	AddIndexer(indexName string, indexer RoleBindingIndexer)
	GetByIndex(indexName, key string) ([]*v1.RoleBinding, error)
}

type RoleBindingIndexer func(obj *v1.RoleBinding) ([]string, error)

type roleBindingController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewRoleBindingController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) RoleBindingController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &roleBindingController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromRoleBindingHandlerToHandler(sync RoleBindingHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v1.RoleBinding
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v1.RoleBinding))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *roleBindingController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v1.RoleBinding))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateRoleBindingDeepCopyOnChange(client RoleBindingClient, obj *v1.RoleBinding, handler func(obj *v1.RoleBinding) (*v1.RoleBinding, error)) (*v1.RoleBinding, error) {
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

func (c *roleBindingController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *roleBindingController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *roleBindingController) OnChange(ctx context.Context, name string, sync RoleBindingHandler) {
	c.AddGenericHandler(ctx, name, FromRoleBindingHandlerToHandler(sync))
}

func (c *roleBindingController) OnRemove(ctx context.Context, name string, sync RoleBindingHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromRoleBindingHandlerToHandler(sync)))
}

func (c *roleBindingController) Enqueue(namespace, name string) {
	c.controller.Enqueue(namespace, name)
}

func (c *roleBindingController) EnqueueAfter(namespace, name string, duration time.Duration) {
	c.controller.EnqueueAfter(namespace, name, duration)
}

func (c *roleBindingController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *roleBindingController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *roleBindingController) Cache() RoleBindingCache {
	return &roleBindingCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *roleBindingController) Create(obj *v1.RoleBinding) (*v1.RoleBinding, error) {
	result := &v1.RoleBinding{}
	return result, c.client.Create(context.TODO(), obj.Namespace, obj, result, metav1.CreateOptions{})
}

func (c *roleBindingController) Update(obj *v1.RoleBinding) (*v1.RoleBinding, error) {
	result := &v1.RoleBinding{}
	return result, c.client.Update(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *roleBindingController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), namespace, name, *options)
}

func (c *roleBindingController) Get(namespace, name string, options metav1.GetOptions) (*v1.RoleBinding, error) {
	result := &v1.RoleBinding{}
	return result, c.client.Get(context.TODO(), namespace, name, result, options)
}

func (c *roleBindingController) List(namespace string, opts metav1.ListOptions) (*v1.RoleBindingList, error) {
	result := &v1.RoleBindingList{}
	return result, c.client.List(context.TODO(), namespace, result, opts)
}

func (c *roleBindingController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), namespace, opts)
}

func (c *roleBindingController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (*v1.RoleBinding, error) {
	result := &v1.RoleBinding{}
	return result, c.client.Patch(context.TODO(), namespace, name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type roleBindingCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *roleBindingCache) Get(namespace, name string) (*v1.RoleBinding, error) {
	obj, exists, err := c.indexer.GetByKey(namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v1.RoleBinding), nil
}

func (c *roleBindingCache) List(namespace string, selector labels.Selector) (ret []*v1.RoleBinding, err error) {

	err = cache.ListAllByNamespace(c.indexer, namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.RoleBinding))
	})

	return ret, err
}

func (c *roleBindingCache) AddIndexer(indexName string, indexer RoleBindingIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v1.RoleBinding))
		},
	}))
}

func (c *roleBindingCache) GetByIndex(indexName, key string) (result []*v1.RoleBinding, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v1.RoleBinding, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v1.RoleBinding))
	}
	return result, nil
}
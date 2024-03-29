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

type GroupMemberHandler func(string, *v3.GroupMember) (*v3.GroupMember, error)

type GroupMemberController interface {
	generic.ControllerMeta
	GroupMemberClient

	OnChange(ctx context.Context, name string, sync GroupMemberHandler)
	OnRemove(ctx context.Context, name string, sync GroupMemberHandler)
	Enqueue(name string)
	EnqueueAfter(name string, duration time.Duration)

	Cache() GroupMemberCache
}

type GroupMemberClient interface {
	Create(*v3.GroupMember) (*v3.GroupMember, error)
	Update(*v3.GroupMember) (*v3.GroupMember, error)

	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*v3.GroupMember, error)
	List(opts metav1.ListOptions) (*v3.GroupMemberList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v3.GroupMember, err error)
}

type GroupMemberCache interface {
	Get(name string) (*v3.GroupMember, error)
	List(selector labels.Selector) ([]*v3.GroupMember, error)

	AddIndexer(indexName string, indexer GroupMemberIndexer)
	GetByIndex(indexName, key string) ([]*v3.GroupMember, error)
}

type GroupMemberIndexer func(obj *v3.GroupMember) ([]string, error)

type groupMemberController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewGroupMemberController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) GroupMemberController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &groupMemberController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromGroupMemberHandlerToHandler(sync GroupMemberHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v3.GroupMember
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v3.GroupMember))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *groupMemberController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v3.GroupMember))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateGroupMemberDeepCopyOnChange(client GroupMemberClient, obj *v3.GroupMember, handler func(obj *v3.GroupMember) (*v3.GroupMember, error)) (*v3.GroupMember, error) {
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

func (c *groupMemberController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *groupMemberController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *groupMemberController) OnChange(ctx context.Context, name string, sync GroupMemberHandler) {
	c.AddGenericHandler(ctx, name, FromGroupMemberHandlerToHandler(sync))
}

func (c *groupMemberController) OnRemove(ctx context.Context, name string, sync GroupMemberHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromGroupMemberHandlerToHandler(sync)))
}

func (c *groupMemberController) Enqueue(name string) {
	c.controller.Enqueue("", name)
}

func (c *groupMemberController) EnqueueAfter(name string, duration time.Duration) {
	c.controller.EnqueueAfter("", name, duration)
}

func (c *groupMemberController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *groupMemberController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *groupMemberController) Cache() GroupMemberCache {
	return &groupMemberCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *groupMemberController) Create(obj *v3.GroupMember) (*v3.GroupMember, error) {
	result := &v3.GroupMember{}
	return result, c.client.Create(context.TODO(), "", obj, result, metav1.CreateOptions{})
}

func (c *groupMemberController) Update(obj *v3.GroupMember) (*v3.GroupMember, error) {
	result := &v3.GroupMember{}
	return result, c.client.Update(context.TODO(), "", obj, result, metav1.UpdateOptions{})
}

func (c *groupMemberController) Delete(name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), "", name, *options)
}

func (c *groupMemberController) Get(name string, options metav1.GetOptions) (*v3.GroupMember, error) {
	result := &v3.GroupMember{}
	return result, c.client.Get(context.TODO(), "", name, result, options)
}

func (c *groupMemberController) List(opts metav1.ListOptions) (*v3.GroupMemberList, error) {
	result := &v3.GroupMemberList{}
	return result, c.client.List(context.TODO(), "", result, opts)
}

func (c *groupMemberController) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), "", opts)
}

func (c *groupMemberController) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (*v3.GroupMember, error) {
	result := &v3.GroupMember{}
	return result, c.client.Patch(context.TODO(), "", name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type groupMemberCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *groupMemberCache) Get(name string) (*v3.GroupMember, error) {
	obj, exists, err := c.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v3.GroupMember), nil
}

func (c *groupMemberCache) List(selector labels.Selector) (ret []*v3.GroupMember, err error) {

	err = cache.ListAll(c.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.GroupMember))
	})

	return ret, err
}

func (c *groupMemberCache) AddIndexer(indexName string, indexer GroupMemberIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v3.GroupMember))
		},
	}))
}

func (c *groupMemberCache) GetByIndex(indexName, key string) (result []*v3.GroupMember, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v3.GroupMember, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v3.GroupMember))
	}
	return result, nil
}

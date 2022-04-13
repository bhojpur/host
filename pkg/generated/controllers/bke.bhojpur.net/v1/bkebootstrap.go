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

	v1 "github.com/bhojpur/host/pkg/apis/bke.bhojpur.net/v1"
	"github.com/bhojpur/host/pkg/common/apply"
	"github.com/bhojpur/host/pkg/common/condition"
	"github.com/bhojpur/host/pkg/common/generic"
	"github.com/bhojpur/host/pkg/common/kv"
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

type BKEBootstrapHandler func(string, *v1.BKEBootstrap) (*v1.BKEBootstrap, error)

type BKEBootstrapController interface {
	generic.ControllerMeta
	BKEBootstrapClient

	OnChange(ctx context.Context, name string, sync BKEBootstrapHandler)
	OnRemove(ctx context.Context, name string, sync BKEBootstrapHandler)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, duration time.Duration)

	Cache() BKEBootstrapCache
}

type BKEBootstrapClient interface {
	Create(*v1.BKEBootstrap) (*v1.BKEBootstrap, error)
	Update(*v1.BKEBootstrap) (*v1.BKEBootstrap, error)
	UpdateStatus(*v1.BKEBootstrap) (*v1.BKEBootstrap, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v1.BKEBootstrap, error)
	List(namespace string, opts metav1.ListOptions) (*v1.BKEBootstrapList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.BKEBootstrap, err error)
}

type BKEBootstrapCache interface {
	Get(namespace, name string) (*v1.BKEBootstrap, error)
	List(namespace string, selector labels.Selector) ([]*v1.BKEBootstrap, error)

	AddIndexer(indexName string, indexer BKEBootstrapIndexer)
	GetByIndex(indexName, key string) ([]*v1.BKEBootstrap, error)
}

type BKEBootstrapIndexer func(obj *v1.BKEBootstrap) ([]string, error)

type bKEBootstrapController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewBKEBootstrapController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) BKEBootstrapController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &bKEBootstrapController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromBKEBootstrapHandlerToHandler(sync BKEBootstrapHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v1.BKEBootstrap
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v1.BKEBootstrap))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *bKEBootstrapController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v1.BKEBootstrap))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateBKEBootstrapDeepCopyOnChange(client BKEBootstrapClient, obj *v1.BKEBootstrap, handler func(obj *v1.BKEBootstrap) (*v1.BKEBootstrap, error)) (*v1.BKEBootstrap, error) {
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

func (c *bKEBootstrapController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *bKEBootstrapController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *bKEBootstrapController) OnChange(ctx context.Context, name string, sync BKEBootstrapHandler) {
	c.AddGenericHandler(ctx, name, FromBKEBootstrapHandlerToHandler(sync))
}

func (c *bKEBootstrapController) OnRemove(ctx context.Context, name string, sync BKEBootstrapHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromBKEBootstrapHandlerToHandler(sync)))
}

func (c *bKEBootstrapController) Enqueue(namespace, name string) {
	c.controller.Enqueue(namespace, name)
}

func (c *bKEBootstrapController) EnqueueAfter(namespace, name string, duration time.Duration) {
	c.controller.EnqueueAfter(namespace, name, duration)
}

func (c *bKEBootstrapController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *bKEBootstrapController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *bKEBootstrapController) Cache() BKEBootstrapCache {
	return &bKEBootstrapCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *bKEBootstrapController) Create(obj *v1.BKEBootstrap) (*v1.BKEBootstrap, error) {
	result := &v1.BKEBootstrap{}
	return result, c.client.Create(context.TODO(), obj.Namespace, obj, result, metav1.CreateOptions{})
}

func (c *bKEBootstrapController) Update(obj *v1.BKEBootstrap) (*v1.BKEBootstrap, error) {
	result := &v1.BKEBootstrap{}
	return result, c.client.Update(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *bKEBootstrapController) UpdateStatus(obj *v1.BKEBootstrap) (*v1.BKEBootstrap, error) {
	result := &v1.BKEBootstrap{}
	return result, c.client.UpdateStatus(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *bKEBootstrapController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), namespace, name, *options)
}

func (c *bKEBootstrapController) Get(namespace, name string, options metav1.GetOptions) (*v1.BKEBootstrap, error) {
	result := &v1.BKEBootstrap{}
	return result, c.client.Get(context.TODO(), namespace, name, result, options)
}

func (c *bKEBootstrapController) List(namespace string, opts metav1.ListOptions) (*v1.BKEBootstrapList, error) {
	result := &v1.BKEBootstrapList{}
	return result, c.client.List(context.TODO(), namespace, result, opts)
}

func (c *bKEBootstrapController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), namespace, opts)
}

func (c *bKEBootstrapController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (*v1.BKEBootstrap, error) {
	result := &v1.BKEBootstrap{}
	return result, c.client.Patch(context.TODO(), namespace, name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type bKEBootstrapCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *bKEBootstrapCache) Get(namespace, name string) (*v1.BKEBootstrap, error) {
	obj, exists, err := c.indexer.GetByKey(namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v1.BKEBootstrap), nil
}

func (c *bKEBootstrapCache) List(namespace string, selector labels.Selector) (ret []*v1.BKEBootstrap, err error) {

	err = cache.ListAllByNamespace(c.indexer, namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.BKEBootstrap))
	})

	return ret, err
}

func (c *bKEBootstrapCache) AddIndexer(indexName string, indexer BKEBootstrapIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v1.BKEBootstrap))
		},
	}))
}

func (c *bKEBootstrapCache) GetByIndex(indexName, key string) (result []*v1.BKEBootstrap, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v1.BKEBootstrap, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v1.BKEBootstrap))
	}
	return result, nil
}

type BKEBootstrapStatusHandler func(obj *v1.BKEBootstrap, status v1.BKEBootstrapStatus) (v1.BKEBootstrapStatus, error)

type BKEBootstrapGeneratingHandler func(obj *v1.BKEBootstrap, status v1.BKEBootstrapStatus) ([]runtime.Object, v1.BKEBootstrapStatus, error)

func RegisterBKEBootstrapStatusHandler(ctx context.Context, controller BKEBootstrapController, condition condition.Cond, name string, handler BKEBootstrapStatusHandler) {
	statusHandler := &bKEBootstrapStatusHandler{
		client:    controller,
		condition: condition,
		handler:   handler,
	}
	controller.AddGenericHandler(ctx, name, FromBKEBootstrapHandlerToHandler(statusHandler.sync))
}

func RegisterBKEBootstrapGeneratingHandler(ctx context.Context, controller BKEBootstrapController, apply apply.Apply,
	condition condition.Cond, name string, handler BKEBootstrapGeneratingHandler, opts *generic.GeneratingHandlerOptions) {
	statusHandler := &bKEBootstrapGeneratingHandler{
		BKEBootstrapGeneratingHandler: handler,
		apply:                         apply,
		name:                          name,
		gvk:                           controller.GroupVersionKind(),
	}
	if opts != nil {
		statusHandler.opts = *opts
	}
	controller.OnChange(ctx, name, statusHandler.Remove)
	RegisterBKEBootstrapStatusHandler(ctx, controller, condition, name, statusHandler.Handle)
}

type bKEBootstrapStatusHandler struct {
	client    BKEBootstrapClient
	condition condition.Cond
	handler   BKEBootstrapStatusHandler
}

func (a *bKEBootstrapStatusHandler) sync(key string, obj *v1.BKEBootstrap) (*v1.BKEBootstrap, error) {
	if obj == nil {
		return obj, nil
	}

	origStatus := obj.Status.DeepCopy()
	obj = obj.DeepCopy()
	newStatus, err := a.handler(obj, obj.Status)
	if err != nil {
		// Revert to old status on error
		newStatus = *origStatus.DeepCopy()
	}

	if a.condition != "" {
		if errors.IsConflict(err) {
			a.condition.SetError(&newStatus, "", nil)
		} else {
			a.condition.SetError(&newStatus, "", err)
		}
	}
	if !equality.Semantic.DeepEqual(origStatus, &newStatus) {
		if a.condition != "" {
			// Since status has changed, update the lastUpdatedTime
			a.condition.LastUpdated(&newStatus, time.Now().UTC().Format(time.RFC3339))
		}

		var newErr error
		obj.Status = newStatus
		newObj, newErr := a.client.UpdateStatus(obj)
		if err == nil {
			err = newErr
		}
		if newErr == nil {
			obj = newObj
		}
	}
	return obj, err
}

type bKEBootstrapGeneratingHandler struct {
	BKEBootstrapGeneratingHandler
	apply apply.Apply
	opts  generic.GeneratingHandlerOptions
	gvk   schema.GroupVersionKind
	name  string
}

func (a *bKEBootstrapGeneratingHandler) Remove(key string, obj *v1.BKEBootstrap) (*v1.BKEBootstrap, error) {
	if obj != nil {
		return obj, nil
	}

	obj = &v1.BKEBootstrap{}
	obj.Namespace, obj.Name = kv.RSplit(key, "/")
	obj.SetGroupVersionKind(a.gvk)

	return nil, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects()
}

func (a *bKEBootstrapGeneratingHandler) Handle(obj *v1.BKEBootstrap, status v1.BKEBootstrapStatus) (v1.BKEBootstrapStatus, error) {
	if !obj.DeletionTimestamp.IsZero() {
		return status, nil
	}

	objs, newStatus, err := a.BKEBootstrapGeneratingHandler(obj, status)
	if err != nil {
		return newStatus, err
	}

	return newStatus, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects(objs...)
}

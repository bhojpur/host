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
	"github.com/bhojpur/host/pkg/common/apply"
	"github.com/bhojpur/host/pkg/common/condition"
	"github.com/bhojpur/host/pkg/common/generic"
	"github.com/bhojpur/host/pkg/common/kv"
	v1 "k8s.io/api/apps/v1"
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

type DaemonSetHandler func(string, *v1.DaemonSet) (*v1.DaemonSet, error)

type DaemonSetController interface {
	generic.ControllerMeta
	DaemonSetClient

	OnChange(ctx context.Context, name string, sync DaemonSetHandler)
	OnRemove(ctx context.Context, name string, sync DaemonSetHandler)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, duration time.Duration)

	Cache() DaemonSetCache
}

type DaemonSetClient interface {
	Create(*v1.DaemonSet) (*v1.DaemonSet, error)
	Update(*v1.DaemonSet) (*v1.DaemonSet, error)
	UpdateStatus(*v1.DaemonSet) (*v1.DaemonSet, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v1.DaemonSet, error)
	List(namespace string, opts metav1.ListOptions) (*v1.DaemonSetList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.DaemonSet, err error)
}

type DaemonSetCache interface {
	Get(namespace, name string) (*v1.DaemonSet, error)
	List(namespace string, selector labels.Selector) ([]*v1.DaemonSet, error)

	AddIndexer(indexName string, indexer DaemonSetIndexer)
	GetByIndex(indexName, key string) ([]*v1.DaemonSet, error)
}

type DaemonSetIndexer func(obj *v1.DaemonSet) ([]string, error)

type daemonSetController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewDaemonSetController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) DaemonSetController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &daemonSetController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromDaemonSetHandlerToHandler(sync DaemonSetHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v1.DaemonSet
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v1.DaemonSet))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *daemonSetController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v1.DaemonSet))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateDaemonSetDeepCopyOnChange(client DaemonSetClient, obj *v1.DaemonSet, handler func(obj *v1.DaemonSet) (*v1.DaemonSet, error)) (*v1.DaemonSet, error) {
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

func (c *daemonSetController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *daemonSetController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *daemonSetController) OnChange(ctx context.Context, name string, sync DaemonSetHandler) {
	c.AddGenericHandler(ctx, name, FromDaemonSetHandlerToHandler(sync))
}

func (c *daemonSetController) OnRemove(ctx context.Context, name string, sync DaemonSetHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromDaemonSetHandlerToHandler(sync)))
}

func (c *daemonSetController) Enqueue(namespace, name string) {
	c.controller.Enqueue(namespace, name)
}

func (c *daemonSetController) EnqueueAfter(namespace, name string, duration time.Duration) {
	c.controller.EnqueueAfter(namespace, name, duration)
}

func (c *daemonSetController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *daemonSetController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *daemonSetController) Cache() DaemonSetCache {
	return &daemonSetCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *daemonSetController) Create(obj *v1.DaemonSet) (*v1.DaemonSet, error) {
	result := &v1.DaemonSet{}
	return result, c.client.Create(context.TODO(), obj.Namespace, obj, result, metav1.CreateOptions{})
}

func (c *daemonSetController) Update(obj *v1.DaemonSet) (*v1.DaemonSet, error) {
	result := &v1.DaemonSet{}
	return result, c.client.Update(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *daemonSetController) UpdateStatus(obj *v1.DaemonSet) (*v1.DaemonSet, error) {
	result := &v1.DaemonSet{}
	return result, c.client.UpdateStatus(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *daemonSetController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), namespace, name, *options)
}

func (c *daemonSetController) Get(namespace, name string, options metav1.GetOptions) (*v1.DaemonSet, error) {
	result := &v1.DaemonSet{}
	return result, c.client.Get(context.TODO(), namespace, name, result, options)
}

func (c *daemonSetController) List(namespace string, opts metav1.ListOptions) (*v1.DaemonSetList, error) {
	result := &v1.DaemonSetList{}
	return result, c.client.List(context.TODO(), namespace, result, opts)
}

func (c *daemonSetController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), namespace, opts)
}

func (c *daemonSetController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (*v1.DaemonSet, error) {
	result := &v1.DaemonSet{}
	return result, c.client.Patch(context.TODO(), namespace, name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type daemonSetCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *daemonSetCache) Get(namespace, name string) (*v1.DaemonSet, error) {
	obj, exists, err := c.indexer.GetByKey(namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v1.DaemonSet), nil
}

func (c *daemonSetCache) List(namespace string, selector labels.Selector) (ret []*v1.DaemonSet, err error) {

	err = cache.ListAllByNamespace(c.indexer, namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.DaemonSet))
	})

	return ret, err
}

func (c *daemonSetCache) AddIndexer(indexName string, indexer DaemonSetIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v1.DaemonSet))
		},
	}))
}

func (c *daemonSetCache) GetByIndex(indexName, key string) (result []*v1.DaemonSet, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v1.DaemonSet, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v1.DaemonSet))
	}
	return result, nil
}

type DaemonSetStatusHandler func(obj *v1.DaemonSet, status v1.DaemonSetStatus) (v1.DaemonSetStatus, error)

type DaemonSetGeneratingHandler func(obj *v1.DaemonSet, status v1.DaemonSetStatus) ([]runtime.Object, v1.DaemonSetStatus, error)

func RegisterDaemonSetStatusHandler(ctx context.Context, controller DaemonSetController, condition condition.Cond, name string, handler DaemonSetStatusHandler) {
	statusHandler := &daemonSetStatusHandler{
		client:    controller,
		condition: condition,
		handler:   handler,
	}
	controller.AddGenericHandler(ctx, name, FromDaemonSetHandlerToHandler(statusHandler.sync))
}

func RegisterDaemonSetGeneratingHandler(ctx context.Context, controller DaemonSetController, apply apply.Apply,
	condition condition.Cond, name string, handler DaemonSetGeneratingHandler, opts *generic.GeneratingHandlerOptions) {
	statusHandler := &daemonSetGeneratingHandler{
		DaemonSetGeneratingHandler: handler,
		apply:                      apply,
		name:                       name,
		gvk:                        controller.GroupVersionKind(),
	}
	if opts != nil {
		statusHandler.opts = *opts
	}
	controller.OnChange(ctx, name, statusHandler.Remove)
	RegisterDaemonSetStatusHandler(ctx, controller, condition, name, statusHandler.Handle)
}

type daemonSetStatusHandler struct {
	client    DaemonSetClient
	condition condition.Cond
	handler   DaemonSetStatusHandler
}

func (a *daemonSetStatusHandler) sync(key string, obj *v1.DaemonSet) (*v1.DaemonSet, error) {
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

type daemonSetGeneratingHandler struct {
	DaemonSetGeneratingHandler
	apply apply.Apply
	opts  generic.GeneratingHandlerOptions
	gvk   schema.GroupVersionKind
	name  string
}

func (a *daemonSetGeneratingHandler) Remove(key string, obj *v1.DaemonSet) (*v1.DaemonSet, error) {
	if obj != nil {
		return obj, nil
	}

	obj = &v1.DaemonSet{}
	obj.Namespace, obj.Name = kv.RSplit(key, "/")
	obj.SetGroupVersionKind(a.gvk)

	return nil, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects()
}

func (a *daemonSetGeneratingHandler) Handle(obj *v1.DaemonSet, status v1.DaemonSetStatus) (v1.DaemonSetStatus, error) {
	if !obj.DeletionTimestamp.IsZero() {
		return status, nil
	}

	objs, newStatus, err := a.DaemonSetGeneratingHandler(obj, status)
	if err != nil {
		return newStatus, err
	}

	return newStatus, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects(objs...)
}
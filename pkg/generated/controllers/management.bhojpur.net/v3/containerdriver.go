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

type ContainerDriverHandler func(string, *v3.ContainerDriver) (*v3.ContainerDriver, error)

type ContainerDriverController interface {
	generic.ControllerMeta
	ContainerDriverClient

	OnChange(ctx context.Context, name string, sync ContainerDriverHandler)
	OnRemove(ctx context.Context, name string, sync ContainerDriverHandler)
	Enqueue(name string)
	EnqueueAfter(name string, duration time.Duration)

	Cache() ContainerDriverCache
}

type ContainerDriverClient interface {
	Create(*v3.ContainerDriver) (*v3.ContainerDriver, error)
	Update(*v3.ContainerDriver) (*v3.ContainerDriver, error)
	UpdateStatus(*v3.ContainerDriver) (*v3.ContainerDriver, error)
	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*v3.ContainerDriver, error)
	List(opts metav1.ListOptions) (*v3.ContainerDriverList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v3.ContainerDriver, err error)
}

type ContainerDriverCache interface {
	Get(name string) (*v3.ContainerDriver, error)
	List(selector labels.Selector) ([]*v3.ContainerDriver, error)

	AddIndexer(indexName string, indexer ContainerDriverIndexer)
	GetByIndex(indexName, key string) ([]*v3.ContainerDriver, error)
}

type ContainerDriverIndexer func(obj *v3.ContainerDriver) ([]string, error)

type containerDriverController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewContainerDriverController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) ContainerDriverController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &containerDriverController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromContainerDriverHandlerToHandler(sync ContainerDriverHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v3.ContainerDriver
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v3.ContainerDriver))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *containerDriverController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v3.ContainerDriver))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateContainerDriverDeepCopyOnChange(client ContainerDriverClient, obj *v3.ContainerDriver, handler func(obj *v3.ContainerDriver) (*v3.ContainerDriver, error)) (*v3.ContainerDriver, error) {
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

func (c *containerDriverController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *containerDriverController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *containerDriverController) OnChange(ctx context.Context, name string, sync ContainerDriverHandler) {
	c.AddGenericHandler(ctx, name, FromContainerDriverHandlerToHandler(sync))
}

func (c *containerDriverController) OnRemove(ctx context.Context, name string, sync ContainerDriverHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromContainerDriverHandlerToHandler(sync)))
}

func (c *containerDriverController) Enqueue(name string) {
	c.controller.Enqueue("", name)
}

func (c *containerDriverController) EnqueueAfter(name string, duration time.Duration) {
	c.controller.EnqueueAfter("", name, duration)
}

func (c *containerDriverController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *containerDriverController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *containerDriverController) Cache() ContainerDriverCache {
	return &containerDriverCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *containerDriverController) Create(obj *v3.ContainerDriver) (*v3.ContainerDriver, error) {
	result := &v3.ContainerDriver{}
	return result, c.client.Create(context.TODO(), "", obj, result, metav1.CreateOptions{})
}

func (c *containerDriverController) Update(obj *v3.ContainerDriver) (*v3.ContainerDriver, error) {
	result := &v3.ContainerDriver{}
	return result, c.client.Update(context.TODO(), "", obj, result, metav1.UpdateOptions{})
}

func (c *containerDriverController) UpdateStatus(obj *v3.ContainerDriver) (*v3.ContainerDriver, error) {
	result := &v3.ContainerDriver{}
	return result, c.client.UpdateStatus(context.TODO(), "", obj, result, metav1.UpdateOptions{})
}

func (c *containerDriverController) Delete(name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), "", name, *options)
}

func (c *containerDriverController) Get(name string, options metav1.GetOptions) (*v3.ContainerDriver, error) {
	result := &v3.ContainerDriver{}
	return result, c.client.Get(context.TODO(), "", name, result, options)
}

func (c *containerDriverController) List(opts metav1.ListOptions) (*v3.ContainerDriverList, error) {
	result := &v3.ContainerDriverList{}
	return result, c.client.List(context.TODO(), "", result, opts)
}

func (c *containerDriverController) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), "", opts)
}

func (c *containerDriverController) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (*v3.ContainerDriver, error) {
	result := &v3.ContainerDriver{}
	return result, c.client.Patch(context.TODO(), "", name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type containerDriverCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *containerDriverCache) Get(name string) (*v3.ContainerDriver, error) {
	obj, exists, err := c.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v3.ContainerDriver), nil
}

func (c *containerDriverCache) List(selector labels.Selector) (ret []*v3.ContainerDriver, err error) {

	err = cache.ListAll(c.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.ContainerDriver))
	})

	return ret, err
}

func (c *containerDriverCache) AddIndexer(indexName string, indexer ContainerDriverIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v3.ContainerDriver))
		},
	}))
}

func (c *containerDriverCache) GetByIndex(indexName, key string) (result []*v3.ContainerDriver, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v3.ContainerDriver, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v3.ContainerDriver))
	}
	return result, nil
}

type ContainerDriverStatusHandler func(obj *v3.ContainerDriver, status v3.ContainerDriverStatus) (v3.ContainerDriverStatus, error)

type ContainerDriverGeneratingHandler func(obj *v3.ContainerDriver, status v3.ContainerDriverStatus) ([]runtime.Object, v3.ContainerDriverStatus, error)

func RegisterContainerDriverStatusHandler(ctx context.Context, controller ContainerDriverController, condition condition.Cond, name string, handler ContainerDriverStatusHandler) {
	statusHandler := &containerDriverStatusHandler{
		client:    controller,
		condition: condition,
		handler:   handler,
	}
	controller.AddGenericHandler(ctx, name, FromContainerDriverHandlerToHandler(statusHandler.sync))
}

func RegisterContainerDriverGeneratingHandler(ctx context.Context, controller ContainerDriverController, apply apply.Apply,
	condition condition.Cond, name string, handler ContainerDriverGeneratingHandler, opts *generic.GeneratingHandlerOptions) {
	statusHandler := &containerDriverGeneratingHandler{
		ContainerDriverGeneratingHandler: handler,
		apply:                            apply,
		name:                             name,
		gvk:                              controller.GroupVersionKind(),
	}
	if opts != nil {
		statusHandler.opts = *opts
	}
	controller.OnChange(ctx, name, statusHandler.Remove)
	RegisterContainerDriverStatusHandler(ctx, controller, condition, name, statusHandler.Handle)
}

type containerDriverStatusHandler struct {
	client    ContainerDriverClient
	condition condition.Cond
	handler   ContainerDriverStatusHandler
}

func (a *containerDriverStatusHandler) sync(key string, obj *v3.ContainerDriver) (*v3.ContainerDriver, error) {
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

type containerDriverGeneratingHandler struct {
	ContainerDriverGeneratingHandler
	apply apply.Apply
	opts  generic.GeneratingHandlerOptions
	gvk   schema.GroupVersionKind
	name  string
}

func (a *containerDriverGeneratingHandler) Remove(key string, obj *v3.ContainerDriver) (*v3.ContainerDriver, error) {
	if obj != nil {
		return obj, nil
	}

	obj = &v3.ContainerDriver{}
	obj.Namespace, obj.Name = kv.RSplit(key, "/")
	obj.SetGroupVersionKind(a.gvk)

	return nil, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects()
}

func (a *containerDriverGeneratingHandler) Handle(obj *v3.ContainerDriver, status v3.ContainerDriverStatus) (v3.ContainerDriverStatus, error) {
	if !obj.DeletionTimestamp.IsZero() {
		return status, nil
	}

	objs, newStatus, err := a.ContainerDriverGeneratingHandler(obj, status)
	if err != nil {
		return newStatus, err
	}

	return newStatus, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects(objs...)
}

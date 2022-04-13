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

type ProjectLoggingHandler func(string, *v3.ProjectLogging) (*v3.ProjectLogging, error)

type ProjectLoggingController interface {
	generic.ControllerMeta
	ProjectLoggingClient

	OnChange(ctx context.Context, name string, sync ProjectLoggingHandler)
	OnRemove(ctx context.Context, name string, sync ProjectLoggingHandler)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, duration time.Duration)

	Cache() ProjectLoggingCache
}

type ProjectLoggingClient interface {
	Create(*v3.ProjectLogging) (*v3.ProjectLogging, error)
	Update(*v3.ProjectLogging) (*v3.ProjectLogging, error)
	UpdateStatus(*v3.ProjectLogging) (*v3.ProjectLogging, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v3.ProjectLogging, error)
	List(namespace string, opts metav1.ListOptions) (*v3.ProjectLoggingList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v3.ProjectLogging, err error)
}

type ProjectLoggingCache interface {
	Get(namespace, name string) (*v3.ProjectLogging, error)
	List(namespace string, selector labels.Selector) ([]*v3.ProjectLogging, error)

	AddIndexer(indexName string, indexer ProjectLoggingIndexer)
	GetByIndex(indexName, key string) ([]*v3.ProjectLogging, error)
}

type ProjectLoggingIndexer func(obj *v3.ProjectLogging) ([]string, error)

type projectLoggingController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewProjectLoggingController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) ProjectLoggingController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &projectLoggingController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromProjectLoggingHandlerToHandler(sync ProjectLoggingHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v3.ProjectLogging
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v3.ProjectLogging))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *projectLoggingController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v3.ProjectLogging))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateProjectLoggingDeepCopyOnChange(client ProjectLoggingClient, obj *v3.ProjectLogging, handler func(obj *v3.ProjectLogging) (*v3.ProjectLogging, error)) (*v3.ProjectLogging, error) {
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

func (c *projectLoggingController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *projectLoggingController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *projectLoggingController) OnChange(ctx context.Context, name string, sync ProjectLoggingHandler) {
	c.AddGenericHandler(ctx, name, FromProjectLoggingHandlerToHandler(sync))
}

func (c *projectLoggingController) OnRemove(ctx context.Context, name string, sync ProjectLoggingHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromProjectLoggingHandlerToHandler(sync)))
}

func (c *projectLoggingController) Enqueue(namespace, name string) {
	c.controller.Enqueue(namespace, name)
}

func (c *projectLoggingController) EnqueueAfter(namespace, name string, duration time.Duration) {
	c.controller.EnqueueAfter(namespace, name, duration)
}

func (c *projectLoggingController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *projectLoggingController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *projectLoggingController) Cache() ProjectLoggingCache {
	return &projectLoggingCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *projectLoggingController) Create(obj *v3.ProjectLogging) (*v3.ProjectLogging, error) {
	result := &v3.ProjectLogging{}
	return result, c.client.Create(context.TODO(), obj.Namespace, obj, result, metav1.CreateOptions{})
}

func (c *projectLoggingController) Update(obj *v3.ProjectLogging) (*v3.ProjectLogging, error) {
	result := &v3.ProjectLogging{}
	return result, c.client.Update(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *projectLoggingController) UpdateStatus(obj *v3.ProjectLogging) (*v3.ProjectLogging, error) {
	result := &v3.ProjectLogging{}
	return result, c.client.UpdateStatus(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *projectLoggingController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), namespace, name, *options)
}

func (c *projectLoggingController) Get(namespace, name string, options metav1.GetOptions) (*v3.ProjectLogging, error) {
	result := &v3.ProjectLogging{}
	return result, c.client.Get(context.TODO(), namespace, name, result, options)
}

func (c *projectLoggingController) List(namespace string, opts metav1.ListOptions) (*v3.ProjectLoggingList, error) {
	result := &v3.ProjectLoggingList{}
	return result, c.client.List(context.TODO(), namespace, result, opts)
}

func (c *projectLoggingController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), namespace, opts)
}

func (c *projectLoggingController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (*v3.ProjectLogging, error) {
	result := &v3.ProjectLogging{}
	return result, c.client.Patch(context.TODO(), namespace, name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type projectLoggingCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *projectLoggingCache) Get(namespace, name string) (*v3.ProjectLogging, error) {
	obj, exists, err := c.indexer.GetByKey(namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v3.ProjectLogging), nil
}

func (c *projectLoggingCache) List(namespace string, selector labels.Selector) (ret []*v3.ProjectLogging, err error) {

	err = cache.ListAllByNamespace(c.indexer, namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.ProjectLogging))
	})

	return ret, err
}

func (c *projectLoggingCache) AddIndexer(indexName string, indexer ProjectLoggingIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v3.ProjectLogging))
		},
	}))
}

func (c *projectLoggingCache) GetByIndex(indexName, key string) (result []*v3.ProjectLogging, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v3.ProjectLogging, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v3.ProjectLogging))
	}
	return result, nil
}

type ProjectLoggingStatusHandler func(obj *v3.ProjectLogging, status v3.ProjectLoggingStatus) (v3.ProjectLoggingStatus, error)

type ProjectLoggingGeneratingHandler func(obj *v3.ProjectLogging, status v3.ProjectLoggingStatus) ([]runtime.Object, v3.ProjectLoggingStatus, error)

func RegisterProjectLoggingStatusHandler(ctx context.Context, controller ProjectLoggingController, condition condition.Cond, name string, handler ProjectLoggingStatusHandler) {
	statusHandler := &projectLoggingStatusHandler{
		client:    controller,
		condition: condition,
		handler:   handler,
	}
	controller.AddGenericHandler(ctx, name, FromProjectLoggingHandlerToHandler(statusHandler.sync))
}

func RegisterProjectLoggingGeneratingHandler(ctx context.Context, controller ProjectLoggingController, apply apply.Apply,
	condition condition.Cond, name string, handler ProjectLoggingGeneratingHandler, opts *generic.GeneratingHandlerOptions) {
	statusHandler := &projectLoggingGeneratingHandler{
		ProjectLoggingGeneratingHandler: handler,
		apply:                           apply,
		name:                            name,
		gvk:                             controller.GroupVersionKind(),
	}
	if opts != nil {
		statusHandler.opts = *opts
	}
	controller.OnChange(ctx, name, statusHandler.Remove)
	RegisterProjectLoggingStatusHandler(ctx, controller, condition, name, statusHandler.Handle)
}

type projectLoggingStatusHandler struct {
	client    ProjectLoggingClient
	condition condition.Cond
	handler   ProjectLoggingStatusHandler
}

func (a *projectLoggingStatusHandler) sync(key string, obj *v3.ProjectLogging) (*v3.ProjectLogging, error) {
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

type projectLoggingGeneratingHandler struct {
	ProjectLoggingGeneratingHandler
	apply apply.Apply
	opts  generic.GeneratingHandlerOptions
	gvk   schema.GroupVersionKind
	name  string
}

func (a *projectLoggingGeneratingHandler) Remove(key string, obj *v3.ProjectLogging) (*v3.ProjectLogging, error) {
	if obj != nil {
		return obj, nil
	}

	obj = &v3.ProjectLogging{}
	obj.Namespace, obj.Name = kv.RSplit(key, "/")
	obj.SetGroupVersionKind(a.gvk)

	return nil, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects()
}

func (a *projectLoggingGeneratingHandler) Handle(obj *v3.ProjectLogging, status v3.ProjectLoggingStatus) (v3.ProjectLoggingStatus, error) {
	if !obj.DeletionTimestamp.IsZero() {
		return status, nil
	}

	objs, newStatus, err := a.ProjectLoggingGeneratingHandler(obj, status)
	if err != nil {
		return newStatus, err
	}

	return newStatus, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects(objs...)
}

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

	v3 "github.com/bhojpur/host/pkg/apis/project.bhojpur.net/v3"
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

type AppRevisionHandler func(string, *v3.AppRevision) (*v3.AppRevision, error)

type AppRevisionController interface {
	generic.ControllerMeta
	AppRevisionClient

	OnChange(ctx context.Context, name string, sync AppRevisionHandler)
	OnRemove(ctx context.Context, name string, sync AppRevisionHandler)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, duration time.Duration)

	Cache() AppRevisionCache
}

type AppRevisionClient interface {
	Create(*v3.AppRevision) (*v3.AppRevision, error)
	Update(*v3.AppRevision) (*v3.AppRevision, error)
	UpdateStatus(*v3.AppRevision) (*v3.AppRevision, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v3.AppRevision, error)
	List(namespace string, opts metav1.ListOptions) (*v3.AppRevisionList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v3.AppRevision, err error)
}

type AppRevisionCache interface {
	Get(namespace, name string) (*v3.AppRevision, error)
	List(namespace string, selector labels.Selector) ([]*v3.AppRevision, error)

	AddIndexer(indexName string, indexer AppRevisionIndexer)
	GetByIndex(indexName, key string) ([]*v3.AppRevision, error)
}

type AppRevisionIndexer func(obj *v3.AppRevision) ([]string, error)

type appRevisionController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewAppRevisionController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) AppRevisionController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &appRevisionController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromAppRevisionHandlerToHandler(sync AppRevisionHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v3.AppRevision
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v3.AppRevision))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *appRevisionController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v3.AppRevision))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateAppRevisionDeepCopyOnChange(client AppRevisionClient, obj *v3.AppRevision, handler func(obj *v3.AppRevision) (*v3.AppRevision, error)) (*v3.AppRevision, error) {
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

func (c *appRevisionController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *appRevisionController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *appRevisionController) OnChange(ctx context.Context, name string, sync AppRevisionHandler) {
	c.AddGenericHandler(ctx, name, FromAppRevisionHandlerToHandler(sync))
}

func (c *appRevisionController) OnRemove(ctx context.Context, name string, sync AppRevisionHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromAppRevisionHandlerToHandler(sync)))
}

func (c *appRevisionController) Enqueue(namespace, name string) {
	c.controller.Enqueue(namespace, name)
}

func (c *appRevisionController) EnqueueAfter(namespace, name string, duration time.Duration) {
	c.controller.EnqueueAfter(namespace, name, duration)
}

func (c *appRevisionController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *appRevisionController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *appRevisionController) Cache() AppRevisionCache {
	return &appRevisionCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *appRevisionController) Create(obj *v3.AppRevision) (*v3.AppRevision, error) {
	result := &v3.AppRevision{}
	return result, c.client.Create(context.TODO(), obj.Namespace, obj, result, metav1.CreateOptions{})
}

func (c *appRevisionController) Update(obj *v3.AppRevision) (*v3.AppRevision, error) {
	result := &v3.AppRevision{}
	return result, c.client.Update(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *appRevisionController) UpdateStatus(obj *v3.AppRevision) (*v3.AppRevision, error) {
	result := &v3.AppRevision{}
	return result, c.client.UpdateStatus(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *appRevisionController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), namespace, name, *options)
}

func (c *appRevisionController) Get(namespace, name string, options metav1.GetOptions) (*v3.AppRevision, error) {
	result := &v3.AppRevision{}
	return result, c.client.Get(context.TODO(), namespace, name, result, options)
}

func (c *appRevisionController) List(namespace string, opts metav1.ListOptions) (*v3.AppRevisionList, error) {
	result := &v3.AppRevisionList{}
	return result, c.client.List(context.TODO(), namespace, result, opts)
}

func (c *appRevisionController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), namespace, opts)
}

func (c *appRevisionController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (*v3.AppRevision, error) {
	result := &v3.AppRevision{}
	return result, c.client.Patch(context.TODO(), namespace, name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type appRevisionCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *appRevisionCache) Get(namespace, name string) (*v3.AppRevision, error) {
	obj, exists, err := c.indexer.GetByKey(namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v3.AppRevision), nil
}

func (c *appRevisionCache) List(namespace string, selector labels.Selector) (ret []*v3.AppRevision, err error) {

	err = cache.ListAllByNamespace(c.indexer, namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.AppRevision))
	})

	return ret, err
}

func (c *appRevisionCache) AddIndexer(indexName string, indexer AppRevisionIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v3.AppRevision))
		},
	}))
}

func (c *appRevisionCache) GetByIndex(indexName, key string) (result []*v3.AppRevision, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v3.AppRevision, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v3.AppRevision))
	}
	return result, nil
}

type AppRevisionStatusHandler func(obj *v3.AppRevision, status v3.AppRevisionStatus) (v3.AppRevisionStatus, error)

type AppRevisionGeneratingHandler func(obj *v3.AppRevision, status v3.AppRevisionStatus) ([]runtime.Object, v3.AppRevisionStatus, error)

func RegisterAppRevisionStatusHandler(ctx context.Context, controller AppRevisionController, condition condition.Cond, name string, handler AppRevisionStatusHandler) {
	statusHandler := &appRevisionStatusHandler{
		client:    controller,
		condition: condition,
		handler:   handler,
	}
	controller.AddGenericHandler(ctx, name, FromAppRevisionHandlerToHandler(statusHandler.sync))
}

func RegisterAppRevisionGeneratingHandler(ctx context.Context, controller AppRevisionController, apply apply.Apply,
	condition condition.Cond, name string, handler AppRevisionGeneratingHandler, opts *generic.GeneratingHandlerOptions) {
	statusHandler := &appRevisionGeneratingHandler{
		AppRevisionGeneratingHandler: handler,
		apply:                        apply,
		name:                         name,
		gvk:                          controller.GroupVersionKind(),
	}
	if opts != nil {
		statusHandler.opts = *opts
	}
	controller.OnChange(ctx, name, statusHandler.Remove)
	RegisterAppRevisionStatusHandler(ctx, controller, condition, name, statusHandler.Handle)
}

type appRevisionStatusHandler struct {
	client    AppRevisionClient
	condition condition.Cond
	handler   AppRevisionStatusHandler
}

func (a *appRevisionStatusHandler) sync(key string, obj *v3.AppRevision) (*v3.AppRevision, error) {
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

type appRevisionGeneratingHandler struct {
	AppRevisionGeneratingHandler
	apply apply.Apply
	opts  generic.GeneratingHandlerOptions
	gvk   schema.GroupVersionKind
	name  string
}

func (a *appRevisionGeneratingHandler) Remove(key string, obj *v3.AppRevision) (*v3.AppRevision, error) {
	if obj != nil {
		return obj, nil
	}

	obj = &v3.AppRevision{}
	obj.Namespace, obj.Name = kv.RSplit(key, "/")
	obj.SetGroupVersionKind(a.gvk)

	return nil, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects()
}

func (a *appRevisionGeneratingHandler) Handle(obj *v3.AppRevision, status v3.AppRevisionStatus) (v3.AppRevisionStatus, error) {
	if !obj.DeletionTimestamp.IsZero() {
		return status, nil
	}

	objs, newStatus, err := a.AppRevisionGeneratingHandler(obj, status)
	if err != nil {
		return newStatus, err
	}

	return newStatus, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects(objs...)
}

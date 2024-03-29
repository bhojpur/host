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

type TemplateVersionHandler func(string, *v3.TemplateVersion) (*v3.TemplateVersion, error)

type TemplateVersionController interface {
	generic.ControllerMeta
	TemplateVersionClient

	OnChange(ctx context.Context, name string, sync TemplateVersionHandler)
	OnRemove(ctx context.Context, name string, sync TemplateVersionHandler)
	Enqueue(name string)
	EnqueueAfter(name string, duration time.Duration)

	Cache() TemplateVersionCache
}

type TemplateVersionClient interface {
	Create(*v3.TemplateVersion) (*v3.TemplateVersion, error)
	Update(*v3.TemplateVersion) (*v3.TemplateVersion, error)
	UpdateStatus(*v3.TemplateVersion) (*v3.TemplateVersion, error)
	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*v3.TemplateVersion, error)
	List(opts metav1.ListOptions) (*v3.TemplateVersionList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v3.TemplateVersion, err error)
}

type TemplateVersionCache interface {
	Get(name string) (*v3.TemplateVersion, error)
	List(selector labels.Selector) ([]*v3.TemplateVersion, error)

	AddIndexer(indexName string, indexer TemplateVersionIndexer)
	GetByIndex(indexName, key string) ([]*v3.TemplateVersion, error)
}

type TemplateVersionIndexer func(obj *v3.TemplateVersion) ([]string, error)

type templateVersionController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewTemplateVersionController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) TemplateVersionController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &templateVersionController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromTemplateVersionHandlerToHandler(sync TemplateVersionHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v3.TemplateVersion
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v3.TemplateVersion))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *templateVersionController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v3.TemplateVersion))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateTemplateVersionDeepCopyOnChange(client TemplateVersionClient, obj *v3.TemplateVersion, handler func(obj *v3.TemplateVersion) (*v3.TemplateVersion, error)) (*v3.TemplateVersion, error) {
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

func (c *templateVersionController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *templateVersionController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *templateVersionController) OnChange(ctx context.Context, name string, sync TemplateVersionHandler) {
	c.AddGenericHandler(ctx, name, FromTemplateVersionHandlerToHandler(sync))
}

func (c *templateVersionController) OnRemove(ctx context.Context, name string, sync TemplateVersionHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromTemplateVersionHandlerToHandler(sync)))
}

func (c *templateVersionController) Enqueue(name string) {
	c.controller.Enqueue("", name)
}

func (c *templateVersionController) EnqueueAfter(name string, duration time.Duration) {
	c.controller.EnqueueAfter("", name, duration)
}

func (c *templateVersionController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *templateVersionController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *templateVersionController) Cache() TemplateVersionCache {
	return &templateVersionCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *templateVersionController) Create(obj *v3.TemplateVersion) (*v3.TemplateVersion, error) {
	result := &v3.TemplateVersion{}
	return result, c.client.Create(context.TODO(), "", obj, result, metav1.CreateOptions{})
}

func (c *templateVersionController) Update(obj *v3.TemplateVersion) (*v3.TemplateVersion, error) {
	result := &v3.TemplateVersion{}
	return result, c.client.Update(context.TODO(), "", obj, result, metav1.UpdateOptions{})
}

func (c *templateVersionController) UpdateStatus(obj *v3.TemplateVersion) (*v3.TemplateVersion, error) {
	result := &v3.TemplateVersion{}
	return result, c.client.UpdateStatus(context.TODO(), "", obj, result, metav1.UpdateOptions{})
}

func (c *templateVersionController) Delete(name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), "", name, *options)
}

func (c *templateVersionController) Get(name string, options metav1.GetOptions) (*v3.TemplateVersion, error) {
	result := &v3.TemplateVersion{}
	return result, c.client.Get(context.TODO(), "", name, result, options)
}

func (c *templateVersionController) List(opts metav1.ListOptions) (*v3.TemplateVersionList, error) {
	result := &v3.TemplateVersionList{}
	return result, c.client.List(context.TODO(), "", result, opts)
}

func (c *templateVersionController) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), "", opts)
}

func (c *templateVersionController) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (*v3.TemplateVersion, error) {
	result := &v3.TemplateVersion{}
	return result, c.client.Patch(context.TODO(), "", name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type templateVersionCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *templateVersionCache) Get(name string) (*v3.TemplateVersion, error) {
	obj, exists, err := c.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v3.TemplateVersion), nil
}

func (c *templateVersionCache) List(selector labels.Selector) (ret []*v3.TemplateVersion, err error) {

	err = cache.ListAll(c.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.TemplateVersion))
	})

	return ret, err
}

func (c *templateVersionCache) AddIndexer(indexName string, indexer TemplateVersionIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v3.TemplateVersion))
		},
	}))
}

func (c *templateVersionCache) GetByIndex(indexName, key string) (result []*v3.TemplateVersion, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v3.TemplateVersion, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v3.TemplateVersion))
	}
	return result, nil
}

type TemplateVersionStatusHandler func(obj *v3.TemplateVersion, status v3.TemplateVersionStatus) (v3.TemplateVersionStatus, error)

type TemplateVersionGeneratingHandler func(obj *v3.TemplateVersion, status v3.TemplateVersionStatus) ([]runtime.Object, v3.TemplateVersionStatus, error)

func RegisterTemplateVersionStatusHandler(ctx context.Context, controller TemplateVersionController, condition condition.Cond, name string, handler TemplateVersionStatusHandler) {
	statusHandler := &templateVersionStatusHandler{
		client:    controller,
		condition: condition,
		handler:   handler,
	}
	controller.AddGenericHandler(ctx, name, FromTemplateVersionHandlerToHandler(statusHandler.sync))
}

func RegisterTemplateVersionGeneratingHandler(ctx context.Context, controller TemplateVersionController, apply apply.Apply,
	condition condition.Cond, name string, handler TemplateVersionGeneratingHandler, opts *generic.GeneratingHandlerOptions) {
	statusHandler := &templateVersionGeneratingHandler{
		TemplateVersionGeneratingHandler: handler,
		apply:                            apply,
		name:                             name,
		gvk:                              controller.GroupVersionKind(),
	}
	if opts != nil {
		statusHandler.opts = *opts
	}
	controller.OnChange(ctx, name, statusHandler.Remove)
	RegisterTemplateVersionStatusHandler(ctx, controller, condition, name, statusHandler.Handle)
}

type templateVersionStatusHandler struct {
	client    TemplateVersionClient
	condition condition.Cond
	handler   TemplateVersionStatusHandler
}

func (a *templateVersionStatusHandler) sync(key string, obj *v3.TemplateVersion) (*v3.TemplateVersion, error) {
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

type templateVersionGeneratingHandler struct {
	TemplateVersionGeneratingHandler
	apply apply.Apply
	opts  generic.GeneratingHandlerOptions
	gvk   schema.GroupVersionKind
	name  string
}

func (a *templateVersionGeneratingHandler) Remove(key string, obj *v3.TemplateVersion) (*v3.TemplateVersion, error) {
	if obj != nil {
		return obj, nil
	}

	obj = &v3.TemplateVersion{}
	obj.Namespace, obj.Name = kv.RSplit(key, "/")
	obj.SetGroupVersionKind(a.gvk)

	return nil, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects()
}

func (a *templateVersionGeneratingHandler) Handle(obj *v3.TemplateVersion, status v3.TemplateVersionStatus) (v3.TemplateVersionStatus, error) {
	if !obj.DeletionTimestamp.IsZero() {
		return status, nil
	}

	objs, newStatus, err := a.TemplateVersionGeneratingHandler(obj, status)
	if err != nil {
		return newStatus, err
	}

	return newStatus, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects(objs...)
}

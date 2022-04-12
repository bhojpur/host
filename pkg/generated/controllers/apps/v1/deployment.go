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

type DeploymentHandler func(string, *v1.Deployment) (*v1.Deployment, error)

type DeploymentController interface {
	generic.ControllerMeta
	DeploymentClient

	OnChange(ctx context.Context, name string, sync DeploymentHandler)
	OnRemove(ctx context.Context, name string, sync DeploymentHandler)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, duration time.Duration)

	Cache() DeploymentCache
}

type DeploymentClient interface {
	Create(*v1.Deployment) (*v1.Deployment, error)
	Update(*v1.Deployment) (*v1.Deployment, error)
	UpdateStatus(*v1.Deployment) (*v1.Deployment, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v1.Deployment, error)
	List(namespace string, opts metav1.ListOptions) (*v1.DeploymentList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.Deployment, err error)
}

type DeploymentCache interface {
	Get(namespace, name string) (*v1.Deployment, error)
	List(namespace string, selector labels.Selector) ([]*v1.Deployment, error)

	AddIndexer(indexName string, indexer DeploymentIndexer)
	GetByIndex(indexName, key string) ([]*v1.Deployment, error)
}

type DeploymentIndexer func(obj *v1.Deployment) ([]string, error)

type deploymentController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewDeploymentController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) DeploymentController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &deploymentController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromDeploymentHandlerToHandler(sync DeploymentHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v1.Deployment
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v1.Deployment))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *deploymentController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v1.Deployment))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateDeploymentDeepCopyOnChange(client DeploymentClient, obj *v1.Deployment, handler func(obj *v1.Deployment) (*v1.Deployment, error)) (*v1.Deployment, error) {
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

func (c *deploymentController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *deploymentController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *deploymentController) OnChange(ctx context.Context, name string, sync DeploymentHandler) {
	c.AddGenericHandler(ctx, name, FromDeploymentHandlerToHandler(sync))
}

func (c *deploymentController) OnRemove(ctx context.Context, name string, sync DeploymentHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromDeploymentHandlerToHandler(sync)))
}

func (c *deploymentController) Enqueue(namespace, name string) {
	c.controller.Enqueue(namespace, name)
}

func (c *deploymentController) EnqueueAfter(namespace, name string, duration time.Duration) {
	c.controller.EnqueueAfter(namespace, name, duration)
}

func (c *deploymentController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *deploymentController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *deploymentController) Cache() DeploymentCache {
	return &deploymentCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *deploymentController) Create(obj *v1.Deployment) (*v1.Deployment, error) {
	result := &v1.Deployment{}
	return result, c.client.Create(context.TODO(), obj.Namespace, obj, result, metav1.CreateOptions{})
}

func (c *deploymentController) Update(obj *v1.Deployment) (*v1.Deployment, error) {
	result := &v1.Deployment{}
	return result, c.client.Update(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *deploymentController) UpdateStatus(obj *v1.Deployment) (*v1.Deployment, error) {
	result := &v1.Deployment{}
	return result, c.client.UpdateStatus(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *deploymentController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), namespace, name, *options)
}

func (c *deploymentController) Get(namespace, name string, options metav1.GetOptions) (*v1.Deployment, error) {
	result := &v1.Deployment{}
	return result, c.client.Get(context.TODO(), namespace, name, result, options)
}

func (c *deploymentController) List(namespace string, opts metav1.ListOptions) (*v1.DeploymentList, error) {
	result := &v1.DeploymentList{}
	return result, c.client.List(context.TODO(), namespace, result, opts)
}

func (c *deploymentController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), namespace, opts)
}

func (c *deploymentController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (*v1.Deployment, error) {
	result := &v1.Deployment{}
	return result, c.client.Patch(context.TODO(), namespace, name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type deploymentCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *deploymentCache) Get(namespace, name string) (*v1.Deployment, error) {
	obj, exists, err := c.indexer.GetByKey(namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v1.Deployment), nil
}

func (c *deploymentCache) List(namespace string, selector labels.Selector) (ret []*v1.Deployment, err error) {

	err = cache.ListAllByNamespace(c.indexer, namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Deployment))
	})

	return ret, err
}

func (c *deploymentCache) AddIndexer(indexName string, indexer DeploymentIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v1.Deployment))
		},
	}))
}

func (c *deploymentCache) GetByIndex(indexName, key string) (result []*v1.Deployment, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v1.Deployment, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v1.Deployment))
	}
	return result, nil
}

type DeploymentStatusHandler func(obj *v1.Deployment, status v1.DeploymentStatus) (v1.DeploymentStatus, error)

type DeploymentGeneratingHandler func(obj *v1.Deployment, status v1.DeploymentStatus) ([]runtime.Object, v1.DeploymentStatus, error)

func RegisterDeploymentStatusHandler(ctx context.Context, controller DeploymentController, condition condition.Cond, name string, handler DeploymentStatusHandler) {
	statusHandler := &deploymentStatusHandler{
		client:    controller,
		condition: condition,
		handler:   handler,
	}
	controller.AddGenericHandler(ctx, name, FromDeploymentHandlerToHandler(statusHandler.sync))
}

func RegisterDeploymentGeneratingHandler(ctx context.Context, controller DeploymentController, apply apply.Apply,
	condition condition.Cond, name string, handler DeploymentGeneratingHandler, opts *generic.GeneratingHandlerOptions) {
	statusHandler := &deploymentGeneratingHandler{
		DeploymentGeneratingHandler: handler,
		apply:                       apply,
		name:                        name,
		gvk:                         controller.GroupVersionKind(),
	}
	if opts != nil {
		statusHandler.opts = *opts
	}
	controller.OnChange(ctx, name, statusHandler.Remove)
	RegisterDeploymentStatusHandler(ctx, controller, condition, name, statusHandler.Handle)
}

type deploymentStatusHandler struct {
	client    DeploymentClient
	condition condition.Cond
	handler   DeploymentStatusHandler
}

func (a *deploymentStatusHandler) sync(key string, obj *v1.Deployment) (*v1.Deployment, error) {
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

type deploymentGeneratingHandler struct {
	DeploymentGeneratingHandler
	apply apply.Apply
	opts  generic.GeneratingHandlerOptions
	gvk   schema.GroupVersionKind
	name  string
}

func (a *deploymentGeneratingHandler) Remove(key string, obj *v1.Deployment) (*v1.Deployment, error) {
	if obj != nil {
		return obj, nil
	}

	obj = &v1.Deployment{}
	obj.Namespace, obj.Name = kv.RSplit(key, "/")
	obj.SetGroupVersionKind(a.gvk)

	return nil, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects()
}

func (a *deploymentGeneratingHandler) Handle(obj *v1.Deployment, status v1.DeploymentStatus) (v1.DeploymentStatus, error) {
	if !obj.DeletionTimestamp.IsZero() {
		return status, nil
	}

	objs, newStatus, err := a.DeploymentGeneratingHandler(obj, status)
	if err != nil {
		return newStatus, err
	}

	return newStatus, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects(objs...)
}
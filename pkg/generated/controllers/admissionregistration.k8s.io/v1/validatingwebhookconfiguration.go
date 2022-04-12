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
	v1 "k8s.io/api/admissionregistration/v1"
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

type ValidatingWebhookConfigurationHandler func(string, *v1.ValidatingWebhookConfiguration) (*v1.ValidatingWebhookConfiguration, error)

type ValidatingWebhookConfigurationController interface {
	generic.ControllerMeta
	ValidatingWebhookConfigurationClient

	OnChange(ctx context.Context, name string, sync ValidatingWebhookConfigurationHandler)
	OnRemove(ctx context.Context, name string, sync ValidatingWebhookConfigurationHandler)
	Enqueue(name string)
	EnqueueAfter(name string, duration time.Duration)

	Cache() ValidatingWebhookConfigurationCache
}

type ValidatingWebhookConfigurationClient interface {
	Create(*v1.ValidatingWebhookConfiguration) (*v1.ValidatingWebhookConfiguration, error)
	Update(*v1.ValidatingWebhookConfiguration) (*v1.ValidatingWebhookConfiguration, error)

	Delete(name string, options *metav1.DeleteOptions) error
	Get(name string, options metav1.GetOptions) (*v1.ValidatingWebhookConfiguration, error)
	List(opts metav1.ListOptions) (*v1.ValidatingWebhookConfigurationList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ValidatingWebhookConfiguration, err error)
}

type ValidatingWebhookConfigurationCache interface {
	Get(name string) (*v1.ValidatingWebhookConfiguration, error)
	List(selector labels.Selector) ([]*v1.ValidatingWebhookConfiguration, error)

	AddIndexer(indexName string, indexer ValidatingWebhookConfigurationIndexer)
	GetByIndex(indexName, key string) ([]*v1.ValidatingWebhookConfiguration, error)
}

type ValidatingWebhookConfigurationIndexer func(obj *v1.ValidatingWebhookConfiguration) ([]string, error)

type validatingWebhookConfigurationController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewValidatingWebhookConfigurationController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) ValidatingWebhookConfigurationController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &validatingWebhookConfigurationController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromValidatingWebhookConfigurationHandlerToHandler(sync ValidatingWebhookConfigurationHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v1.ValidatingWebhookConfiguration
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v1.ValidatingWebhookConfiguration))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *validatingWebhookConfigurationController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v1.ValidatingWebhookConfiguration))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateValidatingWebhookConfigurationDeepCopyOnChange(client ValidatingWebhookConfigurationClient, obj *v1.ValidatingWebhookConfiguration, handler func(obj *v1.ValidatingWebhookConfiguration) (*v1.ValidatingWebhookConfiguration, error)) (*v1.ValidatingWebhookConfiguration, error) {
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

func (c *validatingWebhookConfigurationController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *validatingWebhookConfigurationController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *validatingWebhookConfigurationController) OnChange(ctx context.Context, name string, sync ValidatingWebhookConfigurationHandler) {
	c.AddGenericHandler(ctx, name, FromValidatingWebhookConfigurationHandlerToHandler(sync))
}

func (c *validatingWebhookConfigurationController) OnRemove(ctx context.Context, name string, sync ValidatingWebhookConfigurationHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromValidatingWebhookConfigurationHandlerToHandler(sync)))
}

func (c *validatingWebhookConfigurationController) Enqueue(name string) {
	c.controller.Enqueue("", name)
}

func (c *validatingWebhookConfigurationController) EnqueueAfter(name string, duration time.Duration) {
	c.controller.EnqueueAfter("", name, duration)
}

func (c *validatingWebhookConfigurationController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *validatingWebhookConfigurationController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *validatingWebhookConfigurationController) Cache() ValidatingWebhookConfigurationCache {
	return &validatingWebhookConfigurationCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *validatingWebhookConfigurationController) Create(obj *v1.ValidatingWebhookConfiguration) (*v1.ValidatingWebhookConfiguration, error) {
	result := &v1.ValidatingWebhookConfiguration{}
	return result, c.client.Create(context.TODO(), "", obj, result, metav1.CreateOptions{})
}

func (c *validatingWebhookConfigurationController) Update(obj *v1.ValidatingWebhookConfiguration) (*v1.ValidatingWebhookConfiguration, error) {
	result := &v1.ValidatingWebhookConfiguration{}
	return result, c.client.Update(context.TODO(), "", obj, result, metav1.UpdateOptions{})
}

func (c *validatingWebhookConfigurationController) Delete(name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), "", name, *options)
}

func (c *validatingWebhookConfigurationController) Get(name string, options metav1.GetOptions) (*v1.ValidatingWebhookConfiguration, error) {
	result := &v1.ValidatingWebhookConfiguration{}
	return result, c.client.Get(context.TODO(), "", name, result, options)
}

func (c *validatingWebhookConfigurationController) List(opts metav1.ListOptions) (*v1.ValidatingWebhookConfigurationList, error) {
	result := &v1.ValidatingWebhookConfigurationList{}
	return result, c.client.List(context.TODO(), "", result, opts)
}

func (c *validatingWebhookConfigurationController) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), "", opts)
}

func (c *validatingWebhookConfigurationController) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (*v1.ValidatingWebhookConfiguration, error) {
	result := &v1.ValidatingWebhookConfiguration{}
	return result, c.client.Patch(context.TODO(), "", name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type validatingWebhookConfigurationCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *validatingWebhookConfigurationCache) Get(name string) (*v1.ValidatingWebhookConfiguration, error) {
	obj, exists, err := c.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v1.ValidatingWebhookConfiguration), nil
}

func (c *validatingWebhookConfigurationCache) List(selector labels.Selector) (ret []*v1.ValidatingWebhookConfiguration, err error) {

	err = cache.ListAll(c.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ValidatingWebhookConfiguration))
	})

	return ret, err
}

func (c *validatingWebhookConfigurationCache) AddIndexer(indexName string, indexer ValidatingWebhookConfigurationIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v1.ValidatingWebhookConfiguration))
		},
	}))
}

func (c *validatingWebhookConfigurationCache) GetByIndex(indexName, key string) (result []*v1.ValidatingWebhookConfiguration, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v1.ValidatingWebhookConfiguration, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v1.ValidatingWebhookConfiguration))
	}
	return result, nil
}
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

type ClusterCatalogHandler func(string, *v3.ClusterCatalog) (*v3.ClusterCatalog, error)

type ClusterCatalogController interface {
	generic.ControllerMeta
	ClusterCatalogClient

	OnChange(ctx context.Context, name string, sync ClusterCatalogHandler)
	OnRemove(ctx context.Context, name string, sync ClusterCatalogHandler)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, duration time.Duration)

	Cache() ClusterCatalogCache
}

type ClusterCatalogClient interface {
	Create(*v3.ClusterCatalog) (*v3.ClusterCatalog, error)
	Update(*v3.ClusterCatalog) (*v3.ClusterCatalog, error)

	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v3.ClusterCatalog, error)
	List(namespace string, opts metav1.ListOptions) (*v3.ClusterCatalogList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v3.ClusterCatalog, err error)
}

type ClusterCatalogCache interface {
	Get(namespace, name string) (*v3.ClusterCatalog, error)
	List(namespace string, selector labels.Selector) ([]*v3.ClusterCatalog, error)

	AddIndexer(indexName string, indexer ClusterCatalogIndexer)
	GetByIndex(indexName, key string) ([]*v3.ClusterCatalog, error)
}

type ClusterCatalogIndexer func(obj *v3.ClusterCatalog) ([]string, error)

type clusterCatalogController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewClusterCatalogController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) ClusterCatalogController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &clusterCatalogController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromClusterCatalogHandlerToHandler(sync ClusterCatalogHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v3.ClusterCatalog
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v3.ClusterCatalog))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *clusterCatalogController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v3.ClusterCatalog))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateClusterCatalogDeepCopyOnChange(client ClusterCatalogClient, obj *v3.ClusterCatalog, handler func(obj *v3.ClusterCatalog) (*v3.ClusterCatalog, error)) (*v3.ClusterCatalog, error) {
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

func (c *clusterCatalogController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *clusterCatalogController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *clusterCatalogController) OnChange(ctx context.Context, name string, sync ClusterCatalogHandler) {
	c.AddGenericHandler(ctx, name, FromClusterCatalogHandlerToHandler(sync))
}

func (c *clusterCatalogController) OnRemove(ctx context.Context, name string, sync ClusterCatalogHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromClusterCatalogHandlerToHandler(sync)))
}

func (c *clusterCatalogController) Enqueue(namespace, name string) {
	c.controller.Enqueue(namespace, name)
}

func (c *clusterCatalogController) EnqueueAfter(namespace, name string, duration time.Duration) {
	c.controller.EnqueueAfter(namespace, name, duration)
}

func (c *clusterCatalogController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *clusterCatalogController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *clusterCatalogController) Cache() ClusterCatalogCache {
	return &clusterCatalogCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *clusterCatalogController) Create(obj *v3.ClusterCatalog) (*v3.ClusterCatalog, error) {
	result := &v3.ClusterCatalog{}
	return result, c.client.Create(context.TODO(), obj.Namespace, obj, result, metav1.CreateOptions{})
}

func (c *clusterCatalogController) Update(obj *v3.ClusterCatalog) (*v3.ClusterCatalog, error) {
	result := &v3.ClusterCatalog{}
	return result, c.client.Update(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *clusterCatalogController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), namespace, name, *options)
}

func (c *clusterCatalogController) Get(namespace, name string, options metav1.GetOptions) (*v3.ClusterCatalog, error) {
	result := &v3.ClusterCatalog{}
	return result, c.client.Get(context.TODO(), namespace, name, result, options)
}

func (c *clusterCatalogController) List(namespace string, opts metav1.ListOptions) (*v3.ClusterCatalogList, error) {
	result := &v3.ClusterCatalogList{}
	return result, c.client.List(context.TODO(), namespace, result, opts)
}

func (c *clusterCatalogController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), namespace, opts)
}

func (c *clusterCatalogController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (*v3.ClusterCatalog, error) {
	result := &v3.ClusterCatalog{}
	return result, c.client.Patch(context.TODO(), namespace, name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type clusterCatalogCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *clusterCatalogCache) Get(namespace, name string) (*v3.ClusterCatalog, error) {
	obj, exists, err := c.indexer.GetByKey(namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v3.ClusterCatalog), nil
}

func (c *clusterCatalogCache) List(namespace string, selector labels.Selector) (ret []*v3.ClusterCatalog, err error) {

	err = cache.ListAllByNamespace(c.indexer, namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.ClusterCatalog))
	})

	return ret, err
}

func (c *clusterCatalogCache) AddIndexer(indexName string, indexer ClusterCatalogIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v3.ClusterCatalog))
		},
	}))
}

func (c *clusterCatalogCache) GetByIndex(indexName, key string) (result []*v3.ClusterCatalog, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v3.ClusterCatalog, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v3.ClusterCatalog))
	}
	return result, nil
}

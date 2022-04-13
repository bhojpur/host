package v3

import (
	"context"
	"time"

	"github.com/bhojpur/host/pkg/apis/management.bhojpur.net/v3"
	"github.com/bhojpur/host/pkg/core/controller"
	"github.com/bhojpur/host/pkg/core/objectclient"
	"github.com/bhojpur/host/pkg/core/resource"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

var (
	BkeAddonGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "BkeAddon",
	}
	BkeAddonResource = metav1.APIResource{
		Name:         "bkeaddons",
		SingularName: "bkeaddon",
		Namespaced:   true,

		Kind: BkeAddonGroupVersionKind.Kind,
	}

	BkeAddonGroupVersionResource = schema.GroupVersionResource{
		Group:    GroupName,
		Version:  Version,
		Resource: "bkeaddons",
	}
)

func init() {
	resource.Put(BkeAddonGroupVersionResource)
}

// Deprecated use v3.BkeAddon instead
type BkeAddon = v3.BkeAddon

func NewBkeAddon(namespace, name string, obj v3.BkeAddon) *v3.BkeAddon {
	obj.APIVersion, obj.Kind = BkeAddonGroupVersionKind.ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

type BkeAddonHandlerFunc func(key string, obj *v3.BkeAddon) (runtime.Object, error)

type BkeAddonChangeHandlerFunc func(obj *v3.BkeAddon) (runtime.Object, error)

type BkeAddonLister interface {
	List(namespace string, selector labels.Selector) (ret []*v3.BkeAddon, err error)
	Get(namespace, name string) (*v3.BkeAddon, error)
}

type BkeAddonController interface {
	Generic() controller.GenericController
	Informer() cache.SharedIndexInformer
	Lister() BkeAddonLister
	AddHandler(ctx context.Context, name string, handler BkeAddonHandlerFunc)
	AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync BkeAddonHandlerFunc)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, handler BkeAddonHandlerFunc)
	AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, handler BkeAddonHandlerFunc)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, after time.Duration)
}

type BkeAddonInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*v3.BkeAddon) (*v3.BkeAddon, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*v3.BkeAddon, error)
	Get(name string, opts metav1.GetOptions) (*v3.BkeAddon, error)
	Update(*v3.BkeAddon) (*v3.BkeAddon, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*v3.BkeAddonList, error)
	ListNamespaced(namespace string, opts metav1.ListOptions) (*v3.BkeAddonList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() BkeAddonController
	AddHandler(ctx context.Context, name string, sync BkeAddonHandlerFunc)
	AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync BkeAddonHandlerFunc)
	AddLifecycle(ctx context.Context, name string, lifecycle BkeAddonLifecycle)
	AddFeatureLifecycle(ctx context.Context, enabled func() bool, name string, lifecycle BkeAddonLifecycle)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync BkeAddonHandlerFunc)
	AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, sync BkeAddonHandlerFunc)
	AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle BkeAddonLifecycle)
	AddClusterScopedFeatureLifecycle(ctx context.Context, enabled func() bool, name, clusterName string, lifecycle BkeAddonLifecycle)
}

type bkeAddonLister struct {
	ns         string
	controller *bkeAddonController
}

func (l *bkeAddonLister) List(namespace string, selector labels.Selector) (ret []*v3.BkeAddon, err error) {
	if namespace == "" {
		namespace = l.ns
	}
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*v3.BkeAddon))
	})
	return
}

func (l *bkeAddonLister) Get(namespace, name string) (*v3.BkeAddon, error) {
	var key string
	if namespace != "" {
		key = namespace + "/" + name
	} else {
		key = name
	}
	obj, exists, err := l.controller.Informer().GetIndexer().GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(schema.GroupResource{
			Group:    BkeAddonGroupVersionKind.Group,
			Resource: BkeAddonGroupVersionResource.Resource,
		}, key)
	}
	return obj.(*v3.BkeAddon), nil
}

type bkeAddonController struct {
	ns string
	controller.GenericController
}

func (c *bkeAddonController) Generic() controller.GenericController {
	return c.GenericController
}

func (c *bkeAddonController) Lister() BkeAddonLister {
	return &bkeAddonLister{
		ns:         c.ns,
		controller: c,
	}
}

func (c *bkeAddonController) AddHandler(ctx context.Context, name string, handler BkeAddonHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v3.BkeAddon); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *bkeAddonController) AddFeatureHandler(ctx context.Context, enabled func() bool, name string, handler BkeAddonHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled() {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v3.BkeAddon); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *bkeAddonController) AddClusterScopedHandler(ctx context.Context, name, cluster string, handler BkeAddonHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v3.BkeAddon); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *bkeAddonController) AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, cluster string, handler BkeAddonHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled() {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v3.BkeAddon); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

type bkeAddonFactory struct {
}

func (c bkeAddonFactory) Object() runtime.Object {
	return &v3.BkeAddon{}
}

func (c bkeAddonFactory) List() runtime.Object {
	return &v3.BkeAddonList{}
}

func (s *bkeAddonClient) Controller() BkeAddonController {
	genericController := controller.NewGenericController(s.ns, BkeAddonGroupVersionKind.Kind+"Controller",
		s.client.controllerFactory.ForResourceKind(BkeAddonGroupVersionResource, BkeAddonGroupVersionKind.Kind, true))

	return &bkeAddonController{
		ns:                s.ns,
		GenericController: genericController,
	}
}

type bkeAddonClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   BkeAddonController
}

func (s *bkeAddonClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *bkeAddonClient) Create(o *v3.BkeAddon) (*v3.BkeAddon, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*v3.BkeAddon), err
}

func (s *bkeAddonClient) Get(name string, opts metav1.GetOptions) (*v3.BkeAddon, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*v3.BkeAddon), err
}

func (s *bkeAddonClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*v3.BkeAddon, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*v3.BkeAddon), err
}

func (s *bkeAddonClient) Update(o *v3.BkeAddon) (*v3.BkeAddon, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*v3.BkeAddon), err
}

func (s *bkeAddonClient) UpdateStatus(o *v3.BkeAddon) (*v3.BkeAddon, error) {
	obj, err := s.objectClient.UpdateStatus(o.Name, o)
	return obj.(*v3.BkeAddon), err
}

func (s *bkeAddonClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *bkeAddonClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *bkeAddonClient) List(opts metav1.ListOptions) (*v3.BkeAddonList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*v3.BkeAddonList), err
}

func (s *bkeAddonClient) ListNamespaced(namespace string, opts metav1.ListOptions) (*v3.BkeAddonList, error) {
	obj, err := s.objectClient.ListNamespaced(namespace, opts)
	return obj.(*v3.BkeAddonList), err
}

func (s *bkeAddonClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *bkeAddonClient) Patch(o *v3.BkeAddon, patchType types.PatchType, data []byte, subresources ...string) (*v3.BkeAddon, error) {
	obj, err := s.objectClient.Patch(o.Name, o, patchType, data, subresources...)
	return obj.(*v3.BkeAddon), err
}

func (s *bkeAddonClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *bkeAddonClient) AddHandler(ctx context.Context, name string, sync BkeAddonHandlerFunc) {
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *bkeAddonClient) AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync BkeAddonHandlerFunc) {
	s.Controller().AddFeatureHandler(ctx, enabled, name, sync)
}

func (s *bkeAddonClient) AddLifecycle(ctx context.Context, name string, lifecycle BkeAddonLifecycle) {
	sync := NewBkeAddonLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *bkeAddonClient) AddFeatureLifecycle(ctx context.Context, enabled func() bool, name string, lifecycle BkeAddonLifecycle) {
	sync := NewBkeAddonLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddFeatureHandler(ctx, enabled, name, sync)
}

func (s *bkeAddonClient) AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync BkeAddonHandlerFunc) {
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *bkeAddonClient) AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, sync BkeAddonHandlerFunc) {
	s.Controller().AddClusterScopedFeatureHandler(ctx, enabled, name, clusterName, sync)
}

func (s *bkeAddonClient) AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle BkeAddonLifecycle) {
	sync := NewBkeAddonLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *bkeAddonClient) AddClusterScopedFeatureLifecycle(ctx context.Context, enabled func() bool, name, clusterName string, lifecycle BkeAddonLifecycle) {
	sync := NewBkeAddonLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedFeatureHandler(ctx, enabled, name, clusterName, sync)
}

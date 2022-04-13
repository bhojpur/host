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
	ContainerDriverGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "ContainerDriver",
	}
	ContainerDriverResource = metav1.APIResource{
		Name:         "containerdrivers",
		SingularName: "containerdriver",
		Namespaced:   false,
		Kind:         ContainerDriverGroupVersionKind.Kind,
	}

	ContainerDriverGroupVersionResource = schema.GroupVersionResource{
		Group:    GroupName,
		Version:  Version,
		Resource: "containerdrivers",
	}
)

func init() {
	resource.Put(ContainerDriverGroupVersionResource)
}

// Deprecated use v3.ContainerDriver instead
type ContainerDriver = v3.ContainerDriver

func NewContainerDriver(namespace, name string, obj v3.ContainerDriver) *v3.ContainerDriver {
	obj.APIVersion, obj.Kind = ContainerDriverGroupVersionKind.ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

type ContainerDriverHandlerFunc func(key string, obj *v3.ContainerDriver) (runtime.Object, error)

type ContainerDriverChangeHandlerFunc func(obj *v3.ContainerDriver) (runtime.Object, error)

type ContainerDriverLister interface {
	List(namespace string, selector labels.Selector) (ret []*v3.ContainerDriver, err error)
	Get(namespace, name string) (*v3.ContainerDriver, error)
}

type ContainerDriverController interface {
	Generic() controller.GenericController
	Informer() cache.SharedIndexInformer
	Lister() ContainerDriverLister
	AddHandler(ctx context.Context, name string, handler ContainerDriverHandlerFunc)
	AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync ContainerDriverHandlerFunc)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, handler ContainerDriverHandlerFunc)
	AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, handler ContainerDriverHandlerFunc)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, after time.Duration)
}

type ContainerDriverInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*v3.ContainerDriver) (*v3.ContainerDriver, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*v3.ContainerDriver, error)
	Get(name string, opts metav1.GetOptions) (*v3.ContainerDriver, error)
	Update(*v3.ContainerDriver) (*v3.ContainerDriver, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*v3.ContainerDriverList, error)
	ListNamespaced(namespace string, opts metav1.ListOptions) (*v3.ContainerDriverList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() ContainerDriverController
	AddHandler(ctx context.Context, name string, sync ContainerDriverHandlerFunc)
	AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync ContainerDriverHandlerFunc)
	AddLifecycle(ctx context.Context, name string, lifecycle ContainerDriverLifecycle)
	AddFeatureLifecycle(ctx context.Context, enabled func() bool, name string, lifecycle ContainerDriverLifecycle)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync ContainerDriverHandlerFunc)
	AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, sync ContainerDriverHandlerFunc)
	AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle ContainerDriverLifecycle)
	AddClusterScopedFeatureLifecycle(ctx context.Context, enabled func() bool, name, clusterName string, lifecycle ContainerDriverLifecycle)
}

type containerDriverLister struct {
	ns         string
	controller *containerDriverController
}

func (l *containerDriverLister) List(namespace string, selector labels.Selector) (ret []*v3.ContainerDriver, err error) {
	if namespace == "" {
		namespace = l.ns
	}
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*v3.ContainerDriver))
	})
	return
}

func (l *containerDriverLister) Get(namespace, name string) (*v3.ContainerDriver, error) {
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
			Group:    ContainerDriverGroupVersionKind.Group,
			Resource: ContainerDriverGroupVersionResource.Resource,
		}, key)
	}
	return obj.(*v3.ContainerDriver), nil
}

type containerDriverController struct {
	ns string
	controller.GenericController
}

func (c *containerDriverController) Generic() controller.GenericController {
	return c.GenericController
}

func (c *containerDriverController) Lister() ContainerDriverLister {
	return &containerDriverLister{
		ns:         c.ns,
		controller: c,
	}
}

func (c *containerDriverController) AddHandler(ctx context.Context, name string, handler ContainerDriverHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v3.ContainerDriver); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *containerDriverController) AddFeatureHandler(ctx context.Context, enabled func() bool, name string, handler ContainerDriverHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled() {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v3.ContainerDriver); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *containerDriverController) AddClusterScopedHandler(ctx context.Context, name, cluster string, handler ContainerDriverHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v3.ContainerDriver); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *containerDriverController) AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, cluster string, handler ContainerDriverHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled() {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v3.ContainerDriver); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

type containerDriverFactory struct {
}

func (c containerDriverFactory) Object() runtime.Object {
	return &v3.ContainerDriver{}
}

func (c containerDriverFactory) List() runtime.Object {
	return &v3.ContainerDriverList{}
}

func (s *containerDriverClient) Controller() ContainerDriverController {
	genericController := controller.NewGenericController(s.ns, ContainerDriverGroupVersionKind.Kind+"Controller",
		s.client.controllerFactory.ForResourceKind(ContainerDriverGroupVersionResource, ContainerDriverGroupVersionKind.Kind, false))

	return &containerDriverController{
		ns:                s.ns,
		GenericController: genericController,
	}
}

type containerDriverClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   ContainerDriverController
}

func (s *containerDriverClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *containerDriverClient) Create(o *v3.ContainerDriver) (*v3.ContainerDriver, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*v3.ContainerDriver), err
}

func (s *containerDriverClient) Get(name string, opts metav1.GetOptions) (*v3.ContainerDriver, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*v3.ContainerDriver), err
}

func (s *containerDriverClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*v3.ContainerDriver, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*v3.ContainerDriver), err
}

func (s *containerDriverClient) Update(o *v3.ContainerDriver) (*v3.ContainerDriver, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*v3.ContainerDriver), err
}

func (s *containerDriverClient) UpdateStatus(o *v3.ContainerDriver) (*v3.ContainerDriver, error) {
	obj, err := s.objectClient.UpdateStatus(o.Name, o)
	return obj.(*v3.ContainerDriver), err
}

func (s *containerDriverClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *containerDriverClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *containerDriverClient) List(opts metav1.ListOptions) (*v3.ContainerDriverList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*v3.ContainerDriverList), err
}

func (s *containerDriverClient) ListNamespaced(namespace string, opts metav1.ListOptions) (*v3.ContainerDriverList, error) {
	obj, err := s.objectClient.ListNamespaced(namespace, opts)
	return obj.(*v3.ContainerDriverList), err
}

func (s *containerDriverClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *containerDriverClient) Patch(o *v3.ContainerDriver, patchType types.PatchType, data []byte, subresources ...string) (*v3.ContainerDriver, error) {
	obj, err := s.objectClient.Patch(o.Name, o, patchType, data, subresources...)
	return obj.(*v3.ContainerDriver), err
}

func (s *containerDriverClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *containerDriverClient) AddHandler(ctx context.Context, name string, sync ContainerDriverHandlerFunc) {
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *containerDriverClient) AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync ContainerDriverHandlerFunc) {
	s.Controller().AddFeatureHandler(ctx, enabled, name, sync)
}

func (s *containerDriverClient) AddLifecycle(ctx context.Context, name string, lifecycle ContainerDriverLifecycle) {
	sync := NewContainerDriverLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *containerDriverClient) AddFeatureLifecycle(ctx context.Context, enabled func() bool, name string, lifecycle ContainerDriverLifecycle) {
	sync := NewContainerDriverLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddFeatureHandler(ctx, enabled, name, sync)
}

func (s *containerDriverClient) AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync ContainerDriverHandlerFunc) {
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *containerDriverClient) AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, sync ContainerDriverHandlerFunc) {
	s.Controller().AddClusterScopedFeatureHandler(ctx, enabled, name, clusterName, sync)
}

func (s *containerDriverClient) AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle ContainerDriverLifecycle) {
	sync := NewContainerDriverLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *containerDriverClient) AddClusterScopedFeatureLifecycle(ctx context.Context, enabled func() bool, name, clusterName string, lifecycle ContainerDriverLifecycle) {
	sync := NewContainerDriverLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedFeatureHandler(ctx, enabled, name, clusterName, sync)
}

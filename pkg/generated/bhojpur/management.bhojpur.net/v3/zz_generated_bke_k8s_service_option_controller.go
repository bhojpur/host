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
	BkeK8sServiceOptionGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "BkeK8sServiceOption",
	}
	BkeK8sServiceOptionResource = metav1.APIResource{
		Name:         "bkek8sserviceoptions",
		SingularName: "bkek8sserviceoption",
		Namespaced:   true,

		Kind: BkeK8sServiceOptionGroupVersionKind.Kind,
	}

	BkeK8sServiceOptionGroupVersionResource = schema.GroupVersionResource{
		Group:    GroupName,
		Version:  Version,
		Resource: "bkek8sserviceoptions",
	}
)

func init() {
	resource.Put(BkeK8sServiceOptionGroupVersionResource)
}

// Deprecated use v3.BkeK8sServiceOption instead
type BkeK8sServiceOption = v3.BkeK8sServiceOption

func NewBkeK8sServiceOption(namespace, name string, obj v3.BkeK8sServiceOption) *v3.BkeK8sServiceOption {
	obj.APIVersion, obj.Kind = BkeK8sServiceOptionGroupVersionKind.ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

type BkeK8sServiceOptionHandlerFunc func(key string, obj *v3.BkeK8sServiceOption) (runtime.Object, error)

type BkeK8sServiceOptionChangeHandlerFunc func(obj *v3.BkeK8sServiceOption) (runtime.Object, error)

type BkeK8sServiceOptionLister interface {
	List(namespace string, selector labels.Selector) (ret []*v3.BkeK8sServiceOption, err error)
	Get(namespace, name string) (*v3.BkeK8sServiceOption, error)
}

type BkeK8sServiceOptionController interface {
	Generic() controller.GenericController
	Informer() cache.SharedIndexInformer
	Lister() BkeK8sServiceOptionLister
	AddHandler(ctx context.Context, name string, handler BkeK8sServiceOptionHandlerFunc)
	AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync BkeK8sServiceOptionHandlerFunc)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, handler BkeK8sServiceOptionHandlerFunc)
	AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, handler BkeK8sServiceOptionHandlerFunc)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, after time.Duration)
}

type BkeK8sServiceOptionInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*v3.BkeK8sServiceOption) (*v3.BkeK8sServiceOption, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*v3.BkeK8sServiceOption, error)
	Get(name string, opts metav1.GetOptions) (*v3.BkeK8sServiceOption, error)
	Update(*v3.BkeK8sServiceOption) (*v3.BkeK8sServiceOption, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*v3.BkeK8sServiceOptionList, error)
	ListNamespaced(namespace string, opts metav1.ListOptions) (*v3.BkeK8sServiceOptionList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() BkeK8sServiceOptionController
	AddHandler(ctx context.Context, name string, sync BkeK8sServiceOptionHandlerFunc)
	AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync BkeK8sServiceOptionHandlerFunc)
	AddLifecycle(ctx context.Context, name string, lifecycle BkeK8sServiceOptionLifecycle)
	AddFeatureLifecycle(ctx context.Context, enabled func() bool, name string, lifecycle BkeK8sServiceOptionLifecycle)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync BkeK8sServiceOptionHandlerFunc)
	AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, sync BkeK8sServiceOptionHandlerFunc)
	AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle BkeK8sServiceOptionLifecycle)
	AddClusterScopedFeatureLifecycle(ctx context.Context, enabled func() bool, name, clusterName string, lifecycle BkeK8sServiceOptionLifecycle)
}

type bkeK8sServiceOptionLister struct {
	ns         string
	controller *bkeK8sServiceOptionController
}

func (l *bkeK8sServiceOptionLister) List(namespace string, selector labels.Selector) (ret []*v3.BkeK8sServiceOption, err error) {
	if namespace == "" {
		namespace = l.ns
	}
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*v3.BkeK8sServiceOption))
	})
	return
}

func (l *bkeK8sServiceOptionLister) Get(namespace, name string) (*v3.BkeK8sServiceOption, error) {
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
			Group:    BkeK8sServiceOptionGroupVersionKind.Group,
			Resource: BkeK8sServiceOptionGroupVersionResource.Resource,
		}, key)
	}
	return obj.(*v3.BkeK8sServiceOption), nil
}

type bkeK8sServiceOptionController struct {
	ns string
	controller.GenericController
}

func (c *bkeK8sServiceOptionController) Generic() controller.GenericController {
	return c.GenericController
}

func (c *bkeK8sServiceOptionController) Lister() BkeK8sServiceOptionLister {
	return &bkeK8sServiceOptionLister{
		ns:         c.ns,
		controller: c,
	}
}

func (c *bkeK8sServiceOptionController) AddHandler(ctx context.Context, name string, handler BkeK8sServiceOptionHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v3.BkeK8sServiceOption); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *bkeK8sServiceOptionController) AddFeatureHandler(ctx context.Context, enabled func() bool, name string, handler BkeK8sServiceOptionHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled() {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v3.BkeK8sServiceOption); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *bkeK8sServiceOptionController) AddClusterScopedHandler(ctx context.Context, name, cluster string, handler BkeK8sServiceOptionHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v3.BkeK8sServiceOption); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *bkeK8sServiceOptionController) AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, cluster string, handler BkeK8sServiceOptionHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled() {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v3.BkeK8sServiceOption); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

type bkeK8sServiceOptionFactory struct {
}

func (c bkeK8sServiceOptionFactory) Object() runtime.Object {
	return &v3.BkeK8sServiceOption{}
}

func (c bkeK8sServiceOptionFactory) List() runtime.Object {
	return &v3.BkeK8sServiceOptionList{}
}

func (s *bkeK8sServiceOptionClient) Controller() BkeK8sServiceOptionController {
	genericController := controller.NewGenericController(s.ns, BkeK8sServiceOptionGroupVersionKind.Kind+"Controller",
		s.client.controllerFactory.ForResourceKind(BkeK8sServiceOptionGroupVersionResource, BkeK8sServiceOptionGroupVersionKind.Kind, true))

	return &bkeK8sServiceOptionController{
		ns:                s.ns,
		GenericController: genericController,
	}
}

type bkeK8sServiceOptionClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   BkeK8sServiceOptionController
}

func (s *bkeK8sServiceOptionClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *bkeK8sServiceOptionClient) Create(o *v3.BkeK8sServiceOption) (*v3.BkeK8sServiceOption, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*v3.BkeK8sServiceOption), err
}

func (s *bkeK8sServiceOptionClient) Get(name string, opts metav1.GetOptions) (*v3.BkeK8sServiceOption, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*v3.BkeK8sServiceOption), err
}

func (s *bkeK8sServiceOptionClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*v3.BkeK8sServiceOption, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*v3.BkeK8sServiceOption), err
}

func (s *bkeK8sServiceOptionClient) Update(o *v3.BkeK8sServiceOption) (*v3.BkeK8sServiceOption, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*v3.BkeK8sServiceOption), err
}

func (s *bkeK8sServiceOptionClient) UpdateStatus(o *v3.BkeK8sServiceOption) (*v3.BkeK8sServiceOption, error) {
	obj, err := s.objectClient.UpdateStatus(o.Name, o)
	return obj.(*v3.BkeK8sServiceOption), err
}

func (s *bkeK8sServiceOptionClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *bkeK8sServiceOptionClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *bkeK8sServiceOptionClient) List(opts metav1.ListOptions) (*v3.BkeK8sServiceOptionList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*v3.BkeK8sServiceOptionList), err
}

func (s *bkeK8sServiceOptionClient) ListNamespaced(namespace string, opts metav1.ListOptions) (*v3.BkeK8sServiceOptionList, error) {
	obj, err := s.objectClient.ListNamespaced(namespace, opts)
	return obj.(*v3.BkeK8sServiceOptionList), err
}

func (s *bkeK8sServiceOptionClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *bkeK8sServiceOptionClient) Patch(o *v3.BkeK8sServiceOption, patchType types.PatchType, data []byte, subresources ...string) (*v3.BkeK8sServiceOption, error) {
	obj, err := s.objectClient.Patch(o.Name, o, patchType, data, subresources...)
	return obj.(*v3.BkeK8sServiceOption), err
}

func (s *bkeK8sServiceOptionClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *bkeK8sServiceOptionClient) AddHandler(ctx context.Context, name string, sync BkeK8sServiceOptionHandlerFunc) {
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *bkeK8sServiceOptionClient) AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync BkeK8sServiceOptionHandlerFunc) {
	s.Controller().AddFeatureHandler(ctx, enabled, name, sync)
}

func (s *bkeK8sServiceOptionClient) AddLifecycle(ctx context.Context, name string, lifecycle BkeK8sServiceOptionLifecycle) {
	sync := NewBkeK8sServiceOptionLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *bkeK8sServiceOptionClient) AddFeatureLifecycle(ctx context.Context, enabled func() bool, name string, lifecycle BkeK8sServiceOptionLifecycle) {
	sync := NewBkeK8sServiceOptionLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddFeatureHandler(ctx, enabled, name, sync)
}

func (s *bkeK8sServiceOptionClient) AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync BkeK8sServiceOptionHandlerFunc) {
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *bkeK8sServiceOptionClient) AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, sync BkeK8sServiceOptionHandlerFunc) {
	s.Controller().AddClusterScopedFeatureHandler(ctx, enabled, name, clusterName, sync)
}

func (s *bkeK8sServiceOptionClient) AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle BkeK8sServiceOptionLifecycle) {
	sync := NewBkeK8sServiceOptionLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *bkeK8sServiceOptionClient) AddClusterScopedFeatureLifecycle(ctx context.Context, enabled func() bool, name, clusterName string, lifecycle BkeK8sServiceOptionLifecycle) {
	sync := NewBkeK8sServiceOptionLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedFeatureHandler(ctx, enabled, name, clusterName, sync)
}

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
	BkeK8sSystemImageGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "BkeK8sSystemImage",
	}
	BkeK8sSystemImageResource = metav1.APIResource{
		Name:         "bkek8ssystemimages",
		SingularName: "bkek8ssystemimage",
		Namespaced:   true,

		Kind: BkeK8sSystemImageGroupVersionKind.Kind,
	}

	BkeK8sSystemImageGroupVersionResource = schema.GroupVersionResource{
		Group:    GroupName,
		Version:  Version,
		Resource: "bkek8ssystemimages",
	}
)

func init() {
	resource.Put(BkeK8sSystemImageGroupVersionResource)
}

// Deprecated use v3.BkeK8sSystemImage instead
type BkeK8sSystemImage = v3.BkeK8sSystemImage

func NewBkeK8sSystemImage(namespace, name string, obj v3.BkeK8sSystemImage) *v3.BkeK8sSystemImage {
	obj.APIVersion, obj.Kind = BkeK8sSystemImageGroupVersionKind.ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

type BkeK8sSystemImageHandlerFunc func(key string, obj *v3.BkeK8sSystemImage) (runtime.Object, error)

type BkeK8sSystemImageChangeHandlerFunc func(obj *v3.BkeK8sSystemImage) (runtime.Object, error)

type BkeK8sSystemImageLister interface {
	List(namespace string, selector labels.Selector) (ret []*v3.BkeK8sSystemImage, err error)
	Get(namespace, name string) (*v3.BkeK8sSystemImage, error)
}

type BkeK8sSystemImageController interface {
	Generic() controller.GenericController
	Informer() cache.SharedIndexInformer
	Lister() BkeK8sSystemImageLister
	AddHandler(ctx context.Context, name string, handler BkeK8sSystemImageHandlerFunc)
	AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync BkeK8sSystemImageHandlerFunc)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, handler BkeK8sSystemImageHandlerFunc)
	AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, handler BkeK8sSystemImageHandlerFunc)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, after time.Duration)
}

type BkeK8sSystemImageInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*v3.BkeK8sSystemImage) (*v3.BkeK8sSystemImage, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*v3.BkeK8sSystemImage, error)
	Get(name string, opts metav1.GetOptions) (*v3.BkeK8sSystemImage, error)
	Update(*v3.BkeK8sSystemImage) (*v3.BkeK8sSystemImage, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*v3.BkeK8sSystemImageList, error)
	ListNamespaced(namespace string, opts metav1.ListOptions) (*v3.BkeK8sSystemImageList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() BkeK8sSystemImageController
	AddHandler(ctx context.Context, name string, sync BkeK8sSystemImageHandlerFunc)
	AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync BkeK8sSystemImageHandlerFunc)
	AddLifecycle(ctx context.Context, name string, lifecycle BkeK8sSystemImageLifecycle)
	AddFeatureLifecycle(ctx context.Context, enabled func() bool, name string, lifecycle BkeK8sSystemImageLifecycle)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync BkeK8sSystemImageHandlerFunc)
	AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, sync BkeK8sSystemImageHandlerFunc)
	AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle BkeK8sSystemImageLifecycle)
	AddClusterScopedFeatureLifecycle(ctx context.Context, enabled func() bool, name, clusterName string, lifecycle BkeK8sSystemImageLifecycle)
}

type bkeK8sSystemImageLister struct {
	ns         string
	controller *bkeK8sSystemImageController
}

func (l *bkeK8sSystemImageLister) List(namespace string, selector labels.Selector) (ret []*v3.BkeK8sSystemImage, err error) {
	if namespace == "" {
		namespace = l.ns
	}
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*v3.BkeK8sSystemImage))
	})
	return
}

func (l *bkeK8sSystemImageLister) Get(namespace, name string) (*v3.BkeK8sSystemImage, error) {
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
			Group:    BkeK8sSystemImageGroupVersionKind.Group,
			Resource: BkeK8sSystemImageGroupVersionResource.Resource,
		}, key)
	}
	return obj.(*v3.BkeK8sSystemImage), nil
}

type bkeK8sSystemImageController struct {
	ns string
	controller.GenericController
}

func (c *bkeK8sSystemImageController) Generic() controller.GenericController {
	return c.GenericController
}

func (c *bkeK8sSystemImageController) Lister() BkeK8sSystemImageLister {
	return &bkeK8sSystemImageLister{
		ns:         c.ns,
		controller: c,
	}
}

func (c *bkeK8sSystemImageController) AddHandler(ctx context.Context, name string, handler BkeK8sSystemImageHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v3.BkeK8sSystemImage); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *bkeK8sSystemImageController) AddFeatureHandler(ctx context.Context, enabled func() bool, name string, handler BkeK8sSystemImageHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled() {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v3.BkeK8sSystemImage); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *bkeK8sSystemImageController) AddClusterScopedHandler(ctx context.Context, name, cluster string, handler BkeK8sSystemImageHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v3.BkeK8sSystemImage); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *bkeK8sSystemImageController) AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, cluster string, handler BkeK8sSystemImageHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled() {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v3.BkeK8sSystemImage); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

type bkeK8sSystemImageFactory struct {
}

func (c bkeK8sSystemImageFactory) Object() runtime.Object {
	return &v3.BkeK8sSystemImage{}
}

func (c bkeK8sSystemImageFactory) List() runtime.Object {
	return &v3.BkeK8sSystemImageList{}
}

func (s *bkeK8sSystemImageClient) Controller() BkeK8sSystemImageController {
	genericController := controller.NewGenericController(s.ns, BkeK8sSystemImageGroupVersionKind.Kind+"Controller",
		s.client.controllerFactory.ForResourceKind(BkeK8sSystemImageGroupVersionResource, BkeK8sSystemImageGroupVersionKind.Kind, true))

	return &bkeK8sSystemImageController{
		ns:                s.ns,
		GenericController: genericController,
	}
}

type bkeK8sSystemImageClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   BkeK8sSystemImageController
}

func (s *bkeK8sSystemImageClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *bkeK8sSystemImageClient) Create(o *v3.BkeK8sSystemImage) (*v3.BkeK8sSystemImage, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*v3.BkeK8sSystemImage), err
}

func (s *bkeK8sSystemImageClient) Get(name string, opts metav1.GetOptions) (*v3.BkeK8sSystemImage, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*v3.BkeK8sSystemImage), err
}

func (s *bkeK8sSystemImageClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*v3.BkeK8sSystemImage, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*v3.BkeK8sSystemImage), err
}

func (s *bkeK8sSystemImageClient) Update(o *v3.BkeK8sSystemImage) (*v3.BkeK8sSystemImage, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*v3.BkeK8sSystemImage), err
}

func (s *bkeK8sSystemImageClient) UpdateStatus(o *v3.BkeK8sSystemImage) (*v3.BkeK8sSystemImage, error) {
	obj, err := s.objectClient.UpdateStatus(o.Name, o)
	return obj.(*v3.BkeK8sSystemImage), err
}

func (s *bkeK8sSystemImageClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *bkeK8sSystemImageClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *bkeK8sSystemImageClient) List(opts metav1.ListOptions) (*v3.BkeK8sSystemImageList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*v3.BkeK8sSystemImageList), err
}

func (s *bkeK8sSystemImageClient) ListNamespaced(namespace string, opts metav1.ListOptions) (*v3.BkeK8sSystemImageList, error) {
	obj, err := s.objectClient.ListNamespaced(namespace, opts)
	return obj.(*v3.BkeK8sSystemImageList), err
}

func (s *bkeK8sSystemImageClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *bkeK8sSystemImageClient) Patch(o *v3.BkeK8sSystemImage, patchType types.PatchType, data []byte, subresources ...string) (*v3.BkeK8sSystemImage, error) {
	obj, err := s.objectClient.Patch(o.Name, o, patchType, data, subresources...)
	return obj.(*v3.BkeK8sSystemImage), err
}

func (s *bkeK8sSystemImageClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *bkeK8sSystemImageClient) AddHandler(ctx context.Context, name string, sync BkeK8sSystemImageHandlerFunc) {
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *bkeK8sSystemImageClient) AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync BkeK8sSystemImageHandlerFunc) {
	s.Controller().AddFeatureHandler(ctx, enabled, name, sync)
}

func (s *bkeK8sSystemImageClient) AddLifecycle(ctx context.Context, name string, lifecycle BkeK8sSystemImageLifecycle) {
	sync := NewBkeK8sSystemImageLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *bkeK8sSystemImageClient) AddFeatureLifecycle(ctx context.Context, enabled func() bool, name string, lifecycle BkeK8sSystemImageLifecycle) {
	sync := NewBkeK8sSystemImageLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddFeatureHandler(ctx, enabled, name, sync)
}

func (s *bkeK8sSystemImageClient) AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync BkeK8sSystemImageHandlerFunc) {
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *bkeK8sSystemImageClient) AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, sync BkeK8sSystemImageHandlerFunc) {
	s.Controller().AddClusterScopedFeatureHandler(ctx, enabled, name, clusterName, sync)
}

func (s *bkeK8sSystemImageClient) AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle BkeK8sSystemImageLifecycle) {
	sync := NewBkeK8sSystemImageLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *bkeK8sSystemImageClient) AddClusterScopedFeatureLifecycle(ctx context.Context, enabled func() bool, name, clusterName string, lifecycle BkeK8sSystemImageLifecycle) {
	sync := NewBkeK8sSystemImageLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedFeatureHandler(ctx, enabled, name, clusterName, sync)
}

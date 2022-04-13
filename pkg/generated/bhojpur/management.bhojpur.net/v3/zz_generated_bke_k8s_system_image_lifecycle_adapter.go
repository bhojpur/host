package v3

import (
	"github.com/bhojpur/host/pkg/apis/management.bhojpur.net/v3"
	"github.com/bhojpur/host/pkg/core/lifecycle"
	"github.com/bhojpur/host/pkg/core/resource"
	"k8s.io/apimachinery/pkg/runtime"
)

type BkeK8sSystemImageLifecycle interface {
	Create(obj *v3.BkeK8sSystemImage) (runtime.Object, error)
	Remove(obj *v3.BkeK8sSystemImage) (runtime.Object, error)
	Updated(obj *v3.BkeK8sSystemImage) (runtime.Object, error)
}

type bkeK8sSystemImageLifecycleAdapter struct {
	lifecycle BkeK8sSystemImageLifecycle
}

func (w *bkeK8sSystemImageLifecycleAdapter) HasCreate() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasCreate()
}

func (w *bkeK8sSystemImageLifecycleAdapter) HasFinalize() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasFinalize()
}

func (w *bkeK8sSystemImageLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*v3.BkeK8sSystemImage))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *bkeK8sSystemImageLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*v3.BkeK8sSystemImage))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *bkeK8sSystemImageLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*v3.BkeK8sSystemImage))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewBkeK8sSystemImageLifecycleAdapter(name string, clusterScoped bool, client BkeK8sSystemImageInterface, l BkeK8sSystemImageLifecycle) BkeK8sSystemImageHandlerFunc {
	if clusterScoped {
		resource.PutClusterScoped(BkeK8sSystemImageGroupVersionResource)
	}
	adapter := &bkeK8sSystemImageLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *v3.BkeK8sSystemImage) (runtime.Object, error) {
		newObj, err := syncFn(key, obj)
		if o, ok := newObj.(runtime.Object); ok {
			return o, err
		}
		return nil, err
	}
}

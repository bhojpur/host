package v3

import (
	"github.com/bhojpur/host/pkg/apis/management.bhojpur.net/v3"
	"github.com/bhojpur/host/pkg/core/lifecycle"
	"github.com/bhojpur/host/pkg/core/resource"
	"k8s.io/apimachinery/pkg/runtime"
)

type BkeK8sServiceOptionLifecycle interface {
	Create(obj *v3.BkeK8sServiceOption) (runtime.Object, error)
	Remove(obj *v3.BkeK8sServiceOption) (runtime.Object, error)
	Updated(obj *v3.BkeK8sServiceOption) (runtime.Object, error)
}

type bkeK8sServiceOptionLifecycleAdapter struct {
	lifecycle BkeK8sServiceOptionLifecycle
}

func (w *bkeK8sServiceOptionLifecycleAdapter) HasCreate() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasCreate()
}

func (w *bkeK8sServiceOptionLifecycleAdapter) HasFinalize() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasFinalize()
}

func (w *bkeK8sServiceOptionLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*v3.BkeK8sServiceOption))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *bkeK8sServiceOptionLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*v3.BkeK8sServiceOption))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *bkeK8sServiceOptionLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*v3.BkeK8sServiceOption))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewBkeK8sServiceOptionLifecycleAdapter(name string, clusterScoped bool, client BkeK8sServiceOptionInterface, l BkeK8sServiceOptionLifecycle) BkeK8sServiceOptionHandlerFunc {
	if clusterScoped {
		resource.PutClusterScoped(BkeK8sServiceOptionGroupVersionResource)
	}
	adapter := &bkeK8sServiceOptionLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *v3.BkeK8sServiceOption) (runtime.Object, error) {
		newObj, err := syncFn(key, obj)
		if o, ok := newObj.(runtime.Object); ok {
			return o, err
		}
		return nil, err
	}
}

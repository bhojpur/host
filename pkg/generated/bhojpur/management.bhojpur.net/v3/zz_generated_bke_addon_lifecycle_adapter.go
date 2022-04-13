package v3

import (
	"github.com/bhojpur/host/pkg/apis/management.bhojpur.net/v3"
	"github.com/bhojpur/host/pkg/core/lifecycle"
	"github.com/bhojpur/host/pkg/core/resource"
	"k8s.io/apimachinery/pkg/runtime"
)

type BkeAddonLifecycle interface {
	Create(obj *v3.BkeAddon) (runtime.Object, error)
	Remove(obj *v3.BkeAddon) (runtime.Object, error)
	Updated(obj *v3.BkeAddon) (runtime.Object, error)
}

type bkeAddonLifecycleAdapter struct {
	lifecycle BkeAddonLifecycle
}

func (w *bkeAddonLifecycleAdapter) HasCreate() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasCreate()
}

func (w *bkeAddonLifecycleAdapter) HasFinalize() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasFinalize()
}

func (w *bkeAddonLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*v3.BkeAddon))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *bkeAddonLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*v3.BkeAddon))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *bkeAddonLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*v3.BkeAddon))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewBkeAddonLifecycleAdapter(name string, clusterScoped bool, client BkeAddonInterface, l BkeAddonLifecycle) BkeAddonHandlerFunc {
	if clusterScoped {
		resource.PutClusterScoped(BkeAddonGroupVersionResource)
	}
	adapter := &bkeAddonLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *v3.BkeAddon) (runtime.Object, error) {
		newObj, err := syncFn(key, obj)
		if o, ok := newObj.(runtime.Object); ok {
			return o, err
		}
		return nil, err
	}
}

package v3

import (
	"github.com/bhojpur/host/pkg/apis/management.bhojpur.net/v3"
	"github.com/bhojpur/host/pkg/core/lifecycle"
	"github.com/bhojpur/host/pkg/core/resource"
	"k8s.io/apimachinery/pkg/runtime"
)

type ContainerDriverLifecycle interface {
	Create(obj *v3.ContainerDriver) (runtime.Object, error)
	Remove(obj *v3.ContainerDriver) (runtime.Object, error)
	Updated(obj *v3.ContainerDriver) (runtime.Object, error)
}

type containerDriverLifecycleAdapter struct {
	lifecycle ContainerDriverLifecycle
}

func (w *containerDriverLifecycleAdapter) HasCreate() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasCreate()
}

func (w *containerDriverLifecycleAdapter) HasFinalize() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasFinalize()
}

func (w *containerDriverLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*v3.ContainerDriver))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *containerDriverLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*v3.ContainerDriver))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *containerDriverLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*v3.ContainerDriver))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewContainerDriverLifecycleAdapter(name string, clusterScoped bool, client ContainerDriverInterface, l ContainerDriverLifecycle) ContainerDriverHandlerFunc {
	if clusterScoped {
		resource.PutClusterScoped(ContainerDriverGroupVersionResource)
	}
	adapter := &containerDriverLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *v3.ContainerDriver) (runtime.Object, error) {
		newObj, err := syncFn(key, obj)
		if o, ok := newObj.(runtime.Object); ok {
			return o, err
		}
		return nil, err
	}
}

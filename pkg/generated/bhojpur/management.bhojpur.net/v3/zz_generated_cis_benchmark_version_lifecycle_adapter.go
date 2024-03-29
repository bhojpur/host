package v3

import (
	"github.com/bhojpur/host/pkg/apis/management.bhojpur.net/v3"
	"github.com/bhojpur/host/pkg/core/lifecycle"
	"github.com/bhojpur/host/pkg/core/resource"
	"k8s.io/apimachinery/pkg/runtime"
)

type CisBenchmarkVersionLifecycle interface {
	Create(obj *v3.CisBenchmarkVersion) (runtime.Object, error)
	Remove(obj *v3.CisBenchmarkVersion) (runtime.Object, error)
	Updated(obj *v3.CisBenchmarkVersion) (runtime.Object, error)
}

type cisBenchmarkVersionLifecycleAdapter struct {
	lifecycle CisBenchmarkVersionLifecycle
}

func (w *cisBenchmarkVersionLifecycleAdapter) HasCreate() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasCreate()
}

func (w *cisBenchmarkVersionLifecycleAdapter) HasFinalize() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasFinalize()
}

func (w *cisBenchmarkVersionLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*v3.CisBenchmarkVersion))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *cisBenchmarkVersionLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*v3.CisBenchmarkVersion))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *cisBenchmarkVersionLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*v3.CisBenchmarkVersion))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewCisBenchmarkVersionLifecycleAdapter(name string, clusterScoped bool, client CisBenchmarkVersionInterface, l CisBenchmarkVersionLifecycle) CisBenchmarkVersionHandlerFunc {
	if clusterScoped {
		resource.PutClusterScoped(CisBenchmarkVersionGroupVersionResource)
	}
	adapter := &cisBenchmarkVersionLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *v3.CisBenchmarkVersion) (runtime.Object, error) {
		newObj, err := syncFn(key, obj)
		if o, ok := newObj.(runtime.Object); ok {
			return o, err
		}
		return nil, err
	}
}

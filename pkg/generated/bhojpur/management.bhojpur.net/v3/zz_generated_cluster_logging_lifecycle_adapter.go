package v3

import (
	"github.com/bhojpur/host/pkg/apis/management.bhojpur.net/v3"
	"github.com/bhojpur/host/pkg/core/lifecycle"
	"github.com/bhojpur/host/pkg/core/resource"
	"k8s.io/apimachinery/pkg/runtime"
)

type ClusterLoggingLifecycle interface {
	Create(obj *v3.ClusterLogging) (runtime.Object, error)
	Remove(obj *v3.ClusterLogging) (runtime.Object, error)
	Updated(obj *v3.ClusterLogging) (runtime.Object, error)
}

type clusterLoggingLifecycleAdapter struct {
	lifecycle ClusterLoggingLifecycle
}

func (w *clusterLoggingLifecycleAdapter) HasCreate() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasCreate()
}

func (w *clusterLoggingLifecycleAdapter) HasFinalize() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasFinalize()
}

func (w *clusterLoggingLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*v3.ClusterLogging))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *clusterLoggingLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*v3.ClusterLogging))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *clusterLoggingLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*v3.ClusterLogging))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewClusterLoggingLifecycleAdapter(name string, clusterScoped bool, client ClusterLoggingInterface, l ClusterLoggingLifecycle) ClusterLoggingHandlerFunc {
	if clusterScoped {
		resource.PutClusterScoped(ClusterLoggingGroupVersionResource)
	}
	adapter := &clusterLoggingLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *v3.ClusterLogging) (runtime.Object, error) {
		newObj, err := syncFn(key, obj)
		if o, ok := newObj.(runtime.Object); ok {
			return o, err
		}
		return nil, err
	}
}

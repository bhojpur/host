package v3

import (
	"github.com/bhojpur/host/pkg/apis/project.bhojpur.net/v3"
	"github.com/bhojpur/host/pkg/core/lifecycle"
	"github.com/bhojpur/host/pkg/core/resource"
	"k8s.io/apimachinery/pkg/runtime"
)

type SSHAuthLifecycle interface {
	Create(obj *v3.SSHAuth) (runtime.Object, error)
	Remove(obj *v3.SSHAuth) (runtime.Object, error)
	Updated(obj *v3.SSHAuth) (runtime.Object, error)
}

type sshAuthLifecycleAdapter struct {
	lifecycle SSHAuthLifecycle
}

func (w *sshAuthLifecycleAdapter) HasCreate() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasCreate()
}

func (w *sshAuthLifecycleAdapter) HasFinalize() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasFinalize()
}

func (w *sshAuthLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*v3.SSHAuth))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *sshAuthLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*v3.SSHAuth))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *sshAuthLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*v3.SSHAuth))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewSSHAuthLifecycleAdapter(name string, clusterScoped bool, client SSHAuthInterface, l SSHAuthLifecycle) SSHAuthHandlerFunc {
	if clusterScoped {
		resource.PutClusterScoped(SSHAuthGroupVersionResource)
	}
	adapter := &sshAuthLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *v3.SSHAuth) (runtime.Object, error) {
		newObj, err := syncFn(key, obj)
		if o, ok := newObj.(runtime.Object); ok {
			return o, err
		}
		return nil, err
	}
}

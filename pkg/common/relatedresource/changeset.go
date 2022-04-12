package relatedresource

// Copyright (c) 2018 Bhojpur Consulting Private Limited, India. All rights reserved.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

import (
	"context"
	"time"

	"k8s.io/apimachinery/pkg/api/meta"

	"github.com/bhojpur/host/pkg/common/generic"
	"github.com/bhojpur/host/pkg/common/kv"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/cache"
)

type Key struct {
	Namespace string
	Name      string
}

func NewKey(namespace, name string) Key {
	return Key{
		Namespace: namespace,
		Name:      name,
	}
}

func FromString(key string) Key {
	return NewKey(kv.RSplit(key, "/"))
}

type ControllerWrapper interface {
	Informer() cache.SharedIndexInformer
	AddGenericHandler(ctx context.Context, name string, handler generic.Handler)
}

type ClusterScopedEnqueuer interface {
	Enqueue(name string)
}

type Enqueuer interface {
	Enqueue(namespace, name string)
}

type Resolver func(namespace, name string, obj runtime.Object) ([]Key, error)

func WatchClusterScoped(ctx context.Context, name string, resolve Resolver, enq ClusterScopedEnqueuer, watching ...ControllerWrapper) {
	Watch(ctx, name, resolve, &wrapper{ClusterScopedEnqueuer: enq}, watching...)
}

func Watch(ctx context.Context, name string, resolve Resolver, enq Enqueuer, watching ...ControllerWrapper) {
	for _, c := range watching {
		watch(ctx, name, enq, resolve, c)
	}
}

func watch(ctx context.Context, name string, enq Enqueuer, resolve Resolver, controller ControllerWrapper) {
	runResolve := func(ns, name string, obj runtime.Object) error {
		keys, err := resolve(ns, name, obj)
		if err != nil {
			return err
		}

		for _, key := range keys {
			if key.Name != "" {
				enq.Enqueue(key.Namespace, key.Name)
			}
		}

		return nil
	}

	controller.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		DeleteFunc: func(obj interface{}) {
			ro, ok := obj.(runtime.Object)
			if !ok {
				return
			}

			meta, err := meta.Accessor(ro)
			if err != nil {
				return
			}

			go func() {
				time.Sleep(time.Second)
				runResolve(meta.GetNamespace(), meta.GetName(), ro)
			}()
		},
	})

	controller.AddGenericHandler(ctx, name, func(key string, obj runtime.Object) (runtime.Object, error) {
		ns, name := kv.RSplit(key, "/")
		return obj, runResolve(ns, name, obj)
	})
}

type wrapper struct {
	ClusterScopedEnqueuer
}

func (w *wrapper) Enqueue(namespace, name string) {
	w.ClusterScopedEnqueuer.Enqueue(name)
}

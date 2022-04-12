package generator

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

var lifecycleTemplate = `package {{.schema.Version.Version}}

import (
	{{.importPackage}}
	"k8s.io/apimachinery/pkg/runtime"
	"github.com/bhojpur/host/pkg/core/lifecycle"
	"github.com/bhojpur/host/pkg/core/resource"
)

type {{.schema.CodeName}}Lifecycle interface {
	Create(obj *{{.prefix}}{{.schema.CodeName}}) (runtime.Object, error)
	Remove(obj *{{.prefix}}{{.schema.CodeName}}) (runtime.Object, error)
	Updated(obj *{{.prefix}}{{.schema.CodeName}}) (runtime.Object, error)
}

type {{.schema.ID}}LifecycleAdapter struct {
	lifecycle {{.schema.CodeName}}Lifecycle
}

func (w *{{.schema.ID}}LifecycleAdapter) HasCreate() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasCreate()
}

func (w *{{.schema.ID}}LifecycleAdapter) HasFinalize() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasFinalize()
}

func (w *{{.schema.ID}}LifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*{{.prefix}}{{.schema.CodeName}}))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *{{.schema.ID}}LifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*{{.prefix}}{{.schema.CodeName}}))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *{{.schema.ID}}LifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*{{.prefix}}{{.schema.CodeName}}))
	if o == nil {
		return nil, err
	}
	return o, err
}

func New{{.schema.CodeName}}LifecycleAdapter(name string, clusterScoped bool, client {{.schema.CodeName}}Interface, l {{.schema.CodeName}}Lifecycle) {{.schema.CodeName}}HandlerFunc {
	if clusterScoped {
		resource.PutClusterScoped({{.schema.CodeName}}GroupVersionResource)
	}
	adapter := &{{.schema.ID}}LifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *{{.prefix}}{{.schema.CodeName}}) (runtime.Object, error) {
		newObj, err := syncFn(key, obj)
		if o, ok := newObj.(runtime.Object); ok {
			return o, err
		}
		return nil, err
	}
}
`

package fake

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

	"github.com/bhojpur/host/pkg/common/apply"
	"github.com/bhojpur/host/pkg/common/apply/injectors"
	"github.com/bhojpur/host/pkg/common/objectset"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var _ apply.Apply = (*FakeApply)(nil)

type FakeApply struct {
	Objects []*objectset.ObjectSet
}

func (f *FakeApply) Apply(set *objectset.ObjectSet) error {
	f.Objects = append(f.Objects, set)
	return nil
}

func (f *FakeApply) ApplyObjects(objs ...runtime.Object) error {
	os := objectset.NewObjectSet()
	os.Add(objs...)
	f.Objects = append(f.Objects, os)
	return nil
}

func (f *FakeApply) WithCacheTypes(igs ...apply.InformerGetter) apply.Apply {
	return f
}

func (f *FakeApply) WithIgnorePreviousApplied() apply.Apply {
	return f
}

func (f *FakeApply) WithGVK(gvks ...schema.GroupVersionKind) apply.Apply {
	return f
}

func (f *FakeApply) WithSetID(id string) apply.Apply {
	return f
}

func (f *FakeApply) WithOwner(obj runtime.Object) apply.Apply {
	return f
}

func (f *FakeApply) WithInjector(injs ...injectors.ConfigInjector) apply.Apply {
	return f
}

func (f *FakeApply) WithInjectorName(injs ...string) apply.Apply {
	return f
}

func (f *FakeApply) WithPatcher(gvk schema.GroupVersionKind, patchers apply.Patcher) apply.Apply {
	return f
}

func (f *FakeApply) WithReconciler(gvk schema.GroupVersionKind, reconciler apply.Reconciler) apply.Apply {
	return f
}

func (f *FakeApply) WithStrictCaching() apply.Apply {
	return f
}

func (f *FakeApply) WithDynamicLookup() apply.Apply {
	return f
}

func (f *FakeApply) WithDefaultNamespace(ns string) apply.Apply {
	return f
}

func (f *FakeApply) WithListerNamespace(ns string) apply.Apply {
	return f
}

func (f *FakeApply) WithRestrictClusterScoped() apply.Apply {
	return f
}

func (f *FakeApply) WithSetOwnerReference(controller, block bool) apply.Apply {
	return f
}

func (f *FakeApply) WithRateLimiting(ratelimitingQps float32) apply.Apply {
	return f
}

func (f *FakeApply) WithNoDelete() apply.Apply {
	return f
}

func (f *FakeApply) WithNoDeleteGVK(gvks ...schema.GroupVersionKind) apply.Apply {
	return f
}

func (f *FakeApply) WithContext(ctx context.Context) apply.Apply {
	return f
}

func (f *FakeApply) WithCacheTypeFactory(factory apply.InformerFactory) apply.Apply {
	return f
}

func (f *FakeApply) DryRun(objs ...runtime.Object) (apply.Plan, error) {
	return apply.Plan{}, nil
}

func (f *FakeApply) FindOwner(obj runtime.Object) (runtime.Object, error) {
	return nil, nil
}

func (f *FakeApply) PurgeOrphan(obj runtime.Object) error {
	return nil
}

func (f *FakeApply) WithOwnerKey(key string, gvk schema.GroupVersionKind) apply.Apply {
	return f
}

func (f *FakeApply) WithDiffPatch(gvk schema.GroupVersionKind, namespace, name string, patch []byte) apply.Apply {
	return f
}

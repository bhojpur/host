package lister

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
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/cache"
)

var _ cache.GenericLister = &summaryListerShim{}
var _ cache.GenericNamespaceLister = &summaryNamespaceListerShim{}

// summaryListerShim implements the cache.GenericLister interface.
type summaryListerShim struct {
	lister Lister
}

// NewRuntimeObjectShim returns a new shim for Lister.
// It wraps Lister so that it implements cache.GenericLister interface
func NewRuntimeObjectShim(lister Lister) cache.GenericLister {
	return &summaryListerShim{lister: lister}
}

// List will return all objects across namespaces
func (s *summaryListerShim) List(selector labels.Selector) (ret []runtime.Object, err error) {
	objs, err := s.lister.List(selector)
	if err != nil {
		return nil, err
	}

	ret = make([]runtime.Object, len(objs))
	for index, obj := range objs {
		ret[index] = obj
	}
	return ret, err
}

// Get will attempt to retrieve assuming that name==key
func (s *summaryListerShim) Get(name string) (runtime.Object, error) {
	return s.lister.Get(name)
}

func (s *summaryListerShim) ByNamespace(namespace string) cache.GenericNamespaceLister {
	return &summaryNamespaceListerShim{
		namespaceLister: s.lister.Namespace(namespace),
	}
}

// summaryNamespaceListerShim implements the NamespaceLister interface.
// It wraps NamespaceLister so that it implements cache.GenericNamespaceLister interface
type summaryNamespaceListerShim struct {
	namespaceLister NamespaceLister
}

// List will return all objects in this namespace
func (ns *summaryNamespaceListerShim) List(selector labels.Selector) (ret []runtime.Object, err error) {
	objs, err := ns.namespaceLister.List(selector)
	if err != nil {
		return nil, err
	}

	ret = make([]runtime.Object, len(objs))
	for index, obj := range objs {
		ret[index] = obj
	}
	return ret, err
}

// Get will attempt to retrieve by namespace and name
func (ns *summaryNamespaceListerShim) Get(name string) (runtime.Object, error) {
	return ns.namespaceLister.Get(name)
}

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
	"github.com/bhojpur/host/pkg/common/summary"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/tools/cache"
)

var _ Lister = &summaryLister{}
var _ NamespaceLister = &summaryNamespaceLister{}

// summaryLister implements the Lister interface.
type summaryLister struct {
	indexer cache.Indexer
	gvr     schema.GroupVersionResource
}

// New returns a new Lister.
func New(indexer cache.Indexer, gvr schema.GroupVersionResource) Lister {
	return &summaryLister{indexer: indexer, gvr: gvr}
}

// List lists all resources in the indexer.
func (l *summaryLister) List(selector labels.Selector) (ret []*summary.SummarizedObject, err error) {
	err = cache.ListAll(l.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*summary.SummarizedObject))
	})
	return ret, err
}

// Get retrieves a resource from the indexer with the given name
func (l *summaryLister) Get(name string) (*summary.SummarizedObject, error) {
	obj, exists, err := l.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(l.gvr.GroupResource(), name)
	}
	return obj.(*summary.SummarizedObject), nil
}

// Namespace returns an object that can list and get resources from a given namespace.
func (l *summaryLister) Namespace(namespace string) NamespaceLister {
	return &summaryNamespaceLister{indexer: l.indexer, namespace: namespace, gvr: l.gvr}
}

// summaryNamespaceLister implements the NamespaceLister interface.
type summaryNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
	gvr       schema.GroupVersionResource
}

// List lists all resources in the indexer for a given namespace.
func (l *summaryNamespaceLister) List(selector labels.Selector) (ret []*summary.SummarizedObject, err error) {
	err = cache.ListAllByNamespace(l.indexer, l.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*summary.SummarizedObject))
	})
	return ret, err
}

// Get retrieves a resource from the indexer for a given namespace and name.
func (l *summaryNamespaceLister) Get(name string) (*summary.SummarizedObject, error) {
	obj, exists, err := l.indexer.GetByKey(l.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(l.gvr.GroupResource(), name)
	}
	return obj.(*summary.SummarizedObject), nil
}

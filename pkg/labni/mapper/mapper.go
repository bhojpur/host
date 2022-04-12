package mapper

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
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/discovery/cached/memory"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/restmapper"
)

func New(config *rest.Config) (meta.RESTMapper, error) {
	d, err := discovery.NewDiscoveryClientForConfig(config)
	if err != nil {
		return nil, err
	}
	cached := memory.NewMemCacheClient(d)
	return &retryMapper{
		target: restmapper.NewDeferredDiscoveryRESTMapper(cached),
		cache:  cached,
	}, nil
}

type retryMapper struct {
	target meta.RESTMapper
	cache  discovery.CachedDiscoveryInterface
}

func (r *retryMapper) KindFor(resource schema.GroupVersionResource) (schema.GroupVersionKind, error) {
	result, err := r.target.KindFor(resource)
	if err != nil {
		r.cache.Invalidate()
		return r.target.KindFor(resource)
	}
	return result, err
}

func (r *retryMapper) KindsFor(resource schema.GroupVersionResource) ([]schema.GroupVersionKind, error) {
	result, err := r.target.KindsFor(resource)
	if err != nil {
		r.cache.Invalidate()
		return r.target.KindsFor(resource)
	}
	return result, err
}

func (r *retryMapper) ResourceFor(input schema.GroupVersionResource) (schema.GroupVersionResource, error) {
	result, err := r.target.ResourceFor(input)
	if err != nil {
		r.cache.Invalidate()
		return r.target.ResourceFor(input)
	}
	return result, err
}

func (r *retryMapper) ResourcesFor(input schema.GroupVersionResource) ([]schema.GroupVersionResource, error) {
	result, err := r.target.ResourcesFor(input)
	if err != nil {
		r.cache.Invalidate()
		return r.target.ResourcesFor(input)
	}
	return result, err
}

func (r *retryMapper) RESTMapping(gk schema.GroupKind, versions ...string) (*meta.RESTMapping, error) {
	result, err := r.target.RESTMapping(gk, versions...)
	if err != nil {
		r.cache.Invalidate()
		return r.target.RESTMapping(gk, versions...)
	}
	return result, err
}

func (r *retryMapper) RESTMappings(gk schema.GroupKind, versions ...string) ([]*meta.RESTMapping, error) {
	result, err := r.target.RESTMappings(gk, versions...)
	if err != nil {
		r.cache.Invalidate()
		return r.target.RESTMappings(gk, versions...)
	}
	return result, err
}

func (r *retryMapper) ResourceSingularizer(resource string) (singular string, err error) {
	result, err := r.target.ResourceSingularizer(resource)
	if err != nil {
		r.cache.Invalidate()
		return r.target.ResourceSingularizer(resource)
	}
	return result, err
}

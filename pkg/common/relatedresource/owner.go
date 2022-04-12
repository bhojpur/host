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
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
)

// OwnerResolver Look for owner references that match the apiVersion and kind and resolve to the namespace and
// name of the parent. The namespaced flag is whether the apiVersion/kind referenced is expected to be namespaced
func OwnerResolver(namespaced bool, apiVersion, kind string) Resolver {
	return func(namespace, name string, obj runtime.Object) ([]Key, error) {
		if obj == nil {
			return nil, nil
		}

		meta, err := meta.Accessor(obj)
		if err != nil {
			// ignore err
			return nil, nil
		}

		var result []Key
		for _, owner := range meta.GetOwnerReferences() {
			if owner.Kind == kind && owner.APIVersion == apiVersion {
				ns := ""
				if namespaced {
					ns = meta.GetNamespace()
				}
				result = append(result, Key{
					Namespace: ns,
					Name:      owner.Name,
				})
			}
		}

		return result, nil
	}
}

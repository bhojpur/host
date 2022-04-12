package gvk

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
	"fmt"

	"github.com/bhojpur/host/pkg/common/schemes"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func Get(obj runtime.Object) (schema.GroupVersionKind, error) {
	gvk := obj.GetObjectKind().GroupVersionKind()
	if gvk.Kind != "" {
		return gvk, nil
	}

	gvks, _, err := schemes.All.ObjectKinds(obj)
	if err != nil {
		return schema.GroupVersionKind{}, errors.Wrapf(err, "failed to find gvk for %T, you may need to import the Bhojpur Host generated controller package", obj)
	}

	if len(gvks) == 0 {
		return schema.GroupVersionKind{}, fmt.Errorf("failed to find gvk for %T", obj)
	}

	return gvks[0], nil
}

func Set(objs ...runtime.Object) error {
	for _, obj := range objs {
		if err := setObject(obj); err != nil {
			return err
		}
	}
	return nil
}

func setObject(obj runtime.Object) error {
	gvk := obj.GetObjectKind().GroupVersionKind()
	if gvk.Kind != "" {
		return nil
	}

	gvks, _, err := schemes.All.ObjectKinds(obj)
	if err != nil {
		return err
	}

	if len(gvks) == 0 {
		return nil
	}

	kind := obj.GetObjectKind()
	kind.SetGroupVersionKind(gvks[0])
	return nil
}

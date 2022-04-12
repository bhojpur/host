package crd

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
	"io"
	"os"
	"path/filepath"

	"github.com/bhojpur/host/pkg/common/yaml"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
)

func WriteFile(filename string, crds []CRD) error {
	if err := os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
		return err
	}
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	return Print(f, crds)
}

func Print(out io.Writer, crds []CRD) error {
	obj, err := Objects(crds)
	if err != nil {
		return err
	}

	data, err := yaml.Export(obj...)
	if err != nil {
		return err
	}

	_, err = out.Write(data)
	return err
}

func Objects(crds []CRD) (result []runtime.Object, err error) {
	for _, crdDef := range crds {
		if crdDef.Override == nil {
			crd, err := crdDef.ToCustomResourceDefinition()
			if err != nil {
				return nil, err
			}
			result = append(result, crd)
		} else {
			result = append(result, crdDef.Override)
		}
	}
	return
}

func Create(ctx context.Context, cfg *rest.Config, crds []CRD) error {
	factory, err := NewFactoryFromClient(cfg)
	if err != nil {
		return err
	}

	return factory.BatchCreateCRDs(ctx, crds...).BatchWait()
}

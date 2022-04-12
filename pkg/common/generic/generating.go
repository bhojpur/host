package generic

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
	"github.com/bhojpur/host/pkg/common/apply"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type GeneratingHandlerOptions struct {
	AllowCrossNamespace bool
	AllowClusterScoped  bool
	NoOwnerReference    bool
	DynamicLookup       bool
}

func ConfigureApplyForObject(apply apply.Apply, obj metav1.Object, opts *GeneratingHandlerOptions) apply.Apply {
	if opts == nil {
		opts = &GeneratingHandlerOptions{}
	}

	if opts.DynamicLookup {
		apply = apply.WithDynamicLookup()
	}

	if opts.NoOwnerReference {
		apply = apply.WithSetOwnerReference(true, false)
	}

	if opts.AllowCrossNamespace && !opts.AllowClusterScoped {
		apply = apply.
			WithDefaultNamespace(obj.GetNamespace()).
			WithListerNamespace(obj.GetNamespace())
	}

	if !opts.AllowClusterScoped {
		apply = apply.WithRestrictClusterScoped()
	}

	return apply
}

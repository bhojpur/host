package generators

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
	"strings"

	"k8s.io/code-generator/cmd/client-gen/generators/util"
	"k8s.io/gengo/namer"
	"k8s.io/gengo/types"
)

var (
	Imports = []string{
		"context",
		"time",
		"k8s.io/client-go/rest",
		"github.com/bhojpur/host/pkg/labni/client",
		"github.com/bhojpur/host/pkg/labni/controller",
		"github.com/bhojpur/host/pkg/common/apply",
		"github.com/bhojpur/host/pkg/common/condition",
		"github.com/bhojpur/host/pkg/common/schemes",
		"github.com/bhojpur/host/pkg/common/generic",
		"github.com/bhojpur/host/pkg/common/kv",
		"k8s.io/apimachinery/pkg/api/equality",
		"k8s.io/apimachinery/pkg/api/errors",
		"metav1 \"k8s.io/apimachinery/pkg/apis/meta/v1\"",
		"k8s.io/apimachinery/pkg/labels",
		"k8s.io/apimachinery/pkg/runtime",
		"k8s.io/apimachinery/pkg/runtime/schema",
		"k8s.io/apimachinery/pkg/types",
		"utilruntime \"k8s.io/apimachinery/pkg/util/runtime\"",
		"k8s.io/apimachinery/pkg/watch",
		"k8s.io/client-go/tools/cache",
	}
)

func namespaced(t *types.Type) bool {
	if util.MustParseClientGenTags(t.SecondClosestCommentLines).NonNamespaced {
		return false
	}

	kubeBuilder := false
	for _, line := range t.SecondClosestCommentLines {
		if strings.HasPrefix(line, "+kubebuilder:resource:path=") {
			kubeBuilder = true
			if strings.Contains(line, "scope=Namespaced") {
				return true
			}
		}
	}

	return !kubeBuilder
}

func groupPath(group string) string {
	g := strings.Replace(strings.Split(group, ".")[0], "-", "", -1)
	return groupPackageName(g, "")
}

func groupPackageName(group, groupPackageName string) string {
	if groupPackageName != "" {
		return groupPackageName
	}
	if group == "" {
		return "core"
	}
	return group
}

func upperLowercase(name string) string {
	return namer.IC(strings.ToLower(groupPath(name)))
}

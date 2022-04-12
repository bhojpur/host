package patch

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
	"sync"

	"github.com/bhojpur/host/pkg/common/gvk"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/strategicpatch"
	"k8s.io/client-go/kubernetes/scheme"
)

var (
	patchCache     = map[schema.GroupVersionKind]patchCacheEntry{}
	patchCacheLock = sync.Mutex{}
)

type patchCacheEntry struct {
	patchType types.PatchType
	lookup    strategicpatch.LookupPatchMeta
}

func isJSONPatch(patch []byte) bool {
	// a JSON patch is a list
	return len(patch) > 0 && patch[0] == '['
}

func GetPatchStyle(original, patch []byte) (types.PatchType, strategicpatch.LookupPatchMeta, error) {
	if isJSONPatch(patch) {
		return types.JSONPatchType, nil, nil
	}
	gvk, ok, err := gvk.Detect(original)
	if err != nil {
		return "", nil, err
	}
	if !ok {
		return types.MergePatchType, nil, nil
	}
	return GetMergeStyle(gvk)
}

func GetMergeStyle(gvk schema.GroupVersionKind) (types.PatchType, strategicpatch.LookupPatchMeta, error) {
	var (
		patchType       types.PatchType
		lookupPatchMeta strategicpatch.LookupPatchMeta
	)

	patchCacheLock.Lock()
	entry, ok := patchCache[gvk]
	patchCacheLock.Unlock()

	if ok {
		return entry.patchType, entry.lookup, nil
	}

	versionedObject, err := scheme.Scheme.New(gvk)

	if runtime.IsNotRegisteredError(err) || gvk.Kind == "CustomResourceDefinition" {
		patchType = types.MergePatchType
	} else if err != nil {
		return patchType, nil, err
	} else {
		patchType = types.StrategicMergePatchType
		lookupPatchMeta, err = strategicpatch.NewPatchMetaFromStruct(versionedObject)
		if err != nil {
			return patchType, nil, err
		}
	}

	patchCacheLock.Lock()
	patchCache[gvk] = patchCacheEntry{
		patchType: patchType,
		lookup:    lookupPatchMeta,
	}
	patchCacheLock.Unlock()

	return patchType, lookupPatchMeta, nil
}

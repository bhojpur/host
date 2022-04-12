package trigger

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
	"fmt"
	"sync/atomic"

	"github.com/bhojpur/host/pkg/common/generic"
	"github.com/bhojpur/host/pkg/common/relatedresource"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	counter int64
)

type AllHandler func() error

type Controller interface {
	AddGenericHandler(ctx context.Context, name string, handler generic.Handler)
	GroupVersionKind() schema.GroupVersionKind
	Enqueue(namespace, name string)
}

type Trigger interface {
	Trigger()
	OnTrigger(ctx context.Context, name string, handler AllHandler)
	Key() relatedresource.Key
}

type trigger struct {
	key        string
	controller Controller
}

func New(controller Controller) Trigger {
	return &trigger{
		key:        fmt.Sprintf("__trigger__%d__", atomic.AddInt64(&counter, 1)),
		controller: controller,
	}
}

func (e *trigger) Key() relatedresource.Key {
	return relatedresource.Key{
		Namespace: "__trigger__",
		Name:      e.key,
	}
}

func (e *trigger) Trigger() {
	e.controller.Enqueue("__trigger__", e.key)
}

func (e *trigger) OnTrigger(ctx context.Context, name string, handler AllHandler) {
	e.controller.AddGenericHandler(ctx, name, func(queueKey string, _ runtime.Object) (runtime.Object, error) {
		if queueKey == "__trigger__/"+e.key {
			return nil, handler()
		}
		return nil, nil
	})
	e.Trigger()
}

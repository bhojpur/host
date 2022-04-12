package controller

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
	"errors"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/bhojpur/host/pkg/labni/metrics"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
)

var (
	ErrIgnore = errors.New("ignore handler error")
)

type handlerEntry struct {
	id      int64
	name    string
	handler SharedControllerHandler
}

type SharedHandler struct {
	// keep first because arm32 needs atomic.AddInt64 target to be mem aligned
	idCounter     int64
	controllerGVR string

	lock     sync.RWMutex
	handlers []handlerEntry
}

func (h *SharedHandler) Register(ctx context.Context, name string, handler SharedControllerHandler) {
	h.lock.Lock()
	defer h.lock.Unlock()

	id := atomic.AddInt64(&h.idCounter, 1)
	h.handlers = append(h.handlers, handlerEntry{
		id:      id,
		name:    name,
		handler: handler,
	})

	go func() {
		<-ctx.Done()

		h.lock.Lock()
		defer h.lock.Unlock()

		for i := range h.handlers {
			if h.handlers[i].id == id {
				h.handlers = append(h.handlers[:i], h.handlers[i+1:]...)
				break
			}
		}
	}()
}

func (h *SharedHandler) OnChange(key string, obj runtime.Object) error {
	var (
		errs errorList
	)
	h.lock.RLock()
	handlers := h.handlers
	h.lock.RUnlock()

	for _, handler := range handlers {
		var hasError bool
		reconcileStartTS := time.Now()

		newObj, err := handler.handler.OnChange(key, obj)
		if err != nil && !errors.Is(err, ErrIgnore) {
			errs = append(errs, &handlerError{
				HandlerName: handler.name,
				Err:         err,
			})
			hasError = true
		}
		metrics.IncTotalHandlerExecutions(h.controllerGVR, handler.name, hasError)
		reconcileTime := time.Since(reconcileStartTS)
		metrics.ReportReconcileTime(h.controllerGVR, handler.name, hasError, reconcileTime.Seconds())

		if newObj != nil && !reflect.ValueOf(newObj).IsNil() {
			meta, err := meta.Accessor(newObj)
			if err == nil && meta.GetUID() != "" {
				// avoid using an empty object
				obj = newObj
			} else if err != nil {
				// assign if we can't determine metadata
				obj = newObj
			}
		}
	}

	return errs.ToErr()
}

type errorList []error

func (e errorList) Error() string {
	buf := strings.Builder{}
	for _, err := range e {
		if buf.Len() > 0 {
			buf.WriteString(", ")
		}
		buf.WriteString(err.Error())
	}
	return buf.String()
}

func (e errorList) ToErr() error {
	switch len(e) {
	case 0:
		return nil
	case 1:
		return e[0]
	default:
		return e
	}
}

func (e errorList) Cause() error {
	if len(e) > 0 {
		return e[0]
	}
	return nil
}

type handlerError struct {
	HandlerName string
	Err         error
}

func (h handlerError) Error() string {
	return fmt.Sprintf("handler %s: %v", h.HandlerName, h.Err)
}

func (h handlerError) Cause() error {
	return h.Err
}

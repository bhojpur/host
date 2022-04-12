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
	"strings"
	"time"

	"github.com/bhojpur/host/pkg/core/metrics"
	"github.com/bhojpur/host/pkg/labni/controller"
	errors2 "github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/cache"
)

type HandlerFunc func(key string, obj interface{}) (interface{}, error)

type GenericController interface {
	Informer() cache.SharedIndexInformer
	AddHandler(ctx context.Context, name string, handler HandlerFunc)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, after time.Duration)
}

type genericController struct {
	controller controller.SharedController
	informer   cache.SharedIndexInformer
	name       string
	namespace  string
}

func NewGenericController(namespace, name string, controller controller.SharedController) GenericController {
	return &genericController{
		controller: controller,
		informer:   controller.Informer(),
		name:       name,
		namespace:  namespace,
	}
}

func (g *genericController) Informer() cache.SharedIndexInformer {
	return g.informer
}

func (g *genericController) Enqueue(namespace, name string) {
	g.controller.Enqueue(namespace, name)
}

func (g *genericController) EnqueueAfter(namespace, name string, after time.Duration) {
	g.controller.EnqueueAfter(namespace, name, after)
}

func (g *genericController) AddHandler(ctx context.Context, name string, handler HandlerFunc) {
	g.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(func(key string, obj runtime.Object) (runtime.Object, error) {
		if !isNamespace(g.namespace, obj) {
			return obj, nil
		}
		logrus.Tracef("%s calling handler %s %s", g.name, name, key)
		metrics.IncTotalHandlerExecution(g.name, name)
		result, err := handler(key, obj)
		runtimeObject, _ := result.(runtime.Object)
		if err != nil && !ignoreError(err, false) {
			metrics.IncTotalHandlerFailure(g.name, name, key)
		}
		if _, ok := err.(*ForgetError); ok {
			logrus.Tracef("%v %v completed with dropped err: %v", g.name, key, err)
			return runtimeObject, controller.ErrIgnore
		}
		return runtimeObject, err
	}))
}

func isNamespace(namespace string, obj runtime.Object) bool {
	if namespace == "" || obj == nil {
		return true
	}
	meta, err := meta.Accessor(obj)
	if err != nil {
		// if you can't figure out the namespace, just let it through
		return true
	}
	return meta.GetNamespace() == namespace
}

func ignoreError(err error, checkString bool) bool {
	err = errors2.Cause(err)
	if errors.IsConflict(err) {
		return true
	}
	if err == controller.ErrIgnore {
		return true
	}
	if _, ok := err.(*ForgetError); ok {
		return true
	}
	if checkString {
		return strings.HasSuffix(err.Error(), "please apply your changes to the latest version and try again")
	}
	return false
}

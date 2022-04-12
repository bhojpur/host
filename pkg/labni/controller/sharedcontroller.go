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
	"sync"
	"time"

	"github.com/bhojpur/host/pkg/labni/cache"
	"github.com/bhojpur/host/pkg/labni/client"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	cachetools "k8s.io/client-go/tools/cache"
)

type SharedControllerHandler interface {
	OnChange(key string, obj runtime.Object) (runtime.Object, error)
}

type SharedController interface {
	Controller

	RegisterHandler(ctx context.Context, name string, handler SharedControllerHandler)
	Client() *client.Client
}

type SharedControllerHandlerFunc func(key string, obj runtime.Object) (runtime.Object, error)

func (s SharedControllerHandlerFunc) OnChange(key string, obj runtime.Object) (runtime.Object, error) {
	return s(key, obj)
}

type sharedController struct {
	// this allows one to create a sharedcontroller but it will not actually be started
	// unless some aspect of the controllers informer is accessed or needed to be used
	deferredController func() (Controller, error)
	sharedCacheFactory cache.SharedCacheFactory
	controller         Controller
	gvk                schema.GroupVersionKind
	handler            *SharedHandler
	startLock          sync.Mutex
	started            bool
	startError         error
	client             *client.Client
}

func (s *sharedController) Enqueue(namespace, name string) {
	s.initController().Enqueue(namespace, name)
}

func (s *sharedController) EnqueueAfter(namespace, name string, delay time.Duration) {
	s.initController().EnqueueAfter(namespace, name, delay)
}

func (s *sharedController) EnqueueKey(key string) {
	s.initController().EnqueueKey(key)
}

func (s *sharedController) Informer() cachetools.SharedIndexInformer {
	return s.initController().Informer()
}

func (s *sharedController) Client() *client.Client {
	return s.client
}

func (s *sharedController) initController() Controller {
	s.startLock.Lock()
	defer s.startLock.Unlock()

	if s.controller != nil {
		return s.controller
	}

	controller, err := s.deferredController()
	if err != nil {
		controller = newErrorController()
	}

	s.startError = err
	s.controller = controller
	return s.controller
}

func (s *sharedController) Start(ctx context.Context, workers int) error {
	s.startLock.Lock()
	defer s.startLock.Unlock()

	if s.startError != nil || s.controller == nil {
		return s.startError
	}

	if s.started {
		return nil
	}

	if err := s.controller.Start(ctx, workers); err != nil {
		return err
	}
	s.started = true

	go func() {
		<-ctx.Done()
		s.startLock.Lock()
		defer s.startLock.Unlock()
		s.started = false
	}()

	return nil
}

func (s *sharedController) RegisterHandler(ctx context.Context, name string, handler SharedControllerHandler) {
	// Ensure that controller is initialized
	c := s.initController()

	getHandlerTransaction(ctx).do(func() {
		s.handler.Register(ctx, name, handler)

		s.startLock.Lock()
		defer s.startLock.Unlock()
		if s.started {
			for _, key := range c.Informer().GetStore().ListKeys() {
				c.EnqueueKey(key)
			}
		}
	})
}

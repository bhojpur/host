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
	"context"
	"sync"
	"time"

	"github.com/bhojpur/host/pkg/labni/log"
	"github.com/sirupsen/logrus"

	"github.com/bhojpur/host/pkg/common/schemes"
	"github.com/bhojpur/host/pkg/labni/cache"
	"github.com/bhojpur/host/pkg/labni/client"
	"github.com/bhojpur/host/pkg/labni/controller"
	"k8s.io/apimachinery/pkg/runtime/schema"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/rest"
)

func init() {
	log.Infof = logrus.Infof
	log.Errorf = logrus.Errorf
}

type Factory struct {
	lock              sync.Mutex
	cacheFactory      cache.SharedCacheFactory
	controllerFactory controller.SharedControllerFactory
	threadiness       map[schema.GroupVersionKind]int
	config            *rest.Config
	opts              FactoryOptions
}

type FactoryOptions struct {
	Namespace               string
	Resync                  time.Duration
	SharedCacheFactory      cache.SharedCacheFactory
	SharedControllerFactory controller.SharedControllerFactory
	HealthCallback          func(bool)
}

func NewFactoryFromConfigWithOptions(config *rest.Config, opts *FactoryOptions) (*Factory, error) {
	if opts == nil {
		opts = &FactoryOptions{}
	}

	f := &Factory{
		config:            config,
		threadiness:       map[schema.GroupVersionKind]int{},
		cacheFactory:      opts.SharedCacheFactory,
		controllerFactory: opts.SharedControllerFactory,
		opts:              *opts,
	}

	if f.cacheFactory == nil && f.controllerFactory != nil {
		f.cacheFactory = f.controllerFactory.SharedCacheFactory()
	}

	return f, nil
}

func (c *Factory) SetThreadiness(gvk schema.GroupVersionKind, threadiness int) {
	c.threadiness[gvk] = threadiness
}

func (c *Factory) ControllerFactory() controller.SharedControllerFactory {
	err := c.setControllerFactoryWithLock()
	utilruntime.Must(err)
	return c.controllerFactory
}

func (c *Factory) setControllerFactoryWithLock() error {
	c.lock.Lock()
	defer c.lock.Unlock()

	if c.controllerFactory != nil {
		return nil
	}

	cacheFactory := c.cacheFactory
	if cacheFactory == nil {
		client, err := client.NewSharedClientFactory(c.config, &client.SharedClientFactoryOptions{
			Scheme: schemes.All,
		})
		if err != nil {
			return err
		}

		cacheFactory = cache.NewSharedCachedFactory(client, &cache.SharedCacheFactoryOptions{
			DefaultNamespace: c.opts.Namespace,
			DefaultResync:    c.opts.Resync,
			HealthCallback:   c.opts.HealthCallback,
		})
	}

	c.cacheFactory = cacheFactory
	c.controllerFactory = controller.NewSharedControllerFactory(cacheFactory, &controller.SharedControllerFactoryOptions{
		KindWorkers: c.threadiness,
	})

	return nil
}

func (c *Factory) Sync(ctx context.Context) error {
	if c.cacheFactory != nil {
		c.cacheFactory.Start(ctx)
		c.cacheFactory.WaitForCacheSync(ctx)
	}
	return nil
}

func (c *Factory) Start(ctx context.Context, defaultThreadiness int) error {
	if err := c.Sync(ctx); err != nil {
		return err
	}

	if c.controllerFactory != nil {
		return c.controllerFactory.Start(ctx, defaultThreadiness)
	}

	return nil
}

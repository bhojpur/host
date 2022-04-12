package cache

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

	"github.com/bhojpur/host/pkg/labni/client"
)

const (
	defaultTimeout = 15 * time.Second
)

type healthcheck struct {
	lock     sync.Mutex
	cf       client.SharedClientFactory
	callback func(bool)
}

func (h *healthcheck) ping(ctx context.Context) bool {
	ctx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()
	return h.cf.IsHealthy(ctx)
}

func (h *healthcheck) start(ctx context.Context, cf client.SharedClientFactory) error {
	first, err := h.initialize(cf)
	if err != nil {
		return err
	}
	if first {
		h.ensureHealthy(ctx)
	}
	return nil
}

func (h *healthcheck) initialize(cf client.SharedClientFactory) (bool, error) {
	h.lock.Lock()
	defer h.lock.Unlock()
	if h.cf != nil {
		return false, nil
	}

	h.cf = cf
	return true, nil
}

func (h *healthcheck) ensureHealthy(ctx context.Context) {
	h.lock.Lock()
	defer h.lock.Unlock()
	h.pingUntilGood(ctx)
}

func (h *healthcheck) report(good bool) {
	if h.callback != nil {
		h.callback(good)
	}
}

func (h *healthcheck) pingUntilGood(ctx context.Context) {
	for {
		if h.ping(ctx) {
			h.report(true)
			return
		}

		h.report(false)
		time.Sleep(defaultTimeout)
	}
}

package restwatch

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
	"time"

	"github.com/bhojpur/host/pkg/common/ratelimit"
	"k8s.io/client-go/rest"
)

type WatchClient interface {
	WatchClient() rest.Interface
}

func UnversionedRESTClientFor(config *rest.Config) (rest.Interface, error) {
	// k8s <= 1.16 would not rate limit when calling UnversionedRESTClientFor(config)
	// this keeps that behavior which seems to be relied on in Bhojpur Host.
	if config.QPS == 0.0 && config.RateLimiter == nil {
		config.RateLimiter = ratelimit.None
	}
	client, err := rest.UnversionedRESTClientFor(config)
	if err != nil {
		return nil, err
	}

	if config.Timeout == 0 {
		return client, err
	}

	newConfig := *config
	newConfig.Timeout = 30 * time.Minute
	watchClient, err := rest.UnversionedRESTClientFor(&newConfig)
	if err != nil {
		return nil, err
	}

	return &clientWithWatch{
		RESTClient:  client,
		watchClient: watchClient,
	}, nil
}

type clientWithWatch struct {
	*rest.RESTClient
	watchClient *rest.RESTClient
}

func (c *clientWithWatch) WatchClient() rest.Interface {
	return c.watchClient
}

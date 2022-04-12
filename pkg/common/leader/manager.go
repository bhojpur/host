package leader

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

	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
)

type Manager struct {
	sync.Mutex
	leaderChan    chan struct{}
	leaderStarted bool
	leaderCTX     context.Context
	namespace     string
	name          string
	k8s           kubernetes.Interface
}

func NewManager(namespace, name string, k8s kubernetes.Interface) *Manager {
	return &Manager{
		leaderChan: make(chan struct{}),
		namespace:  namespace,
		name:       name,
		k8s:        k8s,
	}
}

func (m *Manager) Start(ctx context.Context) {
	m.Lock()
	defer m.Unlock()

	if m.leaderStarted {
		return
	}

	m.leaderStarted = true
	go RunOrDie(ctx, m.namespace, m.name, m.k8s, func(ctx context.Context) {
		m.leaderCTX = ctx
		close(m.leaderChan)
	})
}

// OnLeader this function will be called when leadership is acquired.
func (m *Manager) OnLeader(f func(ctx context.Context) error) {
	go func() {
		<-m.leaderChan
		for {
			if err := f(m.leaderCTX); err != nil {
				logrus.Errorf("failed to call leader func: %v", err)
				time.Sleep(5 * time.Second)
				continue
			}
			break
		}
	}()
}

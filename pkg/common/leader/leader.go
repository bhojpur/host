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
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/leaderelection"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
)

type Callback func(cb context.Context)

func RunOrDie(ctx context.Context, namespace, name string, client kubernetes.Interface, cb Callback) {
	if namespace == "" {
		namespace = "kube-system"
	}

	err := run(ctx, namespace, name, client, cb)
	if err != nil {
		logrus.Fatalf("Failed to start leader election for %s", name)
	}
	panic("Failed to start leader election for " + name)
}

func run(ctx context.Context, namespace, name string, client kubernetes.Interface, cb Callback) error {
	id, err := os.Hostname()
	if err != nil {
		return err
	}

	rl, err := resourcelock.New(resourcelock.ConfigMapsResourceLock,
		namespace,
		name,
		client.CoreV1(),
		client.CoordinationV1(),
		resourcelock.ResourceLockConfig{
			Identity: id,
		})
	if err != nil {
		logrus.Fatalf("error creating leader lock for %s: %v", name, err)
	}

	t := time.Second
	if dl := os.Getenv("BHOJPUR_DEV_MODE"); dl != "" {
		t = time.Hour
	}

	leaderelection.RunOrDie(ctx, leaderelection.LeaderElectionConfig{
		Lock:          rl,
		LeaseDuration: 45 * t,
		RenewDeadline: 30 * t,
		RetryPeriod:   2 * t,
		Callbacks: leaderelection.LeaderCallbacks{
			OnStartedLeading: func(ctx context.Context) {
				go cb(ctx)
			},
			OnStoppedLeading: func() {
				logrus.Fatalf("leaderelection lost for %s", name)
			},
		},
		ReleaseOnCancel: true,
	})
	panic("unreachable")
}

package client

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

	"github.com/bhojpur/host/pkg/common/summary"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/dynamic"
)

type summaryClient struct {
	client dynamic.Interface
}

var _ Interface = &summaryClient{}

func NewForDynamicClient(client dynamic.Interface) Interface {
	return &summaryClient{client: client}
}

type summaryResourceClient struct {
	client    dynamic.Interface
	namespace string
	resource  schema.GroupVersionResource
}

func (c *summaryClient) Resource(resource schema.GroupVersionResource) NamespaceableResourceInterface {
	return &summaryResourceClient{client: c.client, resource: resource}
}

func (c *summaryResourceClient) Namespace(ns string) ResourceInterface {
	ret := *c
	ret.namespace = ns
	return &ret
}

func (c *summaryResourceClient) List(ctx context.Context, opts metav1.ListOptions) (*summary.SummarizedObjectList, error) {
	var (
		u   *unstructured.UnstructuredList
		err error
	)

	if c.namespace == "" {
		u, err = c.client.Resource(c.resource).List(ctx, opts)
	} else {
		u, err = c.client.Resource(c.resource).Namespace(c.namespace).List(ctx, opts)
	}
	if err != nil {
		return nil, err
	}

	list := &summary.SummarizedObjectList{
		TypeMeta: metav1.TypeMeta{
			Kind:       u.GetKind(),
			APIVersion: u.GetAPIVersion(),
		},
		ListMeta: metav1.ListMeta{
			ResourceVersion:    u.GetResourceVersion(),
			Continue:           u.GetContinue(),
			RemainingItemCount: u.GetRemainingItemCount(),
		},
	}

	for _, obj := range u.Items {
		list.Items = append(list.Items, *summary.Summarized(&obj))
	}

	return list, nil
}

func (c *summaryResourceClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var (
		resp watch.Interface
		err  error
	)

	eventChan := make(chan watch.Event)

	if c.namespace == "" {
		resp, err = c.client.Resource(c.resource).Watch(ctx, opts)
	} else {
		resp, err = c.client.Resource(c.resource).Namespace(c.namespace).Watch(ctx, opts)
	}
	if err != nil {
		return nil, err
	}

	go func() {
		defer close(eventChan)
		for event := range resp.ResultChan() {
			// don't encode status objects
			if _, ok := event.Object.(*metav1.Status); !ok {
				event.Object = summary.Summarized(event.Object)
			}
			eventChan <- event
		}
	}()

	return &watcher{
		Interface: resp,
		eventChan: eventChan,
	}, nil
}

type watcher struct {
	watch.Interface
	eventChan chan watch.Event
}

func (w watcher) ResultChan() <-chan watch.Event {
	return w.eventChan
}

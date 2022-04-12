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
)

type hTransactionKey struct{}

type HandlerTransaction struct {
	context.Context

	lock   sync.Mutex
	parent context.Context
	todo   []func()
}

func (h *HandlerTransaction) do(f func()) {
	if h == nil {
		f()
		return
	}

	h.lock.Lock()
	defer h.lock.Unlock()

	h.todo = append(h.todo, f)
}

func (h *HandlerTransaction) Commit() {
	h.lock.Lock()
	fs := h.todo
	h.todo = nil
	h.lock.Unlock()

	for _, f := range fs {
		f()
	}
}

func (h *HandlerTransaction) Rollback() {
	h.lock.Lock()
	h.todo = nil
	h.lock.Unlock()
}

func NewHandlerTransaction(ctx context.Context) *HandlerTransaction {
	ht := &HandlerTransaction{
		parent: ctx,
	}
	ctx = context.WithValue(ctx, hTransactionKey{}, ht)
	ht.Context = ctx
	return ht
}

func getHandlerTransaction(ctx context.Context) *HandlerTransaction {
	v, _ := ctx.Value(hTransactionKey{}).(*HandlerTransaction)
	return v
}

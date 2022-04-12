package start

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

	"golang.org/x/sync/errgroup"
)

type Starter interface {
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

func All(ctx context.Context, threadiness int, starters ...Starter) error {
	if err := Sync(ctx, starters...); err != nil {
		return err
	}
	return Start(ctx, threadiness, starters...)
}

func Sync(ctx context.Context, starters ...Starter) error {
	eg, _ := errgroup.WithContext(ctx)
	for _, starter := range starters {
		func(starter Starter) {
			eg.Go(func() error {
				return starter.Sync(ctx)
			})
		}(starter)
	}
	return eg.Wait()
}

func Start(ctx context.Context, threadiness int, starters ...Starter) error {
	for _, starter := range starters {
		if err := starter.Start(ctx, threadiness); err != nil {
			return err
		}
	}
	return nil
}

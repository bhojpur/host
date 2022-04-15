//go:build linux || freebsd
// +build linux freebsd

package dump

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
	"bytes"
	"io"
	"os"
	"os/signal"
	"runtime"

	"github.com/maruel/panicparse/stack"
)

func GoroutineDumpOn(signals ...os.Signal) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, signals...)
	go func() {
		for range c {
			var (
				buf       []byte
				stackSize int
			)
			bufferLen := 16384
			for stackSize == len(buf) {
				buf = make([]byte, bufferLen)
				stackSize = runtime.Stack(buf, true)
				bufferLen *= 2
			}
			buf = buf[:stackSize]
			src := bytes.NewBuffer(buf)
			if goroutines, err := stack.ParseDump(src, os.Stderr); err == nil {
				buckets := stack.SortBuckets(stack.Bucketize(goroutines, stack.AnyValue))
				srcLen, pkgLen := stack.CalcLengths(buckets, true)
				p := &stack.Palette{}
				for _, bucket := range buckets {
					_, _ = io.WriteString(os.Stderr, p.BucketHeader(&bucket, true, len(buckets) > 1))
					_, _ = io.WriteString(os.Stderr, p.StackLines(&bucket.Signature, srcLen, pkgLen, true))
				}
				io.Copy(os.Stderr, bytes.NewBuffer(buf))
			} else {
				io.Copy(os.Stderr, src)
			}
		}
	}()
}

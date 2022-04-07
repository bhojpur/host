package log

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

	"github.com/sirupsen/logrus"
)

type logKey string

const (
	key logKey = "host-logger"
)

type logger interface {
	Debugf(msg string, args ...interface{})
	Infof(msg string, args ...interface{})
	Warnf(msg string, args ...interface{})
}

func SetLogger(ctx context.Context, logger logger) context.Context {
	return context.WithValue(ctx, key, logger)
}

func getLogger(ctx context.Context) logger {
	logger, _ := ctx.Value(key).(logger)
	if logger == nil {
		return logrus.StandardLogger()
	}
	return logger
}

func Infof(ctx context.Context, msg string, args ...interface{}) {
	getLogger(ctx).Infof(msg, args...)
}

func Warnf(ctx context.Context, msg string, args ...interface{}) {
	getLogger(ctx).Warnf(msg, args...)
}

func Debugf(ctx context.Context, msg string, args ...interface{}) {
	getLogger(ctx).Debugf(msg, args...)
}

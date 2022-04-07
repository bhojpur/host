package stream

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
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/sirupsen/logrus"
)

var (
	logs          = map[string]LoggerStream{}
	lock          = sync.Mutex{}
	counter int64 = 1
)

type LogEvent struct {
	Error   bool
	Message string
}

type Logger interface {
	Infof(msg string, args ...interface{})
	Warnf(msg string, args ...interface{})
	Debugf(msg string, args ...interface{})
}

type LoggerStream interface {
	Logger
	ID() string
	Stream() <-chan LogEvent
	Close()
}

func GetLogStream(id string) LoggerStream {
	lock.Lock()
	defer lock.Unlock()
	return logs[id]
}

func NewLogStream() LoggerStream {
	id := atomic.AddInt64(&counter, 1)
	ls := newLoggerStream(strconv.FormatInt(id, 10))

	lock.Lock()
	logs[ls.ID()] = ls
	lock.Unlock()

	return ls
}

type loggerStream struct {
	sync.Mutex
	closed bool
	id     string
	c      chan LogEvent
}

func newLoggerStream(id string) LoggerStream {
	return &loggerStream{
		id: id,
		c:  make(chan LogEvent, 100),
	}
}

func (l *loggerStream) Infof(msg string, args ...interface{}) {
	l.write(false, msg, args...)
}

func (l *loggerStream) Warnf(msg string, args ...interface{}) {
	l.write(true, msg, args...)
}

func (l *loggerStream) Debugf(msg string, args ...interface{}) {
	logrus.Debugf(msg, args...)
}

func (l *loggerStream) write(error bool, msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args...)
	l.Lock()
	if !l.closed {
		l.c <- LogEvent{
			Error:   error,
			Message: msg,
		}
	}
	l.Unlock()
}

func (l *loggerStream) ID() string {
	return l.id
}

func (l *loggerStream) Stream() <-chan LogEvent {
	return l.c
}

func (l *loggerStream) Close() {
	l.Lock()
	if !l.closed {
		close(l.c)
		l.closed = true
	}
	l.Unlock()
	lock.Lock()
	delete(logs, l.id)
	lock.Unlock()
}

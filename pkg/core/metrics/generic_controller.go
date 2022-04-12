package metrics

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
	"github.com/prometheus/client_golang/prometheus"
)

var (
	TotalHandlerExecution = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: "bhojpur_generic_controller",
			Name:      "total_handler_execution",
			Help:      "Total count of hanlder executions",
		},
		[]string{"name", "handlerName"},
	)

	TotalHandlerFailure = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: "bhojpur_generic_controller",
			Name:      "total_handler_failure",
			Help:      "Total count of handler failures",
		},
		[]string{"name", "handlerName", "key"},
	)
)

func IncTotalHandlerExecution(controllerName, handlerName string) {
	if prometheusMetrics {
		TotalHandlerExecution.With(
			prometheus.Labels{
				"name":        controllerName,
				"handlerName": handlerName},
		).Inc()
	}
}

func IncTotalHandlerFailure(controllerName, handlerName, key string) {
	if prometheusMetrics {
		TotalHandlerFailure.With(
			prometheus.Labels{
				"name":        controllerName,
				"handlerName": handlerName,
				"key":         key,
			},
		).Inc()
	}
}

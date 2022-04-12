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
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

var prometheusMetrics = false

const (
	labniSubsystem      = "labni_controller"
	controllerNameLabel = "controller_name"
	handlerNameLabel    = "handler_name"
	hasErrorLabel       = "has_error"

	groupLabel   = "group"
	versionLabel = "version"
	kindLabel    = "kind"
)

var (
	// https://prometheus.io/docs/practices/instrumentation/#use-labels explains logic of having 1 total_requests
	// counter with code label vs a counter for each code

	TotalControllerExecutions = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Subsystem: labniSubsystem,
			Name:      "total_handler_execution",
			Help:      "Total count of handler executions",
		},
		[]string{controllerNameLabel, handlerNameLabel, hasErrorLabel},
	)
	TotalCachedObjects = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: labniSubsystem,
			Name:      "total_cached_object",
			Help:      "Total count of cached objects",
		},
		[]string{groupLabel, versionLabel, kindLabel},
	)

	// reconcileTime is a prometheus histogram metric exposes the duration of reconciliations per controller.
	// controller label refers to the controller name
	reconcileTime = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Subsystem: labniSubsystem,
		Name:      "reconcile_time_seconds",
		Help:      "Histogram of the durations per reconciliation per controller",
	}, []string{controllerNameLabel, handlerNameLabel, hasErrorLabel})
)

func IncTotalHandlerExecutions(controllerName, handlerName string, hasError bool) {
	if prometheusMetrics {
		TotalControllerExecutions.With(
			prometheus.Labels{
				controllerNameLabel: controllerName,
				handlerNameLabel:    handlerName,
				hasErrorLabel:       strconv.FormatBool(hasError),
			},
		).Inc()
	}
}

func IncTotalCachedObjects(group, version, kind string, val float64) {
	if prometheusMetrics {
		TotalCachedObjects.With(
			prometheus.Labels{
				groupLabel:   group,
				versionLabel: version,
				kindLabel:    kind,
			},
		).Set(val)
	}
}

func ReportReconcileTime(controllerName, handlerName string, hasError bool, observeTime float64) {
	if prometheusMetrics {
		reconcileTime.With(
			prometheus.Labels{
				controllerNameLabel: controllerName,
				handlerNameLabel:    handlerName,
				hasErrorLabel:       strconv.FormatBool(hasError),
			},
		).Observe(observeTime)
	}
}

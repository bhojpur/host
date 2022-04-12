package upgrade

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

import "path"

const (
	// AnnotationTTLSecondsAfterFinished is used to store a fallback value for job.spec.ttlSecondsAfterFinished
	AnnotationTTLSecondsAfterFinished = GroupName + `/ttl-seconds-after-finished`

	// AnnotationIncludeInDigest is used to determine parts of the plan to include in the hash for upgrading
	// The value should be a comma-delimited string corresponding to the sections of the plan.
	// For example, a value of "spec.concurrency,spec.upgrade.envs" will include
	// spec.concurrency and spec.upgrade.envs from the plan in the hash to track for upgrades.
	AnnotationIncludeInDigest = GroupName + `/digest`

	// LabelController is the name of the upgrade controller.
	LabelController = GroupName + `/controller`

	// LabelNode is the node being upgraded.
	LabelNode = GroupName + `/node`

	// LabelPlan is the plan being applied.
	LabelPlan = GroupName + `/plan`

	// LabelVersion is the version of the plan being applied.
	LabelVersion = GroupName + `/version`

	// LabelPlanSuffix is used for composing labels specific to a plan.
	LabelPlanSuffix = `plan.` + GroupName
)

func LabelPlanName(name string) string {
	return path.Join(LabelPlanSuffix, name)
}
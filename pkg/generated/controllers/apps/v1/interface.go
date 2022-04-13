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

// Code generated by main. DO NOT EDIT.

package v1

import (
	"github.com/bhojpur/host/pkg/common/schemes"
	"github.com/bhojpur/host/pkg/labni/controller"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func init() {
	schemes.Register(v1.AddToScheme)
}

type Interface interface {
	DaemonSet() DaemonSetController
	Deployment() DeploymentController
	StatefulSet() StatefulSetController
}

func New(controllerFactory controller.SharedControllerFactory) Interface {
	return &version{
		controllerFactory: controllerFactory,
	}
}

type version struct {
	controllerFactory controller.SharedControllerFactory
}

func (c *version) DaemonSet() DaemonSetController {
	return NewDaemonSetController(schema.GroupVersionKind{Group: "apps", Version: "v1", Kind: "DaemonSet"}, "daemonsets", true, c.controllerFactory)
}
func (c *version) Deployment() DeploymentController {
	return NewDeploymentController(schema.GroupVersionKind{Group: "apps", Version: "v1", Kind: "Deployment"}, "deployments", true, c.controllerFactory)
}
func (c *version) StatefulSet() StatefulSetController {
	return NewStatefulSetController(schema.GroupVersionKind{Group: "apps", Version: "v1", Kind: "StatefulSet"}, "statefulsets", true, c.controllerFactory)
}

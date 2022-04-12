package generator

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

var k8sClientTemplate = `package {{.version.Version}}

import (
	"github.com/bhojpur/host/pkg/labni/client"
	"github.com/bhojpur/host/pkg/labni/controller"
	"github.com/bhojpur/host/pkg/core/objectclient"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	{{.importPackage}}
)

type Interface interface {
	{{range .schemas}}
	{{.CodeNamePlural}}Getter{{end}}
}

type Client struct {
	controllerFactory controller.SharedControllerFactory
	clientFactory     client.SharedClientFactory
}

func NewForConfig(cfg rest.Config) (Interface, error) {
	scheme := runtime.NewScheme()
	if err := {{.prefix}}AddToScheme(scheme); err != nil {
		return nil, err
	}
	controllerFactory, err := controller.NewSharedControllerFactoryFromConfig(&cfg, scheme)
	if err != nil {
		return nil, err
	}
	return NewFromControllerFactory(controllerFactory)
}

func NewFromControllerFactory(factory controller.SharedControllerFactory) (Interface, error) {
	return &Client{
		controllerFactory: factory,
		clientFactory:     factory.SharedCacheFactory().SharedClientFactory(),
	}, nil
}

{{range .schemas}}
type {{.CodeNamePlural}}Getter interface {
	{{.CodeNamePlural}}(namespace string) {{.CodeName}}Interface
}

func (c *Client) {{.CodeNamePlural}}(namespace string) {{.CodeName}}Interface {
	sharedClient := c.clientFactory.ForResourceKind({{.CodeName}}GroupVersionResource, {{.CodeName}}GroupVersionKind.Kind, {{ . | namespaced }})
	objectClient := objectclient.NewObjectClient(namespace, sharedClient, &{{.CodeName}}Resource, {{.CodeName}}GroupVersionKind, {{.ID}}Factory{})
	return &{{.ID}}Client{
		ns:           namespace,
		client:       c,
		objectClient: objectClient,
	}
}
{{end}}
`

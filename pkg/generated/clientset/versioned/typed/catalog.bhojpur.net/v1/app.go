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
	"context"
	"time"

	v1 "github.com/bhojpur/host/pkg/apis/catalog.bhojpur.net/v1"
	scheme "github.com/bhojpur/host/pkg/generated/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// AppsGetter has a method to return a AppInterface.
// A group's client should implement this interface.
type AppsGetter interface {
	Apps(namespace string) AppInterface
}

// AppInterface has methods to work with App resources.
type AppInterface interface {
	Create(ctx context.Context, app *v1.App, opts metav1.CreateOptions) (*v1.App, error)
	Update(ctx context.Context, app *v1.App, opts metav1.UpdateOptions) (*v1.App, error)
	UpdateStatus(ctx context.Context, app *v1.App, opts metav1.UpdateOptions) (*v1.App, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.App, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.AppList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.App, err error)
	AppExpansion
}

// apps implements AppInterface
type apps struct {
	client rest.Interface
	ns     string
}

// newApps returns a Apps
func newApps(c *CatalogV1Client, namespace string) *apps {
	return &apps{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the app, and returns the corresponding app object, and an error if there is any.
func (c *apps) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.App, err error) {
	result = &v1.App{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apps").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Apps that match those selectors.
func (c *apps) List(ctx context.Context, opts metav1.ListOptions) (result *v1.AppList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.AppList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("apps").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested apps.
func (c *apps) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("apps").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a app and creates it.  Returns the server's representation of the app, and an error, if there is any.
func (c *apps) Create(ctx context.Context, app *v1.App, opts metav1.CreateOptions) (result *v1.App, err error) {
	result = &v1.App{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("apps").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(app).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a app and updates it. Returns the server's representation of the app, and an error, if there is any.
func (c *apps) Update(ctx context.Context, app *v1.App, opts metav1.UpdateOptions) (result *v1.App, err error) {
	result = &v1.App{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("apps").
		Name(app.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(app).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *apps) UpdateStatus(ctx context.Context, app *v1.App, opts metav1.UpdateOptions) (result *v1.App, err error) {
	result = &v1.App{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("apps").
		Name(app.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(app).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the app and deletes it. Returns an error if one occurs.
func (c *apps) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apps").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *apps) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("apps").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched app.
func (c *apps) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.App, err error) {
	result = &v1.App{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("apps").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
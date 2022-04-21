// Code generated by go-swagger; DO NOT EDIT.

// Copyright (c) 2018 Bhojpur Consulting Private Limited, India. All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Runtime Runtime describes an [OCI compliant](https://github.com/opencontainers/runtime-spec)
// runtime.
//
// The runtime is invoked by the daemon via the `containerd` daemon. OCI
// runtimes act as an interface to the Linux kernel namespaces, cgroups,
// and SELinux.
//
//
// swagger:model Runtime
type Runtime struct {

	// Name and, optional, path, of the OCI executable binary.
	//
	// If the path is omitted, the daemon searches the host's `$PATH` for the
	// binary and uses the first result.
	//
	// Example: /usr/local/bin/my-oci-runtime
	Path string `json:"path,omitempty"`

	// List of command-line arguments to pass to the runtime when invoked.
	//
	// Example: ["--debug","--systemd-cgroup=false"]
	RuntimeArgs []string `json:"runtimeArgs"`
}

// Validate validates this runtime
func (m *Runtime) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this runtime based on context it is used
func (m *Runtime) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Runtime) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Runtime) UnmarshalBinary(b []byte) error {
	var res Runtime
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

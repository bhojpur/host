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
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SwarmInfo Represents generic information about swarm.
//
//
// swagger:model SwarmInfo
type SwarmInfo struct {

	// cluster
	Cluster *ClusterInfo `json:"Cluster,omitempty"`

	// control available
	// Example: true
	ControlAvailable *bool `json:"ControlAvailable,omitempty"`

	// error
	Error string `json:"Error,omitempty"`

	// local node state
	LocalNodeState LocalNodeState `json:"LocalNodeState,omitempty"`

	// Total number of managers in the swarm.
	// Example: 3
	Managers *int64 `json:"Managers,omitempty"`

	// IP address at which this node can be reached by other nodes in the
	// swarm.
	//
	// Example: 10.0.0.46
	NodeAddr string `json:"NodeAddr,omitempty"`

	// Unique identifier of for this node in the swarm.
	// Example: k67qz4598weg5unwwffg6z1m1
	NodeID string `json:"NodeID,omitempty"`

	// Total number of nodes in the swarm.
	// Example: 4
	Nodes *int64 `json:"Nodes,omitempty"`

	// List of ID's and addresses of other managers in the swarm.
	//
	// Example: [{"Addr":"10.0.0.158:2377","NodeID":"71izy0goik036k48jg985xnds"},{"Addr":"10.0.0.159:2377","NodeID":"79y6h1o4gv8n120drcprv5nmc"},{"Addr":"10.0.0.46:2377","NodeID":"k67qz4598weg5unwwffg6z1m1"}]
	RemoteManagers []*PeerNode `json:"RemoteManagers"`
}

// Validate validates this swarm info
func (m *SwarmInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocalNodeState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemoteManagers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SwarmInfo) validateCluster(formats strfmt.Registry) error {
	if swag.IsZero(m.Cluster) { // not required
		return nil
	}

	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Cluster")
			}
			return err
		}
	}

	return nil
}

func (m *SwarmInfo) validateLocalNodeState(formats strfmt.Registry) error {
	if swag.IsZero(m.LocalNodeState) { // not required
		return nil
	}

	if err := m.LocalNodeState.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("LocalNodeState")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("LocalNodeState")
		}
		return err
	}

	return nil
}

func (m *SwarmInfo) validateRemoteManagers(formats strfmt.Registry) error {
	if swag.IsZero(m.RemoteManagers) { // not required
		return nil
	}

	for i := 0; i < len(m.RemoteManagers); i++ {
		if swag.IsZero(m.RemoteManagers[i]) { // not required
			continue
		}

		if m.RemoteManagers[i] != nil {
			if err := m.RemoteManagers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("RemoteManagers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("RemoteManagers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this swarm info based on the context it is used
func (m *SwarmInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocalNodeState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRemoteManagers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SwarmInfo) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.Cluster != nil {
		if err := m.Cluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Cluster")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Cluster")
			}
			return err
		}
	}

	return nil
}

func (m *SwarmInfo) contextValidateLocalNodeState(ctx context.Context, formats strfmt.Registry) error {

	if err := m.LocalNodeState.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("LocalNodeState")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("LocalNodeState")
		}
		return err
	}

	return nil
}

func (m *SwarmInfo) contextValidateRemoteManagers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RemoteManagers); i++ {

		if m.RemoteManagers[i] != nil {
			if err := m.RemoteManagers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("RemoteManagers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("RemoteManagers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SwarmInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SwarmInfo) UnmarshalBinary(b []byte) error {
	var res SwarmInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

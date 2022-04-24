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
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EventMessage SystemEventsResponse
//
// EventMessage represents the information an event contains.
//
//
// swagger:model EventMessage
type EventMessage struct {

	// The type of event
	// Example: create
	Action string `json:"Action,omitempty"`

	// actor
	Actor *EventActor `json:"Actor,omitempty"`

	// The type of object emitting the event
	// Example: container
	// Enum: [builder config container daemon image network node plugin secret service volume]
	Type string `json:"Type,omitempty"`

	// Scope of the event. Bhojpur Host machine events are `local` scope. Cluster (Swarm)
	// events are `swarm` scope.
	//
	// Enum: [local swarm]
	Scope string `json:"scope,omitempty"`

	// Timestamp of event
	// Example: 1629574695
	Time int64 `json:"time,omitempty"`

	// Timestamp of event, with nanosecond accuracy
	// Example: 1629574695515050000
	TimeNano int64 `json:"timeNano,omitempty"`
}

// Validate validates this event message
func (m *EventMessage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EventMessage) validateActor(formats strfmt.Registry) error {
	if swag.IsZero(m.Actor) { // not required
		return nil
	}

	if m.Actor != nil {
		if err := m.Actor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Actor")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Actor")
			}
			return err
		}
	}

	return nil
}

var eventMessageTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["builder","config","container","daemon","image","network","node","plugin","secret","service","volume"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		eventMessageTypeTypePropEnum = append(eventMessageTypeTypePropEnum, v)
	}
}

const (

	// EventMessageTypeBuilder captures enum value "builder"
	EventMessageTypeBuilder string = "builder"

	// EventMessageTypeConfig captures enum value "config"
	EventMessageTypeConfig string = "config"

	// EventMessageTypeContainer captures enum value "container"
	EventMessageTypeContainer string = "container"

	// EventMessageTypeDaemon captures enum value "daemon"
	EventMessageTypeDaemon string = "daemon"

	// EventMessageTypeImage captures enum value "image"
	EventMessageTypeImage string = "image"

	// EventMessageTypeNetwork captures enum value "network"
	EventMessageTypeNetwork string = "network"

	// EventMessageTypeNode captures enum value "node"
	EventMessageTypeNode string = "node"

	// EventMessageTypePlugin captures enum value "plugin"
	EventMessageTypePlugin string = "plugin"

	// EventMessageTypeSecret captures enum value "secret"
	EventMessageTypeSecret string = "secret"

	// EventMessageTypeService captures enum value "service"
	EventMessageTypeService string = "service"

	// EventMessageTypeVolume captures enum value "volume"
	EventMessageTypeVolume string = "volume"
)

// prop value enum
func (m *EventMessage) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, eventMessageTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EventMessage) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("Type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

var eventMessageTypeScopePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["local","swarm"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		eventMessageTypeScopePropEnum = append(eventMessageTypeScopePropEnum, v)
	}
}

const (

	// EventMessageScopeLocal captures enum value "local"
	EventMessageScopeLocal string = "local"

	// EventMessageScopeSwarm captures enum value "swarm"
	EventMessageScopeSwarm string = "swarm"
)

// prop value enum
func (m *EventMessage) validateScopeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, eventMessageTypeScopePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EventMessage) validateScope(formats strfmt.Registry) error {
	if swag.IsZero(m.Scope) { // not required
		return nil
	}

	// value enum
	if err := m.validateScopeEnum("scope", "body", m.Scope); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this event message based on the context it is used
func (m *EventMessage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EventMessage) contextValidateActor(ctx context.Context, formats strfmt.Registry) error {

	if m.Actor != nil {
		if err := m.Actor.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Actor")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Actor")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EventMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventMessage) UnmarshalBinary(b []byte) error {
	var res EventMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
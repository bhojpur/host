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
	"github.com/go-openapi/validate"
)

// TaskState task state
//
// swagger:model TaskState
type TaskState string

func NewTaskState(value TaskState) *TaskState {
	return &value
}

// Pointer returns a pointer to a freshly-allocated TaskState.
func (m TaskState) Pointer() *TaskState {
	return &m
}

const (

	// TaskStateNew captures enum value "new"
	TaskStateNew TaskState = "new"

	// TaskStateAllocated captures enum value "allocated"
	TaskStateAllocated TaskState = "allocated"

	// TaskStatePending captures enum value "pending"
	TaskStatePending TaskState = "pending"

	// TaskStateAssigned captures enum value "assigned"
	TaskStateAssigned TaskState = "assigned"

	// TaskStateAccepted captures enum value "accepted"
	TaskStateAccepted TaskState = "accepted"

	// TaskStatePreparing captures enum value "preparing"
	TaskStatePreparing TaskState = "preparing"

	// TaskStateReady captures enum value "ready"
	TaskStateReady TaskState = "ready"

	// TaskStateStarting captures enum value "starting"
	TaskStateStarting TaskState = "starting"

	// TaskStateRunning captures enum value "running"
	TaskStateRunning TaskState = "running"

	// TaskStateComplete captures enum value "complete"
	TaskStateComplete TaskState = "complete"

	// TaskStateShutdown captures enum value "shutdown"
	TaskStateShutdown TaskState = "shutdown"

	// TaskStateFailed captures enum value "failed"
	TaskStateFailed TaskState = "failed"

	// TaskStateRejected captures enum value "rejected"
	TaskStateRejected TaskState = "rejected"

	// TaskStateRemove captures enum value "remove"
	TaskStateRemove TaskState = "remove"

	// TaskStateOrphaned captures enum value "orphaned"
	TaskStateOrphaned TaskState = "orphaned"
)

// for schema
var taskStateEnum []interface{}

func init() {
	var res []TaskState
	if err := json.Unmarshal([]byte(`["new","allocated","pending","assigned","accepted","preparing","ready","starting","running","complete","shutdown","failed","rejected","remove","orphaned"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		taskStateEnum = append(taskStateEnum, v)
	}
}

func (m TaskState) validateTaskStateEnum(path, location string, value TaskState) error {
	if err := validate.EnumCase(path, location, value, taskStateEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this task state
func (m TaskState) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTaskStateEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this task state based on context it is used
func (m TaskState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
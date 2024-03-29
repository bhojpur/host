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

package container

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/bhojpur/host/pkg/machine/models"
)

// ContainerPruneReader is a Reader for the ContainerPrune structure.
type ContainerPruneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerPruneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContainerPruneOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewContainerPruneInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewContainerPruneOK creates a ContainerPruneOK with default headers values
func NewContainerPruneOK() *ContainerPruneOK {
	return &ContainerPruneOK{}
}

/* ContainerPruneOK describes a response with status code 200, with default header values.

No error
*/
type ContainerPruneOK struct {
	Payload *ContainerPruneOKBody
}

func (o *ContainerPruneOK) Error() string {
	return fmt.Sprintf("[POST /containers/prune][%d] containerPruneOK  %+v", 200, o.Payload)
}
func (o *ContainerPruneOK) GetPayload() *ContainerPruneOKBody {
	return o.Payload
}

func (o *ContainerPruneOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerPruneOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerPruneInternalServerError creates a ContainerPruneInternalServerError with default headers values
func NewContainerPruneInternalServerError() *ContainerPruneInternalServerError {
	return &ContainerPruneInternalServerError{}
}

/* ContainerPruneInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type ContainerPruneInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ContainerPruneInternalServerError) Error() string {
	return fmt.Sprintf("[POST /containers/prune][%d] containerPruneInternalServerError  %+v", 500, o.Payload)
}
func (o *ContainerPruneInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ContainerPruneInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ContainerPruneOKBody ContainerPruneResponse
swagger:model ContainerPruneOKBody
*/
type ContainerPruneOKBody struct {

	// Container IDs that were deleted
	ContainersDeleted []string `json:"ContainersDeleted"`

	// Disk space reclaimed in bytes
	SpaceReclaimed int64 `json:"SpaceReclaimed,omitempty"`
}

// Validate validates this container prune o k body
func (o *ContainerPruneOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container prune o k body based on context it is used
func (o *ContainerPruneOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerPruneOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerPruneOKBody) UnmarshalBinary(b []byte) error {
	var res ContainerPruneOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

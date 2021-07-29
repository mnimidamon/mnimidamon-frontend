// Code generated by go-swagger; DO NOT EDIT.

package computer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"mnimidamonbackend/models"
)

// GetCurrentUserGroupComputersReader is a Reader for the GetCurrentUserGroupComputers structure.
type GetCurrentUserGroupComputersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCurrentUserGroupComputersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCurrentUserGroupComputersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetCurrentUserGroupComputersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCurrentUserGroupComputersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCurrentUserGroupComputersOK creates a GetCurrentUserGroupComputersOK with default headers values
func NewGetCurrentUserGroupComputersOK() *GetCurrentUserGroupComputersOK {
	return &GetCurrentUserGroupComputersOK{}
}

/* GetCurrentUserGroupComputersOK describes a response with status code 200, with default header values.

Array of the denoted group computers.
*/
type GetCurrentUserGroupComputersOK struct {
	Payload []*models.GroupComputer
}

func (o *GetCurrentUserGroupComputersOK) Error() string {
	return fmt.Sprintf("[GET /users/current/groups/{group_id}/computers][%d] getCurrentUserGroupComputersOK  %+v", 200, o.Payload)
}
func (o *GetCurrentUserGroupComputersOK) GetPayload() []*models.GroupComputer {
	return o.Payload
}

func (o *GetCurrentUserGroupComputersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrentUserGroupComputersUnauthorized creates a GetCurrentUserGroupComputersUnauthorized with default headers values
func NewGetCurrentUserGroupComputersUnauthorized() *GetCurrentUserGroupComputersUnauthorized {
	return &GetCurrentUserGroupComputersUnauthorized{}
}

/* GetCurrentUserGroupComputersUnauthorized describes a response with status code 401, with default header values.

Unauthorized.
*/
type GetCurrentUserGroupComputersUnauthorized struct {
	Payload *models.Error
}

func (o *GetCurrentUserGroupComputersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users/current/groups/{group_id}/computers][%d] getCurrentUserGroupComputersUnauthorized  %+v", 401, o.Payload)
}
func (o *GetCurrentUserGroupComputersUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCurrentUserGroupComputersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrentUserGroupComputersInternalServerError creates a GetCurrentUserGroupComputersInternalServerError with default headers values
func NewGetCurrentUserGroupComputersInternalServerError() *GetCurrentUserGroupComputersInternalServerError {
	return &GetCurrentUserGroupComputersInternalServerError{}
}

/* GetCurrentUserGroupComputersInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type GetCurrentUserGroupComputersInternalServerError struct {
	Payload *models.Error
}

func (o *GetCurrentUserGroupComputersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/current/groups/{group_id}/computers][%d] getCurrentUserGroupComputersInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCurrentUserGroupComputersInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCurrentUserGroupComputersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

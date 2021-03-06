// Code generated by go-swagger; DO NOT EDIT.

package group_computer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"mnimidamonbackend/models"
)

// GetGroupComputersOfComputerReader is a Reader for the GetGroupComputersOfComputer structure.
type GetGroupComputersOfComputerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGroupComputersOfComputerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGroupComputersOfComputerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGroupComputersOfComputerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetGroupComputersOfComputerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGroupComputersOfComputerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGroupComputersOfComputerOK creates a GetGroupComputersOfComputerOK with default headers values
func NewGetGroupComputersOfComputerOK() *GetGroupComputersOfComputerOK {
	return &GetGroupComputersOfComputerOK{}
}

/* GetGroupComputersOfComputerOK describes a response with status code 200, with default header values.

The computer
*/
type GetGroupComputersOfComputerOK struct {
	Payload []*models.GroupComputer
}

func (o *GetGroupComputersOfComputerOK) Error() string {
	return fmt.Sprintf("[GET /users/current/computers/{computer_id}/groups][%d] getGroupComputersOfComputerOK  %+v", 200, o.Payload)
}
func (o *GetGroupComputersOfComputerOK) GetPayload() []*models.GroupComputer {
	return o.Payload
}

func (o *GetGroupComputersOfComputerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupComputersOfComputerBadRequest creates a GetGroupComputersOfComputerBadRequest with default headers values
func NewGetGroupComputersOfComputerBadRequest() *GetGroupComputersOfComputerBadRequest {
	return &GetGroupComputersOfComputerBadRequest{}
}

/* GetGroupComputersOfComputerBadRequest describes a response with status code 400, with default header values.

Supplied parameters were not okay.
*/
type GetGroupComputersOfComputerBadRequest struct {
	Payload *models.Error
}

func (o *GetGroupComputersOfComputerBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/current/computers/{computer_id}/groups][%d] getGroupComputersOfComputerBadRequest  %+v", 400, o.Payload)
}
func (o *GetGroupComputersOfComputerBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGroupComputersOfComputerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupComputersOfComputerUnauthorized creates a GetGroupComputersOfComputerUnauthorized with default headers values
func NewGetGroupComputersOfComputerUnauthorized() *GetGroupComputersOfComputerUnauthorized {
	return &GetGroupComputersOfComputerUnauthorized{}
}

/* GetGroupComputersOfComputerUnauthorized describes a response with status code 401, with default header values.

Unauthorized.
*/
type GetGroupComputersOfComputerUnauthorized struct {
	Payload *models.Error
}

func (o *GetGroupComputersOfComputerUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users/current/computers/{computer_id}/groups][%d] getGroupComputersOfComputerUnauthorized  %+v", 401, o.Payload)
}
func (o *GetGroupComputersOfComputerUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGroupComputersOfComputerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupComputersOfComputerInternalServerError creates a GetGroupComputersOfComputerInternalServerError with default headers values
func NewGetGroupComputersOfComputerInternalServerError() *GetGroupComputersOfComputerInternalServerError {
	return &GetGroupComputersOfComputerInternalServerError{}
}

/* GetGroupComputersOfComputerInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type GetGroupComputersOfComputerInternalServerError struct {
	Payload *models.Error
}

func (o *GetGroupComputersOfComputerInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/current/computers/{computer_id}/groups][%d] getGroupComputersOfComputerInternalServerError  %+v", 500, o.Payload)
}
func (o *GetGroupComputersOfComputerInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGroupComputersOfComputerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

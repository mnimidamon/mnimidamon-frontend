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

// GetCurrentUserComputerReader is a Reader for the GetCurrentUserComputer structure.
type GetCurrentUserComputerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCurrentUserComputerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCurrentUserComputerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetCurrentUserComputerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetCurrentUserComputerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCurrentUserComputerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCurrentUserComputerOK creates a GetCurrentUserComputerOK with default headers values
func NewGetCurrentUserComputerOK() *GetCurrentUserComputerOK {
	return &GetCurrentUserComputerOK{}
}

/* GetCurrentUserComputerOK describes a response with status code 200, with default header values.

The computer
*/
type GetCurrentUserComputerOK struct {
	Payload *models.Computer
}

func (o *GetCurrentUserComputerOK) Error() string {
	return fmt.Sprintf("[GET /users/current/computers/{computer_id}][%d] getCurrentUserComputerOK  %+v", 200, o.Payload)
}
func (o *GetCurrentUserComputerOK) GetPayload() *models.Computer {
	return o.Payload
}

func (o *GetCurrentUserComputerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Computer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrentUserComputerBadRequest creates a GetCurrentUserComputerBadRequest with default headers values
func NewGetCurrentUserComputerBadRequest() *GetCurrentUserComputerBadRequest {
	return &GetCurrentUserComputerBadRequest{}
}

/* GetCurrentUserComputerBadRequest describes a response with status code 400, with default header values.

Supplied parameters were not okay.
*/
type GetCurrentUserComputerBadRequest struct {
	Payload *models.Error
}

func (o *GetCurrentUserComputerBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/current/computers/{computer_id}][%d] getCurrentUserComputerBadRequest  %+v", 400, o.Payload)
}
func (o *GetCurrentUserComputerBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCurrentUserComputerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrentUserComputerUnauthorized creates a GetCurrentUserComputerUnauthorized with default headers values
func NewGetCurrentUserComputerUnauthorized() *GetCurrentUserComputerUnauthorized {
	return &GetCurrentUserComputerUnauthorized{}
}

/* GetCurrentUserComputerUnauthorized describes a response with status code 401, with default header values.

Unauthorized.
*/
type GetCurrentUserComputerUnauthorized struct {
	Payload *models.Error
}

func (o *GetCurrentUserComputerUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users/current/computers/{computer_id}][%d] getCurrentUserComputerUnauthorized  %+v", 401, o.Payload)
}
func (o *GetCurrentUserComputerUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCurrentUserComputerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrentUserComputerInternalServerError creates a GetCurrentUserComputerInternalServerError with default headers values
func NewGetCurrentUserComputerInternalServerError() *GetCurrentUserComputerInternalServerError {
	return &GetCurrentUserComputerInternalServerError{}
}

/* GetCurrentUserComputerInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type GetCurrentUserComputerInternalServerError struct {
	Payload *models.Error
}

func (o *GetCurrentUserComputerInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/current/computers/{computer_id}][%d] getCurrentUserComputerInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCurrentUserComputerInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCurrentUserComputerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
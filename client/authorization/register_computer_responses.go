// Code generated by go-swagger; DO NOT EDIT.

package authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"mnimidamonbackend/models"
)

// RegisterComputerReader is a Reader for the RegisterComputer structure.
type RegisterComputerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RegisterComputerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRegisterComputerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRegisterComputerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRegisterComputerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRegisterComputerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRegisterComputerOK creates a RegisterComputerOK with default headers values
func NewRegisterComputerOK() *RegisterComputerOK {
	return &RegisterComputerOK{}
}

/* RegisterComputerOK describes a response with status code 200, with default header values.

Computer successfuly created. Returned Computer object and API key that is used in X-COMP-KEY header.
*/
type RegisterComputerOK struct {
	Payload *models.CreateComputerResponse
}

func (o *RegisterComputerOK) Error() string {
	return fmt.Sprintf("[POST /users/current/computers][%d] registerComputerOK  %+v", 200, o.Payload)
}
func (o *RegisterComputerOK) GetPayload() *models.CreateComputerResponse {
	return o.Payload
}

func (o *RegisterComputerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreateComputerResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegisterComputerBadRequest creates a RegisterComputerBadRequest with default headers values
func NewRegisterComputerBadRequest() *RegisterComputerBadRequest {
	return &RegisterComputerBadRequest{}
}

/* RegisterComputerBadRequest describes a response with status code 400, with default header values.

Supplied parameters were not okay.
*/
type RegisterComputerBadRequest struct {
	Payload *models.Error
}

func (o *RegisterComputerBadRequest) Error() string {
	return fmt.Sprintf("[POST /users/current/computers][%d] registerComputerBadRequest  %+v", 400, o.Payload)
}
func (o *RegisterComputerBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *RegisterComputerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegisterComputerUnauthorized creates a RegisterComputerUnauthorized with default headers values
func NewRegisterComputerUnauthorized() *RegisterComputerUnauthorized {
	return &RegisterComputerUnauthorized{}
}

/* RegisterComputerUnauthorized describes a response with status code 401, with default header values.

Unauthorized.
*/
type RegisterComputerUnauthorized struct {
	Payload *models.Error
}

func (o *RegisterComputerUnauthorized) Error() string {
	return fmt.Sprintf("[POST /users/current/computers][%d] registerComputerUnauthorized  %+v", 401, o.Payload)
}
func (o *RegisterComputerUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *RegisterComputerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRegisterComputerInternalServerError creates a RegisterComputerInternalServerError with default headers values
func NewRegisterComputerInternalServerError() *RegisterComputerInternalServerError {
	return &RegisterComputerInternalServerError{}
}

/* RegisterComputerInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type RegisterComputerInternalServerError struct {
	Payload *models.Error
}

func (o *RegisterComputerInternalServerError) Error() string {
	return fmt.Sprintf("[POST /users/current/computers][%d] registerComputerInternalServerError  %+v", 500, o.Payload)
}
func (o *RegisterComputerInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *RegisterComputerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// JoinComputerToGroupReader is a Reader for the JoinComputerToGroup structure.
type JoinComputerToGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *JoinComputerToGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewJoinComputerToGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewJoinComputerToGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewJoinComputerToGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewJoinComputerToGroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewJoinComputerToGroupOK creates a JoinComputerToGroupOK with default headers values
func NewJoinComputerToGroupOK() *JoinComputerToGroupOK {
	return &JoinComputerToGroupOK{}
}

/* JoinComputerToGroupOK describes a response with status code 200, with default header values.

The group computer
*/
type JoinComputerToGroupOK struct {
	Payload *models.GroupComputer
}

func (o *JoinComputerToGroupOK) Error() string {
	return fmt.Sprintf("[POST /users/current/computers/current/groups/{group_id}/join][%d] joinComputerToGroupOK  %+v", 200, o.Payload)
}
func (o *JoinComputerToGroupOK) GetPayload() *models.GroupComputer {
	return o.Payload
}

func (o *JoinComputerToGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GroupComputer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewJoinComputerToGroupBadRequest creates a JoinComputerToGroupBadRequest with default headers values
func NewJoinComputerToGroupBadRequest() *JoinComputerToGroupBadRequest {
	return &JoinComputerToGroupBadRequest{}
}

/* JoinComputerToGroupBadRequest describes a response with status code 400, with default header values.

Supplied parameters were not okay.
*/
type JoinComputerToGroupBadRequest struct {
	Payload *models.Error
}

func (o *JoinComputerToGroupBadRequest) Error() string {
	return fmt.Sprintf("[POST /users/current/computers/current/groups/{group_id}/join][%d] joinComputerToGroupBadRequest  %+v", 400, o.Payload)
}
func (o *JoinComputerToGroupBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *JoinComputerToGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewJoinComputerToGroupUnauthorized creates a JoinComputerToGroupUnauthorized with default headers values
func NewJoinComputerToGroupUnauthorized() *JoinComputerToGroupUnauthorized {
	return &JoinComputerToGroupUnauthorized{}
}

/* JoinComputerToGroupUnauthorized describes a response with status code 401, with default header values.

Unauthorized.
*/
type JoinComputerToGroupUnauthorized struct {
	Payload *models.Error
}

func (o *JoinComputerToGroupUnauthorized) Error() string {
	return fmt.Sprintf("[POST /users/current/computers/current/groups/{group_id}/join][%d] joinComputerToGroupUnauthorized  %+v", 401, o.Payload)
}
func (o *JoinComputerToGroupUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *JoinComputerToGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewJoinComputerToGroupInternalServerError creates a JoinComputerToGroupInternalServerError with default headers values
func NewJoinComputerToGroupInternalServerError() *JoinComputerToGroupInternalServerError {
	return &JoinComputerToGroupInternalServerError{}
}

/* JoinComputerToGroupInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type JoinComputerToGroupInternalServerError struct {
	Payload *models.Error
}

func (o *JoinComputerToGroupInternalServerError) Error() string {
	return fmt.Sprintf("[POST /users/current/computers/current/groups/{group_id}/join][%d] joinComputerToGroupInternalServerError  %+v", 500, o.Payload)
}
func (o *JoinComputerToGroupInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *JoinComputerToGroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
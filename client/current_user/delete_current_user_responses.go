// Code generated by go-swagger; DO NOT EDIT.

package current_user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"mnimidamonbackend/models"
)

// DeleteCurrentUserReader is a Reader for the DeleteCurrentUser structure.
type DeleteCurrentUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCurrentUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteCurrentUserNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteCurrentUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteCurrentUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteCurrentUserNoContent creates a DeleteCurrentUserNoContent with default headers values
func NewDeleteCurrentUserNoContent() *DeleteCurrentUserNoContent {
	return &DeleteCurrentUserNoContent{}
}

/* DeleteCurrentUserNoContent describes a response with status code 204, with default header values.

Successfuly deleted current user account.
*/
type DeleteCurrentUserNoContent struct {
}

func (o *DeleteCurrentUserNoContent) Error() string {
	return fmt.Sprintf("[DELETE /users/current][%d] deleteCurrentUserNoContent ", 204)
}

func (o *DeleteCurrentUserNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteCurrentUserUnauthorized creates a DeleteCurrentUserUnauthorized with default headers values
func NewDeleteCurrentUserUnauthorized() *DeleteCurrentUserUnauthorized {
	return &DeleteCurrentUserUnauthorized{}
}

/* DeleteCurrentUserUnauthorized describes a response with status code 401, with default header values.

Unauthorized.
*/
type DeleteCurrentUserUnauthorized struct {
	Payload *models.Error
}

func (o *DeleteCurrentUserUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /users/current][%d] deleteCurrentUserUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteCurrentUserUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteCurrentUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCurrentUserInternalServerError creates a DeleteCurrentUserInternalServerError with default headers values
func NewDeleteCurrentUserInternalServerError() *DeleteCurrentUserInternalServerError {
	return &DeleteCurrentUserInternalServerError{}
}

/* DeleteCurrentUserInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type DeleteCurrentUserInternalServerError struct {
	Payload *models.Error
}

func (o *DeleteCurrentUserInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /users/current][%d] deleteCurrentUserInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteCurrentUserInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeleteCurrentUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

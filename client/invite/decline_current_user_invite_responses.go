// Code generated by go-swagger; DO NOT EDIT.

package invite

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"mnimidamonbackend/models"
)

// DeclineCurrentUserInviteReader is a Reader for the DeclineCurrentUserInvite structure.
type DeclineCurrentUserInviteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeclineCurrentUserInviteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeclineCurrentUserInviteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeclineCurrentUserInviteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeclineCurrentUserInviteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeclineCurrentUserInviteNoContent creates a DeclineCurrentUserInviteNoContent with default headers values
func NewDeclineCurrentUserInviteNoContent() *DeclineCurrentUserInviteNoContent {
	return &DeclineCurrentUserInviteNoContent{}
}

/* DeclineCurrentUserInviteNoContent describes a response with status code 204, with default header values.

Invite declined.
*/
type DeclineCurrentUserInviteNoContent struct {
}

func (o *DeclineCurrentUserInviteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /users/current/invites/{group_id}][%d] declineCurrentUserInviteNoContent ", 204)
}

func (o *DeclineCurrentUserInviteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeclineCurrentUserInviteUnauthorized creates a DeclineCurrentUserInviteUnauthorized with default headers values
func NewDeclineCurrentUserInviteUnauthorized() *DeclineCurrentUserInviteUnauthorized {
	return &DeclineCurrentUserInviteUnauthorized{}
}

/* DeclineCurrentUserInviteUnauthorized describes a response with status code 401, with default header values.

Unauthorized.
*/
type DeclineCurrentUserInviteUnauthorized struct {
	Payload *models.Error
}

func (o *DeclineCurrentUserInviteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /users/current/invites/{group_id}][%d] declineCurrentUserInviteUnauthorized  %+v", 401, o.Payload)
}
func (o *DeclineCurrentUserInviteUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeclineCurrentUserInviteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeclineCurrentUserInviteInternalServerError creates a DeclineCurrentUserInviteInternalServerError with default headers values
func NewDeclineCurrentUserInviteInternalServerError() *DeclineCurrentUserInviteInternalServerError {
	return &DeclineCurrentUserInviteInternalServerError{}
}

/* DeclineCurrentUserInviteInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type DeclineCurrentUserInviteInternalServerError struct {
	Payload *models.Error
}

func (o *DeclineCurrentUserInviteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /users/current/invites/{group_id}][%d] declineCurrentUserInviteInternalServerError  %+v", 500, o.Payload)
}
func (o *DeclineCurrentUserInviteInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *DeclineCurrentUserInviteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

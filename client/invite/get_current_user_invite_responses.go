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

// GetCurrentUserInviteReader is a Reader for the GetCurrentUserInvite structure.
type GetCurrentUserInviteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCurrentUserInviteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCurrentUserInviteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetCurrentUserInviteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetCurrentUserInviteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetCurrentUserInviteOK creates a GetCurrentUserInviteOK with default headers values
func NewGetCurrentUserInviteOK() *GetCurrentUserInviteOK {
	return &GetCurrentUserInviteOK{}
}

/* GetCurrentUserInviteOK describes a response with status code 200, with default header values.

Invite object
*/
type GetCurrentUserInviteOK struct {
	Payload *models.Invite
}

func (o *GetCurrentUserInviteOK) Error() string {
	return fmt.Sprintf("[GET /users/current/invites/{group_id}][%d] getCurrentUserInviteOK  %+v", 200, o.Payload)
}
func (o *GetCurrentUserInviteOK) GetPayload() *models.Invite {
	return o.Payload
}

func (o *GetCurrentUserInviteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Invite)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrentUserInviteUnauthorized creates a GetCurrentUserInviteUnauthorized with default headers values
func NewGetCurrentUserInviteUnauthorized() *GetCurrentUserInviteUnauthorized {
	return &GetCurrentUserInviteUnauthorized{}
}

/* GetCurrentUserInviteUnauthorized describes a response with status code 401, with default header values.

Unauthorized.
*/
type GetCurrentUserInviteUnauthorized struct {
	Payload *models.Error
}

func (o *GetCurrentUserInviteUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users/current/invites/{group_id}][%d] getCurrentUserInviteUnauthorized  %+v", 401, o.Payload)
}
func (o *GetCurrentUserInviteUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCurrentUserInviteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCurrentUserInviteInternalServerError creates a GetCurrentUserInviteInternalServerError with default headers values
func NewGetCurrentUserInviteInternalServerError() *GetCurrentUserInviteInternalServerError {
	return &GetCurrentUserInviteInternalServerError{}
}

/* GetCurrentUserInviteInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type GetCurrentUserInviteInternalServerError struct {
	Payload *models.Error
}

func (o *GetCurrentUserInviteInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/current/invites/{group_id}][%d] getCurrentUserInviteInternalServerError  %+v", 500, o.Payload)
}
func (o *GetCurrentUserInviteInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetCurrentUserInviteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

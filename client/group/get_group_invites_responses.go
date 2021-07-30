// Code generated by go-swagger; DO NOT EDIT.

package group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"mnimidamonbackend/models"
)

// GetGroupInvitesReader is a Reader for the GetGroupInvites structure.
type GetGroupInvitesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGroupInvitesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGroupInvitesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetGroupInvitesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGroupInvitesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGroupInvitesOK creates a GetGroupInvitesOK with default headers values
func NewGetGroupInvitesOK() *GetGroupInvitesOK {
	return &GetGroupInvitesOK{}
}

/* GetGroupInvitesOK describes a response with status code 200, with default header values.

Array of active invites of the group.
*/
type GetGroupInvitesOK struct {
	Payload []*models.Invite
}

func (o *GetGroupInvitesOK) Error() string {
	return fmt.Sprintf("[GET /users/current/groups/{group_id}/invites][%d] getGroupInvitesOK  %+v", 200, o.Payload)
}
func (o *GetGroupInvitesOK) GetPayload() []*models.Invite {
	return o.Payload
}

func (o *GetGroupInvitesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupInvitesUnauthorized creates a GetGroupInvitesUnauthorized with default headers values
func NewGetGroupInvitesUnauthorized() *GetGroupInvitesUnauthorized {
	return &GetGroupInvitesUnauthorized{}
}

/* GetGroupInvitesUnauthorized describes a response with status code 401, with default header values.

Unauthorized.
*/
type GetGroupInvitesUnauthorized struct {
	Payload *models.Error
}

func (o *GetGroupInvitesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users/current/groups/{group_id}/invites][%d] getGroupInvitesUnauthorized  %+v", 401, o.Payload)
}
func (o *GetGroupInvitesUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGroupInvitesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupInvitesInternalServerError creates a GetGroupInvitesInternalServerError with default headers values
func NewGetGroupInvitesInternalServerError() *GetGroupInvitesInternalServerError {
	return &GetGroupInvitesInternalServerError{}
}

/* GetGroupInvitesInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type GetGroupInvitesInternalServerError struct {
	Payload *models.Error
}

func (o *GetGroupInvitesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/current/groups/{group_id}/invites][%d] getGroupInvitesInternalServerError  %+v", 500, o.Payload)
}
func (o *GetGroupInvitesInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGroupInvitesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
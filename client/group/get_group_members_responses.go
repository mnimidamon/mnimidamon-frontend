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

// GetGroupMembersReader is a Reader for the GetGroupMembers structure.
type GetGroupMembersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGroupMembersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGroupMembersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGroupMembersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetGroupMembersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGroupMembersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGroupMembersOK creates a GetGroupMembersOK with default headers values
func NewGetGroupMembersOK() *GetGroupMembersOK {
	return &GetGroupMembersOK{}
}

/* GetGroupMembersOK describes a response with status code 200, with default header values.

Array of group members.
*/
type GetGroupMembersOK struct {
	Payload []*models.User
}

func (o *GetGroupMembersOK) Error() string {
	return fmt.Sprintf("[GET /users/current/groups/{group_id}/members][%d] getGroupMembersOK  %+v", 200, o.Payload)
}
func (o *GetGroupMembersOK) GetPayload() []*models.User {
	return o.Payload
}

func (o *GetGroupMembersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupMembersBadRequest creates a GetGroupMembersBadRequest with default headers values
func NewGetGroupMembersBadRequest() *GetGroupMembersBadRequest {
	return &GetGroupMembersBadRequest{}
}

/* GetGroupMembersBadRequest describes a response with status code 400, with default header values.

Supplied parameters were not okay.
*/
type GetGroupMembersBadRequest struct {
	Payload *models.Error
}

func (o *GetGroupMembersBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/current/groups/{group_id}/members][%d] getGroupMembersBadRequest  %+v", 400, o.Payload)
}
func (o *GetGroupMembersBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGroupMembersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupMembersUnauthorized creates a GetGroupMembersUnauthorized with default headers values
func NewGetGroupMembersUnauthorized() *GetGroupMembersUnauthorized {
	return &GetGroupMembersUnauthorized{}
}

/* GetGroupMembersUnauthorized describes a response with status code 401, with default header values.

Unauthorized.
*/
type GetGroupMembersUnauthorized struct {
	Payload *models.Error
}

func (o *GetGroupMembersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users/current/groups/{group_id}/members][%d] getGroupMembersUnauthorized  %+v", 401, o.Payload)
}
func (o *GetGroupMembersUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGroupMembersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupMembersInternalServerError creates a GetGroupMembersInternalServerError with default headers values
func NewGetGroupMembersInternalServerError() *GetGroupMembersInternalServerError {
	return &GetGroupMembersInternalServerError{}
}

/* GetGroupMembersInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type GetGroupMembersInternalServerError struct {
	Payload *models.Error
}

func (o *GetGroupMembersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/current/groups/{group_id}/members][%d] getGroupMembersInternalServerError  %+v", 500, o.Payload)
}
func (o *GetGroupMembersInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGroupMembersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

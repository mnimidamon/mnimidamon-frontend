// Code generated by go-swagger; DO NOT EDIT.

package backup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"mnimidamonbackend/models"
)

// GetGroupBackupsReader is a Reader for the GetGroupBackups structure.
type GetGroupBackupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGroupBackupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGroupBackupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetGroupBackupsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGroupBackupsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGroupBackupsOK creates a GetGroupBackupsOK with default headers values
func NewGetGroupBackupsOK() *GetGroupBackupsOK {
	return &GetGroupBackupsOK{}
}

/* GetGroupBackupsOK describes a response with status code 200, with default header values.

Array of the group backups.
*/
type GetGroupBackupsOK struct {
	Payload []*models.Backup
}

func (o *GetGroupBackupsOK) Error() string {
	return fmt.Sprintf("[GET /users/current/computers/current/groups/{group_id}/backups][%d] getGroupBackupsOK  %+v", 200, o.Payload)
}
func (o *GetGroupBackupsOK) GetPayload() []*models.Backup {
	return o.Payload
}

func (o *GetGroupBackupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupBackupsUnauthorized creates a GetGroupBackupsUnauthorized with default headers values
func NewGetGroupBackupsUnauthorized() *GetGroupBackupsUnauthorized {
	return &GetGroupBackupsUnauthorized{}
}

/* GetGroupBackupsUnauthorized describes a response with status code 401, with default header values.

Unauthorized.
*/
type GetGroupBackupsUnauthorized struct {
	Payload *models.Error
}

func (o *GetGroupBackupsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users/current/computers/current/groups/{group_id}/backups][%d] getGroupBackupsUnauthorized  %+v", 401, o.Payload)
}
func (o *GetGroupBackupsUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGroupBackupsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupBackupsInternalServerError creates a GetGroupBackupsInternalServerError with default headers values
func NewGetGroupBackupsInternalServerError() *GetGroupBackupsInternalServerError {
	return &GetGroupBackupsInternalServerError{}
}

/* GetGroupBackupsInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type GetGroupBackupsInternalServerError struct {
	Payload *models.Error
}

func (o *GetGroupBackupsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/current/computers/current/groups/{group_id}/backups][%d] getGroupBackupsInternalServerError  %+v", 500, o.Payload)
}
func (o *GetGroupBackupsInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGroupBackupsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// GetGroupBackupReader is a Reader for the GetGroupBackup structure.
type GetGroupBackupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGroupBackupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGroupBackupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetGroupBackupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGroupBackupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGroupBackupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGroupBackupOK creates a GetGroupBackupOK with default headers values
func NewGetGroupBackupOK() *GetGroupBackupOK {
	return &GetGroupBackupOK{}
}

/* GetGroupBackupOK describes a response with status code 200, with default header values.

The specified .
*/
type GetGroupBackupOK struct {
	Payload *models.Backup
}

func (o *GetGroupBackupOK) Error() string {
	return fmt.Sprintf("[GET /users/current/computers/current/groups/{group_id}/backups/{backup_id}][%d] getGroupBackupOK  %+v", 200, o.Payload)
}
func (o *GetGroupBackupOK) GetPayload() *models.Backup {
	return o.Payload
}

func (o *GetGroupBackupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Backup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupBackupUnauthorized creates a GetGroupBackupUnauthorized with default headers values
func NewGetGroupBackupUnauthorized() *GetGroupBackupUnauthorized {
	return &GetGroupBackupUnauthorized{}
}

/* GetGroupBackupUnauthorized describes a response with status code 401, with default header values.

Unauthorized.
*/
type GetGroupBackupUnauthorized struct {
	Payload *models.Error
}

func (o *GetGroupBackupUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users/current/computers/current/groups/{group_id}/backups/{backup_id}][%d] getGroupBackupUnauthorized  %+v", 401, o.Payload)
}
func (o *GetGroupBackupUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGroupBackupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupBackupNotFound creates a GetGroupBackupNotFound with default headers values
func NewGetGroupBackupNotFound() *GetGroupBackupNotFound {
	return &GetGroupBackupNotFound{}
}

/* GetGroupBackupNotFound describes a response with status code 404, with default header values.

The specified resource was not found.
*/
type GetGroupBackupNotFound struct {
	Payload *models.Error
}

func (o *GetGroupBackupNotFound) Error() string {
	return fmt.Sprintf("[GET /users/current/computers/current/groups/{group_id}/backups/{backup_id}][%d] getGroupBackupNotFound  %+v", 404, o.Payload)
}
func (o *GetGroupBackupNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGroupBackupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupBackupInternalServerError creates a GetGroupBackupInternalServerError with default headers values
func NewGetGroupBackupInternalServerError() *GetGroupBackupInternalServerError {
	return &GetGroupBackupInternalServerError{}
}

/* GetGroupBackupInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type GetGroupBackupInternalServerError struct {
	Payload *models.Error
}

func (o *GetGroupBackupInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/current/computers/current/groups/{group_id}/backups/{backup_id}][%d] getGroupBackupInternalServerError  %+v", 500, o.Payload)
}
func (o *GetGroupBackupInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetGroupBackupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

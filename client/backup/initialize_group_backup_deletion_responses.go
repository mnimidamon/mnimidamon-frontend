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

// InitializeGroupBackupDeletionReader is a Reader for the InitializeGroupBackupDeletion structure.
type InitializeGroupBackupDeletionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InitializeGroupBackupDeletionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewInitializeGroupBackupDeletionNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewInitializeGroupBackupDeletionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewInitializeGroupBackupDeletionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewInitializeGroupBackupDeletionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInitializeGroupBackupDeletionNoContent creates a InitializeGroupBackupDeletionNoContent with default headers values
func NewInitializeGroupBackupDeletionNoContent() *InitializeGroupBackupDeletionNoContent {
	return &InitializeGroupBackupDeletionNoContent{}
}

/* InitializeGroupBackupDeletionNoContent describes a response with status code 204, with default header values.

Successuful backup deletion.
*/
type InitializeGroupBackupDeletionNoContent struct {
}

func (o *InitializeGroupBackupDeletionNoContent) Error() string {
	return fmt.Sprintf("[DELETE /users/current/computers/current/groups/{group_id}/backups/{backup_id}][%d] initializeGroupBackupDeletionNoContent ", 204)
}

func (o *InitializeGroupBackupDeletionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInitializeGroupBackupDeletionUnauthorized creates a InitializeGroupBackupDeletionUnauthorized with default headers values
func NewInitializeGroupBackupDeletionUnauthorized() *InitializeGroupBackupDeletionUnauthorized {
	return &InitializeGroupBackupDeletionUnauthorized{}
}

/* InitializeGroupBackupDeletionUnauthorized describes a response with status code 401, with default header values.

Unauthorized.
*/
type InitializeGroupBackupDeletionUnauthorized struct {
	Payload *models.Error
}

func (o *InitializeGroupBackupDeletionUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /users/current/computers/current/groups/{group_id}/backups/{backup_id}][%d] initializeGroupBackupDeletionUnauthorized  %+v", 401, o.Payload)
}
func (o *InitializeGroupBackupDeletionUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *InitializeGroupBackupDeletionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInitializeGroupBackupDeletionNotFound creates a InitializeGroupBackupDeletionNotFound with default headers values
func NewInitializeGroupBackupDeletionNotFound() *InitializeGroupBackupDeletionNotFound {
	return &InitializeGroupBackupDeletionNotFound{}
}

/* InitializeGroupBackupDeletionNotFound describes a response with status code 404, with default header values.

The specified resource was not found.
*/
type InitializeGroupBackupDeletionNotFound struct {
	Payload *models.Error
}

func (o *InitializeGroupBackupDeletionNotFound) Error() string {
	return fmt.Sprintf("[DELETE /users/current/computers/current/groups/{group_id}/backups/{backup_id}][%d] initializeGroupBackupDeletionNotFound  %+v", 404, o.Payload)
}
func (o *InitializeGroupBackupDeletionNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *InitializeGroupBackupDeletionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInitializeGroupBackupDeletionInternalServerError creates a InitializeGroupBackupDeletionInternalServerError with default headers values
func NewInitializeGroupBackupDeletionInternalServerError() *InitializeGroupBackupDeletionInternalServerError {
	return &InitializeGroupBackupDeletionInternalServerError{}
}

/* InitializeGroupBackupDeletionInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type InitializeGroupBackupDeletionInternalServerError struct {
	Payload *models.Error
}

func (o *InitializeGroupBackupDeletionInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /users/current/computers/current/groups/{group_id}/backups/{backup_id}][%d] initializeGroupBackupDeletionInternalServerError  %+v", 500, o.Payload)
}
func (o *InitializeGroupBackupDeletionInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *InitializeGroupBackupDeletionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

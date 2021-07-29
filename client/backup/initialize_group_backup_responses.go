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

// InitializeGroupBackupReader is a Reader for the InitializeGroupBackup structure.
type InitializeGroupBackupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InitializeGroupBackupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInitializeGroupBackupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewInitializeGroupBackupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewInitializeGroupBackupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewInitializeGroupBackupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewInitializeGroupBackupOK creates a InitializeGroupBackupOK with default headers values
func NewInitializeGroupBackupOK() *InitializeGroupBackupOK {
	return &InitializeGroupBackupOK{}
}

/* InitializeGroupBackupOK describes a response with status code 200, with default header values.

Newly created backup object.
*/
type InitializeGroupBackupOK struct {
	Payload *models.Backup
}

func (o *InitializeGroupBackupOK) Error() string {
	return fmt.Sprintf("[POST /users/current/computers/current/groups/{group_id}/backups][%d] initializeGroupBackupOK  %+v", 200, o.Payload)
}
func (o *InitializeGroupBackupOK) GetPayload() *models.Backup {
	return o.Payload
}

func (o *InitializeGroupBackupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Backup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInitializeGroupBackupBadRequest creates a InitializeGroupBackupBadRequest with default headers values
func NewInitializeGroupBackupBadRequest() *InitializeGroupBackupBadRequest {
	return &InitializeGroupBackupBadRequest{}
}

/* InitializeGroupBackupBadRequest describes a response with status code 400, with default header values.

If the backup is too big or there is any other problem.
*/
type InitializeGroupBackupBadRequest struct {
	Payload *models.Error
}

func (o *InitializeGroupBackupBadRequest) Error() string {
	return fmt.Sprintf("[POST /users/current/computers/current/groups/{group_id}/backups][%d] initializeGroupBackupBadRequest  %+v", 400, o.Payload)
}
func (o *InitializeGroupBackupBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *InitializeGroupBackupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInitializeGroupBackupUnauthorized creates a InitializeGroupBackupUnauthorized with default headers values
func NewInitializeGroupBackupUnauthorized() *InitializeGroupBackupUnauthorized {
	return &InitializeGroupBackupUnauthorized{}
}

/* InitializeGroupBackupUnauthorized describes a response with status code 401, with default header values.

Unauthorized.
*/
type InitializeGroupBackupUnauthorized struct {
	Payload *models.Error
}

func (o *InitializeGroupBackupUnauthorized) Error() string {
	return fmt.Sprintf("[POST /users/current/computers/current/groups/{group_id}/backups][%d] initializeGroupBackupUnauthorized  %+v", 401, o.Payload)
}
func (o *InitializeGroupBackupUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *InitializeGroupBackupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInitializeGroupBackupInternalServerError creates a InitializeGroupBackupInternalServerError with default headers values
func NewInitializeGroupBackupInternalServerError() *InitializeGroupBackupInternalServerError {
	return &InitializeGroupBackupInternalServerError{}
}

/* InitializeGroupBackupInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type InitializeGroupBackupInternalServerError struct {
	Payload *models.Error
}

func (o *InitializeGroupBackupInternalServerError) Error() string {
	return fmt.Sprintf("[POST /users/current/computers/current/groups/{group_id}/backups][%d] initializeGroupBackupInternalServerError  %+v", 500, o.Payload)
}
func (o *InitializeGroupBackupInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *InitializeGroupBackupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

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

// RequestBackupUploadReader is a Reader for the RequestBackupUpload structure.
type RequestBackupUploadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RequestBackupUploadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRequestBackupUploadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRequestBackupUploadBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewRequestBackupUploadUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRequestBackupUploadNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRequestBackupUploadInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewRequestBackupUploadOK creates a RequestBackupUploadOK with default headers values
func NewRequestBackupUploadOK() *RequestBackupUploadOK {
	return &RequestBackupUploadOK{}
}

/* RequestBackupUploadOK describes a response with status code 200, with default header values.

Upload request flag has been updated
*/
type RequestBackupUploadOK struct {
}

func (o *RequestBackupUploadOK) Error() string {
	return fmt.Sprintf("[PUT /users/current/computers/current/groups/{group_id}/backups/{backup_id}][%d] requestBackupUploadOK ", 200)
}

func (o *RequestBackupUploadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRequestBackupUploadBadRequest creates a RequestBackupUploadBadRequest with default headers values
func NewRequestBackupUploadBadRequest() *RequestBackupUploadBadRequest {
	return &RequestBackupUploadBadRequest{}
}

/* RequestBackupUploadBadRequest describes a response with status code 400, with default header values.

Supplied parameters were not okay.
*/
type RequestBackupUploadBadRequest struct {
	Payload *models.Error
}

func (o *RequestBackupUploadBadRequest) Error() string {
	return fmt.Sprintf("[PUT /users/current/computers/current/groups/{group_id}/backups/{backup_id}][%d] requestBackupUploadBadRequest  %+v", 400, o.Payload)
}
func (o *RequestBackupUploadBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *RequestBackupUploadBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRequestBackupUploadUnauthorized creates a RequestBackupUploadUnauthorized with default headers values
func NewRequestBackupUploadUnauthorized() *RequestBackupUploadUnauthorized {
	return &RequestBackupUploadUnauthorized{}
}

/* RequestBackupUploadUnauthorized describes a response with status code 401, with default header values.

Unauthorized.
*/
type RequestBackupUploadUnauthorized struct {
	Payload *models.Error
}

func (o *RequestBackupUploadUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /users/current/computers/current/groups/{group_id}/backups/{backup_id}][%d] requestBackupUploadUnauthorized  %+v", 401, o.Payload)
}
func (o *RequestBackupUploadUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *RequestBackupUploadUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRequestBackupUploadNotFound creates a RequestBackupUploadNotFound with default headers values
func NewRequestBackupUploadNotFound() *RequestBackupUploadNotFound {
	return &RequestBackupUploadNotFound{}
}

/* RequestBackupUploadNotFound describes a response with status code 404, with default header values.

The specified resource was not found.
*/
type RequestBackupUploadNotFound struct {
	Payload *models.Error
}

func (o *RequestBackupUploadNotFound) Error() string {
	return fmt.Sprintf("[PUT /users/current/computers/current/groups/{group_id}/backups/{backup_id}][%d] requestBackupUploadNotFound  %+v", 404, o.Payload)
}
func (o *RequestBackupUploadNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *RequestBackupUploadNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRequestBackupUploadInternalServerError creates a RequestBackupUploadInternalServerError with default headers values
func NewRequestBackupUploadInternalServerError() *RequestBackupUploadInternalServerError {
	return &RequestBackupUploadInternalServerError{}
}

/* RequestBackupUploadInternalServerError describes a response with status code 500, with default header values.

Internal server error.
*/
type RequestBackupUploadInternalServerError struct {
	Payload *models.Error
}

func (o *RequestBackupUploadInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /users/current/computers/current/groups/{group_id}/backups/{backup_id}][%d] requestBackupUploadInternalServerError  %+v", 500, o.Payload)
}
func (o *RequestBackupUploadInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *RequestBackupUploadInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
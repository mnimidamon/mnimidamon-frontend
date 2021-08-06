// Code generated by go-swagger; DO NOT EDIT.

package backup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewInitializeGroupBackupDeletionParams creates a new InitializeGroupBackupDeletionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInitializeGroupBackupDeletionParams() *InitializeGroupBackupDeletionParams {
	return &InitializeGroupBackupDeletionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInitializeGroupBackupDeletionParamsWithTimeout creates a new InitializeGroupBackupDeletionParams object
// with the ability to set a timeout on a request.
func NewInitializeGroupBackupDeletionParamsWithTimeout(timeout time.Duration) *InitializeGroupBackupDeletionParams {
	return &InitializeGroupBackupDeletionParams{
		timeout: timeout,
	}
}

// NewInitializeGroupBackupDeletionParamsWithContext creates a new InitializeGroupBackupDeletionParams object
// with the ability to set a context for a request.
func NewInitializeGroupBackupDeletionParamsWithContext(ctx context.Context) *InitializeGroupBackupDeletionParams {
	return &InitializeGroupBackupDeletionParams{
		Context: ctx,
	}
}

// NewInitializeGroupBackupDeletionParamsWithHTTPClient creates a new InitializeGroupBackupDeletionParams object
// with the ability to set a custom HTTPClient for a request.
func NewInitializeGroupBackupDeletionParamsWithHTTPClient(client *http.Client) *InitializeGroupBackupDeletionParams {
	return &InitializeGroupBackupDeletionParams{
		HTTPClient: client,
	}
}

/* InitializeGroupBackupDeletionParams contains all the parameters to send to the API endpoint
   for the initialize group backup deletion operation.

   Typically these are written to a http.Request.
*/
type InitializeGroupBackupDeletionParams struct {

	/* BackupID.

	   Numeric ID of the Backup.
	*/
	BackupID int64

	/* GroupID.

	   Numeric ID of the Model.
	*/
	GroupID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the initialize group backup deletion params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InitializeGroupBackupDeletionParams) WithDefaults() *InitializeGroupBackupDeletionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the initialize group backup deletion params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InitializeGroupBackupDeletionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the initialize group backup deletion params
func (o *InitializeGroupBackupDeletionParams) WithTimeout(timeout time.Duration) *InitializeGroupBackupDeletionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the initialize group backup deletion params
func (o *InitializeGroupBackupDeletionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the initialize group backup deletion params
func (o *InitializeGroupBackupDeletionParams) WithContext(ctx context.Context) *InitializeGroupBackupDeletionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the initialize group backup deletion params
func (o *InitializeGroupBackupDeletionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the initialize group backup deletion params
func (o *InitializeGroupBackupDeletionParams) WithHTTPClient(client *http.Client) *InitializeGroupBackupDeletionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the initialize group backup deletion params
func (o *InitializeGroupBackupDeletionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBackupID adds the backupID to the initialize group backup deletion params
func (o *InitializeGroupBackupDeletionParams) WithBackupID(backupID int64) *InitializeGroupBackupDeletionParams {
	o.SetBackupID(backupID)
	return o
}

// SetBackupID adds the backupId to the initialize group backup deletion params
func (o *InitializeGroupBackupDeletionParams) SetBackupID(backupID int64) {
	o.BackupID = backupID
}

// WithGroupID adds the groupID to the initialize group backup deletion params
func (o *InitializeGroupBackupDeletionParams) WithGroupID(groupID int64) *InitializeGroupBackupDeletionParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the initialize group backup deletion params
func (o *InitializeGroupBackupDeletionParams) SetGroupID(groupID int64) {
	o.GroupID = groupID
}

// WriteToRequest writes these params to a swagger request
func (o *InitializeGroupBackupDeletionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param backup_id
	if err := r.SetPathParam("backup_id", swag.FormatInt64(o.BackupID)); err != nil {
		return err
	}

	// path param group_id
	if err := r.SetPathParam("group_id", swag.FormatInt64(o.GroupID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

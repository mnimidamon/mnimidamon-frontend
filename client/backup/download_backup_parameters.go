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

// NewDownloadBackupParams creates a new DownloadBackupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDownloadBackupParams() *DownloadBackupParams {
	return &DownloadBackupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDownloadBackupParamsWithTimeout creates a new DownloadBackupParams object
// with the ability to set a timeout on a request.
func NewDownloadBackupParamsWithTimeout(timeout time.Duration) *DownloadBackupParams {
	return &DownloadBackupParams{
		timeout: timeout,
	}
}

// NewDownloadBackupParamsWithContext creates a new DownloadBackupParams object
// with the ability to set a context for a request.
func NewDownloadBackupParamsWithContext(ctx context.Context) *DownloadBackupParams {
	return &DownloadBackupParams{
		Context: ctx,
	}
}

// NewDownloadBackupParamsWithHTTPClient creates a new DownloadBackupParams object
// with the ability to set a custom HTTPClient for a request.
func NewDownloadBackupParamsWithHTTPClient(client *http.Client) *DownloadBackupParams {
	return &DownloadBackupParams{
		HTTPClient: client,
	}
}

/* DownloadBackupParams contains all the parameters to send to the API endpoint
   for the download backup operation.

   Typically these are written to a http.Request.
*/
type DownloadBackupParams struct {

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

// WithDefaults hydrates default values in the download backup params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DownloadBackupParams) WithDefaults() *DownloadBackupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the download backup params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DownloadBackupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the download backup params
func (o *DownloadBackupParams) WithTimeout(timeout time.Duration) *DownloadBackupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the download backup params
func (o *DownloadBackupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the download backup params
func (o *DownloadBackupParams) WithContext(ctx context.Context) *DownloadBackupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the download backup params
func (o *DownloadBackupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the download backup params
func (o *DownloadBackupParams) WithHTTPClient(client *http.Client) *DownloadBackupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the download backup params
func (o *DownloadBackupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBackupID adds the backupID to the download backup params
func (o *DownloadBackupParams) WithBackupID(backupID int64) *DownloadBackupParams {
	o.SetBackupID(backupID)
	return o
}

// SetBackupID adds the backupId to the download backup params
func (o *DownloadBackupParams) SetBackupID(backupID int64) {
	o.BackupID = backupID
}

// WithGroupID adds the groupID to the download backup params
func (o *DownloadBackupParams) WithGroupID(groupID int64) *DownloadBackupParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the download backup params
func (o *DownloadBackupParams) SetGroupID(groupID int64) {
	o.GroupID = groupID
}

// WriteToRequest writes these params to a swagger request
func (o *DownloadBackupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

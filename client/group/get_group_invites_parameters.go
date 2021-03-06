// Code generated by go-swagger; DO NOT EDIT.

package group

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

// NewGetGroupInvitesParams creates a new GetGroupInvitesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetGroupInvitesParams() *GetGroupInvitesParams {
	return &GetGroupInvitesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetGroupInvitesParamsWithTimeout creates a new GetGroupInvitesParams object
// with the ability to set a timeout on a request.
func NewGetGroupInvitesParamsWithTimeout(timeout time.Duration) *GetGroupInvitesParams {
	return &GetGroupInvitesParams{
		timeout: timeout,
	}
}

// NewGetGroupInvitesParamsWithContext creates a new GetGroupInvitesParams object
// with the ability to set a context for a request.
func NewGetGroupInvitesParamsWithContext(ctx context.Context) *GetGroupInvitesParams {
	return &GetGroupInvitesParams{
		Context: ctx,
	}
}

// NewGetGroupInvitesParamsWithHTTPClient creates a new GetGroupInvitesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetGroupInvitesParamsWithHTTPClient(client *http.Client) *GetGroupInvitesParams {
	return &GetGroupInvitesParams{
		HTTPClient: client,
	}
}

/* GetGroupInvitesParams contains all the parameters to send to the API endpoint
   for the get group invites operation.

   Typically these are written to a http.Request.
*/
type GetGroupInvitesParams struct {

	/* GroupID.

	   Numeric ID of the Group.
	*/
	GroupID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get group invites params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGroupInvitesParams) WithDefaults() *GetGroupInvitesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get group invites params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGroupInvitesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get group invites params
func (o *GetGroupInvitesParams) WithTimeout(timeout time.Duration) *GetGroupInvitesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get group invites params
func (o *GetGroupInvitesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get group invites params
func (o *GetGroupInvitesParams) WithContext(ctx context.Context) *GetGroupInvitesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get group invites params
func (o *GetGroupInvitesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get group invites params
func (o *GetGroupInvitesParams) WithHTTPClient(client *http.Client) *GetGroupInvitesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get group invites params
func (o *GetGroupInvitesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroupID adds the groupID to the get group invites params
func (o *GetGroupInvitesParams) WithGroupID(groupID int64) *GetGroupInvitesParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the get group invites params
func (o *GetGroupInvitesParams) SetGroupID(groupID int64) {
	o.GroupID = groupID
}

// WriteToRequest writes these params to a swagger request
func (o *GetGroupInvitesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param group_id
	if err := r.SetPathParam("group_id", swag.FormatInt64(o.GroupID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

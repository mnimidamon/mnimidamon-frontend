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

// NewGetGroupMembersParams creates a new GetGroupMembersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetGroupMembersParams() *GetGroupMembersParams {
	return &GetGroupMembersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetGroupMembersParamsWithTimeout creates a new GetGroupMembersParams object
// with the ability to set a timeout on a request.
func NewGetGroupMembersParamsWithTimeout(timeout time.Duration) *GetGroupMembersParams {
	return &GetGroupMembersParams{
		timeout: timeout,
	}
}

// NewGetGroupMembersParamsWithContext creates a new GetGroupMembersParams object
// with the ability to set a context for a request.
func NewGetGroupMembersParamsWithContext(ctx context.Context) *GetGroupMembersParams {
	return &GetGroupMembersParams{
		Context: ctx,
	}
}

// NewGetGroupMembersParamsWithHTTPClient creates a new GetGroupMembersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetGroupMembersParamsWithHTTPClient(client *http.Client) *GetGroupMembersParams {
	return &GetGroupMembersParams{
		HTTPClient: client,
	}
}

/* GetGroupMembersParams contains all the parameters to send to the API endpoint
   for the get group members operation.

   Typically these are written to a http.Request.
*/
type GetGroupMembersParams struct {

	/* GroupID.

	   Numeric ID of the Model.
	*/
	GroupID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get group members params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGroupMembersParams) WithDefaults() *GetGroupMembersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get group members params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGroupMembersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get group members params
func (o *GetGroupMembersParams) WithTimeout(timeout time.Duration) *GetGroupMembersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get group members params
func (o *GetGroupMembersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get group members params
func (o *GetGroupMembersParams) WithContext(ctx context.Context) *GetGroupMembersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get group members params
func (o *GetGroupMembersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get group members params
func (o *GetGroupMembersParams) WithHTTPClient(client *http.Client) *GetGroupMembersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get group members params
func (o *GetGroupMembersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroupID adds the groupID to the get group members params
func (o *GetGroupMembersParams) WithGroupID(groupID int64) *GetGroupMembersParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the get group members params
func (o *GetGroupMembersParams) SetGroupID(groupID int64) {
	o.GroupID = groupID
}

// WriteToRequest writes these params to a swagger request
func (o *GetGroupMembersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

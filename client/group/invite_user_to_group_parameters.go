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

	"mnimidamonbackend/models"
)

// NewInviteUserToGroupParams creates a new InviteUserToGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInviteUserToGroupParams() *InviteUserToGroupParams {
	return &InviteUserToGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInviteUserToGroupParamsWithTimeout creates a new InviteUserToGroupParams object
// with the ability to set a timeout on a request.
func NewInviteUserToGroupParamsWithTimeout(timeout time.Duration) *InviteUserToGroupParams {
	return &InviteUserToGroupParams{
		timeout: timeout,
	}
}

// NewInviteUserToGroupParamsWithContext creates a new InviteUserToGroupParams object
// with the ability to set a context for a request.
func NewInviteUserToGroupParamsWithContext(ctx context.Context) *InviteUserToGroupParams {
	return &InviteUserToGroupParams{
		Context: ctx,
	}
}

// NewInviteUserToGroupParamsWithHTTPClient creates a new InviteUserToGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewInviteUserToGroupParamsWithHTTPClient(client *http.Client) *InviteUserToGroupParams {
	return &InviteUserToGroupParams{
		HTTPClient: client,
	}
}

/* InviteUserToGroupParams contains all the parameters to send to the API endpoint
   for the invite user to group operation.

   Typically these are written to a http.Request.
*/
type InviteUserToGroupParams struct {

	/* Body.

	   Payload to invite a user
	*/
	Body *models.InviteUserPayload

	/* GroupID.

	   Numeric ID of the Group.
	*/
	GroupID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the invite user to group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InviteUserToGroupParams) WithDefaults() *InviteUserToGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the invite user to group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InviteUserToGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the invite user to group params
func (o *InviteUserToGroupParams) WithTimeout(timeout time.Duration) *InviteUserToGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the invite user to group params
func (o *InviteUserToGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the invite user to group params
func (o *InviteUserToGroupParams) WithContext(ctx context.Context) *InviteUserToGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the invite user to group params
func (o *InviteUserToGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the invite user to group params
func (o *InviteUserToGroupParams) WithHTTPClient(client *http.Client) *InviteUserToGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the invite user to group params
func (o *InviteUserToGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the invite user to group params
func (o *InviteUserToGroupParams) WithBody(body *models.InviteUserPayload) *InviteUserToGroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the invite user to group params
func (o *InviteUserToGroupParams) SetBody(body *models.InviteUserPayload) {
	o.Body = body
}

// WithGroupID adds the groupID to the invite user to group params
func (o *InviteUserToGroupParams) WithGroupID(groupID int64) *InviteUserToGroupParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the invite user to group params
func (o *InviteUserToGroupParams) SetGroupID(groupID int64) {
	o.GroupID = groupID
}

// WriteToRequest writes these params to a swagger request
func (o *InviteUserToGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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

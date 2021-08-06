// Code generated by go-swagger; DO NOT EDIT.

package group_computer

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

// NewJoinComputerToGroupParams creates a new JoinComputerToGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewJoinComputerToGroupParams() *JoinComputerToGroupParams {
	return &JoinComputerToGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewJoinComputerToGroupParamsWithTimeout creates a new JoinComputerToGroupParams object
// with the ability to set a timeout on a request.
func NewJoinComputerToGroupParamsWithTimeout(timeout time.Duration) *JoinComputerToGroupParams {
	return &JoinComputerToGroupParams{
		timeout: timeout,
	}
}

// NewJoinComputerToGroupParamsWithContext creates a new JoinComputerToGroupParams object
// with the ability to set a context for a request.
func NewJoinComputerToGroupParamsWithContext(ctx context.Context) *JoinComputerToGroupParams {
	return &JoinComputerToGroupParams{
		Context: ctx,
	}
}

// NewJoinComputerToGroupParamsWithHTTPClient creates a new JoinComputerToGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewJoinComputerToGroupParamsWithHTTPClient(client *http.Client) *JoinComputerToGroupParams {
	return &JoinComputerToGroupParams{
		HTTPClient: client,
	}
}

/* JoinComputerToGroupParams contains all the parameters to send to the API endpoint
   for the join computer to group operation.

   Typically these are written to a http.Request.
*/
type JoinComputerToGroupParams struct {

	/* Body.

	   Model creation payload.
	*/
	Body *models.CreateGroupComputerPayload

	/* GroupID.

	   Numeric ID of the Model.
	*/
	GroupID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the join computer to group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *JoinComputerToGroupParams) WithDefaults() *JoinComputerToGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the join computer to group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *JoinComputerToGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the join computer to group params
func (o *JoinComputerToGroupParams) WithTimeout(timeout time.Duration) *JoinComputerToGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the join computer to group params
func (o *JoinComputerToGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the join computer to group params
func (o *JoinComputerToGroupParams) WithContext(ctx context.Context) *JoinComputerToGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the join computer to group params
func (o *JoinComputerToGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the join computer to group params
func (o *JoinComputerToGroupParams) WithHTTPClient(client *http.Client) *JoinComputerToGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the join computer to group params
func (o *JoinComputerToGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the join computer to group params
func (o *JoinComputerToGroupParams) WithBody(body *models.CreateGroupComputerPayload) *JoinComputerToGroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the join computer to group params
func (o *JoinComputerToGroupParams) SetBody(body *models.CreateGroupComputerPayload) {
	o.Body = body
}

// WithGroupID adds the groupID to the join computer to group params
func (o *JoinComputerToGroupParams) WithGroupID(groupID int64) *JoinComputerToGroupParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the join computer to group params
func (o *JoinComputerToGroupParams) SetGroupID(groupID int64) {
	o.GroupID = groupID
}

// WriteToRequest writes these params to a swagger request
func (o *JoinComputerToGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

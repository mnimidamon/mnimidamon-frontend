// Code generated by go-swagger; DO NOT EDIT.

package current_user

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new current user API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for current user API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteCurrentUser(params *DeleteCurrentUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteCurrentUserAccepted, error)

	GetCurrentUser(params *GetCurrentUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCurrentUserOK, error)

	GetCurrentUserGroups(params *GetCurrentUserGroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCurrentUserGroupsOK, error)

	GetCurrentUserInvites(params *GetCurrentUserInvitesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCurrentUserInvitesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteCurrentUser deletes current user account
*/
func (a *Client) DeleteCurrentUser(params *DeleteCurrentUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteCurrentUserAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCurrentUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteCurrentUser",
		Method:             "DELETE",
		PathPattern:        "/users/current",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteCurrentUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteCurrentUserAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteCurrentUser: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCurrentUser gets current user profile
*/
func (a *Client) GetCurrentUser(params *GetCurrentUserParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCurrentUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCurrentUserParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCurrentUser",
		Method:             "GET",
		PathPattern:        "/users/current",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCurrentUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCurrentUserOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCurrentUser: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCurrentUserGroups gets current user groups
*/
func (a *Client) GetCurrentUserGroups(params *GetCurrentUserGroupsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCurrentUserGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCurrentUserGroupsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCurrentUserGroups",
		Method:             "GET",
		PathPattern:        "/users/current/groups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCurrentUserGroupsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCurrentUserGroupsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCurrentUserGroups: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetCurrentUserInvites gets group invites of current user
*/
func (a *Client) GetCurrentUserInvites(params *GetCurrentUserInvitesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetCurrentUserInvitesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCurrentUserInvitesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getCurrentUserInvites",
		Method:             "GET",
		PathPattern:        "/users/current/invites",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCurrentUserInvitesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCurrentUserInvitesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCurrentUserInvites: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
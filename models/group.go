// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Group Object that represents a Group.
//
// swagger:model Group
type Group struct {

	// Numeric identificator of the Group.
	// Example: 42
	// Read Only: true
	GroupID int64 `json:"group_id,omitempty"`

	// Name of the Group.
	// Example: damons
	// Max Length: 12
	// Min Length: 3
	Name string `json:"name,omitempty"`
}

// Validate validates this group
func (m *Group) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Group) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", m.Name, 3); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", m.Name, 12); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this group based on the context it is used
func (m *Group) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGroupID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Group) contextValidateGroupID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "group_id", "body", int64(m.GroupID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Group) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Group) UnmarshalBinary(b []byte) error {
	var res Group
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

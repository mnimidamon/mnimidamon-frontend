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

// GroupComputer Object that represents a Group Computer.
//
// swagger:model GroupComputer
type GroupComputer struct {

	// computer
	Computer *Computer `json:"computer,omitempty"`

	// Numeric identificator of the Computer.
	// Example: 42
	// Read Only: true
	ComputerID int64 `json:"computer_id,omitempty"`

	// Numeric identificatior of the Group.
	// Example: 42
	// Read Only: true
	GroupID int64 `json:"group_id,omitempty"`

	// How much space in MB does the User contribute to the Group.
	// Example: 1024
	StorageSize int64 `json:"storage_size,omitempty"`
}

// Validate validates this group computer
func (m *GroupComputer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateComputer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GroupComputer) validateComputer(formats strfmt.Registry) error {
	if swag.IsZero(m.Computer) { // not required
		return nil
	}

	if m.Computer != nil {
		if err := m.Computer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("computer")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this group computer based on the context it is used
func (m *GroupComputer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateComputer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateComputerID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGroupID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GroupComputer) contextValidateComputer(ctx context.Context, formats strfmt.Registry) error {

	if m.Computer != nil {
		if err := m.Computer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("computer")
			}
			return err
		}
	}

	return nil
}

func (m *GroupComputer) contextValidateComputerID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "computer_id", "body", int64(m.ComputerID)); err != nil {
		return err
	}

	return nil
}

func (m *GroupComputer) contextValidateGroupID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "group_id", "body", int64(m.GroupID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GroupComputer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GroupComputer) UnmarshalBinary(b []byte) error {
	var res GroupComputer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

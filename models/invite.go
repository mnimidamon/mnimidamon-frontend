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

// Invite Object that represents an active invite to a group.
//
// swagger:model Invite
type Invite struct {

	// The invitation date.
	// Format: date
	Date strfmt.Date `json:"date,omitempty"`

	// group
	Group *Group `json:"group,omitempty"`

	// Numeric identificator of the Invite.
	// Example: 42
	// Read Only: true
	InviteID int64 `json:"invite_id,omitempty"`

	// user
	User *User `json:"user,omitempty"`
}

// Validate validates this invite
func (m *Invite) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Invite) validateDate(formats strfmt.Registry) error {
	if swag.IsZero(m.Date) { // not required
		return nil
	}

	if err := validate.FormatOf("date", "body", "date", m.Date.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Invite) validateGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.Group) { // not required
		return nil
	}

	if m.Group != nil {
		if err := m.Group.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("group")
			}
			return err
		}
	}

	return nil
}

func (m *Invite) validateUser(formats strfmt.Registry) error {
	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this invite based on the context it is used
func (m *Invite) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInviteID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Invite) contextValidateGroup(ctx context.Context, formats strfmt.Registry) error {

	if m.Group != nil {
		if err := m.Group.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("group")
			}
			return err
		}
	}

	return nil
}

func (m *Invite) contextValidateInviteID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "invite_id", "body", int64(m.InviteID)); err != nil {
		return err
	}

	return nil
}

func (m *Invite) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

	if m.User != nil {
		if err := m.User.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Invite) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Invite) UnmarshalBinary(b []byte) error {
	var res Invite
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

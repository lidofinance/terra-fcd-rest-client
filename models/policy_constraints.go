// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PolicyConstraints policy constraints
//
// swagger:model PolicyConstraints
type PolicyConstraints struct {

	// cap
	Cap *PolicyConstraintsCap `json:"cap,omitempty"`

	// 0.025%
	// Example: 0.00025
	ChangeMax float32 `json:"change_max,omitempty"`

	// 1%
	// Example: 0.01
	RateMax float32 `json:"rate_max,omitempty"`

	// 0.05%
	// Example: 0.0005
	RateMin float32 `json:"rate_min,omitempty"`
}

// Validate validates this policy constraints
func (m *PolicyConstraints) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCap(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PolicyConstraints) validateCap(formats strfmt.Registry) error {
	if swag.IsZero(m.Cap) { // not required
		return nil
	}

	if m.Cap != nil {
		if err := m.Cap.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cap")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cap")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this policy constraints based on the context it is used
func (m *PolicyConstraints) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCap(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PolicyConstraints) contextValidateCap(ctx context.Context, formats strfmt.Registry) error {

	if m.Cap != nil {
		if err := m.Cap.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cap")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cap")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PolicyConstraints) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicyConstraints) UnmarshalBinary(b []byte) error {
	var res PolicyConstraints
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// PolicyConstraintsCap policy constraints cap
//
// swagger:model PolicyConstraintsCap
type PolicyConstraintsCap struct {

	// amount
	// Example: 50
	Amount string `json:"amount,omitempty"`

	// denom
	// Example: uluna
	Denom string `json:"denom,omitempty"`
}

// Validate validates this policy constraints cap
func (m *PolicyConstraintsCap) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this policy constraints cap based on context it is used
func (m *PolicyConstraintsCap) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PolicyConstraintsCap) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicyConstraintsCap) UnmarshalBinary(b []byte) error {
	var res PolicyConstraintsCap
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
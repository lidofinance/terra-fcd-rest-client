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

// ValidatorVotingPower validator voting power
//
// swagger:model validator.votingPower
type ValidatorVotingPower struct {

	// amount
	// Required: true
	Amount *string `json:"amount"`

	// weight
	// Required: true
	Weight *string `json:"weight"`
}

// Validate validates this validator voting power
func (m *ValidatorVotingPower) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeight(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ValidatorVotingPower) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *ValidatorVotingPower) validateWeight(formats strfmt.Registry) error {

	if err := validate.Required("weight", "body", m.Weight); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this validator voting power based on context it is used
func (m *ValidatorVotingPower) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ValidatorVotingPower) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ValidatorVotingPower) UnmarshalBinary(b []byte) error {
	var res ValidatorVotingPower
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

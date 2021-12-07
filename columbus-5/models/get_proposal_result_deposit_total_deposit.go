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

// GetProposalResultDepositTotalDeposit get proposal result deposit total deposit
//
// swagger:model getProposalResult.deposit.totalDeposit
type GetProposalResultDepositTotalDeposit struct {

	// denom amount
	// Required: true
	Amount *string `json:"amount"`

	// denom name
	// Required: true
	Denom *string `json:"denom"`
}

// Validate validates this get proposal result deposit total deposit
func (m *GetProposalResultDepositTotalDeposit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDenom(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetProposalResultDepositTotalDeposit) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *GetProposalResultDepositTotalDeposit) validateDenom(formats strfmt.Registry) error {

	if err := validate.Required("denom", "body", m.Denom); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get proposal result deposit total deposit based on context it is used
func (m *GetProposalResultDepositTotalDeposit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetProposalResultDepositTotalDeposit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetProposalResultDepositTotalDeposit) UnmarshalBinary(b []byte) error {
	var res GetProposalResultDepositTotalDeposit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

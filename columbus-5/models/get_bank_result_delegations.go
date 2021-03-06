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

// GetBankResultDelegations get bank result delegations
//
// swagger:model getBankResult.delegations
type GetBankResultDelegations struct {

	// delegation amount
	// Required: true
	Amount *string `json:"amount"`

	// delegator address
	// Required: true
	DelegatorAddress *string `json:"delegator_address"`

	// delegation share
	// Required: true
	Shares *string `json:"shares"`

	// validator address
	// Required: true
	ValidatorAddress *string `json:"validator_address"`
}

// Validate validates this get bank result delegations
func (m *GetBankResultDelegations) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDelegatorAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShares(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidatorAddress(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetBankResultDelegations) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *GetBankResultDelegations) validateDelegatorAddress(formats strfmt.Registry) error {

	if err := validate.Required("delegator_address", "body", m.DelegatorAddress); err != nil {
		return err
	}

	return nil
}

func (m *GetBankResultDelegations) validateShares(formats strfmt.Registry) error {

	if err := validate.Required("shares", "body", m.Shares); err != nil {
		return err
	}

	return nil
}

func (m *GetBankResultDelegations) validateValidatorAddress(formats strfmt.Registry) error {

	if err := validate.Required("validator_address", "body", m.ValidatorAddress); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get bank result delegations based on context it is used
func (m *GetBankResultDelegations) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetBankResultDelegations) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetBankResultDelegations) UnmarshalBinary(b []byte) error {
	var res GetBankResultDelegations
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

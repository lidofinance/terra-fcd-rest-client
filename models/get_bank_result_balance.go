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

// GetBankResultBalance get bank result balance
//
// swagger:model getBankResult.balance
type GetBankResultBalance struct {

	// Available amount
	// Required: true
	Amount *string `json:"amount"`

	// Delegatable amount
	// Required: true
	Delegatable *string `json:"delegatable"`

	// Delegated amount of the vesting amount
	// Required: true
	DelegatedVesting *string `json:"delegatedVesting"`

	// Coin denomination
	// Required: true
	Denom *string `json:"denom"`

	// Freed amount of the vesting amount
	// Required: true
	FreedVesting *string `json:"freedVesting"`

	// Amount not yet freed
	// Required: true
	RemainingVesting *string `json:"remainingVesting"`

	// Amount in unbonding state
	// Required: true
	Unbonding *string `json:"unbonding"`
}

// Validate validates this get bank result balance
func (m *GetBankResultBalance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDelegatable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDelegatedVesting(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDenom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFreedVesting(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemainingVesting(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnbonding(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetBankResultBalance) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *GetBankResultBalance) validateDelegatable(formats strfmt.Registry) error {

	if err := validate.Required("delegatable", "body", m.Delegatable); err != nil {
		return err
	}

	return nil
}

func (m *GetBankResultBalance) validateDelegatedVesting(formats strfmt.Registry) error {

	if err := validate.Required("delegatedVesting", "body", m.DelegatedVesting); err != nil {
		return err
	}

	return nil
}

func (m *GetBankResultBalance) validateDenom(formats strfmt.Registry) error {

	if err := validate.Required("denom", "body", m.Denom); err != nil {
		return err
	}

	return nil
}

func (m *GetBankResultBalance) validateFreedVesting(formats strfmt.Registry) error {

	if err := validate.Required("freedVesting", "body", m.FreedVesting); err != nil {
		return err
	}

	return nil
}

func (m *GetBankResultBalance) validateRemainingVesting(formats strfmt.Registry) error {

	if err := validate.Required("remainingVesting", "body", m.RemainingVesting); err != nil {
		return err
	}

	return nil
}

func (m *GetBankResultBalance) validateUnbonding(formats strfmt.Registry) error {

	if err := validate.Required("unbonding", "body", m.Unbonding); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get bank result balance based on context it is used
func (m *GetBankResultBalance) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetBankResultBalance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetBankResultBalance) UnmarshalBinary(b []byte) error {
	var res GetBankResultBalance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
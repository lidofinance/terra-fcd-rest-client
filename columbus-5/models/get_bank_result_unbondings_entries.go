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

// GetBankResultUnbondingsEntries get bank result unbondings entries
//
// swagger:model getBankResult.unbondings.entries
type GetBankResultUnbondingsEntries struct {

	// current balance
	// Required: true
	Balance *string `json:"balance"`

	// unbonding completion time
	// Required: true
	CompletionTime *string `json:"completion_time"`

	// block height
	// Required: true
	CreatingHeight *string `json:"creating_height"`

	// initial balancd
	// Required: true
	InitialBalance *string `json:"initial_balance"`
}

// Validate validates this get bank result unbondings entries
func (m *GetBankResultUnbondingsEntries) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBalance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCompletionTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatingHeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitialBalance(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetBankResultUnbondingsEntries) validateBalance(formats strfmt.Registry) error {

	if err := validate.Required("balance", "body", m.Balance); err != nil {
		return err
	}

	return nil
}

func (m *GetBankResultUnbondingsEntries) validateCompletionTime(formats strfmt.Registry) error {

	if err := validate.Required("completion_time", "body", m.CompletionTime); err != nil {
		return err
	}

	return nil
}

func (m *GetBankResultUnbondingsEntries) validateCreatingHeight(formats strfmt.Registry) error {

	if err := validate.Required("creating_height", "body", m.CreatingHeight); err != nil {
		return err
	}

	return nil
}

func (m *GetBankResultUnbondingsEntries) validateInitialBalance(formats strfmt.Registry) error {

	if err := validate.Required("initial_balance", "body", m.InitialBalance); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get bank result unbondings entries based on context it is used
func (m *GetBankResultUnbondingsEntries) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetBankResultUnbondingsEntries) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetBankResultUnbondingsEntries) UnmarshalBinary(b []byte) error {
	var res GetBankResultUnbondingsEntries
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

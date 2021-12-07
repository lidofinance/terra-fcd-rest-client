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

// TxsTxValueFeeAmount txs tx value fee amount
//
// swagger:model txs.tx.value.fee.amount
type TxsTxValueFeeAmount struct {

	// amount
	// Required: true
	Amount *string `json:"amount"`

	// denom
	// Required: true
	Denom *string `json:"denom"`
}

// Validate validates this txs tx value fee amount
func (m *TxsTxValueFeeAmount) Validate(formats strfmt.Registry) error {
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

func (m *TxsTxValueFeeAmount) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *TxsTxValueFeeAmount) validateDenom(formats strfmt.Registry) error {

	if err := validate.Required("denom", "body", m.Denom); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this txs tx value fee amount based on context it is used
func (m *TxsTxValueFeeAmount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TxsTxValueFeeAmount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TxsTxValueFeeAmount) UnmarshalBinary(b []byte) error {
	var res TxsTxValueFeeAmount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

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

// GetTxListResultTxsTxValueFeeAmount get tx list result txs tx value fee amount
//
// swagger:model getTxListResult.txs.tx.value.fee.amount
type GetTxListResultTxsTxValueFeeAmount struct {

	// amount
	// Required: true
	Amount *string `json:"amount"`

	// denom
	// Required: true
	Denom *string `json:"denom"`
}

// Validate validates this get tx list result txs tx value fee amount
func (m *GetTxListResultTxsTxValueFeeAmount) Validate(formats strfmt.Registry) error {
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

func (m *GetTxListResultTxsTxValueFeeAmount) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *GetTxListResultTxsTxValueFeeAmount) validateDenom(formats strfmt.Registry) error {

	if err := validate.Required("denom", "body", m.Denom); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get tx list result txs tx value fee amount based on context it is used
func (m *GetTxListResultTxsTxValueFeeAmount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetTxListResultTxsTxValueFeeAmount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetTxListResultTxsTxValueFeeAmount) UnmarshalBinary(b []byte) error {
	var res GetTxListResultTxsTxValueFeeAmount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

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

// TxsTxValueMsgValueOutputsCoins txs tx value msg value outputs coins
//
// swagger:model txs.tx.value.msg.value.outputs.coins
type TxsTxValueMsgValueOutputsCoins struct {

	// amount
	// Required: true
	Amount *string `json:"amount"`

	// deonm
	// Required: true
	Deonm *string `json:"deonm"`
}

// Validate validates this txs tx value msg value outputs coins
func (m *TxsTxValueMsgValueOutputsCoins) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeonm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TxsTxValueMsgValueOutputsCoins) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *TxsTxValueMsgValueOutputsCoins) validateDeonm(formats strfmt.Registry) error {

	if err := validate.Required("deonm", "body", m.Deonm); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this txs tx value msg value outputs coins based on context it is used
func (m *TxsTxValueMsgValueOutputsCoins) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TxsTxValueMsgValueOutputsCoins) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TxsTxValueMsgValueOutputsCoins) UnmarshalBinary(b []byte) error {
	var res TxsTxValueMsgValueOutputsCoins
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

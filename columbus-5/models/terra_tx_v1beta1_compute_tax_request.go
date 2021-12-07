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

// TerraTxV1beta1ComputeTaxRequest ComputeTaxRequest is the request type for the Service.ComputeTax
// RPC method.
//
// swagger:model terra.tx.v1beta1.ComputeTaxRequest
type TerraTxV1beta1ComputeTaxRequest struct {

	// tx is the transaction to simulate.
	// Deprecated. Send raw tx bytes instead.
	Tx *CosmosTxV1beta1Tx `json:"tx,omitempty"`

	// tx_bytes is the raw transaction.
	// Format: byte
	TxBytes strfmt.Base64 `json:"tx_bytes,omitempty"`
}

// Validate validates this terra tx v1beta1 compute tax request
func (m *TerraTxV1beta1ComputeTaxRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTx(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TerraTxV1beta1ComputeTaxRequest) validateTx(formats strfmt.Registry) error {
	if swag.IsZero(m.Tx) { // not required
		return nil
	}

	if m.Tx != nil {
		if err := m.Tx.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tx")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tx")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this terra tx v1beta1 compute tax request based on the context it is used
func (m *TerraTxV1beta1ComputeTaxRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTx(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TerraTxV1beta1ComputeTaxRequest) contextValidateTx(ctx context.Context, formats strfmt.Registry) error {

	if m.Tx != nil {
		if err := m.Tx.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tx")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tx")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TerraTxV1beta1ComputeTaxRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TerraTxV1beta1ComputeTaxRequest) UnmarshalBinary(b []byte) error {
	var res TerraTxV1beta1ComputeTaxRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

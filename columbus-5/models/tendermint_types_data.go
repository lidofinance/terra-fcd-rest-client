// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TendermintTypesData Data contains the set of transactions included in the block
//
// swagger:model tendermint.types.Data
type TendermintTypesData struct {

	// Txs that will be applied by state @ block.Height+1.
	// NOTE: not all txs here are valid.  We're just agreeing on the order first.
	// This means that block.AppHash does not include these txs.
	Txs []strfmt.Base64 `json:"txs"`
}

// Validate validates this tendermint types data
func (m *TendermintTypesData) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this tendermint types data based on context it is used
func (m *TendermintTypesData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TendermintTypesData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TendermintTypesData) UnmarshalBinary(b []byte) error {
	var res TendermintTypesData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

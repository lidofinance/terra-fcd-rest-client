// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TerraTreasuryV1beta1QueryTaxCapsResponseItem QueryTaxCapsResponseItem is response item type for the
// Query/TaxCaps RPC method.
//
// swagger:model terra.treasury.v1beta1.QueryTaxCapsResponseItem
type TerraTreasuryV1beta1QueryTaxCapsResponseItem struct {

	// denom
	Denom string `json:"denom,omitempty"`

	// tax cap
	TaxCap string `json:"tax_cap,omitempty"`
}

// Validate validates this terra treasury v1beta1 query tax caps response item
func (m *TerraTreasuryV1beta1QueryTaxCapsResponseItem) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this terra treasury v1beta1 query tax caps response item based on context it is used
func (m *TerraTreasuryV1beta1QueryTaxCapsResponseItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TerraTreasuryV1beta1QueryTaxCapsResponseItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TerraTreasuryV1beta1QueryTaxCapsResponseItem) UnmarshalBinary(b []byte) error {
	var res TerraTreasuryV1beta1QueryTaxCapsResponseItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

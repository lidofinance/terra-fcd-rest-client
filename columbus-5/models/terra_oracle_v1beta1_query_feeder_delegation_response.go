// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TerraOracleV1beta1QueryFeederDelegationResponse QueryFeederDelegationResponse is response type for the
// Query/FeederDelegation RPC method.
//
// swagger:model terra.oracle.v1beta1.QueryFeederDelegationResponse
type TerraOracleV1beta1QueryFeederDelegationResponse struct {

	// feeder_addr defines the feeder delegation of a validator
	FeederAddr string `json:"feeder_addr,omitempty"`
}

// Validate validates this terra oracle v1beta1 query feeder delegation response
func (m *TerraOracleV1beta1QueryFeederDelegationResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this terra oracle v1beta1 query feeder delegation response based on context it is used
func (m *TerraOracleV1beta1QueryFeederDelegationResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TerraOracleV1beta1QueryFeederDelegationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TerraOracleV1beta1QueryFeederDelegationResponse) UnmarshalBinary(b []byte) error {
	var res TerraOracleV1beta1QueryFeederDelegationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

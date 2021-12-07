// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CosmosBankV1beta1SendEnabled SendEnabled maps coin denom to a send_enabled status (whether a denom is
// sendable).
//
// swagger:model cosmos.bank.v1beta1.SendEnabled
type CosmosBankV1beta1SendEnabled struct {

	// denom
	Denom string `json:"denom,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`
}

// Validate validates this cosmos bank v1beta1 send enabled
func (m *CosmosBankV1beta1SendEnabled) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cosmos bank v1beta1 send enabled based on context it is used
func (m *CosmosBankV1beta1SendEnabled) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CosmosBankV1beta1SendEnabled) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosBankV1beta1SendEnabled) UnmarshalBinary(b []byte) error {
	var res CosmosBankV1beta1SendEnabled
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
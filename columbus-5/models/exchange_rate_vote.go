// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ExchangeRateVote exchange rate vote
//
// swagger:model ExchangeRateVote
type ExchangeRateVote struct {

	// denom
	// Example: ukrw
	Denom string `json:"denom,omitempty"`

	// exchange rate
	// Example: 0.01241
	ExchangeRate float64 `json:"exchange_rate,omitempty"`

	// bech32 encoded address
	// Example: terravaloper1wg2mlrxdmnnkkykgqg4znky86nyrtc45q7a85l
	Voter string `json:"voter,omitempty"`
}

// Validate validates this exchange rate vote
func (m *ExchangeRateVote) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this exchange rate vote based on context it is used
func (m *ExchangeRateVote) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ExchangeRateVote) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExchangeRateVote) UnmarshalBinary(b []byte) error {
	var res ExchangeRateVote
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

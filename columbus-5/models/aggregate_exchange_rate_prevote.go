// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AggregateExchangeRatePrevote aggregate exchange rate prevote
//
// swagger:model AggregateExchangeRatePrevote
type AggregateExchangeRatePrevote struct {

	// hash
	// Example: 061bf1e27dfff121f40c826e593c8a28ec299a02
	Hash string `json:"hash,omitempty"`

	// submit block
	// Example: 1
	SubmitBlock float64 `json:"submit_block,omitempty"`

	// bech32 encoded address
	// Example: terravaloper1wg2mlrxdmnnkkykgqg4znky86nyrtc45q7a85l
	Voter string `json:"voter,omitempty"`
}

// Validate validates this aggregate exchange rate prevote
func (m *AggregateExchangeRatePrevote) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this aggregate exchange rate prevote based on context it is used
func (m *AggregateExchangeRatePrevote) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AggregateExchangeRatePrevote) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AggregateExchangeRatePrevote) UnmarshalBinary(b []byte) error {
	var res AggregateExchangeRatePrevote
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
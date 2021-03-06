// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TendermintVersionConsensus Consensus captures the consensus rules for processing a block in the blockchain,
// including all blockchain data structures and the rules of the application's
// state transition machine.
//
// swagger:model tendermint.version.Consensus
type TendermintVersionConsensus struct {

	// app
	App string `json:"app,omitempty"`

	// block
	Block string `json:"block,omitempty"`
}

// Validate validates this tendermint version consensus
func (m *TendermintVersionConsensus) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this tendermint version consensus based on context it is used
func (m *TendermintVersionConsensus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TendermintVersionConsensus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TendermintVersionConsensus) UnmarshalBinary(b []byte) error {
	var res TendermintVersionConsensus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

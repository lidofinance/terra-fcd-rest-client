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

// TerraOracleV1beta1QueryAggregatePrevoteResponse QueryAggregatePrevoteResponse is response type for the
// Query/AggregatePrevote RPC method.
//
// swagger:model terra.oracle.v1beta1.QueryAggregatePrevoteResponse
type TerraOracleV1beta1QueryAggregatePrevoteResponse struct {

	// aggregate prevote
	AggregatePrevote *TerraOracleV1beta1QueryAggregatePrevoteResponseAggregatePrevote `json:"aggregate_prevote,omitempty"`
}

// Validate validates this terra oracle v1beta1 query aggregate prevote response
func (m *TerraOracleV1beta1QueryAggregatePrevoteResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAggregatePrevote(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TerraOracleV1beta1QueryAggregatePrevoteResponse) validateAggregatePrevote(formats strfmt.Registry) error {
	if swag.IsZero(m.AggregatePrevote) { // not required
		return nil
	}

	if m.AggregatePrevote != nil {
		if err := m.AggregatePrevote.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aggregate_prevote")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aggregate_prevote")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this terra oracle v1beta1 query aggregate prevote response based on the context it is used
func (m *TerraOracleV1beta1QueryAggregatePrevoteResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAggregatePrevote(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TerraOracleV1beta1QueryAggregatePrevoteResponse) contextValidateAggregatePrevote(ctx context.Context, formats strfmt.Registry) error {

	if m.AggregatePrevote != nil {
		if err := m.AggregatePrevote.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aggregate_prevote")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aggregate_prevote")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TerraOracleV1beta1QueryAggregatePrevoteResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TerraOracleV1beta1QueryAggregatePrevoteResponse) UnmarshalBinary(b []byte) error {
	var res TerraOracleV1beta1QueryAggregatePrevoteResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TerraOracleV1beta1QueryAggregatePrevoteResponseAggregatePrevote struct for aggregate prevoting on the ExchangeRateVote.
// The purpose of aggregate prevote is to hide vote exchange rates with hash
// which is formatted as hex string in SHA256("{salt}:{exchange rate}{denom},...,{exchange rate}{denom}:{voter}")
//
// swagger:model TerraOracleV1beta1QueryAggregatePrevoteResponseAggregatePrevote
type TerraOracleV1beta1QueryAggregatePrevoteResponseAggregatePrevote struct {

	// hash
	Hash string `json:"hash,omitempty"`

	// submit block
	SubmitBlock string `json:"submit_block,omitempty"`

	// voter
	Voter string `json:"voter,omitempty"`
}

// Validate validates this terra oracle v1beta1 query aggregate prevote response aggregate prevote
func (m *TerraOracleV1beta1QueryAggregatePrevoteResponseAggregatePrevote) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this terra oracle v1beta1 query aggregate prevote response aggregate prevote based on context it is used
func (m *TerraOracleV1beta1QueryAggregatePrevoteResponseAggregatePrevote) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TerraOracleV1beta1QueryAggregatePrevoteResponseAggregatePrevote) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TerraOracleV1beta1QueryAggregatePrevoteResponseAggregatePrevote) UnmarshalBinary(b []byte) error {
	var res TerraOracleV1beta1QueryAggregatePrevoteResponseAggregatePrevote
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

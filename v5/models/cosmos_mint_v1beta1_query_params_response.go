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

// CosmosMintV1beta1QueryParamsResponse QueryParamsResponse is the response type for the Query/Params RPC method.
//
// swagger:model cosmos.mint.v1beta1.QueryParamsResponse
type CosmosMintV1beta1QueryParamsResponse struct {

	// params
	Params *CosmosMintV1beta1QueryParamsResponseParams `json:"params,omitempty"`
}

// Validate validates this cosmos mint v1beta1 query params response
func (m *CosmosMintV1beta1QueryParamsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosMintV1beta1QueryParamsResponse) validateParams(formats strfmt.Registry) error {
	if swag.IsZero(m.Params) { // not required
		return nil
	}

	if m.Params != nil {
		if err := m.Params.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("params")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("params")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cosmos mint v1beta1 query params response based on the context it is used
func (m *CosmosMintV1beta1QueryParamsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosMintV1beta1QueryParamsResponse) contextValidateParams(ctx context.Context, formats strfmt.Registry) error {

	if m.Params != nil {
		if err := m.Params.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("params")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("params")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CosmosMintV1beta1QueryParamsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosMintV1beta1QueryParamsResponse) UnmarshalBinary(b []byte) error {
	var res CosmosMintV1beta1QueryParamsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CosmosMintV1beta1QueryParamsResponseParams params defines the parameters of the module.
//
// swagger:model CosmosMintV1beta1QueryParamsResponseParams
type CosmosMintV1beta1QueryParamsResponseParams struct {

	// expected blocks per year
	BlocksPerYear string `json:"blocks_per_year,omitempty"`

	// goal of percent bonded atoms
	GoalBonded string `json:"goal_bonded,omitempty"`

	// maximum inflation rate
	InflationMax string `json:"inflation_max,omitempty"`

	// minimum inflation rate
	InflationMin string `json:"inflation_min,omitempty"`

	// maximum annual change in inflation rate
	InflationRateChange string `json:"inflation_rate_change,omitempty"`

	// type of coin to mint
	MintDenom string `json:"mint_denom,omitempty"`
}

// Validate validates this cosmos mint v1beta1 query params response params
func (m *CosmosMintV1beta1QueryParamsResponseParams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cosmos mint v1beta1 query params response params based on context it is used
func (m *CosmosMintV1beta1QueryParamsResponseParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CosmosMintV1beta1QueryParamsResponseParams) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosMintV1beta1QueryParamsResponseParams) UnmarshalBinary(b []byte) error {
	var res CosmosMintV1beta1QueryParamsResponseParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
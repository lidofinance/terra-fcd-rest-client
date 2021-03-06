// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CosmosParamsV1beta1ParamChange ParamChange defines an individual parameter change, for use in
// ParameterChangeProposal.
//
// swagger:model cosmos.params.v1beta1.ParamChange
type CosmosParamsV1beta1ParamChange struct {

	// key
	Key string `json:"key,omitempty"`

	// subspace
	Subspace string `json:"subspace,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this cosmos params v1beta1 param change
func (m *CosmosParamsV1beta1ParamChange) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cosmos params v1beta1 param change based on context it is used
func (m *CosmosParamsV1beta1ParamChange) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CosmosParamsV1beta1ParamChange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosParamsV1beta1ParamChange) UnmarshalBinary(b []byte) error {
	var res CosmosParamsV1beta1ParamChange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

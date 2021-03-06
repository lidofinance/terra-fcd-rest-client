// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CosmosDistributionV1beta1QueryDelegatorValidatorsResponse QueryDelegatorValidatorsResponse is the response type for the
// Query/DelegatorValidators RPC method.
//
// swagger:model cosmos.distribution.v1beta1.QueryDelegatorValidatorsResponse
type CosmosDistributionV1beta1QueryDelegatorValidatorsResponse struct {

	// validators defines the validators a delegator is delegating for.
	Validators []string `json:"validators"`
}

// Validate validates this cosmos distribution v1beta1 query delegator validators response
func (m *CosmosDistributionV1beta1QueryDelegatorValidatorsResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cosmos distribution v1beta1 query delegator validators response based on context it is used
func (m *CosmosDistributionV1beta1QueryDelegatorValidatorsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CosmosDistributionV1beta1QueryDelegatorValidatorsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosDistributionV1beta1QueryDelegatorValidatorsResponse) UnmarshalBinary(b []byte) error {
	var res CosmosDistributionV1beta1QueryDelegatorValidatorsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

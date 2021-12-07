// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CosmosGovV1beta1TallyResult TallyResult defines a standard tally for a governance proposal.
//
// swagger:model cosmos.gov.v1beta1.TallyResult
type CosmosGovV1beta1TallyResult struct {

	// abstain
	Abstain string `json:"abstain,omitempty"`

	// no
	No string `json:"no,omitempty"`

	// no with veto
	NoWithVeto string `json:"no_with_veto,omitempty"`

	// yes
	Yes string `json:"yes,omitempty"`
}

// Validate validates this cosmos gov v1beta1 tally result
func (m *CosmosGovV1beta1TallyResult) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cosmos gov v1beta1 tally result based on context it is used
func (m *CosmosGovV1beta1TallyResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CosmosGovV1beta1TallyResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosGovV1beta1TallyResult) UnmarshalBinary(b []byte) error {
	var res CosmosGovV1beta1TallyResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CosmosMintV1beta1QueryAnnualProvisionsResponse QueryAnnualProvisionsResponse is the response type for the
// Query/AnnualProvisions RPC method.
//
// swagger:model cosmos.mint.v1beta1.QueryAnnualProvisionsResponse
type CosmosMintV1beta1QueryAnnualProvisionsResponse struct {

	// annual_provisions is the current minting annual provisions value.
	// Format: byte
	AnnualProvisions strfmt.Base64 `json:"annual_provisions,omitempty"`
}

// Validate validates this cosmos mint v1beta1 query annual provisions response
func (m *CosmosMintV1beta1QueryAnnualProvisionsResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cosmos mint v1beta1 query annual provisions response based on context it is used
func (m *CosmosMintV1beta1QueryAnnualProvisionsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CosmosMintV1beta1QueryAnnualProvisionsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosMintV1beta1QueryAnnualProvisionsResponse) UnmarshalBinary(b []byte) error {
	var res CosmosMintV1beta1QueryAnnualProvisionsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

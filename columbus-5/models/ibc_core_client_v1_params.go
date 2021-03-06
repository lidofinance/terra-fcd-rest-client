// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IbcCoreClientV1Params Params defines the set of IBC light client parameters.
//
// swagger:model ibc.core.client.v1.Params
type IbcCoreClientV1Params struct {

	// allowed_clients defines the list of allowed client state types.
	AllowedClients []string `json:"allowed_clients"`
}

// Validate validates this ibc core client v1 params
func (m *IbcCoreClientV1Params) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ibc core client v1 params based on context it is used
func (m *IbcCoreClientV1Params) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IbcCoreClientV1Params) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IbcCoreClientV1Params) UnmarshalBinary(b []byte) error {
	var res IbcCoreClientV1Params
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

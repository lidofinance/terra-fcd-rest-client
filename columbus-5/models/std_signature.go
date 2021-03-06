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

// StdSignature std signature
//
// swagger:model StdSignature
type StdSignature struct {

	// pub key
	PubKey *StdSignaturePubKey `json:"pub_key,omitempty"`

	// signature
	// Example: MEUCIQD02fsDPra8MtbRsyB1w7bqTM55Wu138zQbFcWx4+CFyAIge5WNPfKIuvzBZ69MyqHsqD8S1IwiEp+iUb6VSdtlpgY=
	Signature string `json:"signature,omitempty"`
}

// Validate validates this std signature
func (m *StdSignature) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePubKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StdSignature) validatePubKey(formats strfmt.Registry) error {
	if swag.IsZero(m.PubKey) { // not required
		return nil
	}

	if m.PubKey != nil {
		if err := m.PubKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pub_key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pub_key")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this std signature based on the context it is used
func (m *StdSignature) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePubKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StdSignature) contextValidatePubKey(ctx context.Context, formats strfmt.Registry) error {

	if m.PubKey != nil {
		if err := m.PubKey.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pub_key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pub_key")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StdSignature) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StdSignature) UnmarshalBinary(b []byte) error {
	var res StdSignature
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// StdSignaturePubKey std signature pub key
//
// swagger:model StdSignaturePubKey
type StdSignaturePubKey struct {

	// type
	// Example: tendermint/PubKeySecp256k1
	Type string `json:"type,omitempty"`

	// value
	// Example: Avz04VhtKJh8ACCVzlI8aTosGy0ikFXKIVHQ3jKMrosH
	Value string `json:"value,omitempty"`
}

// Validate validates this std signature pub key
func (m *StdSignaturePubKey) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this std signature pub key based on context it is used
func (m *StdSignaturePubKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StdSignaturePubKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StdSignaturePubKey) UnmarshalBinary(b []byte) error {
	var res StdSignaturePubKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

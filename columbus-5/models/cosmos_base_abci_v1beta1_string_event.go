// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CosmosBaseAbciV1beta1StringEvent StringEvent defines en Event object wrapper where all the attributes
// contain key/value pairs that are strings instead of raw bytes.
//
// swagger:model cosmos.base.abci.v1beta1.StringEvent
type CosmosBaseAbciV1beta1StringEvent struct {

	// attributes
	Attributes []*CosmosBaseAbciV1beta1StringEventAttributesItems0 `json:"attributes"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this cosmos base abci v1beta1 string event
func (m *CosmosBaseAbciV1beta1StringEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosBaseAbciV1beta1StringEvent) validateAttributes(formats strfmt.Registry) error {
	if swag.IsZero(m.Attributes) { // not required
		return nil
	}

	for i := 0; i < len(m.Attributes); i++ {
		if swag.IsZero(m.Attributes[i]) { // not required
			continue
		}

		if m.Attributes[i] != nil {
			if err := m.Attributes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attributes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attributes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this cosmos base abci v1beta1 string event based on the context it is used
func (m *CosmosBaseAbciV1beta1StringEvent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosBaseAbciV1beta1StringEvent) contextValidateAttributes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Attributes); i++ {

		if m.Attributes[i] != nil {
			if err := m.Attributes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attributes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attributes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CosmosBaseAbciV1beta1StringEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosBaseAbciV1beta1StringEvent) UnmarshalBinary(b []byte) error {
	var res CosmosBaseAbciV1beta1StringEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CosmosBaseAbciV1beta1StringEventAttributesItems0 Attribute defines an attribute wrapper where the key and value are
// strings instead of raw bytes.
//
// swagger:model CosmosBaseAbciV1beta1StringEventAttributesItems0
type CosmosBaseAbciV1beta1StringEventAttributesItems0 struct {

	// key
	Key string `json:"key,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this cosmos base abci v1beta1 string event attributes items0
func (m *CosmosBaseAbciV1beta1StringEventAttributesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cosmos base abci v1beta1 string event attributes items0 based on context it is used
func (m *CosmosBaseAbciV1beta1StringEventAttributesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CosmosBaseAbciV1beta1StringEventAttributesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosBaseAbciV1beta1StringEventAttributesItems0) UnmarshalBinary(b []byte) error {
	var res CosmosBaseAbciV1beta1StringEventAttributesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

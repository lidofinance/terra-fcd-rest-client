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

// CosmosBankV1beta1Metadata Metadata represents a struct that describes
// a basic token.
//
// swagger:model cosmos.bank.v1beta1.Metadata
type CosmosBankV1beta1Metadata struct {

	// base represents the base denom (should be the DenomUnit with exponent = 0).
	Base string `json:"base,omitempty"`

	// denom_units represents the list of DenomUnit's for a given coin
	DenomUnits []*CosmosBankV1beta1MetadataDenomUnitsItems0 `json:"denom_units"`

	// description
	Description string `json:"description,omitempty"`

	// display indicates the suggested denom that should be
	// displayed in clients.
	Display string `json:"display,omitempty"`

	// name defines the name of the token (eg: Cosmos Atom)
	Name string `json:"name,omitempty"`

	// symbol is the token symbol usually shown on exchanges (eg: ATOM). This can
	// be the same as the display.
	Symbol string `json:"symbol,omitempty"`
}

// Validate validates this cosmos bank v1beta1 metadata
func (m *CosmosBankV1beta1Metadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDenomUnits(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosBankV1beta1Metadata) validateDenomUnits(formats strfmt.Registry) error {
	if swag.IsZero(m.DenomUnits) { // not required
		return nil
	}

	for i := 0; i < len(m.DenomUnits); i++ {
		if swag.IsZero(m.DenomUnits[i]) { // not required
			continue
		}

		if m.DenomUnits[i] != nil {
			if err := m.DenomUnits[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("denom_units" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("denom_units" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this cosmos bank v1beta1 metadata based on the context it is used
func (m *CosmosBankV1beta1Metadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDenomUnits(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosBankV1beta1Metadata) contextValidateDenomUnits(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DenomUnits); i++ {

		if m.DenomUnits[i] != nil {
			if err := m.DenomUnits[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("denom_units" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("denom_units" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CosmosBankV1beta1Metadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosBankV1beta1Metadata) UnmarshalBinary(b []byte) error {
	var res CosmosBankV1beta1Metadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CosmosBankV1beta1MetadataDenomUnitsItems0 DenomUnit represents a struct that describes a given
// denomination unit of the basic token.
//
// swagger:model CosmosBankV1beta1MetadataDenomUnitsItems0
type CosmosBankV1beta1MetadataDenomUnitsItems0 struct {

	// aliases is a list of string aliases for the given denom
	Aliases []string `json:"aliases"`

	// denom represents the string name of the given denom unit (e.g uatom).
	Denom string `json:"denom,omitempty"`

	// exponent represents power of 10 exponent that one must
	// raise the base_denom to in order to equal the given DenomUnit's denom
	// 1 denom = 1^exponent base_denom
	// (e.g. with a base_denom of uatom, one can create a DenomUnit of 'atom' with
	// exponent = 6, thus: 1 atom = 10^6 uatom).
	Exponent int64 `json:"exponent,omitempty"`
}

// Validate validates this cosmos bank v1beta1 metadata denom units items0
func (m *CosmosBankV1beta1MetadataDenomUnitsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cosmos bank v1beta1 metadata denom units items0 based on context it is used
func (m *CosmosBankV1beta1MetadataDenomUnitsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CosmosBankV1beta1MetadataDenomUnitsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosBankV1beta1MetadataDenomUnitsItems0) UnmarshalBinary(b []byte) error {
	var res CosmosBankV1beta1MetadataDenomUnitsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

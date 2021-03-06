// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Rates rates
//
// swagger:model rates
type Rates struct {

	// Coin denomination
	// Required: true
	Denom *string `json:"denom"`

	// one day variation
	// Required: true
	OneDayVariation *string `json:"oneDayVariation"`

	// one day variation rate
	// Required: true
	OneDayVariationRate *string `json:"oneDayVariationRate"`

	// Current swap rate
	// Required: true
	Swaprate *string `json:"swaprate"`
}

// Validate validates this rates
func (m *Rates) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDenom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOneDayVariation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOneDayVariationRate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSwaprate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Rates) validateDenom(formats strfmt.Registry) error {

	if err := validate.Required("denom", "body", m.Denom); err != nil {
		return err
	}

	return nil
}

func (m *Rates) validateOneDayVariation(formats strfmt.Registry) error {

	if err := validate.Required("oneDayVariation", "body", m.OneDayVariation); err != nil {
		return err
	}

	return nil
}

func (m *Rates) validateOneDayVariationRate(formats strfmt.Registry) error {

	if err := validate.Required("oneDayVariationRate", "body", m.OneDayVariationRate); err != nil {
		return err
	}

	return nil
}

func (m *Rates) validateSwaprate(formats strfmt.Registry) error {

	if err := validate.Required("swaprate", "body", m.Swaprate); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this rates based on context it is used
func (m *Rates) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Rates) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Rates) UnmarshalBinary(b []byte) error {
	var res Rates
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

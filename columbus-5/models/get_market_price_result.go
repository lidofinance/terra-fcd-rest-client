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
	"github.com/go-openapi/validate"
)

// GetMarketPriceResult get market price result
//
// swagger:model getMarketPriceResult
type GetMarketPriceResult struct {

	// last price
	// Required: true
	LastPrice *float64 `json:"lastPrice"`

	// one day variation
	// Required: true
	OneDayVariation *string `json:"oneDayVariation"`

	// one day variation rate
	// Required: true
	OneDayVariationRate *string `json:"oneDayVariationRate"`

	// Price history
	// Required: true
	Prices []*GetMarketPriceResultPrices `json:"prices"`
}

// Validate validates this get market price result
func (m *GetMarketPriceResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastPrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOneDayVariation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOneDayVariationRate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetMarketPriceResult) validateLastPrice(formats strfmt.Registry) error {

	if err := validate.Required("lastPrice", "body", m.LastPrice); err != nil {
		return err
	}

	return nil
}

func (m *GetMarketPriceResult) validateOneDayVariation(formats strfmt.Registry) error {

	if err := validate.Required("oneDayVariation", "body", m.OneDayVariation); err != nil {
		return err
	}

	return nil
}

func (m *GetMarketPriceResult) validateOneDayVariationRate(formats strfmt.Registry) error {

	if err := validate.Required("oneDayVariationRate", "body", m.OneDayVariationRate); err != nil {
		return err
	}

	return nil
}

func (m *GetMarketPriceResult) validatePrices(formats strfmt.Registry) error {

	if err := validate.Required("prices", "body", m.Prices); err != nil {
		return err
	}

	for i := 0; i < len(m.Prices); i++ {
		if swag.IsZero(m.Prices[i]) { // not required
			continue
		}

		if m.Prices[i] != nil {
			if err := m.Prices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("prices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("prices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get market price result based on the context it is used
func (m *GetMarketPriceResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePrices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetMarketPriceResult) contextValidatePrices(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Prices); i++ {

		if m.Prices[i] != nil {
			if err := m.Prices[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("prices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("prices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetMarketPriceResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetMarketPriceResult) UnmarshalBinary(b []byte) error {
	var res GetMarketPriceResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

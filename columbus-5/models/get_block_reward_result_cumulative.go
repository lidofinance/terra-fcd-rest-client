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

// GetBlockRewardResultCumulative get block reward result cumulative
//
// swagger:model getBlockRewardResult.cumulative
type GetBlockRewardResultCumulative struct {

	// cumulative reward
	// Required: true
	BlockReward *float64 `json:"blockReward"`

	// unix timestamp
	// Required: true
	Datetime *float64 `json:"datetime"`
}

// Validate validates this get block reward result cumulative
func (m *GetBlockRewardResultCumulative) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBlockReward(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatetime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetBlockRewardResultCumulative) validateBlockReward(formats strfmt.Registry) error {

	if err := validate.Required("blockReward", "body", m.BlockReward); err != nil {
		return err
	}

	return nil
}

func (m *GetBlockRewardResultCumulative) validateDatetime(formats strfmt.Registry) error {

	if err := validate.Required("datetime", "body", m.Datetime); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get block reward result cumulative based on context it is used
func (m *GetBlockRewardResultCumulative) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetBlockRewardResultCumulative) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetBlockRewardResultCumulative) UnmarshalBinary(b []byte) error {
	var res GetBlockRewardResultCumulative
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

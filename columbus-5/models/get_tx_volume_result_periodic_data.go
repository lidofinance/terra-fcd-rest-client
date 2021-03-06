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

// GetTxVolumeResultPeriodicData get tx volume result periodic data
//
// swagger:model getTxVolumeResult.periodic.data
type GetTxVolumeResultPeriodicData struct {

	// unix time
	// Required: true
	Datetime *float64 `json:"datetime"`

	// periodic tx volume
	// Required: true
	TxVolume *string `json:"txVolume"`
}

// Validate validates this get tx volume result periodic data
func (m *GetTxVolumeResultPeriodicData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatetime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTxVolume(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetTxVolumeResultPeriodicData) validateDatetime(formats strfmt.Registry) error {

	if err := validate.Required("datetime", "body", m.Datetime); err != nil {
		return err
	}

	return nil
}

func (m *GetTxVolumeResultPeriodicData) validateTxVolume(formats strfmt.Registry) error {

	if err := validate.Required("txVolume", "body", m.TxVolume); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get tx volume result periodic data based on context it is used
func (m *GetTxVolumeResultPeriodicData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetTxVolumeResultPeriodicData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetTxVolumeResultPeriodicData) UnmarshalBinary(b []byte) error {
	var res GetTxVolumeResultPeriodicData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

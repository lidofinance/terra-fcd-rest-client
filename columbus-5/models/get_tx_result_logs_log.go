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

// GetTxResultLogsLog get tx result logs log
//
// swagger:model getTxResult.logs.log
type GetTxResultLogsLog struct {

	// tax
	// Required: true
	Tax *string `json:"tax"`
}

// Validate validates this get tx result logs log
func (m *GetTxResultLogsLog) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTax(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetTxResultLogsLog) validateTax(formats strfmt.Registry) error {

	if err := validate.Required("tax", "body", m.Tax); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get tx result logs log based on context it is used
func (m *GetTxResultLogsLog) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetTxResultLogsLog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetTxResultLogsLog) UnmarshalBinary(b []byte) error {
	var res GetTxResultLogsLog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

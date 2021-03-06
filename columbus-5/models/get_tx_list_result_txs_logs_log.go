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

// GetTxListResultTxsLogsLog get tx list result txs logs log
//
// swagger:model getTxListResult.txs.logs.log
type GetTxListResultTxsLogsLog struct {

	// tax
	// Required: true
	Tax *string `json:"tax"`
}

// Validate validates this get tx list result txs logs log
func (m *GetTxListResultTxsLogsLog) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTax(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetTxListResultTxsLogsLog) validateTax(formats strfmt.Registry) error {

	if err := validate.Required("tax", "body", m.Tax); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get tx list result txs logs log based on context it is used
func (m *GetTxListResultTxsLogsLog) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetTxListResultTxsLogsLog) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetTxListResultTxsLogsLog) UnmarshalBinary(b []byte) error {
	var res GetTxListResultTxsLogsLog
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

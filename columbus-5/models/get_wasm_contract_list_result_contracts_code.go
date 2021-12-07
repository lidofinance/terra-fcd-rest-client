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

// GetWasmContractListResultContractsCode get wasm contract list result contracts code
//
// swagger:model getWasmContractListResult.contracts.code
type GetWasmContractListResultContractsCode struct {

	// sent code id
	// Required: true
	CodeID *string `json:"code_id"`

	// code info
	// Required: true
	Info *GetWasmContractListResultContractsCodeInfo `json:"info"`

	// sender
	// Required: true
	Sender *string `json:"sender"`

	// timestamp
	// Required: true
	Timestamp *string `json:"timestamp"`

	// txhash
	// Required: true
	Txhash *string `json:"txhash"`
}

// Validate validates this get wasm contract list result contracts code
func (m *GetWasmContractListResultContractsCode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCodeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSender(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTxhash(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetWasmContractListResultContractsCode) validateCodeID(formats strfmt.Registry) error {

	if err := validate.Required("code_id", "body", m.CodeID); err != nil {
		return err
	}

	return nil
}

func (m *GetWasmContractListResultContractsCode) validateInfo(formats strfmt.Registry) error {

	if err := validate.Required("info", "body", m.Info); err != nil {
		return err
	}

	if m.Info != nil {
		if err := m.Info.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("info")
			}
			return err
		}
	}

	return nil
}

func (m *GetWasmContractListResultContractsCode) validateSender(formats strfmt.Registry) error {

	if err := validate.Required("sender", "body", m.Sender); err != nil {
		return err
	}

	return nil
}

func (m *GetWasmContractListResultContractsCode) validateTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	return nil
}

func (m *GetWasmContractListResultContractsCode) validateTxhash(formats strfmt.Registry) error {

	if err := validate.Required("txhash", "body", m.Txhash); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get wasm contract list result contracts code based on the context it is used
func (m *GetWasmContractListResultContractsCode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetWasmContractListResultContractsCode) contextValidateInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.Info != nil {
		if err := m.Info.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("info")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetWasmContractListResultContractsCode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetWasmContractListResultContractsCode) UnmarshalBinary(b []byte) error {
	var res GetWasmContractListResultContractsCode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

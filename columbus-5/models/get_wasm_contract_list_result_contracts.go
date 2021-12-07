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

// GetWasmContractListResultContracts get wasm contract list result contracts
//
// swagger:model getWasmContractListResult.contracts
type GetWasmContractListResultContracts struct {

	// code details info
	// Required: true
	Code *GetWasmContractListResultContractsCode `json:"code"`

	// sent code id
	// Required: true
	CodeID *string `json:"code_id"`

	// contract address
	// Required: true
	ContractAddress *string `json:"contractAddress"`

	// code info
	// Required: true
	Info *GetWasmContractListResultContractsInfo `json:"info"`

	// contract initialization message
	// Required: true
	InitMsg *string `json:"init_msg"`

	// contract migrate message
	// Required: true
	MigrateMsg *string `json:"migrate_msg"`

	// owner
	// Required: true
	Owner *string `json:"owner"`

	// timestamp
	// Required: true
	Timestamp *string `json:"timestamp"`

	// txhash
	// Required: true
	Txhash *string `json:"txhash"`
}

// Validate validates this get wasm contract list result contracts
func (m *GetWasmContractListResultContracts) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCodeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContractAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitMsg(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMigrateMsg(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
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

func (m *GetWasmContractListResultContracts) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	if m.Code != nil {
		if err := m.Code.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("code")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("code")
			}
			return err
		}
	}

	return nil
}

func (m *GetWasmContractListResultContracts) validateCodeID(formats strfmt.Registry) error {

	if err := validate.Required("code_id", "body", m.CodeID); err != nil {
		return err
	}

	return nil
}

func (m *GetWasmContractListResultContracts) validateContractAddress(formats strfmt.Registry) error {

	if err := validate.Required("contractAddress", "body", m.ContractAddress); err != nil {
		return err
	}

	return nil
}

func (m *GetWasmContractListResultContracts) validateInfo(formats strfmt.Registry) error {

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

func (m *GetWasmContractListResultContracts) validateInitMsg(formats strfmt.Registry) error {

	if err := validate.Required("init_msg", "body", m.InitMsg); err != nil {
		return err
	}

	return nil
}

func (m *GetWasmContractListResultContracts) validateMigrateMsg(formats strfmt.Registry) error {

	if err := validate.Required("migrate_msg", "body", m.MigrateMsg); err != nil {
		return err
	}

	return nil
}

func (m *GetWasmContractListResultContracts) validateOwner(formats strfmt.Registry) error {

	if err := validate.Required("owner", "body", m.Owner); err != nil {
		return err
	}

	return nil
}

func (m *GetWasmContractListResultContracts) validateTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("timestamp", "body", m.Timestamp); err != nil {
		return err
	}

	return nil
}

func (m *GetWasmContractListResultContracts) validateTxhash(formats strfmt.Registry) error {

	if err := validate.Required("txhash", "body", m.Txhash); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get wasm contract list result contracts based on the context it is used
func (m *GetWasmContractListResultContracts) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetWasmContractListResultContracts) contextValidateCode(ctx context.Context, formats strfmt.Registry) error {

	if m.Code != nil {
		if err := m.Code.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("code")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("code")
			}
			return err
		}
	}

	return nil
}

func (m *GetWasmContractListResultContracts) contextValidateInfo(ctx context.Context, formats strfmt.Registry) error {

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
func (m *GetWasmContractListResultContracts) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetWasmContractListResultContracts) UnmarshalBinary(b []byte) error {
	var res GetWasmContractListResultContracts
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

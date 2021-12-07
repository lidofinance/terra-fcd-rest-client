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

// MigrateCodeReq migrate code req
//
// swagger:model MigrateCodeReq
type MigrateCodeReq struct {

	// base req
	BaseReq *MigrateCodeReqBaseReq `json:"base_req,omitempty"`

	// wasm bytes
	// Example: Avz04VhtKJh8ACCVzlI8aTosGy0ikFXKIVHQ3jKMrosH
	WasmBytes string `json:"wasm_bytes,omitempty"`
}

// Validate validates this migrate code req
func (m *MigrateCodeReq) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBaseReq(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MigrateCodeReq) validateBaseReq(formats strfmt.Registry) error {
	if swag.IsZero(m.BaseReq) { // not required
		return nil
	}

	if m.BaseReq != nil {
		if err := m.BaseReq.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("base_req")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("base_req")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this migrate code req based on the context it is used
func (m *MigrateCodeReq) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBaseReq(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MigrateCodeReq) contextValidateBaseReq(ctx context.Context, formats strfmt.Registry) error {

	if m.BaseReq != nil {
		if err := m.BaseReq.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("base_req")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("base_req")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MigrateCodeReq) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MigrateCodeReq) UnmarshalBinary(b []byte) error {
	var res MigrateCodeReq
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MigrateCodeReqBaseReq migrate code req base req
//
// swagger:model MigrateCodeReqBaseReq
type MigrateCodeReqBaseReq struct {

	// account number
	// Example: 0
	AccountNumber string `json:"account_number,omitempty"`

	// chain id
	// Example: Columbus-5
	ChainID string `json:"chain_id,omitempty"`

	// fees
	Fees []*MigrateCodeReqBaseReqFeesItems0 `json:"fees"`

	// Sender address or Keybase name to generate a transaction
	// Example: terra1wg2mlrxdmnnkkykgqg4znky86nyrtc45q336yv
	From string `json:"from,omitempty"`

	// gas
	// Example: 200000
	Gas string `json:"gas,omitempty"`

	// gas adjustment
	// Example: 1.2
	GasAdjustment string `json:"gas_adjustment,omitempty"`

	// memo
	// Example: Sent via Terra Station 🚀
	Memo string `json:"memo,omitempty"`

	// sequence
	// Example: 1
	Sequence string `json:"sequence,omitempty"`

	// Estimate gas for a transaction (cannot be used in conjunction with generate_only)
	// Example: false
	Simulate bool `json:"simulate,omitempty"`
}

// Validate validates this migrate code req base req
func (m *MigrateCodeReqBaseReq) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFees(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MigrateCodeReqBaseReq) validateFees(formats strfmt.Registry) error {
	if swag.IsZero(m.Fees) { // not required
		return nil
	}

	for i := 0; i < len(m.Fees); i++ {
		if swag.IsZero(m.Fees[i]) { // not required
			continue
		}

		if m.Fees[i] != nil {
			if err := m.Fees[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("base_req" + "." + "fees" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("base_req" + "." + "fees" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this migrate code req base req based on the context it is used
func (m *MigrateCodeReqBaseReq) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFees(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MigrateCodeReqBaseReq) contextValidateFees(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Fees); i++ {

		if m.Fees[i] != nil {
			if err := m.Fees[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("base_req" + "." + "fees" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("base_req" + "." + "fees" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MigrateCodeReqBaseReq) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MigrateCodeReqBaseReq) UnmarshalBinary(b []byte) error {
	var res MigrateCodeReqBaseReq
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MigrateCodeReqBaseReqFeesItems0 migrate code req base req fees items0
//
// swagger:model MigrateCodeReqBaseReqFeesItems0
type MigrateCodeReqBaseReqFeesItems0 struct {

	// amount
	// Example: 50
	Amount string `json:"amount,omitempty"`

	// denom
	// Example: uluna
	Denom string `json:"denom,omitempty"`
}

// Validate validates this migrate code req base req fees items0
func (m *MigrateCodeReqBaseReqFeesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this migrate code req base req fees items0 based on context it is used
func (m *MigrateCodeReqBaseReqFeesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MigrateCodeReqBaseReqFeesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MigrateCodeReqBaseReqFeesItems0) UnmarshalBinary(b []byte) error {
	var res MigrateCodeReqBaseReqFeesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

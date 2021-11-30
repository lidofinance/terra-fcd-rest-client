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

// AggregatePrevoteReq aggregate prevote req
//
// swagger:model AggregatePrevoteReq
type AggregatePrevoteReq struct {

	// base req
	BaseReq *AggregatePrevoteReqBaseReq `json:"base_req,omitempty"`

	// exchange rates of Luna in denom currencies are to make aggregate prevote hash; this field is required to submit prevote in case absence of hash
	// Example: 1000.0ukrw,1.2uusd,0.99usdr
	ExchangeRates string `json:"exchange_rates,omitempty"`

	// hex string; hash of next vote; empty == skip prevote
	// Example: 061bf1e27dfff121f40c826e593c8a28ec299a02
	Hash string `json:"hash,omitempty"`

	// salt is to make prevote hash; this field is required to submit prevote in case  absence of hash
	// Example: abcd
	Salt string `json:"salt,omitempty"`
}

// Validate validates this aggregate prevote req
func (m *AggregatePrevoteReq) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBaseReq(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AggregatePrevoteReq) validateBaseReq(formats strfmt.Registry) error {
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

// ContextValidate validate this aggregate prevote req based on the context it is used
func (m *AggregatePrevoteReq) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBaseReq(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AggregatePrevoteReq) contextValidateBaseReq(ctx context.Context, formats strfmt.Registry) error {

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
func (m *AggregatePrevoteReq) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AggregatePrevoteReq) UnmarshalBinary(b []byte) error {
	var res AggregatePrevoteReq
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AggregatePrevoteReqBaseReq aggregate prevote req base req
//
// swagger:model AggregatePrevoteReqBaseReq
type AggregatePrevoteReqBaseReq struct {

	// account number
	// Example: 0
	AccountNumber string `json:"account_number,omitempty"`

	// chain id
	// Example: Columbus-5
	ChainID string `json:"chain_id,omitempty"`

	// fees
	Fees []*AggregatePrevoteReqBaseReqFeesItems0 `json:"fees"`

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

// Validate validates this aggregate prevote req base req
func (m *AggregatePrevoteReqBaseReq) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFees(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AggregatePrevoteReqBaseReq) validateFees(formats strfmt.Registry) error {
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

// ContextValidate validate this aggregate prevote req base req based on the context it is used
func (m *AggregatePrevoteReqBaseReq) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFees(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AggregatePrevoteReqBaseReq) contextValidateFees(ctx context.Context, formats strfmt.Registry) error {

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
func (m *AggregatePrevoteReqBaseReq) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AggregatePrevoteReqBaseReq) UnmarshalBinary(b []byte) error {
	var res AggregatePrevoteReqBaseReq
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AggregatePrevoteReqBaseReqFeesItems0 aggregate prevote req base req fees items0
//
// swagger:model AggregatePrevoteReqBaseReqFeesItems0
type AggregatePrevoteReqBaseReqFeesItems0 struct {

	// amount
	// Example: 50
	Amount string `json:"amount,omitempty"`

	// denom
	// Example: uluna
	Denom string `json:"denom,omitempty"`
}

// Validate validates this aggregate prevote req base req fees items0
func (m *AggregatePrevoteReqBaseReqFeesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this aggregate prevote req base req fees items0 based on context it is used
func (m *AggregatePrevoteReqBaseReqFeesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AggregatePrevoteReqBaseReqFeesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AggregatePrevoteReqBaseReqFeesItems0) UnmarshalBinary(b []byte) error {
	var res AggregatePrevoteReqBaseReqFeesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
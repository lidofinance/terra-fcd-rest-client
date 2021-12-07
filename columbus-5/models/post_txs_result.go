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

// PostTxsResult post txs result
//
// swagger:model postTxsResult
type PostTxsResult struct {

	// tx info
	// Required: true
	CheckTx *PostTxsResultCheckTx `json:"check_tx"`

	// tx info
	// Required: true
	DeliverTx *PostTxsResultDeliverTx `json:"deliver_tx"`

	// Tx hash
	// Required: true
	Hash *string `json:"hash"`

	// Block height
	// Required: true
	Height *float64 `json:"height"`
}

// Validate validates this post txs result
func (m *PostTxsResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCheckTx(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeliverTx(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHeight(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostTxsResult) validateCheckTx(formats strfmt.Registry) error {

	if err := validate.Required("check_tx", "body", m.CheckTx); err != nil {
		return err
	}

	if m.CheckTx != nil {
		if err := m.CheckTx.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("check_tx")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("check_tx")
			}
			return err
		}
	}

	return nil
}

func (m *PostTxsResult) validateDeliverTx(formats strfmt.Registry) error {

	if err := validate.Required("deliver_tx", "body", m.DeliverTx); err != nil {
		return err
	}

	if m.DeliverTx != nil {
		if err := m.DeliverTx.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deliver_tx")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deliver_tx")
			}
			return err
		}
	}

	return nil
}

func (m *PostTxsResult) validateHash(formats strfmt.Registry) error {

	if err := validate.Required("hash", "body", m.Hash); err != nil {
		return err
	}

	return nil
}

func (m *PostTxsResult) validateHeight(formats strfmt.Registry) error {

	if err := validate.Required("height", "body", m.Height); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this post txs result based on the context it is used
func (m *PostTxsResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCheckTx(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeliverTx(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostTxsResult) contextValidateCheckTx(ctx context.Context, formats strfmt.Registry) error {

	if m.CheckTx != nil {
		if err := m.CheckTx.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("check_tx")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("check_tx")
			}
			return err
		}
	}

	return nil
}

func (m *PostTxsResult) contextValidateDeliverTx(ctx context.Context, formats strfmt.Registry) error {

	if m.DeliverTx != nil {
		if err := m.DeliverTx.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deliver_tx")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deliver_tx")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostTxsResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostTxsResult) UnmarshalBinary(b []byte) error {
	var res PostTxsResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

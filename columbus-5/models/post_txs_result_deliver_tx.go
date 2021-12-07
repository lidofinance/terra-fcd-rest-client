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

// PostTxsResultDeliverTx post txs result deliver tx
//
// swagger:model postTxsResult.deliver_tx
type PostTxsResultDeliverTx struct {

	// code
	// Required: true
	Code *float64 `json:"code"`

	// data
	// Required: true
	Data *string `json:"data"`

	// gas used
	// Required: true
	GasUsed *float64 `json:"gas_used"`

	// gas wanted
	// Required: true
	GasWanted *float64 `json:"gas_wanted"`

	// info
	// Required: true
	Info *string `json:"info"`

	// log
	// Required: true
	Log *string `json:"log"`

	// tags
	// Required: true
	Tags []string `json:"tags"`
}

// Validate validates this post txs result deliver tx
func (m *PostTxsResultDeliverTx) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGasUsed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGasWanted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLog(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostTxsResultDeliverTx) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *PostTxsResultDeliverTx) validateData(formats strfmt.Registry) error {

	if err := validate.Required("data", "body", m.Data); err != nil {
		return err
	}

	return nil
}

func (m *PostTxsResultDeliverTx) validateGasUsed(formats strfmt.Registry) error {

	if err := validate.Required("gas_used", "body", m.GasUsed); err != nil {
		return err
	}

	return nil
}

func (m *PostTxsResultDeliverTx) validateGasWanted(formats strfmt.Registry) error {

	if err := validate.Required("gas_wanted", "body", m.GasWanted); err != nil {
		return err
	}

	return nil
}

func (m *PostTxsResultDeliverTx) validateInfo(formats strfmt.Registry) error {

	if err := validate.Required("info", "body", m.Info); err != nil {
		return err
	}

	return nil
}

func (m *PostTxsResultDeliverTx) validateLog(formats strfmt.Registry) error {

	if err := validate.Required("log", "body", m.Log); err != nil {
		return err
	}

	return nil
}

func (m *PostTxsResultDeliverTx) validateTags(formats strfmt.Registry) error {

	if err := validate.Required("tags", "body", m.Tags); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post txs result deliver tx based on context it is used
func (m *PostTxsResultDeliverTx) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostTxsResultDeliverTx) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostTxsResultDeliverTx) UnmarshalBinary(b []byte) error {
	var res PostTxsResultDeliverTx
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
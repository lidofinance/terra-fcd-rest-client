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
	"github.com/go-openapi/validate"
)

// GetWasmContractListResult get wasm contract list result
//
// swagger:model getWasmContractListResult
type GetWasmContractListResult struct {

	// contracts info
	// Required: true
	Contracts []*GetWasmContractListResultContracts `json:"contracts"`

	// Per page item limit
	// Required: true
	Limit *float64 `json:"limit"`
}

// Validate validates this get wasm contract list result
func (m *GetWasmContractListResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContracts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLimit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetWasmContractListResult) validateContracts(formats strfmt.Registry) error {

	if err := validate.Required("contracts", "body", m.Contracts); err != nil {
		return err
	}

	for i := 0; i < len(m.Contracts); i++ {
		if swag.IsZero(m.Contracts[i]) { // not required
			continue
		}

		if m.Contracts[i] != nil {
			if err := m.Contracts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("contracts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("contracts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GetWasmContractListResult) validateLimit(formats strfmt.Registry) error {

	if err := validate.Required("limit", "body", m.Limit); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get wasm contract list result based on the context it is used
func (m *GetWasmContractListResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContracts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetWasmContractListResult) contextValidateContracts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Contracts); i++ {

		if m.Contracts[i] != nil {
			if err := m.Contracts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("contracts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("contracts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetWasmContractListResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetWasmContractListResult) UnmarshalBinary(b []byte) error {
	var res GetWasmContractListResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

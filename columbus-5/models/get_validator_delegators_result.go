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

// GetValidatorDelegatorsResult get validator delegators result
//
// swagger:model getValidatorDelegatorsResult
type GetValidatorDelegatorsResult struct {

	// Delegator list
	// Required: true
	Delegator []*GetValidatorDelegatorsResultDelegator `json:"delegator"`

	// limit
	// Required: true
	Limit *float64 `json:"limit"`

	// page
	// Required: true
	Page *float64 `json:"page"`
}

// Validate validates this get validator delegators result
func (m *GetValidatorDelegatorsResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDelegator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetValidatorDelegatorsResult) validateDelegator(formats strfmt.Registry) error {

	if err := validate.Required("delegator", "body", m.Delegator); err != nil {
		return err
	}

	for i := 0; i < len(m.Delegator); i++ {
		if swag.IsZero(m.Delegator[i]) { // not required
			continue
		}

		if m.Delegator[i] != nil {
			if err := m.Delegator[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("delegator" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("delegator" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GetValidatorDelegatorsResult) validateLimit(formats strfmt.Registry) error {

	if err := validate.Required("limit", "body", m.Limit); err != nil {
		return err
	}

	return nil
}

func (m *GetValidatorDelegatorsResult) validatePage(formats strfmt.Registry) error {

	if err := validate.Required("page", "body", m.Page); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get validator delegators result based on the context it is used
func (m *GetValidatorDelegatorsResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDelegator(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetValidatorDelegatorsResult) contextValidateDelegator(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Delegator); i++ {

		if m.Delegator[i] != nil {
			if err := m.Delegator[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("delegator" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("delegator" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetValidatorDelegatorsResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetValidatorDelegatorsResult) UnmarshalBinary(b []byte) error {
	var res GetValidatorDelegatorsResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

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

// GetValidatorClaimsResult get validator claims result
//
// swagger:model getValidatorClaimsResult
type GetValidatorClaimsResult struct {

	// Claim list
	// Required: true
	Claims []*GetValidatorClaimsResultClaims `json:"claims"`

	// limit
	// Required: true
	Limit *float64 `json:"limit"`

	// page
	// Required: true
	Page *float64 `json:"page"`
}

// Validate validates this get validator claims result
func (m *GetValidatorClaimsResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClaims(formats); err != nil {
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

func (m *GetValidatorClaimsResult) validateClaims(formats strfmt.Registry) error {

	if err := validate.Required("claims", "body", m.Claims); err != nil {
		return err
	}

	for i := 0; i < len(m.Claims); i++ {
		if swag.IsZero(m.Claims[i]) { // not required
			continue
		}

		if m.Claims[i] != nil {
			if err := m.Claims[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("claims" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("claims" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GetValidatorClaimsResult) validateLimit(formats strfmt.Registry) error {

	if err := validate.Required("limit", "body", m.Limit); err != nil {
		return err
	}

	return nil
}

func (m *GetValidatorClaimsResult) validatePage(formats strfmt.Registry) error {

	if err := validate.Required("page", "body", m.Page); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get validator claims result based on the context it is used
func (m *GetValidatorClaimsResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClaims(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetValidatorClaimsResult) contextValidateClaims(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Claims); i++ {

		if m.Claims[i] != nil {
			if err := m.Claims[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("claims" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("claims" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetValidatorClaimsResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetValidatorClaimsResult) UnmarshalBinary(b []byte) error {
	var res GetValidatorClaimsResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

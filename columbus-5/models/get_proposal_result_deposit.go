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

// GetProposalResultDeposit get proposal result deposit
//
// swagger:model getProposalResult.deposit
type GetProposalResultDeposit struct {

	// deposit end time
	// Required: true
	DepositEndTime *string `json:"depositEndTime"`

	// Minimum deposit
	// Required: true
	MinDeposit []*GetProposalResultDepositMinDeposit `json:"minDeposit"`

	// total deposit info
	// Required: true
	TotalDeposit []*GetProposalResultDepositTotalDeposit `json:"totalDeposit"`
}

// Validate validates this get proposal result deposit
func (m *GetProposalResultDeposit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDepositEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinDeposit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalDeposit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetProposalResultDeposit) validateDepositEndTime(formats strfmt.Registry) error {

	if err := validate.Required("depositEndTime", "body", m.DepositEndTime); err != nil {
		return err
	}

	return nil
}

func (m *GetProposalResultDeposit) validateMinDeposit(formats strfmt.Registry) error {

	if err := validate.Required("minDeposit", "body", m.MinDeposit); err != nil {
		return err
	}

	for i := 0; i < len(m.MinDeposit); i++ {
		if swag.IsZero(m.MinDeposit[i]) { // not required
			continue
		}

		if m.MinDeposit[i] != nil {
			if err := m.MinDeposit[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("minDeposit" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("minDeposit" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GetProposalResultDeposit) validateTotalDeposit(formats strfmt.Registry) error {

	if err := validate.Required("totalDeposit", "body", m.TotalDeposit); err != nil {
		return err
	}

	for i := 0; i < len(m.TotalDeposit); i++ {
		if swag.IsZero(m.TotalDeposit[i]) { // not required
			continue
		}

		if m.TotalDeposit[i] != nil {
			if err := m.TotalDeposit[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("totalDeposit" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("totalDeposit" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get proposal result deposit based on the context it is used
func (m *GetProposalResultDeposit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMinDeposit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTotalDeposit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetProposalResultDeposit) contextValidateMinDeposit(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MinDeposit); i++ {

		if m.MinDeposit[i] != nil {
			if err := m.MinDeposit[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("minDeposit" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("minDeposit" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GetProposalResultDeposit) contextValidateTotalDeposit(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TotalDeposit); i++ {

		if m.TotalDeposit[i] != nil {
			if err := m.TotalDeposit[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("totalDeposit" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("totalDeposit" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetProposalResultDeposit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetProposalResultDeposit) UnmarshalBinary(b []byte) error {
	var res GetProposalResultDeposit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

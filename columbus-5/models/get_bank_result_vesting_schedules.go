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

// GetBankResultVestingSchedules get bank result vesting schedules
//
// swagger:model getBankResult.vesting.schedules
type GetBankResultVestingSchedules struct {

	// vesting amount
	// Required: true
	Amount *string `json:"amount"`

	// vesting end time
	// Required: true
	EndTime *string `json:"endTime"`

	// vesting ratio
	// Required: true
	Ratio *string `json:"ratio"`

	// vestring start time
	// Required: true
	StartTime *string `json:"startTime"`
}

// Validate validates this get bank result vesting schedules
func (m *GetBankResultVestingSchedules) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRatio(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetBankResultVestingSchedules) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *GetBankResultVestingSchedules) validateEndTime(formats strfmt.Registry) error {

	if err := validate.Required("endTime", "body", m.EndTime); err != nil {
		return err
	}

	return nil
}

func (m *GetBankResultVestingSchedules) validateRatio(formats strfmt.Registry) error {

	if err := validate.Required("ratio", "body", m.Ratio); err != nil {
		return err
	}

	return nil
}

func (m *GetBankResultVestingSchedules) validateStartTime(formats strfmt.Registry) error {

	if err := validate.Required("startTime", "body", m.StartTime); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get bank result vesting schedules based on context it is used
func (m *GetBankResultVestingSchedules) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetBankResultVestingSchedules) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetBankResultVestingSchedules) UnmarshalBinary(b []byte) error {
	var res GetBankResultVestingSchedules
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

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

// GetValidatorDetailResultMyUndelegation get validator detail result my undelegation
//
// swagger:model getValidatorDetailResult.myUndelegation
type GetValidatorDetailResultMyUndelegation struct {

	// undelegation amount
	// Required: true
	Amount *string `json:"amount"`

	// block height
	// Required: true
	CreationHeight *string `json:"creationHeight"`

	// undelegation release date time
	// Required: true
	ReleaseTime *string `json:"releaseTime"`

	// validator address
	// Required: true
	ValidatorAddress *string `json:"validatorAddress"`

	// validator name
	// Required: true
	ValidatorName *string `json:"validatorName"`
}

// Validate validates this get validator detail result my undelegation
func (m *GetValidatorDetailResultMyUndelegation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationHeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReleaseTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidatorAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidatorName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetValidatorDetailResultMyUndelegation) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *GetValidatorDetailResultMyUndelegation) validateCreationHeight(formats strfmt.Registry) error {

	if err := validate.Required("creationHeight", "body", m.CreationHeight); err != nil {
		return err
	}

	return nil
}

func (m *GetValidatorDetailResultMyUndelegation) validateReleaseTime(formats strfmt.Registry) error {

	if err := validate.Required("releaseTime", "body", m.ReleaseTime); err != nil {
		return err
	}

	return nil
}

func (m *GetValidatorDetailResultMyUndelegation) validateValidatorAddress(formats strfmt.Registry) error {

	if err := validate.Required("validatorAddress", "body", m.ValidatorAddress); err != nil {
		return err
	}

	return nil
}

func (m *GetValidatorDetailResultMyUndelegation) validateValidatorName(formats strfmt.Registry) error {

	if err := validate.Required("validatorName", "body", m.ValidatorName); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get validator detail result my undelegation based on context it is used
func (m *GetValidatorDetailResultMyUndelegation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetValidatorDetailResultMyUndelegation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetValidatorDetailResultMyUndelegation) UnmarshalBinary(b []byte) error {
	var res GetValidatorDetailResultMyUndelegation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

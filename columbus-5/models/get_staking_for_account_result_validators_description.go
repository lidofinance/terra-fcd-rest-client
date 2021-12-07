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

// GetStakingForAccountResultValidatorsDescription get staking for account result validators description
//
// swagger:model getStakingForAccountResult.validators.description
type GetStakingForAccountResultValidatorsDescription struct {

	// details
	// Required: true
	Details *string `json:"details"`

	// identity
	// Required: true
	Identity *string `json:"identity"`

	// moniker
	// Required: true
	Moniker *string `json:"moniker"`

	// profile icon
	// Required: true
	ProfileIcon *string `json:"profileIcon"`

	// website
	// Required: true
	Website *string `json:"website"`
}

// Validate validates this get staking for account result validators description
func (m *GetStakingForAccountResultValidatorsDescription) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMoniker(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProfileIcon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWebsite(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetStakingForAccountResultValidatorsDescription) validateDetails(formats strfmt.Registry) error {

	if err := validate.Required("details", "body", m.Details); err != nil {
		return err
	}

	return nil
}

func (m *GetStakingForAccountResultValidatorsDescription) validateIdentity(formats strfmt.Registry) error {

	if err := validate.Required("identity", "body", m.Identity); err != nil {
		return err
	}

	return nil
}

func (m *GetStakingForAccountResultValidatorsDescription) validateMoniker(formats strfmt.Registry) error {

	if err := validate.Required("moniker", "body", m.Moniker); err != nil {
		return err
	}

	return nil
}

func (m *GetStakingForAccountResultValidatorsDescription) validateProfileIcon(formats strfmt.Registry) error {

	if err := validate.Required("profileIcon", "body", m.ProfileIcon); err != nil {
		return err
	}

	return nil
}

func (m *GetStakingForAccountResultValidatorsDescription) validateWebsite(formats strfmt.Registry) error {

	if err := validate.Required("website", "body", m.Website); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get staking for account result validators description based on context it is used
func (m *GetStakingForAccountResultValidatorsDescription) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetStakingForAccountResultValidatorsDescription) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetStakingForAccountResultValidatorsDescription) UnmarshalBinary(b []byte) error {
	var res GetStakingForAccountResultValidatorsDescription
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

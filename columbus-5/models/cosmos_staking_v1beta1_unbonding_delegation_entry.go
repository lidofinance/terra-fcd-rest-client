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

// CosmosStakingV1beta1UnbondingDelegationEntry UnbondingDelegationEntry defines an unbonding object with relevant metadata.
//
// swagger:model cosmos.staking.v1beta1.UnbondingDelegationEntry
type CosmosStakingV1beta1UnbondingDelegationEntry struct {

	// balance defines the tokens to receive at completion.
	Balance string `json:"balance,omitempty"`

	// completion_time is the unix time for unbonding completion.
	// Format: date-time
	CompletionTime strfmt.DateTime `json:"completion_time,omitempty"`

	// creation_height is the height which the unbonding took place.
	CreationHeight string `json:"creation_height,omitempty"`

	// initial_balance defines the tokens initially scheduled to receive at completion.
	InitialBalance string `json:"initial_balance,omitempty"`
}

// Validate validates this cosmos staking v1beta1 unbonding delegation entry
func (m *CosmosStakingV1beta1UnbondingDelegationEntry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompletionTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosStakingV1beta1UnbondingDelegationEntry) validateCompletionTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CompletionTime) { // not required
		return nil
	}

	if err := validate.FormatOf("completion_time", "body", "date-time", m.CompletionTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this cosmos staking v1beta1 unbonding delegation entry based on context it is used
func (m *CosmosStakingV1beta1UnbondingDelegationEntry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CosmosStakingV1beta1UnbondingDelegationEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosStakingV1beta1UnbondingDelegationEntry) UnmarshalBinary(b []byte) error {
	var res CosmosStakingV1beta1UnbondingDelegationEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

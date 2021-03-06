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

// CosmosStakingV1beta1RedelegationEntry RedelegationEntry defines a redelegation object with relevant metadata.
//
// swagger:model cosmos.staking.v1beta1.RedelegationEntry
type CosmosStakingV1beta1RedelegationEntry struct {

	// completion_time defines the unix time for redelegation completion.
	// Format: date-time
	CompletionTime strfmt.DateTime `json:"completion_time,omitempty"`

	// creation_height  defines the height which the redelegation took place.
	CreationHeight string `json:"creation_height,omitempty"`

	// initial_balance defines the initial balance when redelegation started.
	InitialBalance string `json:"initial_balance,omitempty"`

	// shares_dst is the amount of destination-validator shares created by redelegation.
	SharesDst string `json:"shares_dst,omitempty"`
}

// Validate validates this cosmos staking v1beta1 redelegation entry
func (m *CosmosStakingV1beta1RedelegationEntry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompletionTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosStakingV1beta1RedelegationEntry) validateCompletionTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CompletionTime) { // not required
		return nil
	}

	if err := validate.FormatOf("completion_time", "body", "date-time", m.CompletionTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this cosmos staking v1beta1 redelegation entry based on context it is used
func (m *CosmosStakingV1beta1RedelegationEntry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CosmosStakingV1beta1RedelegationEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosStakingV1beta1RedelegationEntry) UnmarshalBinary(b []byte) error {
	var res CosmosStakingV1beta1RedelegationEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

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
)

// DelegationDelegatorReward delegation delegator reward
//
// swagger:model DelegationDelegatorReward
type DelegationDelegatorReward struct {

	// reward
	Reward []*DelegationDelegatorRewardRewardItems0 `json:"reward"`

	// bech32 encoded address
	// Example: terravaloper1wg2mlrxdmnnkkykgqg4znky86nyrtc45q7a85l
	ValidatorAddress string `json:"validator_address,omitempty"`
}

// Validate validates this delegation delegator reward
func (m *DelegationDelegatorReward) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReward(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DelegationDelegatorReward) validateReward(formats strfmt.Registry) error {
	if swag.IsZero(m.Reward) { // not required
		return nil
	}

	for i := 0; i < len(m.Reward); i++ {
		if swag.IsZero(m.Reward[i]) { // not required
			continue
		}

		if m.Reward[i] != nil {
			if err := m.Reward[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("reward" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("reward" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this delegation delegator reward based on the context it is used
func (m *DelegationDelegatorReward) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateReward(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DelegationDelegatorReward) contextValidateReward(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Reward); i++ {

		if m.Reward[i] != nil {
			if err := m.Reward[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("reward" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("reward" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DelegationDelegatorReward) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DelegationDelegatorReward) UnmarshalBinary(b []byte) error {
	var res DelegationDelegatorReward
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// DelegationDelegatorRewardRewardItems0 delegation delegator reward reward items0
//
// swagger:model DelegationDelegatorRewardRewardItems0
type DelegationDelegatorRewardRewardItems0 struct {

	// amount
	// Example: 50
	Amount string `json:"amount,omitempty"`

	// denom
	// Example: uluna
	Denom string `json:"denom,omitempty"`
}

// Validate validates this delegation delegator reward reward items0
func (m *DelegationDelegatorRewardRewardItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delegation delegator reward reward items0 based on context it is used
func (m *DelegationDelegatorRewardRewardItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DelegationDelegatorRewardRewardItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DelegationDelegatorRewardRewardItems0) UnmarshalBinary(b []byte) error {
	var res DelegationDelegatorRewardRewardItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

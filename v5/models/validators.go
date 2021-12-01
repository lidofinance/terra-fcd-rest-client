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

// Validators validators
//
// swagger:model validators
type Validators struct {

	// commission info
	// Required: true
	CommissionInfo *ValidatorsCommissionInfo `json:"commissionInfo"`

	// consensus pubkey
	// Required: true
	ConsensusPubkey *string `json:"consensusPubkey"`

	// delegator shares
	// Required: true
	DelegatorShares *string `json:"delegatorShares"`

	// description
	// Required: true
	Description *ValidatorsDescription `json:"description"`

	// The amount of user delegation to this validator
	// Required: true
	MyDelegation *string `json:"myDelegation"`

	// Undelegation information of user in progress in this validator
	// Required: true
	MyUndelegation *string `json:"myUndelegation"`

	// operator address
	// Required: true
	OperatorAddress *string `json:"operatorAddress"`

	// rewards pool
	// Required: true
	RewardsPool *ValidatorsRewardsPool `json:"rewardsPool"`

	// staking return
	// Required: true
	StakingReturn *string `json:"stakingReturn"`

	// status
	// Required: true
	Status *string `json:"status"`

	// tokens
	// Required: true
	Tokens *string `json:"tokens"`

	// up time
	// Required: true
	UpTime *float64 `json:"upTime"`

	// voting power
	// Required: true
	VotingPower *ValidatorsVotingPower `json:"votingPower"`
}

// Validate validates this validators
func (m *Validators) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommissionInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConsensusPubkey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDelegatorShares(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMyDelegation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMyUndelegation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperatorAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRewardsPool(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStakingReturn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTokens(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVotingPower(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Validators) validateCommissionInfo(formats strfmt.Registry) error {

	if err := validate.Required("commissionInfo", "body", m.CommissionInfo); err != nil {
		return err
	}

	if m.CommissionInfo != nil {
		if err := m.CommissionInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("commissionInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("commissionInfo")
			}
			return err
		}
	}

	return nil
}

func (m *Validators) validateConsensusPubkey(formats strfmt.Registry) error {

	if err := validate.Required("consensusPubkey", "body", m.ConsensusPubkey); err != nil {
		return err
	}

	return nil
}

func (m *Validators) validateDelegatorShares(formats strfmt.Registry) error {

	if err := validate.Required("delegatorShares", "body", m.DelegatorShares); err != nil {
		return err
	}

	return nil
}

func (m *Validators) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	if m.Description != nil {
		if err := m.Description.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("description")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("description")
			}
			return err
		}
	}

	return nil
}

func (m *Validators) validateMyDelegation(formats strfmt.Registry) error {

	if err := validate.Required("myDelegation", "body", m.MyDelegation); err != nil {
		return err
	}

	return nil
}

func (m *Validators) validateMyUndelegation(formats strfmt.Registry) error {

	if err := validate.Required("myUndelegation", "body", m.MyUndelegation); err != nil {
		return err
	}

	return nil
}

func (m *Validators) validateOperatorAddress(formats strfmt.Registry) error {

	if err := validate.Required("operatorAddress", "body", m.OperatorAddress); err != nil {
		return err
	}

	return nil
}

func (m *Validators) validateRewardsPool(formats strfmt.Registry) error {

	if err := validate.Required("rewardsPool", "body", m.RewardsPool); err != nil {
		return err
	}

	if m.RewardsPool != nil {
		if err := m.RewardsPool.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rewardsPool")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rewardsPool")
			}
			return err
		}
	}

	return nil
}

func (m *Validators) validateStakingReturn(formats strfmt.Registry) error {

	if err := validate.Required("stakingReturn", "body", m.StakingReturn); err != nil {
		return err
	}

	return nil
}

func (m *Validators) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *Validators) validateTokens(formats strfmt.Registry) error {

	if err := validate.Required("tokens", "body", m.Tokens); err != nil {
		return err
	}

	return nil
}

func (m *Validators) validateUpTime(formats strfmt.Registry) error {

	if err := validate.Required("upTime", "body", m.UpTime); err != nil {
		return err
	}

	return nil
}

func (m *Validators) validateVotingPower(formats strfmt.Registry) error {

	if err := validate.Required("votingPower", "body", m.VotingPower); err != nil {
		return err
	}

	if m.VotingPower != nil {
		if err := m.VotingPower.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("votingPower")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("votingPower")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this validators based on the context it is used
func (m *Validators) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCommissionInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDescription(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRewardsPool(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVotingPower(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Validators) contextValidateCommissionInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.CommissionInfo != nil {
		if err := m.CommissionInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("commissionInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("commissionInfo")
			}
			return err
		}
	}

	return nil
}

func (m *Validators) contextValidateDescription(ctx context.Context, formats strfmt.Registry) error {

	if m.Description != nil {
		if err := m.Description.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("description")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("description")
			}
			return err
		}
	}

	return nil
}

func (m *Validators) contextValidateRewardsPool(ctx context.Context, formats strfmt.Registry) error {

	if m.RewardsPool != nil {
		if err := m.RewardsPool.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rewardsPool")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rewardsPool")
			}
			return err
		}
	}

	return nil
}

func (m *Validators) contextValidateVotingPower(ctx context.Context, formats strfmt.Registry) error {

	if m.VotingPower != nil {
		if err := m.VotingPower.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("votingPower")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("votingPower")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Validators) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Validators) UnmarshalBinary(b []byte) error {
	var res Validators
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
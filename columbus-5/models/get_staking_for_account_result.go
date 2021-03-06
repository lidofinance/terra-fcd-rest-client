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

// GetStakingForAccountResult get staking for account result
//
// swagger:model getStakingForAccountResult
type GetStakingForAccountResult struct {

	// Users total luna amount
	// Required: true
	AvailableLuna *string `json:"availableLuna"`

	// Amount staked by user
	// Required: true
	DelegationTotal *string `json:"delegationTotal"`

	// Users delegations list
	// Required: true
	MyDelegations []*GetStakingForAccountResultMyDelegations `json:"myDelegations"`

	// User's reward info
	// Required: true
	Rewards *GetStakingForAccountResultRewards `json:"rewards"`

	// Undelegation information in progress by user
	// Required: true
	Undelegations []*GetStakingForAccountResultUndelegations `json:"undelegations"`

	// validators
	// Required: true
	Validators []*GetStakingForAccountResultValidators `json:"validators"`
}

// Validate validates this get staking for account result
func (m *GetStakingForAccountResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailableLuna(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDelegationTotal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMyDelegations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRewards(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUndelegations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValidators(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetStakingForAccountResult) validateAvailableLuna(formats strfmt.Registry) error {

	if err := validate.Required("availableLuna", "body", m.AvailableLuna); err != nil {
		return err
	}

	return nil
}

func (m *GetStakingForAccountResult) validateDelegationTotal(formats strfmt.Registry) error {

	if err := validate.Required("delegationTotal", "body", m.DelegationTotal); err != nil {
		return err
	}

	return nil
}

func (m *GetStakingForAccountResult) validateMyDelegations(formats strfmt.Registry) error {

	if err := validate.Required("myDelegations", "body", m.MyDelegations); err != nil {
		return err
	}

	for i := 0; i < len(m.MyDelegations); i++ {
		if swag.IsZero(m.MyDelegations[i]) { // not required
			continue
		}

		if m.MyDelegations[i] != nil {
			if err := m.MyDelegations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("myDelegations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("myDelegations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GetStakingForAccountResult) validateRewards(formats strfmt.Registry) error {

	if err := validate.Required("rewards", "body", m.Rewards); err != nil {
		return err
	}

	if m.Rewards != nil {
		if err := m.Rewards.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rewards")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rewards")
			}
			return err
		}
	}

	return nil
}

func (m *GetStakingForAccountResult) validateUndelegations(formats strfmt.Registry) error {

	if err := validate.Required("undelegations", "body", m.Undelegations); err != nil {
		return err
	}

	for i := 0; i < len(m.Undelegations); i++ {
		if swag.IsZero(m.Undelegations[i]) { // not required
			continue
		}

		if m.Undelegations[i] != nil {
			if err := m.Undelegations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("undelegations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("undelegations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GetStakingForAccountResult) validateValidators(formats strfmt.Registry) error {

	if err := validate.Required("validators", "body", m.Validators); err != nil {
		return err
	}

	for i := 0; i < len(m.Validators); i++ {
		if swag.IsZero(m.Validators[i]) { // not required
			continue
		}

		if m.Validators[i] != nil {
			if err := m.Validators[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("validators" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("validators" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get staking for account result based on the context it is used
func (m *GetStakingForAccountResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMyDelegations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRewards(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUndelegations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateValidators(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetStakingForAccountResult) contextValidateMyDelegations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.MyDelegations); i++ {

		if m.MyDelegations[i] != nil {
			if err := m.MyDelegations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("myDelegations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("myDelegations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GetStakingForAccountResult) contextValidateRewards(ctx context.Context, formats strfmt.Registry) error {

	if m.Rewards != nil {
		if err := m.Rewards.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rewards")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("rewards")
			}
			return err
		}
	}

	return nil
}

func (m *GetStakingForAccountResult) contextValidateUndelegations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Undelegations); i++ {

		if m.Undelegations[i] != nil {
			if err := m.Undelegations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("undelegations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("undelegations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GetStakingForAccountResult) contextValidateValidators(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Validators); i++ {

		if m.Validators[i] != nil {
			if err := m.Validators[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("validators" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("validators" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetStakingForAccountResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetStakingForAccountResult) UnmarshalBinary(b []byte) error {
	var res GetStakingForAccountResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

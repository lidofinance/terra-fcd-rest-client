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

// BaseVestingAccount base vesting account
//
// swagger:model BaseVestingAccount
type BaseVestingAccount struct {

	// base account
	BaseAccount *BaseVestingAccountBaseAccount `json:"BaseAccount,omitempty"`

	// delegated free
	DelegatedFree []*BaseVestingAccountDelegatedFreeItems0 `json:"delegated_free"`

	// delegated vesting
	DelegatedVesting []*BaseVestingAccountDelegatedVestingItems0 `json:"delegated_vesting"`

	// end time
	// Example: 0
	EndTime string `json:"end_time,omitempty"`

	// original vesting
	OriginalVesting []*BaseVestingAccountOriginalVestingItems0 `json:"original_vesting"`
}

// Validate validates this base vesting account
func (m *BaseVestingAccount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBaseAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDelegatedFree(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDelegatedVesting(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginalVesting(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BaseVestingAccount) validateBaseAccount(formats strfmt.Registry) error {
	if swag.IsZero(m.BaseAccount) { // not required
		return nil
	}

	if m.BaseAccount != nil {
		if err := m.BaseAccount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("BaseAccount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("BaseAccount")
			}
			return err
		}
	}

	return nil
}

func (m *BaseVestingAccount) validateDelegatedFree(formats strfmt.Registry) error {
	if swag.IsZero(m.DelegatedFree) { // not required
		return nil
	}

	for i := 0; i < len(m.DelegatedFree); i++ {
		if swag.IsZero(m.DelegatedFree[i]) { // not required
			continue
		}

		if m.DelegatedFree[i] != nil {
			if err := m.DelegatedFree[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("delegated_free" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("delegated_free" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BaseVestingAccount) validateDelegatedVesting(formats strfmt.Registry) error {
	if swag.IsZero(m.DelegatedVesting) { // not required
		return nil
	}

	for i := 0; i < len(m.DelegatedVesting); i++ {
		if swag.IsZero(m.DelegatedVesting[i]) { // not required
			continue
		}

		if m.DelegatedVesting[i] != nil {
			if err := m.DelegatedVesting[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("delegated_vesting" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("delegated_vesting" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BaseVestingAccount) validateOriginalVesting(formats strfmt.Registry) error {
	if swag.IsZero(m.OriginalVesting) { // not required
		return nil
	}

	for i := 0; i < len(m.OriginalVesting); i++ {
		if swag.IsZero(m.OriginalVesting[i]) { // not required
			continue
		}

		if m.OriginalVesting[i] != nil {
			if err := m.OriginalVesting[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("original_vesting" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("original_vesting" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this base vesting account based on the context it is used
func (m *BaseVestingAccount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBaseAccount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDelegatedFree(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDelegatedVesting(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOriginalVesting(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BaseVestingAccount) contextValidateBaseAccount(ctx context.Context, formats strfmt.Registry) error {

	if m.BaseAccount != nil {
		if err := m.BaseAccount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("BaseAccount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("BaseAccount")
			}
			return err
		}
	}

	return nil
}

func (m *BaseVestingAccount) contextValidateDelegatedFree(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DelegatedFree); i++ {

		if m.DelegatedFree[i] != nil {
			if err := m.DelegatedFree[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("delegated_free" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("delegated_free" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BaseVestingAccount) contextValidateDelegatedVesting(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DelegatedVesting); i++ {

		if m.DelegatedVesting[i] != nil {
			if err := m.DelegatedVesting[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("delegated_vesting" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("delegated_vesting" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BaseVestingAccount) contextValidateOriginalVesting(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.OriginalVesting); i++ {

		if m.OriginalVesting[i] != nil {
			if err := m.OriginalVesting[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("original_vesting" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("original_vesting" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BaseVestingAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BaseVestingAccount) UnmarshalBinary(b []byte) error {
	var res BaseVestingAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BaseVestingAccountBaseAccount base vesting account base account
//
// swagger:model BaseVestingAccountBaseAccount
type BaseVestingAccountBaseAccount struct {

	// account number
	AccountNumber string `json:"account_number,omitempty"`

	// address
	Address string `json:"address,omitempty"`

	// public key
	PublicKey string `json:"public_key,omitempty"`

	// sequence
	Sequence string `json:"sequence,omitempty"`
}

// Validate validates this base vesting account base account
func (m *BaseVestingAccountBaseAccount) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this base vesting account base account based on context it is used
func (m *BaseVestingAccountBaseAccount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BaseVestingAccountBaseAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BaseVestingAccountBaseAccount) UnmarshalBinary(b []byte) error {
	var res BaseVestingAccountBaseAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BaseVestingAccountDelegatedFreeItems0 base vesting account delegated free items0
//
// swagger:model BaseVestingAccountDelegatedFreeItems0
type BaseVestingAccountDelegatedFreeItems0 struct {

	// amount
	// Example: 50
	Amount string `json:"amount,omitempty"`

	// denom
	// Example: uluna
	Denom string `json:"denom,omitempty"`
}

// Validate validates this base vesting account delegated free items0
func (m *BaseVestingAccountDelegatedFreeItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this base vesting account delegated free items0 based on context it is used
func (m *BaseVestingAccountDelegatedFreeItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BaseVestingAccountDelegatedFreeItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BaseVestingAccountDelegatedFreeItems0) UnmarshalBinary(b []byte) error {
	var res BaseVestingAccountDelegatedFreeItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BaseVestingAccountDelegatedVestingItems0 base vesting account delegated vesting items0
//
// swagger:model BaseVestingAccountDelegatedVestingItems0
type BaseVestingAccountDelegatedVestingItems0 struct {

	// amount
	// Example: 50
	Amount string `json:"amount,omitempty"`

	// denom
	// Example: uluna
	Denom string `json:"denom,omitempty"`
}

// Validate validates this base vesting account delegated vesting items0
func (m *BaseVestingAccountDelegatedVestingItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this base vesting account delegated vesting items0 based on context it is used
func (m *BaseVestingAccountDelegatedVestingItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BaseVestingAccountDelegatedVestingItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BaseVestingAccountDelegatedVestingItems0) UnmarshalBinary(b []byte) error {
	var res BaseVestingAccountDelegatedVestingItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BaseVestingAccountOriginalVestingItems0 base vesting account original vesting items0
//
// swagger:model BaseVestingAccountOriginalVestingItems0
type BaseVestingAccountOriginalVestingItems0 struct {

	// amount
	// Example: 50
	Amount string `json:"amount,omitempty"`

	// denom
	// Example: uluna
	Denom string `json:"denom,omitempty"`
}

// Validate validates this base vesting account original vesting items0
func (m *BaseVestingAccountOriginalVestingItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this base vesting account original vesting items0 based on context it is used
func (m *BaseVestingAccountOriginalVestingItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BaseVestingAccountOriginalVestingItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BaseVestingAccountOriginalVestingItems0) UnmarshalBinary(b []byte) error {
	var res BaseVestingAccountOriginalVestingItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

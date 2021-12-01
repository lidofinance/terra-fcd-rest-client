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

// LazyGradedVestingAccount lazy graded vesting account
//
// swagger:model LazyGradedVestingAccount
type LazyGradedVestingAccount struct {

	// type
	// Example: core/LazyGradedVestingAccount
	Type string `json:"type,omitempty"`

	// value
	Value *LazyGradedVestingAccountValue `json:"value,omitempty"`
}

// Validate validates this lazy graded vesting account
func (m *LazyGradedVestingAccount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LazyGradedVestingAccount) validateValue(formats strfmt.Registry) error {
	if swag.IsZero(m.Value) { // not required
		return nil
	}

	if m.Value != nil {
		if err := m.Value.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("value")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("value")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this lazy graded vesting account based on the context it is used
func (m *LazyGradedVestingAccount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LazyGradedVestingAccount) contextValidateValue(ctx context.Context, formats strfmt.Registry) error {

	if m.Value != nil {
		if err := m.Value.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("value")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("value")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LazyGradedVestingAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LazyGradedVestingAccount) UnmarshalBinary(b []byte) error {
	var res LazyGradedVestingAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LazyGradedVestingAccountValue lazy graded vesting account value
//
// swagger:model LazyGradedVestingAccountValue
type LazyGradedVestingAccountValue struct {

	// base vesting account
	BaseVestingAccount *LazyGradedVestingAccountValueBaseVestingAccount `json:"BaseVestingAccount,omitempty"`

	// vesting schedules
	VestingSchedules []*LazyGradedVestingAccountValueVestingSchedulesItems0 `json:"vesting_schedules"`
}

// Validate validates this lazy graded vesting account value
func (m *LazyGradedVestingAccountValue) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBaseVestingAccount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVestingSchedules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LazyGradedVestingAccountValue) validateBaseVestingAccount(formats strfmt.Registry) error {
	if swag.IsZero(m.BaseVestingAccount) { // not required
		return nil
	}

	if m.BaseVestingAccount != nil {
		if err := m.BaseVestingAccount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("value" + "." + "BaseVestingAccount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("value" + "." + "BaseVestingAccount")
			}
			return err
		}
	}

	return nil
}

func (m *LazyGradedVestingAccountValue) validateVestingSchedules(formats strfmt.Registry) error {
	if swag.IsZero(m.VestingSchedules) { // not required
		return nil
	}

	for i := 0; i < len(m.VestingSchedules); i++ {
		if swag.IsZero(m.VestingSchedules[i]) { // not required
			continue
		}

		if m.VestingSchedules[i] != nil {
			if err := m.VestingSchedules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("value" + "." + "vesting_schedules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("value" + "." + "vesting_schedules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this lazy graded vesting account value based on the context it is used
func (m *LazyGradedVestingAccountValue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBaseVestingAccount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVestingSchedules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LazyGradedVestingAccountValue) contextValidateBaseVestingAccount(ctx context.Context, formats strfmt.Registry) error {

	if m.BaseVestingAccount != nil {
		if err := m.BaseVestingAccount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("value" + "." + "BaseVestingAccount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("value" + "." + "BaseVestingAccount")
			}
			return err
		}
	}

	return nil
}

func (m *LazyGradedVestingAccountValue) contextValidateVestingSchedules(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VestingSchedules); i++ {

		if m.VestingSchedules[i] != nil {
			if err := m.VestingSchedules[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("value" + "." + "vesting_schedules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("value" + "." + "vesting_schedules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LazyGradedVestingAccountValue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LazyGradedVestingAccountValue) UnmarshalBinary(b []byte) error {
	var res LazyGradedVestingAccountValue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LazyGradedVestingAccountValueBaseVestingAccount lazy graded vesting account value base vesting account
//
// swagger:model LazyGradedVestingAccountValueBaseVestingAccount
type LazyGradedVestingAccountValueBaseVestingAccount struct {

	// base account
	BaseAccount *LazyGradedVestingAccountValueBaseVestingAccountBaseAccount `json:"BaseAccount,omitempty"`

	// delegated free
	DelegatedFree []*LazyGradedVestingAccountValueBaseVestingAccountDelegatedFreeItems0 `json:"delegated_free"`

	// delegated vesting
	DelegatedVesting []*LazyGradedVestingAccountValueBaseVestingAccountDelegatedVestingItems0 `json:"delegated_vesting"`

	// end time
	// Example: 0
	EndTime string `json:"end_time,omitempty"`

	// original vesting
	OriginalVesting []*LazyGradedVestingAccountValueBaseVestingAccountOriginalVestingItems0 `json:"original_vesting"`
}

// Validate validates this lazy graded vesting account value base vesting account
func (m *LazyGradedVestingAccountValueBaseVestingAccount) Validate(formats strfmt.Registry) error {
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

func (m *LazyGradedVestingAccountValueBaseVestingAccount) validateBaseAccount(formats strfmt.Registry) error {
	if swag.IsZero(m.BaseAccount) { // not required
		return nil
	}

	if m.BaseAccount != nil {
		if err := m.BaseAccount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("value" + "." + "BaseVestingAccount" + "." + "BaseAccount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("value" + "." + "BaseVestingAccount" + "." + "BaseAccount")
			}
			return err
		}
	}

	return nil
}

func (m *LazyGradedVestingAccountValueBaseVestingAccount) validateDelegatedFree(formats strfmt.Registry) error {
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
					return ve.ValidateName("value" + "." + "BaseVestingAccount" + "." + "delegated_free" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("value" + "." + "BaseVestingAccount" + "." + "delegated_free" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LazyGradedVestingAccountValueBaseVestingAccount) validateDelegatedVesting(formats strfmt.Registry) error {
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
					return ve.ValidateName("value" + "." + "BaseVestingAccount" + "." + "delegated_vesting" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("value" + "." + "BaseVestingAccount" + "." + "delegated_vesting" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LazyGradedVestingAccountValueBaseVestingAccount) validateOriginalVesting(formats strfmt.Registry) error {
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
					return ve.ValidateName("value" + "." + "BaseVestingAccount" + "." + "original_vesting" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("value" + "." + "BaseVestingAccount" + "." + "original_vesting" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this lazy graded vesting account value base vesting account based on the context it is used
func (m *LazyGradedVestingAccountValueBaseVestingAccount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *LazyGradedVestingAccountValueBaseVestingAccount) contextValidateBaseAccount(ctx context.Context, formats strfmt.Registry) error {

	if m.BaseAccount != nil {
		if err := m.BaseAccount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("value" + "." + "BaseVestingAccount" + "." + "BaseAccount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("value" + "." + "BaseVestingAccount" + "." + "BaseAccount")
			}
			return err
		}
	}

	return nil
}

func (m *LazyGradedVestingAccountValueBaseVestingAccount) contextValidateDelegatedFree(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DelegatedFree); i++ {

		if m.DelegatedFree[i] != nil {
			if err := m.DelegatedFree[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("value" + "." + "BaseVestingAccount" + "." + "delegated_free" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("value" + "." + "BaseVestingAccount" + "." + "delegated_free" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LazyGradedVestingAccountValueBaseVestingAccount) contextValidateDelegatedVesting(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DelegatedVesting); i++ {

		if m.DelegatedVesting[i] != nil {
			if err := m.DelegatedVesting[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("value" + "." + "BaseVestingAccount" + "." + "delegated_vesting" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("value" + "." + "BaseVestingAccount" + "." + "delegated_vesting" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LazyGradedVestingAccountValueBaseVestingAccount) contextValidateOriginalVesting(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.OriginalVesting); i++ {

		if m.OriginalVesting[i] != nil {
			if err := m.OriginalVesting[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("value" + "." + "BaseVestingAccount" + "." + "original_vesting" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("value" + "." + "BaseVestingAccount" + "." + "original_vesting" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LazyGradedVestingAccountValueBaseVestingAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LazyGradedVestingAccountValueBaseVestingAccount) UnmarshalBinary(b []byte) error {
	var res LazyGradedVestingAccountValueBaseVestingAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LazyGradedVestingAccountValueBaseVestingAccountBaseAccount lazy graded vesting account value base vesting account base account
//
// swagger:model LazyGradedVestingAccountValueBaseVestingAccountBaseAccount
type LazyGradedVestingAccountValueBaseVestingAccountBaseAccount struct {

	// account number
	AccountNumber string `json:"account_number,omitempty"`

	// address
	Address string `json:"address,omitempty"`

	// public key
	PublicKey string `json:"public_key,omitempty"`

	// sequence
	Sequence string `json:"sequence,omitempty"`
}

// Validate validates this lazy graded vesting account value base vesting account base account
func (m *LazyGradedVestingAccountValueBaseVestingAccountBaseAccount) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this lazy graded vesting account value base vesting account base account based on context it is used
func (m *LazyGradedVestingAccountValueBaseVestingAccountBaseAccount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LazyGradedVestingAccountValueBaseVestingAccountBaseAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LazyGradedVestingAccountValueBaseVestingAccountBaseAccount) UnmarshalBinary(b []byte) error {
	var res LazyGradedVestingAccountValueBaseVestingAccountBaseAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LazyGradedVestingAccountValueBaseVestingAccountDelegatedFreeItems0 lazy graded vesting account value base vesting account delegated free items0
//
// swagger:model LazyGradedVestingAccountValueBaseVestingAccountDelegatedFreeItems0
type LazyGradedVestingAccountValueBaseVestingAccountDelegatedFreeItems0 struct {

	// amount
	// Example: 50
	Amount string `json:"amount,omitempty"`

	// denom
	// Example: uluna
	Denom string `json:"denom,omitempty"`
}

// Validate validates this lazy graded vesting account value base vesting account delegated free items0
func (m *LazyGradedVestingAccountValueBaseVestingAccountDelegatedFreeItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this lazy graded vesting account value base vesting account delegated free items0 based on context it is used
func (m *LazyGradedVestingAccountValueBaseVestingAccountDelegatedFreeItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LazyGradedVestingAccountValueBaseVestingAccountDelegatedFreeItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LazyGradedVestingAccountValueBaseVestingAccountDelegatedFreeItems0) UnmarshalBinary(b []byte) error {
	var res LazyGradedVestingAccountValueBaseVestingAccountDelegatedFreeItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LazyGradedVestingAccountValueBaseVestingAccountDelegatedVestingItems0 lazy graded vesting account value base vesting account delegated vesting items0
//
// swagger:model LazyGradedVestingAccountValueBaseVestingAccountDelegatedVestingItems0
type LazyGradedVestingAccountValueBaseVestingAccountDelegatedVestingItems0 struct {

	// amount
	// Example: 50
	Amount string `json:"amount,omitempty"`

	// denom
	// Example: uluna
	Denom string `json:"denom,omitempty"`
}

// Validate validates this lazy graded vesting account value base vesting account delegated vesting items0
func (m *LazyGradedVestingAccountValueBaseVestingAccountDelegatedVestingItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this lazy graded vesting account value base vesting account delegated vesting items0 based on context it is used
func (m *LazyGradedVestingAccountValueBaseVestingAccountDelegatedVestingItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LazyGradedVestingAccountValueBaseVestingAccountDelegatedVestingItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LazyGradedVestingAccountValueBaseVestingAccountDelegatedVestingItems0) UnmarshalBinary(b []byte) error {
	var res LazyGradedVestingAccountValueBaseVestingAccountDelegatedVestingItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LazyGradedVestingAccountValueBaseVestingAccountOriginalVestingItems0 lazy graded vesting account value base vesting account original vesting items0
//
// swagger:model LazyGradedVestingAccountValueBaseVestingAccountOriginalVestingItems0
type LazyGradedVestingAccountValueBaseVestingAccountOriginalVestingItems0 struct {

	// amount
	// Example: 50
	Amount string `json:"amount,omitempty"`

	// denom
	// Example: uluna
	Denom string `json:"denom,omitempty"`
}

// Validate validates this lazy graded vesting account value base vesting account original vesting items0
func (m *LazyGradedVestingAccountValueBaseVestingAccountOriginalVestingItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this lazy graded vesting account value base vesting account original vesting items0 based on context it is used
func (m *LazyGradedVestingAccountValueBaseVestingAccountOriginalVestingItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LazyGradedVestingAccountValueBaseVestingAccountOriginalVestingItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LazyGradedVestingAccountValueBaseVestingAccountOriginalVestingItems0) UnmarshalBinary(b []byte) error {
	var res LazyGradedVestingAccountValueBaseVestingAccountOriginalVestingItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LazyGradedVestingAccountValueVestingSchedulesItems0 lazy graded vesting account value vesting schedules items0
//
// swagger:model LazyGradedVestingAccountValueVestingSchedulesItems0
type LazyGradedVestingAccountValueVestingSchedulesItems0 struct {

	// denom
	// Example: usdr
	Denom string `json:"denom,omitempty"`

	// lazy schedules
	LazySchedules []*LazyGradedVestingAccountValueVestingSchedulesItems0LazySchedulesItems0 `json:"lazy_schedules"`
}

// Validate validates this lazy graded vesting account value vesting schedules items0
func (m *LazyGradedVestingAccountValueVestingSchedulesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLazySchedules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LazyGradedVestingAccountValueVestingSchedulesItems0) validateLazySchedules(formats strfmt.Registry) error {
	if swag.IsZero(m.LazySchedules) { // not required
		return nil
	}

	for i := 0; i < len(m.LazySchedules); i++ {
		if swag.IsZero(m.LazySchedules[i]) { // not required
			continue
		}

		if m.LazySchedules[i] != nil {
			if err := m.LazySchedules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("lazy_schedules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("lazy_schedules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this lazy graded vesting account value vesting schedules items0 based on the context it is used
func (m *LazyGradedVestingAccountValueVestingSchedulesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLazySchedules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LazyGradedVestingAccountValueVestingSchedulesItems0) contextValidateLazySchedules(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LazySchedules); i++ {

		if m.LazySchedules[i] != nil {
			if err := m.LazySchedules[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("lazy_schedules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("lazy_schedules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LazyGradedVestingAccountValueVestingSchedulesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LazyGradedVestingAccountValueVestingSchedulesItems0) UnmarshalBinary(b []byte) error {
	var res LazyGradedVestingAccountValueVestingSchedulesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// LazyGradedVestingAccountValueVestingSchedulesItems0LazySchedulesItems0 lazy graded vesting account value vesting schedules items0 lazy schedules items0
//
// swagger:model LazyGradedVestingAccountValueVestingSchedulesItems0LazySchedulesItems0
type LazyGradedVestingAccountValueVestingSchedulesItems0LazySchedulesItems0 struct {

	// end time
	// Example: 1556085600
	EndTime string `json:"end_time,omitempty"`

	// ratio
	// Example: 0.100000000000000000
	Ratio string `json:"ratio,omitempty"`

	// start time
	// Example: 1556085600
	StartTime string `json:"start_time,omitempty"`
}

// Validate validates this lazy graded vesting account value vesting schedules items0 lazy schedules items0
func (m *LazyGradedVestingAccountValueVestingSchedulesItems0LazySchedulesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this lazy graded vesting account value vesting schedules items0 lazy schedules items0 based on context it is used
func (m *LazyGradedVestingAccountValueVestingSchedulesItems0LazySchedulesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LazyGradedVestingAccountValueVestingSchedulesItems0LazySchedulesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LazyGradedVestingAccountValueVestingSchedulesItems0LazySchedulesItems0) UnmarshalBinary(b []byte) error {
	var res LazyGradedVestingAccountValueVestingSchedulesItems0LazySchedulesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
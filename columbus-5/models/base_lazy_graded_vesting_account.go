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

// BaseLazyGradedVestingAccount base lazy graded vesting account
//
// swagger:model BaseLazyGradedVestingAccount
type BaseLazyGradedVestingAccount struct {

	// base vesting account
	BaseVestingAccount *BaseLazyGradedVestingAccountBaseVestingAccount `json:"BaseVestingAccount,omitempty"`

	// vesting schedules
	VestingSchedules []*BaseLazyGradedVestingAccountVestingSchedulesItems0 `json:"vesting_schedules"`
}

// Validate validates this base lazy graded vesting account
func (m *BaseLazyGradedVestingAccount) Validate(formats strfmt.Registry) error {
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

func (m *BaseLazyGradedVestingAccount) validateBaseVestingAccount(formats strfmt.Registry) error {
	if swag.IsZero(m.BaseVestingAccount) { // not required
		return nil
	}

	if m.BaseVestingAccount != nil {
		if err := m.BaseVestingAccount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("BaseVestingAccount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("BaseVestingAccount")
			}
			return err
		}
	}

	return nil
}

func (m *BaseLazyGradedVestingAccount) validateVestingSchedules(formats strfmt.Registry) error {
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
					return ve.ValidateName("vesting_schedules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vesting_schedules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this base lazy graded vesting account based on the context it is used
func (m *BaseLazyGradedVestingAccount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *BaseLazyGradedVestingAccount) contextValidateBaseVestingAccount(ctx context.Context, formats strfmt.Registry) error {

	if m.BaseVestingAccount != nil {
		if err := m.BaseVestingAccount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("BaseVestingAccount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("BaseVestingAccount")
			}
			return err
		}
	}

	return nil
}

func (m *BaseLazyGradedVestingAccount) contextValidateVestingSchedules(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.VestingSchedules); i++ {

		if m.VestingSchedules[i] != nil {
			if err := m.VestingSchedules[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vesting_schedules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vesting_schedules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BaseLazyGradedVestingAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BaseLazyGradedVestingAccount) UnmarshalBinary(b []byte) error {
	var res BaseLazyGradedVestingAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BaseLazyGradedVestingAccountBaseVestingAccount base lazy graded vesting account base vesting account
//
// swagger:model BaseLazyGradedVestingAccountBaseVestingAccount
type BaseLazyGradedVestingAccountBaseVestingAccount struct {

	// base account
	BaseAccount *BaseLazyGradedVestingAccountBaseVestingAccountBaseAccount `json:"BaseAccount,omitempty"`

	// delegated free
	DelegatedFree []*BaseLazyGradedVestingAccountBaseVestingAccountDelegatedFreeItems0 `json:"delegated_free"`

	// delegated vesting
	DelegatedVesting []*BaseLazyGradedVestingAccountBaseVestingAccountDelegatedVestingItems0 `json:"delegated_vesting"`

	// end time
	// Example: 0
	EndTime string `json:"end_time,omitempty"`

	// original vesting
	OriginalVesting []*BaseLazyGradedVestingAccountBaseVestingAccountOriginalVestingItems0 `json:"original_vesting"`
}

// Validate validates this base lazy graded vesting account base vesting account
func (m *BaseLazyGradedVestingAccountBaseVestingAccount) Validate(formats strfmt.Registry) error {
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

func (m *BaseLazyGradedVestingAccountBaseVestingAccount) validateBaseAccount(formats strfmt.Registry) error {
	if swag.IsZero(m.BaseAccount) { // not required
		return nil
	}

	if m.BaseAccount != nil {
		if err := m.BaseAccount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("BaseVestingAccount" + "." + "BaseAccount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("BaseVestingAccount" + "." + "BaseAccount")
			}
			return err
		}
	}

	return nil
}

func (m *BaseLazyGradedVestingAccountBaseVestingAccount) validateDelegatedFree(formats strfmt.Registry) error {
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
					return ve.ValidateName("BaseVestingAccount" + "." + "delegated_free" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("BaseVestingAccount" + "." + "delegated_free" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BaseLazyGradedVestingAccountBaseVestingAccount) validateDelegatedVesting(formats strfmt.Registry) error {
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
					return ve.ValidateName("BaseVestingAccount" + "." + "delegated_vesting" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("BaseVestingAccount" + "." + "delegated_vesting" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BaseLazyGradedVestingAccountBaseVestingAccount) validateOriginalVesting(formats strfmt.Registry) error {
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
					return ve.ValidateName("BaseVestingAccount" + "." + "original_vesting" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("BaseVestingAccount" + "." + "original_vesting" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this base lazy graded vesting account base vesting account based on the context it is used
func (m *BaseLazyGradedVestingAccountBaseVestingAccount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *BaseLazyGradedVestingAccountBaseVestingAccount) contextValidateBaseAccount(ctx context.Context, formats strfmt.Registry) error {

	if m.BaseAccount != nil {
		if err := m.BaseAccount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("BaseVestingAccount" + "." + "BaseAccount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("BaseVestingAccount" + "." + "BaseAccount")
			}
			return err
		}
	}

	return nil
}

func (m *BaseLazyGradedVestingAccountBaseVestingAccount) contextValidateDelegatedFree(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DelegatedFree); i++ {

		if m.DelegatedFree[i] != nil {
			if err := m.DelegatedFree[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("BaseVestingAccount" + "." + "delegated_free" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("BaseVestingAccount" + "." + "delegated_free" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BaseLazyGradedVestingAccountBaseVestingAccount) contextValidateDelegatedVesting(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DelegatedVesting); i++ {

		if m.DelegatedVesting[i] != nil {
			if err := m.DelegatedVesting[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("BaseVestingAccount" + "." + "delegated_vesting" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("BaseVestingAccount" + "." + "delegated_vesting" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BaseLazyGradedVestingAccountBaseVestingAccount) contextValidateOriginalVesting(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.OriginalVesting); i++ {

		if m.OriginalVesting[i] != nil {
			if err := m.OriginalVesting[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("BaseVestingAccount" + "." + "original_vesting" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("BaseVestingAccount" + "." + "original_vesting" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BaseLazyGradedVestingAccountBaseVestingAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BaseLazyGradedVestingAccountBaseVestingAccount) UnmarshalBinary(b []byte) error {
	var res BaseLazyGradedVestingAccountBaseVestingAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BaseLazyGradedVestingAccountBaseVestingAccountBaseAccount base lazy graded vesting account base vesting account base account
//
// swagger:model BaseLazyGradedVestingAccountBaseVestingAccountBaseAccount
type BaseLazyGradedVestingAccountBaseVestingAccountBaseAccount struct {

	// account number
	AccountNumber string `json:"account_number,omitempty"`

	// address
	Address string `json:"address,omitempty"`

	// public key
	PublicKey string `json:"public_key,omitempty"`

	// sequence
	Sequence string `json:"sequence,omitempty"`
}

// Validate validates this base lazy graded vesting account base vesting account base account
func (m *BaseLazyGradedVestingAccountBaseVestingAccountBaseAccount) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this base lazy graded vesting account base vesting account base account based on context it is used
func (m *BaseLazyGradedVestingAccountBaseVestingAccountBaseAccount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BaseLazyGradedVestingAccountBaseVestingAccountBaseAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BaseLazyGradedVestingAccountBaseVestingAccountBaseAccount) UnmarshalBinary(b []byte) error {
	var res BaseLazyGradedVestingAccountBaseVestingAccountBaseAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BaseLazyGradedVestingAccountBaseVestingAccountDelegatedFreeItems0 base lazy graded vesting account base vesting account delegated free items0
//
// swagger:model BaseLazyGradedVestingAccountBaseVestingAccountDelegatedFreeItems0
type BaseLazyGradedVestingAccountBaseVestingAccountDelegatedFreeItems0 struct {

	// amount
	// Example: 50
	Amount string `json:"amount,omitempty"`

	// denom
	// Example: uluna
	Denom string `json:"denom,omitempty"`
}

// Validate validates this base lazy graded vesting account base vesting account delegated free items0
func (m *BaseLazyGradedVestingAccountBaseVestingAccountDelegatedFreeItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this base lazy graded vesting account base vesting account delegated free items0 based on context it is used
func (m *BaseLazyGradedVestingAccountBaseVestingAccountDelegatedFreeItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BaseLazyGradedVestingAccountBaseVestingAccountDelegatedFreeItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BaseLazyGradedVestingAccountBaseVestingAccountDelegatedFreeItems0) UnmarshalBinary(b []byte) error {
	var res BaseLazyGradedVestingAccountBaseVestingAccountDelegatedFreeItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BaseLazyGradedVestingAccountBaseVestingAccountDelegatedVestingItems0 base lazy graded vesting account base vesting account delegated vesting items0
//
// swagger:model BaseLazyGradedVestingAccountBaseVestingAccountDelegatedVestingItems0
type BaseLazyGradedVestingAccountBaseVestingAccountDelegatedVestingItems0 struct {

	// amount
	// Example: 50
	Amount string `json:"amount,omitempty"`

	// denom
	// Example: uluna
	Denom string `json:"denom,omitempty"`
}

// Validate validates this base lazy graded vesting account base vesting account delegated vesting items0
func (m *BaseLazyGradedVestingAccountBaseVestingAccountDelegatedVestingItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this base lazy graded vesting account base vesting account delegated vesting items0 based on context it is used
func (m *BaseLazyGradedVestingAccountBaseVestingAccountDelegatedVestingItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BaseLazyGradedVestingAccountBaseVestingAccountDelegatedVestingItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BaseLazyGradedVestingAccountBaseVestingAccountDelegatedVestingItems0) UnmarshalBinary(b []byte) error {
	var res BaseLazyGradedVestingAccountBaseVestingAccountDelegatedVestingItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BaseLazyGradedVestingAccountBaseVestingAccountOriginalVestingItems0 base lazy graded vesting account base vesting account original vesting items0
//
// swagger:model BaseLazyGradedVestingAccountBaseVestingAccountOriginalVestingItems0
type BaseLazyGradedVestingAccountBaseVestingAccountOriginalVestingItems0 struct {

	// amount
	// Example: 50
	Amount string `json:"amount,omitempty"`

	// denom
	// Example: uluna
	Denom string `json:"denom,omitempty"`
}

// Validate validates this base lazy graded vesting account base vesting account original vesting items0
func (m *BaseLazyGradedVestingAccountBaseVestingAccountOriginalVestingItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this base lazy graded vesting account base vesting account original vesting items0 based on context it is used
func (m *BaseLazyGradedVestingAccountBaseVestingAccountOriginalVestingItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BaseLazyGradedVestingAccountBaseVestingAccountOriginalVestingItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BaseLazyGradedVestingAccountBaseVestingAccountOriginalVestingItems0) UnmarshalBinary(b []byte) error {
	var res BaseLazyGradedVestingAccountBaseVestingAccountOriginalVestingItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BaseLazyGradedVestingAccountVestingSchedulesItems0 base lazy graded vesting account vesting schedules items0
//
// swagger:model BaseLazyGradedVestingAccountVestingSchedulesItems0
type BaseLazyGradedVestingAccountVestingSchedulesItems0 struct {

	// denom
	// Example: usdr
	Denom string `json:"denom,omitempty"`

	// lazy schedules
	LazySchedules []*BaseLazyGradedVestingAccountVestingSchedulesItems0LazySchedulesItems0 `json:"lazy_schedules"`
}

// Validate validates this base lazy graded vesting account vesting schedules items0
func (m *BaseLazyGradedVestingAccountVestingSchedulesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLazySchedules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BaseLazyGradedVestingAccountVestingSchedulesItems0) validateLazySchedules(formats strfmt.Registry) error {
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

// ContextValidate validate this base lazy graded vesting account vesting schedules items0 based on the context it is used
func (m *BaseLazyGradedVestingAccountVestingSchedulesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLazySchedules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BaseLazyGradedVestingAccountVestingSchedulesItems0) contextValidateLazySchedules(ctx context.Context, formats strfmt.Registry) error {

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
func (m *BaseLazyGradedVestingAccountVestingSchedulesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BaseLazyGradedVestingAccountVestingSchedulesItems0) UnmarshalBinary(b []byte) error {
	var res BaseLazyGradedVestingAccountVestingSchedulesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BaseLazyGradedVestingAccountVestingSchedulesItems0LazySchedulesItems0 base lazy graded vesting account vesting schedules items0 lazy schedules items0
//
// swagger:model BaseLazyGradedVestingAccountVestingSchedulesItems0LazySchedulesItems0
type BaseLazyGradedVestingAccountVestingSchedulesItems0LazySchedulesItems0 struct {

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

// Validate validates this base lazy graded vesting account vesting schedules items0 lazy schedules items0
func (m *BaseLazyGradedVestingAccountVestingSchedulesItems0LazySchedulesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this base lazy graded vesting account vesting schedules items0 lazy schedules items0 based on context it is used
func (m *BaseLazyGradedVestingAccountVestingSchedulesItems0LazySchedulesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BaseLazyGradedVestingAccountVestingSchedulesItems0LazySchedulesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BaseLazyGradedVestingAccountVestingSchedulesItems0LazySchedulesItems0) UnmarshalBinary(b []byte) error {
	var res BaseLazyGradedVestingAccountVestingSchedulesItems0LazySchedulesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

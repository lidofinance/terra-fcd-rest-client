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

// CosmosDistributionV1beta1ValidatorAccumulatedCommission ValidatorAccumulatedCommission represents accumulated commission
// for a validator kept as a running counter, can be withdrawn at any time.
//
// swagger:model cosmos.distribution.v1beta1.ValidatorAccumulatedCommission
type CosmosDistributionV1beta1ValidatorAccumulatedCommission struct {

	// commission
	Commission []*CosmosDistributionV1beta1ValidatorAccumulatedCommissionCommissionItems0 `json:"commission"`
}

// Validate validates this cosmos distribution v1beta1 validator accumulated commission
func (m *CosmosDistributionV1beta1ValidatorAccumulatedCommission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommission(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosDistributionV1beta1ValidatorAccumulatedCommission) validateCommission(formats strfmt.Registry) error {
	if swag.IsZero(m.Commission) { // not required
		return nil
	}

	for i := 0; i < len(m.Commission); i++ {
		if swag.IsZero(m.Commission[i]) { // not required
			continue
		}

		if m.Commission[i] != nil {
			if err := m.Commission[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("commission" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("commission" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this cosmos distribution v1beta1 validator accumulated commission based on the context it is used
func (m *CosmosDistributionV1beta1ValidatorAccumulatedCommission) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCommission(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosDistributionV1beta1ValidatorAccumulatedCommission) contextValidateCommission(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Commission); i++ {

		if m.Commission[i] != nil {
			if err := m.Commission[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("commission" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("commission" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CosmosDistributionV1beta1ValidatorAccumulatedCommission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosDistributionV1beta1ValidatorAccumulatedCommission) UnmarshalBinary(b []byte) error {
	var res CosmosDistributionV1beta1ValidatorAccumulatedCommission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CosmosDistributionV1beta1ValidatorAccumulatedCommissionCommissionItems0 DecCoin defines a token with a denomination and a decimal amount.
//
// NOTE: The amount field is an Dec which implements the custom method
// signatures required by gogoproto.
//
// swagger:model CosmosDistributionV1beta1ValidatorAccumulatedCommissionCommissionItems0
type CosmosDistributionV1beta1ValidatorAccumulatedCommissionCommissionItems0 struct {

	// amount
	Amount string `json:"amount,omitempty"`

	// denom
	Denom string `json:"denom,omitempty"`
}

// Validate validates this cosmos distribution v1beta1 validator accumulated commission commission items0
func (m *CosmosDistributionV1beta1ValidatorAccumulatedCommissionCommissionItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cosmos distribution v1beta1 validator accumulated commission commission items0 based on context it is used
func (m *CosmosDistributionV1beta1ValidatorAccumulatedCommissionCommissionItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CosmosDistributionV1beta1ValidatorAccumulatedCommissionCommissionItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosDistributionV1beta1ValidatorAccumulatedCommissionCommissionItems0) UnmarshalBinary(b []byte) error {
	var res CosmosDistributionV1beta1ValidatorAccumulatedCommissionCommissionItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

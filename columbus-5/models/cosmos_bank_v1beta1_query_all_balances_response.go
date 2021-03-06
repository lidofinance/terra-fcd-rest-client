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

// CosmosBankV1beta1QueryAllBalancesResponse QueryAllBalancesResponse is the response type for the Query/AllBalances RPC
// method.
//
// swagger:model cosmos.bank.v1beta1.QueryAllBalancesResponse
type CosmosBankV1beta1QueryAllBalancesResponse struct {

	// balances is the balances of all the coins.
	Balances []*CosmosBankV1beta1QueryAllBalancesResponseBalancesItems0 `json:"balances"`

	// pagination
	Pagination *CosmosBankV1beta1QueryAllBalancesResponsePagination `json:"pagination,omitempty"`
}

// Validate validates this cosmos bank v1beta1 query all balances response
func (m *CosmosBankV1beta1QueryAllBalancesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBalances(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePagination(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosBankV1beta1QueryAllBalancesResponse) validateBalances(formats strfmt.Registry) error {
	if swag.IsZero(m.Balances) { // not required
		return nil
	}

	for i := 0; i < len(m.Balances); i++ {
		if swag.IsZero(m.Balances[i]) { // not required
			continue
		}

		if m.Balances[i] != nil {
			if err := m.Balances[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("balances" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("balances" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CosmosBankV1beta1QueryAllBalancesResponse) validatePagination(formats strfmt.Registry) error {
	if swag.IsZero(m.Pagination) { // not required
		return nil
	}

	if m.Pagination != nil {
		if err := m.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pagination")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cosmos bank v1beta1 query all balances response based on the context it is used
func (m *CosmosBankV1beta1QueryAllBalancesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBalances(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePagination(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosBankV1beta1QueryAllBalancesResponse) contextValidateBalances(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Balances); i++ {

		if m.Balances[i] != nil {
			if err := m.Balances[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("balances" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("balances" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CosmosBankV1beta1QueryAllBalancesResponse) contextValidatePagination(ctx context.Context, formats strfmt.Registry) error {

	if m.Pagination != nil {
		if err := m.Pagination.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pagination")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CosmosBankV1beta1QueryAllBalancesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosBankV1beta1QueryAllBalancesResponse) UnmarshalBinary(b []byte) error {
	var res CosmosBankV1beta1QueryAllBalancesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CosmosBankV1beta1QueryAllBalancesResponseBalancesItems0 Coin defines a token with a denomination and an amount.
//
// NOTE: The amount field is an Int which implements the custom method
// signatures required by gogoproto.
//
// swagger:model CosmosBankV1beta1QueryAllBalancesResponseBalancesItems0
type CosmosBankV1beta1QueryAllBalancesResponseBalancesItems0 struct {

	// amount
	Amount string `json:"amount,omitempty"`

	// denom
	Denom string `json:"denom,omitempty"`
}

// Validate validates this cosmos bank v1beta1 query all balances response balances items0
func (m *CosmosBankV1beta1QueryAllBalancesResponseBalancesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cosmos bank v1beta1 query all balances response balances items0 based on context it is used
func (m *CosmosBankV1beta1QueryAllBalancesResponseBalancesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CosmosBankV1beta1QueryAllBalancesResponseBalancesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosBankV1beta1QueryAllBalancesResponseBalancesItems0) UnmarshalBinary(b []byte) error {
	var res CosmosBankV1beta1QueryAllBalancesResponseBalancesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CosmosBankV1beta1QueryAllBalancesResponsePagination pagination defines the pagination in the response.
//
// swagger:model CosmosBankV1beta1QueryAllBalancesResponsePagination
type CosmosBankV1beta1QueryAllBalancesResponsePagination struct {

	// next_key is the key to be passed to PageRequest.key to
	// query the next page most efficiently
	// Format: byte
	NextKey strfmt.Base64 `json:"next_key,omitempty"`

	// total is total number of results available if PageRequest.count_total
	// was set, its value is undefined otherwise
	Total string `json:"total,omitempty"`
}

// Validate validates this cosmos bank v1beta1 query all balances response pagination
func (m *CosmosBankV1beta1QueryAllBalancesResponsePagination) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cosmos bank v1beta1 query all balances response pagination based on context it is used
func (m *CosmosBankV1beta1QueryAllBalancesResponsePagination) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CosmosBankV1beta1QueryAllBalancesResponsePagination) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosBankV1beta1QueryAllBalancesResponsePagination) UnmarshalBinary(b []byte) error {
	var res CosmosBankV1beta1QueryAllBalancesResponsePagination
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

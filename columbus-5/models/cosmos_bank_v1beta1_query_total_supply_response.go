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

// CosmosBankV1beta1QueryTotalSupplyResponse QueryTotalSupplyResponse is the response type for the Query/TotalSupply RPC
// method
//
// swagger:model cosmos.bank.v1beta1.QueryTotalSupplyResponse
type CosmosBankV1beta1QueryTotalSupplyResponse struct {

	// pagination
	Pagination *CosmosBankV1beta1QueryTotalSupplyResponsePagination `json:"pagination,omitempty"`

	// supply is the supply of the coins
	Supply []*CosmosBankV1beta1QueryTotalSupplyResponseSupplyItems0 `json:"supply"`
}

// Validate validates this cosmos bank v1beta1 query total supply response
func (m *CosmosBankV1beta1QueryTotalSupplyResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePagination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSupply(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosBankV1beta1QueryTotalSupplyResponse) validatePagination(formats strfmt.Registry) error {
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

func (m *CosmosBankV1beta1QueryTotalSupplyResponse) validateSupply(formats strfmt.Registry) error {
	if swag.IsZero(m.Supply) { // not required
		return nil
	}

	for i := 0; i < len(m.Supply); i++ {
		if swag.IsZero(m.Supply[i]) { // not required
			continue
		}

		if m.Supply[i] != nil {
			if err := m.Supply[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("supply" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("supply" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this cosmos bank v1beta1 query total supply response based on the context it is used
func (m *CosmosBankV1beta1QueryTotalSupplyResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePagination(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSupply(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosBankV1beta1QueryTotalSupplyResponse) contextValidatePagination(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CosmosBankV1beta1QueryTotalSupplyResponse) contextValidateSupply(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Supply); i++ {

		if m.Supply[i] != nil {
			if err := m.Supply[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("supply" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("supply" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CosmosBankV1beta1QueryTotalSupplyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosBankV1beta1QueryTotalSupplyResponse) UnmarshalBinary(b []byte) error {
	var res CosmosBankV1beta1QueryTotalSupplyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CosmosBankV1beta1QueryTotalSupplyResponsePagination pagination defines the pagination in the response.
//
// swagger:model CosmosBankV1beta1QueryTotalSupplyResponsePagination
type CosmosBankV1beta1QueryTotalSupplyResponsePagination struct {

	// next_key is the key to be passed to PageRequest.key to
	// query the next page most efficiently
	// Format: byte
	NextKey strfmt.Base64 `json:"next_key,omitempty"`

	// total is total number of results available if PageRequest.count_total
	// was set, its value is undefined otherwise
	Total string `json:"total,omitempty"`
}

// Validate validates this cosmos bank v1beta1 query total supply response pagination
func (m *CosmosBankV1beta1QueryTotalSupplyResponsePagination) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cosmos bank v1beta1 query total supply response pagination based on context it is used
func (m *CosmosBankV1beta1QueryTotalSupplyResponsePagination) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CosmosBankV1beta1QueryTotalSupplyResponsePagination) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosBankV1beta1QueryTotalSupplyResponsePagination) UnmarshalBinary(b []byte) error {
	var res CosmosBankV1beta1QueryTotalSupplyResponsePagination
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CosmosBankV1beta1QueryTotalSupplyResponseSupplyItems0 Coin defines a token with a denomination and an amount.
//
// NOTE: The amount field is an Int which implements the custom method
// signatures required by gogoproto.
//
// swagger:model CosmosBankV1beta1QueryTotalSupplyResponseSupplyItems0
type CosmosBankV1beta1QueryTotalSupplyResponseSupplyItems0 struct {

	// amount
	Amount string `json:"amount,omitempty"`

	// denom
	Denom string `json:"denom,omitempty"`
}

// Validate validates this cosmos bank v1beta1 query total supply response supply items0
func (m *CosmosBankV1beta1QueryTotalSupplyResponseSupplyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cosmos bank v1beta1 query total supply response supply items0 based on context it is used
func (m *CosmosBankV1beta1QueryTotalSupplyResponseSupplyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CosmosBankV1beta1QueryTotalSupplyResponseSupplyItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosBankV1beta1QueryTotalSupplyResponseSupplyItems0) UnmarshalBinary(b []byte) error {
	var res CosmosBankV1beta1QueryTotalSupplyResponseSupplyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

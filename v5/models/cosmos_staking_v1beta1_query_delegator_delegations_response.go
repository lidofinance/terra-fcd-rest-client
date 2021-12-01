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

// CosmosStakingV1beta1QueryDelegatorDelegationsResponse QueryDelegatorDelegationsResponse is response type for the
// Query/DelegatorDelegations RPC method.
//
// swagger:model cosmos.staking.v1beta1.QueryDelegatorDelegationsResponse
type CosmosStakingV1beta1QueryDelegatorDelegationsResponse struct {

	// delegation_responses defines all the delegations' info of a delegator.
	DelegationResponses []*CosmosStakingV1beta1QueryDelegatorDelegationsResponseDelegationResponsesItems0 `json:"delegation_responses"`

	// pagination
	Pagination *CosmosStakingV1beta1QueryDelegatorDelegationsResponsePagination `json:"pagination,omitempty"`
}

// Validate validates this cosmos staking v1beta1 query delegator delegations response
func (m *CosmosStakingV1beta1QueryDelegatorDelegationsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDelegationResponses(formats); err != nil {
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

func (m *CosmosStakingV1beta1QueryDelegatorDelegationsResponse) validateDelegationResponses(formats strfmt.Registry) error {
	if swag.IsZero(m.DelegationResponses) { // not required
		return nil
	}

	for i := 0; i < len(m.DelegationResponses); i++ {
		if swag.IsZero(m.DelegationResponses[i]) { // not required
			continue
		}

		if m.DelegationResponses[i] != nil {
			if err := m.DelegationResponses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("delegation_responses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("delegation_responses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CosmosStakingV1beta1QueryDelegatorDelegationsResponse) validatePagination(formats strfmt.Registry) error {
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

// ContextValidate validate this cosmos staking v1beta1 query delegator delegations response based on the context it is used
func (m *CosmosStakingV1beta1QueryDelegatorDelegationsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDelegationResponses(ctx, formats); err != nil {
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

func (m *CosmosStakingV1beta1QueryDelegatorDelegationsResponse) contextValidateDelegationResponses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DelegationResponses); i++ {

		if m.DelegationResponses[i] != nil {
			if err := m.DelegationResponses[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("delegation_responses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("delegation_responses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CosmosStakingV1beta1QueryDelegatorDelegationsResponse) contextValidatePagination(ctx context.Context, formats strfmt.Registry) error {

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
func (m *CosmosStakingV1beta1QueryDelegatorDelegationsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosStakingV1beta1QueryDelegatorDelegationsResponse) UnmarshalBinary(b []byte) error {
	var res CosmosStakingV1beta1QueryDelegatorDelegationsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CosmosStakingV1beta1QueryDelegatorDelegationsResponseDelegationResponsesItems0 DelegationResponse is equivalent to Delegation except that it contains a
// balance in addition to shares which is more suitable for client responses.
//
// swagger:model CosmosStakingV1beta1QueryDelegatorDelegationsResponseDelegationResponsesItems0
type CosmosStakingV1beta1QueryDelegatorDelegationsResponseDelegationResponsesItems0 struct {

	// balance
	Balance *CosmosStakingV1beta1QueryDelegatorDelegationsResponseDelegationResponsesItems0Balance `json:"balance,omitempty"`

	// delegation
	Delegation *CosmosStakingV1beta1QueryDelegatorDelegationsResponseDelegationResponsesItems0Delegation `json:"delegation,omitempty"`
}

// Validate validates this cosmos staking v1beta1 query delegator delegations response delegation responses items0
func (m *CosmosStakingV1beta1QueryDelegatorDelegationsResponseDelegationResponsesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBalance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDelegation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosStakingV1beta1QueryDelegatorDelegationsResponseDelegationResponsesItems0) validateBalance(formats strfmt.Registry) error {
	if swag.IsZero(m.Balance) { // not required
		return nil
	}

	if m.Balance != nil {
		if err := m.Balance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("balance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("balance")
			}
			return err
		}
	}

	return nil
}

func (m *CosmosStakingV1beta1QueryDelegatorDelegationsResponseDelegationResponsesItems0) validateDelegation(formats strfmt.Registry) error {
	if swag.IsZero(m.Delegation) { // not required
		return nil
	}

	if m.Delegation != nil {
		if err := m.Delegation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("delegation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("delegation")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cosmos staking v1beta1 query delegator delegations response delegation responses items0 based on the context it is used
func (m *CosmosStakingV1beta1QueryDelegatorDelegationsResponseDelegationResponsesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBalance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDelegation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosStakingV1beta1QueryDelegatorDelegationsResponseDelegationResponsesItems0) contextValidateBalance(ctx context.Context, formats strfmt.Registry) error {

	if m.Balance != nil {
		if err := m.Balance.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("balance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("balance")
			}
			return err
		}
	}

	return nil
}

func (m *CosmosStakingV1beta1QueryDelegatorDelegationsResponseDelegationResponsesItems0) contextValidateDelegation(ctx context.Context, formats strfmt.Registry) error {

	if m.Delegation != nil {
		if err := m.Delegation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("delegation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("delegation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CosmosStakingV1beta1QueryDelegatorDelegationsResponseDelegationResponsesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosStakingV1beta1QueryDelegatorDelegationsResponseDelegationResponsesItems0) UnmarshalBinary(b []byte) error {
	var res CosmosStakingV1beta1QueryDelegatorDelegationsResponseDelegationResponsesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CosmosStakingV1beta1QueryDelegatorDelegationsResponseDelegationResponsesItems0Balance Coin defines a token with a denomination and an amount.
//
// NOTE: The amount field is an Int which implements the custom method
// signatures required by gogoproto.
//
// swagger:model CosmosStakingV1beta1QueryDelegatorDelegationsResponseDelegationResponsesItems0Balance
type CosmosStakingV1beta1QueryDelegatorDelegationsResponseDelegationResponsesItems0Balance struct {

	// amount
	Amount string `json:"amount,omitempty"`

	// denom
	Denom string `json:"denom,omitempty"`
}

// Validate validates this cosmos staking v1beta1 query delegator delegations response delegation responses items0 balance
func (m *CosmosStakingV1beta1QueryDelegatorDelegationsResponseDelegationResponsesItems0Balance) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cosmos staking v1beta1 query delegator delegations response delegation responses items0 balance based on context it is used
func (m *CosmosStakingV1beta1QueryDelegatorDelegationsResponseDelegationResponsesItems0Balance) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CosmosStakingV1beta1QueryDelegatorDelegationsResponseDelegationResponsesItems0Balance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosStakingV1beta1QueryDelegatorDelegationsResponseDelegationResponsesItems0Balance) UnmarshalBinary(b []byte) error {
	var res CosmosStakingV1beta1QueryDelegatorDelegationsResponseDelegationResponsesItems0Balance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CosmosStakingV1beta1QueryDelegatorDelegationsResponseDelegationResponsesItems0Delegation Delegation represents the bond with tokens held by an account. It is
// owned by one delegator, and is associated with the voting power of one
// validator.
//
// swagger:model CosmosStakingV1beta1QueryDelegatorDelegationsResponseDelegationResponsesItems0Delegation
type CosmosStakingV1beta1QueryDelegatorDelegationsResponseDelegationResponsesItems0Delegation struct {

	// delegator_address is the bech32-encoded address of the delegator.
	DelegatorAddress string `json:"delegator_address,omitempty"`

	// shares define the delegation shares received.
	Shares string `json:"shares,omitempty"`

	// validator_address is the bech32-encoded address of the validator.
	ValidatorAddress string `json:"validator_address,omitempty"`
}

// Validate validates this cosmos staking v1beta1 query delegator delegations response delegation responses items0 delegation
func (m *CosmosStakingV1beta1QueryDelegatorDelegationsResponseDelegationResponsesItems0Delegation) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cosmos staking v1beta1 query delegator delegations response delegation responses items0 delegation based on context it is used
func (m *CosmosStakingV1beta1QueryDelegatorDelegationsResponseDelegationResponsesItems0Delegation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CosmosStakingV1beta1QueryDelegatorDelegationsResponseDelegationResponsesItems0Delegation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosStakingV1beta1QueryDelegatorDelegationsResponseDelegationResponsesItems0Delegation) UnmarshalBinary(b []byte) error {
	var res CosmosStakingV1beta1QueryDelegatorDelegationsResponseDelegationResponsesItems0Delegation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CosmosStakingV1beta1QueryDelegatorDelegationsResponsePagination pagination defines the pagination in the response.
//
// swagger:model CosmosStakingV1beta1QueryDelegatorDelegationsResponsePagination
type CosmosStakingV1beta1QueryDelegatorDelegationsResponsePagination struct {

	// next_key is the key to be passed to PageRequest.key to
	// query the next page most efficiently
	// Format: byte
	NextKey strfmt.Base64 `json:"next_key,omitempty"`

	// total is total number of results available if PageRequest.count_total
	// was set, its value is undefined otherwise
	Total string `json:"total,omitempty"`
}

// Validate validates this cosmos staking v1beta1 query delegator delegations response pagination
func (m *CosmosStakingV1beta1QueryDelegatorDelegationsResponsePagination) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cosmos staking v1beta1 query delegator delegations response pagination based on context it is used
func (m *CosmosStakingV1beta1QueryDelegatorDelegationsResponsePagination) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CosmosStakingV1beta1QueryDelegatorDelegationsResponsePagination) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosStakingV1beta1QueryDelegatorDelegationsResponsePagination) UnmarshalBinary(b []byte) error {
	var res CosmosStakingV1beta1QueryDelegatorDelegationsResponsePagination
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
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

// Redelegation redelegation
//
// swagger:model Redelegation
type Redelegation struct {

	// delegator address
	DelegatorAddress string `json:"delegator_address,omitempty"`

	// entries
	Entries []*RedelegationEntriesItems0 `json:"entries"`

	// validator dst address
	ValidatorDstAddress string `json:"validator_dst_address,omitempty"`

	// validator src address
	ValidatorSrcAddress string `json:"validator_src_address,omitempty"`
}

// Validate validates this redelegation
func (m *Redelegation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntries(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Redelegation) validateEntries(formats strfmt.Registry) error {
	if swag.IsZero(m.Entries) { // not required
		return nil
	}

	for i := 0; i < len(m.Entries); i++ {
		if swag.IsZero(m.Entries[i]) { // not required
			continue
		}

		if m.Entries[i] != nil {
			if err := m.Entries[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("entries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("entries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this redelegation based on the context it is used
func (m *Redelegation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEntries(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Redelegation) contextValidateEntries(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Entries); i++ {

		if m.Entries[i] != nil {
			if err := m.Entries[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("entries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("entries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Redelegation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Redelegation) UnmarshalBinary(b []byte) error {
	var res Redelegation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// RedelegationEntriesItems0 redelegation entries items0
//
// swagger:model RedelegationEntriesItems0
type RedelegationEntriesItems0 struct {

	// balance
	Balance string `json:"balance,omitempty"`

	// completion time
	CompletionTime int64 `json:"completion_time,omitempty"`

	// creation height
	CreationHeight int64 `json:"creation_height,omitempty"`

	// initial balance
	InitialBalance string `json:"initial_balance,omitempty"`

	// shares dst
	SharesDst string `json:"shares_dst,omitempty"`
}

// Validate validates this redelegation entries items0
func (m *RedelegationEntriesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this redelegation entries items0 based on context it is used
func (m *RedelegationEntriesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RedelegationEntriesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RedelegationEntriesItems0) UnmarshalBinary(b []byte) error {
	var res RedelegationEntriesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

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

// CosmosStakingV1beta1RedelegationResponse RedelegationResponse is equivalent to a Redelegation except that its entries
// contain a balance in addition to shares which is more suitable for client
// responses.
//
// swagger:model cosmos.staking.v1beta1.RedelegationResponse
type CosmosStakingV1beta1RedelegationResponse struct {

	// entries
	Entries []*CosmosStakingV1beta1RedelegationResponseEntriesItems0 `json:"entries"`

	// redelegation
	Redelegation *CosmosStakingV1beta1RedelegationResponseRedelegation `json:"redelegation,omitempty"`
}

// Validate validates this cosmos staking v1beta1 redelegation response
func (m *CosmosStakingV1beta1RedelegationResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntries(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRedelegation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosStakingV1beta1RedelegationResponse) validateEntries(formats strfmt.Registry) error {
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

func (m *CosmosStakingV1beta1RedelegationResponse) validateRedelegation(formats strfmt.Registry) error {
	if swag.IsZero(m.Redelegation) { // not required
		return nil
	}

	if m.Redelegation != nil {
		if err := m.Redelegation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("redelegation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("redelegation")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cosmos staking v1beta1 redelegation response based on the context it is used
func (m *CosmosStakingV1beta1RedelegationResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEntries(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRedelegation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosStakingV1beta1RedelegationResponse) contextValidateEntries(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CosmosStakingV1beta1RedelegationResponse) contextValidateRedelegation(ctx context.Context, formats strfmt.Registry) error {

	if m.Redelegation != nil {
		if err := m.Redelegation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("redelegation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("redelegation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CosmosStakingV1beta1RedelegationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosStakingV1beta1RedelegationResponse) UnmarshalBinary(b []byte) error {
	var res CosmosStakingV1beta1RedelegationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CosmosStakingV1beta1RedelegationResponseEntriesItems0 RedelegationEntryResponse is equivalent to a RedelegationEntry except that it
// contains a balance in addition to shares which is more suitable for client
// responses.
//
// swagger:model CosmosStakingV1beta1RedelegationResponseEntriesItems0
type CosmosStakingV1beta1RedelegationResponseEntriesItems0 struct {

	// balance
	Balance string `json:"balance,omitempty"`

	// redelegation entry
	RedelegationEntry *CosmosStakingV1beta1RedelegationResponseEntriesItems0RedelegationEntry `json:"redelegation_entry,omitempty"`
}

// Validate validates this cosmos staking v1beta1 redelegation response entries items0
func (m *CosmosStakingV1beta1RedelegationResponseEntriesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRedelegationEntry(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosStakingV1beta1RedelegationResponseEntriesItems0) validateRedelegationEntry(formats strfmt.Registry) error {
	if swag.IsZero(m.RedelegationEntry) { // not required
		return nil
	}

	if m.RedelegationEntry != nil {
		if err := m.RedelegationEntry.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("redelegation_entry")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("redelegation_entry")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cosmos staking v1beta1 redelegation response entries items0 based on the context it is used
func (m *CosmosStakingV1beta1RedelegationResponseEntriesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRedelegationEntry(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosStakingV1beta1RedelegationResponseEntriesItems0) contextValidateRedelegationEntry(ctx context.Context, formats strfmt.Registry) error {

	if m.RedelegationEntry != nil {
		if err := m.RedelegationEntry.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("redelegation_entry")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("redelegation_entry")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CosmosStakingV1beta1RedelegationResponseEntriesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosStakingV1beta1RedelegationResponseEntriesItems0) UnmarshalBinary(b []byte) error {
	var res CosmosStakingV1beta1RedelegationResponseEntriesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CosmosStakingV1beta1RedelegationResponseEntriesItems0RedelegationEntry RedelegationEntry defines a redelegation object with relevant metadata.
//
// swagger:model CosmosStakingV1beta1RedelegationResponseEntriesItems0RedelegationEntry
type CosmosStakingV1beta1RedelegationResponseEntriesItems0RedelegationEntry struct {

	// completion_time defines the unix time for redelegation completion.
	// Format: date-time
	CompletionTime strfmt.DateTime `json:"completion_time,omitempty"`

	// creation_height  defines the height which the redelegation took place.
	CreationHeight string `json:"creation_height,omitempty"`

	// initial_balance defines the initial balance when redelegation started.
	InitialBalance string `json:"initial_balance,omitempty"`

	// shares_dst is the amount of destination-validator shares created by redelegation.
	SharesDst string `json:"shares_dst,omitempty"`
}

// Validate validates this cosmos staking v1beta1 redelegation response entries items0 redelegation entry
func (m *CosmosStakingV1beta1RedelegationResponseEntriesItems0RedelegationEntry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompletionTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosStakingV1beta1RedelegationResponseEntriesItems0RedelegationEntry) validateCompletionTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CompletionTime) { // not required
		return nil
	}

	if err := validate.FormatOf("redelegation_entry"+"."+"completion_time", "body", "date-time", m.CompletionTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this cosmos staking v1beta1 redelegation response entries items0 redelegation entry based on context it is used
func (m *CosmosStakingV1beta1RedelegationResponseEntriesItems0RedelegationEntry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CosmosStakingV1beta1RedelegationResponseEntriesItems0RedelegationEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosStakingV1beta1RedelegationResponseEntriesItems0RedelegationEntry) UnmarshalBinary(b []byte) error {
	var res CosmosStakingV1beta1RedelegationResponseEntriesItems0RedelegationEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CosmosStakingV1beta1RedelegationResponseRedelegation Redelegation contains the list of a particular delegator's redelegating bonds
// from a particular source validator to a particular destination validator.
//
// swagger:model CosmosStakingV1beta1RedelegationResponseRedelegation
type CosmosStakingV1beta1RedelegationResponseRedelegation struct {

	// delegator_address is the bech32-encoded address of the delegator.
	DelegatorAddress string `json:"delegator_address,omitempty"`

	// entries are the redelegation entries.
	Entries []*CosmosStakingV1beta1RedelegationResponseRedelegationEntriesItems0 `json:"entries"`

	// validator_dst_address is the validator redelegation destination operator address.
	ValidatorDstAddress string `json:"validator_dst_address,omitempty"`

	// validator_src_address is the validator redelegation source operator address.
	ValidatorSrcAddress string `json:"validator_src_address,omitempty"`
}

// Validate validates this cosmos staking v1beta1 redelegation response redelegation
func (m *CosmosStakingV1beta1RedelegationResponseRedelegation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntries(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosStakingV1beta1RedelegationResponseRedelegation) validateEntries(formats strfmt.Registry) error {
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
					return ve.ValidateName("redelegation" + "." + "entries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("redelegation" + "." + "entries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this cosmos staking v1beta1 redelegation response redelegation based on the context it is used
func (m *CosmosStakingV1beta1RedelegationResponseRedelegation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEntries(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosStakingV1beta1RedelegationResponseRedelegation) contextValidateEntries(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Entries); i++ {

		if m.Entries[i] != nil {
			if err := m.Entries[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("redelegation" + "." + "entries" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("redelegation" + "." + "entries" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CosmosStakingV1beta1RedelegationResponseRedelegation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosStakingV1beta1RedelegationResponseRedelegation) UnmarshalBinary(b []byte) error {
	var res CosmosStakingV1beta1RedelegationResponseRedelegation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CosmosStakingV1beta1RedelegationResponseRedelegationEntriesItems0 RedelegationEntry defines a redelegation object with relevant metadata.
//
// swagger:model CosmosStakingV1beta1RedelegationResponseRedelegationEntriesItems0
type CosmosStakingV1beta1RedelegationResponseRedelegationEntriesItems0 struct {

	// completion_time defines the unix time for redelegation completion.
	// Format: date-time
	CompletionTime strfmt.DateTime `json:"completion_time,omitempty"`

	// creation_height  defines the height which the redelegation took place.
	CreationHeight string `json:"creation_height,omitempty"`

	// initial_balance defines the initial balance when redelegation started.
	InitialBalance string `json:"initial_balance,omitempty"`

	// shares_dst is the amount of destination-validator shares created by redelegation.
	SharesDst string `json:"shares_dst,omitempty"`
}

// Validate validates this cosmos staking v1beta1 redelegation response redelegation entries items0
func (m *CosmosStakingV1beta1RedelegationResponseRedelegationEntriesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompletionTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosStakingV1beta1RedelegationResponseRedelegationEntriesItems0) validateCompletionTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CompletionTime) { // not required
		return nil
	}

	if err := validate.FormatOf("completion_time", "body", "date-time", m.CompletionTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this cosmos staking v1beta1 redelegation response redelegation entries items0 based on context it is used
func (m *CosmosStakingV1beta1RedelegationResponseRedelegationEntriesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CosmosStakingV1beta1RedelegationResponseRedelegationEntriesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosStakingV1beta1RedelegationResponseRedelegationEntriesItems0) UnmarshalBinary(b []byte) error {
	var res CosmosStakingV1beta1RedelegationResponseRedelegationEntriesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

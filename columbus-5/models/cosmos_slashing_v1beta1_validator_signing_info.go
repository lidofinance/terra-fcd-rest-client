// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CosmosSlashingV1beta1ValidatorSigningInfo ValidatorSigningInfo defines a validator's signing info for monitoring their
// liveness activity.
//
// swagger:model cosmos.slashing.v1beta1.ValidatorSigningInfo
type CosmosSlashingV1beta1ValidatorSigningInfo struct {

	// address
	Address string `json:"address,omitempty"`

	// Index which is incremented each time the validator was a bonded
	// in a block and may have signed a precommit or not. This in conjunction with the
	// `SignedBlocksWindow` param determines the index in the `MissedBlocksBitArray`.
	IndexOffset string `json:"index_offset,omitempty"`

	// Timestamp until which the validator is jailed due to liveness downtime.
	// Format: date-time
	JailedUntil strfmt.DateTime `json:"jailed_until,omitempty"`

	// A counter kept to avoid unnecessary array reads.
	// Note that `Sum(MissedBlocksBitArray)` always equals `MissedBlocksCounter`.
	MissedBlocksCounter string `json:"missed_blocks_counter,omitempty"`

	// Height at which validator was first a candidate OR was unjailed
	StartHeight string `json:"start_height,omitempty"`

	// Whether or not a validator has been tombstoned (killed out of validator set). It is set
	// once the validator commits an equivocation or for any other configured misbehiavor.
	Tombstoned bool `json:"tombstoned,omitempty"`
}

// Validate validates this cosmos slashing v1beta1 validator signing info
func (m *CosmosSlashingV1beta1ValidatorSigningInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateJailedUntil(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosSlashingV1beta1ValidatorSigningInfo) validateJailedUntil(formats strfmt.Registry) error {
	if swag.IsZero(m.JailedUntil) { // not required
		return nil
	}

	if err := validate.FormatOf("jailed_until", "body", "date-time", m.JailedUntil.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this cosmos slashing v1beta1 validator signing info based on context it is used
func (m *CosmosSlashingV1beta1ValidatorSigningInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CosmosSlashingV1beta1ValidatorSigningInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosSlashingV1beta1ValidatorSigningInfo) UnmarshalBinary(b []byte) error {
	var res CosmosSlashingV1beta1ValidatorSigningInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

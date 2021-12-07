// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IbcCoreChannelV1QueryPacketCommitmentResponse QueryPacketCommitmentResponse defines the client query response for a packet
// which also includes a proof and the height from which the proof was
// retrieved
//
// swagger:model ibc.core.channel.v1.QueryPacketCommitmentResponse
type IbcCoreChannelV1QueryPacketCommitmentResponse struct {

	// packet associated with the request fields
	// Format: byte
	Commitment strfmt.Base64 `json:"commitment,omitempty"`

	// merkle proof of existence
	// Format: byte
	Proof strfmt.Base64 `json:"proof,omitempty"`

	// proof height
	ProofHeight *IbcCoreChannelV1QueryPacketCommitmentResponseProofHeight `json:"proof_height,omitempty"`
}

// Validate validates this ibc core channel v1 query packet commitment response
func (m *IbcCoreChannelV1QueryPacketCommitmentResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProofHeight(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IbcCoreChannelV1QueryPacketCommitmentResponse) validateProofHeight(formats strfmt.Registry) error {
	if swag.IsZero(m.ProofHeight) { // not required
		return nil
	}

	if m.ProofHeight != nil {
		if err := m.ProofHeight.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proof_height")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("proof_height")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ibc core channel v1 query packet commitment response based on the context it is used
func (m *IbcCoreChannelV1QueryPacketCommitmentResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateProofHeight(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IbcCoreChannelV1QueryPacketCommitmentResponse) contextValidateProofHeight(ctx context.Context, formats strfmt.Registry) error {

	if m.ProofHeight != nil {
		if err := m.ProofHeight.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proof_height")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("proof_height")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IbcCoreChannelV1QueryPacketCommitmentResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IbcCoreChannelV1QueryPacketCommitmentResponse) UnmarshalBinary(b []byte) error {
	var res IbcCoreChannelV1QueryPacketCommitmentResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IbcCoreChannelV1QueryPacketCommitmentResponseProofHeight height at which the proof was retrieved
//
// Normally the RevisionHeight is incremented at each height while keeping
// RevisionNumber the same. However some consensus algorithms may choose to
// reset the height in certain conditions e.g. hard forks, state-machine
// breaking changes In these cases, the RevisionNumber is incremented so that
// height continues to be monitonically increasing even as the RevisionHeight
// gets reset
//
// swagger:model IbcCoreChannelV1QueryPacketCommitmentResponseProofHeight
type IbcCoreChannelV1QueryPacketCommitmentResponseProofHeight struct {

	// the height within the given revision
	RevisionHeight string `json:"revision_height,omitempty"`

	// the revision that the client is currently on
	RevisionNumber string `json:"revision_number,omitempty"`
}

// Validate validates this ibc core channel v1 query packet commitment response proof height
func (m *IbcCoreChannelV1QueryPacketCommitmentResponseProofHeight) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ibc core channel v1 query packet commitment response proof height based on context it is used
func (m *IbcCoreChannelV1QueryPacketCommitmentResponseProofHeight) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IbcCoreChannelV1QueryPacketCommitmentResponseProofHeight) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IbcCoreChannelV1QueryPacketCommitmentResponseProofHeight) UnmarshalBinary(b []byte) error {
	var res IbcCoreChannelV1QueryPacketCommitmentResponseProofHeight
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

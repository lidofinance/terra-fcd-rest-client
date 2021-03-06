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

// BlockHeader block header
//
// swagger:model BlockHeader
type BlockHeader struct {

	// app hash
	// Example: EE5F3404034C524501629B56E0DDC38FAD651F04
	AppHash string `json:"app_hash,omitempty"`

	// chain id
	// Example: columbus-5
	ChainID string `json:"chain_id,omitempty"`

	// consensus hash
	// Example: EE5F3404034C524501629B56E0DDC38FAD651F04
	ConsensusHash string `json:"consensus_hash,omitempty"`

	// data hash
	// Example: EE5F3404034C524501629B56E0DDC38FAD651F04
	DataHash string `json:"data_hash,omitempty"`

	// evidence hash
	// Example: EE5F3404034C524501629B56E0DDC38FAD651F04
	EvidenceHash string `json:"evidence_hash,omitempty"`

	// height
	// Example: 1
	Height string `json:"height,omitempty"`

	// last block id
	LastBlockID *BlockID `json:"last_block_id,omitempty"`

	// last commit hash
	// Example: EE5F3404034C524501629B56E0DDC38FAD651F04
	LastCommitHash string `json:"last_commit_hash,omitempty"`

	// last results hash
	// Example: EE5F3404034C524501629B56E0DDC38FAD651F04
	LastResultsHash string `json:"last_results_hash,omitempty"`

	// next validators hash
	// Example: EE5F3404034C524501629B56E0DDC38FAD651F04
	NextValidatorsHash string `json:"next_validators_hash,omitempty"`

	// bech32 encoded address
	// Example: terra1wg2mlrxdmnnkkykgqg4znky86nyrtc45q336yv
	ProposerAddress string `json:"proposer_address,omitempty"`

	// time
	// Example: 2017-12-30T05:53:09.287+01:00
	Time string `json:"time,omitempty"`

	// validators hash
	// Example: EE5F3404034C524501629B56E0DDC38FAD651F04
	ValidatorsHash string `json:"validators_hash,omitempty"`

	// version
	Version *BlockHeaderVersion `json:"version,omitempty"`
}

// Validate validates this block header
func (m *BlockHeader) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastBlockID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BlockHeader) validateLastBlockID(formats strfmt.Registry) error {
	if swag.IsZero(m.LastBlockID) { // not required
		return nil
	}

	if m.LastBlockID != nil {
		if err := m.LastBlockID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("last_block_id")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("last_block_id")
			}
			return err
		}
	}

	return nil
}

func (m *BlockHeader) validateVersion(formats strfmt.Registry) error {
	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if m.Version != nil {
		if err := m.Version.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("version")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("version")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this block header based on the context it is used
func (m *BlockHeader) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLastBlockID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BlockHeader) contextValidateLastBlockID(ctx context.Context, formats strfmt.Registry) error {

	if m.LastBlockID != nil {
		if err := m.LastBlockID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("last_block_id")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("last_block_id")
			}
			return err
		}
	}

	return nil
}

func (m *BlockHeader) contextValidateVersion(ctx context.Context, formats strfmt.Registry) error {

	if m.Version != nil {
		if err := m.Version.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("version")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("version")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BlockHeader) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BlockHeader) UnmarshalBinary(b []byte) error {
	var res BlockHeader
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// BlockHeaderVersion block header version
//
// swagger:model BlockHeaderVersion
type BlockHeaderVersion struct {

	// app
	// Example: 0
	App string `json:"app,omitempty"`

	// block
	// Example: 10
	Block string `json:"block,omitempty"`
}

// Validate validates this block header version
func (m *BlockHeaderVersion) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this block header version based on context it is used
func (m *BlockHeaderVersion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BlockHeaderVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BlockHeaderVersion) UnmarshalBinary(b []byte) error {
	var res BlockHeaderVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TendermintTypesSignedHeader tendermint types signed header
//
// swagger:model tendermint.types.SignedHeader
type TendermintTypesSignedHeader struct {

	// commit
	Commit *TendermintTypesSignedHeaderCommit `json:"commit,omitempty"`

	// header
	Header *TendermintTypesSignedHeaderHeader `json:"header,omitempty"`
}

// Validate validates this tendermint types signed header
func (m *TendermintTypesSignedHeader) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHeader(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TendermintTypesSignedHeader) validateCommit(formats strfmt.Registry) error {
	if swag.IsZero(m.Commit) { // not required
		return nil
	}

	if m.Commit != nil {
		if err := m.Commit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("commit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("commit")
			}
			return err
		}
	}

	return nil
}

func (m *TendermintTypesSignedHeader) validateHeader(formats strfmt.Registry) error {
	if swag.IsZero(m.Header) { // not required
		return nil
	}

	if m.Header != nil {
		if err := m.Header.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("header")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("header")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this tendermint types signed header based on the context it is used
func (m *TendermintTypesSignedHeader) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCommit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHeader(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TendermintTypesSignedHeader) contextValidateCommit(ctx context.Context, formats strfmt.Registry) error {

	if m.Commit != nil {
		if err := m.Commit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("commit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("commit")
			}
			return err
		}
	}

	return nil
}

func (m *TendermintTypesSignedHeader) contextValidateHeader(ctx context.Context, formats strfmt.Registry) error {

	if m.Header != nil {
		if err := m.Header.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("header")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("header")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TendermintTypesSignedHeader) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TendermintTypesSignedHeader) UnmarshalBinary(b []byte) error {
	var res TendermintTypesSignedHeader
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TendermintTypesSignedHeaderCommit Commit contains the evidence that a block was committed by a set of validators.
//
// swagger:model TendermintTypesSignedHeaderCommit
type TendermintTypesSignedHeaderCommit struct {

	// block id
	BlockID *TendermintTypesSignedHeaderCommitBlockID `json:"block_id,omitempty"`

	// height
	Height string `json:"height,omitempty"`

	// round
	Round int32 `json:"round,omitempty"`

	// signatures
	Signatures []*TendermintTypesSignedHeaderCommitSignaturesItems0 `json:"signatures"`
}

// Validate validates this tendermint types signed header commit
func (m *TendermintTypesSignedHeaderCommit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBlockID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSignatures(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TendermintTypesSignedHeaderCommit) validateBlockID(formats strfmt.Registry) error {
	if swag.IsZero(m.BlockID) { // not required
		return nil
	}

	if m.BlockID != nil {
		if err := m.BlockID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("commit" + "." + "block_id")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("commit" + "." + "block_id")
			}
			return err
		}
	}

	return nil
}

func (m *TendermintTypesSignedHeaderCommit) validateSignatures(formats strfmt.Registry) error {
	if swag.IsZero(m.Signatures) { // not required
		return nil
	}

	for i := 0; i < len(m.Signatures); i++ {
		if swag.IsZero(m.Signatures[i]) { // not required
			continue
		}

		if m.Signatures[i] != nil {
			if err := m.Signatures[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("commit" + "." + "signatures" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("commit" + "." + "signatures" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this tendermint types signed header commit based on the context it is used
func (m *TendermintTypesSignedHeaderCommit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBlockID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSignatures(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TendermintTypesSignedHeaderCommit) contextValidateBlockID(ctx context.Context, formats strfmt.Registry) error {

	if m.BlockID != nil {
		if err := m.BlockID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("commit" + "." + "block_id")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("commit" + "." + "block_id")
			}
			return err
		}
	}

	return nil
}

func (m *TendermintTypesSignedHeaderCommit) contextValidateSignatures(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Signatures); i++ {

		if m.Signatures[i] != nil {
			if err := m.Signatures[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("commit" + "." + "signatures" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("commit" + "." + "signatures" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TendermintTypesSignedHeaderCommit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TendermintTypesSignedHeaderCommit) UnmarshalBinary(b []byte) error {
	var res TendermintTypesSignedHeaderCommit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TendermintTypesSignedHeaderCommitBlockID BlockID
//
// swagger:model TendermintTypesSignedHeaderCommitBlockID
type TendermintTypesSignedHeaderCommitBlockID struct {

	// hash
	// Format: byte
	Hash strfmt.Base64 `json:"hash,omitempty"`

	// part set header
	PartSetHeader *TendermintTypesSignedHeaderCommitBlockIDPartSetHeader `json:"part_set_header,omitempty"`
}

// Validate validates this tendermint types signed header commit block ID
func (m *TendermintTypesSignedHeaderCommitBlockID) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePartSetHeader(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TendermintTypesSignedHeaderCommitBlockID) validatePartSetHeader(formats strfmt.Registry) error {
	if swag.IsZero(m.PartSetHeader) { // not required
		return nil
	}

	if m.PartSetHeader != nil {
		if err := m.PartSetHeader.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("commit" + "." + "block_id" + "." + "part_set_header")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("commit" + "." + "block_id" + "." + "part_set_header")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this tendermint types signed header commit block ID based on the context it is used
func (m *TendermintTypesSignedHeaderCommitBlockID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePartSetHeader(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TendermintTypesSignedHeaderCommitBlockID) contextValidatePartSetHeader(ctx context.Context, formats strfmt.Registry) error {

	if m.PartSetHeader != nil {
		if err := m.PartSetHeader.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("commit" + "." + "block_id" + "." + "part_set_header")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("commit" + "." + "block_id" + "." + "part_set_header")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TendermintTypesSignedHeaderCommitBlockID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TendermintTypesSignedHeaderCommitBlockID) UnmarshalBinary(b []byte) error {
	var res TendermintTypesSignedHeaderCommitBlockID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TendermintTypesSignedHeaderCommitBlockIDPartSetHeader PartsetHeader
//
// swagger:model TendermintTypesSignedHeaderCommitBlockIDPartSetHeader
type TendermintTypesSignedHeaderCommitBlockIDPartSetHeader struct {

	// hash
	// Format: byte
	Hash strfmt.Base64 `json:"hash,omitempty"`

	// total
	Total int64 `json:"total,omitempty"`
}

// Validate validates this tendermint types signed header commit block ID part set header
func (m *TendermintTypesSignedHeaderCommitBlockIDPartSetHeader) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this tendermint types signed header commit block ID part set header based on context it is used
func (m *TendermintTypesSignedHeaderCommitBlockIDPartSetHeader) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TendermintTypesSignedHeaderCommitBlockIDPartSetHeader) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TendermintTypesSignedHeaderCommitBlockIDPartSetHeader) UnmarshalBinary(b []byte) error {
	var res TendermintTypesSignedHeaderCommitBlockIDPartSetHeader
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TendermintTypesSignedHeaderCommitSignaturesItems0 CommitSig is a part of the Vote included in a Commit.
//
// swagger:model TendermintTypesSignedHeaderCommitSignaturesItems0
type TendermintTypesSignedHeaderCommitSignaturesItems0 struct {

	// BlockIdFlag indicates which BlcokID the signature is for
	// Enum: [BLOCK_ID_FLAG_UNKNOWN BLOCK_ID_FLAG_ABSENT BLOCK_ID_FLAG_COMMIT BLOCK_ID_FLAG_NIL]
	BlockIDFlag *string `json:"block_id_flag,omitempty"`

	// signature
	// Format: byte
	Signature strfmt.Base64 `json:"signature,omitempty"`

	// timestamp
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`

	// validator address
	// Format: byte
	ValidatorAddress strfmt.Base64 `json:"validator_address,omitempty"`
}

// Validate validates this tendermint types signed header commit signatures items0
func (m *TendermintTypesSignedHeaderCommitSignaturesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBlockIDFlag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var tendermintTypesSignedHeaderCommitSignaturesItems0TypeBlockIDFlagPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["BLOCK_ID_FLAG_UNKNOWN","BLOCK_ID_FLAG_ABSENT","BLOCK_ID_FLAG_COMMIT","BLOCK_ID_FLAG_NIL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tendermintTypesSignedHeaderCommitSignaturesItems0TypeBlockIDFlagPropEnum = append(tendermintTypesSignedHeaderCommitSignaturesItems0TypeBlockIDFlagPropEnum, v)
	}
}

const (

	// TendermintTypesSignedHeaderCommitSignaturesItems0BlockIDFlagBLOCKIDFLAGUNKNOWN captures enum value "BLOCK_ID_FLAG_UNKNOWN"
	TendermintTypesSignedHeaderCommitSignaturesItems0BlockIDFlagBLOCKIDFLAGUNKNOWN string = "BLOCK_ID_FLAG_UNKNOWN"

	// TendermintTypesSignedHeaderCommitSignaturesItems0BlockIDFlagBLOCKIDFLAGABSENT captures enum value "BLOCK_ID_FLAG_ABSENT"
	TendermintTypesSignedHeaderCommitSignaturesItems0BlockIDFlagBLOCKIDFLAGABSENT string = "BLOCK_ID_FLAG_ABSENT"

	// TendermintTypesSignedHeaderCommitSignaturesItems0BlockIDFlagBLOCKIDFLAGCOMMIT captures enum value "BLOCK_ID_FLAG_COMMIT"
	TendermintTypesSignedHeaderCommitSignaturesItems0BlockIDFlagBLOCKIDFLAGCOMMIT string = "BLOCK_ID_FLAG_COMMIT"

	// TendermintTypesSignedHeaderCommitSignaturesItems0BlockIDFlagBLOCKIDFLAGNIL captures enum value "BLOCK_ID_FLAG_NIL"
	TendermintTypesSignedHeaderCommitSignaturesItems0BlockIDFlagBLOCKIDFLAGNIL string = "BLOCK_ID_FLAG_NIL"
)

// prop value enum
func (m *TendermintTypesSignedHeaderCommitSignaturesItems0) validateBlockIDFlagEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, tendermintTypesSignedHeaderCommitSignaturesItems0TypeBlockIDFlagPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TendermintTypesSignedHeaderCommitSignaturesItems0) validateBlockIDFlag(formats strfmt.Registry) error {
	if swag.IsZero(m.BlockIDFlag) { // not required
		return nil
	}

	// value enum
	if err := m.validateBlockIDFlagEnum("block_id_flag", "body", *m.BlockIDFlag); err != nil {
		return err
	}

	return nil
}

func (m *TendermintTypesSignedHeaderCommitSignaturesItems0) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this tendermint types signed header commit signatures items0 based on context it is used
func (m *TendermintTypesSignedHeaderCommitSignaturesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TendermintTypesSignedHeaderCommitSignaturesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TendermintTypesSignedHeaderCommitSignaturesItems0) UnmarshalBinary(b []byte) error {
	var res TendermintTypesSignedHeaderCommitSignaturesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TendermintTypesSignedHeaderHeader Header defines the structure of a Tendermint block header.
//
// swagger:model TendermintTypesSignedHeaderHeader
type TendermintTypesSignedHeaderHeader struct {

	// app hash
	// Format: byte
	AppHash strfmt.Base64 `json:"app_hash,omitempty"`

	// chain id
	ChainID string `json:"chain_id,omitempty"`

	// consensus hash
	// Format: byte
	ConsensusHash strfmt.Base64 `json:"consensus_hash,omitempty"`

	// data hash
	// Format: byte
	DataHash strfmt.Base64 `json:"data_hash,omitempty"`

	// consensus info
	// Format: byte
	EvidenceHash strfmt.Base64 `json:"evidence_hash,omitempty"`

	// height
	Height string `json:"height,omitempty"`

	// last block id
	LastBlockID *TendermintTypesSignedHeaderHeaderLastBlockID `json:"last_block_id,omitempty"`

	// hashes of block data
	// Format: byte
	LastCommitHash strfmt.Base64 `json:"last_commit_hash,omitempty"`

	// last results hash
	// Format: byte
	LastResultsHash strfmt.Base64 `json:"last_results_hash,omitempty"`

	// next validators hash
	// Format: byte
	NextValidatorsHash strfmt.Base64 `json:"next_validators_hash,omitempty"`

	// proposer address
	// Format: byte
	ProposerAddress strfmt.Base64 `json:"proposer_address,omitempty"`

	// time
	// Format: date-time
	Time strfmt.DateTime `json:"time,omitempty"`

	// hashes from the app output from the prev block
	// Format: byte
	ValidatorsHash strfmt.Base64 `json:"validators_hash,omitempty"`

	// version
	Version *TendermintTypesSignedHeaderHeaderVersion `json:"version,omitempty"`
}

// Validate validates this tendermint types signed header header
func (m *TendermintTypesSignedHeaderHeader) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastBlockID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTime(formats); err != nil {
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

func (m *TendermintTypesSignedHeaderHeader) validateLastBlockID(formats strfmt.Registry) error {
	if swag.IsZero(m.LastBlockID) { // not required
		return nil
	}

	if m.LastBlockID != nil {
		if err := m.LastBlockID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("header" + "." + "last_block_id")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("header" + "." + "last_block_id")
			}
			return err
		}
	}

	return nil
}

func (m *TendermintTypesSignedHeaderHeader) validateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.Time) { // not required
		return nil
	}

	if err := validate.FormatOf("header"+"."+"time", "body", "date-time", m.Time.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TendermintTypesSignedHeaderHeader) validateVersion(formats strfmt.Registry) error {
	if swag.IsZero(m.Version) { // not required
		return nil
	}

	if m.Version != nil {
		if err := m.Version.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("header" + "." + "version")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("header" + "." + "version")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this tendermint types signed header header based on the context it is used
func (m *TendermintTypesSignedHeaderHeader) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *TendermintTypesSignedHeaderHeader) contextValidateLastBlockID(ctx context.Context, formats strfmt.Registry) error {

	if m.LastBlockID != nil {
		if err := m.LastBlockID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("header" + "." + "last_block_id")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("header" + "." + "last_block_id")
			}
			return err
		}
	}

	return nil
}

func (m *TendermintTypesSignedHeaderHeader) contextValidateVersion(ctx context.Context, formats strfmt.Registry) error {

	if m.Version != nil {
		if err := m.Version.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("header" + "." + "version")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("header" + "." + "version")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TendermintTypesSignedHeaderHeader) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TendermintTypesSignedHeaderHeader) UnmarshalBinary(b []byte) error {
	var res TendermintTypesSignedHeaderHeader
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TendermintTypesSignedHeaderHeaderLastBlockID BlockID
//
// swagger:model TendermintTypesSignedHeaderHeaderLastBlockID
type TendermintTypesSignedHeaderHeaderLastBlockID struct {

	// hash
	// Format: byte
	Hash strfmt.Base64 `json:"hash,omitempty"`

	// part set header
	PartSetHeader *TendermintTypesSignedHeaderHeaderLastBlockIDPartSetHeader `json:"part_set_header,omitempty"`
}

// Validate validates this tendermint types signed header header last block ID
func (m *TendermintTypesSignedHeaderHeaderLastBlockID) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePartSetHeader(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TendermintTypesSignedHeaderHeaderLastBlockID) validatePartSetHeader(formats strfmt.Registry) error {
	if swag.IsZero(m.PartSetHeader) { // not required
		return nil
	}

	if m.PartSetHeader != nil {
		if err := m.PartSetHeader.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("header" + "." + "last_block_id" + "." + "part_set_header")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("header" + "." + "last_block_id" + "." + "part_set_header")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this tendermint types signed header header last block ID based on the context it is used
func (m *TendermintTypesSignedHeaderHeaderLastBlockID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePartSetHeader(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TendermintTypesSignedHeaderHeaderLastBlockID) contextValidatePartSetHeader(ctx context.Context, formats strfmt.Registry) error {

	if m.PartSetHeader != nil {
		if err := m.PartSetHeader.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("header" + "." + "last_block_id" + "." + "part_set_header")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("header" + "." + "last_block_id" + "." + "part_set_header")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TendermintTypesSignedHeaderHeaderLastBlockID) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TendermintTypesSignedHeaderHeaderLastBlockID) UnmarshalBinary(b []byte) error {
	var res TendermintTypesSignedHeaderHeaderLastBlockID
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TendermintTypesSignedHeaderHeaderLastBlockIDPartSetHeader PartsetHeader
//
// swagger:model TendermintTypesSignedHeaderHeaderLastBlockIDPartSetHeader
type TendermintTypesSignedHeaderHeaderLastBlockIDPartSetHeader struct {

	// hash
	// Format: byte
	Hash strfmt.Base64 `json:"hash,omitempty"`

	// total
	Total int64 `json:"total,omitempty"`
}

// Validate validates this tendermint types signed header header last block ID part set header
func (m *TendermintTypesSignedHeaderHeaderLastBlockIDPartSetHeader) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this tendermint types signed header header last block ID part set header based on context it is used
func (m *TendermintTypesSignedHeaderHeaderLastBlockIDPartSetHeader) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TendermintTypesSignedHeaderHeaderLastBlockIDPartSetHeader) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TendermintTypesSignedHeaderHeaderLastBlockIDPartSetHeader) UnmarshalBinary(b []byte) error {
	var res TendermintTypesSignedHeaderHeaderLastBlockIDPartSetHeader
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TendermintTypesSignedHeaderHeaderVersion basic block info
//
// Consensus captures the consensus rules for processing a block in the blockchain,
// including all blockchain data structures and the rules of the application's
// state transition machine.
//
// swagger:model TendermintTypesSignedHeaderHeaderVersion
type TendermintTypesSignedHeaderHeaderVersion struct {

	// app
	App string `json:"app,omitempty"`

	// block
	Block string `json:"block,omitempty"`
}

// Validate validates this tendermint types signed header header version
func (m *TendermintTypesSignedHeaderHeaderVersion) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this tendermint types signed header header version based on context it is used
func (m *TendermintTypesSignedHeaderHeaderVersion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TendermintTypesSignedHeaderHeaderVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TendermintTypesSignedHeaderHeaderVersion) UnmarshalBinary(b []byte) error {
	var res TendermintTypesSignedHeaderHeaderVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
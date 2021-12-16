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

// BlockQuery block query
//
// swagger:model BlockQuery
type BlockQuery struct {

	// block
	Block *Block `json:"block,omitempty"`

	// block id
	BlockID *BlockID `json:"block_id,omitempty"`
}

// Validate validates this block query
func (m *BlockQuery) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBlock(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBlockID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BlockQuery) validateBlock(formats strfmt.Registry) error {
	if swag.IsZero(m.Block) { // not required
		return nil
	}

	if m.Block != nil {
		if err := m.Block.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("block")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("block")
			}
			return err
		}
	}

	return nil
}

func (m *BlockQuery) validateBlockID(formats strfmt.Registry) error {
	if swag.IsZero(m.BlockID) { // not required
		return nil
	}

	if m.BlockID != nil {
		if err := m.BlockID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("block_id")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("block_id")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this block query based on the context it is used
func (m *BlockQuery) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBlock(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBlockID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BlockQuery) contextValidateBlock(ctx context.Context, formats strfmt.Registry) error {

	if m.Block != nil {
		if err := m.Block.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("block")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("block")
			}
			return err
		}
	}

	return nil
}

func (m *BlockQuery) contextValidateBlockID(ctx context.Context, formats strfmt.Registry) error {

	if m.BlockID != nil {
		if err := m.BlockID.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("block_id")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("block_id")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BlockQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BlockQuery) UnmarshalBinary(b []byte) error {
	var res BlockQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

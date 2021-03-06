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

// GetProposalListResultProposals get proposal list result proposals
//
// swagger:model getProposalListResult.proposals
type GetProposalListResultProposals struct {

	// deposit
	// Required: true
	Deposit *GetProposalListResultProposalsDeposit `json:"deposit"`

	// description
	// Required: true
	Description *string `json:"description"`

	// id
	// Required: true
	ID *string `json:"id"`

	// Proposer information
	// Required: true
	Proposer *GetProposalListResultProposalsProposer `json:"proposer"`

	// Proposal status
	// Required: true
	Status *string `json:"status"`

	// submit time
	// Required: true
	SubmitTime *string `json:"submitTime"`

	// title
	// Required: true
	Title *string `json:"title"`

	// Proposal type
	// Required: true
	Type *string `json:"type"`

	// vote
	// Required: true
	Vote *GetProposalListResultProposalsVote `json:"vote"`
}

// Validate validates this get proposal list result proposals
func (m *GetProposalListResultProposals) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDeposit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProposer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubmitTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVote(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetProposalListResultProposals) validateDeposit(formats strfmt.Registry) error {

	if err := validate.Required("deposit", "body", m.Deposit); err != nil {
		return err
	}

	if m.Deposit != nil {
		if err := m.Deposit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deposit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deposit")
			}
			return err
		}
	}

	return nil
}

func (m *GetProposalListResultProposals) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *GetProposalListResultProposals) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *GetProposalListResultProposals) validateProposer(formats strfmt.Registry) error {

	if err := validate.Required("proposer", "body", m.Proposer); err != nil {
		return err
	}

	if m.Proposer != nil {
		if err := m.Proposer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proposer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("proposer")
			}
			return err
		}
	}

	return nil
}

func (m *GetProposalListResultProposals) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *GetProposalListResultProposals) validateSubmitTime(formats strfmt.Registry) error {

	if err := validate.Required("submitTime", "body", m.SubmitTime); err != nil {
		return err
	}

	return nil
}

func (m *GetProposalListResultProposals) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

func (m *GetProposalListResultProposals) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *GetProposalListResultProposals) validateVote(formats strfmt.Registry) error {

	if err := validate.Required("vote", "body", m.Vote); err != nil {
		return err
	}

	if m.Vote != nil {
		if err := m.Vote.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vote")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vote")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get proposal list result proposals based on the context it is used
func (m *GetProposalListResultProposals) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDeposit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProposer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVote(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetProposalListResultProposals) contextValidateDeposit(ctx context.Context, formats strfmt.Registry) error {

	if m.Deposit != nil {
		if err := m.Deposit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deposit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deposit")
			}
			return err
		}
	}

	return nil
}

func (m *GetProposalListResultProposals) contextValidateProposer(ctx context.Context, formats strfmt.Registry) error {

	if m.Proposer != nil {
		if err := m.Proposer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proposer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("proposer")
			}
			return err
		}
	}

	return nil
}

func (m *GetProposalListResultProposals) contextValidateVote(ctx context.Context, formats strfmt.Registry) error {

	if m.Vote != nil {
		if err := m.Vote.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vote")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("vote")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetProposalListResultProposals) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetProposalListResultProposals) UnmarshalBinary(b []byte) error {
	var res GetProposalListResultProposals
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

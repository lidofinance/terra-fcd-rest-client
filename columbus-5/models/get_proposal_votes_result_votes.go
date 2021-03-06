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

// GetProposalVotesResultVotes get proposal votes result votes
//
// swagger:model getProposalVotesResult.votes
type GetProposalVotesResultVotes struct {

	// 'Yes', 'No', 'NoWithVeto', 'Abstain'
	// Required: true
	Answer *string `json:"answer"`

	// Txhash of the vote transaction
	Txhash string `json:"txhash,omitempty"`

	// Voter information
	// Required: true
	Voter *GetProposalVotesResultVotesVoter `json:"voter"`
}

// Validate validates this get proposal votes result votes
func (m *GetProposalVotesResultVotes) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAnswer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVoter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetProposalVotesResultVotes) validateAnswer(formats strfmt.Registry) error {

	if err := validate.Required("answer", "body", m.Answer); err != nil {
		return err
	}

	return nil
}

func (m *GetProposalVotesResultVotes) validateVoter(formats strfmt.Registry) error {

	if err := validate.Required("voter", "body", m.Voter); err != nil {
		return err
	}

	if m.Voter != nil {
		if err := m.Voter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("voter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("voter")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get proposal votes result votes based on the context it is used
func (m *GetProposalVotesResultVotes) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVoter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetProposalVotesResultVotes) contextValidateVoter(ctx context.Context, formats strfmt.Registry) error {

	if m.Voter != nil {
		if err := m.Voter.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("voter")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("voter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetProposalVotesResultVotes) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetProposalVotesResultVotes) UnmarshalBinary(b []byte) error {
	var res GetProposalVotesResultVotes
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

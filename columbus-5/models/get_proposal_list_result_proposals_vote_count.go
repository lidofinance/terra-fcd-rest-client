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

// GetProposalListResultProposalsVoteCount get proposal list result proposals vote count
//
// swagger:model getProposalListResult.proposals.vote.count
type GetProposalListResultProposalsVoteCount struct {

	// vote count
	// Required: true
	Abstain *float64 `json:"Abstain"`

	// vote count
	// Required: true
	No *float64 `json:"No"`

	// vote count
	// Required: true
	NoWithVeto *float64 `json:"NoWithVeto"`

	// vote count
	// Required: true
	Yes *float64 `json:"Yes"`
}

// Validate validates this get proposal list result proposals vote count
func (m *GetProposalListResultProposalsVoteCount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAbstain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNoWithVeto(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateYes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetProposalListResultProposalsVoteCount) validateAbstain(formats strfmt.Registry) error {

	if err := validate.Required("Abstain", "body", m.Abstain); err != nil {
		return err
	}

	return nil
}

func (m *GetProposalListResultProposalsVoteCount) validateNo(formats strfmt.Registry) error {

	if err := validate.Required("No", "body", m.No); err != nil {
		return err
	}

	return nil
}

func (m *GetProposalListResultProposalsVoteCount) validateNoWithVeto(formats strfmt.Registry) error {

	if err := validate.Required("NoWithVeto", "body", m.NoWithVeto); err != nil {
		return err
	}

	return nil
}

func (m *GetProposalListResultProposalsVoteCount) validateYes(formats strfmt.Registry) error {

	if err := validate.Required("Yes", "body", m.Yes); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get proposal list result proposals vote count based on context it is used
func (m *GetProposalListResultProposalsVoteCount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetProposalListResultProposalsVoteCount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetProposalListResultProposalsVoteCount) UnmarshalBinary(b []byte) error {
	var res GetProposalListResultProposalsVoteCount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

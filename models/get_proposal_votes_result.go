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

// GetProposalVotesResult get proposal votes result
//
// swagger:model getProposalVotesResult
type GetProposalVotesResult struct {

	// limit
	// Required: true
	Limit *float64 `json:"limit"`

	// page
	// Required: true
	Page *float64 `json:"page"`

	// total cnt
	// Required: true
	TotalCnt *float64 `json:"totalCnt"`

	// Vote list
	// Required: true
	Votes []*GetProposalVotesResultVotes `json:"votes"`
}

// Validate validates this get proposal votes result
func (m *GetProposalVotesResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalCnt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVotes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetProposalVotesResult) validateLimit(formats strfmt.Registry) error {

	if err := validate.Required("limit", "body", m.Limit); err != nil {
		return err
	}

	return nil
}

func (m *GetProposalVotesResult) validatePage(formats strfmt.Registry) error {

	if err := validate.Required("page", "body", m.Page); err != nil {
		return err
	}

	return nil
}

func (m *GetProposalVotesResult) validateTotalCnt(formats strfmt.Registry) error {

	if err := validate.Required("totalCnt", "body", m.TotalCnt); err != nil {
		return err
	}

	return nil
}

func (m *GetProposalVotesResult) validateVotes(formats strfmt.Registry) error {

	if err := validate.Required("votes", "body", m.Votes); err != nil {
		return err
	}

	for i := 0; i < len(m.Votes); i++ {
		if swag.IsZero(m.Votes[i]) { // not required
			continue
		}

		if m.Votes[i] != nil {
			if err := m.Votes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("votes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("votes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get proposal votes result based on the context it is used
func (m *GetProposalVotesResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVotes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetProposalVotesResult) contextValidateVotes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Votes); i++ {

		if m.Votes[i] != nil {
			if err := m.Votes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("votes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("votes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetProposalVotesResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetProposalVotesResult) UnmarshalBinary(b []byte) error {
	var res GetProposalVotesResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
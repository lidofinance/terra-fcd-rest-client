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
)

// TextProposal text proposal
//
// swagger:model TextProposal
type TextProposal struct {

	// content
	Content *TextProposalContent `json:"content,omitempty"`

	// deposit end time
	DepositEndTime string `json:"deposit_end_time,omitempty"`

	// final tally result
	FinalTallyResult *TextProposalFinalTallyResult `json:"final_tally_result,omitempty"`

	// id
	// Example: 1
	ID string `json:"id,omitempty"`

	// status
	Status float64 `json:"status,omitempty"`

	// submit time
	SubmitTime string `json:"submit_time,omitempty"`

	// total deposit
	TotalDeposit []*TextProposalTotalDepositItems0 `json:"total_deposit"`

	// voting end time
	VotingEndTime string `json:"voting_end_time,omitempty"`

	// voting start time
	VotingStartTime string `json:"voting_start_time,omitempty"`
}

// Validate validates this text proposal
func (m *TextProposal) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFinalTallyResult(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotalDeposit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TextProposal) validateContent(formats strfmt.Registry) error {
	if swag.IsZero(m.Content) { // not required
		return nil
	}

	if m.Content != nil {
		if err := m.Content.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("content")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("content")
			}
			return err
		}
	}

	return nil
}

func (m *TextProposal) validateFinalTallyResult(formats strfmt.Registry) error {
	if swag.IsZero(m.FinalTallyResult) { // not required
		return nil
	}

	if m.FinalTallyResult != nil {
		if err := m.FinalTallyResult.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("final_tally_result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("final_tally_result")
			}
			return err
		}
	}

	return nil
}

func (m *TextProposal) validateTotalDeposit(formats strfmt.Registry) error {
	if swag.IsZero(m.TotalDeposit) { // not required
		return nil
	}

	for i := 0; i < len(m.TotalDeposit); i++ {
		if swag.IsZero(m.TotalDeposit[i]) { // not required
			continue
		}

		if m.TotalDeposit[i] != nil {
			if err := m.TotalDeposit[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("total_deposit" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("total_deposit" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this text proposal based on the context it is used
func (m *TextProposal) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFinalTallyResult(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTotalDeposit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TextProposal) contextValidateContent(ctx context.Context, formats strfmt.Registry) error {

	if m.Content != nil {
		if err := m.Content.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("content")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("content")
			}
			return err
		}
	}

	return nil
}

func (m *TextProposal) contextValidateFinalTallyResult(ctx context.Context, formats strfmt.Registry) error {

	if m.FinalTallyResult != nil {
		if err := m.FinalTallyResult.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("final_tally_result")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("final_tally_result")
			}
			return err
		}
	}

	return nil
}

func (m *TextProposal) contextValidateTotalDeposit(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TotalDeposit); i++ {

		if m.TotalDeposit[i] != nil {
			if err := m.TotalDeposit[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("total_deposit" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("total_deposit" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TextProposal) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TextProposal) UnmarshalBinary(b []byte) error {
	var res TextProposal
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TextProposalContent text proposal content
//
// swagger:model TextProposalContent
type TextProposalContent struct {

	// type
	// Example: gov/TextProposal
	Type string `json:"type,omitempty"`

	// value
	Value *TextProposalContentValue `json:"value,omitempty"`
}

// Validate validates this text proposal content
func (m *TextProposalContent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TextProposalContent) validateValue(formats strfmt.Registry) error {
	if swag.IsZero(m.Value) { // not required
		return nil
	}

	if m.Value != nil {
		if err := m.Value.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("content" + "." + "value")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("content" + "." + "value")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this text proposal content based on the context it is used
func (m *TextProposalContent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TextProposalContent) contextValidateValue(ctx context.Context, formats strfmt.Registry) error {

	if m.Value != nil {
		if err := m.Value.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("content" + "." + "value")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("content" + "." + "value")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TextProposalContent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TextProposalContent) UnmarshalBinary(b []byte) error {
	var res TextProposalContent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TextProposalContentValue text proposal content value
//
// swagger:model TextProposalContentValue
type TextProposalContentValue struct {

	// description
	Description string `json:"description,omitempty"`

	// title
	Title string `json:"title,omitempty"`
}

// Validate validates this text proposal content value
func (m *TextProposalContentValue) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this text proposal content value based on context it is used
func (m *TextProposalContentValue) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TextProposalContentValue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TextProposalContentValue) UnmarshalBinary(b []byte) error {
	var res TextProposalContentValue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TextProposalFinalTallyResult text proposal final tally result
//
// swagger:model TextProposalFinalTallyResult
type TextProposalFinalTallyResult struct {

	// abstain
	// Example: 0.0000000000
	Abstain string `json:"abstain,omitempty"`

	// no
	// Example: 0.0000000000
	No string `json:"no,omitempty"`

	// no with veto
	// Example: 0.0000000000
	NoWithVeto string `json:"no_with_veto,omitempty"`

	// yes
	// Example: 0.0000000000
	Yes string `json:"yes,omitempty"`
}

// Validate validates this text proposal final tally result
func (m *TextProposalFinalTallyResult) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this text proposal final tally result based on context it is used
func (m *TextProposalFinalTallyResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TextProposalFinalTallyResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TextProposalFinalTallyResult) UnmarshalBinary(b []byte) error {
	var res TextProposalFinalTallyResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// TextProposalTotalDepositItems0 text proposal total deposit items0
//
// swagger:model TextProposalTotalDepositItems0
type TextProposalTotalDepositItems0 struct {

	// amount
	// Example: 50
	Amount string `json:"amount,omitempty"`

	// denom
	// Example: uluna
	Denom string `json:"denom,omitempty"`
}

// Validate validates this text proposal total deposit items0
func (m *TextProposalTotalDepositItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this text proposal total deposit items0 based on context it is used
func (m *TextProposalTotalDepositItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TextProposalTotalDepositItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TextProposalTotalDepositItems0) UnmarshalBinary(b []byte) error {
	var res TextProposalTotalDepositItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
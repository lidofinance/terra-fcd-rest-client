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

// CosmosGovV1beta1QueryVoteResponse QueryVoteResponse is the response type for the Query/Vote RPC method.
//
// swagger:model cosmos.gov.v1beta1.QueryVoteResponse
type CosmosGovV1beta1QueryVoteResponse struct {

	// vote
	Vote *CosmosGovV1beta1QueryVoteResponseVote `json:"vote,omitempty"`
}

// Validate validates this cosmos gov v1beta1 query vote response
func (m *CosmosGovV1beta1QueryVoteResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVote(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosGovV1beta1QueryVoteResponse) validateVote(formats strfmt.Registry) error {
	if swag.IsZero(m.Vote) { // not required
		return nil
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

// ContextValidate validate this cosmos gov v1beta1 query vote response based on the context it is used
func (m *CosmosGovV1beta1QueryVoteResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateVote(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosGovV1beta1QueryVoteResponse) contextValidateVote(ctx context.Context, formats strfmt.Registry) error {

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
func (m *CosmosGovV1beta1QueryVoteResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosGovV1beta1QueryVoteResponse) UnmarshalBinary(b []byte) error {
	var res CosmosGovV1beta1QueryVoteResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CosmosGovV1beta1QueryVoteResponseVote Vote defines a vote on a governance proposal.
// A Vote consists of a proposal ID, the voter, and the vote option.
//
// swagger:model CosmosGovV1beta1QueryVoteResponseVote
type CosmosGovV1beta1QueryVoteResponseVote struct {

	// Deprecated: Prefer to use `options` instead. This field is set in queries
	// if and only if `len(options) == 1` and that option has weight 1. In all
	// other cases, this field will default to VOTE_OPTION_UNSPECIFIED.
	// Enum: [VOTE_OPTION_UNSPECIFIED VOTE_OPTION_YES VOTE_OPTION_ABSTAIN VOTE_OPTION_NO VOTE_OPTION_NO_WITH_VETO]
	Option string `json:"option,omitempty"`

	// options
	Options []*CosmosGovV1beta1QueryVoteResponseVoteOptionsItems0 `json:"options"`

	// proposal id
	ProposalID string `json:"proposal_id,omitempty"`

	// voter
	Voter string `json:"voter,omitempty"`
}

// Validate validates this cosmos gov v1beta1 query vote response vote
func (m *CosmosGovV1beta1QueryVoteResponseVote) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOption(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOptions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var cosmosGovV1beta1QueryVoteResponseVoteTypeOptionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["VOTE_OPTION_UNSPECIFIED","VOTE_OPTION_YES","VOTE_OPTION_ABSTAIN","VOTE_OPTION_NO","VOTE_OPTION_NO_WITH_VETO"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cosmosGovV1beta1QueryVoteResponseVoteTypeOptionPropEnum = append(cosmosGovV1beta1QueryVoteResponseVoteTypeOptionPropEnum, v)
	}
}

const (

	// CosmosGovV1beta1QueryVoteResponseVoteOptionVOTEOPTIONUNSPECIFIED captures enum value "VOTE_OPTION_UNSPECIFIED"
	CosmosGovV1beta1QueryVoteResponseVoteOptionVOTEOPTIONUNSPECIFIED string = "VOTE_OPTION_UNSPECIFIED"

	// CosmosGovV1beta1QueryVoteResponseVoteOptionVOTEOPTIONYES captures enum value "VOTE_OPTION_YES"
	CosmosGovV1beta1QueryVoteResponseVoteOptionVOTEOPTIONYES string = "VOTE_OPTION_YES"

	// CosmosGovV1beta1QueryVoteResponseVoteOptionVOTEOPTIONABSTAIN captures enum value "VOTE_OPTION_ABSTAIN"
	CosmosGovV1beta1QueryVoteResponseVoteOptionVOTEOPTIONABSTAIN string = "VOTE_OPTION_ABSTAIN"

	// CosmosGovV1beta1QueryVoteResponseVoteOptionVOTEOPTIONNO captures enum value "VOTE_OPTION_NO"
	CosmosGovV1beta1QueryVoteResponseVoteOptionVOTEOPTIONNO string = "VOTE_OPTION_NO"

	// CosmosGovV1beta1QueryVoteResponseVoteOptionVOTEOPTIONNOWITHVETO captures enum value "VOTE_OPTION_NO_WITH_VETO"
	CosmosGovV1beta1QueryVoteResponseVoteOptionVOTEOPTIONNOWITHVETO string = "VOTE_OPTION_NO_WITH_VETO"
)

// prop value enum
func (m *CosmosGovV1beta1QueryVoteResponseVote) validateOptionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cosmosGovV1beta1QueryVoteResponseVoteTypeOptionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CosmosGovV1beta1QueryVoteResponseVote) validateOption(formats strfmt.Registry) error {
	if swag.IsZero(m.Option) { // not required
		return nil
	}

	// value enum
	if err := m.validateOptionEnum("vote"+"."+"option", "body", m.Option); err != nil {
		return err
	}

	return nil
}

func (m *CosmosGovV1beta1QueryVoteResponseVote) validateOptions(formats strfmt.Registry) error {
	if swag.IsZero(m.Options) { // not required
		return nil
	}

	for i := 0; i < len(m.Options); i++ {
		if swag.IsZero(m.Options[i]) { // not required
			continue
		}

		if m.Options[i] != nil {
			if err := m.Options[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vote" + "." + "options" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vote" + "." + "options" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this cosmos gov v1beta1 query vote response vote based on the context it is used
func (m *CosmosGovV1beta1QueryVoteResponseVote) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosGovV1beta1QueryVoteResponseVote) contextValidateOptions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Options); i++ {

		if m.Options[i] != nil {
			if err := m.Options[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vote" + "." + "options" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("vote" + "." + "options" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CosmosGovV1beta1QueryVoteResponseVote) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosGovV1beta1QueryVoteResponseVote) UnmarshalBinary(b []byte) error {
	var res CosmosGovV1beta1QueryVoteResponseVote
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CosmosGovV1beta1QueryVoteResponseVoteOptionsItems0 WeightedVoteOption defines a unit of vote for vote split.
//
// swagger:model CosmosGovV1beta1QueryVoteResponseVoteOptionsItems0
type CosmosGovV1beta1QueryVoteResponseVoteOptionsItems0 struct {

	// VoteOption enumerates the valid vote options for a given governance proposal.
	//
	//  - VOTE_OPTION_UNSPECIFIED: VOTE_OPTION_UNSPECIFIED defines a no-op vote option.
	//  - VOTE_OPTION_YES: VOTE_OPTION_YES defines a yes vote option.
	//  - VOTE_OPTION_ABSTAIN: VOTE_OPTION_ABSTAIN defines an abstain vote option.
	//  - VOTE_OPTION_NO: VOTE_OPTION_NO defines a no vote option.
	//  - VOTE_OPTION_NO_WITH_VETO: VOTE_OPTION_NO_WITH_VETO defines a no with veto vote option.
	// Enum: [VOTE_OPTION_UNSPECIFIED VOTE_OPTION_YES VOTE_OPTION_ABSTAIN VOTE_OPTION_NO VOTE_OPTION_NO_WITH_VETO]
	Option string `json:"option,omitempty"`

	// weight
	Weight string `json:"weight,omitempty"`
}

// Validate validates this cosmos gov v1beta1 query vote response vote options items0
func (m *CosmosGovV1beta1QueryVoteResponseVoteOptionsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOption(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var cosmosGovV1beta1QueryVoteResponseVoteOptionsItems0TypeOptionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["VOTE_OPTION_UNSPECIFIED","VOTE_OPTION_YES","VOTE_OPTION_ABSTAIN","VOTE_OPTION_NO","VOTE_OPTION_NO_WITH_VETO"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cosmosGovV1beta1QueryVoteResponseVoteOptionsItems0TypeOptionPropEnum = append(cosmosGovV1beta1QueryVoteResponseVoteOptionsItems0TypeOptionPropEnum, v)
	}
}

const (

	// CosmosGovV1beta1QueryVoteResponseVoteOptionsItems0OptionVOTEOPTIONUNSPECIFIED captures enum value "VOTE_OPTION_UNSPECIFIED"
	CosmosGovV1beta1QueryVoteResponseVoteOptionsItems0OptionVOTEOPTIONUNSPECIFIED string = "VOTE_OPTION_UNSPECIFIED"

	// CosmosGovV1beta1QueryVoteResponseVoteOptionsItems0OptionVOTEOPTIONYES captures enum value "VOTE_OPTION_YES"
	CosmosGovV1beta1QueryVoteResponseVoteOptionsItems0OptionVOTEOPTIONYES string = "VOTE_OPTION_YES"

	// CosmosGovV1beta1QueryVoteResponseVoteOptionsItems0OptionVOTEOPTIONABSTAIN captures enum value "VOTE_OPTION_ABSTAIN"
	CosmosGovV1beta1QueryVoteResponseVoteOptionsItems0OptionVOTEOPTIONABSTAIN string = "VOTE_OPTION_ABSTAIN"

	// CosmosGovV1beta1QueryVoteResponseVoteOptionsItems0OptionVOTEOPTIONNO captures enum value "VOTE_OPTION_NO"
	CosmosGovV1beta1QueryVoteResponseVoteOptionsItems0OptionVOTEOPTIONNO string = "VOTE_OPTION_NO"

	// CosmosGovV1beta1QueryVoteResponseVoteOptionsItems0OptionVOTEOPTIONNOWITHVETO captures enum value "VOTE_OPTION_NO_WITH_VETO"
	CosmosGovV1beta1QueryVoteResponseVoteOptionsItems0OptionVOTEOPTIONNOWITHVETO string = "VOTE_OPTION_NO_WITH_VETO"
)

// prop value enum
func (m *CosmosGovV1beta1QueryVoteResponseVoteOptionsItems0) validateOptionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cosmosGovV1beta1QueryVoteResponseVoteOptionsItems0TypeOptionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CosmosGovV1beta1QueryVoteResponseVoteOptionsItems0) validateOption(formats strfmt.Registry) error {
	if swag.IsZero(m.Option) { // not required
		return nil
	}

	// value enum
	if err := m.validateOptionEnum("option", "body", m.Option); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this cosmos gov v1beta1 query vote response vote options items0 based on context it is used
func (m *CosmosGovV1beta1QueryVoteResponseVoteOptionsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CosmosGovV1beta1QueryVoteResponseVoteOptionsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosGovV1beta1QueryVoteResponseVoteOptionsItems0) UnmarshalBinary(b []byte) error {
	var res CosmosGovV1beta1QueryVoteResponseVoteOptionsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

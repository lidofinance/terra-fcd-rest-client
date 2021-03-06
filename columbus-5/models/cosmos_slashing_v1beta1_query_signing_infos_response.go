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

// CosmosSlashingV1beta1QuerySigningInfosResponse QuerySigningInfosResponse is the response type for the Query/SigningInfos RPC
// method
//
// swagger:model cosmos.slashing.v1beta1.QuerySigningInfosResponse
type CosmosSlashingV1beta1QuerySigningInfosResponse struct {

	// info is the signing info of all validators
	Info []*CosmosSlashingV1beta1QuerySigningInfosResponseInfoItems0 `json:"info"`

	// pagination
	Pagination *CosmosSlashingV1beta1QuerySigningInfosResponsePagination `json:"pagination,omitempty"`
}

// Validate validates this cosmos slashing v1beta1 query signing infos response
func (m *CosmosSlashingV1beta1QuerySigningInfosResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePagination(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosSlashingV1beta1QuerySigningInfosResponse) validateInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.Info) { // not required
		return nil
	}

	for i := 0; i < len(m.Info); i++ {
		if swag.IsZero(m.Info[i]) { // not required
			continue
		}

		if m.Info[i] != nil {
			if err := m.Info[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("info" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CosmosSlashingV1beta1QuerySigningInfosResponse) validatePagination(formats strfmt.Registry) error {
	if swag.IsZero(m.Pagination) { // not required
		return nil
	}

	if m.Pagination != nil {
		if err := m.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pagination")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cosmos slashing v1beta1 query signing infos response based on the context it is used
func (m *CosmosSlashingV1beta1QuerySigningInfosResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePagination(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosSlashingV1beta1QuerySigningInfosResponse) contextValidateInfo(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Info); i++ {

		if m.Info[i] != nil {
			if err := m.Info[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("info" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("info" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CosmosSlashingV1beta1QuerySigningInfosResponse) contextValidatePagination(ctx context.Context, formats strfmt.Registry) error {

	if m.Pagination != nil {
		if err := m.Pagination.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pagination")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CosmosSlashingV1beta1QuerySigningInfosResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosSlashingV1beta1QuerySigningInfosResponse) UnmarshalBinary(b []byte) error {
	var res CosmosSlashingV1beta1QuerySigningInfosResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CosmosSlashingV1beta1QuerySigningInfosResponseInfoItems0 ValidatorSigningInfo defines a validator's signing info for monitoring their
// liveness activity.
//
// swagger:model CosmosSlashingV1beta1QuerySigningInfosResponseInfoItems0
type CosmosSlashingV1beta1QuerySigningInfosResponseInfoItems0 struct {

	// address
	Address string `json:"address,omitempty"`

	// Index which is incremented each time the validator was a bonded
	// in a block and may have signed a precommit or not. This in conjunction with the
	// `SignedBlocksWindow` param determines the index in the `MissedBlocksBitArray`.
	IndexOffset string `json:"index_offset,omitempty"`

	// Timestamp until which the validator is jailed due to liveness downtime.
	// Format: date-time
	JailedUntil strfmt.DateTime `json:"jailed_until,omitempty"`

	// A counter kept to avoid unnecessary array reads.
	// Note that `Sum(MissedBlocksBitArray)` always equals `MissedBlocksCounter`.
	MissedBlocksCounter string `json:"missed_blocks_counter,omitempty"`

	// Height at which validator was first a candidate OR was unjailed
	StartHeight string `json:"start_height,omitempty"`

	// Whether or not a validator has been tombstoned (killed out of validator set). It is set
	// once the validator commits an equivocation or for any other configured misbehiavor.
	Tombstoned bool `json:"tombstoned,omitempty"`
}

// Validate validates this cosmos slashing v1beta1 query signing infos response info items0
func (m *CosmosSlashingV1beta1QuerySigningInfosResponseInfoItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateJailedUntil(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosSlashingV1beta1QuerySigningInfosResponseInfoItems0) validateJailedUntil(formats strfmt.Registry) error {
	if swag.IsZero(m.JailedUntil) { // not required
		return nil
	}

	if err := validate.FormatOf("jailed_until", "body", "date-time", m.JailedUntil.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this cosmos slashing v1beta1 query signing infos response info items0 based on context it is used
func (m *CosmosSlashingV1beta1QuerySigningInfosResponseInfoItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CosmosSlashingV1beta1QuerySigningInfosResponseInfoItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosSlashingV1beta1QuerySigningInfosResponseInfoItems0) UnmarshalBinary(b []byte) error {
	var res CosmosSlashingV1beta1QuerySigningInfosResponseInfoItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CosmosSlashingV1beta1QuerySigningInfosResponsePagination PageResponse is to be embedded in gRPC response messages where the
// corresponding request message has used PageRequest.
//
//  message SomeResponse {
//          repeated Bar results = 1;
//          PageResponse page = 2;
//  }
//
// swagger:model CosmosSlashingV1beta1QuerySigningInfosResponsePagination
type CosmosSlashingV1beta1QuerySigningInfosResponsePagination struct {

	// next_key is the key to be passed to PageRequest.key to
	// query the next page most efficiently
	// Format: byte
	NextKey strfmt.Base64 `json:"next_key,omitempty"`

	// total is total number of results available if PageRequest.count_total
	// was set, its value is undefined otherwise
	Total string `json:"total,omitempty"`
}

// Validate validates this cosmos slashing v1beta1 query signing infos response pagination
func (m *CosmosSlashingV1beta1QuerySigningInfosResponsePagination) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cosmos slashing v1beta1 query signing infos response pagination based on context it is used
func (m *CosmosSlashingV1beta1QuerySigningInfosResponsePagination) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CosmosSlashingV1beta1QuerySigningInfosResponsePagination) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosSlashingV1beta1QuerySigningInfosResponsePagination) UnmarshalBinary(b []byte) error {
	var res CosmosSlashingV1beta1QuerySigningInfosResponsePagination
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

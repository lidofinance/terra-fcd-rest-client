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

// CosmosFeegrantV1beta1QueryAllowancesResponse QueryAllowancesResponse is the response type for the Query/Allowances RPC method.
//
// swagger:model cosmos.feegrant.v1beta1.QueryAllowancesResponse
type CosmosFeegrantV1beta1QueryAllowancesResponse struct {

	// allowances are allowance's granted for grantee by granter.
	Allowances []*CosmosFeegrantV1beta1QueryAllowancesResponseAllowancesItems0 `json:"allowances"`

	// pagination
	Pagination *CosmosFeegrantV1beta1QueryAllowancesResponsePagination `json:"pagination,omitempty"`
}

// Validate validates this cosmos feegrant v1beta1 query allowances response
func (m *CosmosFeegrantV1beta1QueryAllowancesResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowances(formats); err != nil {
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

func (m *CosmosFeegrantV1beta1QueryAllowancesResponse) validateAllowances(formats strfmt.Registry) error {
	if swag.IsZero(m.Allowances) { // not required
		return nil
	}

	for i := 0; i < len(m.Allowances); i++ {
		if swag.IsZero(m.Allowances[i]) { // not required
			continue
		}

		if m.Allowances[i] != nil {
			if err := m.Allowances[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("allowances" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("allowances" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CosmosFeegrantV1beta1QueryAllowancesResponse) validatePagination(formats strfmt.Registry) error {
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

// ContextValidate validate this cosmos feegrant v1beta1 query allowances response based on the context it is used
func (m *CosmosFeegrantV1beta1QueryAllowancesResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAllowances(ctx, formats); err != nil {
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

func (m *CosmosFeegrantV1beta1QueryAllowancesResponse) contextValidateAllowances(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Allowances); i++ {

		if m.Allowances[i] != nil {
			if err := m.Allowances[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("allowances" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("allowances" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CosmosFeegrantV1beta1QueryAllowancesResponse) contextValidatePagination(ctx context.Context, formats strfmt.Registry) error {

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
func (m *CosmosFeegrantV1beta1QueryAllowancesResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosFeegrantV1beta1QueryAllowancesResponse) UnmarshalBinary(b []byte) error {
	var res CosmosFeegrantV1beta1QueryAllowancesResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CosmosFeegrantV1beta1QueryAllowancesResponseAllowancesItems0 Grant is stored in the KVStore to record a grant with full context
//
// swagger:model CosmosFeegrantV1beta1QueryAllowancesResponseAllowancesItems0
type CosmosFeegrantV1beta1QueryAllowancesResponseAllowancesItems0 struct {

	// allowance
	Allowance *CosmosFeegrantV1beta1QueryAllowancesResponseAllowancesItems0Allowance `json:"allowance,omitempty"`

	// grantee is the address of the user being granted an allowance of another user's funds.
	Grantee string `json:"grantee,omitempty"`

	// granter is the address of the user granting an allowance of their funds.
	Granter string `json:"granter,omitempty"`
}

// Validate validates this cosmos feegrant v1beta1 query allowances response allowances items0
func (m *CosmosFeegrantV1beta1QueryAllowancesResponseAllowancesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllowance(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosFeegrantV1beta1QueryAllowancesResponseAllowancesItems0) validateAllowance(formats strfmt.Registry) error {
	if swag.IsZero(m.Allowance) { // not required
		return nil
	}

	if m.Allowance != nil {
		if err := m.Allowance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("allowance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("allowance")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cosmos feegrant v1beta1 query allowances response allowances items0 based on the context it is used
func (m *CosmosFeegrantV1beta1QueryAllowancesResponseAllowancesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAllowance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosFeegrantV1beta1QueryAllowancesResponseAllowancesItems0) contextValidateAllowance(ctx context.Context, formats strfmt.Registry) error {

	if m.Allowance != nil {
		if err := m.Allowance.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("allowance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("allowance")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CosmosFeegrantV1beta1QueryAllowancesResponseAllowancesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosFeegrantV1beta1QueryAllowancesResponseAllowancesItems0) UnmarshalBinary(b []byte) error {
	var res CosmosFeegrantV1beta1QueryAllowancesResponseAllowancesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CosmosFeegrantV1beta1QueryAllowancesResponseAllowancesItems0Allowance allowance can be any of basic and filtered fee allowance.
//
// swagger:model CosmosFeegrantV1beta1QueryAllowancesResponseAllowancesItems0Allowance
type CosmosFeegrantV1beta1QueryAllowancesResponseAllowancesItems0Allowance struct {

	// A URL/resource name that uniquely identifies the type of the serialized
	// protocol buffer message. This string must contain at least
	// one "/" character. The last segment of the URL's path must represent
	// the fully qualified name of the type (as in
	// `path/google.protobuf.Duration`). The name should be in a canonical form
	// (e.g., leading "." is not accepted).
	//
	// In practice, teams usually precompile into the binary all types that they
	// expect it to use in the context of Any. However, for URLs which use the
	// scheme `http`, `https`, or no scheme, one can optionally set up a type
	// server that maps type URLs to message definitions as follows:
	//
	// * If no scheme is provided, `https` is assumed.
	// * An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//   value in binary format, or produce an error.
	// * Applications are allowed to cache lookup results based on the
	//   URL, or have them precompiled into a binary to avoid any
	//   lookup. Therefore, binary compatibility needs to be preserved
	//   on changes to types. (Use versioned type names to manage
	//   breaking changes.)
	//
	// Note: this functionality is not currently available in the official
	// protobuf release, and it is not used for type URLs beginning with
	// type.googleapis.com.
	//
	// Schemes other than `http`, `https` (or the empty scheme) might be
	// used with implementation specific semantics.
	TypeURL string `json:"type_url,omitempty"`

	// Must be a valid serialized protocol buffer of the above specified type.
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this cosmos feegrant v1beta1 query allowances response allowances items0 allowance
func (m *CosmosFeegrantV1beta1QueryAllowancesResponseAllowancesItems0Allowance) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cosmos feegrant v1beta1 query allowances response allowances items0 allowance based on context it is used
func (m *CosmosFeegrantV1beta1QueryAllowancesResponseAllowancesItems0Allowance) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CosmosFeegrantV1beta1QueryAllowancesResponseAllowancesItems0Allowance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosFeegrantV1beta1QueryAllowancesResponseAllowancesItems0Allowance) UnmarshalBinary(b []byte) error {
	var res CosmosFeegrantV1beta1QueryAllowancesResponseAllowancesItems0Allowance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CosmosFeegrantV1beta1QueryAllowancesResponsePagination pagination defines an pagination for the response.
//
// swagger:model CosmosFeegrantV1beta1QueryAllowancesResponsePagination
type CosmosFeegrantV1beta1QueryAllowancesResponsePagination struct {

	// next_key is the key to be passed to PageRequest.key to
	// query the next page most efficiently
	// Format: byte
	NextKey strfmt.Base64 `json:"next_key,omitempty"`

	// total is total number of results available if PageRequest.count_total
	// was set, its value is undefined otherwise
	Total string `json:"total,omitempty"`
}

// Validate validates this cosmos feegrant v1beta1 query allowances response pagination
func (m *CosmosFeegrantV1beta1QueryAllowancesResponsePagination) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cosmos feegrant v1beta1 query allowances response pagination based on context it is used
func (m *CosmosFeegrantV1beta1QueryAllowancesResponsePagination) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CosmosFeegrantV1beta1QueryAllowancesResponsePagination) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosFeegrantV1beta1QueryAllowancesResponsePagination) UnmarshalBinary(b []byte) error {
	var res CosmosFeegrantV1beta1QueryAllowancesResponsePagination
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

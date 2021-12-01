// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CosmosBaseQueryV1beta1PageRequest PageRequest is to be embedded in gRPC request messages for efficient
// pagination. Ex:
//
// message SomeRequest {
//          Foo some_parameter = 1;
//          PageRequest pagination = 2;
//  }
//
// swagger:model cosmos.base.query.v1beta1.PageRequest
type CosmosBaseQueryV1beta1PageRequest struct {

	// count_total is set to true  to indicate that the result set should include
	// a count of the total number of items available for pagination in UIs.
	// count_total is only respected when offset is used. It is ignored when key
	// is set.
	CountTotal bool `json:"count_total,omitempty"`

	// key is a value returned in PageResponse.next_key to begin
	// querying the next page most efficiently. Only one of offset or key
	// should be set.
	// Format: byte
	Key strfmt.Base64 `json:"key,omitempty"`

	// limit is the total number of results to be returned in the result page.
	// If left empty it will default to a value to be set by each app.
	Limit string `json:"limit,omitempty"`

	// offset is a numeric offset that can be used when key is unavailable.
	// It is less efficient than using key. Only one of offset or key should
	// be set.
	Offset string `json:"offset,omitempty"`

	// reverse is set to true if results are to be returned in the descending order.
	Reverse bool `json:"reverse,omitempty"`
}

// Validate validates this cosmos base query v1beta1 page request
func (m *CosmosBaseQueryV1beta1PageRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cosmos base query v1beta1 page request based on context it is used
func (m *CosmosBaseQueryV1beta1PageRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CosmosBaseQueryV1beta1PageRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosBaseQueryV1beta1PageRequest) UnmarshalBinary(b []byte) error {
	var res CosmosBaseQueryV1beta1PageRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
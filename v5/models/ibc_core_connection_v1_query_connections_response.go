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

// IbcCoreConnectionV1QueryConnectionsResponse QueryConnectionsResponse is the response type for the Query/Connections RPC
// method.
//
// swagger:model ibc.core.connection.v1.QueryConnectionsResponse
type IbcCoreConnectionV1QueryConnectionsResponse struct {

	// list of stored connections of the chain.
	Connections []*IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0 `json:"connections"`

	// height
	Height *IbcCoreConnectionV1QueryConnectionsResponseHeight `json:"height,omitempty"`

	// pagination
	Pagination *IbcCoreConnectionV1QueryConnectionsResponsePagination `json:"pagination,omitempty"`
}

// Validate validates this ibc core connection v1 query connections response
func (m *IbcCoreConnectionV1QueryConnectionsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnections(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHeight(formats); err != nil {
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

func (m *IbcCoreConnectionV1QueryConnectionsResponse) validateConnections(formats strfmt.Registry) error {
	if swag.IsZero(m.Connections) { // not required
		return nil
	}

	for i := 0; i < len(m.Connections); i++ {
		if swag.IsZero(m.Connections[i]) { // not required
			continue
		}

		if m.Connections[i] != nil {
			if err := m.Connections[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("connections" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("connections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IbcCoreConnectionV1QueryConnectionsResponse) validateHeight(formats strfmt.Registry) error {
	if swag.IsZero(m.Height) { // not required
		return nil
	}

	if m.Height != nil {
		if err := m.Height.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("height")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("height")
			}
			return err
		}
	}

	return nil
}

func (m *IbcCoreConnectionV1QueryConnectionsResponse) validatePagination(formats strfmt.Registry) error {
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

// ContextValidate validate this ibc core connection v1 query connections response based on the context it is used
func (m *IbcCoreConnectionV1QueryConnectionsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConnections(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHeight(ctx, formats); err != nil {
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

func (m *IbcCoreConnectionV1QueryConnectionsResponse) contextValidateConnections(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Connections); i++ {

		if m.Connections[i] != nil {
			if err := m.Connections[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("connections" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("connections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IbcCoreConnectionV1QueryConnectionsResponse) contextValidateHeight(ctx context.Context, formats strfmt.Registry) error {

	if m.Height != nil {
		if err := m.Height.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("height")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("height")
			}
			return err
		}
	}

	return nil
}

func (m *IbcCoreConnectionV1QueryConnectionsResponse) contextValidatePagination(ctx context.Context, formats strfmt.Registry) error {

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
func (m *IbcCoreConnectionV1QueryConnectionsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IbcCoreConnectionV1QueryConnectionsResponse) UnmarshalBinary(b []byte) error {
	var res IbcCoreConnectionV1QueryConnectionsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0 IdentifiedConnection defines a connection with additional connection
// identifier field.
//
// swagger:model IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0
type IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0 struct {

	// client associated with this connection.
	ClientID string `json:"client_id,omitempty"`

	// counterparty
	Counterparty *IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0Counterparty `json:"counterparty,omitempty"`

	// delay period associated with this connection.
	DelayPeriod string `json:"delay_period,omitempty"`

	// connection identifier.
	ID string `json:"id,omitempty"`

	// current state of the connection end.
	// Enum: [STATE_UNINITIALIZED_UNSPECIFIED STATE_INIT STATE_TRYOPEN STATE_OPEN]
	State *string `json:"state,omitempty"`

	// IBC version which can be utilised to determine encodings or protocols for
	// channels or packets utilising this connection
	Versions []*IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0VersionsItems0 `json:"versions"`
}

// Validate validates this ibc core connection v1 query connections response connections items0
func (m *IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCounterparty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0) validateCounterparty(formats strfmt.Registry) error {
	if swag.IsZero(m.Counterparty) { // not required
		return nil
	}

	if m.Counterparty != nil {
		if err := m.Counterparty.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("counterparty")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("counterparty")
			}
			return err
		}
	}

	return nil
}

var ibcCoreConnectionV1QueryConnectionsResponseConnectionsItems0TypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["STATE_UNINITIALIZED_UNSPECIFIED","STATE_INIT","STATE_TRYOPEN","STATE_OPEN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ibcCoreConnectionV1QueryConnectionsResponseConnectionsItems0TypeStatePropEnum = append(ibcCoreConnectionV1QueryConnectionsResponseConnectionsItems0TypeStatePropEnum, v)
	}
}

const (

	// IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0StateSTATEUNINITIALIZEDUNSPECIFIED captures enum value "STATE_UNINITIALIZED_UNSPECIFIED"
	IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0StateSTATEUNINITIALIZEDUNSPECIFIED string = "STATE_UNINITIALIZED_UNSPECIFIED"

	// IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0StateSTATEINIT captures enum value "STATE_INIT"
	IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0StateSTATEINIT string = "STATE_INIT"

	// IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0StateSTATETRYOPEN captures enum value "STATE_TRYOPEN"
	IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0StateSTATETRYOPEN string = "STATE_TRYOPEN"

	// IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0StateSTATEOPEN captures enum value "STATE_OPEN"
	IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0StateSTATEOPEN string = "STATE_OPEN"
)

// prop value enum
func (m *IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ibcCoreConnectionV1QueryConnectionsResponseConnectionsItems0TypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

func (m *IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0) validateVersions(formats strfmt.Registry) error {
	if swag.IsZero(m.Versions) { // not required
		return nil
	}

	for i := 0; i < len(m.Versions); i++ {
		if swag.IsZero(m.Versions[i]) { // not required
			continue
		}

		if m.Versions[i] != nil {
			if err := m.Versions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("versions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("versions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this ibc core connection v1 query connections response connections items0 based on the context it is used
func (m *IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCounterparty(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVersions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0) contextValidateCounterparty(ctx context.Context, formats strfmt.Registry) error {

	if m.Counterparty != nil {
		if err := m.Counterparty.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("counterparty")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("counterparty")
			}
			return err
		}
	}

	return nil
}

func (m *IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0) contextValidateVersions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Versions); i++ {

		if m.Versions[i] != nil {
			if err := m.Versions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("versions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("versions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0) UnmarshalBinary(b []byte) error {
	var res IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0Counterparty counterparty chain associated with this connection.
//
// swagger:model IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0Counterparty
type IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0Counterparty struct {

	// identifies the client on the counterparty chain associated with a given
	// connection.
	ClientID string `json:"client_id,omitempty"`

	// identifies the connection end on the counterparty chain associated with a
	// given connection.
	ConnectionID string `json:"connection_id,omitempty"`

	// prefix
	Prefix *IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0CounterpartyPrefix `json:"prefix,omitempty"`
}

// Validate validates this ibc core connection v1 query connections response connections items0 counterparty
func (m *IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0Counterparty) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrefix(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0Counterparty) validatePrefix(formats strfmt.Registry) error {
	if swag.IsZero(m.Prefix) { // not required
		return nil
	}

	if m.Prefix != nil {
		if err := m.Prefix.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("counterparty" + "." + "prefix")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("counterparty" + "." + "prefix")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ibc core connection v1 query connections response connections items0 counterparty based on the context it is used
func (m *IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0Counterparty) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePrefix(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0Counterparty) contextValidatePrefix(ctx context.Context, formats strfmt.Registry) error {

	if m.Prefix != nil {
		if err := m.Prefix.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("counterparty" + "." + "prefix")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("counterparty" + "." + "prefix")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0Counterparty) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0Counterparty) UnmarshalBinary(b []byte) error {
	var res IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0Counterparty
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0CounterpartyPrefix MerklePrefix is merkle path prefixed to the key.
// The constructed key from the Path and the key will be append(Path.KeyPath,
// append(Path.KeyPrefix, key...))
//
// commitment merkle prefix of the counterparty chain.
//
// swagger:model IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0CounterpartyPrefix
type IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0CounterpartyPrefix struct {

	// key prefix
	// Format: byte
	KeyPrefix strfmt.Base64 `json:"key_prefix,omitempty"`
}

// Validate validates this ibc core connection v1 query connections response connections items0 counterparty prefix
func (m *IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0CounterpartyPrefix) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ibc core connection v1 query connections response connections items0 counterparty prefix based on context it is used
func (m *IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0CounterpartyPrefix) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0CounterpartyPrefix) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0CounterpartyPrefix) UnmarshalBinary(b []byte) error {
	var res IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0CounterpartyPrefix
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0VersionsItems0 Version defines the versioning scheme used to negotiate the IBC verison in
// the connection handshake.
//
// swagger:model IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0VersionsItems0
type IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0VersionsItems0 struct {

	// list of features compatible with the specified identifier
	Features []string `json:"features"`

	// unique version identifier
	Identifier string `json:"identifier,omitempty"`
}

// Validate validates this ibc core connection v1 query connections response connections items0 versions items0
func (m *IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0VersionsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ibc core connection v1 query connections response connections items0 versions items0 based on context it is used
func (m *IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0VersionsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0VersionsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0VersionsItems0) UnmarshalBinary(b []byte) error {
	var res IbcCoreConnectionV1QueryConnectionsResponseConnectionsItems0VersionsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IbcCoreConnectionV1QueryConnectionsResponseHeight query block height
//
// Normally the RevisionHeight is incremented at each height while keeping
// RevisionNumber the same. However some consensus algorithms may choose to
// reset the height in certain conditions e.g. hard forks, state-machine
// breaking changes In these cases, the RevisionNumber is incremented so that
// height continues to be monitonically increasing even as the RevisionHeight
// gets reset
//
// swagger:model IbcCoreConnectionV1QueryConnectionsResponseHeight
type IbcCoreConnectionV1QueryConnectionsResponseHeight struct {

	// the height within the given revision
	RevisionHeight string `json:"revision_height,omitempty"`

	// the revision that the client is currently on
	RevisionNumber string `json:"revision_number,omitempty"`
}

// Validate validates this ibc core connection v1 query connections response height
func (m *IbcCoreConnectionV1QueryConnectionsResponseHeight) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ibc core connection v1 query connections response height based on context it is used
func (m *IbcCoreConnectionV1QueryConnectionsResponseHeight) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IbcCoreConnectionV1QueryConnectionsResponseHeight) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IbcCoreConnectionV1QueryConnectionsResponseHeight) UnmarshalBinary(b []byte) error {
	var res IbcCoreConnectionV1QueryConnectionsResponseHeight
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IbcCoreConnectionV1QueryConnectionsResponsePagination pagination response
//
// PageResponse is to be embedded in gRPC response messages where the
// corresponding request message has used PageRequest.
//
//  message SomeResponse {
//          repeated Bar results = 1;
//          PageResponse page = 2;
//  }
//
// swagger:model IbcCoreConnectionV1QueryConnectionsResponsePagination
type IbcCoreConnectionV1QueryConnectionsResponsePagination struct {

	// next_key is the key to be passed to PageRequest.key to
	// query the next page most efficiently
	// Format: byte
	NextKey strfmt.Base64 `json:"next_key,omitempty"`

	// total is total number of results available if PageRequest.count_total
	// was set, its value is undefined otherwise
	Total string `json:"total,omitempty"`
}

// Validate validates this ibc core connection v1 query connections response pagination
func (m *IbcCoreConnectionV1QueryConnectionsResponsePagination) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ibc core connection v1 query connections response pagination based on context it is used
func (m *IbcCoreConnectionV1QueryConnectionsResponsePagination) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IbcCoreConnectionV1QueryConnectionsResponsePagination) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IbcCoreConnectionV1QueryConnectionsResponsePagination) UnmarshalBinary(b []byte) error {
	var res IbcCoreConnectionV1QueryConnectionsResponsePagination
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
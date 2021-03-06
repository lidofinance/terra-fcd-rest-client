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

// IbcCoreConnectionV1ConnectionEnd ConnectionEnd defines a stateful object on a chain connected to another
// separate one.
// NOTE: there must only be 2 defined ConnectionEnds to establish
// a connection between two chains.
//
// swagger:model ibc.core.connection.v1.ConnectionEnd
type IbcCoreConnectionV1ConnectionEnd struct {

	// client associated with this connection.
	ClientID string `json:"client_id,omitempty"`

	// counterparty
	Counterparty *IbcCoreConnectionV1ConnectionEndCounterparty `json:"counterparty,omitempty"`

	// delay period that must pass before a consensus state can be used for
	// packet-verification NOTE: delay period logic is only implemented by some
	// clients.
	DelayPeriod string `json:"delay_period,omitempty"`

	// current state of the connection end.
	// Enum: [STATE_UNINITIALIZED_UNSPECIFIED STATE_INIT STATE_TRYOPEN STATE_OPEN]
	State *string `json:"state,omitempty"`

	// IBC version which can be utilised to determine encodings or protocols for
	// channels or packets utilising this connection.
	Versions []*IbcCoreConnectionV1ConnectionEndVersionsItems0 `json:"versions"`
}

// Validate validates this ibc core connection v1 connection end
func (m *IbcCoreConnectionV1ConnectionEnd) Validate(formats strfmt.Registry) error {
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

func (m *IbcCoreConnectionV1ConnectionEnd) validateCounterparty(formats strfmt.Registry) error {
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

var ibcCoreConnectionV1ConnectionEndTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["STATE_UNINITIALIZED_UNSPECIFIED","STATE_INIT","STATE_TRYOPEN","STATE_OPEN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ibcCoreConnectionV1ConnectionEndTypeStatePropEnum = append(ibcCoreConnectionV1ConnectionEndTypeStatePropEnum, v)
	}
}

const (

	// IbcCoreConnectionV1ConnectionEndStateSTATEUNINITIALIZEDUNSPECIFIED captures enum value "STATE_UNINITIALIZED_UNSPECIFIED"
	IbcCoreConnectionV1ConnectionEndStateSTATEUNINITIALIZEDUNSPECIFIED string = "STATE_UNINITIALIZED_UNSPECIFIED"

	// IbcCoreConnectionV1ConnectionEndStateSTATEINIT captures enum value "STATE_INIT"
	IbcCoreConnectionV1ConnectionEndStateSTATEINIT string = "STATE_INIT"

	// IbcCoreConnectionV1ConnectionEndStateSTATETRYOPEN captures enum value "STATE_TRYOPEN"
	IbcCoreConnectionV1ConnectionEndStateSTATETRYOPEN string = "STATE_TRYOPEN"

	// IbcCoreConnectionV1ConnectionEndStateSTATEOPEN captures enum value "STATE_OPEN"
	IbcCoreConnectionV1ConnectionEndStateSTATEOPEN string = "STATE_OPEN"
)

// prop value enum
func (m *IbcCoreConnectionV1ConnectionEnd) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ibcCoreConnectionV1ConnectionEndTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IbcCoreConnectionV1ConnectionEnd) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

func (m *IbcCoreConnectionV1ConnectionEnd) validateVersions(formats strfmt.Registry) error {
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

// ContextValidate validate this ibc core connection v1 connection end based on the context it is used
func (m *IbcCoreConnectionV1ConnectionEnd) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *IbcCoreConnectionV1ConnectionEnd) contextValidateCounterparty(ctx context.Context, formats strfmt.Registry) error {

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

func (m *IbcCoreConnectionV1ConnectionEnd) contextValidateVersions(ctx context.Context, formats strfmt.Registry) error {

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
func (m *IbcCoreConnectionV1ConnectionEnd) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IbcCoreConnectionV1ConnectionEnd) UnmarshalBinary(b []byte) error {
	var res IbcCoreConnectionV1ConnectionEnd
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IbcCoreConnectionV1ConnectionEndCounterparty counterparty chain associated with this connection.
//
// swagger:model IbcCoreConnectionV1ConnectionEndCounterparty
type IbcCoreConnectionV1ConnectionEndCounterparty struct {

	// identifies the client on the counterparty chain associated with a given
	// connection.
	ClientID string `json:"client_id,omitempty"`

	// identifies the connection end on the counterparty chain associated with a
	// given connection.
	ConnectionID string `json:"connection_id,omitempty"`

	// prefix
	Prefix *IbcCoreConnectionV1ConnectionEndCounterpartyPrefix `json:"prefix,omitempty"`
}

// Validate validates this ibc core connection v1 connection end counterparty
func (m *IbcCoreConnectionV1ConnectionEndCounterparty) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrefix(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IbcCoreConnectionV1ConnectionEndCounterparty) validatePrefix(formats strfmt.Registry) error {
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

// ContextValidate validate this ibc core connection v1 connection end counterparty based on the context it is used
func (m *IbcCoreConnectionV1ConnectionEndCounterparty) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePrefix(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IbcCoreConnectionV1ConnectionEndCounterparty) contextValidatePrefix(ctx context.Context, formats strfmt.Registry) error {

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
func (m *IbcCoreConnectionV1ConnectionEndCounterparty) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IbcCoreConnectionV1ConnectionEndCounterparty) UnmarshalBinary(b []byte) error {
	var res IbcCoreConnectionV1ConnectionEndCounterparty
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IbcCoreConnectionV1ConnectionEndCounterpartyPrefix MerklePrefix is merkle path prefixed to the key.
// The constructed key from the Path and the key will be append(Path.KeyPath,
// append(Path.KeyPrefix, key...))
//
// commitment merkle prefix of the counterparty chain.
//
// swagger:model IbcCoreConnectionV1ConnectionEndCounterpartyPrefix
type IbcCoreConnectionV1ConnectionEndCounterpartyPrefix struct {

	// key prefix
	// Format: byte
	KeyPrefix strfmt.Base64 `json:"key_prefix,omitempty"`
}

// Validate validates this ibc core connection v1 connection end counterparty prefix
func (m *IbcCoreConnectionV1ConnectionEndCounterpartyPrefix) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ibc core connection v1 connection end counterparty prefix based on context it is used
func (m *IbcCoreConnectionV1ConnectionEndCounterpartyPrefix) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IbcCoreConnectionV1ConnectionEndCounterpartyPrefix) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IbcCoreConnectionV1ConnectionEndCounterpartyPrefix) UnmarshalBinary(b []byte) error {
	var res IbcCoreConnectionV1ConnectionEndCounterpartyPrefix
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IbcCoreConnectionV1ConnectionEndVersionsItems0 Version defines the versioning scheme used to negotiate the IBC verison in
// the connection handshake.
//
// swagger:model IbcCoreConnectionV1ConnectionEndVersionsItems0
type IbcCoreConnectionV1ConnectionEndVersionsItems0 struct {

	// list of features compatible with the specified identifier
	Features []string `json:"features"`

	// unique version identifier
	Identifier string `json:"identifier,omitempty"`
}

// Validate validates this ibc core connection v1 connection end versions items0
func (m *IbcCoreConnectionV1ConnectionEndVersionsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ibc core connection v1 connection end versions items0 based on context it is used
func (m *IbcCoreConnectionV1ConnectionEndVersionsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IbcCoreConnectionV1ConnectionEndVersionsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IbcCoreConnectionV1ConnectionEndVersionsItems0) UnmarshalBinary(b []byte) error {
	var res IbcCoreConnectionV1ConnectionEndVersionsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

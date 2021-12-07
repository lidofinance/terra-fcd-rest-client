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

// IbcCoreConnectionV1QueryConnectionResponse QueryConnectionResponse is the response type for the Query/Connection RPC
// method. Besides the connection end, it includes a proof and the height from
// which the proof was retrieved.
//
// swagger:model ibc.core.connection.v1.QueryConnectionResponse
type IbcCoreConnectionV1QueryConnectionResponse struct {

	// connection
	Connection *IbcCoreConnectionV1QueryConnectionResponseConnection `json:"connection,omitempty"`

	// merkle proof of existence
	// Format: byte
	Proof strfmt.Base64 `json:"proof,omitempty"`

	// proof height
	ProofHeight *IbcCoreConnectionV1QueryConnectionResponseProofHeight `json:"proof_height,omitempty"`
}

// Validate validates this ibc core connection v1 query connection response
func (m *IbcCoreConnectionV1QueryConnectionResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProofHeight(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IbcCoreConnectionV1QueryConnectionResponse) validateConnection(formats strfmt.Registry) error {
	if swag.IsZero(m.Connection) { // not required
		return nil
	}

	if m.Connection != nil {
		if err := m.Connection.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connection")
			}
			return err
		}
	}

	return nil
}

func (m *IbcCoreConnectionV1QueryConnectionResponse) validateProofHeight(formats strfmt.Registry) error {
	if swag.IsZero(m.ProofHeight) { // not required
		return nil
	}

	if m.ProofHeight != nil {
		if err := m.ProofHeight.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proof_height")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("proof_height")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ibc core connection v1 query connection response based on the context it is used
func (m *IbcCoreConnectionV1QueryConnectionResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConnection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProofHeight(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IbcCoreConnectionV1QueryConnectionResponse) contextValidateConnection(ctx context.Context, formats strfmt.Registry) error {

	if m.Connection != nil {
		if err := m.Connection.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connection")
			}
			return err
		}
	}

	return nil
}

func (m *IbcCoreConnectionV1QueryConnectionResponse) contextValidateProofHeight(ctx context.Context, formats strfmt.Registry) error {

	if m.ProofHeight != nil {
		if err := m.ProofHeight.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("proof_height")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("proof_height")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IbcCoreConnectionV1QueryConnectionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IbcCoreConnectionV1QueryConnectionResponse) UnmarshalBinary(b []byte) error {
	var res IbcCoreConnectionV1QueryConnectionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IbcCoreConnectionV1QueryConnectionResponseConnection connection associated with the request identifier
//
// ConnectionEnd defines a stateful object on a chain connected to another
// separate one.
// NOTE: there must only be 2 defined ConnectionEnds to establish
// a connection between two chains.
//
// swagger:model IbcCoreConnectionV1QueryConnectionResponseConnection
type IbcCoreConnectionV1QueryConnectionResponseConnection struct {

	// client associated with this connection.
	ClientID string `json:"client_id,omitempty"`

	// counterparty
	Counterparty *IbcCoreConnectionV1QueryConnectionResponseConnectionCounterparty `json:"counterparty,omitempty"`

	// delay period that must pass before a consensus state can be used for
	// packet-verification NOTE: delay period logic is only implemented by some
	// clients.
	DelayPeriod string `json:"delay_period,omitempty"`

	// current state of the connection end.
	// Enum: [STATE_UNINITIALIZED_UNSPECIFIED STATE_INIT STATE_TRYOPEN STATE_OPEN]
	State *string `json:"state,omitempty"`

	// IBC version which can be utilised to determine encodings or protocols for
	// channels or packets utilising this connection.
	Versions []*IbcCoreConnectionV1QueryConnectionResponseConnectionVersionsItems0 `json:"versions"`
}

// Validate validates this ibc core connection v1 query connection response connection
func (m *IbcCoreConnectionV1QueryConnectionResponseConnection) Validate(formats strfmt.Registry) error {
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

func (m *IbcCoreConnectionV1QueryConnectionResponseConnection) validateCounterparty(formats strfmt.Registry) error {
	if swag.IsZero(m.Counterparty) { // not required
		return nil
	}

	if m.Counterparty != nil {
		if err := m.Counterparty.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connection" + "." + "counterparty")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connection" + "." + "counterparty")
			}
			return err
		}
	}

	return nil
}

var ibcCoreConnectionV1QueryConnectionResponseConnectionTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["STATE_UNINITIALIZED_UNSPECIFIED","STATE_INIT","STATE_TRYOPEN","STATE_OPEN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ibcCoreConnectionV1QueryConnectionResponseConnectionTypeStatePropEnum = append(ibcCoreConnectionV1QueryConnectionResponseConnectionTypeStatePropEnum, v)
	}
}

const (

	// IbcCoreConnectionV1QueryConnectionResponseConnectionStateSTATEUNINITIALIZEDUNSPECIFIED captures enum value "STATE_UNINITIALIZED_UNSPECIFIED"
	IbcCoreConnectionV1QueryConnectionResponseConnectionStateSTATEUNINITIALIZEDUNSPECIFIED string = "STATE_UNINITIALIZED_UNSPECIFIED"

	// IbcCoreConnectionV1QueryConnectionResponseConnectionStateSTATEINIT captures enum value "STATE_INIT"
	IbcCoreConnectionV1QueryConnectionResponseConnectionStateSTATEINIT string = "STATE_INIT"

	// IbcCoreConnectionV1QueryConnectionResponseConnectionStateSTATETRYOPEN captures enum value "STATE_TRYOPEN"
	IbcCoreConnectionV1QueryConnectionResponseConnectionStateSTATETRYOPEN string = "STATE_TRYOPEN"

	// IbcCoreConnectionV1QueryConnectionResponseConnectionStateSTATEOPEN captures enum value "STATE_OPEN"
	IbcCoreConnectionV1QueryConnectionResponseConnectionStateSTATEOPEN string = "STATE_OPEN"
)

// prop value enum
func (m *IbcCoreConnectionV1QueryConnectionResponseConnection) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ibcCoreConnectionV1QueryConnectionResponseConnectionTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IbcCoreConnectionV1QueryConnectionResponseConnection) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("connection"+"."+"state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

func (m *IbcCoreConnectionV1QueryConnectionResponseConnection) validateVersions(formats strfmt.Registry) error {
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
					return ve.ValidateName("connection" + "." + "versions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("connection" + "." + "versions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this ibc core connection v1 query connection response connection based on the context it is used
func (m *IbcCoreConnectionV1QueryConnectionResponseConnection) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *IbcCoreConnectionV1QueryConnectionResponseConnection) contextValidateCounterparty(ctx context.Context, formats strfmt.Registry) error {

	if m.Counterparty != nil {
		if err := m.Counterparty.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connection" + "." + "counterparty")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connection" + "." + "counterparty")
			}
			return err
		}
	}

	return nil
}

func (m *IbcCoreConnectionV1QueryConnectionResponseConnection) contextValidateVersions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Versions); i++ {

		if m.Versions[i] != nil {
			if err := m.Versions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("connection" + "." + "versions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("connection" + "." + "versions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *IbcCoreConnectionV1QueryConnectionResponseConnection) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IbcCoreConnectionV1QueryConnectionResponseConnection) UnmarshalBinary(b []byte) error {
	var res IbcCoreConnectionV1QueryConnectionResponseConnection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IbcCoreConnectionV1QueryConnectionResponseConnectionCounterparty counterparty chain associated with this connection.
//
// swagger:model IbcCoreConnectionV1QueryConnectionResponseConnectionCounterparty
type IbcCoreConnectionV1QueryConnectionResponseConnectionCounterparty struct {

	// identifies the client on the counterparty chain associated with a given
	// connection.
	ClientID string `json:"client_id,omitempty"`

	// identifies the connection end on the counterparty chain associated with a
	// given connection.
	ConnectionID string `json:"connection_id,omitempty"`

	// prefix
	Prefix *IbcCoreConnectionV1QueryConnectionResponseConnectionCounterpartyPrefix `json:"prefix,omitempty"`
}

// Validate validates this ibc core connection v1 query connection response connection counterparty
func (m *IbcCoreConnectionV1QueryConnectionResponseConnectionCounterparty) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrefix(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IbcCoreConnectionV1QueryConnectionResponseConnectionCounterparty) validatePrefix(formats strfmt.Registry) error {
	if swag.IsZero(m.Prefix) { // not required
		return nil
	}

	if m.Prefix != nil {
		if err := m.Prefix.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connection" + "." + "counterparty" + "." + "prefix")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connection" + "." + "counterparty" + "." + "prefix")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ibc core connection v1 query connection response connection counterparty based on the context it is used
func (m *IbcCoreConnectionV1QueryConnectionResponseConnectionCounterparty) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePrefix(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IbcCoreConnectionV1QueryConnectionResponseConnectionCounterparty) contextValidatePrefix(ctx context.Context, formats strfmt.Registry) error {

	if m.Prefix != nil {
		if err := m.Prefix.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connection" + "." + "counterparty" + "." + "prefix")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connection" + "." + "counterparty" + "." + "prefix")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IbcCoreConnectionV1QueryConnectionResponseConnectionCounterparty) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IbcCoreConnectionV1QueryConnectionResponseConnectionCounterparty) UnmarshalBinary(b []byte) error {
	var res IbcCoreConnectionV1QueryConnectionResponseConnectionCounterparty
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IbcCoreConnectionV1QueryConnectionResponseConnectionCounterpartyPrefix MerklePrefix is merkle path prefixed to the key.
// The constructed key from the Path and the key will be append(Path.KeyPath,
// append(Path.KeyPrefix, key...))
//
// commitment merkle prefix of the counterparty chain.
//
// swagger:model IbcCoreConnectionV1QueryConnectionResponseConnectionCounterpartyPrefix
type IbcCoreConnectionV1QueryConnectionResponseConnectionCounterpartyPrefix struct {

	// key prefix
	// Format: byte
	KeyPrefix strfmt.Base64 `json:"key_prefix,omitempty"`
}

// Validate validates this ibc core connection v1 query connection response connection counterparty prefix
func (m *IbcCoreConnectionV1QueryConnectionResponseConnectionCounterpartyPrefix) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ibc core connection v1 query connection response connection counterparty prefix based on context it is used
func (m *IbcCoreConnectionV1QueryConnectionResponseConnectionCounterpartyPrefix) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IbcCoreConnectionV1QueryConnectionResponseConnectionCounterpartyPrefix) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IbcCoreConnectionV1QueryConnectionResponseConnectionCounterpartyPrefix) UnmarshalBinary(b []byte) error {
	var res IbcCoreConnectionV1QueryConnectionResponseConnectionCounterpartyPrefix
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IbcCoreConnectionV1QueryConnectionResponseConnectionVersionsItems0 Version defines the versioning scheme used to negotiate the IBC verison in
// the connection handshake.
//
// swagger:model IbcCoreConnectionV1QueryConnectionResponseConnectionVersionsItems0
type IbcCoreConnectionV1QueryConnectionResponseConnectionVersionsItems0 struct {

	// list of features compatible with the specified identifier
	Features []string `json:"features"`

	// unique version identifier
	Identifier string `json:"identifier,omitempty"`
}

// Validate validates this ibc core connection v1 query connection response connection versions items0
func (m *IbcCoreConnectionV1QueryConnectionResponseConnectionVersionsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ibc core connection v1 query connection response connection versions items0 based on context it is used
func (m *IbcCoreConnectionV1QueryConnectionResponseConnectionVersionsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IbcCoreConnectionV1QueryConnectionResponseConnectionVersionsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IbcCoreConnectionV1QueryConnectionResponseConnectionVersionsItems0) UnmarshalBinary(b []byte) error {
	var res IbcCoreConnectionV1QueryConnectionResponseConnectionVersionsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IbcCoreConnectionV1QueryConnectionResponseProofHeight height at which the proof was retrieved
//
// Normally the RevisionHeight is incremented at each height while keeping
// RevisionNumber the same. However some consensus algorithms may choose to
// reset the height in certain conditions e.g. hard forks, state-machine
// breaking changes In these cases, the RevisionNumber is incremented so that
// height continues to be monitonically increasing even as the RevisionHeight
// gets reset
//
// swagger:model IbcCoreConnectionV1QueryConnectionResponseProofHeight
type IbcCoreConnectionV1QueryConnectionResponseProofHeight struct {

	// the height within the given revision
	RevisionHeight string `json:"revision_height,omitempty"`

	// the revision that the client is currently on
	RevisionNumber string `json:"revision_number,omitempty"`
}

// Validate validates this ibc core connection v1 query connection response proof height
func (m *IbcCoreConnectionV1QueryConnectionResponseProofHeight) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ibc core connection v1 query connection response proof height based on context it is used
func (m *IbcCoreConnectionV1QueryConnectionResponseProofHeight) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IbcCoreConnectionV1QueryConnectionResponseProofHeight) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IbcCoreConnectionV1QueryConnectionResponseProofHeight) UnmarshalBinary(b []byte) error {
	var res IbcCoreConnectionV1QueryConnectionResponseProofHeight
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

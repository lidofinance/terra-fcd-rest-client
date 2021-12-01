// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IbcCoreConnectionV1QueryConnectionClientStateResponse QueryConnectionClientStateResponse is the response type for the
// Query/ConnectionClientState RPC method
//
// swagger:model ibc.core.connection.v1.QueryConnectionClientStateResponse
type IbcCoreConnectionV1QueryConnectionClientStateResponse struct {

	// identified client state
	IdentifiedClientState *IbcCoreConnectionV1QueryConnectionClientStateResponseIdentifiedClientState `json:"identified_client_state,omitempty"`

	// merkle proof of existence
	// Format: byte
	Proof strfmt.Base64 `json:"proof,omitempty"`

	// proof height
	ProofHeight *IbcCoreConnectionV1QueryConnectionClientStateResponseProofHeight `json:"proof_height,omitempty"`
}

// Validate validates this ibc core connection v1 query connection client state response
func (m *IbcCoreConnectionV1QueryConnectionClientStateResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIdentifiedClientState(formats); err != nil {
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

func (m *IbcCoreConnectionV1QueryConnectionClientStateResponse) validateIdentifiedClientState(formats strfmt.Registry) error {
	if swag.IsZero(m.IdentifiedClientState) { // not required
		return nil
	}

	if m.IdentifiedClientState != nil {
		if err := m.IdentifiedClientState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identified_client_state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("identified_client_state")
			}
			return err
		}
	}

	return nil
}

func (m *IbcCoreConnectionV1QueryConnectionClientStateResponse) validateProofHeight(formats strfmt.Registry) error {
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

// ContextValidate validate this ibc core connection v1 query connection client state response based on the context it is used
func (m *IbcCoreConnectionV1QueryConnectionClientStateResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIdentifiedClientState(ctx, formats); err != nil {
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

func (m *IbcCoreConnectionV1QueryConnectionClientStateResponse) contextValidateIdentifiedClientState(ctx context.Context, formats strfmt.Registry) error {

	if m.IdentifiedClientState != nil {
		if err := m.IdentifiedClientState.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identified_client_state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("identified_client_state")
			}
			return err
		}
	}

	return nil
}

func (m *IbcCoreConnectionV1QueryConnectionClientStateResponse) contextValidateProofHeight(ctx context.Context, formats strfmt.Registry) error {

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
func (m *IbcCoreConnectionV1QueryConnectionClientStateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IbcCoreConnectionV1QueryConnectionClientStateResponse) UnmarshalBinary(b []byte) error {
	var res IbcCoreConnectionV1QueryConnectionClientStateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IbcCoreConnectionV1QueryConnectionClientStateResponseIdentifiedClientState client state associated with the channel
//
// IdentifiedClientState defines a client state with an additional client
// identifier field.
//
// swagger:model IbcCoreConnectionV1QueryConnectionClientStateResponseIdentifiedClientState
type IbcCoreConnectionV1QueryConnectionClientStateResponseIdentifiedClientState struct {

	// client identifier
	ClientID string `json:"client_id,omitempty"`

	// client state
	ClientState *IbcCoreConnectionV1QueryConnectionClientStateResponseIdentifiedClientStateClientState `json:"client_state,omitempty"`
}

// Validate validates this ibc core connection v1 query connection client state response identified client state
func (m *IbcCoreConnectionV1QueryConnectionClientStateResponseIdentifiedClientState) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClientState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IbcCoreConnectionV1QueryConnectionClientStateResponseIdentifiedClientState) validateClientState(formats strfmt.Registry) error {
	if swag.IsZero(m.ClientState) { // not required
		return nil
	}

	if m.ClientState != nil {
		if err := m.ClientState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identified_client_state" + "." + "client_state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("identified_client_state" + "." + "client_state")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ibc core connection v1 query connection client state response identified client state based on the context it is used
func (m *IbcCoreConnectionV1QueryConnectionClientStateResponseIdentifiedClientState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClientState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IbcCoreConnectionV1QueryConnectionClientStateResponseIdentifiedClientState) contextValidateClientState(ctx context.Context, formats strfmt.Registry) error {

	if m.ClientState != nil {
		if err := m.ClientState.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("identified_client_state" + "." + "client_state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("identified_client_state" + "." + "client_state")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IbcCoreConnectionV1QueryConnectionClientStateResponseIdentifiedClientState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IbcCoreConnectionV1QueryConnectionClientStateResponseIdentifiedClientState) UnmarshalBinary(b []byte) error {
	var res IbcCoreConnectionV1QueryConnectionClientStateResponseIdentifiedClientState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IbcCoreConnectionV1QueryConnectionClientStateResponseIdentifiedClientStateClientState client state
//
// `Any` contains an arbitrary serialized protocol buffer message along with a
// URL that describes the type of the serialized message.
//
// Protobuf library provides support to pack/unpack Any values in the form
// of utility functions or additional generated methods of the Any type.
//
// Example 1: Pack and unpack a message in C++.
//
//     Foo foo = ...;
//     Any any;
//     any.PackFrom(foo);
//     ...
//     if (any.UnpackTo(&foo)) {
//       ...
//     }
//
// Example 2: Pack and unpack a message in Java.
//
//     Foo foo = ...;
//     Any any = Any.pack(foo);
//     ...
//     if (any.is(Foo.class)) {
//       foo = any.unpack(Foo.class);
//     }
//
//  Example 3: Pack and unpack a message in Python.
//
//     foo = Foo(...)
//     any = Any()
//     any.Pack(foo)
//     ...
//     if any.Is(Foo.DESCRIPTOR):
//       any.Unpack(foo)
//       ...
//
//  Example 4: Pack and unpack a message in Go
//
//      foo := &pb.Foo{...}
//      any, err := ptypes.MarshalAny(foo)
//      ...
//      foo := &pb.Foo{}
//      if err := ptypes.UnmarshalAny(any, foo); err != nil {
//        ...
//      }
//
// The pack methods provided by protobuf library will by default use
// 'type.googleapis.com/full.type.name' as the type URL and the unpack
// methods only use the fully qualified type name after the last '/'
// in the type URL, for example "foo.bar.com/x/y.z" will yield type
// name "y.z".
//
//
// JSON
// ====
// The JSON representation of an `Any` value uses the regular
// representation of the deserialized, embedded message, with an
// additional field `@type` which contains the type URL. Example:
//
//     package google.profile;
//     message Person {
//       string first_name = 1;
//       string last_name = 2;
//     }
//
//     {
//       "@type": "type.googleapis.com/google.profile.Person",
//       "firstName": <string>,
//       "lastName": <string>
//     }
//
// If the embedded message type is well-known and has a custom JSON
// representation, that representation will be embedded adding a field
// `value` which holds the custom JSON in addition to the `@type`
// field. Example (for message [google.protobuf.Duration][]):
//
//     {
//       "@type": "type.googleapis.com/google.protobuf.Duration",
//       "value": "1.212s"
//     }
//
// swagger:model IbcCoreConnectionV1QueryConnectionClientStateResponseIdentifiedClientStateClientState
type IbcCoreConnectionV1QueryConnectionClientStateResponseIdentifiedClientStateClientState struct {

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

// Validate validates this ibc core connection v1 query connection client state response identified client state client state
func (m *IbcCoreConnectionV1QueryConnectionClientStateResponseIdentifiedClientStateClientState) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ibc core connection v1 query connection client state response identified client state client state based on context it is used
func (m *IbcCoreConnectionV1QueryConnectionClientStateResponseIdentifiedClientStateClientState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IbcCoreConnectionV1QueryConnectionClientStateResponseIdentifiedClientStateClientState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IbcCoreConnectionV1QueryConnectionClientStateResponseIdentifiedClientStateClientState) UnmarshalBinary(b []byte) error {
	var res IbcCoreConnectionV1QueryConnectionClientStateResponseIdentifiedClientStateClientState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IbcCoreConnectionV1QueryConnectionClientStateResponseProofHeight height at which the proof was retrieved
//
// Normally the RevisionHeight is incremented at each height while keeping
// RevisionNumber the same. However some consensus algorithms may choose to
// reset the height in certain conditions e.g. hard forks, state-machine
// breaking changes In these cases, the RevisionNumber is incremented so that
// height continues to be monitonically increasing even as the RevisionHeight
// gets reset
//
// swagger:model IbcCoreConnectionV1QueryConnectionClientStateResponseProofHeight
type IbcCoreConnectionV1QueryConnectionClientStateResponseProofHeight struct {

	// the height within the given revision
	RevisionHeight string `json:"revision_height,omitempty"`

	// the revision that the client is currently on
	RevisionNumber string `json:"revision_number,omitempty"`
}

// Validate validates this ibc core connection v1 query connection client state response proof height
func (m *IbcCoreConnectionV1QueryConnectionClientStateResponseProofHeight) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ibc core connection v1 query connection client state response proof height based on context it is used
func (m *IbcCoreConnectionV1QueryConnectionClientStateResponseProofHeight) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IbcCoreConnectionV1QueryConnectionClientStateResponseProofHeight) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IbcCoreConnectionV1QueryConnectionClientStateResponseProofHeight) UnmarshalBinary(b []byte) error {
	var res IbcCoreConnectionV1QueryConnectionClientStateResponseProofHeight
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
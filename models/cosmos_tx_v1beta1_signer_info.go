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

// CosmosTxV1beta1SignerInfo SignerInfo describes the public key and signing mode of a single top-level
// signer.
//
// swagger:model cosmos.tx.v1beta1.SignerInfo
type CosmosTxV1beta1SignerInfo struct {

	// mode_info describes the signing mode of the signer and is a nested
	// structure to support nested multisig pubkey's
	ModeInfo *CosmosTxV1beta1ModeInfo `json:"mode_info,omitempty"`

	// public key
	PublicKey *CosmosTxV1beta1SignerInfoPublicKey `json:"public_key,omitempty"`

	// sequence is the sequence of the account, which describes the
	// number of committed transactions signed by a given address. It is used to
	// prevent replay attacks.
	Sequence string `json:"sequence,omitempty"`
}

// Validate validates this cosmos tx v1beta1 signer info
func (m *CosmosTxV1beta1SignerInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateModeInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublicKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosTxV1beta1SignerInfo) validateModeInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.ModeInfo) { // not required
		return nil
	}

	if m.ModeInfo != nil {
		if err := m.ModeInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mode_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mode_info")
			}
			return err
		}
	}

	return nil
}

func (m *CosmosTxV1beta1SignerInfo) validatePublicKey(formats strfmt.Registry) error {
	if swag.IsZero(m.PublicKey) { // not required
		return nil
	}

	if m.PublicKey != nil {
		if err := m.PublicKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("public_key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("public_key")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cosmos tx v1beta1 signer info based on the context it is used
func (m *CosmosTxV1beta1SignerInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateModeInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePublicKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosTxV1beta1SignerInfo) contextValidateModeInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.ModeInfo != nil {
		if err := m.ModeInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mode_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mode_info")
			}
			return err
		}
	}

	return nil
}

func (m *CosmosTxV1beta1SignerInfo) contextValidatePublicKey(ctx context.Context, formats strfmt.Registry) error {

	if m.PublicKey != nil {
		if err := m.PublicKey.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("public_key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("public_key")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CosmosTxV1beta1SignerInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosTxV1beta1SignerInfo) UnmarshalBinary(b []byte) error {
	var res CosmosTxV1beta1SignerInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CosmosTxV1beta1SignerInfoPublicKey `Any` contains an arbitrary serialized protocol buffer message along with a
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
// swagger:model CosmosTxV1beta1SignerInfoPublicKey
type CosmosTxV1beta1SignerInfoPublicKey struct {

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

// Validate validates this cosmos tx v1beta1 signer info public key
func (m *CosmosTxV1beta1SignerInfoPublicKey) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cosmos tx v1beta1 signer info public key based on context it is used
func (m *CosmosTxV1beta1SignerInfoPublicKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CosmosTxV1beta1SignerInfoPublicKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosTxV1beta1SignerInfoPublicKey) UnmarshalBinary(b []byte) error {
	var res CosmosTxV1beta1SignerInfoPublicKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
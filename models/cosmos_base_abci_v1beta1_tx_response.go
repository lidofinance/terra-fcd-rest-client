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

// CosmosBaseAbciV1beta1TxResponse TxResponse defines a structure containing relevant tx data and metadata. The
// tags are stringified and the log is JSON decoded.
//
// swagger:model cosmos.base.abci.v1beta1.TxResponse
type CosmosBaseAbciV1beta1TxResponse struct {

	// Response code.
	Code int64 `json:"code,omitempty"`

	// Namespace for the Code
	Codespace string `json:"codespace,omitempty"`

	// Result bytes, if any.
	Data string `json:"data,omitempty"`

	// Amount of gas consumed by transaction.
	GasUsed string `json:"gas_used,omitempty"`

	// Amount of gas requested for transaction.
	GasWanted string `json:"gas_wanted,omitempty"`

	// The block height
	Height string `json:"height,omitempty"`

	// Additional information. May be non-deterministic.
	Info string `json:"info,omitempty"`

	// The output of the application's logger (typed). May be non-deterministic.
	Logs []*CosmosBaseAbciV1beta1TxResponseLogsItems0 `json:"logs"`

	// The output of the application's logger (raw string). May be
	// non-deterministic.
	RawLog string `json:"raw_log,omitempty"`

	// Time of the previous block. For heights > 1, it's the weighted median of
	// the timestamps of the valid votes in the block.LastCommit. For height == 1,
	// it's genesis time.
	Timestamp string `json:"timestamp,omitempty"`

	// tx
	Tx *CosmosBaseAbciV1beta1TxResponseTx `json:"tx,omitempty"`

	// The transaction hash.
	Txhash string `json:"txhash,omitempty"`
}

// Validate validates this cosmos base abci v1beta1 tx response
func (m *CosmosBaseAbciV1beta1TxResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLogs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTx(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosBaseAbciV1beta1TxResponse) validateLogs(formats strfmt.Registry) error {
	if swag.IsZero(m.Logs) { // not required
		return nil
	}

	for i := 0; i < len(m.Logs); i++ {
		if swag.IsZero(m.Logs[i]) { // not required
			continue
		}

		if m.Logs[i] != nil {
			if err := m.Logs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("logs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("logs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CosmosBaseAbciV1beta1TxResponse) validateTx(formats strfmt.Registry) error {
	if swag.IsZero(m.Tx) { // not required
		return nil
	}

	if m.Tx != nil {
		if err := m.Tx.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tx")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tx")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cosmos base abci v1beta1 tx response based on the context it is used
func (m *CosmosBaseAbciV1beta1TxResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLogs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTx(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosBaseAbciV1beta1TxResponse) contextValidateLogs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Logs); i++ {

		if m.Logs[i] != nil {
			if err := m.Logs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("logs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("logs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CosmosBaseAbciV1beta1TxResponse) contextValidateTx(ctx context.Context, formats strfmt.Registry) error {

	if m.Tx != nil {
		if err := m.Tx.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tx")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tx")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CosmosBaseAbciV1beta1TxResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosBaseAbciV1beta1TxResponse) UnmarshalBinary(b []byte) error {
	var res CosmosBaseAbciV1beta1TxResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CosmosBaseAbciV1beta1TxResponseLogsItems0 ABCIMessageLog defines a structure containing an indexed tx ABCI message log.
//
// swagger:model CosmosBaseAbciV1beta1TxResponseLogsItems0
type CosmosBaseAbciV1beta1TxResponseLogsItems0 struct {

	// Events contains a slice of Event objects that were emitted during some
	// execution.
	Events []*CosmosBaseAbciV1beta1TxResponseLogsItems0EventsItems0 `json:"events"`

	// log
	Log string `json:"log,omitempty"`

	// msg index
	MsgIndex int64 `json:"msg_index,omitempty"`
}

// Validate validates this cosmos base abci v1beta1 tx response logs items0
func (m *CosmosBaseAbciV1beta1TxResponseLogsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEvents(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosBaseAbciV1beta1TxResponseLogsItems0) validateEvents(formats strfmt.Registry) error {
	if swag.IsZero(m.Events) { // not required
		return nil
	}

	for i := 0; i < len(m.Events); i++ {
		if swag.IsZero(m.Events[i]) { // not required
			continue
		}

		if m.Events[i] != nil {
			if err := m.Events[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("events" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("events" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this cosmos base abci v1beta1 tx response logs items0 based on the context it is used
func (m *CosmosBaseAbciV1beta1TxResponseLogsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEvents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosBaseAbciV1beta1TxResponseLogsItems0) contextValidateEvents(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Events); i++ {

		if m.Events[i] != nil {
			if err := m.Events[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("events" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("events" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CosmosBaseAbciV1beta1TxResponseLogsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosBaseAbciV1beta1TxResponseLogsItems0) UnmarshalBinary(b []byte) error {
	var res CosmosBaseAbciV1beta1TxResponseLogsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CosmosBaseAbciV1beta1TxResponseLogsItems0EventsItems0 StringEvent defines en Event object wrapper where all the attributes
// contain key/value pairs that are strings instead of raw bytes.
//
// swagger:model CosmosBaseAbciV1beta1TxResponseLogsItems0EventsItems0
type CosmosBaseAbciV1beta1TxResponseLogsItems0EventsItems0 struct {

	// attributes
	Attributes []*CosmosBaseAbciV1beta1TxResponseLogsItems0EventsItems0AttributesItems0 `json:"attributes"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this cosmos base abci v1beta1 tx response logs items0 events items0
func (m *CosmosBaseAbciV1beta1TxResponseLogsItems0EventsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttributes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosBaseAbciV1beta1TxResponseLogsItems0EventsItems0) validateAttributes(formats strfmt.Registry) error {
	if swag.IsZero(m.Attributes) { // not required
		return nil
	}

	for i := 0; i < len(m.Attributes); i++ {
		if swag.IsZero(m.Attributes[i]) { // not required
			continue
		}

		if m.Attributes[i] != nil {
			if err := m.Attributes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attributes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attributes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this cosmos base abci v1beta1 tx response logs items0 events items0 based on the context it is used
func (m *CosmosBaseAbciV1beta1TxResponseLogsItems0EventsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttributes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosBaseAbciV1beta1TxResponseLogsItems0EventsItems0) contextValidateAttributes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Attributes); i++ {

		if m.Attributes[i] != nil {
			if err := m.Attributes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("attributes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("attributes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CosmosBaseAbciV1beta1TxResponseLogsItems0EventsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosBaseAbciV1beta1TxResponseLogsItems0EventsItems0) UnmarshalBinary(b []byte) error {
	var res CosmosBaseAbciV1beta1TxResponseLogsItems0EventsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CosmosBaseAbciV1beta1TxResponseLogsItems0EventsItems0AttributesItems0 Attribute defines an attribute wrapper where the key and value are
// strings instead of raw bytes.
//
// swagger:model CosmosBaseAbciV1beta1TxResponseLogsItems0EventsItems0AttributesItems0
type CosmosBaseAbciV1beta1TxResponseLogsItems0EventsItems0AttributesItems0 struct {

	// key
	Key string `json:"key,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this cosmos base abci v1beta1 tx response logs items0 events items0 attributes items0
func (m *CosmosBaseAbciV1beta1TxResponseLogsItems0EventsItems0AttributesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cosmos base abci v1beta1 tx response logs items0 events items0 attributes items0 based on context it is used
func (m *CosmosBaseAbciV1beta1TxResponseLogsItems0EventsItems0AttributesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CosmosBaseAbciV1beta1TxResponseLogsItems0EventsItems0AttributesItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosBaseAbciV1beta1TxResponseLogsItems0EventsItems0AttributesItems0) UnmarshalBinary(b []byte) error {
	var res CosmosBaseAbciV1beta1TxResponseLogsItems0EventsItems0AttributesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CosmosBaseAbciV1beta1TxResponseTx `Any` contains an arbitrary serialized protocol buffer message along with a
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
// swagger:model CosmosBaseAbciV1beta1TxResponseTx
type CosmosBaseAbciV1beta1TxResponseTx struct {

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

// Validate validates this cosmos base abci v1beta1 tx response tx
func (m *CosmosBaseAbciV1beta1TxResponseTx) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cosmos base abci v1beta1 tx response tx based on context it is used
func (m *CosmosBaseAbciV1beta1TxResponseTx) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CosmosBaseAbciV1beta1TxResponseTx) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosBaseAbciV1beta1TxResponseTx) UnmarshalBinary(b []byte) error {
	var res CosmosBaseAbciV1beta1TxResponseTx
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
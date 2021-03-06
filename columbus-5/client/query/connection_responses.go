// Code generated by go-swagger; DO NOT EDIT.

package query

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConnectionReader is a Reader for the Connection structure.
type ConnectionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConnectionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConnectionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewConnectionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConnectionOK creates a ConnectionOK with default headers values
func NewConnectionOK() *ConnectionOK {
	return &ConnectionOK{}
}

/* ConnectionOK describes a response with status code 200, with default header values.

A successful response.
*/
type ConnectionOK struct {
	Payload *ConnectionOKBody
}

func (o *ConnectionOK) Error() string {
	return fmt.Sprintf("[GET /ibc/core/connection/v1/connections/{connection_id}][%d] connectionOK  %+v", 200, o.Payload)
}
func (o *ConnectionOK) GetPayload() *ConnectionOKBody {
	return o.Payload
}

func (o *ConnectionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ConnectionOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConnectionDefault creates a ConnectionDefault with default headers values
func NewConnectionDefault(code int) *ConnectionDefault {
	return &ConnectionDefault{
		_statusCode: code,
	}
}

/* ConnectionDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type ConnectionDefault struct {
	_statusCode int

	Payload *ConnectionDefaultBody
}

// Code gets the status code for the connection default response
func (o *ConnectionDefault) Code() int {
	return o._statusCode
}

func (o *ConnectionDefault) Error() string {
	return fmt.Sprintf("[GET /ibc/core/connection/v1/connections/{connection_id}][%d] Connection default  %+v", o._statusCode, o.Payload)
}
func (o *ConnectionDefault) GetPayload() *ConnectionDefaultBody {
	return o.Payload
}

func (o *ConnectionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ConnectionDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ConnectionDefaultBody connection default body
swagger:model ConnectionDefaultBody
*/
type ConnectionDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// details
	Details []*ConnectionDefaultBodyDetailsItems0 `json:"details"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this connection default body
func (o *ConnectionDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ConnectionDefaultBody) validateDetails(formats strfmt.Registry) error {
	if swag.IsZero(o.Details) { // not required
		return nil
	}

	for i := 0; i < len(o.Details); i++ {
		if swag.IsZero(o.Details[i]) { // not required
			continue
		}

		if o.Details[i] != nil {
			if err := o.Details[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Connection default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Connection default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this connection default body based on the context it is used
func (o *ConnectionDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ConnectionDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Connection default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Connection default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ConnectionDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ConnectionDefaultBody) UnmarshalBinary(b []byte) error {
	var res ConnectionDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ConnectionDefaultBodyDetailsItems0 `Any` contains an arbitrary serialized protocol buffer message along with a
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
swagger:model ConnectionDefaultBodyDetailsItems0
*/
type ConnectionDefaultBodyDetailsItems0 struct {

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

// Validate validates this connection default body details items0
func (o *ConnectionDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this connection default body details items0 based on context it is used
func (o *ConnectionDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ConnectionDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ConnectionDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res ConnectionDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ConnectionOKBody QueryConnectionResponse is the response type for the Query/Connection RPC
// method. Besides the connection end, it includes a proof and the height from
// which the proof was retrieved.
swagger:model ConnectionOKBody
*/
type ConnectionOKBody struct {

	// connection
	Connection *ConnectionOKBodyConnection `json:"connection,omitempty"`

	// merkle proof of existence
	// Format: byte
	Proof strfmt.Base64 `json:"proof,omitempty"`

	// proof height
	ProofHeight *ConnectionOKBodyProofHeight `json:"proof_height,omitempty"`
}

// Validate validates this connection o k body
func (o *ConnectionOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateConnection(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateProofHeight(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ConnectionOKBody) validateConnection(formats strfmt.Registry) error {
	if swag.IsZero(o.Connection) { // not required
		return nil
	}

	if o.Connection != nil {
		if err := o.Connection.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connectionOK" + "." + "connection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connectionOK" + "." + "connection")
			}
			return err
		}
	}

	return nil
}

func (o *ConnectionOKBody) validateProofHeight(formats strfmt.Registry) error {
	if swag.IsZero(o.ProofHeight) { // not required
		return nil
	}

	if o.ProofHeight != nil {
		if err := o.ProofHeight.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connectionOK" + "." + "proof_height")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connectionOK" + "." + "proof_height")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this connection o k body based on the context it is used
func (o *ConnectionOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateConnection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateProofHeight(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ConnectionOKBody) contextValidateConnection(ctx context.Context, formats strfmt.Registry) error {

	if o.Connection != nil {
		if err := o.Connection.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connectionOK" + "." + "connection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connectionOK" + "." + "connection")
			}
			return err
		}
	}

	return nil
}

func (o *ConnectionOKBody) contextValidateProofHeight(ctx context.Context, formats strfmt.Registry) error {

	if o.ProofHeight != nil {
		if err := o.ProofHeight.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connectionOK" + "." + "proof_height")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connectionOK" + "." + "proof_height")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ConnectionOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ConnectionOKBody) UnmarshalBinary(b []byte) error {
	var res ConnectionOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ConnectionOKBodyConnection connection associated with the request identifier
//
// ConnectionEnd defines a stateful object on a chain connected to another
// separate one.
// NOTE: there must only be 2 defined ConnectionEnds to establish
// a connection between two chains.
swagger:model ConnectionOKBodyConnection
*/
type ConnectionOKBodyConnection struct {

	// client associated with this connection.
	ClientID string `json:"client_id,omitempty"`

	// counterparty
	Counterparty *ConnectionOKBodyConnectionCounterparty `json:"counterparty,omitempty"`

	// delay period that must pass before a consensus state can be used for
	// packet-verification NOTE: delay period logic is only implemented by some
	// clients.
	DelayPeriod string `json:"delay_period,omitempty"`

	// current state of the connection end.
	// Enum: [STATE_UNINITIALIZED_UNSPECIFIED STATE_INIT STATE_TRYOPEN STATE_OPEN]
	State *string `json:"state,omitempty"`

	// IBC version which can be utilised to determine encodings or protocols for
	// channels or packets utilising this connection.
	Versions []*ConnectionOKBodyConnectionVersionsItems0 `json:"versions"`
}

// Validate validates this connection o k body connection
func (o *ConnectionOKBodyConnection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCounterparty(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateVersions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ConnectionOKBodyConnection) validateCounterparty(formats strfmt.Registry) error {
	if swag.IsZero(o.Counterparty) { // not required
		return nil
	}

	if o.Counterparty != nil {
		if err := o.Counterparty.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connectionOK" + "." + "connection" + "." + "counterparty")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connectionOK" + "." + "connection" + "." + "counterparty")
			}
			return err
		}
	}

	return nil
}

var connectionOKBodyConnectionTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["STATE_UNINITIALIZED_UNSPECIFIED","STATE_INIT","STATE_TRYOPEN","STATE_OPEN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		connectionOKBodyConnectionTypeStatePropEnum = append(connectionOKBodyConnectionTypeStatePropEnum, v)
	}
}

const (

	// ConnectionOKBodyConnectionStateSTATEUNINITIALIZEDUNSPECIFIED captures enum value "STATE_UNINITIALIZED_UNSPECIFIED"
	ConnectionOKBodyConnectionStateSTATEUNINITIALIZEDUNSPECIFIED string = "STATE_UNINITIALIZED_UNSPECIFIED"

	// ConnectionOKBodyConnectionStateSTATEINIT captures enum value "STATE_INIT"
	ConnectionOKBodyConnectionStateSTATEINIT string = "STATE_INIT"

	// ConnectionOKBodyConnectionStateSTATETRYOPEN captures enum value "STATE_TRYOPEN"
	ConnectionOKBodyConnectionStateSTATETRYOPEN string = "STATE_TRYOPEN"

	// ConnectionOKBodyConnectionStateSTATEOPEN captures enum value "STATE_OPEN"
	ConnectionOKBodyConnectionStateSTATEOPEN string = "STATE_OPEN"
)

// prop value enum
func (o *ConnectionOKBodyConnection) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, connectionOKBodyConnectionTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ConnectionOKBodyConnection) validateState(formats strfmt.Registry) error {
	if swag.IsZero(o.State) { // not required
		return nil
	}

	// value enum
	if err := o.validateStateEnum("connectionOK"+"."+"connection"+"."+"state", "body", *o.State); err != nil {
		return err
	}

	return nil
}

func (o *ConnectionOKBodyConnection) validateVersions(formats strfmt.Registry) error {
	if swag.IsZero(o.Versions) { // not required
		return nil
	}

	for i := 0; i < len(o.Versions); i++ {
		if swag.IsZero(o.Versions[i]) { // not required
			continue
		}

		if o.Versions[i] != nil {
			if err := o.Versions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("connectionOK" + "." + "connection" + "." + "versions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("connectionOK" + "." + "connection" + "." + "versions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this connection o k body connection based on the context it is used
func (o *ConnectionOKBodyConnection) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateCounterparty(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateVersions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ConnectionOKBodyConnection) contextValidateCounterparty(ctx context.Context, formats strfmt.Registry) error {

	if o.Counterparty != nil {
		if err := o.Counterparty.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connectionOK" + "." + "connection" + "." + "counterparty")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connectionOK" + "." + "connection" + "." + "counterparty")
			}
			return err
		}
	}

	return nil
}

func (o *ConnectionOKBodyConnection) contextValidateVersions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Versions); i++ {

		if o.Versions[i] != nil {
			if err := o.Versions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("connectionOK" + "." + "connection" + "." + "versions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("connectionOK" + "." + "connection" + "." + "versions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ConnectionOKBodyConnection) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ConnectionOKBodyConnection) UnmarshalBinary(b []byte) error {
	var res ConnectionOKBodyConnection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ConnectionOKBodyConnectionCounterparty counterparty chain associated with this connection.
swagger:model ConnectionOKBodyConnectionCounterparty
*/
type ConnectionOKBodyConnectionCounterparty struct {

	// identifies the client on the counterparty chain associated with a given
	// connection.
	ClientID string `json:"client_id,omitempty"`

	// identifies the connection end on the counterparty chain associated with a
	// given connection.
	ConnectionID string `json:"connection_id,omitempty"`

	// prefix
	Prefix *ConnectionOKBodyConnectionCounterpartyPrefix `json:"prefix,omitempty"`
}

// Validate validates this connection o k body connection counterparty
func (o *ConnectionOKBodyConnectionCounterparty) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePrefix(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ConnectionOKBodyConnectionCounterparty) validatePrefix(formats strfmt.Registry) error {
	if swag.IsZero(o.Prefix) { // not required
		return nil
	}

	if o.Prefix != nil {
		if err := o.Prefix.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connectionOK" + "." + "connection" + "." + "counterparty" + "." + "prefix")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connectionOK" + "." + "connection" + "." + "counterparty" + "." + "prefix")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this connection o k body connection counterparty based on the context it is used
func (o *ConnectionOKBodyConnectionCounterparty) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePrefix(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ConnectionOKBodyConnectionCounterparty) contextValidatePrefix(ctx context.Context, formats strfmt.Registry) error {

	if o.Prefix != nil {
		if err := o.Prefix.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connectionOK" + "." + "connection" + "." + "counterparty" + "." + "prefix")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connectionOK" + "." + "connection" + "." + "counterparty" + "." + "prefix")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ConnectionOKBodyConnectionCounterparty) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ConnectionOKBodyConnectionCounterparty) UnmarshalBinary(b []byte) error {
	var res ConnectionOKBodyConnectionCounterparty
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ConnectionOKBodyConnectionCounterpartyPrefix MerklePrefix is merkle path prefixed to the key.
// The constructed key from the Path and the key will be append(Path.KeyPath,
// append(Path.KeyPrefix, key...))
//
// commitment merkle prefix of the counterparty chain.
swagger:model ConnectionOKBodyConnectionCounterpartyPrefix
*/
type ConnectionOKBodyConnectionCounterpartyPrefix struct {

	// key prefix
	// Format: byte
	KeyPrefix strfmt.Base64 `json:"key_prefix,omitempty"`
}

// Validate validates this connection o k body connection counterparty prefix
func (o *ConnectionOKBodyConnectionCounterpartyPrefix) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this connection o k body connection counterparty prefix based on context it is used
func (o *ConnectionOKBodyConnectionCounterpartyPrefix) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ConnectionOKBodyConnectionCounterpartyPrefix) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ConnectionOKBodyConnectionCounterpartyPrefix) UnmarshalBinary(b []byte) error {
	var res ConnectionOKBodyConnectionCounterpartyPrefix
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ConnectionOKBodyConnectionVersionsItems0 Version defines the versioning scheme used to negotiate the IBC verison in
// the connection handshake.
swagger:model ConnectionOKBodyConnectionVersionsItems0
*/
type ConnectionOKBodyConnectionVersionsItems0 struct {

	// list of features compatible with the specified identifier
	Features []string `json:"features"`

	// unique version identifier
	Identifier string `json:"identifier,omitempty"`
}

// Validate validates this connection o k body connection versions items0
func (o *ConnectionOKBodyConnectionVersionsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this connection o k body connection versions items0 based on context it is used
func (o *ConnectionOKBodyConnectionVersionsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ConnectionOKBodyConnectionVersionsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ConnectionOKBodyConnectionVersionsItems0) UnmarshalBinary(b []byte) error {
	var res ConnectionOKBodyConnectionVersionsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ConnectionOKBodyProofHeight height at which the proof was retrieved
//
// Normally the RevisionHeight is incremented at each height while keeping
// RevisionNumber the same. However some consensus algorithms may choose to
// reset the height in certain conditions e.g. hard forks, state-machine
// breaking changes In these cases, the RevisionNumber is incremented so that
// height continues to be monitonically increasing even as the RevisionHeight
// gets reset
swagger:model ConnectionOKBodyProofHeight
*/
type ConnectionOKBodyProofHeight struct {

	// the height within the given revision
	RevisionHeight string `json:"revision_height,omitempty"`

	// the revision that the client is currently on
	RevisionNumber string `json:"revision_number,omitempty"`
}

// Validate validates this connection o k body proof height
func (o *ConnectionOKBodyProofHeight) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this connection o k body proof height based on context it is used
func (o *ConnectionOKBodyProofHeight) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ConnectionOKBodyProofHeight) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ConnectionOKBodyProofHeight) UnmarshalBinary(b []byte) error {
	var res ConnectionOKBodyProofHeight
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

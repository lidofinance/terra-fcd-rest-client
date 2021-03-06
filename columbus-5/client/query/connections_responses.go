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

// ConnectionsReader is a Reader for the Connections structure.
type ConnectionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConnectionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConnectionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewConnectionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConnectionsOK creates a ConnectionsOK with default headers values
func NewConnectionsOK() *ConnectionsOK {
	return &ConnectionsOK{}
}

/* ConnectionsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ConnectionsOK struct {
	Payload *ConnectionsOKBody
}

func (o *ConnectionsOK) Error() string {
	return fmt.Sprintf("[GET /ibc/core/connection/v1/connections][%d] connectionsOK  %+v", 200, o.Payload)
}
func (o *ConnectionsOK) GetPayload() *ConnectionsOKBody {
	return o.Payload
}

func (o *ConnectionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ConnectionsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConnectionsDefault creates a ConnectionsDefault with default headers values
func NewConnectionsDefault(code int) *ConnectionsDefault {
	return &ConnectionsDefault{
		_statusCode: code,
	}
}

/* ConnectionsDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type ConnectionsDefault struct {
	_statusCode int

	Payload *ConnectionsDefaultBody
}

// Code gets the status code for the connections default response
func (o *ConnectionsDefault) Code() int {
	return o._statusCode
}

func (o *ConnectionsDefault) Error() string {
	return fmt.Sprintf("[GET /ibc/core/connection/v1/connections][%d] Connections default  %+v", o._statusCode, o.Payload)
}
func (o *ConnectionsDefault) GetPayload() *ConnectionsDefaultBody {
	return o.Payload
}

func (o *ConnectionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ConnectionsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ConnectionsDefaultBody connections default body
swagger:model ConnectionsDefaultBody
*/
type ConnectionsDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// details
	Details []*ConnectionsDefaultBodyDetailsItems0 `json:"details"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this connections default body
func (o *ConnectionsDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ConnectionsDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("Connections default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Connections default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this connections default body based on the context it is used
func (o *ConnectionsDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ConnectionsDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Connections default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Connections default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ConnectionsDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ConnectionsDefaultBody) UnmarshalBinary(b []byte) error {
	var res ConnectionsDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ConnectionsDefaultBodyDetailsItems0 `Any` contains an arbitrary serialized protocol buffer message along with a
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
swagger:model ConnectionsDefaultBodyDetailsItems0
*/
type ConnectionsDefaultBodyDetailsItems0 struct {

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

// Validate validates this connections default body details items0
func (o *ConnectionsDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this connections default body details items0 based on context it is used
func (o *ConnectionsDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ConnectionsDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ConnectionsDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res ConnectionsDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ConnectionsOKBody QueryConnectionsResponse is the response type for the Query/Connections RPC
// method.
swagger:model ConnectionsOKBody
*/
type ConnectionsOKBody struct {

	// list of stored connections of the chain.
	Connections []*ConnectionsOKBodyConnectionsItems0 `json:"connections"`

	// height
	Height *ConnectionsOKBodyHeight `json:"height,omitempty"`

	// pagination
	Pagination *ConnectionsOKBodyPagination `json:"pagination,omitempty"`
}

// Validate validates this connections o k body
func (o *ConnectionsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateConnections(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateHeight(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePagination(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ConnectionsOKBody) validateConnections(formats strfmt.Registry) error {
	if swag.IsZero(o.Connections) { // not required
		return nil
	}

	for i := 0; i < len(o.Connections); i++ {
		if swag.IsZero(o.Connections[i]) { // not required
			continue
		}

		if o.Connections[i] != nil {
			if err := o.Connections[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("connectionsOK" + "." + "connections" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("connectionsOK" + "." + "connections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *ConnectionsOKBody) validateHeight(formats strfmt.Registry) error {
	if swag.IsZero(o.Height) { // not required
		return nil
	}

	if o.Height != nil {
		if err := o.Height.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connectionsOK" + "." + "height")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connectionsOK" + "." + "height")
			}
			return err
		}
	}

	return nil
}

func (o *ConnectionsOKBody) validatePagination(formats strfmt.Registry) error {
	if swag.IsZero(o.Pagination) { // not required
		return nil
	}

	if o.Pagination != nil {
		if err := o.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connectionsOK" + "." + "pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connectionsOK" + "." + "pagination")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this connections o k body based on the context it is used
func (o *ConnectionsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateConnections(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateHeight(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidatePagination(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ConnectionsOKBody) contextValidateConnections(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Connections); i++ {

		if o.Connections[i] != nil {
			if err := o.Connections[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("connectionsOK" + "." + "connections" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("connectionsOK" + "." + "connections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *ConnectionsOKBody) contextValidateHeight(ctx context.Context, formats strfmt.Registry) error {

	if o.Height != nil {
		if err := o.Height.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connectionsOK" + "." + "height")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connectionsOK" + "." + "height")
			}
			return err
		}
	}

	return nil
}

func (o *ConnectionsOKBody) contextValidatePagination(ctx context.Context, formats strfmt.Registry) error {

	if o.Pagination != nil {
		if err := o.Pagination.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connectionsOK" + "." + "pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connectionsOK" + "." + "pagination")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ConnectionsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ConnectionsOKBody) UnmarshalBinary(b []byte) error {
	var res ConnectionsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ConnectionsOKBodyConnectionsItems0 IdentifiedConnection defines a connection with additional connection
// identifier field.
swagger:model ConnectionsOKBodyConnectionsItems0
*/
type ConnectionsOKBodyConnectionsItems0 struct {

	// client associated with this connection.
	ClientID string `json:"client_id,omitempty"`

	// counterparty
	Counterparty *ConnectionsOKBodyConnectionsItems0Counterparty `json:"counterparty,omitempty"`

	// delay period associated with this connection.
	DelayPeriod string `json:"delay_period,omitempty"`

	// connection identifier.
	ID string `json:"id,omitempty"`

	// current state of the connection end.
	// Enum: [STATE_UNINITIALIZED_UNSPECIFIED STATE_INIT STATE_TRYOPEN STATE_OPEN]
	State *string `json:"state,omitempty"`

	// IBC version which can be utilised to determine encodings or protocols for
	// channels or packets utilising this connection
	Versions []*ConnectionsOKBodyConnectionsItems0VersionsItems0 `json:"versions"`
}

// Validate validates this connections o k body connections items0
func (o *ConnectionsOKBodyConnectionsItems0) Validate(formats strfmt.Registry) error {
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

func (o *ConnectionsOKBodyConnectionsItems0) validateCounterparty(formats strfmt.Registry) error {
	if swag.IsZero(o.Counterparty) { // not required
		return nil
	}

	if o.Counterparty != nil {
		if err := o.Counterparty.Validate(formats); err != nil {
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

var connectionsOKBodyConnectionsItems0TypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["STATE_UNINITIALIZED_UNSPECIFIED","STATE_INIT","STATE_TRYOPEN","STATE_OPEN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		connectionsOKBodyConnectionsItems0TypeStatePropEnum = append(connectionsOKBodyConnectionsItems0TypeStatePropEnum, v)
	}
}

const (

	// ConnectionsOKBodyConnectionsItems0StateSTATEUNINITIALIZEDUNSPECIFIED captures enum value "STATE_UNINITIALIZED_UNSPECIFIED"
	ConnectionsOKBodyConnectionsItems0StateSTATEUNINITIALIZEDUNSPECIFIED string = "STATE_UNINITIALIZED_UNSPECIFIED"

	// ConnectionsOKBodyConnectionsItems0StateSTATEINIT captures enum value "STATE_INIT"
	ConnectionsOKBodyConnectionsItems0StateSTATEINIT string = "STATE_INIT"

	// ConnectionsOKBodyConnectionsItems0StateSTATETRYOPEN captures enum value "STATE_TRYOPEN"
	ConnectionsOKBodyConnectionsItems0StateSTATETRYOPEN string = "STATE_TRYOPEN"

	// ConnectionsOKBodyConnectionsItems0StateSTATEOPEN captures enum value "STATE_OPEN"
	ConnectionsOKBodyConnectionsItems0StateSTATEOPEN string = "STATE_OPEN"
)

// prop value enum
func (o *ConnectionsOKBodyConnectionsItems0) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, connectionsOKBodyConnectionsItems0TypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *ConnectionsOKBodyConnectionsItems0) validateState(formats strfmt.Registry) error {
	if swag.IsZero(o.State) { // not required
		return nil
	}

	// value enum
	if err := o.validateStateEnum("state", "body", *o.State); err != nil {
		return err
	}

	return nil
}

func (o *ConnectionsOKBodyConnectionsItems0) validateVersions(formats strfmt.Registry) error {
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

// ContextValidate validate this connections o k body connections items0 based on the context it is used
func (o *ConnectionsOKBodyConnectionsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *ConnectionsOKBodyConnectionsItems0) contextValidateCounterparty(ctx context.Context, formats strfmt.Registry) error {

	if o.Counterparty != nil {
		if err := o.Counterparty.ContextValidate(ctx, formats); err != nil {
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

func (o *ConnectionsOKBodyConnectionsItems0) contextValidateVersions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Versions); i++ {

		if o.Versions[i] != nil {
			if err := o.Versions[i].ContextValidate(ctx, formats); err != nil {
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
func (o *ConnectionsOKBodyConnectionsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ConnectionsOKBodyConnectionsItems0) UnmarshalBinary(b []byte) error {
	var res ConnectionsOKBodyConnectionsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ConnectionsOKBodyConnectionsItems0Counterparty counterparty chain associated with this connection.
swagger:model ConnectionsOKBodyConnectionsItems0Counterparty
*/
type ConnectionsOKBodyConnectionsItems0Counterparty struct {

	// identifies the client on the counterparty chain associated with a given
	// connection.
	ClientID string `json:"client_id,omitempty"`

	// identifies the connection end on the counterparty chain associated with a
	// given connection.
	ConnectionID string `json:"connection_id,omitempty"`

	// prefix
	Prefix *ConnectionsOKBodyConnectionsItems0CounterpartyPrefix `json:"prefix,omitempty"`
}

// Validate validates this connections o k body connections items0 counterparty
func (o *ConnectionsOKBodyConnectionsItems0Counterparty) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePrefix(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ConnectionsOKBodyConnectionsItems0Counterparty) validatePrefix(formats strfmt.Registry) error {
	if swag.IsZero(o.Prefix) { // not required
		return nil
	}

	if o.Prefix != nil {
		if err := o.Prefix.Validate(formats); err != nil {
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

// ContextValidate validate this connections o k body connections items0 counterparty based on the context it is used
func (o *ConnectionsOKBodyConnectionsItems0Counterparty) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePrefix(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ConnectionsOKBodyConnectionsItems0Counterparty) contextValidatePrefix(ctx context.Context, formats strfmt.Registry) error {

	if o.Prefix != nil {
		if err := o.Prefix.ContextValidate(ctx, formats); err != nil {
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
func (o *ConnectionsOKBodyConnectionsItems0Counterparty) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ConnectionsOKBodyConnectionsItems0Counterparty) UnmarshalBinary(b []byte) error {
	var res ConnectionsOKBodyConnectionsItems0Counterparty
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ConnectionsOKBodyConnectionsItems0CounterpartyPrefix MerklePrefix is merkle path prefixed to the key.
// The constructed key from the Path and the key will be append(Path.KeyPath,
// append(Path.KeyPrefix, key...))
//
// commitment merkle prefix of the counterparty chain.
swagger:model ConnectionsOKBodyConnectionsItems0CounterpartyPrefix
*/
type ConnectionsOKBodyConnectionsItems0CounterpartyPrefix struct {

	// key prefix
	// Format: byte
	KeyPrefix strfmt.Base64 `json:"key_prefix,omitempty"`
}

// Validate validates this connections o k body connections items0 counterparty prefix
func (o *ConnectionsOKBodyConnectionsItems0CounterpartyPrefix) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this connections o k body connections items0 counterparty prefix based on context it is used
func (o *ConnectionsOKBodyConnectionsItems0CounterpartyPrefix) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ConnectionsOKBodyConnectionsItems0CounterpartyPrefix) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ConnectionsOKBodyConnectionsItems0CounterpartyPrefix) UnmarshalBinary(b []byte) error {
	var res ConnectionsOKBodyConnectionsItems0CounterpartyPrefix
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ConnectionsOKBodyConnectionsItems0VersionsItems0 Version defines the versioning scheme used to negotiate the IBC verison in
// the connection handshake.
swagger:model ConnectionsOKBodyConnectionsItems0VersionsItems0
*/
type ConnectionsOKBodyConnectionsItems0VersionsItems0 struct {

	// list of features compatible with the specified identifier
	Features []string `json:"features"`

	// unique version identifier
	Identifier string `json:"identifier,omitempty"`
}

// Validate validates this connections o k body connections items0 versions items0
func (o *ConnectionsOKBodyConnectionsItems0VersionsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this connections o k body connections items0 versions items0 based on context it is used
func (o *ConnectionsOKBodyConnectionsItems0VersionsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ConnectionsOKBodyConnectionsItems0VersionsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ConnectionsOKBodyConnectionsItems0VersionsItems0) UnmarshalBinary(b []byte) error {
	var res ConnectionsOKBodyConnectionsItems0VersionsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ConnectionsOKBodyHeight query block height
//
// Normally the RevisionHeight is incremented at each height while keeping
// RevisionNumber the same. However some consensus algorithms may choose to
// reset the height in certain conditions e.g. hard forks, state-machine
// breaking changes In these cases, the RevisionNumber is incremented so that
// height continues to be monitonically increasing even as the RevisionHeight
// gets reset
swagger:model ConnectionsOKBodyHeight
*/
type ConnectionsOKBodyHeight struct {

	// the height within the given revision
	RevisionHeight string `json:"revision_height,omitempty"`

	// the revision that the client is currently on
	RevisionNumber string `json:"revision_number,omitempty"`
}

// Validate validates this connections o k body height
func (o *ConnectionsOKBodyHeight) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this connections o k body height based on context it is used
func (o *ConnectionsOKBodyHeight) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ConnectionsOKBodyHeight) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ConnectionsOKBodyHeight) UnmarshalBinary(b []byte) error {
	var res ConnectionsOKBodyHeight
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ConnectionsOKBodyPagination pagination response
//
// PageResponse is to be embedded in gRPC response messages where the
// corresponding request message has used PageRequest.
//
//  message SomeResponse {
//          repeated Bar results = 1;
//          PageResponse page = 2;
//  }
swagger:model ConnectionsOKBodyPagination
*/
type ConnectionsOKBodyPagination struct {

	// next_key is the key to be passed to PageRequest.key to
	// query the next page most efficiently
	// Format: byte
	NextKey strfmt.Base64 `json:"next_key,omitempty"`

	// total is total number of results available if PageRequest.count_total
	// was set, its value is undefined otherwise
	Total string `json:"total,omitempty"`
}

// Validate validates this connections o k body pagination
func (o *ConnectionsOKBodyPagination) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this connections o k body pagination based on context it is used
func (o *ConnectionsOKBodyPagination) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ConnectionsOKBodyPagination) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ConnectionsOKBodyPagination) UnmarshalBinary(b []byte) error {
	var res ConnectionsOKBodyPagination
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

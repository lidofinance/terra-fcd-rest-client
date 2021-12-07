// Code generated by go-swagger; DO NOT EDIT.

package query

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ChannelConsensusStateReader is a Reader for the ChannelConsensusState structure.
type ChannelConsensusStateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChannelConsensusStateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChannelConsensusStateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewChannelConsensusStateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewChannelConsensusStateOK creates a ChannelConsensusStateOK with default headers values
func NewChannelConsensusStateOK() *ChannelConsensusStateOK {
	return &ChannelConsensusStateOK{}
}

/* ChannelConsensusStateOK describes a response with status code 200, with default header values.

A successful response.
*/
type ChannelConsensusStateOK struct {
	Payload *ChannelConsensusStateOKBody
}

func (o *ChannelConsensusStateOK) Error() string {
	return fmt.Sprintf("[GET /ibc/core/channel/v1/channels/{channel_id}/ports/{port_id}/consensus_state/revision/{revision_number}/height/{revision_height}][%d] channelConsensusStateOK  %+v", 200, o.Payload)
}
func (o *ChannelConsensusStateOK) GetPayload() *ChannelConsensusStateOKBody {
	return o.Payload
}

func (o *ChannelConsensusStateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ChannelConsensusStateOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChannelConsensusStateDefault creates a ChannelConsensusStateDefault with default headers values
func NewChannelConsensusStateDefault(code int) *ChannelConsensusStateDefault {
	return &ChannelConsensusStateDefault{
		_statusCode: code,
	}
}

/* ChannelConsensusStateDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type ChannelConsensusStateDefault struct {
	_statusCode int

	Payload *ChannelConsensusStateDefaultBody
}

// Code gets the status code for the channel consensus state default response
func (o *ChannelConsensusStateDefault) Code() int {
	return o._statusCode
}

func (o *ChannelConsensusStateDefault) Error() string {
	return fmt.Sprintf("[GET /ibc/core/channel/v1/channels/{channel_id}/ports/{port_id}/consensus_state/revision/{revision_number}/height/{revision_height}][%d] ChannelConsensusState default  %+v", o._statusCode, o.Payload)
}
func (o *ChannelConsensusStateDefault) GetPayload() *ChannelConsensusStateDefaultBody {
	return o.Payload
}

func (o *ChannelConsensusStateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ChannelConsensusStateDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ChannelConsensusStateDefaultBody channel consensus state default body
swagger:model ChannelConsensusStateDefaultBody
*/
type ChannelConsensusStateDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// details
	Details []*ChannelConsensusStateDefaultBodyDetailsItems0 `json:"details"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this channel consensus state default body
func (o *ChannelConsensusStateDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChannelConsensusStateDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("ChannelConsensusState default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ChannelConsensusState default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this channel consensus state default body based on the context it is used
func (o *ChannelConsensusStateDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChannelConsensusStateDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ChannelConsensusState default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ChannelConsensusState default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChannelConsensusStateDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChannelConsensusStateDefaultBody) UnmarshalBinary(b []byte) error {
	var res ChannelConsensusStateDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChannelConsensusStateDefaultBodyDetailsItems0 `Any` contains an arbitrary serialized protocol buffer message along with a
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
swagger:model ChannelConsensusStateDefaultBodyDetailsItems0
*/
type ChannelConsensusStateDefaultBodyDetailsItems0 struct {

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

// Validate validates this channel consensus state default body details items0
func (o *ChannelConsensusStateDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this channel consensus state default body details items0 based on context it is used
func (o *ChannelConsensusStateDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChannelConsensusStateDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChannelConsensusStateDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res ChannelConsensusStateDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChannelConsensusStateOKBody QueryChannelClientStateResponse is the Response type for the
// Query/QueryChannelClientState RPC method
swagger:model ChannelConsensusStateOKBody
*/
type ChannelConsensusStateOKBody struct {

	// client ID associated with the consensus state
	ClientID string `json:"client_id,omitempty"`

	// consensus state
	ConsensusState *ChannelConsensusStateOKBodyConsensusState `json:"consensus_state,omitempty"`

	// merkle proof of existence
	// Format: byte
	Proof strfmt.Base64 `json:"proof,omitempty"`

	// proof height
	ProofHeight *ChannelConsensusStateOKBodyProofHeight `json:"proof_height,omitempty"`
}

// Validate validates this channel consensus state o k body
func (o *ChannelConsensusStateOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateConsensusState(formats); err != nil {
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

func (o *ChannelConsensusStateOKBody) validateConsensusState(formats strfmt.Registry) error {
	if swag.IsZero(o.ConsensusState) { // not required
		return nil
	}

	if o.ConsensusState != nil {
		if err := o.ConsensusState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("channelConsensusStateOK" + "." + "consensus_state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("channelConsensusStateOK" + "." + "consensus_state")
			}
			return err
		}
	}

	return nil
}

func (o *ChannelConsensusStateOKBody) validateProofHeight(formats strfmt.Registry) error {
	if swag.IsZero(o.ProofHeight) { // not required
		return nil
	}

	if o.ProofHeight != nil {
		if err := o.ProofHeight.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("channelConsensusStateOK" + "." + "proof_height")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("channelConsensusStateOK" + "." + "proof_height")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this channel consensus state o k body based on the context it is used
func (o *ChannelConsensusStateOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateConsensusState(ctx, formats); err != nil {
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

func (o *ChannelConsensusStateOKBody) contextValidateConsensusState(ctx context.Context, formats strfmt.Registry) error {

	if o.ConsensusState != nil {
		if err := o.ConsensusState.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("channelConsensusStateOK" + "." + "consensus_state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("channelConsensusStateOK" + "." + "consensus_state")
			}
			return err
		}
	}

	return nil
}

func (o *ChannelConsensusStateOKBody) contextValidateProofHeight(ctx context.Context, formats strfmt.Registry) error {

	if o.ProofHeight != nil {
		if err := o.ProofHeight.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("channelConsensusStateOK" + "." + "proof_height")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("channelConsensusStateOK" + "." + "proof_height")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChannelConsensusStateOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChannelConsensusStateOKBody) UnmarshalBinary(b []byte) error {
	var res ChannelConsensusStateOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChannelConsensusStateOKBodyConsensusState consensus state associated with the channel
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
swagger:model ChannelConsensusStateOKBodyConsensusState
*/
type ChannelConsensusStateOKBodyConsensusState struct {

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

// Validate validates this channel consensus state o k body consensus state
func (o *ChannelConsensusStateOKBodyConsensusState) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this channel consensus state o k body consensus state based on context it is used
func (o *ChannelConsensusStateOKBodyConsensusState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChannelConsensusStateOKBodyConsensusState) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChannelConsensusStateOKBodyConsensusState) UnmarshalBinary(b []byte) error {
	var res ChannelConsensusStateOKBodyConsensusState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChannelConsensusStateOKBodyProofHeight height at which the proof was retrieved
//
// Normally the RevisionHeight is incremented at each height while keeping
// RevisionNumber the same. However some consensus algorithms may choose to
// reset the height in certain conditions e.g. hard forks, state-machine
// breaking changes In these cases, the RevisionNumber is incremented so that
// height continues to be monitonically increasing even as the RevisionHeight
// gets reset
swagger:model ChannelConsensusStateOKBodyProofHeight
*/
type ChannelConsensusStateOKBodyProofHeight struct {

	// the height within the given revision
	RevisionHeight string `json:"revision_height,omitempty"`

	// the revision that the client is currently on
	RevisionNumber string `json:"revision_number,omitempty"`
}

// Validate validates this channel consensus state o k body proof height
func (o *ChannelConsensusStateOKBodyProofHeight) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this channel consensus state o k body proof height based on context it is used
func (o *ChannelConsensusStateOKBodyProofHeight) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChannelConsensusStateOKBodyProofHeight) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChannelConsensusStateOKBodyProofHeight) UnmarshalBinary(b []byte) error {
	var res ChannelConsensusStateOKBodyProofHeight
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

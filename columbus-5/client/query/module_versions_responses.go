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

// ModuleVersionsReader is a Reader for the ModuleVersions structure.
type ModuleVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModuleVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewModuleVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewModuleVersionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewModuleVersionsOK creates a ModuleVersionsOK with default headers values
func NewModuleVersionsOK() *ModuleVersionsOK {
	return &ModuleVersionsOK{}
}

/* ModuleVersionsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ModuleVersionsOK struct {
	Payload *ModuleVersionsOKBody
}

func (o *ModuleVersionsOK) Error() string {
	return fmt.Sprintf("[GET /cosmos/upgrade/v1beta1/module_versions][%d] moduleVersionsOK  %+v", 200, o.Payload)
}
func (o *ModuleVersionsOK) GetPayload() *ModuleVersionsOKBody {
	return o.Payload
}

func (o *ModuleVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ModuleVersionsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewModuleVersionsDefault creates a ModuleVersionsDefault with default headers values
func NewModuleVersionsDefault(code int) *ModuleVersionsDefault {
	return &ModuleVersionsDefault{
		_statusCode: code,
	}
}

/* ModuleVersionsDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type ModuleVersionsDefault struct {
	_statusCode int

	Payload *ModuleVersionsDefaultBody
}

// Code gets the status code for the module versions default response
func (o *ModuleVersionsDefault) Code() int {
	return o._statusCode
}

func (o *ModuleVersionsDefault) Error() string {
	return fmt.Sprintf("[GET /cosmos/upgrade/v1beta1/module_versions][%d] ModuleVersions default  %+v", o._statusCode, o.Payload)
}
func (o *ModuleVersionsDefault) GetPayload() *ModuleVersionsDefaultBody {
	return o.Payload
}

func (o *ModuleVersionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ModuleVersionsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ModuleVersionsDefaultBody module versions default body
swagger:model ModuleVersionsDefaultBody
*/
type ModuleVersionsDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// details
	Details []*ModuleVersionsDefaultBodyDetailsItems0 `json:"details"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this module versions default body
func (o *ModuleVersionsDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ModuleVersionsDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("ModuleVersions default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ModuleVersions default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this module versions default body based on the context it is used
func (o *ModuleVersionsDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ModuleVersionsDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ModuleVersions default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ModuleVersions default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ModuleVersionsDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ModuleVersionsDefaultBody) UnmarshalBinary(b []byte) error {
	var res ModuleVersionsDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ModuleVersionsDefaultBodyDetailsItems0 `Any` contains an arbitrary serialized protocol buffer message along with a
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
swagger:model ModuleVersionsDefaultBodyDetailsItems0
*/
type ModuleVersionsDefaultBodyDetailsItems0 struct {

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

// Validate validates this module versions default body details items0
func (o *ModuleVersionsDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this module versions default body details items0 based on context it is used
func (o *ModuleVersionsDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ModuleVersionsDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ModuleVersionsDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res ModuleVersionsDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ModuleVersionsOKBody QueryModuleVersionsResponse is the response type for the Query/ModuleVersions
// RPC method.
swagger:model ModuleVersionsOKBody
*/
type ModuleVersionsOKBody struct {

	// module_versions is a list of module names with their consensus versions.
	ModuleVersions []*ModuleVersionsOKBodyModuleVersionsItems0 `json:"module_versions"`
}

// Validate validates this module versions o k body
func (o *ModuleVersionsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateModuleVersions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ModuleVersionsOKBody) validateModuleVersions(formats strfmt.Registry) error {
	if swag.IsZero(o.ModuleVersions) { // not required
		return nil
	}

	for i := 0; i < len(o.ModuleVersions); i++ {
		if swag.IsZero(o.ModuleVersions[i]) { // not required
			continue
		}

		if o.ModuleVersions[i] != nil {
			if err := o.ModuleVersions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("moduleVersionsOK" + "." + "module_versions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("moduleVersionsOK" + "." + "module_versions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this module versions o k body based on the context it is used
func (o *ModuleVersionsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateModuleVersions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ModuleVersionsOKBody) contextValidateModuleVersions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.ModuleVersions); i++ {

		if o.ModuleVersions[i] != nil {
			if err := o.ModuleVersions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("moduleVersionsOK" + "." + "module_versions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("moduleVersionsOK" + "." + "module_versions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ModuleVersionsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ModuleVersionsOKBody) UnmarshalBinary(b []byte) error {
	var res ModuleVersionsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ModuleVersionsOKBodyModuleVersionsItems0 ModuleVersion specifies a module and its consensus version.
swagger:model ModuleVersionsOKBodyModuleVersionsItems0
*/
type ModuleVersionsOKBodyModuleVersionsItems0 struct {

	// name of the app module
	Name string `json:"name,omitempty"`

	// consensus version of the app module
	Version string `json:"version,omitempty"`
}

// Validate validates this module versions o k body module versions items0
func (o *ModuleVersionsOKBodyModuleVersionsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this module versions o k body module versions items0 based on context it is used
func (o *ModuleVersionsOKBodyModuleVersionsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ModuleVersionsOKBodyModuleVersionsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ModuleVersionsOKBodyModuleVersionsItems0) UnmarshalBinary(b []byte) error {
	var res ModuleVersionsOKBodyModuleVersionsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

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

// ValidatorDelegationsReader is a Reader for the ValidatorDelegations structure.
type ValidatorDelegationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidatorDelegationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewValidatorDelegationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewValidatorDelegationsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewValidatorDelegationsOK creates a ValidatorDelegationsOK with default headers values
func NewValidatorDelegationsOK() *ValidatorDelegationsOK {
	return &ValidatorDelegationsOK{}
}

/* ValidatorDelegationsOK describes a response with status code 200, with default header values.

A successful response.
*/
type ValidatorDelegationsOK struct {
	Payload *ValidatorDelegationsOKBody
}

func (o *ValidatorDelegationsOK) Error() string {
	return fmt.Sprintf("[GET /cosmos/staking/v1beta1/validators/{validator_addr}/delegations][%d] validatorDelegationsOK  %+v", 200, o.Payload)
}
func (o *ValidatorDelegationsOK) GetPayload() *ValidatorDelegationsOKBody {
	return o.Payload
}

func (o *ValidatorDelegationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ValidatorDelegationsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidatorDelegationsDefault creates a ValidatorDelegationsDefault with default headers values
func NewValidatorDelegationsDefault(code int) *ValidatorDelegationsDefault {
	return &ValidatorDelegationsDefault{
		_statusCode: code,
	}
}

/* ValidatorDelegationsDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type ValidatorDelegationsDefault struct {
	_statusCode int

	Payload *ValidatorDelegationsDefaultBody
}

// Code gets the status code for the validator delegations default response
func (o *ValidatorDelegationsDefault) Code() int {
	return o._statusCode
}

func (o *ValidatorDelegationsDefault) Error() string {
	return fmt.Sprintf("[GET /cosmos/staking/v1beta1/validators/{validator_addr}/delegations][%d] ValidatorDelegations default  %+v", o._statusCode, o.Payload)
}
func (o *ValidatorDelegationsDefault) GetPayload() *ValidatorDelegationsDefaultBody {
	return o.Payload
}

func (o *ValidatorDelegationsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ValidatorDelegationsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ValidatorDelegationsDefaultBody validator delegations default body
swagger:model ValidatorDelegationsDefaultBody
*/
type ValidatorDelegationsDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// details
	Details []*ValidatorDelegationsDefaultBodyDetailsItems0 `json:"details"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this validator delegations default body
func (o *ValidatorDelegationsDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ValidatorDelegationsDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("ValidatorDelegations default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ValidatorDelegations default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this validator delegations default body based on the context it is used
func (o *ValidatorDelegationsDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ValidatorDelegationsDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ValidatorDelegations default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ValidatorDelegations default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ValidatorDelegationsDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ValidatorDelegationsDefaultBody) UnmarshalBinary(b []byte) error {
	var res ValidatorDelegationsDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ValidatorDelegationsDefaultBodyDetailsItems0 `Any` contains an arbitrary serialized protocol buffer message along with a
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
swagger:model ValidatorDelegationsDefaultBodyDetailsItems0
*/
type ValidatorDelegationsDefaultBodyDetailsItems0 struct {

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

// Validate validates this validator delegations default body details items0
func (o *ValidatorDelegationsDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this validator delegations default body details items0 based on context it is used
func (o *ValidatorDelegationsDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ValidatorDelegationsDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ValidatorDelegationsDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res ValidatorDelegationsDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ValidatorDelegationsOKBody QueryValidatorDelegationsResponse is response type for the
// Query/ValidatorDelegations RPC method
swagger:model ValidatorDelegationsOKBody
*/
type ValidatorDelegationsOKBody struct {

	// delegation responses
	DelegationResponses []*ValidatorDelegationsOKBodyDelegationResponsesItems0 `json:"delegation_responses"`

	// pagination
	Pagination *ValidatorDelegationsOKBodyPagination `json:"pagination,omitempty"`
}

// Validate validates this validator delegations o k body
func (o *ValidatorDelegationsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDelegationResponses(formats); err != nil {
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

func (o *ValidatorDelegationsOKBody) validateDelegationResponses(formats strfmt.Registry) error {
	if swag.IsZero(o.DelegationResponses) { // not required
		return nil
	}

	for i := 0; i < len(o.DelegationResponses); i++ {
		if swag.IsZero(o.DelegationResponses[i]) { // not required
			continue
		}

		if o.DelegationResponses[i] != nil {
			if err := o.DelegationResponses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("validatorDelegationsOK" + "." + "delegation_responses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("validatorDelegationsOK" + "." + "delegation_responses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *ValidatorDelegationsOKBody) validatePagination(formats strfmt.Registry) error {
	if swag.IsZero(o.Pagination) { // not required
		return nil
	}

	if o.Pagination != nil {
		if err := o.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("validatorDelegationsOK" + "." + "pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("validatorDelegationsOK" + "." + "pagination")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this validator delegations o k body based on the context it is used
func (o *ValidatorDelegationsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDelegationResponses(ctx, formats); err != nil {
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

func (o *ValidatorDelegationsOKBody) contextValidateDelegationResponses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.DelegationResponses); i++ {

		if o.DelegationResponses[i] != nil {
			if err := o.DelegationResponses[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("validatorDelegationsOK" + "." + "delegation_responses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("validatorDelegationsOK" + "." + "delegation_responses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *ValidatorDelegationsOKBody) contextValidatePagination(ctx context.Context, formats strfmt.Registry) error {

	if o.Pagination != nil {
		if err := o.Pagination.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("validatorDelegationsOK" + "." + "pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("validatorDelegationsOK" + "." + "pagination")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ValidatorDelegationsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ValidatorDelegationsOKBody) UnmarshalBinary(b []byte) error {
	var res ValidatorDelegationsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ValidatorDelegationsOKBodyDelegationResponsesItems0 DelegationResponse is equivalent to Delegation except that it contains a
// balance in addition to shares which is more suitable for client responses.
swagger:model ValidatorDelegationsOKBodyDelegationResponsesItems0
*/
type ValidatorDelegationsOKBodyDelegationResponsesItems0 struct {

	// balance
	Balance *ValidatorDelegationsOKBodyDelegationResponsesItems0Balance `json:"balance,omitempty"`

	// delegation
	Delegation *ValidatorDelegationsOKBodyDelegationResponsesItems0Delegation `json:"delegation,omitempty"`
}

// Validate validates this validator delegations o k body delegation responses items0
func (o *ValidatorDelegationsOKBodyDelegationResponsesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBalance(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDelegation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ValidatorDelegationsOKBodyDelegationResponsesItems0) validateBalance(formats strfmt.Registry) error {
	if swag.IsZero(o.Balance) { // not required
		return nil
	}

	if o.Balance != nil {
		if err := o.Balance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("balance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("balance")
			}
			return err
		}
	}

	return nil
}

func (o *ValidatorDelegationsOKBodyDelegationResponsesItems0) validateDelegation(formats strfmt.Registry) error {
	if swag.IsZero(o.Delegation) { // not required
		return nil
	}

	if o.Delegation != nil {
		if err := o.Delegation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("delegation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("delegation")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this validator delegations o k body delegation responses items0 based on the context it is used
func (o *ValidatorDelegationsOKBodyDelegationResponsesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateBalance(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateDelegation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ValidatorDelegationsOKBodyDelegationResponsesItems0) contextValidateBalance(ctx context.Context, formats strfmt.Registry) error {

	if o.Balance != nil {
		if err := o.Balance.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("balance")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("balance")
			}
			return err
		}
	}

	return nil
}

func (o *ValidatorDelegationsOKBodyDelegationResponsesItems0) contextValidateDelegation(ctx context.Context, formats strfmt.Registry) error {

	if o.Delegation != nil {
		if err := o.Delegation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("delegation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("delegation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ValidatorDelegationsOKBodyDelegationResponsesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ValidatorDelegationsOKBodyDelegationResponsesItems0) UnmarshalBinary(b []byte) error {
	var res ValidatorDelegationsOKBodyDelegationResponsesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ValidatorDelegationsOKBodyDelegationResponsesItems0Balance Coin defines a token with a denomination and an amount.
//
// NOTE: The amount field is an Int which implements the custom method
// signatures required by gogoproto.
swagger:model ValidatorDelegationsOKBodyDelegationResponsesItems0Balance
*/
type ValidatorDelegationsOKBodyDelegationResponsesItems0Balance struct {

	// amount
	Amount string `json:"amount,omitempty"`

	// denom
	Denom string `json:"denom,omitempty"`
}

// Validate validates this validator delegations o k body delegation responses items0 balance
func (o *ValidatorDelegationsOKBodyDelegationResponsesItems0Balance) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this validator delegations o k body delegation responses items0 balance based on context it is used
func (o *ValidatorDelegationsOKBodyDelegationResponsesItems0Balance) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ValidatorDelegationsOKBodyDelegationResponsesItems0Balance) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ValidatorDelegationsOKBodyDelegationResponsesItems0Balance) UnmarshalBinary(b []byte) error {
	var res ValidatorDelegationsOKBodyDelegationResponsesItems0Balance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ValidatorDelegationsOKBodyDelegationResponsesItems0Delegation Delegation represents the bond with tokens held by an account. It is
// owned by one delegator, and is associated with the voting power of one
// validator.
swagger:model ValidatorDelegationsOKBodyDelegationResponsesItems0Delegation
*/
type ValidatorDelegationsOKBodyDelegationResponsesItems0Delegation struct {

	// delegator_address is the bech32-encoded address of the delegator.
	DelegatorAddress string `json:"delegator_address,omitempty"`

	// shares define the delegation shares received.
	Shares string `json:"shares,omitempty"`

	// validator_address is the bech32-encoded address of the validator.
	ValidatorAddress string `json:"validator_address,omitempty"`
}

// Validate validates this validator delegations o k body delegation responses items0 delegation
func (o *ValidatorDelegationsOKBodyDelegationResponsesItems0Delegation) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this validator delegations o k body delegation responses items0 delegation based on context it is used
func (o *ValidatorDelegationsOKBodyDelegationResponsesItems0Delegation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ValidatorDelegationsOKBodyDelegationResponsesItems0Delegation) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ValidatorDelegationsOKBodyDelegationResponsesItems0Delegation) UnmarshalBinary(b []byte) error {
	var res ValidatorDelegationsOKBodyDelegationResponsesItems0Delegation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ValidatorDelegationsOKBodyPagination pagination defines the pagination in the response.
swagger:model ValidatorDelegationsOKBodyPagination
*/
type ValidatorDelegationsOKBodyPagination struct {

	// next_key is the key to be passed to PageRequest.key to
	// query the next page most efficiently
	// Format: byte
	NextKey strfmt.Base64 `json:"next_key,omitempty"`

	// total is total number of results available if PageRequest.count_total
	// was set, its value is undefined otherwise
	Total string `json:"total,omitempty"`
}

// Validate validates this validator delegations o k body pagination
func (o *ValidatorDelegationsOKBodyPagination) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this validator delegations o k body pagination based on context it is used
func (o *ValidatorDelegationsOKBodyPagination) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ValidatorDelegationsOKBodyPagination) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ValidatorDelegationsOKBodyPagination) UnmarshalBinary(b []byte) error {
	var res ValidatorDelegationsOKBodyPagination
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

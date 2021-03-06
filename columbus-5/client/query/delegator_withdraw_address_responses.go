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

// DelegatorWithdrawAddressReader is a Reader for the DelegatorWithdrawAddress structure.
type DelegatorWithdrawAddressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DelegatorWithdrawAddressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDelegatorWithdrawAddressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDelegatorWithdrawAddressDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDelegatorWithdrawAddressOK creates a DelegatorWithdrawAddressOK with default headers values
func NewDelegatorWithdrawAddressOK() *DelegatorWithdrawAddressOK {
	return &DelegatorWithdrawAddressOK{}
}

/* DelegatorWithdrawAddressOK describes a response with status code 200, with default header values.

A successful response.
*/
type DelegatorWithdrawAddressOK struct {
	Payload *DelegatorWithdrawAddressOKBody
}

func (o *DelegatorWithdrawAddressOK) Error() string {
	return fmt.Sprintf("[GET /cosmos/distribution/v1beta1/delegators/{delegator_address}/withdraw_address][%d] delegatorWithdrawAddressOK  %+v", 200, o.Payload)
}
func (o *DelegatorWithdrawAddressOK) GetPayload() *DelegatorWithdrawAddressOKBody {
	return o.Payload
}

func (o *DelegatorWithdrawAddressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DelegatorWithdrawAddressOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDelegatorWithdrawAddressDefault creates a DelegatorWithdrawAddressDefault with default headers values
func NewDelegatorWithdrawAddressDefault(code int) *DelegatorWithdrawAddressDefault {
	return &DelegatorWithdrawAddressDefault{
		_statusCode: code,
	}
}

/* DelegatorWithdrawAddressDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type DelegatorWithdrawAddressDefault struct {
	_statusCode int

	Payload *DelegatorWithdrawAddressDefaultBody
}

// Code gets the status code for the delegator withdraw address default response
func (o *DelegatorWithdrawAddressDefault) Code() int {
	return o._statusCode
}

func (o *DelegatorWithdrawAddressDefault) Error() string {
	return fmt.Sprintf("[GET /cosmos/distribution/v1beta1/delegators/{delegator_address}/withdraw_address][%d] DelegatorWithdrawAddress default  %+v", o._statusCode, o.Payload)
}
func (o *DelegatorWithdrawAddressDefault) GetPayload() *DelegatorWithdrawAddressDefaultBody {
	return o.Payload
}

func (o *DelegatorWithdrawAddressDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DelegatorWithdrawAddressDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*DelegatorWithdrawAddressDefaultBody delegator withdraw address default body
swagger:model DelegatorWithdrawAddressDefaultBody
*/
type DelegatorWithdrawAddressDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// details
	Details []*DelegatorWithdrawAddressDefaultBodyDetailsItems0 `json:"details"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this delegator withdraw address default body
func (o *DelegatorWithdrawAddressDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DelegatorWithdrawAddressDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("DelegatorWithdrawAddress default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("DelegatorWithdrawAddress default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this delegator withdraw address default body based on the context it is used
func (o *DelegatorWithdrawAddressDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DelegatorWithdrawAddressDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("DelegatorWithdrawAddress default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("DelegatorWithdrawAddress default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DelegatorWithdrawAddressDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DelegatorWithdrawAddressDefaultBody) UnmarshalBinary(b []byte) error {
	var res DelegatorWithdrawAddressDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DelegatorWithdrawAddressDefaultBodyDetailsItems0 delegator withdraw address default body details items0
swagger:model DelegatorWithdrawAddressDefaultBodyDetailsItems0
*/
type DelegatorWithdrawAddressDefaultBodyDetailsItems0 struct {

	// type url
	TypeURL string `json:"type_url,omitempty"`

	// value
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this delegator withdraw address default body details items0
func (o *DelegatorWithdrawAddressDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delegator withdraw address default body details items0 based on context it is used
func (o *DelegatorWithdrawAddressDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DelegatorWithdrawAddressDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DelegatorWithdrawAddressDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res DelegatorWithdrawAddressDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*DelegatorWithdrawAddressOKBody QueryDelegatorWithdrawAddressResponse is the response type for the
// Query/DelegatorWithdrawAddress RPC method.
swagger:model DelegatorWithdrawAddressOKBody
*/
type DelegatorWithdrawAddressOKBody struct {

	// withdraw_address defines the delegator address to query for.
	WithdrawAddress string `json:"withdraw_address,omitempty"`
}

// Validate validates this delegator withdraw address o k body
func (o *DelegatorWithdrawAddressOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this delegator withdraw address o k body based on context it is used
func (o *DelegatorWithdrawAddressOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *DelegatorWithdrawAddressOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DelegatorWithdrawAddressOKBody) UnmarshalBinary(b []byte) error {
	var res DelegatorWithdrawAddressOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

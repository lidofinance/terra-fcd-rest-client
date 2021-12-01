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

// SupplyOfReader is a Reader for the SupplyOf structure.
type SupplyOfReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SupplyOfReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSupplyOfOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSupplyOfDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSupplyOfOK creates a SupplyOfOK with default headers values
func NewSupplyOfOK() *SupplyOfOK {
	return &SupplyOfOK{}
}

/* SupplyOfOK describes a response with status code 200, with default header values.

A successful response.
*/
type SupplyOfOK struct {
	Payload *SupplyOfOKBody
}

func (o *SupplyOfOK) Error() string {
	return fmt.Sprintf("[GET /cosmos/bank/v1beta1/supply/{denom}][%d] supplyOfOK  %+v", 200, o.Payload)
}
func (o *SupplyOfOK) GetPayload() *SupplyOfOKBody {
	return o.Payload
}

func (o *SupplyOfOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SupplyOfOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSupplyOfDefault creates a SupplyOfDefault with default headers values
func NewSupplyOfDefault(code int) *SupplyOfDefault {
	return &SupplyOfDefault{
		_statusCode: code,
	}
}

/* SupplyOfDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type SupplyOfDefault struct {
	_statusCode int

	Payload *SupplyOfDefaultBody
}

// Code gets the status code for the supply of default response
func (o *SupplyOfDefault) Code() int {
	return o._statusCode
}

func (o *SupplyOfDefault) Error() string {
	return fmt.Sprintf("[GET /cosmos/bank/v1beta1/supply/{denom}][%d] SupplyOf default  %+v", o._statusCode, o.Payload)
}
func (o *SupplyOfDefault) GetPayload() *SupplyOfDefaultBody {
	return o.Payload
}

func (o *SupplyOfDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SupplyOfDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*SupplyOfDefaultBody supply of default body
swagger:model SupplyOfDefaultBody
*/
type SupplyOfDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// details
	Details []*SupplyOfDefaultBodyDetailsItems0 `json:"details"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this supply of default body
func (o *SupplyOfDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SupplyOfDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("SupplyOf default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("SupplyOf default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this supply of default body based on the context it is used
func (o *SupplyOfDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SupplyOfDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("SupplyOf default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("SupplyOf default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *SupplyOfDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SupplyOfDefaultBody) UnmarshalBinary(b []byte) error {
	var res SupplyOfDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*SupplyOfDefaultBodyDetailsItems0 supply of default body details items0
swagger:model SupplyOfDefaultBodyDetailsItems0
*/
type SupplyOfDefaultBodyDetailsItems0 struct {

	// type url
	TypeURL string `json:"type_url,omitempty"`

	// value
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this supply of default body details items0
func (o *SupplyOfDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this supply of default body details items0 based on context it is used
func (o *SupplyOfDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SupplyOfDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SupplyOfDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res SupplyOfDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*SupplyOfOKBody QuerySupplyOfResponse is the response type for the Query/SupplyOf RPC method.
swagger:model SupplyOfOKBody
*/
type SupplyOfOKBody struct {

	// amount
	Amount *SupplyOfOKBodyAmount `json:"amount,omitempty"`
}

// Validate validates this supply of o k body
func (o *SupplyOfOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SupplyOfOKBody) validateAmount(formats strfmt.Registry) error {
	if swag.IsZero(o.Amount) { // not required
		return nil
	}

	if o.Amount != nil {
		if err := o.Amount.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("supplyOfOK" + "." + "amount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("supplyOfOK" + "." + "amount")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this supply of o k body based on the context it is used
func (o *SupplyOfOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SupplyOfOKBody) contextValidateAmount(ctx context.Context, formats strfmt.Registry) error {

	if o.Amount != nil {
		if err := o.Amount.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("supplyOfOK" + "." + "amount")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("supplyOfOK" + "." + "amount")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SupplyOfOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SupplyOfOKBody) UnmarshalBinary(b []byte) error {
	var res SupplyOfOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*SupplyOfOKBodyAmount Coin defines a token with a denomination and an amount.
//
// NOTE: The amount field is an Int which implements the custom method
// signatures required by gogoproto.
swagger:model SupplyOfOKBodyAmount
*/
type SupplyOfOKBodyAmount struct {

	// amount
	Amount string `json:"amount,omitempty"`

	// denom
	Denom string `json:"denom,omitempty"`
}

// Validate validates this supply of o k body amount
func (o *SupplyOfOKBodyAmount) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this supply of o k body amount based on context it is used
func (o *SupplyOfOKBodyAmount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SupplyOfOKBodyAmount) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SupplyOfOKBodyAmount) UnmarshalBinary(b []byte) error {
	var res SupplyOfOKBodyAmount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
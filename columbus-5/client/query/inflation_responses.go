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

// InflationReader is a Reader for the Inflation structure.
type InflationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InflationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInflationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewInflationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewInflationOK creates a InflationOK with default headers values
func NewInflationOK() *InflationOK {
	return &InflationOK{}
}

/* InflationOK describes a response with status code 200, with default header values.

A successful response.
*/
type InflationOK struct {
	Payload *InflationOKBody
}

func (o *InflationOK) Error() string {
	return fmt.Sprintf("[GET /cosmos/mint/v1beta1/inflation][%d] inflationOK  %+v", 200, o.Payload)
}
func (o *InflationOK) GetPayload() *InflationOKBody {
	return o.Payload
}

func (o *InflationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(InflationOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInflationDefault creates a InflationDefault with default headers values
func NewInflationDefault(code int) *InflationDefault {
	return &InflationDefault{
		_statusCode: code,
	}
}

/* InflationDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type InflationDefault struct {
	_statusCode int

	Payload *InflationDefaultBody
}

// Code gets the status code for the inflation default response
func (o *InflationDefault) Code() int {
	return o._statusCode
}

func (o *InflationDefault) Error() string {
	return fmt.Sprintf("[GET /cosmos/mint/v1beta1/inflation][%d] Inflation default  %+v", o._statusCode, o.Payload)
}
func (o *InflationDefault) GetPayload() *InflationDefaultBody {
	return o.Payload
}

func (o *InflationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(InflationDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*InflationDefaultBody inflation default body
swagger:model InflationDefaultBody
*/
type InflationDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// details
	Details []*InflationDefaultBodyDetailsItems0 `json:"details"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this inflation default body
func (o *InflationDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *InflationDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("Inflation default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Inflation default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this inflation default body based on the context it is used
func (o *InflationDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *InflationDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Inflation default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Inflation default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *InflationDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InflationDefaultBody) UnmarshalBinary(b []byte) error {
	var res InflationDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*InflationDefaultBodyDetailsItems0 inflation default body details items0
swagger:model InflationDefaultBodyDetailsItems0
*/
type InflationDefaultBodyDetailsItems0 struct {

	// type url
	TypeURL string `json:"type_url,omitempty"`

	// value
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this inflation default body details items0
func (o *InflationDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this inflation default body details items0 based on context it is used
func (o *InflationDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *InflationDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InflationDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res InflationDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*InflationOKBody QueryInflationResponse is the response type for the Query/Inflation RPC
// method.
swagger:model InflationOKBody
*/
type InflationOKBody struct {

	// inflation is the current minting inflation value.
	// Format: byte
	Inflation strfmt.Base64 `json:"inflation,omitempty"`
}

// Validate validates this inflation o k body
func (o *InflationOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this inflation o k body based on context it is used
func (o *InflationOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *InflationOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *InflationOKBody) UnmarshalBinary(b []byte) error {
	var res InflationOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

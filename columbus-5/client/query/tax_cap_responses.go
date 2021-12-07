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

// TaxCapReader is a Reader for the TaxCap structure.
type TaxCapReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TaxCapReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTaxCapOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTaxCapDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTaxCapOK creates a TaxCapOK with default headers values
func NewTaxCapOK() *TaxCapOK {
	return &TaxCapOK{}
}

/* TaxCapOK describes a response with status code 200, with default header values.

A successful response.
*/
type TaxCapOK struct {
	Payload *TaxCapOKBody
}

func (o *TaxCapOK) Error() string {
	return fmt.Sprintf("[GET /terra/treasury/v1beta1/tax_caps/{denom}][%d] taxCapOK  %+v", 200, o.Payload)
}
func (o *TaxCapOK) GetPayload() *TaxCapOKBody {
	return o.Payload
}

func (o *TaxCapOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(TaxCapOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTaxCapDefault creates a TaxCapDefault with default headers values
func NewTaxCapDefault(code int) *TaxCapDefault {
	return &TaxCapDefault{
		_statusCode: code,
	}
}

/* TaxCapDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type TaxCapDefault struct {
	_statusCode int

	Payload *TaxCapDefaultBody
}

// Code gets the status code for the tax cap default response
func (o *TaxCapDefault) Code() int {
	return o._statusCode
}

func (o *TaxCapDefault) Error() string {
	return fmt.Sprintf("[GET /terra/treasury/v1beta1/tax_caps/{denom}][%d] TaxCap default  %+v", o._statusCode, o.Payload)
}
func (o *TaxCapDefault) GetPayload() *TaxCapDefaultBody {
	return o.Payload
}

func (o *TaxCapDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(TaxCapDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*TaxCapDefaultBody tax cap default body
swagger:model TaxCapDefaultBody
*/
type TaxCapDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// details
	Details []*TaxCapDefaultBodyDetailsItems0 `json:"details"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this tax cap default body
func (o *TaxCapDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TaxCapDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("TaxCap default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("TaxCap default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this tax cap default body based on the context it is used
func (o *TaxCapDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TaxCapDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("TaxCap default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("TaxCap default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *TaxCapDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TaxCapDefaultBody) UnmarshalBinary(b []byte) error {
	var res TaxCapDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*TaxCapDefaultBodyDetailsItems0 tax cap default body details items0
swagger:model TaxCapDefaultBodyDetailsItems0
*/
type TaxCapDefaultBodyDetailsItems0 struct {

	// type url
	TypeURL string `json:"type_url,omitempty"`

	// value
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this tax cap default body details items0
func (o *TaxCapDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this tax cap default body details items0 based on context it is used
func (o *TaxCapDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *TaxCapDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TaxCapDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res TaxCapDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*TaxCapOKBody QueryTaxCapResponse is response type for the
// Query/TaxCap RPC method.
swagger:model TaxCapOKBody
*/
type TaxCapOKBody struct {

	// tax cap
	TaxCap string `json:"tax_cap,omitempty"`
}

// Validate validates this tax cap o k body
func (o *TaxCapOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this tax cap o k body based on context it is used
func (o *TaxCapOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *TaxCapOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *TaxCapOKBody) UnmarshalBinary(b []byte) error {
	var res TaxCapOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

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

// MintParamsReader is a Reader for the MintParams structure.
type MintParamsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MintParamsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMintParamsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMintParamsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMintParamsOK creates a MintParamsOK with default headers values
func NewMintParamsOK() *MintParamsOK {
	return &MintParamsOK{}
}

/* MintParamsOK describes a response with status code 200, with default header values.

A successful response.
*/
type MintParamsOK struct {
	Payload *MintParamsOKBody
}

func (o *MintParamsOK) Error() string {
	return fmt.Sprintf("[GET /cosmos/mint/v1beta1/params][%d] mintParamsOK  %+v", 200, o.Payload)
}
func (o *MintParamsOK) GetPayload() *MintParamsOKBody {
	return o.Payload
}

func (o *MintParamsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(MintParamsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMintParamsDefault creates a MintParamsDefault with default headers values
func NewMintParamsDefault(code int) *MintParamsDefault {
	return &MintParamsDefault{
		_statusCode: code,
	}
}

/* MintParamsDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type MintParamsDefault struct {
	_statusCode int

	Payload *MintParamsDefaultBody
}

// Code gets the status code for the mint params default response
func (o *MintParamsDefault) Code() int {
	return o._statusCode
}

func (o *MintParamsDefault) Error() string {
	return fmt.Sprintf("[GET /cosmos/mint/v1beta1/params][%d] MintParams default  %+v", o._statusCode, o.Payload)
}
func (o *MintParamsDefault) GetPayload() *MintParamsDefaultBody {
	return o.Payload
}

func (o *MintParamsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(MintParamsDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*MintParamsDefaultBody mint params default body
swagger:model MintParamsDefaultBody
*/
type MintParamsDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// details
	Details []*MintParamsDefaultBodyDetailsItems0 `json:"details"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this mint params default body
func (o *MintParamsDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *MintParamsDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("MintParams default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("MintParams default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this mint params default body based on the context it is used
func (o *MintParamsDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *MintParamsDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("MintParams default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("MintParams default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *MintParamsDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *MintParamsDefaultBody) UnmarshalBinary(b []byte) error {
	var res MintParamsDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*MintParamsDefaultBodyDetailsItems0 mint params default body details items0
swagger:model MintParamsDefaultBodyDetailsItems0
*/
type MintParamsDefaultBodyDetailsItems0 struct {

	// type url
	TypeURL string `json:"type_url,omitempty"`

	// value
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this mint params default body details items0
func (o *MintParamsDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this mint params default body details items0 based on context it is used
func (o *MintParamsDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *MintParamsDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *MintParamsDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res MintParamsDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*MintParamsOKBody QueryParamsResponse is the response type for the Query/Params RPC method.
swagger:model MintParamsOKBody
*/
type MintParamsOKBody struct {

	// params
	Params *MintParamsOKBodyParams `json:"params,omitempty"`
}

// Validate validates this mint params o k body
func (o *MintParamsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateParams(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *MintParamsOKBody) validateParams(formats strfmt.Registry) error {
	if swag.IsZero(o.Params) { // not required
		return nil
	}

	if o.Params != nil {
		if err := o.Params.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mintParamsOK" + "." + "params")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mintParamsOK" + "." + "params")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this mint params o k body based on the context it is used
func (o *MintParamsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateParams(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *MintParamsOKBody) contextValidateParams(ctx context.Context, formats strfmt.Registry) error {

	if o.Params != nil {
		if err := o.Params.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mintParamsOK" + "." + "params")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mintParamsOK" + "." + "params")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *MintParamsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *MintParamsOKBody) UnmarshalBinary(b []byte) error {
	var res MintParamsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*MintParamsOKBodyParams params defines the parameters of the module.
swagger:model MintParamsOKBodyParams
*/
type MintParamsOKBodyParams struct {

	// expected blocks per year
	BlocksPerYear string `json:"blocks_per_year,omitempty"`

	// goal of percent bonded atoms
	GoalBonded string `json:"goal_bonded,omitempty"`

	// maximum inflation rate
	InflationMax string `json:"inflation_max,omitempty"`

	// minimum inflation rate
	InflationMin string `json:"inflation_min,omitempty"`

	// maximum annual change in inflation rate
	InflationRateChange string `json:"inflation_rate_change,omitempty"`

	// type of coin to mint
	MintDenom string `json:"mint_denom,omitempty"`
}

// Validate validates this mint params o k body params
func (o *MintParamsOKBodyParams) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this mint params o k body params based on context it is used
func (o *MintParamsOKBodyParams) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *MintParamsOKBodyParams) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *MintParamsOKBodyParams) UnmarshalBinary(b []byte) error {
	var res MintParamsOKBodyParams
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

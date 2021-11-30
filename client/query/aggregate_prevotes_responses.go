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

// AggregatePrevotesReader is a Reader for the AggregatePrevotes structure.
type AggregatePrevotesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AggregatePrevotesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAggregatePrevotesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAggregatePrevotesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAggregatePrevotesOK creates a AggregatePrevotesOK with default headers values
func NewAggregatePrevotesOK() *AggregatePrevotesOK {
	return &AggregatePrevotesOK{}
}

/* AggregatePrevotesOK describes a response with status code 200, with default header values.

A successful response.
*/
type AggregatePrevotesOK struct {
	Payload *AggregatePrevotesOKBody
}

func (o *AggregatePrevotesOK) Error() string {
	return fmt.Sprintf("[GET /terra/oracle/v1beta1/validators/aggregate_prevotes][%d] aggregatePrevotesOK  %+v", 200, o.Payload)
}
func (o *AggregatePrevotesOK) GetPayload() *AggregatePrevotesOKBody {
	return o.Payload
}

func (o *AggregatePrevotesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AggregatePrevotesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAggregatePrevotesDefault creates a AggregatePrevotesDefault with default headers values
func NewAggregatePrevotesDefault(code int) *AggregatePrevotesDefault {
	return &AggregatePrevotesDefault{
		_statusCode: code,
	}
}

/* AggregatePrevotesDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type AggregatePrevotesDefault struct {
	_statusCode int

	Payload *AggregatePrevotesDefaultBody
}

// Code gets the status code for the aggregate prevotes default response
func (o *AggregatePrevotesDefault) Code() int {
	return o._statusCode
}

func (o *AggregatePrevotesDefault) Error() string {
	return fmt.Sprintf("[GET /terra/oracle/v1beta1/validators/aggregate_prevotes][%d] AggregatePrevotes default  %+v", o._statusCode, o.Payload)
}
func (o *AggregatePrevotesDefault) GetPayload() *AggregatePrevotesDefaultBody {
	return o.Payload
}

func (o *AggregatePrevotesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AggregatePrevotesDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AggregatePrevotesDefaultBody aggregate prevotes default body
swagger:model AggregatePrevotesDefaultBody
*/
type AggregatePrevotesDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// details
	Details []*AggregatePrevotesDefaultBodyDetailsItems0 `json:"details"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this aggregate prevotes default body
func (o *AggregatePrevotesDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AggregatePrevotesDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("AggregatePrevotes default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AggregatePrevotes default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this aggregate prevotes default body based on the context it is used
func (o *AggregatePrevotesDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AggregatePrevotesDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("AggregatePrevotes default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("AggregatePrevotes default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *AggregatePrevotesDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AggregatePrevotesDefaultBody) UnmarshalBinary(b []byte) error {
	var res AggregatePrevotesDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AggregatePrevotesDefaultBodyDetailsItems0 aggregate prevotes default body details items0
swagger:model AggregatePrevotesDefaultBodyDetailsItems0
*/
type AggregatePrevotesDefaultBodyDetailsItems0 struct {

	// type url
	TypeURL string `json:"type_url,omitempty"`

	// value
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this aggregate prevotes default body details items0
func (o *AggregatePrevotesDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this aggregate prevotes default body details items0 based on context it is used
func (o *AggregatePrevotesDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AggregatePrevotesDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AggregatePrevotesDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res AggregatePrevotesDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AggregatePrevotesOKBody QueryAggregatePrevotesResponse is response type for the
// Query/AggregatePrevotes RPC method.
swagger:model AggregatePrevotesOKBody
*/
type AggregatePrevotesOKBody struct {

	// aggregate_prevotes defines all oracle aggregate prevotes submitted in the current vote period
	AggregatePrevotes []*AggregatePrevotesOKBodyAggregatePrevotesItems0 `json:"aggregate_prevotes"`
}

// Validate validates this aggregate prevotes o k body
func (o *AggregatePrevotesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAggregatePrevotes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AggregatePrevotesOKBody) validateAggregatePrevotes(formats strfmt.Registry) error {
	if swag.IsZero(o.AggregatePrevotes) { // not required
		return nil
	}

	for i := 0; i < len(o.AggregatePrevotes); i++ {
		if swag.IsZero(o.AggregatePrevotes[i]) { // not required
			continue
		}

		if o.AggregatePrevotes[i] != nil {
			if err := o.AggregatePrevotes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("aggregatePrevotesOK" + "." + "aggregate_prevotes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("aggregatePrevotesOK" + "." + "aggregate_prevotes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this aggregate prevotes o k body based on the context it is used
func (o *AggregatePrevotesOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAggregatePrevotes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AggregatePrevotesOKBody) contextValidateAggregatePrevotes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.AggregatePrevotes); i++ {

		if o.AggregatePrevotes[i] != nil {
			if err := o.AggregatePrevotes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("aggregatePrevotesOK" + "." + "aggregate_prevotes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("aggregatePrevotesOK" + "." + "aggregate_prevotes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *AggregatePrevotesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AggregatePrevotesOKBody) UnmarshalBinary(b []byte) error {
	var res AggregatePrevotesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AggregatePrevotesOKBodyAggregatePrevotesItems0 struct for aggregate prevoting on the ExchangeRateVote.
// The purpose of aggregate prevote is to hide vote exchange rates with hash
// which is formatted as hex string in SHA256("{salt}:{exchange rate}{denom},...,{exchange rate}{denom}:{voter}")
swagger:model AggregatePrevotesOKBodyAggregatePrevotesItems0
*/
type AggregatePrevotesOKBodyAggregatePrevotesItems0 struct {

	// hash
	Hash string `json:"hash,omitempty"`

	// submit block
	SubmitBlock string `json:"submit_block,omitempty"`

	// voter
	Voter string `json:"voter,omitempty"`
}

// Validate validates this aggregate prevotes o k body aggregate prevotes items0
func (o *AggregatePrevotesOKBodyAggregatePrevotesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this aggregate prevotes o k body aggregate prevotes items0 based on context it is used
func (o *AggregatePrevotesOKBodyAggregatePrevotesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AggregatePrevotesOKBodyAggregatePrevotesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AggregatePrevotesOKBodyAggregatePrevotesItems0) UnmarshalBinary(b []byte) error {
	var res AggregatePrevotesOKBodyAggregatePrevotesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
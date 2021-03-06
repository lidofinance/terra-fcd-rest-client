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

// ValidatorCommissionReader is a Reader for the ValidatorCommission structure.
type ValidatorCommissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ValidatorCommissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewValidatorCommissionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewValidatorCommissionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewValidatorCommissionOK creates a ValidatorCommissionOK with default headers values
func NewValidatorCommissionOK() *ValidatorCommissionOK {
	return &ValidatorCommissionOK{}
}

/* ValidatorCommissionOK describes a response with status code 200, with default header values.

A successful response.
*/
type ValidatorCommissionOK struct {
	Payload *ValidatorCommissionOKBody
}

func (o *ValidatorCommissionOK) Error() string {
	return fmt.Sprintf("[GET /cosmos/distribution/v1beta1/validators/{validator_address}/commission][%d] validatorCommissionOK  %+v", 200, o.Payload)
}
func (o *ValidatorCommissionOK) GetPayload() *ValidatorCommissionOKBody {
	return o.Payload
}

func (o *ValidatorCommissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ValidatorCommissionOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewValidatorCommissionDefault creates a ValidatorCommissionDefault with default headers values
func NewValidatorCommissionDefault(code int) *ValidatorCommissionDefault {
	return &ValidatorCommissionDefault{
		_statusCode: code,
	}
}

/* ValidatorCommissionDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type ValidatorCommissionDefault struct {
	_statusCode int

	Payload *ValidatorCommissionDefaultBody
}

// Code gets the status code for the validator commission default response
func (o *ValidatorCommissionDefault) Code() int {
	return o._statusCode
}

func (o *ValidatorCommissionDefault) Error() string {
	return fmt.Sprintf("[GET /cosmos/distribution/v1beta1/validators/{validator_address}/commission][%d] ValidatorCommission default  %+v", o._statusCode, o.Payload)
}
func (o *ValidatorCommissionDefault) GetPayload() *ValidatorCommissionDefaultBody {
	return o.Payload
}

func (o *ValidatorCommissionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ValidatorCommissionDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ValidatorCommissionDefaultBody validator commission default body
swagger:model ValidatorCommissionDefaultBody
*/
type ValidatorCommissionDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// details
	Details []*ValidatorCommissionDefaultBodyDetailsItems0 `json:"details"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this validator commission default body
func (o *ValidatorCommissionDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ValidatorCommissionDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("ValidatorCommission default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ValidatorCommission default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this validator commission default body based on the context it is used
func (o *ValidatorCommissionDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ValidatorCommissionDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ValidatorCommission default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ValidatorCommission default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ValidatorCommissionDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ValidatorCommissionDefaultBody) UnmarshalBinary(b []byte) error {
	var res ValidatorCommissionDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ValidatorCommissionDefaultBodyDetailsItems0 validator commission default body details items0
swagger:model ValidatorCommissionDefaultBodyDetailsItems0
*/
type ValidatorCommissionDefaultBodyDetailsItems0 struct {

	// type url
	TypeURL string `json:"type_url,omitempty"`

	// value
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this validator commission default body details items0
func (o *ValidatorCommissionDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this validator commission default body details items0 based on context it is used
func (o *ValidatorCommissionDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ValidatorCommissionDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ValidatorCommissionDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res ValidatorCommissionDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ValidatorCommissionOKBody QueryValidatorCommissionResponse is the response type for the
// Query/ValidatorCommission RPC method
swagger:model ValidatorCommissionOKBody
*/
type ValidatorCommissionOKBody struct {

	// commission
	Commission *ValidatorCommissionOKBodyCommission `json:"commission,omitempty"`
}

// Validate validates this validator commission o k body
func (o *ValidatorCommissionOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCommission(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ValidatorCommissionOKBody) validateCommission(formats strfmt.Registry) error {
	if swag.IsZero(o.Commission) { // not required
		return nil
	}

	if o.Commission != nil {
		if err := o.Commission.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("validatorCommissionOK" + "." + "commission")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("validatorCommissionOK" + "." + "commission")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this validator commission o k body based on the context it is used
func (o *ValidatorCommissionOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateCommission(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ValidatorCommissionOKBody) contextValidateCommission(ctx context.Context, formats strfmt.Registry) error {

	if o.Commission != nil {
		if err := o.Commission.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("validatorCommissionOK" + "." + "commission")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("validatorCommissionOK" + "." + "commission")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ValidatorCommissionOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ValidatorCommissionOKBody) UnmarshalBinary(b []byte) error {
	var res ValidatorCommissionOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ValidatorCommissionOKBodyCommission commission defines the commision the validator received.
swagger:model ValidatorCommissionOKBodyCommission
*/
type ValidatorCommissionOKBodyCommission struct {

	// commission
	Commission []*ValidatorCommissionOKBodyCommissionCommissionItems0 `json:"commission"`
}

// Validate validates this validator commission o k body commission
func (o *ValidatorCommissionOKBodyCommission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCommission(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ValidatorCommissionOKBodyCommission) validateCommission(formats strfmt.Registry) error {
	if swag.IsZero(o.Commission) { // not required
		return nil
	}

	for i := 0; i < len(o.Commission); i++ {
		if swag.IsZero(o.Commission[i]) { // not required
			continue
		}

		if o.Commission[i] != nil {
			if err := o.Commission[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("validatorCommissionOK" + "." + "commission" + "." + "commission" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("validatorCommissionOK" + "." + "commission" + "." + "commission" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this validator commission o k body commission based on the context it is used
func (o *ValidatorCommissionOKBodyCommission) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateCommission(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ValidatorCommissionOKBodyCommission) contextValidateCommission(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Commission); i++ {

		if o.Commission[i] != nil {
			if err := o.Commission[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("validatorCommissionOK" + "." + "commission" + "." + "commission" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("validatorCommissionOK" + "." + "commission" + "." + "commission" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ValidatorCommissionOKBodyCommission) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ValidatorCommissionOKBodyCommission) UnmarshalBinary(b []byte) error {
	var res ValidatorCommissionOKBodyCommission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ValidatorCommissionOKBodyCommissionCommissionItems0 DecCoin defines a token with a denomination and a decimal amount.
//
// NOTE: The amount field is an Dec which implements the custom method
// signatures required by gogoproto.
swagger:model ValidatorCommissionOKBodyCommissionCommissionItems0
*/
type ValidatorCommissionOKBodyCommissionCommissionItems0 struct {

	// amount
	Amount string `json:"amount,omitempty"`

	// denom
	Denom string `json:"denom,omitempty"`
}

// Validate validates this validator commission o k body commission commission items0
func (o *ValidatorCommissionOKBodyCommissionCommissionItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this validator commission o k body commission commission items0 based on context it is used
func (o *ValidatorCommissionOKBodyCommissionCommissionItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ValidatorCommissionOKBodyCommissionCommissionItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ValidatorCommissionOKBodyCommissionCommissionItems0) UnmarshalBinary(b []byte) error {
	var res ValidatorCommissionOKBodyCommissionCommissionItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

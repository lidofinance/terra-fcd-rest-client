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
	"github.com/go-openapi/validate"
)

// SigningInfoReader is a Reader for the SigningInfo structure.
type SigningInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SigningInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSigningInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSigningInfoDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSigningInfoOK creates a SigningInfoOK with default headers values
func NewSigningInfoOK() *SigningInfoOK {
	return &SigningInfoOK{}
}

/* SigningInfoOK describes a response with status code 200, with default header values.

A successful response.
*/
type SigningInfoOK struct {
	Payload *SigningInfoOKBody
}

func (o *SigningInfoOK) Error() string {
	return fmt.Sprintf("[GET /cosmos/slashing/v1beta1/signing_infos/{cons_address}][%d] signingInfoOK  %+v", 200, o.Payload)
}
func (o *SigningInfoOK) GetPayload() *SigningInfoOKBody {
	return o.Payload
}

func (o *SigningInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SigningInfoOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSigningInfoDefault creates a SigningInfoDefault with default headers values
func NewSigningInfoDefault(code int) *SigningInfoDefault {
	return &SigningInfoDefault{
		_statusCode: code,
	}
}

/* SigningInfoDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type SigningInfoDefault struct {
	_statusCode int

	Payload *SigningInfoDefaultBody
}

// Code gets the status code for the signing info default response
func (o *SigningInfoDefault) Code() int {
	return o._statusCode
}

func (o *SigningInfoDefault) Error() string {
	return fmt.Sprintf("[GET /cosmos/slashing/v1beta1/signing_infos/{cons_address}][%d] SigningInfo default  %+v", o._statusCode, o.Payload)
}
func (o *SigningInfoDefault) GetPayload() *SigningInfoDefaultBody {
	return o.Payload
}

func (o *SigningInfoDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SigningInfoDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*SigningInfoDefaultBody signing info default body
swagger:model SigningInfoDefaultBody
*/
type SigningInfoDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// details
	Details []*SigningInfoDefaultBodyDetailsItems0 `json:"details"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this signing info default body
func (o *SigningInfoDefaultBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SigningInfoDefaultBody) validateDetails(formats strfmt.Registry) error {
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
					return ve.ValidateName("SigningInfo default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("SigningInfo default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this signing info default body based on the context it is used
func (o *SigningInfoDefaultBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SigningInfoDefaultBody) contextValidateDetails(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Details); i++ {

		if o.Details[i] != nil {
			if err := o.Details[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("SigningInfo default" + "." + "details" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("SigningInfo default" + "." + "details" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *SigningInfoDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SigningInfoDefaultBody) UnmarshalBinary(b []byte) error {
	var res SigningInfoDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*SigningInfoDefaultBodyDetailsItems0 signing info default body details items0
swagger:model SigningInfoDefaultBodyDetailsItems0
*/
type SigningInfoDefaultBodyDetailsItems0 struct {

	// type url
	TypeURL string `json:"type_url,omitempty"`

	// value
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this signing info default body details items0
func (o *SigningInfoDefaultBodyDetailsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this signing info default body details items0 based on context it is used
func (o *SigningInfoDefaultBodyDetailsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SigningInfoDefaultBodyDetailsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SigningInfoDefaultBodyDetailsItems0) UnmarshalBinary(b []byte) error {
	var res SigningInfoDefaultBodyDetailsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*SigningInfoOKBody QuerySigningInfoResponse is the response type for the Query/SigningInfo RPC
// method
swagger:model SigningInfoOKBody
*/
type SigningInfoOKBody struct {

	// val signing info
	ValSigningInfo *SigningInfoOKBodyValSigningInfo `json:"val_signing_info,omitempty"`
}

// Validate validates this signing info o k body
func (o *SigningInfoOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateValSigningInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SigningInfoOKBody) validateValSigningInfo(formats strfmt.Registry) error {
	if swag.IsZero(o.ValSigningInfo) { // not required
		return nil
	}

	if o.ValSigningInfo != nil {
		if err := o.ValSigningInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("signingInfoOK" + "." + "val_signing_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("signingInfoOK" + "." + "val_signing_info")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this signing info o k body based on the context it is used
func (o *SigningInfoOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateValSigningInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SigningInfoOKBody) contextValidateValSigningInfo(ctx context.Context, formats strfmt.Registry) error {

	if o.ValSigningInfo != nil {
		if err := o.ValSigningInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("signingInfoOK" + "." + "val_signing_info")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("signingInfoOK" + "." + "val_signing_info")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *SigningInfoOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SigningInfoOKBody) UnmarshalBinary(b []byte) error {
	var res SigningInfoOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*SigningInfoOKBodyValSigningInfo val_signing_info is the signing info of requested val cons address
//
// ValidatorSigningInfo defines a validator's signing info for monitoring their
// liveness activity.
swagger:model SigningInfoOKBodyValSigningInfo
*/
type SigningInfoOKBodyValSigningInfo struct {

	// address
	Address string `json:"address,omitempty"`

	// Index which is incremented each time the validator was a bonded
	// in a block and may have signed a precommit or not. This in conjunction with the
	// `SignedBlocksWindow` param determines the index in the `MissedBlocksBitArray`.
	IndexOffset string `json:"index_offset,omitempty"`

	// Timestamp until which the validator is jailed due to liveness downtime.
	// Format: date-time
	JailedUntil strfmt.DateTime `json:"jailed_until,omitempty"`

	// A counter kept to avoid unnecessary array reads.
	// Note that `Sum(MissedBlocksBitArray)` always equals `MissedBlocksCounter`.
	MissedBlocksCounter string `json:"missed_blocks_counter,omitempty"`

	// Height at which validator was first a candidate OR was unjailed
	StartHeight string `json:"start_height,omitempty"`

	// Whether or not a validator has been tombstoned (killed out of validator set). It is set
	// once the validator commits an equivocation or for any other configured misbehiavor.
	Tombstoned bool `json:"tombstoned,omitempty"`
}

// Validate validates this signing info o k body val signing info
func (o *SigningInfoOKBodyValSigningInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateJailedUntil(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SigningInfoOKBodyValSigningInfo) validateJailedUntil(formats strfmt.Registry) error {
	if swag.IsZero(o.JailedUntil) { // not required
		return nil
	}

	if err := validate.FormatOf("signingInfoOK"+"."+"val_signing_info"+"."+"jailed_until", "body", "date-time", o.JailedUntil.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this signing info o k body val signing info based on context it is used
func (o *SigningInfoOKBodyValSigningInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SigningInfoOKBodyValSigningInfo) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SigningInfoOKBodyValSigningInfo) UnmarshalBinary(b []byte) error {
	var res SigningInfoOKBodyValSigningInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

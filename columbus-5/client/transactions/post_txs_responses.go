// Code generated by go-swagger; DO NOT EDIT.

package transactions

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

// PostTxsReader is a Reader for the PostTxs structure.
type PostTxsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTxsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostTxsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewPostTxsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostTxsOK creates a PostTxsOK with default headers values
func NewPostTxsOK() *PostTxsOK {
	return &PostTxsOK{}
}

/* PostTxsOK describes a response with status code 200, with default header values.

Tx broadcasting result
*/
type PostTxsOK struct {
	Payload *PostTxsOKBody
}

func (o *PostTxsOK) Error() string {
	return fmt.Sprintf("[POST /txs][%d] postTxsOK  %+v", 200, o.Payload)
}
func (o *PostTxsOK) GetPayload() *PostTxsOKBody {
	return o.Payload
}

func (o *PostTxsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostTxsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTxsInternalServerError creates a PostTxsInternalServerError with default headers values
func NewPostTxsInternalServerError() *PostTxsInternalServerError {
	return &PostTxsInternalServerError{}
}

/* PostTxsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostTxsInternalServerError struct {
}

func (o *PostTxsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /txs][%d] postTxsInternalServerError ", 500)
}

func (o *PostTxsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PostTxsBody post txs body
swagger:model PostTxsBody
*/
type PostTxsBody struct {

	// bech32 encoded address
	// Example: terra1wg2mlrxdmnnkkykgqg4znky86nyrtc45q336yv
	FeeGranter string `json:"fee_granter,omitempty"`

	// mode
	// Example: block
	// Required: true
	Mode *string `json:"mode"`

	// sequences
	Sequences []string `json:"sequences"`

	// tx
	// Required: true
	Tx *PostTxsParamsBodyTx `json:"tx"`
}

// Validate validates this post txs body
func (o *PostTxsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTx(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostTxsBody) validateMode(formats strfmt.Registry) error {

	if err := validate.Required("txBroadcast"+"."+"mode", "body", o.Mode); err != nil {
		return err
	}

	return nil
}

func (o *PostTxsBody) validateTx(formats strfmt.Registry) error {

	if err := validate.Required("txBroadcast"+"."+"tx", "body", o.Tx); err != nil {
		return err
	}

	if o.Tx != nil {
		if err := o.Tx.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("txBroadcast" + "." + "tx")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("txBroadcast" + "." + "tx")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post txs body based on the context it is used
func (o *PostTxsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateTx(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostTxsBody) contextValidateTx(ctx context.Context, formats strfmt.Registry) error {

	if o.Tx != nil {
		if err := o.Tx.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("txBroadcast" + "." + "tx")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("txBroadcast" + "." + "tx")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostTxsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostTxsBody) UnmarshalBinary(b []byte) error {
	var res PostTxsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostTxsOKBody post txs o k body
swagger:model PostTxsOKBody
*/
type PostTxsOKBody struct {

	// check tx
	CheckTx *PostTxsOKBodyCheckTx `json:"check_tx,omitempty"`

	// deliver tx
	DeliverTx *PostTxsOKBodyDeliverTx `json:"deliver_tx,omitempty"`

	// hash
	// Example: EE5F3404034C524501629B56E0DDC38FAD651F04
	Hash string `json:"hash,omitempty"`

	// height
	Height int64 `json:"height,omitempty"`
}

// Validate validates this post txs o k body
func (o *PostTxsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCheckTx(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDeliverTx(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostTxsOKBody) validateCheckTx(formats strfmt.Registry) error {
	if swag.IsZero(o.CheckTx) { // not required
		return nil
	}

	if o.CheckTx != nil {
		if err := o.CheckTx.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postTxsOK" + "." + "check_tx")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postTxsOK" + "." + "check_tx")
			}
			return err
		}
	}

	return nil
}

func (o *PostTxsOKBody) validateDeliverTx(formats strfmt.Registry) error {
	if swag.IsZero(o.DeliverTx) { // not required
		return nil
	}

	if o.DeliverTx != nil {
		if err := o.DeliverTx.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postTxsOK" + "." + "deliver_tx")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postTxsOK" + "." + "deliver_tx")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post txs o k body based on the context it is used
func (o *PostTxsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateCheckTx(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateDeliverTx(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostTxsOKBody) contextValidateCheckTx(ctx context.Context, formats strfmt.Registry) error {

	if o.CheckTx != nil {
		if err := o.CheckTx.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postTxsOK" + "." + "check_tx")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postTxsOK" + "." + "check_tx")
			}
			return err
		}
	}

	return nil
}

func (o *PostTxsOKBody) contextValidateDeliverTx(ctx context.Context, formats strfmt.Registry) error {

	if o.DeliverTx != nil {
		if err := o.DeliverTx.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postTxsOK" + "." + "deliver_tx")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postTxsOK" + "." + "deliver_tx")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostTxsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostTxsOKBody) UnmarshalBinary(b []byte) error {
	var res PostTxsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostTxsOKBodyCheckTx post txs o k body check tx
// Example: {"code":0,"data":"data","gas_used":5000,"gas_wanted":10000,"info":"info","log":"log","tags":["",""]}
swagger:model PostTxsOKBodyCheckTx
*/
type PostTxsOKBodyCheckTx struct {

	// code
	Code int64 `json:"code,omitempty"`

	// data
	Data string `json:"data,omitempty"`

	// gas used
	GasUsed int64 `json:"gas_used,omitempty"`

	// gas wanted
	GasWanted int64 `json:"gas_wanted,omitempty"`

	// info
	Info string `json:"info,omitempty"`

	// log
	Log string `json:"log,omitempty"`

	// tags
	Tags []*PostTxsOKBodyCheckTxTagsItems0 `json:"tags"`
}

// Validate validates this post txs o k body check tx
func (o *PostTxsOKBodyCheckTx) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostTxsOKBodyCheckTx) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(o.Tags) { // not required
		return nil
	}

	for i := 0; i < len(o.Tags); i++ {
		if swag.IsZero(o.Tags[i]) { // not required
			continue
		}

		if o.Tags[i] != nil {
			if err := o.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("postTxsOK" + "." + "check_tx" + "." + "tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("postTxsOK" + "." + "check_tx" + "." + "tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this post txs o k body check tx based on the context it is used
func (o *PostTxsOKBodyCheckTx) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostTxsOKBodyCheckTx) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Tags); i++ {

		if o.Tags[i] != nil {
			if err := o.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("postTxsOK" + "." + "check_tx" + "." + "tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("postTxsOK" + "." + "check_tx" + "." + "tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostTxsOKBodyCheckTx) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostTxsOKBodyCheckTx) UnmarshalBinary(b []byte) error {
	var res PostTxsOKBodyCheckTx
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostTxsOKBodyCheckTxTagsItems0 post txs o k body check tx tags items0
swagger:model PostTxsOKBodyCheckTxTagsItems0
*/
type PostTxsOKBodyCheckTxTagsItems0 struct {

	// key
	Key string `json:"key,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this post txs o k body check tx tags items0
func (o *PostTxsOKBodyCheckTxTagsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post txs o k body check tx tags items0 based on context it is used
func (o *PostTxsOKBodyCheckTxTagsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostTxsOKBodyCheckTxTagsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostTxsOKBodyCheckTxTagsItems0) UnmarshalBinary(b []byte) error {
	var res PostTxsOKBodyCheckTxTagsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostTxsOKBodyDeliverTx post txs o k body deliver tx
// Example: {"code":5,"data":"data","gas_used":5000,"gas_wanted":10000,"info":"info","log":"log","tags":["",""]}
swagger:model PostTxsOKBodyDeliverTx
*/
type PostTxsOKBodyDeliverTx struct {

	// code
	Code int64 `json:"code,omitempty"`

	// data
	Data string `json:"data,omitempty"`

	// gas used
	GasUsed int64 `json:"gas_used,omitempty"`

	// gas wanted
	GasWanted int64 `json:"gas_wanted,omitempty"`

	// info
	Info string `json:"info,omitempty"`

	// log
	Log string `json:"log,omitempty"`

	// tags
	Tags []*PostTxsOKBodyDeliverTxTagsItems0 `json:"tags"`
}

// Validate validates this post txs o k body deliver tx
func (o *PostTxsOKBodyDeliverTx) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostTxsOKBodyDeliverTx) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(o.Tags) { // not required
		return nil
	}

	for i := 0; i < len(o.Tags); i++ {
		if swag.IsZero(o.Tags[i]) { // not required
			continue
		}

		if o.Tags[i] != nil {
			if err := o.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("postTxsOK" + "." + "deliver_tx" + "." + "tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("postTxsOK" + "." + "deliver_tx" + "." + "tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this post txs o k body deliver tx based on the context it is used
func (o *PostTxsOKBodyDeliverTx) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostTxsOKBodyDeliverTx) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Tags); i++ {

		if o.Tags[i] != nil {
			if err := o.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("postTxsOK" + "." + "deliver_tx" + "." + "tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("postTxsOK" + "." + "deliver_tx" + "." + "tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostTxsOKBodyDeliverTx) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostTxsOKBodyDeliverTx) UnmarshalBinary(b []byte) error {
	var res PostTxsOKBodyDeliverTx
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostTxsOKBodyDeliverTxTagsItems0 post txs o k body deliver tx tags items0
swagger:model PostTxsOKBodyDeliverTxTagsItems0
*/
type PostTxsOKBodyDeliverTxTagsItems0 struct {

	// key
	Key string `json:"key,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this post txs o k body deliver tx tags items0
func (o *PostTxsOKBodyDeliverTxTagsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post txs o k body deliver tx tags items0 based on context it is used
func (o *PostTxsOKBodyDeliverTxTagsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostTxsOKBodyDeliverTxTagsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostTxsOKBodyDeliverTxTagsItems0) UnmarshalBinary(b []byte) error {
	var res PostTxsOKBodyDeliverTxTagsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostTxsParamsBodyTx post txs params body tx
swagger:model PostTxsParamsBodyTx
*/
type PostTxsParamsBodyTx struct {

	// fee
	Fee *PostTxsParamsBodyTxFee `json:"fee,omitempty"`

	// memo
	Memo string `json:"memo,omitempty"`

	// msg
	Msg []string `json:"msg"`

	// signature
	Signature *PostTxsParamsBodyTxSignature `json:"signature,omitempty"`
}

// Validate validates this post txs params body tx
func (o *PostTxsParamsBodyTx) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFee(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSignature(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostTxsParamsBodyTx) validateFee(formats strfmt.Registry) error {
	if swag.IsZero(o.Fee) { // not required
		return nil
	}

	if o.Fee != nil {
		if err := o.Fee.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("txBroadcast" + "." + "tx" + "." + "fee")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("txBroadcast" + "." + "tx" + "." + "fee")
			}
			return err
		}
	}

	return nil
}

func (o *PostTxsParamsBodyTx) validateSignature(formats strfmt.Registry) error {
	if swag.IsZero(o.Signature) { // not required
		return nil
	}

	if o.Signature != nil {
		if err := o.Signature.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("txBroadcast" + "." + "tx" + "." + "signature")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("txBroadcast" + "." + "tx" + "." + "signature")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post txs params body tx based on the context it is used
func (o *PostTxsParamsBodyTx) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateFee(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSignature(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostTxsParamsBodyTx) contextValidateFee(ctx context.Context, formats strfmt.Registry) error {

	if o.Fee != nil {
		if err := o.Fee.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("txBroadcast" + "." + "tx" + "." + "fee")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("txBroadcast" + "." + "tx" + "." + "fee")
			}
			return err
		}
	}

	return nil
}

func (o *PostTxsParamsBodyTx) contextValidateSignature(ctx context.Context, formats strfmt.Registry) error {

	if o.Signature != nil {
		if err := o.Signature.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("txBroadcast" + "." + "tx" + "." + "signature")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("txBroadcast" + "." + "tx" + "." + "signature")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostTxsParamsBodyTx) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostTxsParamsBodyTx) UnmarshalBinary(b []byte) error {
	var res PostTxsParamsBodyTx
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostTxsParamsBodyTxFee post txs params body tx fee
swagger:model PostTxsParamsBodyTxFee
*/
type PostTxsParamsBodyTxFee struct {

	// amount
	Amount []*PostTxsParamsBodyTxFeeAmountItems0 `json:"amount"`

	// gas
	Gas string `json:"gas,omitempty"`
}

// Validate validates this post txs params body tx fee
func (o *PostTxsParamsBodyTxFee) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostTxsParamsBodyTxFee) validateAmount(formats strfmt.Registry) error {
	if swag.IsZero(o.Amount) { // not required
		return nil
	}

	for i := 0; i < len(o.Amount); i++ {
		if swag.IsZero(o.Amount[i]) { // not required
			continue
		}

		if o.Amount[i] != nil {
			if err := o.Amount[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("txBroadcast" + "." + "tx" + "." + "fee" + "." + "amount" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("txBroadcast" + "." + "tx" + "." + "fee" + "." + "amount" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this post txs params body tx fee based on the context it is used
func (o *PostTxsParamsBodyTxFee) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostTxsParamsBodyTxFee) contextValidateAmount(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Amount); i++ {

		if o.Amount[i] != nil {
			if err := o.Amount[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("txBroadcast" + "." + "tx" + "." + "fee" + "." + "amount" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("txBroadcast" + "." + "tx" + "." + "fee" + "." + "amount" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostTxsParamsBodyTxFee) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostTxsParamsBodyTxFee) UnmarshalBinary(b []byte) error {
	var res PostTxsParamsBodyTxFee
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostTxsParamsBodyTxFeeAmountItems0 post txs params body tx fee amount items0
swagger:model PostTxsParamsBodyTxFeeAmountItems0
*/
type PostTxsParamsBodyTxFeeAmountItems0 struct {

	// amount
	// Example: 50
	Amount string `json:"amount,omitempty"`

	// denom
	// Example: uluna
	Denom string `json:"denom,omitempty"`
}

// Validate validates this post txs params body tx fee amount items0
func (o *PostTxsParamsBodyTxFeeAmountItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post txs params body tx fee amount items0 based on context it is used
func (o *PostTxsParamsBodyTxFeeAmountItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostTxsParamsBodyTxFeeAmountItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostTxsParamsBodyTxFeeAmountItems0) UnmarshalBinary(b []byte) error {
	var res PostTxsParamsBodyTxFeeAmountItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostTxsParamsBodyTxSignature post txs params body tx signature
swagger:model PostTxsParamsBodyTxSignature
*/
type PostTxsParamsBodyTxSignature struct {

	// account number
	// Example: 0
	AccountNumber string `json:"account_number,omitempty"`

	// pub key
	PubKey *PostTxsParamsBodyTxSignaturePubKey `json:"pub_key,omitempty"`

	// sequence
	// Example: 0
	Sequence string `json:"sequence,omitempty"`

	// signature
	// Example: MEUCIQD02fsDPra8MtbRsyB1w7bqTM55Wu138zQbFcWx4+CFyAIge5WNPfKIuvzBZ69MyqHsqD8S1IwiEp+iUb6VSdtlpgY=
	Signature string `json:"signature,omitempty"`
}

// Validate validates this post txs params body tx signature
func (o *PostTxsParamsBodyTxSignature) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePubKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostTxsParamsBodyTxSignature) validatePubKey(formats strfmt.Registry) error {
	if swag.IsZero(o.PubKey) { // not required
		return nil
	}

	if o.PubKey != nil {
		if err := o.PubKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("txBroadcast" + "." + "tx" + "." + "signature" + "." + "pub_key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("txBroadcast" + "." + "tx" + "." + "signature" + "." + "pub_key")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post txs params body tx signature based on the context it is used
func (o *PostTxsParamsBodyTxSignature) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePubKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostTxsParamsBodyTxSignature) contextValidatePubKey(ctx context.Context, formats strfmt.Registry) error {

	if o.PubKey != nil {
		if err := o.PubKey.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("txBroadcast" + "." + "tx" + "." + "signature" + "." + "pub_key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("txBroadcast" + "." + "tx" + "." + "signature" + "." + "pub_key")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostTxsParamsBodyTxSignature) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostTxsParamsBodyTxSignature) UnmarshalBinary(b []byte) error {
	var res PostTxsParamsBodyTxSignature
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostTxsParamsBodyTxSignaturePubKey post txs params body tx signature pub key
swagger:model PostTxsParamsBodyTxSignaturePubKey
*/
type PostTxsParamsBodyTxSignaturePubKey struct {

	// type
	// Example: tendermint/PubKeySecp256k1
	Type string `json:"type,omitempty"`

	// value
	// Example: Avz04VhtKJh8ACCVzlI8aTosGy0ikFXKIVHQ3jKMrosH
	Value string `json:"value,omitempty"`
}

// Validate validates this post txs params body tx signature pub key
func (o *PostTxsParamsBodyTxSignaturePubKey) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post txs params body tx signature pub key based on context it is used
func (o *PostTxsParamsBodyTxSignaturePubKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostTxsParamsBodyTxSignaturePubKey) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostTxsParamsBodyTxSignaturePubKey) UnmarshalBinary(b []byte) error {
	var res PostTxsParamsBodyTxSignaturePubKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package wasm

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

// PostWasmCodesCodeIDMigrateReader is a Reader for the PostWasmCodesCodeIDMigrate structure.
type PostWasmCodesCodeIDMigrateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWasmCodesCodeIDMigrateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostWasmCodesCodeIDMigrateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostWasmCodesCodeIDMigrateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostWasmCodesCodeIDMigrateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostWasmCodesCodeIDMigrateOK creates a PostWasmCodesCodeIDMigrateOK with default headers values
func NewPostWasmCodesCodeIDMigrateOK() *PostWasmCodesCodeIDMigrateOK {
	return &PostWasmCodesCodeIDMigrateOK{}
}

/* PostWasmCodesCodeIDMigrateOK describes a response with status code 200, with default header values.

OK
*/
type PostWasmCodesCodeIDMigrateOK struct {
	Payload *PostWasmCodesCodeIDMigrateOKBody
}

func (o *PostWasmCodesCodeIDMigrateOK) Error() string {
	return fmt.Sprintf("[POST /wasm/codes/{codeID}/migrate][%d] postWasmCodesCodeIdMigrateOK  %+v", 200, o.Payload)
}
func (o *PostWasmCodesCodeIDMigrateOK) GetPayload() *PostWasmCodesCodeIDMigrateOKBody {
	return o.Payload
}

func (o *PostWasmCodesCodeIDMigrateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostWasmCodesCodeIDMigrateOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWasmCodesCodeIDMigrateBadRequest creates a PostWasmCodesCodeIDMigrateBadRequest with default headers values
func NewPostWasmCodesCodeIDMigrateBadRequest() *PostWasmCodesCodeIDMigrateBadRequest {
	return &PostWasmCodesCodeIDMigrateBadRequest{}
}

/* PostWasmCodesCodeIDMigrateBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PostWasmCodesCodeIDMigrateBadRequest struct {
}

func (o *PostWasmCodesCodeIDMigrateBadRequest) Error() string {
	return fmt.Sprintf("[POST /wasm/codes/{codeID}/migrate][%d] postWasmCodesCodeIdMigrateBadRequest ", 400)
}

func (o *PostWasmCodesCodeIDMigrateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostWasmCodesCodeIDMigrateInternalServerError creates a PostWasmCodesCodeIDMigrateInternalServerError with default headers values
func NewPostWasmCodesCodeIDMigrateInternalServerError() *PostWasmCodesCodeIDMigrateInternalServerError {
	return &PostWasmCodesCodeIDMigrateInternalServerError{}
}

/* PostWasmCodesCodeIDMigrateInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostWasmCodesCodeIDMigrateInternalServerError struct {
}

func (o *PostWasmCodesCodeIDMigrateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /wasm/codes/{codeID}/migrate][%d] postWasmCodesCodeIdMigrateInternalServerError ", 500)
}

func (o *PostWasmCodesCodeIDMigrateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PostWasmCodesCodeIDMigrateBody post wasm codes code ID migrate body
swagger:model PostWasmCodesCodeIDMigrateBody
*/
type PostWasmCodesCodeIDMigrateBody struct {

	// base req
	BaseReq *PostWasmCodesCodeIDMigrateParamsBodyBaseReq `json:"base_req,omitempty"`

	// wasm bytes
	// Example: Avz04VhtKJh8ACCVzlI8aTosGy0ikFXKIVHQ3jKMrosH
	WasmBytes string `json:"wasm_bytes,omitempty"`
}

// Validate validates this post wasm codes code ID migrate body
func (o *PostWasmCodesCodeIDMigrateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBaseReq(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostWasmCodesCodeIDMigrateBody) validateBaseReq(formats strfmt.Registry) error {
	if swag.IsZero(o.BaseReq) { // not required
		return nil
	}

	if o.BaseReq != nil {
		if err := o.BaseReq.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("migrate contract request body" + "." + "base_req")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("migrate contract request body" + "." + "base_req")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post wasm codes code ID migrate body based on the context it is used
func (o *PostWasmCodesCodeIDMigrateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateBaseReq(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostWasmCodesCodeIDMigrateBody) contextValidateBaseReq(ctx context.Context, formats strfmt.Registry) error {

	if o.BaseReq != nil {
		if err := o.BaseReq.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("migrate contract request body" + "." + "base_req")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("migrate contract request body" + "." + "base_req")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostWasmCodesCodeIDMigrateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostWasmCodesCodeIDMigrateBody) UnmarshalBinary(b []byte) error {
	var res PostWasmCodesCodeIDMigrateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostWasmCodesCodeIDMigrateOKBody post wasm codes code ID migrate o k body
swagger:model PostWasmCodesCodeIDMigrateOKBody
*/
type PostWasmCodesCodeIDMigrateOKBody struct {

	// fee
	Fee *PostWasmCodesCodeIDMigrateOKBodyFee `json:"fee,omitempty"`

	// memo
	Memo string `json:"memo,omitempty"`

	// msg
	Msg []string `json:"msg"`

	// signature
	Signature *PostWasmCodesCodeIDMigrateOKBodySignature `json:"signature,omitempty"`
}

// Validate validates this post wasm codes code ID migrate o k body
func (o *PostWasmCodesCodeIDMigrateOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PostWasmCodesCodeIDMigrateOKBody) validateFee(formats strfmt.Registry) error {
	if swag.IsZero(o.Fee) { // not required
		return nil
	}

	if o.Fee != nil {
		if err := o.Fee.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postWasmCodesCodeIdMigrateOK" + "." + "fee")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postWasmCodesCodeIdMigrateOK" + "." + "fee")
			}
			return err
		}
	}

	return nil
}

func (o *PostWasmCodesCodeIDMigrateOKBody) validateSignature(formats strfmt.Registry) error {
	if swag.IsZero(o.Signature) { // not required
		return nil
	}

	if o.Signature != nil {
		if err := o.Signature.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postWasmCodesCodeIdMigrateOK" + "." + "signature")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postWasmCodesCodeIdMigrateOK" + "." + "signature")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post wasm codes code ID migrate o k body based on the context it is used
func (o *PostWasmCodesCodeIDMigrateOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *PostWasmCodesCodeIDMigrateOKBody) contextValidateFee(ctx context.Context, formats strfmt.Registry) error {

	if o.Fee != nil {
		if err := o.Fee.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postWasmCodesCodeIdMigrateOK" + "." + "fee")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postWasmCodesCodeIdMigrateOK" + "." + "fee")
			}
			return err
		}
	}

	return nil
}

func (o *PostWasmCodesCodeIDMigrateOKBody) contextValidateSignature(ctx context.Context, formats strfmt.Registry) error {

	if o.Signature != nil {
		if err := o.Signature.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postWasmCodesCodeIdMigrateOK" + "." + "signature")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postWasmCodesCodeIdMigrateOK" + "." + "signature")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostWasmCodesCodeIDMigrateOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostWasmCodesCodeIDMigrateOKBody) UnmarshalBinary(b []byte) error {
	var res PostWasmCodesCodeIDMigrateOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostWasmCodesCodeIDMigrateOKBodyFee post wasm codes code ID migrate o k body fee
swagger:model PostWasmCodesCodeIDMigrateOKBodyFee
*/
type PostWasmCodesCodeIDMigrateOKBodyFee struct {

	// amount
	Amount []*PostWasmCodesCodeIDMigrateOKBodyFeeAmountItems0 `json:"amount"`

	// gas
	Gas string `json:"gas,omitempty"`
}

// Validate validates this post wasm codes code ID migrate o k body fee
func (o *PostWasmCodesCodeIDMigrateOKBodyFee) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostWasmCodesCodeIDMigrateOKBodyFee) validateAmount(formats strfmt.Registry) error {
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
					return ve.ValidateName("postWasmCodesCodeIdMigrateOK" + "." + "fee" + "." + "amount" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("postWasmCodesCodeIdMigrateOK" + "." + "fee" + "." + "amount" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this post wasm codes code ID migrate o k body fee based on the context it is used
func (o *PostWasmCodesCodeIDMigrateOKBodyFee) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostWasmCodesCodeIDMigrateOKBodyFee) contextValidateAmount(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Amount); i++ {

		if o.Amount[i] != nil {
			if err := o.Amount[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("postWasmCodesCodeIdMigrateOK" + "." + "fee" + "." + "amount" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("postWasmCodesCodeIdMigrateOK" + "." + "fee" + "." + "amount" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostWasmCodesCodeIDMigrateOKBodyFee) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostWasmCodesCodeIDMigrateOKBodyFee) UnmarshalBinary(b []byte) error {
	var res PostWasmCodesCodeIDMigrateOKBodyFee
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostWasmCodesCodeIDMigrateOKBodyFeeAmountItems0 post wasm codes code ID migrate o k body fee amount items0
swagger:model PostWasmCodesCodeIDMigrateOKBodyFeeAmountItems0
*/
type PostWasmCodesCodeIDMigrateOKBodyFeeAmountItems0 struct {

	// amount
	// Example: 50
	Amount string `json:"amount,omitempty"`

	// denom
	// Example: uluna
	Denom string `json:"denom,omitempty"`
}

// Validate validates this post wasm codes code ID migrate o k body fee amount items0
func (o *PostWasmCodesCodeIDMigrateOKBodyFeeAmountItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post wasm codes code ID migrate o k body fee amount items0 based on context it is used
func (o *PostWasmCodesCodeIDMigrateOKBodyFeeAmountItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostWasmCodesCodeIDMigrateOKBodyFeeAmountItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostWasmCodesCodeIDMigrateOKBodyFeeAmountItems0) UnmarshalBinary(b []byte) error {
	var res PostWasmCodesCodeIDMigrateOKBodyFeeAmountItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostWasmCodesCodeIDMigrateOKBodySignature post wasm codes code ID migrate o k body signature
swagger:model PostWasmCodesCodeIDMigrateOKBodySignature
*/
type PostWasmCodesCodeIDMigrateOKBodySignature struct {

	// account number
	// Example: 0
	AccountNumber string `json:"account_number,omitempty"`

	// pub key
	PubKey *PostWasmCodesCodeIDMigrateOKBodySignaturePubKey `json:"pub_key,omitempty"`

	// sequence
	// Example: 0
	Sequence string `json:"sequence,omitempty"`

	// signature
	// Example: MEUCIQD02fsDPra8MtbRsyB1w7bqTM55Wu138zQbFcWx4+CFyAIge5WNPfKIuvzBZ69MyqHsqD8S1IwiEp+iUb6VSdtlpgY=
	Signature string `json:"signature,omitempty"`
}

// Validate validates this post wasm codes code ID migrate o k body signature
func (o *PostWasmCodesCodeIDMigrateOKBodySignature) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePubKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostWasmCodesCodeIDMigrateOKBodySignature) validatePubKey(formats strfmt.Registry) error {
	if swag.IsZero(o.PubKey) { // not required
		return nil
	}

	if o.PubKey != nil {
		if err := o.PubKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postWasmCodesCodeIdMigrateOK" + "." + "signature" + "." + "pub_key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postWasmCodesCodeIdMigrateOK" + "." + "signature" + "." + "pub_key")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post wasm codes code ID migrate o k body signature based on the context it is used
func (o *PostWasmCodesCodeIDMigrateOKBodySignature) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePubKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostWasmCodesCodeIDMigrateOKBodySignature) contextValidatePubKey(ctx context.Context, formats strfmt.Registry) error {

	if o.PubKey != nil {
		if err := o.PubKey.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postWasmCodesCodeIdMigrateOK" + "." + "signature" + "." + "pub_key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postWasmCodesCodeIdMigrateOK" + "." + "signature" + "." + "pub_key")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostWasmCodesCodeIDMigrateOKBodySignature) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostWasmCodesCodeIDMigrateOKBodySignature) UnmarshalBinary(b []byte) error {
	var res PostWasmCodesCodeIDMigrateOKBodySignature
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostWasmCodesCodeIDMigrateOKBodySignaturePubKey post wasm codes code ID migrate o k body signature pub key
swagger:model PostWasmCodesCodeIDMigrateOKBodySignaturePubKey
*/
type PostWasmCodesCodeIDMigrateOKBodySignaturePubKey struct {

	// type
	// Example: tendermint/PubKeySecp256k1
	Type string `json:"type,omitempty"`

	// value
	// Example: Avz04VhtKJh8ACCVzlI8aTosGy0ikFXKIVHQ3jKMrosH
	Value string `json:"value,omitempty"`
}

// Validate validates this post wasm codes code ID migrate o k body signature pub key
func (o *PostWasmCodesCodeIDMigrateOKBodySignaturePubKey) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post wasm codes code ID migrate o k body signature pub key based on context it is used
func (o *PostWasmCodesCodeIDMigrateOKBodySignaturePubKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostWasmCodesCodeIDMigrateOKBodySignaturePubKey) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostWasmCodesCodeIDMigrateOKBodySignaturePubKey) UnmarshalBinary(b []byte) error {
	var res PostWasmCodesCodeIDMigrateOKBodySignaturePubKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostWasmCodesCodeIDMigrateParamsBodyBaseReq post wasm codes code ID migrate params body base req
swagger:model PostWasmCodesCodeIDMigrateParamsBodyBaseReq
*/
type PostWasmCodesCodeIDMigrateParamsBodyBaseReq struct {

	// account number
	// Example: 0
	AccountNumber string `json:"account_number,omitempty"`

	// chain id
	// Example: Columbus-5
	ChainID string `json:"chain_id,omitempty"`

	// fees
	Fees []*PostWasmCodesCodeIDMigrateParamsBodyBaseReqFeesItems0 `json:"fees"`

	// Sender address or Keybase name to generate a transaction
	// Example: terra1wg2mlrxdmnnkkykgqg4znky86nyrtc45q336yv
	From string `json:"from,omitempty"`

	// gas
	// Example: 200000
	Gas string `json:"gas,omitempty"`

	// gas adjustment
	// Example: 1.2
	GasAdjustment string `json:"gas_adjustment,omitempty"`

	// memo
	// Example: Sent via Terra Station 🚀
	Memo string `json:"memo,omitempty"`

	// sequence
	// Example: 1
	Sequence string `json:"sequence,omitempty"`

	// Estimate gas for a transaction (cannot be used in conjunction with generate_only)
	// Example: false
	Simulate bool `json:"simulate,omitempty"`
}

// Validate validates this post wasm codes code ID migrate params body base req
func (o *PostWasmCodesCodeIDMigrateParamsBodyBaseReq) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFees(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostWasmCodesCodeIDMigrateParamsBodyBaseReq) validateFees(formats strfmt.Registry) error {
	if swag.IsZero(o.Fees) { // not required
		return nil
	}

	for i := 0; i < len(o.Fees); i++ {
		if swag.IsZero(o.Fees[i]) { // not required
			continue
		}

		if o.Fees[i] != nil {
			if err := o.Fees[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("migrate contract request body" + "." + "base_req" + "." + "fees" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("migrate contract request body" + "." + "base_req" + "." + "fees" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this post wasm codes code ID migrate params body base req based on the context it is used
func (o *PostWasmCodesCodeIDMigrateParamsBodyBaseReq) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateFees(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostWasmCodesCodeIDMigrateParamsBodyBaseReq) contextValidateFees(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Fees); i++ {

		if o.Fees[i] != nil {
			if err := o.Fees[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("migrate contract request body" + "." + "base_req" + "." + "fees" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("migrate contract request body" + "." + "base_req" + "." + "fees" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostWasmCodesCodeIDMigrateParamsBodyBaseReq) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostWasmCodesCodeIDMigrateParamsBodyBaseReq) UnmarshalBinary(b []byte) error {
	var res PostWasmCodesCodeIDMigrateParamsBodyBaseReq
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostWasmCodesCodeIDMigrateParamsBodyBaseReqFeesItems0 post wasm codes code ID migrate params body base req fees items0
swagger:model PostWasmCodesCodeIDMigrateParamsBodyBaseReqFeesItems0
*/
type PostWasmCodesCodeIDMigrateParamsBodyBaseReqFeesItems0 struct {

	// amount
	// Example: 50
	Amount string `json:"amount,omitempty"`

	// denom
	// Example: uluna
	Denom string `json:"denom,omitempty"`
}

// Validate validates this post wasm codes code ID migrate params body base req fees items0
func (o *PostWasmCodesCodeIDMigrateParamsBodyBaseReqFeesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post wasm codes code ID migrate params body base req fees items0 based on context it is used
func (o *PostWasmCodesCodeIDMigrateParamsBodyBaseReqFeesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostWasmCodesCodeIDMigrateParamsBodyBaseReqFeesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostWasmCodesCodeIDMigrateParamsBodyBaseReqFeesItems0) UnmarshalBinary(b []byte) error {
	var res PostWasmCodesCodeIDMigrateParamsBodyBaseReqFeesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

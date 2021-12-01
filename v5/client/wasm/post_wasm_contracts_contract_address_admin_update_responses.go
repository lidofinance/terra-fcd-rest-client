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

// PostWasmContractsContractAddressAdminUpdateReader is a Reader for the PostWasmContractsContractAddressAdminUpdate structure.
type PostWasmContractsContractAddressAdminUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostWasmContractsContractAddressAdminUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostWasmContractsContractAddressAdminUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostWasmContractsContractAddressAdminUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostWasmContractsContractAddressAdminUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostWasmContractsContractAddressAdminUpdateOK creates a PostWasmContractsContractAddressAdminUpdateOK with default headers values
func NewPostWasmContractsContractAddressAdminUpdateOK() *PostWasmContractsContractAddressAdminUpdateOK {
	return &PostWasmContractsContractAddressAdminUpdateOK{}
}

/* PostWasmContractsContractAddressAdminUpdateOK describes a response with status code 200, with default header values.

OK
*/
type PostWasmContractsContractAddressAdminUpdateOK struct {
	Payload *PostWasmContractsContractAddressAdminUpdateOKBody
}

func (o *PostWasmContractsContractAddressAdminUpdateOK) Error() string {
	return fmt.Sprintf("[POST /wasm/contracts/{contractAddress}/admin/update][%d] postWasmContractsContractAddressAdminUpdateOK  %+v", 200, o.Payload)
}
func (o *PostWasmContractsContractAddressAdminUpdateOK) GetPayload() *PostWasmContractsContractAddressAdminUpdateOKBody {
	return o.Payload
}

func (o *PostWasmContractsContractAddressAdminUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostWasmContractsContractAddressAdminUpdateOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostWasmContractsContractAddressAdminUpdateBadRequest creates a PostWasmContractsContractAddressAdminUpdateBadRequest with default headers values
func NewPostWasmContractsContractAddressAdminUpdateBadRequest() *PostWasmContractsContractAddressAdminUpdateBadRequest {
	return &PostWasmContractsContractAddressAdminUpdateBadRequest{}
}

/* PostWasmContractsContractAddressAdminUpdateBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PostWasmContractsContractAddressAdminUpdateBadRequest struct {
}

func (o *PostWasmContractsContractAddressAdminUpdateBadRequest) Error() string {
	return fmt.Sprintf("[POST /wasm/contracts/{contractAddress}/admin/update][%d] postWasmContractsContractAddressAdminUpdateBadRequest ", 400)
}

func (o *PostWasmContractsContractAddressAdminUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostWasmContractsContractAddressAdminUpdateInternalServerError creates a PostWasmContractsContractAddressAdminUpdateInternalServerError with default headers values
func NewPostWasmContractsContractAddressAdminUpdateInternalServerError() *PostWasmContractsContractAddressAdminUpdateInternalServerError {
	return &PostWasmContractsContractAddressAdminUpdateInternalServerError{}
}

/* PostWasmContractsContractAddressAdminUpdateInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostWasmContractsContractAddressAdminUpdateInternalServerError struct {
}

func (o *PostWasmContractsContractAddressAdminUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /wasm/contracts/{contractAddress}/admin/update][%d] postWasmContractsContractAddressAdminUpdateInternalServerError ", 500)
}

func (o *PostWasmContractsContractAddressAdminUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PostWasmContractsContractAddressAdminUpdateBody post wasm contracts contract address admin update body
swagger:model PostWasmContractsContractAddressAdminUpdateBody
*/
type PostWasmContractsContractAddressAdminUpdateBody struct {

	// base req
	BaseReq *PostWasmContractsContractAddressAdminUpdateParamsBodyBaseReq `json:"base_req,omitempty"`

	// bech32 encoded address
	// Example: terra1wg2mlrxdmnnkkykgqg4znky86nyrtc45q336yv
	NewAdmin string `json:"new_admin,omitempty"`
}

// Validate validates this post wasm contracts contract address admin update body
func (o *PostWasmContractsContractAddressAdminUpdateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBaseReq(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostWasmContractsContractAddressAdminUpdateBody) validateBaseReq(formats strfmt.Registry) error {
	if swag.IsZero(o.BaseReq) { // not required
		return nil
	}

	if o.BaseReq != nil {
		if err := o.BaseReq.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("update contract admin request body" + "." + "base_req")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("update contract admin request body" + "." + "base_req")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post wasm contracts contract address admin update body based on the context it is used
func (o *PostWasmContractsContractAddressAdminUpdateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateBaseReq(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostWasmContractsContractAddressAdminUpdateBody) contextValidateBaseReq(ctx context.Context, formats strfmt.Registry) error {

	if o.BaseReq != nil {
		if err := o.BaseReq.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("update contract admin request body" + "." + "base_req")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("update contract admin request body" + "." + "base_req")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostWasmContractsContractAddressAdminUpdateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostWasmContractsContractAddressAdminUpdateBody) UnmarshalBinary(b []byte) error {
	var res PostWasmContractsContractAddressAdminUpdateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostWasmContractsContractAddressAdminUpdateOKBody post wasm contracts contract address admin update o k body
swagger:model PostWasmContractsContractAddressAdminUpdateOKBody
*/
type PostWasmContractsContractAddressAdminUpdateOKBody struct {

	// fee
	Fee *PostWasmContractsContractAddressAdminUpdateOKBodyFee `json:"fee,omitempty"`

	// memo
	Memo string `json:"memo,omitempty"`

	// msg
	Msg []string `json:"msg"`

	// signature
	Signature *PostWasmContractsContractAddressAdminUpdateOKBodySignature `json:"signature,omitempty"`
}

// Validate validates this post wasm contracts contract address admin update o k body
func (o *PostWasmContractsContractAddressAdminUpdateOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PostWasmContractsContractAddressAdminUpdateOKBody) validateFee(formats strfmt.Registry) error {
	if swag.IsZero(o.Fee) { // not required
		return nil
	}

	if o.Fee != nil {
		if err := o.Fee.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postWasmContractsContractAddressAdminUpdateOK" + "." + "fee")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postWasmContractsContractAddressAdminUpdateOK" + "." + "fee")
			}
			return err
		}
	}

	return nil
}

func (o *PostWasmContractsContractAddressAdminUpdateOKBody) validateSignature(formats strfmt.Registry) error {
	if swag.IsZero(o.Signature) { // not required
		return nil
	}

	if o.Signature != nil {
		if err := o.Signature.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postWasmContractsContractAddressAdminUpdateOK" + "." + "signature")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postWasmContractsContractAddressAdminUpdateOK" + "." + "signature")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post wasm contracts contract address admin update o k body based on the context it is used
func (o *PostWasmContractsContractAddressAdminUpdateOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *PostWasmContractsContractAddressAdminUpdateOKBody) contextValidateFee(ctx context.Context, formats strfmt.Registry) error {

	if o.Fee != nil {
		if err := o.Fee.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postWasmContractsContractAddressAdminUpdateOK" + "." + "fee")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postWasmContractsContractAddressAdminUpdateOK" + "." + "fee")
			}
			return err
		}
	}

	return nil
}

func (o *PostWasmContractsContractAddressAdminUpdateOKBody) contextValidateSignature(ctx context.Context, formats strfmt.Registry) error {

	if o.Signature != nil {
		if err := o.Signature.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postWasmContractsContractAddressAdminUpdateOK" + "." + "signature")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postWasmContractsContractAddressAdminUpdateOK" + "." + "signature")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostWasmContractsContractAddressAdminUpdateOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostWasmContractsContractAddressAdminUpdateOKBody) UnmarshalBinary(b []byte) error {
	var res PostWasmContractsContractAddressAdminUpdateOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostWasmContractsContractAddressAdminUpdateOKBodyFee post wasm contracts contract address admin update o k body fee
swagger:model PostWasmContractsContractAddressAdminUpdateOKBodyFee
*/
type PostWasmContractsContractAddressAdminUpdateOKBodyFee struct {

	// amount
	Amount []*PostWasmContractsContractAddressAdminUpdateOKBodyFeeAmountItems0 `json:"amount"`

	// gas
	Gas string `json:"gas,omitempty"`
}

// Validate validates this post wasm contracts contract address admin update o k body fee
func (o *PostWasmContractsContractAddressAdminUpdateOKBodyFee) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostWasmContractsContractAddressAdminUpdateOKBodyFee) validateAmount(formats strfmt.Registry) error {
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
					return ve.ValidateName("postWasmContractsContractAddressAdminUpdateOK" + "." + "fee" + "." + "amount" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("postWasmContractsContractAddressAdminUpdateOK" + "." + "fee" + "." + "amount" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this post wasm contracts contract address admin update o k body fee based on the context it is used
func (o *PostWasmContractsContractAddressAdminUpdateOKBodyFee) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostWasmContractsContractAddressAdminUpdateOKBodyFee) contextValidateAmount(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Amount); i++ {

		if o.Amount[i] != nil {
			if err := o.Amount[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("postWasmContractsContractAddressAdminUpdateOK" + "." + "fee" + "." + "amount" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("postWasmContractsContractAddressAdminUpdateOK" + "." + "fee" + "." + "amount" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostWasmContractsContractAddressAdminUpdateOKBodyFee) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostWasmContractsContractAddressAdminUpdateOKBodyFee) UnmarshalBinary(b []byte) error {
	var res PostWasmContractsContractAddressAdminUpdateOKBodyFee
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostWasmContractsContractAddressAdminUpdateOKBodyFeeAmountItems0 post wasm contracts contract address admin update o k body fee amount items0
swagger:model PostWasmContractsContractAddressAdminUpdateOKBodyFeeAmountItems0
*/
type PostWasmContractsContractAddressAdminUpdateOKBodyFeeAmountItems0 struct {

	// amount
	// Example: 50
	Amount string `json:"amount,omitempty"`

	// denom
	// Example: uluna
	Denom string `json:"denom,omitempty"`
}

// Validate validates this post wasm contracts contract address admin update o k body fee amount items0
func (o *PostWasmContractsContractAddressAdminUpdateOKBodyFeeAmountItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post wasm contracts contract address admin update o k body fee amount items0 based on context it is used
func (o *PostWasmContractsContractAddressAdminUpdateOKBodyFeeAmountItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostWasmContractsContractAddressAdminUpdateOKBodyFeeAmountItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostWasmContractsContractAddressAdminUpdateOKBodyFeeAmountItems0) UnmarshalBinary(b []byte) error {
	var res PostWasmContractsContractAddressAdminUpdateOKBodyFeeAmountItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostWasmContractsContractAddressAdminUpdateOKBodySignature post wasm contracts contract address admin update o k body signature
swagger:model PostWasmContractsContractAddressAdminUpdateOKBodySignature
*/
type PostWasmContractsContractAddressAdminUpdateOKBodySignature struct {

	// account number
	// Example: 0
	AccountNumber string `json:"account_number,omitempty"`

	// pub key
	PubKey *PostWasmContractsContractAddressAdminUpdateOKBodySignaturePubKey `json:"pub_key,omitempty"`

	// sequence
	// Example: 0
	Sequence string `json:"sequence,omitempty"`

	// signature
	// Example: MEUCIQD02fsDPra8MtbRsyB1w7bqTM55Wu138zQbFcWx4+CFyAIge5WNPfKIuvzBZ69MyqHsqD8S1IwiEp+iUb6VSdtlpgY=
	Signature string `json:"signature,omitempty"`
}

// Validate validates this post wasm contracts contract address admin update o k body signature
func (o *PostWasmContractsContractAddressAdminUpdateOKBodySignature) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePubKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostWasmContractsContractAddressAdminUpdateOKBodySignature) validatePubKey(formats strfmt.Registry) error {
	if swag.IsZero(o.PubKey) { // not required
		return nil
	}

	if o.PubKey != nil {
		if err := o.PubKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postWasmContractsContractAddressAdminUpdateOK" + "." + "signature" + "." + "pub_key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postWasmContractsContractAddressAdminUpdateOK" + "." + "signature" + "." + "pub_key")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post wasm contracts contract address admin update o k body signature based on the context it is used
func (o *PostWasmContractsContractAddressAdminUpdateOKBodySignature) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePubKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostWasmContractsContractAddressAdminUpdateOKBodySignature) contextValidatePubKey(ctx context.Context, formats strfmt.Registry) error {

	if o.PubKey != nil {
		if err := o.PubKey.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postWasmContractsContractAddressAdminUpdateOK" + "." + "signature" + "." + "pub_key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postWasmContractsContractAddressAdminUpdateOK" + "." + "signature" + "." + "pub_key")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostWasmContractsContractAddressAdminUpdateOKBodySignature) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostWasmContractsContractAddressAdminUpdateOKBodySignature) UnmarshalBinary(b []byte) error {
	var res PostWasmContractsContractAddressAdminUpdateOKBodySignature
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostWasmContractsContractAddressAdminUpdateOKBodySignaturePubKey post wasm contracts contract address admin update o k body signature pub key
swagger:model PostWasmContractsContractAddressAdminUpdateOKBodySignaturePubKey
*/
type PostWasmContractsContractAddressAdminUpdateOKBodySignaturePubKey struct {

	// type
	// Example: tendermint/PubKeySecp256k1
	Type string `json:"type,omitempty"`

	// value
	// Example: Avz04VhtKJh8ACCVzlI8aTosGy0ikFXKIVHQ3jKMrosH
	Value string `json:"value,omitempty"`
}

// Validate validates this post wasm contracts contract address admin update o k body signature pub key
func (o *PostWasmContractsContractAddressAdminUpdateOKBodySignaturePubKey) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post wasm contracts contract address admin update o k body signature pub key based on context it is used
func (o *PostWasmContractsContractAddressAdminUpdateOKBodySignaturePubKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostWasmContractsContractAddressAdminUpdateOKBodySignaturePubKey) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostWasmContractsContractAddressAdminUpdateOKBodySignaturePubKey) UnmarshalBinary(b []byte) error {
	var res PostWasmContractsContractAddressAdminUpdateOKBodySignaturePubKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostWasmContractsContractAddressAdminUpdateParamsBodyBaseReq post wasm contracts contract address admin update params body base req
swagger:model PostWasmContractsContractAddressAdminUpdateParamsBodyBaseReq
*/
type PostWasmContractsContractAddressAdminUpdateParamsBodyBaseReq struct {

	// account number
	// Example: 0
	AccountNumber string `json:"account_number,omitempty"`

	// chain id
	// Example: Columbus-5
	ChainID string `json:"chain_id,omitempty"`

	// fees
	Fees []*PostWasmContractsContractAddressAdminUpdateParamsBodyBaseReqFeesItems0 `json:"fees"`

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

// Validate validates this post wasm contracts contract address admin update params body base req
func (o *PostWasmContractsContractAddressAdminUpdateParamsBodyBaseReq) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFees(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostWasmContractsContractAddressAdminUpdateParamsBodyBaseReq) validateFees(formats strfmt.Registry) error {
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
					return ve.ValidateName("update contract admin request body" + "." + "base_req" + "." + "fees" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("update contract admin request body" + "." + "base_req" + "." + "fees" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this post wasm contracts contract address admin update params body base req based on the context it is used
func (o *PostWasmContractsContractAddressAdminUpdateParamsBodyBaseReq) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateFees(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostWasmContractsContractAddressAdminUpdateParamsBodyBaseReq) contextValidateFees(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Fees); i++ {

		if o.Fees[i] != nil {
			if err := o.Fees[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("update contract admin request body" + "." + "base_req" + "." + "fees" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("update contract admin request body" + "." + "base_req" + "." + "fees" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostWasmContractsContractAddressAdminUpdateParamsBodyBaseReq) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostWasmContractsContractAddressAdminUpdateParamsBodyBaseReq) UnmarshalBinary(b []byte) error {
	var res PostWasmContractsContractAddressAdminUpdateParamsBodyBaseReq
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostWasmContractsContractAddressAdminUpdateParamsBodyBaseReqFeesItems0 post wasm contracts contract address admin update params body base req fees items0
swagger:model PostWasmContractsContractAddressAdminUpdateParamsBodyBaseReqFeesItems0
*/
type PostWasmContractsContractAddressAdminUpdateParamsBodyBaseReqFeesItems0 struct {

	// amount
	// Example: 50
	Amount string `json:"amount,omitempty"`

	// denom
	// Example: uluna
	Denom string `json:"denom,omitempty"`
}

// Validate validates this post wasm contracts contract address admin update params body base req fees items0
func (o *PostWasmContractsContractAddressAdminUpdateParamsBodyBaseReqFeesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post wasm contracts contract address admin update params body base req fees items0 based on context it is used
func (o *PostWasmContractsContractAddressAdminUpdateParamsBodyBaseReqFeesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostWasmContractsContractAddressAdminUpdateParamsBodyBaseReqFeesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostWasmContractsContractAddressAdminUpdateParamsBodyBaseReqFeesItems0) UnmarshalBinary(b []byte) error {
	var res PostWasmContractsContractAddressAdminUpdateParamsBodyBaseReqFeesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
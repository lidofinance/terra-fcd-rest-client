// Code generated by go-swagger; DO NOT EDIT.

package governance

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

// PostGovProposalsParamChangeReader is a Reader for the PostGovProposalsParamChange structure.
type PostGovProposalsParamChangeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostGovProposalsParamChangeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostGovProposalsParamChangeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostGovProposalsParamChangeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostGovProposalsParamChangeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostGovProposalsParamChangeOK creates a PostGovProposalsParamChangeOK with default headers values
func NewPostGovProposalsParamChangeOK() *PostGovProposalsParamChangeOK {
	return &PostGovProposalsParamChangeOK{}
}

/* PostGovProposalsParamChangeOK describes a response with status code 200, with default header values.

The transaction was successfully generated
*/
type PostGovProposalsParamChangeOK struct {
	Payload *PostGovProposalsParamChangeOKBody
}

func (o *PostGovProposalsParamChangeOK) Error() string {
	return fmt.Sprintf("[POST /gov/proposals/param_change][%d] postGovProposalsParamChangeOK  %+v", 200, o.Payload)
}
func (o *PostGovProposalsParamChangeOK) GetPayload() *PostGovProposalsParamChangeOKBody {
	return o.Payload
}

func (o *PostGovProposalsParamChangeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostGovProposalsParamChangeOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGovProposalsParamChangeBadRequest creates a PostGovProposalsParamChangeBadRequest with default headers values
func NewPostGovProposalsParamChangeBadRequest() *PostGovProposalsParamChangeBadRequest {
	return &PostGovProposalsParamChangeBadRequest{}
}

/* PostGovProposalsParamChangeBadRequest describes a response with status code 400, with default header values.

Invalid proposal body
*/
type PostGovProposalsParamChangeBadRequest struct {
}

func (o *PostGovProposalsParamChangeBadRequest) Error() string {
	return fmt.Sprintf("[POST /gov/proposals/param_change][%d] postGovProposalsParamChangeBadRequest ", 400)
}

func (o *PostGovProposalsParamChangeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostGovProposalsParamChangeInternalServerError creates a PostGovProposalsParamChangeInternalServerError with default headers values
func NewPostGovProposalsParamChangeInternalServerError() *PostGovProposalsParamChangeInternalServerError {
	return &PostGovProposalsParamChangeInternalServerError{}
}

/* PostGovProposalsParamChangeInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostGovProposalsParamChangeInternalServerError struct {
}

func (o *PostGovProposalsParamChangeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /gov/proposals/param_change][%d] postGovProposalsParamChangeInternalServerError ", 500)
}

func (o *PostGovProposalsParamChangeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PostGovProposalsParamChangeBody post gov proposals param change body
swagger:model PostGovProposalsParamChangeBody
*/
type PostGovProposalsParamChangeBody struct {

	// base req
	BaseReq *PostGovProposalsParamChangeParamsBodyBaseReq `json:"base_req,omitempty"`

	// changes
	Changes []*PostGovProposalsParamChangeParamsBodyChangesItems0 `json:"changes"`

	// deposit
	Deposit []*PostGovProposalsParamChangeParamsBodyDepositItems0 `json:"deposit"`

	// description
	Description string `json:"description,omitempty"`

	// bech32 encoded address
	// Example: terra1wg2mlrxdmnnkkykgqg4znky86nyrtc45q336yv
	Proposer string `json:"proposer,omitempty"`

	// title
	Title string `json:"title,omitempty"`
}

// Validate validates this post gov proposals param change body
func (o *PostGovProposalsParamChangeBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBaseReq(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateChanges(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDeposit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostGovProposalsParamChangeBody) validateBaseReq(formats strfmt.Registry) error {
	if swag.IsZero(o.BaseReq) { // not required
		return nil
	}

	if o.BaseReq != nil {
		if err := o.BaseReq.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("post_proposal_body" + "." + "base_req")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("post_proposal_body" + "." + "base_req")
			}
			return err
		}
	}

	return nil
}

func (o *PostGovProposalsParamChangeBody) validateChanges(formats strfmt.Registry) error {
	if swag.IsZero(o.Changes) { // not required
		return nil
	}

	for i := 0; i < len(o.Changes); i++ {
		if swag.IsZero(o.Changes[i]) { // not required
			continue
		}

		if o.Changes[i] != nil {
			if err := o.Changes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("post_proposal_body" + "." + "changes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("post_proposal_body" + "." + "changes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *PostGovProposalsParamChangeBody) validateDeposit(formats strfmt.Registry) error {
	if swag.IsZero(o.Deposit) { // not required
		return nil
	}

	for i := 0; i < len(o.Deposit); i++ {
		if swag.IsZero(o.Deposit[i]) { // not required
			continue
		}

		if o.Deposit[i] != nil {
			if err := o.Deposit[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("post_proposal_body" + "." + "deposit" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("post_proposal_body" + "." + "deposit" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this post gov proposals param change body based on the context it is used
func (o *PostGovProposalsParamChangeBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateBaseReq(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateChanges(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateDeposit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostGovProposalsParamChangeBody) contextValidateBaseReq(ctx context.Context, formats strfmt.Registry) error {

	if o.BaseReq != nil {
		if err := o.BaseReq.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("post_proposal_body" + "." + "base_req")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("post_proposal_body" + "." + "base_req")
			}
			return err
		}
	}

	return nil
}

func (o *PostGovProposalsParamChangeBody) contextValidateChanges(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Changes); i++ {

		if o.Changes[i] != nil {
			if err := o.Changes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("post_proposal_body" + "." + "changes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("post_proposal_body" + "." + "changes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *PostGovProposalsParamChangeBody) contextValidateDeposit(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Deposit); i++ {

		if o.Deposit[i] != nil {
			if err := o.Deposit[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("post_proposal_body" + "." + "deposit" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("post_proposal_body" + "." + "deposit" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostGovProposalsParamChangeBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostGovProposalsParamChangeBody) UnmarshalBinary(b []byte) error {
	var res PostGovProposalsParamChangeBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostGovProposalsParamChangeOKBody post gov proposals param change o k body
swagger:model PostGovProposalsParamChangeOKBody
*/
type PostGovProposalsParamChangeOKBody struct {

	// fee
	Fee *PostGovProposalsParamChangeOKBodyFee `json:"fee,omitempty"`

	// memo
	Memo string `json:"memo,omitempty"`

	// msg
	Msg []string `json:"msg"`

	// signature
	Signature *PostGovProposalsParamChangeOKBodySignature `json:"signature,omitempty"`
}

// Validate validates this post gov proposals param change o k body
func (o *PostGovProposalsParamChangeOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PostGovProposalsParamChangeOKBody) validateFee(formats strfmt.Registry) error {
	if swag.IsZero(o.Fee) { // not required
		return nil
	}

	if o.Fee != nil {
		if err := o.Fee.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postGovProposalsParamChangeOK" + "." + "fee")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postGovProposalsParamChangeOK" + "." + "fee")
			}
			return err
		}
	}

	return nil
}

func (o *PostGovProposalsParamChangeOKBody) validateSignature(formats strfmt.Registry) error {
	if swag.IsZero(o.Signature) { // not required
		return nil
	}

	if o.Signature != nil {
		if err := o.Signature.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postGovProposalsParamChangeOK" + "." + "signature")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postGovProposalsParamChangeOK" + "." + "signature")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post gov proposals param change o k body based on the context it is used
func (o *PostGovProposalsParamChangeOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *PostGovProposalsParamChangeOKBody) contextValidateFee(ctx context.Context, formats strfmt.Registry) error {

	if o.Fee != nil {
		if err := o.Fee.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postGovProposalsParamChangeOK" + "." + "fee")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postGovProposalsParamChangeOK" + "." + "fee")
			}
			return err
		}
	}

	return nil
}

func (o *PostGovProposalsParamChangeOKBody) contextValidateSignature(ctx context.Context, formats strfmt.Registry) error {

	if o.Signature != nil {
		if err := o.Signature.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postGovProposalsParamChangeOK" + "." + "signature")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postGovProposalsParamChangeOK" + "." + "signature")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostGovProposalsParamChangeOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostGovProposalsParamChangeOKBody) UnmarshalBinary(b []byte) error {
	var res PostGovProposalsParamChangeOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostGovProposalsParamChangeOKBodyFee post gov proposals param change o k body fee
swagger:model PostGovProposalsParamChangeOKBodyFee
*/
type PostGovProposalsParamChangeOKBodyFee struct {

	// amount
	Amount []*PostGovProposalsParamChangeOKBodyFeeAmountItems0 `json:"amount"`

	// gas
	Gas string `json:"gas,omitempty"`
}

// Validate validates this post gov proposals param change o k body fee
func (o *PostGovProposalsParamChangeOKBodyFee) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostGovProposalsParamChangeOKBodyFee) validateAmount(formats strfmt.Registry) error {
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
					return ve.ValidateName("postGovProposalsParamChangeOK" + "." + "fee" + "." + "amount" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("postGovProposalsParamChangeOK" + "." + "fee" + "." + "amount" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this post gov proposals param change o k body fee based on the context it is used
func (o *PostGovProposalsParamChangeOKBodyFee) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostGovProposalsParamChangeOKBodyFee) contextValidateAmount(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Amount); i++ {

		if o.Amount[i] != nil {
			if err := o.Amount[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("postGovProposalsParamChangeOK" + "." + "fee" + "." + "amount" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("postGovProposalsParamChangeOK" + "." + "fee" + "." + "amount" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostGovProposalsParamChangeOKBodyFee) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostGovProposalsParamChangeOKBodyFee) UnmarshalBinary(b []byte) error {
	var res PostGovProposalsParamChangeOKBodyFee
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostGovProposalsParamChangeOKBodyFeeAmountItems0 post gov proposals param change o k body fee amount items0
swagger:model PostGovProposalsParamChangeOKBodyFeeAmountItems0
*/
type PostGovProposalsParamChangeOKBodyFeeAmountItems0 struct {

	// amount
	// Example: 50
	Amount string `json:"amount,omitempty"`

	// denom
	// Example: uluna
	Denom string `json:"denom,omitempty"`
}

// Validate validates this post gov proposals param change o k body fee amount items0
func (o *PostGovProposalsParamChangeOKBodyFeeAmountItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post gov proposals param change o k body fee amount items0 based on context it is used
func (o *PostGovProposalsParamChangeOKBodyFeeAmountItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostGovProposalsParamChangeOKBodyFeeAmountItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostGovProposalsParamChangeOKBodyFeeAmountItems0) UnmarshalBinary(b []byte) error {
	var res PostGovProposalsParamChangeOKBodyFeeAmountItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostGovProposalsParamChangeOKBodySignature post gov proposals param change o k body signature
swagger:model PostGovProposalsParamChangeOKBodySignature
*/
type PostGovProposalsParamChangeOKBodySignature struct {

	// account number
	// Example: 0
	AccountNumber string `json:"account_number,omitempty"`

	// pub key
	PubKey *PostGovProposalsParamChangeOKBodySignaturePubKey `json:"pub_key,omitempty"`

	// sequence
	// Example: 0
	Sequence string `json:"sequence,omitempty"`

	// signature
	// Example: MEUCIQD02fsDPra8MtbRsyB1w7bqTM55Wu138zQbFcWx4+CFyAIge5WNPfKIuvzBZ69MyqHsqD8S1IwiEp+iUb6VSdtlpgY=
	Signature string `json:"signature,omitempty"`
}

// Validate validates this post gov proposals param change o k body signature
func (o *PostGovProposalsParamChangeOKBodySignature) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePubKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostGovProposalsParamChangeOKBodySignature) validatePubKey(formats strfmt.Registry) error {
	if swag.IsZero(o.PubKey) { // not required
		return nil
	}

	if o.PubKey != nil {
		if err := o.PubKey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postGovProposalsParamChangeOK" + "." + "signature" + "." + "pub_key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postGovProposalsParamChangeOK" + "." + "signature" + "." + "pub_key")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post gov proposals param change o k body signature based on the context it is used
func (o *PostGovProposalsParamChangeOKBodySignature) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePubKey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostGovProposalsParamChangeOKBodySignature) contextValidatePubKey(ctx context.Context, formats strfmt.Registry) error {

	if o.PubKey != nil {
		if err := o.PubKey.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("postGovProposalsParamChangeOK" + "." + "signature" + "." + "pub_key")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("postGovProposalsParamChangeOK" + "." + "signature" + "." + "pub_key")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostGovProposalsParamChangeOKBodySignature) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostGovProposalsParamChangeOKBodySignature) UnmarshalBinary(b []byte) error {
	var res PostGovProposalsParamChangeOKBodySignature
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostGovProposalsParamChangeOKBodySignaturePubKey post gov proposals param change o k body signature pub key
swagger:model PostGovProposalsParamChangeOKBodySignaturePubKey
*/
type PostGovProposalsParamChangeOKBodySignaturePubKey struct {

	// type
	// Example: tendermint/PubKeySecp256k1
	Type string `json:"type,omitempty"`

	// value
	// Example: Avz04VhtKJh8ACCVzlI8aTosGy0ikFXKIVHQ3jKMrosH
	Value string `json:"value,omitempty"`
}

// Validate validates this post gov proposals param change o k body signature pub key
func (o *PostGovProposalsParamChangeOKBodySignaturePubKey) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post gov proposals param change o k body signature pub key based on context it is used
func (o *PostGovProposalsParamChangeOKBodySignaturePubKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostGovProposalsParamChangeOKBodySignaturePubKey) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostGovProposalsParamChangeOKBodySignaturePubKey) UnmarshalBinary(b []byte) error {
	var res PostGovProposalsParamChangeOKBodySignaturePubKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostGovProposalsParamChangeParamsBodyBaseReq post gov proposals param change params body base req
swagger:model PostGovProposalsParamChangeParamsBodyBaseReq
*/
type PostGovProposalsParamChangeParamsBodyBaseReq struct {

	// account number
	// Example: 0
	AccountNumber string `json:"account_number,omitempty"`

	// chain id
	// Example: Columbus-5
	ChainID string `json:"chain_id,omitempty"`

	// fees
	Fees []*PostGovProposalsParamChangeParamsBodyBaseReqFeesItems0 `json:"fees"`

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

// Validate validates this post gov proposals param change params body base req
func (o *PostGovProposalsParamChangeParamsBodyBaseReq) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFees(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostGovProposalsParamChangeParamsBodyBaseReq) validateFees(formats strfmt.Registry) error {
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
					return ve.ValidateName("post_proposal_body" + "." + "base_req" + "." + "fees" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("post_proposal_body" + "." + "base_req" + "." + "fees" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this post gov proposals param change params body base req based on the context it is used
func (o *PostGovProposalsParamChangeParamsBodyBaseReq) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateFees(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostGovProposalsParamChangeParamsBodyBaseReq) contextValidateFees(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Fees); i++ {

		if o.Fees[i] != nil {
			if err := o.Fees[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("post_proposal_body" + "." + "base_req" + "." + "fees" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("post_proposal_body" + "." + "base_req" + "." + "fees" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *PostGovProposalsParamChangeParamsBodyBaseReq) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostGovProposalsParamChangeParamsBodyBaseReq) UnmarshalBinary(b []byte) error {
	var res PostGovProposalsParamChangeParamsBodyBaseReq
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostGovProposalsParamChangeParamsBodyBaseReqFeesItems0 post gov proposals param change params body base req fees items0
swagger:model PostGovProposalsParamChangeParamsBodyBaseReqFeesItems0
*/
type PostGovProposalsParamChangeParamsBodyBaseReqFeesItems0 struct {

	// amount
	// Example: 50
	Amount string `json:"amount,omitempty"`

	// denom
	// Example: uluna
	Denom string `json:"denom,omitempty"`
}

// Validate validates this post gov proposals param change params body base req fees items0
func (o *PostGovProposalsParamChangeParamsBodyBaseReqFeesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post gov proposals param change params body base req fees items0 based on context it is used
func (o *PostGovProposalsParamChangeParamsBodyBaseReqFeesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostGovProposalsParamChangeParamsBodyBaseReqFeesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostGovProposalsParamChangeParamsBodyBaseReqFeesItems0) UnmarshalBinary(b []byte) error {
	var res PostGovProposalsParamChangeParamsBodyBaseReqFeesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostGovProposalsParamChangeParamsBodyChangesItems0 post gov proposals param change params body changes items0
swagger:model PostGovProposalsParamChangeParamsBodyChangesItems0
*/
type PostGovProposalsParamChangeParamsBodyChangesItems0 struct {

	// key
	// Example: MaxValidators
	Key string `json:"key,omitempty"`

	// subkey
	Subkey string `json:"subkey,omitempty"`

	// subspace
	// Example: staking
	Subspace string `json:"subspace,omitempty"`

	// value
	Value interface{} `json:"value,omitempty"`
}

// Validate validates this post gov proposals param change params body changes items0
func (o *PostGovProposalsParamChangeParamsBodyChangesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post gov proposals param change params body changes items0 based on context it is used
func (o *PostGovProposalsParamChangeParamsBodyChangesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostGovProposalsParamChangeParamsBodyChangesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostGovProposalsParamChangeParamsBodyChangesItems0) UnmarshalBinary(b []byte) error {
	var res PostGovProposalsParamChangeParamsBodyChangesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*PostGovProposalsParamChangeParamsBodyDepositItems0 post gov proposals param change params body deposit items0
swagger:model PostGovProposalsParamChangeParamsBodyDepositItems0
*/
type PostGovProposalsParamChangeParamsBodyDepositItems0 struct {

	// amount
	// Example: 50
	Amount string `json:"amount,omitempty"`

	// denom
	// Example: uluna
	Denom string `json:"denom,omitempty"`
}

// Validate validates this post gov proposals param change params body deposit items0
func (o *PostGovProposalsParamChangeParamsBodyDepositItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post gov proposals param change params body deposit items0 based on context it is used
func (o *PostGovProposalsParamChangeParamsBodyDepositItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostGovProposalsParamChangeParamsBodyDepositItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostGovProposalsParamChangeParamsBodyDepositItems0) UnmarshalBinary(b []byte) error {
	var res PostGovProposalsParamChangeParamsBodyDepositItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

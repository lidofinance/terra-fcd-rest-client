// Code generated by go-swagger; DO NOT EDIT.

package distribution

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

// GetDistributionDelegatorsDelegatorAddrRewardsReader is a Reader for the GetDistributionDelegatorsDelegatorAddrRewards structure.
type GetDistributionDelegatorsDelegatorAddrRewardsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDistributionDelegatorsDelegatorAddrRewardsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDistributionDelegatorsDelegatorAddrRewardsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetDistributionDelegatorsDelegatorAddrRewardsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetDistributionDelegatorsDelegatorAddrRewardsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDistributionDelegatorsDelegatorAddrRewardsOK creates a GetDistributionDelegatorsDelegatorAddrRewardsOK with default headers values
func NewGetDistributionDelegatorsDelegatorAddrRewardsOK() *GetDistributionDelegatorsDelegatorAddrRewardsOK {
	return &GetDistributionDelegatorsDelegatorAddrRewardsOK{}
}

/* GetDistributionDelegatorsDelegatorAddrRewardsOK describes a response with status code 200, with default header values.

OK
*/
type GetDistributionDelegatorsDelegatorAddrRewardsOK struct {
	Payload *GetDistributionDelegatorsDelegatorAddrRewardsOKBody
}

func (o *GetDistributionDelegatorsDelegatorAddrRewardsOK) Error() string {
	return fmt.Sprintf("[GET /distribution/delegators/{delegatorAddr}/rewards][%d] getDistributionDelegatorsDelegatorAddrRewardsOK  %+v", 200, o.Payload)
}
func (o *GetDistributionDelegatorsDelegatorAddrRewardsOK) GetPayload() *GetDistributionDelegatorsDelegatorAddrRewardsOKBody {
	return o.Payload
}

func (o *GetDistributionDelegatorsDelegatorAddrRewardsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetDistributionDelegatorsDelegatorAddrRewardsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDistributionDelegatorsDelegatorAddrRewardsBadRequest creates a GetDistributionDelegatorsDelegatorAddrRewardsBadRequest with default headers values
func NewGetDistributionDelegatorsDelegatorAddrRewardsBadRequest() *GetDistributionDelegatorsDelegatorAddrRewardsBadRequest {
	return &GetDistributionDelegatorsDelegatorAddrRewardsBadRequest{}
}

/* GetDistributionDelegatorsDelegatorAddrRewardsBadRequest describes a response with status code 400, with default header values.

Invalid delegator address
*/
type GetDistributionDelegatorsDelegatorAddrRewardsBadRequest struct {
}

func (o *GetDistributionDelegatorsDelegatorAddrRewardsBadRequest) Error() string {
	return fmt.Sprintf("[GET /distribution/delegators/{delegatorAddr}/rewards][%d] getDistributionDelegatorsDelegatorAddrRewardsBadRequest ", 400)
}

func (o *GetDistributionDelegatorsDelegatorAddrRewardsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetDistributionDelegatorsDelegatorAddrRewardsInternalServerError creates a GetDistributionDelegatorsDelegatorAddrRewardsInternalServerError with default headers values
func NewGetDistributionDelegatorsDelegatorAddrRewardsInternalServerError() *GetDistributionDelegatorsDelegatorAddrRewardsInternalServerError {
	return &GetDistributionDelegatorsDelegatorAddrRewardsInternalServerError{}
}

/* GetDistributionDelegatorsDelegatorAddrRewardsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetDistributionDelegatorsDelegatorAddrRewardsInternalServerError struct {
}

func (o *GetDistributionDelegatorsDelegatorAddrRewardsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /distribution/delegators/{delegatorAddr}/rewards][%d] getDistributionDelegatorsDelegatorAddrRewardsInternalServerError ", 500)
}

func (o *GetDistributionDelegatorsDelegatorAddrRewardsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetDistributionDelegatorsDelegatorAddrRewardsOKBody get distribution delegators delegator addr rewards o k body
swagger:model GetDistributionDelegatorsDelegatorAddrRewardsOKBody
*/
type GetDistributionDelegatorsDelegatorAddrRewardsOKBody struct {

	// rewards
	Rewards []*GetDistributionDelegatorsDelegatorAddrRewardsOKBodyRewardsItems0 `json:"rewards"`

	// total
	Total []*GetDistributionDelegatorsDelegatorAddrRewardsOKBodyTotalItems0 `json:"total"`
}

// Validate validates this get distribution delegators delegator addr rewards o k body
func (o *GetDistributionDelegatorsDelegatorAddrRewardsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRewards(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTotal(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetDistributionDelegatorsDelegatorAddrRewardsOKBody) validateRewards(formats strfmt.Registry) error {
	if swag.IsZero(o.Rewards) { // not required
		return nil
	}

	for i := 0; i < len(o.Rewards); i++ {
		if swag.IsZero(o.Rewards[i]) { // not required
			continue
		}

		if o.Rewards[i] != nil {
			if err := o.Rewards[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getDistributionDelegatorsDelegatorAddrRewardsOK" + "." + "rewards" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getDistributionDelegatorsDelegatorAddrRewardsOK" + "." + "rewards" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetDistributionDelegatorsDelegatorAddrRewardsOKBody) validateTotal(formats strfmt.Registry) error {
	if swag.IsZero(o.Total) { // not required
		return nil
	}

	for i := 0; i < len(o.Total); i++ {
		if swag.IsZero(o.Total[i]) { // not required
			continue
		}

		if o.Total[i] != nil {
			if err := o.Total[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getDistributionDelegatorsDelegatorAddrRewardsOK" + "." + "total" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getDistributionDelegatorsDelegatorAddrRewardsOK" + "." + "total" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get distribution delegators delegator addr rewards o k body based on the context it is used
func (o *GetDistributionDelegatorsDelegatorAddrRewardsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRewards(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateTotal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetDistributionDelegatorsDelegatorAddrRewardsOKBody) contextValidateRewards(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Rewards); i++ {

		if o.Rewards[i] != nil {
			if err := o.Rewards[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getDistributionDelegatorsDelegatorAddrRewardsOK" + "." + "rewards" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getDistributionDelegatorsDelegatorAddrRewardsOK" + "." + "rewards" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetDistributionDelegatorsDelegatorAddrRewardsOKBody) contextValidateTotal(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Total); i++ {

		if o.Total[i] != nil {
			if err := o.Total[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getDistributionDelegatorsDelegatorAddrRewardsOK" + "." + "total" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getDistributionDelegatorsDelegatorAddrRewardsOK" + "." + "total" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetDistributionDelegatorsDelegatorAddrRewardsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDistributionDelegatorsDelegatorAddrRewardsOKBody) UnmarshalBinary(b []byte) error {
	var res GetDistributionDelegatorsDelegatorAddrRewardsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetDistributionDelegatorsDelegatorAddrRewardsOKBodyRewardsItems0 get distribution delegators delegator addr rewards o k body rewards items0
swagger:model GetDistributionDelegatorsDelegatorAddrRewardsOKBodyRewardsItems0
*/
type GetDistributionDelegatorsDelegatorAddrRewardsOKBodyRewardsItems0 struct {

	// reward
	Reward []*GetDistributionDelegatorsDelegatorAddrRewardsOKBodyRewardsItems0RewardItems0 `json:"reward"`

	// bech32 encoded address
	// Example: terravaloper1wg2mlrxdmnnkkykgqg4znky86nyrtc45q7a85l
	ValidatorAddress string `json:"validator_address,omitempty"`
}

// Validate validates this get distribution delegators delegator addr rewards o k body rewards items0
func (o *GetDistributionDelegatorsDelegatorAddrRewardsOKBodyRewardsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateReward(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetDistributionDelegatorsDelegatorAddrRewardsOKBodyRewardsItems0) validateReward(formats strfmt.Registry) error {
	if swag.IsZero(o.Reward) { // not required
		return nil
	}

	for i := 0; i < len(o.Reward); i++ {
		if swag.IsZero(o.Reward[i]) { // not required
			continue
		}

		if o.Reward[i] != nil {
			if err := o.Reward[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("reward" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("reward" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get distribution delegators delegator addr rewards o k body rewards items0 based on the context it is used
func (o *GetDistributionDelegatorsDelegatorAddrRewardsOKBodyRewardsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateReward(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetDistributionDelegatorsDelegatorAddrRewardsOKBodyRewardsItems0) contextValidateReward(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Reward); i++ {

		if o.Reward[i] != nil {
			if err := o.Reward[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("reward" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("reward" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetDistributionDelegatorsDelegatorAddrRewardsOKBodyRewardsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDistributionDelegatorsDelegatorAddrRewardsOKBodyRewardsItems0) UnmarshalBinary(b []byte) error {
	var res GetDistributionDelegatorsDelegatorAddrRewardsOKBodyRewardsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetDistributionDelegatorsDelegatorAddrRewardsOKBodyRewardsItems0RewardItems0 get distribution delegators delegator addr rewards o k body rewards items0 reward items0
swagger:model GetDistributionDelegatorsDelegatorAddrRewardsOKBodyRewardsItems0RewardItems0
*/
type GetDistributionDelegatorsDelegatorAddrRewardsOKBodyRewardsItems0RewardItems0 struct {

	// amount
	// Example: 50
	Amount string `json:"amount,omitempty"`

	// denom
	// Example: uluna
	Denom string `json:"denom,omitempty"`
}

// Validate validates this get distribution delegators delegator addr rewards o k body rewards items0 reward items0
func (o *GetDistributionDelegatorsDelegatorAddrRewardsOKBodyRewardsItems0RewardItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get distribution delegators delegator addr rewards o k body rewards items0 reward items0 based on context it is used
func (o *GetDistributionDelegatorsDelegatorAddrRewardsOKBodyRewardsItems0RewardItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetDistributionDelegatorsDelegatorAddrRewardsOKBodyRewardsItems0RewardItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDistributionDelegatorsDelegatorAddrRewardsOKBodyRewardsItems0RewardItems0) UnmarshalBinary(b []byte) error {
	var res GetDistributionDelegatorsDelegatorAddrRewardsOKBodyRewardsItems0RewardItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetDistributionDelegatorsDelegatorAddrRewardsOKBodyTotalItems0 get distribution delegators delegator addr rewards o k body total items0
swagger:model GetDistributionDelegatorsDelegatorAddrRewardsOKBodyTotalItems0
*/
type GetDistributionDelegatorsDelegatorAddrRewardsOKBodyTotalItems0 struct {

	// amount
	// Example: 50
	Amount string `json:"amount,omitempty"`

	// denom
	// Example: uluna
	Denom string `json:"denom,omitempty"`
}

// Validate validates this get distribution delegators delegator addr rewards o k body total items0
func (o *GetDistributionDelegatorsDelegatorAddrRewardsOKBodyTotalItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get distribution delegators delegator addr rewards o k body total items0 based on context it is used
func (o *GetDistributionDelegatorsDelegatorAddrRewardsOKBodyTotalItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetDistributionDelegatorsDelegatorAddrRewardsOKBodyTotalItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDistributionDelegatorsDelegatorAddrRewardsOKBodyTotalItems0) UnmarshalBinary(b []byte) error {
	var res GetDistributionDelegatorsDelegatorAddrRewardsOKBodyTotalItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

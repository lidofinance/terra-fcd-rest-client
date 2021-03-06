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

// GetGovProposalsProposalIDDepositsDepositorReader is a Reader for the GetGovProposalsProposalIDDepositsDepositor structure.
type GetGovProposalsProposalIDDepositsDepositorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGovProposalsProposalIDDepositsDepositorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGovProposalsProposalIDDepositsDepositorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGovProposalsProposalIDDepositsDepositorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGovProposalsProposalIDDepositsDepositorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGovProposalsProposalIDDepositsDepositorInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGovProposalsProposalIDDepositsDepositorOK creates a GetGovProposalsProposalIDDepositsDepositorOK with default headers values
func NewGetGovProposalsProposalIDDepositsDepositorOK() *GetGovProposalsProposalIDDepositsDepositorOK {
	return &GetGovProposalsProposalIDDepositsDepositorOK{}
}

/* GetGovProposalsProposalIDDepositsDepositorOK describes a response with status code 200, with default header values.

OK
*/
type GetGovProposalsProposalIDDepositsDepositorOK struct {
	Payload *GetGovProposalsProposalIDDepositsDepositorOKBody
}

func (o *GetGovProposalsProposalIDDepositsDepositorOK) Error() string {
	return fmt.Sprintf("[GET /gov/proposals/{proposalId}/deposits/{depositor}][%d] getGovProposalsProposalIdDepositsDepositorOK  %+v", 200, o.Payload)
}
func (o *GetGovProposalsProposalIDDepositsDepositorOK) GetPayload() *GetGovProposalsProposalIDDepositsDepositorOKBody {
	return o.Payload
}

func (o *GetGovProposalsProposalIDDepositsDepositorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetGovProposalsProposalIDDepositsDepositorOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGovProposalsProposalIDDepositsDepositorBadRequest creates a GetGovProposalsProposalIDDepositsDepositorBadRequest with default headers values
func NewGetGovProposalsProposalIDDepositsDepositorBadRequest() *GetGovProposalsProposalIDDepositsDepositorBadRequest {
	return &GetGovProposalsProposalIDDepositsDepositorBadRequest{}
}

/* GetGovProposalsProposalIDDepositsDepositorBadRequest describes a response with status code 400, with default header values.

Invalid proposal id or depositor address
*/
type GetGovProposalsProposalIDDepositsDepositorBadRequest struct {
}

func (o *GetGovProposalsProposalIDDepositsDepositorBadRequest) Error() string {
	return fmt.Sprintf("[GET /gov/proposals/{proposalId}/deposits/{depositor}][%d] getGovProposalsProposalIdDepositsDepositorBadRequest ", 400)
}

func (o *GetGovProposalsProposalIDDepositsDepositorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetGovProposalsProposalIDDepositsDepositorNotFound creates a GetGovProposalsProposalIDDepositsDepositorNotFound with default headers values
func NewGetGovProposalsProposalIDDepositsDepositorNotFound() *GetGovProposalsProposalIDDepositsDepositorNotFound {
	return &GetGovProposalsProposalIDDepositsDepositorNotFound{}
}

/* GetGovProposalsProposalIDDepositsDepositorNotFound describes a response with status code 404, with default header values.

Found no deposit
*/
type GetGovProposalsProposalIDDepositsDepositorNotFound struct {
}

func (o *GetGovProposalsProposalIDDepositsDepositorNotFound) Error() string {
	return fmt.Sprintf("[GET /gov/proposals/{proposalId}/deposits/{depositor}][%d] getGovProposalsProposalIdDepositsDepositorNotFound ", 404)
}

func (o *GetGovProposalsProposalIDDepositsDepositorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetGovProposalsProposalIDDepositsDepositorInternalServerError creates a GetGovProposalsProposalIDDepositsDepositorInternalServerError with default headers values
func NewGetGovProposalsProposalIDDepositsDepositorInternalServerError() *GetGovProposalsProposalIDDepositsDepositorInternalServerError {
	return &GetGovProposalsProposalIDDepositsDepositorInternalServerError{}
}

/* GetGovProposalsProposalIDDepositsDepositorInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetGovProposalsProposalIDDepositsDepositorInternalServerError struct {
}

func (o *GetGovProposalsProposalIDDepositsDepositorInternalServerError) Error() string {
	return fmt.Sprintf("[GET /gov/proposals/{proposalId}/deposits/{depositor}][%d] getGovProposalsProposalIdDepositsDepositorInternalServerError ", 500)
}

func (o *GetGovProposalsProposalIDDepositsDepositorInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetGovProposalsProposalIDDepositsDepositorOKBody get gov proposals proposal ID deposits depositor o k body
swagger:model GetGovProposalsProposalIDDepositsDepositorOKBody
*/
type GetGovProposalsProposalIDDepositsDepositorOKBody struct {

	// amount
	Amount []*GetGovProposalsProposalIDDepositsDepositorOKBodyAmountItems0 `json:"amount"`

	// bech32 encoded address
	// Example: terra1wg2mlrxdmnnkkykgqg4znky86nyrtc45q336yv
	Depositor string `json:"depositor,omitempty"`

	// proposal id
	ProposalID string `json:"proposal_id,omitempty"`
}

// Validate validates this get gov proposals proposal ID deposits depositor o k body
func (o *GetGovProposalsProposalIDDepositsDepositorOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetGovProposalsProposalIDDepositsDepositorOKBody) validateAmount(formats strfmt.Registry) error {
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
					return ve.ValidateName("getGovProposalsProposalIdDepositsDepositorOK" + "." + "amount" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getGovProposalsProposalIdDepositsDepositorOK" + "." + "amount" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get gov proposals proposal ID deposits depositor o k body based on the context it is used
func (o *GetGovProposalsProposalIDDepositsDepositorOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAmount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetGovProposalsProposalIDDepositsDepositorOKBody) contextValidateAmount(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Amount); i++ {

		if o.Amount[i] != nil {
			if err := o.Amount[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getGovProposalsProposalIdDepositsDepositorOK" + "." + "amount" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getGovProposalsProposalIdDepositsDepositorOK" + "." + "amount" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetGovProposalsProposalIDDepositsDepositorOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetGovProposalsProposalIDDepositsDepositorOKBody) UnmarshalBinary(b []byte) error {
	var res GetGovProposalsProposalIDDepositsDepositorOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetGovProposalsProposalIDDepositsDepositorOKBodyAmountItems0 get gov proposals proposal ID deposits depositor o k body amount items0
swagger:model GetGovProposalsProposalIDDepositsDepositorOKBodyAmountItems0
*/
type GetGovProposalsProposalIDDepositsDepositorOKBodyAmountItems0 struct {

	// amount
	// Example: 50
	Amount string `json:"amount,omitempty"`

	// denom
	// Example: uluna
	Denom string `json:"denom,omitempty"`
}

// Validate validates this get gov proposals proposal ID deposits depositor o k body amount items0
func (o *GetGovProposalsProposalIDDepositsDepositorOKBodyAmountItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get gov proposals proposal ID deposits depositor o k body amount items0 based on context it is used
func (o *GetGovProposalsProposalIDDepositsDepositorOKBodyAmountItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetGovProposalsProposalIDDepositsDepositorOKBodyAmountItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetGovProposalsProposalIDDepositsDepositorOKBodyAmountItems0) UnmarshalBinary(b []byte) error {
	var res GetGovProposalsProposalIDDepositsDepositorOKBodyAmountItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package governance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetGovProposalsProposalIDProposerReader is a Reader for the GetGovProposalsProposalIDProposer structure.
type GetGovProposalsProposalIDProposerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGovProposalsProposalIDProposerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGovProposalsProposalIDProposerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGovProposalsProposalIDProposerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGovProposalsProposalIDProposerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGovProposalsProposalIDProposerOK creates a GetGovProposalsProposalIDProposerOK with default headers values
func NewGetGovProposalsProposalIDProposerOK() *GetGovProposalsProposalIDProposerOK {
	return &GetGovProposalsProposalIDProposerOK{}
}

/* GetGovProposalsProposalIDProposerOK describes a response with status code 200, with default header values.

OK
*/
type GetGovProposalsProposalIDProposerOK struct {
	Payload *GetGovProposalsProposalIDProposerOKBody
}

func (o *GetGovProposalsProposalIDProposerOK) Error() string {
	return fmt.Sprintf("[GET /gov/proposals/{proposalId}/proposer][%d] getGovProposalsProposalIdProposerOK  %+v", 200, o.Payload)
}
func (o *GetGovProposalsProposalIDProposerOK) GetPayload() *GetGovProposalsProposalIDProposerOKBody {
	return o.Payload
}

func (o *GetGovProposalsProposalIDProposerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetGovProposalsProposalIDProposerOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGovProposalsProposalIDProposerBadRequest creates a GetGovProposalsProposalIDProposerBadRequest with default headers values
func NewGetGovProposalsProposalIDProposerBadRequest() *GetGovProposalsProposalIDProposerBadRequest {
	return &GetGovProposalsProposalIDProposerBadRequest{}
}

/* GetGovProposalsProposalIDProposerBadRequest describes a response with status code 400, with default header values.

Invalid proposal ID
*/
type GetGovProposalsProposalIDProposerBadRequest struct {
}

func (o *GetGovProposalsProposalIDProposerBadRequest) Error() string {
	return fmt.Sprintf("[GET /gov/proposals/{proposalId}/proposer][%d] getGovProposalsProposalIdProposerBadRequest ", 400)
}

func (o *GetGovProposalsProposalIDProposerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetGovProposalsProposalIDProposerInternalServerError creates a GetGovProposalsProposalIDProposerInternalServerError with default headers values
func NewGetGovProposalsProposalIDProposerInternalServerError() *GetGovProposalsProposalIDProposerInternalServerError {
	return &GetGovProposalsProposalIDProposerInternalServerError{}
}

/* GetGovProposalsProposalIDProposerInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetGovProposalsProposalIDProposerInternalServerError struct {
}

func (o *GetGovProposalsProposalIDProposerInternalServerError) Error() string {
	return fmt.Sprintf("[GET /gov/proposals/{proposalId}/proposer][%d] getGovProposalsProposalIdProposerInternalServerError ", 500)
}

func (o *GetGovProposalsProposalIDProposerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetGovProposalsProposalIDProposerOKBody get gov proposals proposal ID proposer o k body
swagger:model GetGovProposalsProposalIDProposerOKBody
*/
type GetGovProposalsProposalIDProposerOKBody struct {

	// proposal id
	ProposalID string `json:"proposal_id,omitempty"`

	// proposer
	Proposer string `json:"proposer,omitempty"`
}

// Validate validates this get gov proposals proposal ID proposer o k body
func (o *GetGovProposalsProposalIDProposerOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get gov proposals proposal ID proposer o k body based on context it is used
func (o *GetGovProposalsProposalIDProposerOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetGovProposalsProposalIDProposerOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetGovProposalsProposalIDProposerOKBody) UnmarshalBinary(b []byte) error {
	var res GetGovProposalsProposalIDProposerOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

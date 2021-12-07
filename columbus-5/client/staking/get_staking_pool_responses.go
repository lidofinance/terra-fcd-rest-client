// Code generated by go-swagger; DO NOT EDIT.

package staking

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

// GetStakingPoolReader is a Reader for the GetStakingPool structure.
type GetStakingPoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStakingPoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetStakingPoolOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetStakingPoolInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetStakingPoolOK creates a GetStakingPoolOK with default headers values
func NewGetStakingPoolOK() *GetStakingPoolOK {
	return &GetStakingPoolOK{}
}

/* GetStakingPoolOK describes a response with status code 200, with default header values.

OK
*/
type GetStakingPoolOK struct {
	Payload *GetStakingPoolOKBody
}

func (o *GetStakingPoolOK) Error() string {
	return fmt.Sprintf("[GET /staking/pool][%d] getStakingPoolOK  %+v", 200, o.Payload)
}
func (o *GetStakingPoolOK) GetPayload() *GetStakingPoolOKBody {
	return o.Payload
}

func (o *GetStakingPoolOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetStakingPoolOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStakingPoolInternalServerError creates a GetStakingPoolInternalServerError with default headers values
func NewGetStakingPoolInternalServerError() *GetStakingPoolInternalServerError {
	return &GetStakingPoolInternalServerError{}
}

/* GetStakingPoolInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetStakingPoolInternalServerError struct {
}

func (o *GetStakingPoolInternalServerError) Error() string {
	return fmt.Sprintf("[GET /staking/pool][%d] getStakingPoolInternalServerError ", 500)
}

func (o *GetStakingPoolInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetStakingPoolOKBody get staking pool o k body
swagger:model GetStakingPoolOKBody
*/
type GetStakingPoolOKBody struct {

	// bonded tokens
	BondedTokens string `json:"bonded_tokens,omitempty"`

	// not bonded tokens
	NotBondedTokens string `json:"not_bonded_tokens,omitempty"`
}

// Validate validates this get staking pool o k body
func (o *GetStakingPoolOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get staking pool o k body based on context it is used
func (o *GetStakingPoolOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetStakingPoolOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetStakingPoolOKBody) UnmarshalBinary(b []byte) error {
	var res GetStakingPoolOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
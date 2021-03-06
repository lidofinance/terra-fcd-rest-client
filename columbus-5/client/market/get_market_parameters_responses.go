// Code generated by go-swagger; DO NOT EDIT.

package market

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

// GetMarketParametersReader is a Reader for the GetMarketParameters structure.
type GetMarketParametersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMarketParametersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMarketParametersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetMarketParametersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetMarketParametersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetMarketParametersOK creates a GetMarketParametersOK with default headers values
func NewGetMarketParametersOK() *GetMarketParametersOK {
	return &GetMarketParametersOK{}
}

/* GetMarketParametersOK describes a response with status code 200, with default header values.

OK
*/
type GetMarketParametersOK struct {
	Payload *GetMarketParametersOKBody
}

func (o *GetMarketParametersOK) Error() string {
	return fmt.Sprintf("[GET /market/parameters][%d] getMarketParametersOK  %+v", 200, o.Payload)
}
func (o *GetMarketParametersOK) GetPayload() *GetMarketParametersOKBody {
	return o.Payload
}

func (o *GetMarketParametersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetMarketParametersOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMarketParametersBadRequest creates a GetMarketParametersBadRequest with default headers values
func NewGetMarketParametersBadRequest() *GetMarketParametersBadRequest {
	return &GetMarketParametersBadRequest{}
}

/* GetMarketParametersBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetMarketParametersBadRequest struct {
}

func (o *GetMarketParametersBadRequest) Error() string {
	return fmt.Sprintf("[GET /market/parameters][%d] getMarketParametersBadRequest ", 400)
}

func (o *GetMarketParametersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetMarketParametersInternalServerError creates a GetMarketParametersInternalServerError with default headers values
func NewGetMarketParametersInternalServerError() *GetMarketParametersInternalServerError {
	return &GetMarketParametersInternalServerError{}
}

/* GetMarketParametersInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetMarketParametersInternalServerError struct {
}

func (o *GetMarketParametersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /market/parameters][%d] getMarketParametersInternalServerError ", 500)
}

func (o *GetMarketParametersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetMarketParametersOKBody get market parameters o k body
swagger:model GetMarketParametersOKBody
*/
type GetMarketParametersOKBody struct {

	// base pool
	// Example: 1000000.0
	BasePool float64 `json:"base_pool,omitempty"`

	// min stability spread
	// Example: 0.02
	MinStabilitySpread float64 `json:"min_stability_spread,omitempty"`

	// pool recovery period
	// Example: 16000
	PoolRecoveryPeriod float64 `json:"pool_recovery_period,omitempty"`
}

// Validate validates this get market parameters o k body
func (o *GetMarketParametersOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get market parameters o k body based on context it is used
func (o *GetMarketParametersOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetMarketParametersOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetMarketParametersOKBody) UnmarshalBinary(b []byte) error {
	var res GetMarketParametersOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

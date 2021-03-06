// Code generated by go-swagger; DO NOT EDIT.

package treasury

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/lidofinance/terra-fcd-rest-client/columbus-5/models"
)

// GetV1RichlistDenomReader is a Reader for the GetV1RichlistDenom structure.
type GetV1RichlistDenomReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1RichlistDenomReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1RichlistDenomOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1RichlistDenomOK creates a GetV1RichlistDenomOK with default headers values
func NewGetV1RichlistDenomOK() *GetV1RichlistDenomOK {
	return &GetV1RichlistDenomOK{}
}

/* GetV1RichlistDenomOK describes a response with status code 200, with default header values.

Success
*/
type GetV1RichlistDenomOK struct {
	Payload []*models.Accounts
}

func (o *GetV1RichlistDenomOK) Error() string {
	return fmt.Sprintf("[GET /v1/richlist/{denom}][%d] getV1RichlistDenomOK  %+v", 200, o.Payload)
}
func (o *GetV1RichlistDenomOK) GetPayload() []*models.Accounts {
	return o.Payload
}

func (o *GetV1RichlistDenomOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

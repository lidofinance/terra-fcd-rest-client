// Code generated by go-swagger; DO NOT EDIT.

package treasury

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetV1CirculatingsupplyDenomReader is a Reader for the GetV1CirculatingsupplyDenom structure.
type GetV1CirculatingsupplyDenomReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1CirculatingsupplyDenomReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1CirculatingsupplyDenomOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1CirculatingsupplyDenomOK creates a GetV1CirculatingsupplyDenomOK with default headers values
func NewGetV1CirculatingsupplyDenomOK() *GetV1CirculatingsupplyDenomOK {
	return &GetV1CirculatingsupplyDenomOK{}
}

/* GetV1CirculatingsupplyDenomOK describes a response with status code 200, with default header values.

Success
*/
type GetV1CirculatingsupplyDenomOK struct {
	Payload float64
}

func (o *GetV1CirculatingsupplyDenomOK) Error() string {
	return fmt.Sprintf("[GET /v1/circulatingsupply/{denom}][%d] getV1CirculatingsupplyDenomOK  %+v", 200, o.Payload)
}
func (o *GetV1CirculatingsupplyDenomOK) GetPayload() float64 {
	return o.Payload
}

func (o *GetV1CirculatingsupplyDenomOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

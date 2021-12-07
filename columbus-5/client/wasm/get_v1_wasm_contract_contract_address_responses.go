// Code generated by go-swagger; DO NOT EDIT.

package wasm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/lidofinance/terra-fcd-rest-client/columbus-5/models"
)

// GetV1WasmContractContractAddressReader is a Reader for the GetV1WasmContractContractAddress structure.
type GetV1WasmContractContractAddressReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1WasmContractContractAddressReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1WasmContractContractAddressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1WasmContractContractAddressOK creates a GetV1WasmContractContractAddressOK with default headers values
func NewGetV1WasmContractContractAddressOK() *GetV1WasmContractContractAddressOK {
	return &GetV1WasmContractContractAddressOK{}
}

/* GetV1WasmContractContractAddressOK describes a response with status code 200, with default header values.

Success
*/
type GetV1WasmContractContractAddressOK struct {
	Payload *models.GetIndividualWasmContractResult
}

func (o *GetV1WasmContractContractAddressOK) Error() string {
	return fmt.Sprintf("[GET /v1/wasm/contract/{contractAddress}][%d] getV1WasmContractContractAddressOK  %+v", 200, o.Payload)
}
func (o *GetV1WasmContractContractAddressOK) GetPayload() *models.GetIndividualWasmContractResult {
	return o.Payload
}

func (o *GetV1WasmContractContractAddressOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetIndividualWasmContractResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
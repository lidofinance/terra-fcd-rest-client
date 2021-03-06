// Code generated by go-swagger; DO NOT EDIT.

package wasm

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

// GetWasmCodesCodeIDReader is a Reader for the GetWasmCodesCodeID structure.
type GetWasmCodesCodeIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWasmCodesCodeIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWasmCodesCodeIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWasmCodesCodeIDOK creates a GetWasmCodesCodeIDOK with default headers values
func NewGetWasmCodesCodeIDOK() *GetWasmCodesCodeIDOK {
	return &GetWasmCodesCodeIDOK{}
}

/* GetWasmCodesCodeIDOK describes a response with status code 200, with default header values.

OK
*/
type GetWasmCodesCodeIDOK struct {
	Payload *GetWasmCodesCodeIDOKBody
}

func (o *GetWasmCodesCodeIDOK) Error() string {
	return fmt.Sprintf("[GET /wasm/codes/{codeID}][%d] getWasmCodesCodeIdOK  %+v", 200, o.Payload)
}
func (o *GetWasmCodesCodeIDOK) GetPayload() *GetWasmCodesCodeIDOKBody {
	return o.Payload
}

func (o *GetWasmCodesCodeIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetWasmCodesCodeIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetWasmCodesCodeIDOKBody get wasm codes code ID o k body
swagger:model GetWasmCodesCodeIDOKBody
*/
type GetWasmCodesCodeIDOKBody struct {

	// code hash
	CodeHash string `json:"code_hash,omitempty"`

	// creator
	Creator string `json:"creator,omitempty"`
}

// Validate validates this get wasm codes code ID o k body
func (o *GetWasmCodesCodeIDOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get wasm codes code ID o k body based on context it is used
func (o *GetWasmCodesCodeIDOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetWasmCodesCodeIDOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetWasmCodesCodeIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetWasmCodesCodeIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package dashboard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/lidofinance/terra-fcd-rest-client/columbus-5/models"
)

// GetV1DashboardRegisteredAccountsReader is a Reader for the GetV1DashboardRegisteredAccounts structure.
type GetV1DashboardRegisteredAccountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1DashboardRegisteredAccountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1DashboardRegisteredAccountsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1DashboardRegisteredAccountsOK creates a GetV1DashboardRegisteredAccountsOK with default headers values
func NewGetV1DashboardRegisteredAccountsOK() *GetV1DashboardRegisteredAccountsOK {
	return &GetV1DashboardRegisteredAccountsOK{}
}

/* GetV1DashboardRegisteredAccountsOK describes a response with status code 200, with default header values.

Success
*/
type GetV1DashboardRegisteredAccountsOK struct {
	Payload *models.GetRegisteredAccountsResult
}

func (o *GetV1DashboardRegisteredAccountsOK) Error() string {
	return fmt.Sprintf("[GET /v1/dashboard/registered_accounts][%d] getV1DashboardRegisteredAccountsOK  %+v", 200, o.Payload)
}
func (o *GetV1DashboardRegisteredAccountsOK) GetPayload() *models.GetRegisteredAccountsResult {
	return o.Payload
}

func (o *GetV1DashboardRegisteredAccountsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetRegisteredAccountsResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

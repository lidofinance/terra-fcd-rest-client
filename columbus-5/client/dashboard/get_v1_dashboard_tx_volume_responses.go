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

// GetV1DashboardTxVolumeReader is a Reader for the GetV1DashboardTxVolume structure.
type GetV1DashboardTxVolumeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1DashboardTxVolumeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1DashboardTxVolumeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1DashboardTxVolumeOK creates a GetV1DashboardTxVolumeOK with default headers values
func NewGetV1DashboardTxVolumeOK() *GetV1DashboardTxVolumeOK {
	return &GetV1DashboardTxVolumeOK{}
}

/* GetV1DashboardTxVolumeOK describes a response with status code 200, with default header values.

Success
*/
type GetV1DashboardTxVolumeOK struct {
	Payload *models.GetTxVolumeResult
}

func (o *GetV1DashboardTxVolumeOK) Error() string {
	return fmt.Sprintf("[GET /v1/dashboard/tx_volume][%d] getV1DashboardTxVolumeOK  %+v", 200, o.Payload)
}
func (o *GetV1DashboardTxVolumeOK) GetPayload() *models.GetTxVolumeResult {
	return o.Payload
}

func (o *GetV1DashboardTxVolumeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetTxVolumeResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

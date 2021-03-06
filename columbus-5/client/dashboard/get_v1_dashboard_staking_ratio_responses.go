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

// GetV1DashboardStakingRatioReader is a Reader for the GetV1DashboardStakingRatio structure.
type GetV1DashboardStakingRatioReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1DashboardStakingRatioReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1DashboardStakingRatioOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1DashboardStakingRatioOK creates a GetV1DashboardStakingRatioOK with default headers values
func NewGetV1DashboardStakingRatioOK() *GetV1DashboardStakingRatioOK {
	return &GetV1DashboardStakingRatioOK{}
}

/* GetV1DashboardStakingRatioOK describes a response with status code 200, with default header values.

Success
*/
type GetV1DashboardStakingRatioOK struct {
	Payload []*models.StakingHistory
}

func (o *GetV1DashboardStakingRatioOK) Error() string {
	return fmt.Sprintf("[GET /v1/dashboard/staking_ratio][%d] getV1DashboardStakingRatioOK  %+v", 200, o.Payload)
}
func (o *GetV1DashboardStakingRatioOK) GetPayload() []*models.StakingHistory {
	return o.Payload
}

func (o *GetV1DashboardStakingRatioOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

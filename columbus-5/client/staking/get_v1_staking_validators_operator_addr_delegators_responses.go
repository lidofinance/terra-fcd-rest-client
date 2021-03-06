// Code generated by go-swagger; DO NOT EDIT.

package staking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/lidofinance/terra-fcd-rest-client/columbus-5/models"
)

// GetV1StakingValidatorsOperatorAddrDelegatorsReader is a Reader for the GetV1StakingValidatorsOperatorAddrDelegators structure.
type GetV1StakingValidatorsOperatorAddrDelegatorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV1StakingValidatorsOperatorAddrDelegatorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetV1StakingValidatorsOperatorAddrDelegatorsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetV1StakingValidatorsOperatorAddrDelegatorsOK creates a GetV1StakingValidatorsOperatorAddrDelegatorsOK with default headers values
func NewGetV1StakingValidatorsOperatorAddrDelegatorsOK() *GetV1StakingValidatorsOperatorAddrDelegatorsOK {
	return &GetV1StakingValidatorsOperatorAddrDelegatorsOK{}
}

/* GetV1StakingValidatorsOperatorAddrDelegatorsOK describes a response with status code 200, with default header values.

Success
*/
type GetV1StakingValidatorsOperatorAddrDelegatorsOK struct {
	Payload *models.GetValidatorDelegatorsResult
}

func (o *GetV1StakingValidatorsOperatorAddrDelegatorsOK) Error() string {
	return fmt.Sprintf("[GET /v1/staking/validators/{operatorAddr}/delegators][%d] getV1StakingValidatorsOperatorAddrDelegatorsOK  %+v", 200, o.Payload)
}
func (o *GetV1StakingValidatorsOperatorAddrDelegatorsOK) GetPayload() *models.GetValidatorDelegatorsResult {
	return o.Payload
}

func (o *GetV1StakingValidatorsOperatorAddrDelegatorsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetValidatorDelegatorsResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

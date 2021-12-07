// Code generated by go-swagger; DO NOT EDIT.

package staking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetV1StakingValidatorsParams creates a new GetV1StakingValidatorsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1StakingValidatorsParams() *GetV1StakingValidatorsParams {
	return &GetV1StakingValidatorsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1StakingValidatorsParamsWithTimeout creates a new GetV1StakingValidatorsParams object
// with the ability to set a timeout on a request.
func NewGetV1StakingValidatorsParamsWithTimeout(timeout time.Duration) *GetV1StakingValidatorsParams {
	return &GetV1StakingValidatorsParams{
		timeout: timeout,
	}
}

// NewGetV1StakingValidatorsParamsWithContext creates a new GetV1StakingValidatorsParams object
// with the ability to set a context for a request.
func NewGetV1StakingValidatorsParamsWithContext(ctx context.Context) *GetV1StakingValidatorsParams {
	return &GetV1StakingValidatorsParams{
		Context: ctx,
	}
}

// NewGetV1StakingValidatorsParamsWithHTTPClient creates a new GetV1StakingValidatorsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1StakingValidatorsParamsWithHTTPClient(client *http.Client) *GetV1StakingValidatorsParams {
	return &GetV1StakingValidatorsParams{
		HTTPClient: client,
	}
}

/* GetV1StakingValidatorsParams contains all the parameters to send to the API endpoint
   for the get v1 staking validators operation.

   Typically these are written to a http.Request.
*/
type GetV1StakingValidatorsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 staking validators params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1StakingValidatorsParams) WithDefaults() *GetV1StakingValidatorsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 staking validators params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1StakingValidatorsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 staking validators params
func (o *GetV1StakingValidatorsParams) WithTimeout(timeout time.Duration) *GetV1StakingValidatorsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 staking validators params
func (o *GetV1StakingValidatorsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 staking validators params
func (o *GetV1StakingValidatorsParams) WithContext(ctx context.Context) *GetV1StakingValidatorsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 staking validators params
func (o *GetV1StakingValidatorsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 staking validators params
func (o *GetV1StakingValidatorsParams) WithHTTPClient(client *http.Client) *GetV1StakingValidatorsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 staking validators params
func (o *GetV1StakingValidatorsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1StakingValidatorsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
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

// NewGetStakingParametersParams creates a new GetStakingParametersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetStakingParametersParams() *GetStakingParametersParams {
	return &GetStakingParametersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetStakingParametersParamsWithTimeout creates a new GetStakingParametersParams object
// with the ability to set a timeout on a request.
func NewGetStakingParametersParamsWithTimeout(timeout time.Duration) *GetStakingParametersParams {
	return &GetStakingParametersParams{
		timeout: timeout,
	}
}

// NewGetStakingParametersParamsWithContext creates a new GetStakingParametersParams object
// with the ability to set a context for a request.
func NewGetStakingParametersParamsWithContext(ctx context.Context) *GetStakingParametersParams {
	return &GetStakingParametersParams{
		Context: ctx,
	}
}

// NewGetStakingParametersParamsWithHTTPClient creates a new GetStakingParametersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetStakingParametersParamsWithHTTPClient(client *http.Client) *GetStakingParametersParams {
	return &GetStakingParametersParams{
		HTTPClient: client,
	}
}

/* GetStakingParametersParams contains all the parameters to send to the API endpoint
   for the get staking parameters operation.

   Typically these are written to a http.Request.
*/
type GetStakingParametersParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get staking parameters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStakingParametersParams) WithDefaults() *GetStakingParametersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get staking parameters params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStakingParametersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get staking parameters params
func (o *GetStakingParametersParams) WithTimeout(timeout time.Duration) *GetStakingParametersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get staking parameters params
func (o *GetStakingParametersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get staking parameters params
func (o *GetStakingParametersParams) WithContext(ctx context.Context) *GetStakingParametersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get staking parameters params
func (o *GetStakingParametersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get staking parameters params
func (o *GetStakingParametersParams) WithHTTPClient(client *http.Client) *GetStakingParametersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get staking parameters params
func (o *GetStakingParametersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetStakingParametersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package dashboard

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

// NewGetV1DashboardStakingReturnParams creates a new GetV1DashboardStakingReturnParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1DashboardStakingReturnParams() *GetV1DashboardStakingReturnParams {
	return &GetV1DashboardStakingReturnParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1DashboardStakingReturnParamsWithTimeout creates a new GetV1DashboardStakingReturnParams object
// with the ability to set a timeout on a request.
func NewGetV1DashboardStakingReturnParamsWithTimeout(timeout time.Duration) *GetV1DashboardStakingReturnParams {
	return &GetV1DashboardStakingReturnParams{
		timeout: timeout,
	}
}

// NewGetV1DashboardStakingReturnParamsWithContext creates a new GetV1DashboardStakingReturnParams object
// with the ability to set a context for a request.
func NewGetV1DashboardStakingReturnParamsWithContext(ctx context.Context) *GetV1DashboardStakingReturnParams {
	return &GetV1DashboardStakingReturnParams{
		Context: ctx,
	}
}

// NewGetV1DashboardStakingReturnParamsWithHTTPClient creates a new GetV1DashboardStakingReturnParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1DashboardStakingReturnParamsWithHTTPClient(client *http.Client) *GetV1DashboardStakingReturnParams {
	return &GetV1DashboardStakingReturnParams{
		HTTPClient: client,
	}
}

/* GetV1DashboardStakingReturnParams contains all the parameters to send to the API endpoint
   for the get v1 dashboard staking return operation.

   Typically these are written to a http.Request.
*/
type GetV1DashboardStakingReturnParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 dashboard staking return params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1DashboardStakingReturnParams) WithDefaults() *GetV1DashboardStakingReturnParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 dashboard staking return params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1DashboardStakingReturnParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 dashboard staking return params
func (o *GetV1DashboardStakingReturnParams) WithTimeout(timeout time.Duration) *GetV1DashboardStakingReturnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 dashboard staking return params
func (o *GetV1DashboardStakingReturnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 dashboard staking return params
func (o *GetV1DashboardStakingReturnParams) WithContext(ctx context.Context) *GetV1DashboardStakingReturnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 dashboard staking return params
func (o *GetV1DashboardStakingReturnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 dashboard staking return params
func (o *GetV1DashboardStakingReturnParams) WithHTTPClient(client *http.Client) *GetV1DashboardStakingReturnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 dashboard staking return params
func (o *GetV1DashboardStakingReturnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1DashboardStakingReturnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package governance

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

// NewGetGovParametersVotingParams creates a new GetGovParametersVotingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetGovParametersVotingParams() *GetGovParametersVotingParams {
	return &GetGovParametersVotingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetGovParametersVotingParamsWithTimeout creates a new GetGovParametersVotingParams object
// with the ability to set a timeout on a request.
func NewGetGovParametersVotingParamsWithTimeout(timeout time.Duration) *GetGovParametersVotingParams {
	return &GetGovParametersVotingParams{
		timeout: timeout,
	}
}

// NewGetGovParametersVotingParamsWithContext creates a new GetGovParametersVotingParams object
// with the ability to set a context for a request.
func NewGetGovParametersVotingParamsWithContext(ctx context.Context) *GetGovParametersVotingParams {
	return &GetGovParametersVotingParams{
		Context: ctx,
	}
}

// NewGetGovParametersVotingParamsWithHTTPClient creates a new GetGovParametersVotingParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetGovParametersVotingParamsWithHTTPClient(client *http.Client) *GetGovParametersVotingParams {
	return &GetGovParametersVotingParams{
		HTTPClient: client,
	}
}

/* GetGovParametersVotingParams contains all the parameters to send to the API endpoint
   for the get gov parameters voting operation.

   Typically these are written to a http.Request.
*/
type GetGovParametersVotingParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get gov parameters voting params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGovParametersVotingParams) WithDefaults() *GetGovParametersVotingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get gov parameters voting params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetGovParametersVotingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get gov parameters voting params
func (o *GetGovParametersVotingParams) WithTimeout(timeout time.Duration) *GetGovParametersVotingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get gov parameters voting params
func (o *GetGovParametersVotingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get gov parameters voting params
func (o *GetGovParametersVotingParams) WithContext(ctx context.Context) *GetGovParametersVotingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get gov parameters voting params
func (o *GetGovParametersVotingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get gov parameters voting params
func (o *GetGovParametersVotingParams) WithHTTPClient(client *http.Client) *GetGovParametersVotingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get gov parameters voting params
func (o *GetGovParametersVotingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetGovParametersVotingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

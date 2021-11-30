// Code generated by go-swagger; DO NOT EDIT.

package query

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

// NewAggregateVotesParams creates a new AggregateVotesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAggregateVotesParams() *AggregateVotesParams {
	return &AggregateVotesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAggregateVotesParamsWithTimeout creates a new AggregateVotesParams object
// with the ability to set a timeout on a request.
func NewAggregateVotesParamsWithTimeout(timeout time.Duration) *AggregateVotesParams {
	return &AggregateVotesParams{
		timeout: timeout,
	}
}

// NewAggregateVotesParamsWithContext creates a new AggregateVotesParams object
// with the ability to set a context for a request.
func NewAggregateVotesParamsWithContext(ctx context.Context) *AggregateVotesParams {
	return &AggregateVotesParams{
		Context: ctx,
	}
}

// NewAggregateVotesParamsWithHTTPClient creates a new AggregateVotesParams object
// with the ability to set a custom HTTPClient for a request.
func NewAggregateVotesParamsWithHTTPClient(client *http.Client) *AggregateVotesParams {
	return &AggregateVotesParams{
		HTTPClient: client,
	}
}

/* AggregateVotesParams contains all the parameters to send to the API endpoint
   for the aggregate votes operation.

   Typically these are written to a http.Request.
*/
type AggregateVotesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the aggregate votes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AggregateVotesParams) WithDefaults() *AggregateVotesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the aggregate votes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AggregateVotesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the aggregate votes params
func (o *AggregateVotesParams) WithTimeout(timeout time.Duration) *AggregateVotesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the aggregate votes params
func (o *AggregateVotesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the aggregate votes params
func (o *AggregateVotesParams) WithContext(ctx context.Context) *AggregateVotesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the aggregate votes params
func (o *AggregateVotesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the aggregate votes params
func (o *AggregateVotesParams) WithHTTPClient(client *http.Client) *AggregateVotesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the aggregate votes params
func (o *AggregateVotesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *AggregateVotesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
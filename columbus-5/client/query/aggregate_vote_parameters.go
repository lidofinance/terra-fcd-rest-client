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

// NewAggregateVoteParams creates a new AggregateVoteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAggregateVoteParams() *AggregateVoteParams {
	return &AggregateVoteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAggregateVoteParamsWithTimeout creates a new AggregateVoteParams object
// with the ability to set a timeout on a request.
func NewAggregateVoteParamsWithTimeout(timeout time.Duration) *AggregateVoteParams {
	return &AggregateVoteParams{
		timeout: timeout,
	}
}

// NewAggregateVoteParamsWithContext creates a new AggregateVoteParams object
// with the ability to set a context for a request.
func NewAggregateVoteParamsWithContext(ctx context.Context) *AggregateVoteParams {
	return &AggregateVoteParams{
		Context: ctx,
	}
}

// NewAggregateVoteParamsWithHTTPClient creates a new AggregateVoteParams object
// with the ability to set a custom HTTPClient for a request.
func NewAggregateVoteParamsWithHTTPClient(client *http.Client) *AggregateVoteParams {
	return &AggregateVoteParams{
		HTTPClient: client,
	}
}

/* AggregateVoteParams contains all the parameters to send to the API endpoint
   for the aggregate vote operation.

   Typically these are written to a http.Request.
*/
type AggregateVoteParams struct {

	/* ValidatorAddr.

	   validator defines the validator address to query for.
	*/
	ValidatorAddr string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the aggregate vote params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AggregateVoteParams) WithDefaults() *AggregateVoteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the aggregate vote params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AggregateVoteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the aggregate vote params
func (o *AggregateVoteParams) WithTimeout(timeout time.Duration) *AggregateVoteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the aggregate vote params
func (o *AggregateVoteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the aggregate vote params
func (o *AggregateVoteParams) WithContext(ctx context.Context) *AggregateVoteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the aggregate vote params
func (o *AggregateVoteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the aggregate vote params
func (o *AggregateVoteParams) WithHTTPClient(client *http.Client) *AggregateVoteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the aggregate vote params
func (o *AggregateVoteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithValidatorAddr adds the validatorAddr to the aggregate vote params
func (o *AggregateVoteParams) WithValidatorAddr(validatorAddr string) *AggregateVoteParams {
	o.SetValidatorAddr(validatorAddr)
	return o
}

// SetValidatorAddr adds the validatorAddr to the aggregate vote params
func (o *AggregateVoteParams) SetValidatorAddr(validatorAddr string) {
	o.ValidatorAddr = validatorAddr
}

// WriteToRequest writes these params to a swagger request
func (o *AggregateVoteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param validator_addr
	if err := r.SetPathParam("validator_addr", o.ValidatorAddr); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

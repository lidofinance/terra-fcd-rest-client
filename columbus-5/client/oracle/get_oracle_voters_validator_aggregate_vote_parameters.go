// Code generated by go-swagger; DO NOT EDIT.

package oracle

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

// NewGetOracleVotersValidatorAggregateVoteParams creates a new GetOracleVotersValidatorAggregateVoteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOracleVotersValidatorAggregateVoteParams() *GetOracleVotersValidatorAggregateVoteParams {
	return &GetOracleVotersValidatorAggregateVoteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOracleVotersValidatorAggregateVoteParamsWithTimeout creates a new GetOracleVotersValidatorAggregateVoteParams object
// with the ability to set a timeout on a request.
func NewGetOracleVotersValidatorAggregateVoteParamsWithTimeout(timeout time.Duration) *GetOracleVotersValidatorAggregateVoteParams {
	return &GetOracleVotersValidatorAggregateVoteParams{
		timeout: timeout,
	}
}

// NewGetOracleVotersValidatorAggregateVoteParamsWithContext creates a new GetOracleVotersValidatorAggregateVoteParams object
// with the ability to set a context for a request.
func NewGetOracleVotersValidatorAggregateVoteParamsWithContext(ctx context.Context) *GetOracleVotersValidatorAggregateVoteParams {
	return &GetOracleVotersValidatorAggregateVoteParams{
		Context: ctx,
	}
}

// NewGetOracleVotersValidatorAggregateVoteParamsWithHTTPClient creates a new GetOracleVotersValidatorAggregateVoteParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOracleVotersValidatorAggregateVoteParamsWithHTTPClient(client *http.Client) *GetOracleVotersValidatorAggregateVoteParams {
	return &GetOracleVotersValidatorAggregateVoteParams{
		HTTPClient: client,
	}
}

/* GetOracleVotersValidatorAggregateVoteParams contains all the parameters to send to the API endpoint
   for the get oracle voters validator aggregate vote operation.

   Typically these are written to a http.Request.
*/
type GetOracleVotersValidatorAggregateVoteParams struct {

	/* Validator.

	   oracle operator
	*/
	Validator string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get oracle voters validator aggregate vote params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOracleVotersValidatorAggregateVoteParams) WithDefaults() *GetOracleVotersValidatorAggregateVoteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get oracle voters validator aggregate vote params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOracleVotersValidatorAggregateVoteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get oracle voters validator aggregate vote params
func (o *GetOracleVotersValidatorAggregateVoteParams) WithTimeout(timeout time.Duration) *GetOracleVotersValidatorAggregateVoteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get oracle voters validator aggregate vote params
func (o *GetOracleVotersValidatorAggregateVoteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get oracle voters validator aggregate vote params
func (o *GetOracleVotersValidatorAggregateVoteParams) WithContext(ctx context.Context) *GetOracleVotersValidatorAggregateVoteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get oracle voters validator aggregate vote params
func (o *GetOracleVotersValidatorAggregateVoteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get oracle voters validator aggregate vote params
func (o *GetOracleVotersValidatorAggregateVoteParams) WithHTTPClient(client *http.Client) *GetOracleVotersValidatorAggregateVoteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get oracle voters validator aggregate vote params
func (o *GetOracleVotersValidatorAggregateVoteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithValidator adds the validator to the get oracle voters validator aggregate vote params
func (o *GetOracleVotersValidatorAggregateVoteParams) WithValidator(validator string) *GetOracleVotersValidatorAggregateVoteParams {
	o.SetValidator(validator)
	return o
}

// SetValidator adds the validator to the get oracle voters validator aggregate vote params
func (o *GetOracleVotersValidatorAggregateVoteParams) SetValidator(validator string) {
	o.Validator = validator
}

// WriteToRequest writes these params to a swagger request
func (o *GetOracleVotersValidatorAggregateVoteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param validator
	if err := r.SetPathParam("validator", o.Validator); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
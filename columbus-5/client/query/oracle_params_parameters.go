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

// NewOracleParamsParams creates a new OracleParamsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewOracleParamsParams() *OracleParamsParams {
	return &OracleParamsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewOracleParamsParamsWithTimeout creates a new OracleParamsParams object
// with the ability to set a timeout on a request.
func NewOracleParamsParamsWithTimeout(timeout time.Duration) *OracleParamsParams {
	return &OracleParamsParams{
		timeout: timeout,
	}
}

// NewOracleParamsParamsWithContext creates a new OracleParamsParams object
// with the ability to set a context for a request.
func NewOracleParamsParamsWithContext(ctx context.Context) *OracleParamsParams {
	return &OracleParamsParams{
		Context: ctx,
	}
}

// NewOracleParamsParamsWithHTTPClient creates a new OracleParamsParams object
// with the ability to set a custom HTTPClient for a request.
func NewOracleParamsParamsWithHTTPClient(client *http.Client) *OracleParamsParams {
	return &OracleParamsParams{
		HTTPClient: client,
	}
}

/* OracleParamsParams contains all the parameters to send to the API endpoint
   for the oracle params operation.

   Typically these are written to a http.Request.
*/
type OracleParamsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the oracle params params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OracleParamsParams) WithDefaults() *OracleParamsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the oracle params params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *OracleParamsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the oracle params params
func (o *OracleParamsParams) WithTimeout(timeout time.Duration) *OracleParamsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the oracle params params
func (o *OracleParamsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the oracle params params
func (o *OracleParamsParams) WithContext(ctx context.Context) *OracleParamsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the oracle params params
func (o *OracleParamsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the oracle params params
func (o *OracleParamsParams) WithHTTPClient(client *http.Client) *OracleParamsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the oracle params params
func (o *OracleParamsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *OracleParamsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

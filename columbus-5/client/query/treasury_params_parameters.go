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

// NewTreasuryParamsParams creates a new TreasuryParamsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTreasuryParamsParams() *TreasuryParamsParams {
	return &TreasuryParamsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTreasuryParamsParamsWithTimeout creates a new TreasuryParamsParams object
// with the ability to set a timeout on a request.
func NewTreasuryParamsParamsWithTimeout(timeout time.Duration) *TreasuryParamsParams {
	return &TreasuryParamsParams{
		timeout: timeout,
	}
}

// NewTreasuryParamsParamsWithContext creates a new TreasuryParamsParams object
// with the ability to set a context for a request.
func NewTreasuryParamsParamsWithContext(ctx context.Context) *TreasuryParamsParams {
	return &TreasuryParamsParams{
		Context: ctx,
	}
}

// NewTreasuryParamsParamsWithHTTPClient creates a new TreasuryParamsParams object
// with the ability to set a custom HTTPClient for a request.
func NewTreasuryParamsParamsWithHTTPClient(client *http.Client) *TreasuryParamsParams {
	return &TreasuryParamsParams{
		HTTPClient: client,
	}
}

/* TreasuryParamsParams contains all the parameters to send to the API endpoint
   for the treasury params operation.

   Typically these are written to a http.Request.
*/
type TreasuryParamsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the treasury params params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TreasuryParamsParams) WithDefaults() *TreasuryParamsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the treasury params params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TreasuryParamsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the treasury params params
func (o *TreasuryParamsParams) WithTimeout(timeout time.Duration) *TreasuryParamsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the treasury params params
func (o *TreasuryParamsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the treasury params params
func (o *TreasuryParamsParams) WithContext(ctx context.Context) *TreasuryParamsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the treasury params params
func (o *TreasuryParamsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the treasury params params
func (o *TreasuryParamsParams) WithHTTPClient(client *http.Client) *TreasuryParamsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the treasury params params
func (o *TreasuryParamsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *TreasuryParamsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

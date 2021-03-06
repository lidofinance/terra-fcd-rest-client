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

// NewExchangeRateParams creates a new ExchangeRateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExchangeRateParams() *ExchangeRateParams {
	return &ExchangeRateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExchangeRateParamsWithTimeout creates a new ExchangeRateParams object
// with the ability to set a timeout on a request.
func NewExchangeRateParamsWithTimeout(timeout time.Duration) *ExchangeRateParams {
	return &ExchangeRateParams{
		timeout: timeout,
	}
}

// NewExchangeRateParamsWithContext creates a new ExchangeRateParams object
// with the ability to set a context for a request.
func NewExchangeRateParamsWithContext(ctx context.Context) *ExchangeRateParams {
	return &ExchangeRateParams{
		Context: ctx,
	}
}

// NewExchangeRateParamsWithHTTPClient creates a new ExchangeRateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExchangeRateParamsWithHTTPClient(client *http.Client) *ExchangeRateParams {
	return &ExchangeRateParams{
		HTTPClient: client,
	}
}

/* ExchangeRateParams contains all the parameters to send to the API endpoint
   for the exchange rate operation.

   Typically these are written to a http.Request.
*/
type ExchangeRateParams struct {

	/* Denom.

	   denom defines the denomination to query for.
	*/
	Denom string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the exchange rate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExchangeRateParams) WithDefaults() *ExchangeRateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the exchange rate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExchangeRateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the exchange rate params
func (o *ExchangeRateParams) WithTimeout(timeout time.Duration) *ExchangeRateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the exchange rate params
func (o *ExchangeRateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the exchange rate params
func (o *ExchangeRateParams) WithContext(ctx context.Context) *ExchangeRateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the exchange rate params
func (o *ExchangeRateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the exchange rate params
func (o *ExchangeRateParams) WithHTTPClient(client *http.Client) *ExchangeRateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the exchange rate params
func (o *ExchangeRateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDenom adds the denom to the exchange rate params
func (o *ExchangeRateParams) WithDenom(denom string) *ExchangeRateParams {
	o.SetDenom(denom)
	return o
}

// SetDenom adds the denom to the exchange rate params
func (o *ExchangeRateParams) SetDenom(denom string) {
	o.Denom = denom
}

// WriteToRequest writes these params to a swagger request
func (o *ExchangeRateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param denom
	if err := r.SetPathParam("denom", o.Denom); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

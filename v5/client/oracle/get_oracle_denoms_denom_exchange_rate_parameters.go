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

// NewGetOracleDenomsDenomExchangeRateParams creates a new GetOracleDenomsDenomExchangeRateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOracleDenomsDenomExchangeRateParams() *GetOracleDenomsDenomExchangeRateParams {
	return &GetOracleDenomsDenomExchangeRateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOracleDenomsDenomExchangeRateParamsWithTimeout creates a new GetOracleDenomsDenomExchangeRateParams object
// with the ability to set a timeout on a request.
func NewGetOracleDenomsDenomExchangeRateParamsWithTimeout(timeout time.Duration) *GetOracleDenomsDenomExchangeRateParams {
	return &GetOracleDenomsDenomExchangeRateParams{
		timeout: timeout,
	}
}

// NewGetOracleDenomsDenomExchangeRateParamsWithContext creates a new GetOracleDenomsDenomExchangeRateParams object
// with the ability to set a context for a request.
func NewGetOracleDenomsDenomExchangeRateParamsWithContext(ctx context.Context) *GetOracleDenomsDenomExchangeRateParams {
	return &GetOracleDenomsDenomExchangeRateParams{
		Context: ctx,
	}
}

// NewGetOracleDenomsDenomExchangeRateParamsWithHTTPClient creates a new GetOracleDenomsDenomExchangeRateParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOracleDenomsDenomExchangeRateParamsWithHTTPClient(client *http.Client) *GetOracleDenomsDenomExchangeRateParams {
	return &GetOracleDenomsDenomExchangeRateParams{
		HTTPClient: client,
	}
}

/* GetOracleDenomsDenomExchangeRateParams contains all the parameters to send to the API endpoint
   for the get oracle denoms denom exchange rate operation.

   Typically these are written to a http.Request.
*/
type GetOracleDenomsDenomExchangeRateParams struct {

	/* Denom.

	   The coin denom to get
	*/
	Denom string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get oracle denoms denom exchange rate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOracleDenomsDenomExchangeRateParams) WithDefaults() *GetOracleDenomsDenomExchangeRateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get oracle denoms denom exchange rate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOracleDenomsDenomExchangeRateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get oracle denoms denom exchange rate params
func (o *GetOracleDenomsDenomExchangeRateParams) WithTimeout(timeout time.Duration) *GetOracleDenomsDenomExchangeRateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get oracle denoms denom exchange rate params
func (o *GetOracleDenomsDenomExchangeRateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get oracle denoms denom exchange rate params
func (o *GetOracleDenomsDenomExchangeRateParams) WithContext(ctx context.Context) *GetOracleDenomsDenomExchangeRateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get oracle denoms denom exchange rate params
func (o *GetOracleDenomsDenomExchangeRateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get oracle denoms denom exchange rate params
func (o *GetOracleDenomsDenomExchangeRateParams) WithHTTPClient(client *http.Client) *GetOracleDenomsDenomExchangeRateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get oracle denoms denom exchange rate params
func (o *GetOracleDenomsDenomExchangeRateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDenom adds the denom to the get oracle denoms denom exchange rate params
func (o *GetOracleDenomsDenomExchangeRateParams) WithDenom(denom string) *GetOracleDenomsDenomExchangeRateParams {
	o.SetDenom(denom)
	return o
}

// SetDenom adds the denom to the get oracle denoms denom exchange rate params
func (o *GetOracleDenomsDenomExchangeRateParams) SetDenom(denom string) {
	o.Denom = denom
}

// WriteToRequest writes these params to a swagger request
func (o *GetOracleDenomsDenomExchangeRateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
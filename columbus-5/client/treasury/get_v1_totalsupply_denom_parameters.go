// Code generated by go-swagger; DO NOT EDIT.

package treasury

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

// NewGetV1TotalsupplyDenomParams creates a new GetV1TotalsupplyDenomParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1TotalsupplyDenomParams() *GetV1TotalsupplyDenomParams {
	return &GetV1TotalsupplyDenomParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1TotalsupplyDenomParamsWithTimeout creates a new GetV1TotalsupplyDenomParams object
// with the ability to set a timeout on a request.
func NewGetV1TotalsupplyDenomParamsWithTimeout(timeout time.Duration) *GetV1TotalsupplyDenomParams {
	return &GetV1TotalsupplyDenomParams{
		timeout: timeout,
	}
}

// NewGetV1TotalsupplyDenomParamsWithContext creates a new GetV1TotalsupplyDenomParams object
// with the ability to set a context for a request.
func NewGetV1TotalsupplyDenomParamsWithContext(ctx context.Context) *GetV1TotalsupplyDenomParams {
	return &GetV1TotalsupplyDenomParams{
		Context: ctx,
	}
}

// NewGetV1TotalsupplyDenomParamsWithHTTPClient creates a new GetV1TotalsupplyDenomParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1TotalsupplyDenomParamsWithHTTPClient(client *http.Client) *GetV1TotalsupplyDenomParams {
	return &GetV1TotalsupplyDenomParams{
		HTTPClient: client,
	}
}

/* GetV1TotalsupplyDenomParams contains all the parameters to send to the API endpoint
   for the get v1 totalsupply denom operation.

   Typically these are written to a http.Request.
*/
type GetV1TotalsupplyDenomParams struct {

	/* Denom.

	   Coin denomination
	*/
	Denom string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 totalsupply denom params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1TotalsupplyDenomParams) WithDefaults() *GetV1TotalsupplyDenomParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 totalsupply denom params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1TotalsupplyDenomParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 totalsupply denom params
func (o *GetV1TotalsupplyDenomParams) WithTimeout(timeout time.Duration) *GetV1TotalsupplyDenomParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 totalsupply denom params
func (o *GetV1TotalsupplyDenomParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 totalsupply denom params
func (o *GetV1TotalsupplyDenomParams) WithContext(ctx context.Context) *GetV1TotalsupplyDenomParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 totalsupply denom params
func (o *GetV1TotalsupplyDenomParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 totalsupply denom params
func (o *GetV1TotalsupplyDenomParams) WithHTTPClient(client *http.Client) *GetV1TotalsupplyDenomParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 totalsupply denom params
func (o *GetV1TotalsupplyDenomParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDenom adds the denom to the get v1 totalsupply denom params
func (o *GetV1TotalsupplyDenomParams) WithDenom(denom string) *GetV1TotalsupplyDenomParams {
	o.SetDenom(denom)
	return o
}

// SetDenom adds the denom to the get v1 totalsupply denom params
func (o *GetV1TotalsupplyDenomParams) SetDenom(denom string) {
	o.Denom = denom
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1TotalsupplyDenomParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

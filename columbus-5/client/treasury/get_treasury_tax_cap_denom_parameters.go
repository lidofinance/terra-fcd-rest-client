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

// NewGetTreasuryTaxCapDenomParams creates a new GetTreasuryTaxCapDenomParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTreasuryTaxCapDenomParams() *GetTreasuryTaxCapDenomParams {
	return &GetTreasuryTaxCapDenomParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTreasuryTaxCapDenomParamsWithTimeout creates a new GetTreasuryTaxCapDenomParams object
// with the ability to set a timeout on a request.
func NewGetTreasuryTaxCapDenomParamsWithTimeout(timeout time.Duration) *GetTreasuryTaxCapDenomParams {
	return &GetTreasuryTaxCapDenomParams{
		timeout: timeout,
	}
}

// NewGetTreasuryTaxCapDenomParamsWithContext creates a new GetTreasuryTaxCapDenomParams object
// with the ability to set a context for a request.
func NewGetTreasuryTaxCapDenomParamsWithContext(ctx context.Context) *GetTreasuryTaxCapDenomParams {
	return &GetTreasuryTaxCapDenomParams{
		Context: ctx,
	}
}

// NewGetTreasuryTaxCapDenomParamsWithHTTPClient creates a new GetTreasuryTaxCapDenomParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTreasuryTaxCapDenomParamsWithHTTPClient(client *http.Client) *GetTreasuryTaxCapDenomParams {
	return &GetTreasuryTaxCapDenomParams{
		HTTPClient: client,
	}
}

/* GetTreasuryTaxCapDenomParams contains all the parameters to send to the API endpoint
   for the get treasury tax cap denom operation.

   Typically these are written to a http.Request.
*/
type GetTreasuryTaxCapDenomParams struct {

	/* Denom.

	   Denom
	*/
	Denom string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get treasury tax cap denom params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTreasuryTaxCapDenomParams) WithDefaults() *GetTreasuryTaxCapDenomParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get treasury tax cap denom params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTreasuryTaxCapDenomParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get treasury tax cap denom params
func (o *GetTreasuryTaxCapDenomParams) WithTimeout(timeout time.Duration) *GetTreasuryTaxCapDenomParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get treasury tax cap denom params
func (o *GetTreasuryTaxCapDenomParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get treasury tax cap denom params
func (o *GetTreasuryTaxCapDenomParams) WithContext(ctx context.Context) *GetTreasuryTaxCapDenomParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get treasury tax cap denom params
func (o *GetTreasuryTaxCapDenomParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get treasury tax cap denom params
func (o *GetTreasuryTaxCapDenomParams) WithHTTPClient(client *http.Client) *GetTreasuryTaxCapDenomParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get treasury tax cap denom params
func (o *GetTreasuryTaxCapDenomParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDenom adds the denom to the get treasury tax cap denom params
func (o *GetTreasuryTaxCapDenomParams) WithDenom(denom string) *GetTreasuryTaxCapDenomParams {
	o.SetDenom(denom)
	return o
}

// SetDenom adds the denom to the get treasury tax cap denom params
func (o *GetTreasuryTaxCapDenomParams) SetDenom(denom string) {
	o.Denom = denom
}

// WriteToRequest writes these params to a swagger request
func (o *GetTreasuryTaxCapDenomParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
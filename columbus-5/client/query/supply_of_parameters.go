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

// NewSupplyOfParams creates a new SupplyOfParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSupplyOfParams() *SupplyOfParams {
	return &SupplyOfParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSupplyOfParamsWithTimeout creates a new SupplyOfParams object
// with the ability to set a timeout on a request.
func NewSupplyOfParamsWithTimeout(timeout time.Duration) *SupplyOfParams {
	return &SupplyOfParams{
		timeout: timeout,
	}
}

// NewSupplyOfParamsWithContext creates a new SupplyOfParams object
// with the ability to set a context for a request.
func NewSupplyOfParamsWithContext(ctx context.Context) *SupplyOfParams {
	return &SupplyOfParams{
		Context: ctx,
	}
}

// NewSupplyOfParamsWithHTTPClient creates a new SupplyOfParams object
// with the ability to set a custom HTTPClient for a request.
func NewSupplyOfParamsWithHTTPClient(client *http.Client) *SupplyOfParams {
	return &SupplyOfParams{
		HTTPClient: client,
	}
}

/* SupplyOfParams contains all the parameters to send to the API endpoint
   for the supply of operation.

   Typically these are written to a http.Request.
*/
type SupplyOfParams struct {

	/* Denom.

	   denom is the coin denom to query balances for.
	*/
	Denom string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the supply of params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SupplyOfParams) WithDefaults() *SupplyOfParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the supply of params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SupplyOfParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the supply of params
func (o *SupplyOfParams) WithTimeout(timeout time.Duration) *SupplyOfParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the supply of params
func (o *SupplyOfParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the supply of params
func (o *SupplyOfParams) WithContext(ctx context.Context) *SupplyOfParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the supply of params
func (o *SupplyOfParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the supply of params
func (o *SupplyOfParams) WithHTTPClient(client *http.Client) *SupplyOfParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the supply of params
func (o *SupplyOfParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDenom adds the denom to the supply of params
func (o *SupplyOfParams) WithDenom(denom string) *SupplyOfParams {
	o.SetDenom(denom)
	return o
}

// SetDenom adds the denom to the supply of params
func (o *SupplyOfParams) SetDenom(denom string) {
	o.Denom = denom
}

// WriteToRequest writes these params to a swagger request
func (o *SupplyOfParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

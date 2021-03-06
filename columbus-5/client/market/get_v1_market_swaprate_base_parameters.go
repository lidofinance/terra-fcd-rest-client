// Code generated by go-swagger; DO NOT EDIT.

package market

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

// NewGetV1MarketSwaprateBaseParams creates a new GetV1MarketSwaprateBaseParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1MarketSwaprateBaseParams() *GetV1MarketSwaprateBaseParams {
	return &GetV1MarketSwaprateBaseParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1MarketSwaprateBaseParamsWithTimeout creates a new GetV1MarketSwaprateBaseParams object
// with the ability to set a timeout on a request.
func NewGetV1MarketSwaprateBaseParamsWithTimeout(timeout time.Duration) *GetV1MarketSwaprateBaseParams {
	return &GetV1MarketSwaprateBaseParams{
		timeout: timeout,
	}
}

// NewGetV1MarketSwaprateBaseParamsWithContext creates a new GetV1MarketSwaprateBaseParams object
// with the ability to set a context for a request.
func NewGetV1MarketSwaprateBaseParamsWithContext(ctx context.Context) *GetV1MarketSwaprateBaseParams {
	return &GetV1MarketSwaprateBaseParams{
		Context: ctx,
	}
}

// NewGetV1MarketSwaprateBaseParamsWithHTTPClient creates a new GetV1MarketSwaprateBaseParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1MarketSwaprateBaseParamsWithHTTPClient(client *http.Client) *GetV1MarketSwaprateBaseParams {
	return &GetV1MarketSwaprateBaseParams{
		HTTPClient: client,
	}
}

/* GetV1MarketSwaprateBaseParams contains all the parameters to send to the API endpoint
   for the get v1 market swaprate base operation.

   Typically these are written to a http.Request.
*/
type GetV1MarketSwaprateBaseParams struct {

	/* Base.

	   Coin denomination
	*/
	Base string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 market swaprate base params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1MarketSwaprateBaseParams) WithDefaults() *GetV1MarketSwaprateBaseParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 market swaprate base params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1MarketSwaprateBaseParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 market swaprate base params
func (o *GetV1MarketSwaprateBaseParams) WithTimeout(timeout time.Duration) *GetV1MarketSwaprateBaseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 market swaprate base params
func (o *GetV1MarketSwaprateBaseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 market swaprate base params
func (o *GetV1MarketSwaprateBaseParams) WithContext(ctx context.Context) *GetV1MarketSwaprateBaseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 market swaprate base params
func (o *GetV1MarketSwaprateBaseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 market swaprate base params
func (o *GetV1MarketSwaprateBaseParams) WithHTTPClient(client *http.Client) *GetV1MarketSwaprateBaseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 market swaprate base params
func (o *GetV1MarketSwaprateBaseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBase adds the base to the get v1 market swaprate base params
func (o *GetV1MarketSwaprateBaseParams) WithBase(base string) *GetV1MarketSwaprateBaseParams {
	o.SetBase(base)
	return o
}

// SetBase adds the base to the get v1 market swaprate base params
func (o *GetV1MarketSwaprateBaseParams) SetBase(base string) {
	o.Base = base
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1MarketSwaprateBaseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param base
	if err := r.SetPathParam("base", o.Base); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

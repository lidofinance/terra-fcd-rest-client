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

// NewGovParamsParams creates a new GovParamsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGovParamsParams() *GovParamsParams {
	return &GovParamsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGovParamsParamsWithTimeout creates a new GovParamsParams object
// with the ability to set a timeout on a request.
func NewGovParamsParamsWithTimeout(timeout time.Duration) *GovParamsParams {
	return &GovParamsParams{
		timeout: timeout,
	}
}

// NewGovParamsParamsWithContext creates a new GovParamsParams object
// with the ability to set a context for a request.
func NewGovParamsParamsWithContext(ctx context.Context) *GovParamsParams {
	return &GovParamsParams{
		Context: ctx,
	}
}

// NewGovParamsParamsWithHTTPClient creates a new GovParamsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGovParamsParamsWithHTTPClient(client *http.Client) *GovParamsParams {
	return &GovParamsParams{
		HTTPClient: client,
	}
}

/* GovParamsParams contains all the parameters to send to the API endpoint
   for the gov params operation.

   Typically these are written to a http.Request.
*/
type GovParamsParams struct {

	/* ParamsType.

	     params_type defines which parameters to query for, can be one of "voting",
	"tallying" or "deposit".
	*/
	ParamsType string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the gov params params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GovParamsParams) WithDefaults() *GovParamsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the gov params params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GovParamsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the gov params params
func (o *GovParamsParams) WithTimeout(timeout time.Duration) *GovParamsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the gov params params
func (o *GovParamsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the gov params params
func (o *GovParamsParams) WithContext(ctx context.Context) *GovParamsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the gov params params
func (o *GovParamsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the gov params params
func (o *GovParamsParams) WithHTTPClient(client *http.Client) *GovParamsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the gov params params
func (o *GovParamsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithParamsType adds the paramsType to the gov params params
func (o *GovParamsParams) WithParamsType(paramsType string) *GovParamsParams {
	o.SetParamsType(paramsType)
	return o
}

// SetParamsType adds the paramsType to the gov params params
func (o *GovParamsParams) SetParamsType(paramsType string) {
	o.ParamsType = paramsType
}

// WriteToRequest writes these params to a swagger request
func (o *GovParamsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param params_type
	if err := r.SetPathParam("params_type", o.ParamsType); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
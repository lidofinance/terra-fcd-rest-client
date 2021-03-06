// Code generated by go-swagger; DO NOT EDIT.

package dashboard

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

// NewGetV1DashboardActiveAccountsParams creates a new GetV1DashboardActiveAccountsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1DashboardActiveAccountsParams() *GetV1DashboardActiveAccountsParams {
	return &GetV1DashboardActiveAccountsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1DashboardActiveAccountsParamsWithTimeout creates a new GetV1DashboardActiveAccountsParams object
// with the ability to set a timeout on a request.
func NewGetV1DashboardActiveAccountsParamsWithTimeout(timeout time.Duration) *GetV1DashboardActiveAccountsParams {
	return &GetV1DashboardActiveAccountsParams{
		timeout: timeout,
	}
}

// NewGetV1DashboardActiveAccountsParamsWithContext creates a new GetV1DashboardActiveAccountsParams object
// with the ability to set a context for a request.
func NewGetV1DashboardActiveAccountsParamsWithContext(ctx context.Context) *GetV1DashboardActiveAccountsParams {
	return &GetV1DashboardActiveAccountsParams{
		Context: ctx,
	}
}

// NewGetV1DashboardActiveAccountsParamsWithHTTPClient creates a new GetV1DashboardActiveAccountsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1DashboardActiveAccountsParamsWithHTTPClient(client *http.Client) *GetV1DashboardActiveAccountsParams {
	return &GetV1DashboardActiveAccountsParams{
		HTTPClient: client,
	}
}

/* GetV1DashboardActiveAccountsParams contains all the parameters to send to the API endpoint
   for the get v1 dashboard active accounts operation.

   Typically these are written to a http.Request.
*/
type GetV1DashboardActiveAccountsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 dashboard active accounts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1DashboardActiveAccountsParams) WithDefaults() *GetV1DashboardActiveAccountsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 dashboard active accounts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1DashboardActiveAccountsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 dashboard active accounts params
func (o *GetV1DashboardActiveAccountsParams) WithTimeout(timeout time.Duration) *GetV1DashboardActiveAccountsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 dashboard active accounts params
func (o *GetV1DashboardActiveAccountsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 dashboard active accounts params
func (o *GetV1DashboardActiveAccountsParams) WithContext(ctx context.Context) *GetV1DashboardActiveAccountsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 dashboard active accounts params
func (o *GetV1DashboardActiveAccountsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 dashboard active accounts params
func (o *GetV1DashboardActiveAccountsParams) WithHTTPClient(client *http.Client) *GetV1DashboardActiveAccountsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 dashboard active accounts params
func (o *GetV1DashboardActiveAccountsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1DashboardActiveAccountsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

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

// NewAppliedPlanParams creates a new AppliedPlanParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAppliedPlanParams() *AppliedPlanParams {
	return &AppliedPlanParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAppliedPlanParamsWithTimeout creates a new AppliedPlanParams object
// with the ability to set a timeout on a request.
func NewAppliedPlanParamsWithTimeout(timeout time.Duration) *AppliedPlanParams {
	return &AppliedPlanParams{
		timeout: timeout,
	}
}

// NewAppliedPlanParamsWithContext creates a new AppliedPlanParams object
// with the ability to set a context for a request.
func NewAppliedPlanParamsWithContext(ctx context.Context) *AppliedPlanParams {
	return &AppliedPlanParams{
		Context: ctx,
	}
}

// NewAppliedPlanParamsWithHTTPClient creates a new AppliedPlanParams object
// with the ability to set a custom HTTPClient for a request.
func NewAppliedPlanParamsWithHTTPClient(client *http.Client) *AppliedPlanParams {
	return &AppliedPlanParams{
		HTTPClient: client,
	}
}

/* AppliedPlanParams contains all the parameters to send to the API endpoint
   for the applied plan operation.

   Typically these are written to a http.Request.
*/
type AppliedPlanParams struct {

	/* Name.

	   name is the name of the applied plan to query for.
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the applied plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AppliedPlanParams) WithDefaults() *AppliedPlanParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the applied plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AppliedPlanParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the applied plan params
func (o *AppliedPlanParams) WithTimeout(timeout time.Duration) *AppliedPlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the applied plan params
func (o *AppliedPlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the applied plan params
func (o *AppliedPlanParams) WithContext(ctx context.Context) *AppliedPlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the applied plan params
func (o *AppliedPlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the applied plan params
func (o *AppliedPlanParams) WithHTTPClient(client *http.Client) *AppliedPlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the applied plan params
func (o *AppliedPlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the applied plan params
func (o *AppliedPlanParams) WithName(name string) *AppliedPlanParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the applied plan params
func (o *AppliedPlanParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *AppliedPlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

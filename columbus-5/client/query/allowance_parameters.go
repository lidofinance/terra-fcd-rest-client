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

// NewAllowanceParams creates a new AllowanceParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAllowanceParams() *AllowanceParams {
	return &AllowanceParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAllowanceParamsWithTimeout creates a new AllowanceParams object
// with the ability to set a timeout on a request.
func NewAllowanceParamsWithTimeout(timeout time.Duration) *AllowanceParams {
	return &AllowanceParams{
		timeout: timeout,
	}
}

// NewAllowanceParamsWithContext creates a new AllowanceParams object
// with the ability to set a context for a request.
func NewAllowanceParamsWithContext(ctx context.Context) *AllowanceParams {
	return &AllowanceParams{
		Context: ctx,
	}
}

// NewAllowanceParamsWithHTTPClient creates a new AllowanceParams object
// with the ability to set a custom HTTPClient for a request.
func NewAllowanceParamsWithHTTPClient(client *http.Client) *AllowanceParams {
	return &AllowanceParams{
		HTTPClient: client,
	}
}

/* AllowanceParams contains all the parameters to send to the API endpoint
   for the allowance operation.

   Typically these are written to a http.Request.
*/
type AllowanceParams struct {

	/* Grantee.

	   grantee is the address of the user being granted an allowance of another user's funds.
	*/
	Grantee string

	/* Granter.

	   granter is the address of the user granting an allowance of their funds.
	*/
	Granter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the allowance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AllowanceParams) WithDefaults() *AllowanceParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the allowance params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AllowanceParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the allowance params
func (o *AllowanceParams) WithTimeout(timeout time.Duration) *AllowanceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the allowance params
func (o *AllowanceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the allowance params
func (o *AllowanceParams) WithContext(ctx context.Context) *AllowanceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the allowance params
func (o *AllowanceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the allowance params
func (o *AllowanceParams) WithHTTPClient(client *http.Client) *AllowanceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the allowance params
func (o *AllowanceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGrantee adds the grantee to the allowance params
func (o *AllowanceParams) WithGrantee(grantee string) *AllowanceParams {
	o.SetGrantee(grantee)
	return o
}

// SetGrantee adds the grantee to the allowance params
func (o *AllowanceParams) SetGrantee(grantee string) {
	o.Grantee = grantee
}

// WithGranter adds the granter to the allowance params
func (o *AllowanceParams) WithGranter(granter string) *AllowanceParams {
	o.SetGranter(granter)
	return o
}

// SetGranter adds the granter to the allowance params
func (o *AllowanceParams) SetGranter(granter string) {
	o.Granter = granter
}

// WriteToRequest writes these params to a swagger request
func (o *AllowanceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param grantee
	if err := r.SetPathParam("grantee", o.Grantee); err != nil {
		return err
	}

	// path param granter
	if err := r.SetPathParam("granter", o.Granter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

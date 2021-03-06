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

// NewAuthParamsParams creates a new AuthParamsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAuthParamsParams() *AuthParamsParams {
	return &AuthParamsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAuthParamsParamsWithTimeout creates a new AuthParamsParams object
// with the ability to set a timeout on a request.
func NewAuthParamsParamsWithTimeout(timeout time.Duration) *AuthParamsParams {
	return &AuthParamsParams{
		timeout: timeout,
	}
}

// NewAuthParamsParamsWithContext creates a new AuthParamsParams object
// with the ability to set a context for a request.
func NewAuthParamsParamsWithContext(ctx context.Context) *AuthParamsParams {
	return &AuthParamsParams{
		Context: ctx,
	}
}

// NewAuthParamsParamsWithHTTPClient creates a new AuthParamsParams object
// with the ability to set a custom HTTPClient for a request.
func NewAuthParamsParamsWithHTTPClient(client *http.Client) *AuthParamsParams {
	return &AuthParamsParams{
		HTTPClient: client,
	}
}

/* AuthParamsParams contains all the parameters to send to the API endpoint
   for the auth params operation.

   Typically these are written to a http.Request.
*/
type AuthParamsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the auth params params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthParamsParams) WithDefaults() *AuthParamsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the auth params params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthParamsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the auth params params
func (o *AuthParamsParams) WithTimeout(timeout time.Duration) *AuthParamsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the auth params params
func (o *AuthParamsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the auth params params
func (o *AuthParamsParams) WithContext(ctx context.Context) *AuthParamsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the auth params params
func (o *AuthParamsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the auth params params
func (o *AuthParamsParams) WithHTTPClient(client *http.Client) *AuthParamsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the auth params params
func (o *AuthParamsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *AuthParamsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

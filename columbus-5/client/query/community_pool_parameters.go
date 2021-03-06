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

// NewCommunityPoolParams creates a new CommunityPoolParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCommunityPoolParams() *CommunityPoolParams {
	return &CommunityPoolParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCommunityPoolParamsWithTimeout creates a new CommunityPoolParams object
// with the ability to set a timeout on a request.
func NewCommunityPoolParamsWithTimeout(timeout time.Duration) *CommunityPoolParams {
	return &CommunityPoolParams{
		timeout: timeout,
	}
}

// NewCommunityPoolParamsWithContext creates a new CommunityPoolParams object
// with the ability to set a context for a request.
func NewCommunityPoolParamsWithContext(ctx context.Context) *CommunityPoolParams {
	return &CommunityPoolParams{
		Context: ctx,
	}
}

// NewCommunityPoolParamsWithHTTPClient creates a new CommunityPoolParams object
// with the ability to set a custom HTTPClient for a request.
func NewCommunityPoolParamsWithHTTPClient(client *http.Client) *CommunityPoolParams {
	return &CommunityPoolParams{
		HTTPClient: client,
	}
}

/* CommunityPoolParams contains all the parameters to send to the API endpoint
   for the community pool operation.

   Typically these are written to a http.Request.
*/
type CommunityPoolParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the community pool params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CommunityPoolParams) WithDefaults() *CommunityPoolParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the community pool params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CommunityPoolParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the community pool params
func (o *CommunityPoolParams) WithTimeout(timeout time.Duration) *CommunityPoolParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the community pool params
func (o *CommunityPoolParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the community pool params
func (o *CommunityPoolParams) WithContext(ctx context.Context) *CommunityPoolParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the community pool params
func (o *CommunityPoolParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the community pool params
func (o *CommunityPoolParams) WithHTTPClient(client *http.Client) *CommunityPoolParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the community pool params
func (o *CommunityPoolParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CommunityPoolParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

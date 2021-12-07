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

// NewHistoricalInfoParams creates a new HistoricalInfoParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHistoricalInfoParams() *HistoricalInfoParams {
	return &HistoricalInfoParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHistoricalInfoParamsWithTimeout creates a new HistoricalInfoParams object
// with the ability to set a timeout on a request.
func NewHistoricalInfoParamsWithTimeout(timeout time.Duration) *HistoricalInfoParams {
	return &HistoricalInfoParams{
		timeout: timeout,
	}
}

// NewHistoricalInfoParamsWithContext creates a new HistoricalInfoParams object
// with the ability to set a context for a request.
func NewHistoricalInfoParamsWithContext(ctx context.Context) *HistoricalInfoParams {
	return &HistoricalInfoParams{
		Context: ctx,
	}
}

// NewHistoricalInfoParamsWithHTTPClient creates a new HistoricalInfoParams object
// with the ability to set a custom HTTPClient for a request.
func NewHistoricalInfoParamsWithHTTPClient(client *http.Client) *HistoricalInfoParams {
	return &HistoricalInfoParams{
		HTTPClient: client,
	}
}

/* HistoricalInfoParams contains all the parameters to send to the API endpoint
   for the historical info operation.

   Typically these are written to a http.Request.
*/
type HistoricalInfoParams struct {

	/* Height.

	   height defines at which height to query the historical info.

	   Format: int64
	*/
	Height string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the historical info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HistoricalInfoParams) WithDefaults() *HistoricalInfoParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the historical info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HistoricalInfoParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the historical info params
func (o *HistoricalInfoParams) WithTimeout(timeout time.Duration) *HistoricalInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the historical info params
func (o *HistoricalInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the historical info params
func (o *HistoricalInfoParams) WithContext(ctx context.Context) *HistoricalInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the historical info params
func (o *HistoricalInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the historical info params
func (o *HistoricalInfoParams) WithHTTPClient(client *http.Client) *HistoricalInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the historical info params
func (o *HistoricalInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHeight adds the height to the historical info params
func (o *HistoricalInfoParams) WithHeight(height string) *HistoricalInfoParams {
	o.SetHeight(height)
	return o
}

// SetHeight adds the height to the historical info params
func (o *HistoricalInfoParams) SetHeight(height string) {
	o.Height = height
}

// WriteToRequest writes these params to a swagger request
func (o *HistoricalInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param height
	if err := r.SetPathParam("height", o.Height); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

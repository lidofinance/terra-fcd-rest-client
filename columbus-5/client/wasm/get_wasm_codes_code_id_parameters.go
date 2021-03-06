// Code generated by go-swagger; DO NOT EDIT.

package wasm

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
	"github.com/go-openapi/swag"
)

// NewGetWasmCodesCodeIDParams creates a new GetWasmCodesCodeIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWasmCodesCodeIDParams() *GetWasmCodesCodeIDParams {
	return &GetWasmCodesCodeIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWasmCodesCodeIDParamsWithTimeout creates a new GetWasmCodesCodeIDParams object
// with the ability to set a timeout on a request.
func NewGetWasmCodesCodeIDParamsWithTimeout(timeout time.Duration) *GetWasmCodesCodeIDParams {
	return &GetWasmCodesCodeIDParams{
		timeout: timeout,
	}
}

// NewGetWasmCodesCodeIDParamsWithContext creates a new GetWasmCodesCodeIDParams object
// with the ability to set a context for a request.
func NewGetWasmCodesCodeIDParamsWithContext(ctx context.Context) *GetWasmCodesCodeIDParams {
	return &GetWasmCodesCodeIDParams{
		Context: ctx,
	}
}

// NewGetWasmCodesCodeIDParamsWithHTTPClient creates a new GetWasmCodesCodeIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWasmCodesCodeIDParamsWithHTTPClient(client *http.Client) *GetWasmCodesCodeIDParams {
	return &GetWasmCodesCodeIDParams{
		HTTPClient: client,
	}
}

/* GetWasmCodesCodeIDParams contains all the parameters to send to the API endpoint
   for the get wasm codes code ID operation.

   Typically these are written to a http.Request.
*/
type GetWasmCodesCodeIDParams struct {

	/* CodeID.

	   code ID you want to instantiate
	*/
	CodeID float64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get wasm codes code ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWasmCodesCodeIDParams) WithDefaults() *GetWasmCodesCodeIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get wasm codes code ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWasmCodesCodeIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get wasm codes code ID params
func (o *GetWasmCodesCodeIDParams) WithTimeout(timeout time.Duration) *GetWasmCodesCodeIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get wasm codes code ID params
func (o *GetWasmCodesCodeIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get wasm codes code ID params
func (o *GetWasmCodesCodeIDParams) WithContext(ctx context.Context) *GetWasmCodesCodeIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get wasm codes code ID params
func (o *GetWasmCodesCodeIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get wasm codes code ID params
func (o *GetWasmCodesCodeIDParams) WithHTTPClient(client *http.Client) *GetWasmCodesCodeIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get wasm codes code ID params
func (o *GetWasmCodesCodeIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCodeID adds the codeID to the get wasm codes code ID params
func (o *GetWasmCodesCodeIDParams) WithCodeID(codeID float64) *GetWasmCodesCodeIDParams {
	o.SetCodeID(codeID)
	return o
}

// SetCodeID adds the codeId to the get wasm codes code ID params
func (o *GetWasmCodesCodeIDParams) SetCodeID(codeID float64) {
	o.CodeID = codeID
}

// WriteToRequest writes these params to a swagger request
func (o *GetWasmCodesCodeIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param codeID
	if err := r.SetPathParam("codeID", swag.FormatFloat64(o.CodeID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

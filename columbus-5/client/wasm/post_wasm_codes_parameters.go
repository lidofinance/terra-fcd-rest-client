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
)

// NewPostWasmCodesParams creates a new PostWasmCodesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostWasmCodesParams() *PostWasmCodesParams {
	return &PostWasmCodesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostWasmCodesParamsWithTimeout creates a new PostWasmCodesParams object
// with the ability to set a timeout on a request.
func NewPostWasmCodesParamsWithTimeout(timeout time.Duration) *PostWasmCodesParams {
	return &PostWasmCodesParams{
		timeout: timeout,
	}
}

// NewPostWasmCodesParamsWithContext creates a new PostWasmCodesParams object
// with the ability to set a context for a request.
func NewPostWasmCodesParamsWithContext(ctx context.Context) *PostWasmCodesParams {
	return &PostWasmCodesParams{
		Context: ctx,
	}
}

// NewPostWasmCodesParamsWithHTTPClient creates a new PostWasmCodesParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostWasmCodesParamsWithHTTPClient(client *http.Client) *PostWasmCodesParams {
	return &PostWasmCodesParams{
		HTTPClient: client,
	}
}

/* PostWasmCodesParams contains all the parameters to send to the API endpoint
   for the post wasm codes operation.

   Typically these are written to a http.Request.
*/
type PostWasmCodesParams struct {

	// StoreCodeRequestBody.
	StoreCodeRequestBody PostWasmCodesBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post wasm codes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostWasmCodesParams) WithDefaults() *PostWasmCodesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post wasm codes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostWasmCodesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post wasm codes params
func (o *PostWasmCodesParams) WithTimeout(timeout time.Duration) *PostWasmCodesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post wasm codes params
func (o *PostWasmCodesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post wasm codes params
func (o *PostWasmCodesParams) WithContext(ctx context.Context) *PostWasmCodesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post wasm codes params
func (o *PostWasmCodesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post wasm codes params
func (o *PostWasmCodesParams) WithHTTPClient(client *http.Client) *PostWasmCodesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post wasm codes params
func (o *PostWasmCodesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithStoreCodeRequestBody adds the storeCodeRequestBody to the post wasm codes params
func (o *PostWasmCodesParams) WithStoreCodeRequestBody(storeCodeRequestBody PostWasmCodesBody) *PostWasmCodesParams {
	o.SetStoreCodeRequestBody(storeCodeRequestBody)
	return o
}

// SetStoreCodeRequestBody adds the storeCodeRequestBody to the post wasm codes params
func (o *PostWasmCodesParams) SetStoreCodeRequestBody(storeCodeRequestBody PostWasmCodesBody) {
	o.StoreCodeRequestBody = storeCodeRequestBody
}

// WriteToRequest writes these params to a swagger request
func (o *PostWasmCodesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.StoreCodeRequestBody); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

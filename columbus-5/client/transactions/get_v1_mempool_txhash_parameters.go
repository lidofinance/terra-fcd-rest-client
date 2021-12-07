// Code generated by go-swagger; DO NOT EDIT.

package transactions

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

// NewGetV1MempoolTxhashParams creates a new GetV1MempoolTxhashParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1MempoolTxhashParams() *GetV1MempoolTxhashParams {
	return &GetV1MempoolTxhashParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1MempoolTxhashParamsWithTimeout creates a new GetV1MempoolTxhashParams object
// with the ability to set a timeout on a request.
func NewGetV1MempoolTxhashParamsWithTimeout(timeout time.Duration) *GetV1MempoolTxhashParams {
	return &GetV1MempoolTxhashParams{
		timeout: timeout,
	}
}

// NewGetV1MempoolTxhashParamsWithContext creates a new GetV1MempoolTxhashParams object
// with the ability to set a context for a request.
func NewGetV1MempoolTxhashParamsWithContext(ctx context.Context) *GetV1MempoolTxhashParams {
	return &GetV1MempoolTxhashParams{
		Context: ctx,
	}
}

// NewGetV1MempoolTxhashParamsWithHTTPClient creates a new GetV1MempoolTxhashParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1MempoolTxhashParamsWithHTTPClient(client *http.Client) *GetV1MempoolTxhashParams {
	return &GetV1MempoolTxhashParams{
		HTTPClient: client,
	}
}

/* GetV1MempoolTxhashParams contains all the parameters to send to the API endpoint
   for the get v1 mempool txhash operation.

   Typically these are written to a http.Request.
*/
type GetV1MempoolTxhashParams struct {

	/* Txhash.

	   Tx Hash
	*/
	Txhash string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 mempool txhash params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1MempoolTxhashParams) WithDefaults() *GetV1MempoolTxhashParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 mempool txhash params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1MempoolTxhashParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 mempool txhash params
func (o *GetV1MempoolTxhashParams) WithTimeout(timeout time.Duration) *GetV1MempoolTxhashParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 mempool txhash params
func (o *GetV1MempoolTxhashParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 mempool txhash params
func (o *GetV1MempoolTxhashParams) WithContext(ctx context.Context) *GetV1MempoolTxhashParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 mempool txhash params
func (o *GetV1MempoolTxhashParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 mempool txhash params
func (o *GetV1MempoolTxhashParams) WithHTTPClient(client *http.Client) *GetV1MempoolTxhashParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 mempool txhash params
func (o *GetV1MempoolTxhashParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithTxhash adds the txhash to the get v1 mempool txhash params
func (o *GetV1MempoolTxhashParams) WithTxhash(txhash string) *GetV1MempoolTxhashParams {
	o.SetTxhash(txhash)
	return o
}

// SetTxhash adds the txhash to the get v1 mempool txhash params
func (o *GetV1MempoolTxhashParams) SetTxhash(txhash string) {
	o.Txhash = txhash
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1MempoolTxhashParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param txhash
	if err := r.SetPathParam("txhash", o.Txhash); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

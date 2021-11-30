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

// NewGetWasmContractsContractAddressParams creates a new GetWasmContractsContractAddressParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWasmContractsContractAddressParams() *GetWasmContractsContractAddressParams {
	return &GetWasmContractsContractAddressParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWasmContractsContractAddressParamsWithTimeout creates a new GetWasmContractsContractAddressParams object
// with the ability to set a timeout on a request.
func NewGetWasmContractsContractAddressParamsWithTimeout(timeout time.Duration) *GetWasmContractsContractAddressParams {
	return &GetWasmContractsContractAddressParams{
		timeout: timeout,
	}
}

// NewGetWasmContractsContractAddressParamsWithContext creates a new GetWasmContractsContractAddressParams object
// with the ability to set a context for a request.
func NewGetWasmContractsContractAddressParamsWithContext(ctx context.Context) *GetWasmContractsContractAddressParams {
	return &GetWasmContractsContractAddressParams{
		Context: ctx,
	}
}

// NewGetWasmContractsContractAddressParamsWithHTTPClient creates a new GetWasmContractsContractAddressParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWasmContractsContractAddressParamsWithHTTPClient(client *http.Client) *GetWasmContractsContractAddressParams {
	return &GetWasmContractsContractAddressParams{
		HTTPClient: client,
	}
}

/* GetWasmContractsContractAddressParams contains all the parameters to send to the API endpoint
   for the get wasm contracts contract address operation.

   Typically these are written to a http.Request.
*/
type GetWasmContractsContractAddressParams struct {

	/* ContractAddress.

	   contract address you want to execute
	*/
	ContractAddress string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get wasm contracts contract address params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWasmContractsContractAddressParams) WithDefaults() *GetWasmContractsContractAddressParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get wasm contracts contract address params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWasmContractsContractAddressParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get wasm contracts contract address params
func (o *GetWasmContractsContractAddressParams) WithTimeout(timeout time.Duration) *GetWasmContractsContractAddressParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get wasm contracts contract address params
func (o *GetWasmContractsContractAddressParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get wasm contracts contract address params
func (o *GetWasmContractsContractAddressParams) WithContext(ctx context.Context) *GetWasmContractsContractAddressParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get wasm contracts contract address params
func (o *GetWasmContractsContractAddressParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get wasm contracts contract address params
func (o *GetWasmContractsContractAddressParams) WithHTTPClient(client *http.Client) *GetWasmContractsContractAddressParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get wasm contracts contract address params
func (o *GetWasmContractsContractAddressParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContractAddress adds the contractAddress to the get wasm contracts contract address params
func (o *GetWasmContractsContractAddressParams) WithContractAddress(contractAddress string) *GetWasmContractsContractAddressParams {
	o.SetContractAddress(contractAddress)
	return o
}

// SetContractAddress adds the contractAddress to the get wasm contracts contract address params
func (o *GetWasmContractsContractAddressParams) SetContractAddress(contractAddress string) {
	o.ContractAddress = contractAddress
}

// WriteToRequest writes these params to a swagger request
func (o *GetWasmContractsContractAddressParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param contractAddress
	if err := r.SetPathParam("contractAddress", o.ContractAddress); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
// Code generated by go-swagger; DO NOT EDIT.

package distribution

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

// NewPostDistributionDelegatorsDelegatorAddrWithdrawAddressParams creates a new PostDistributionDelegatorsDelegatorAddrWithdrawAddressParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostDistributionDelegatorsDelegatorAddrWithdrawAddressParams() *PostDistributionDelegatorsDelegatorAddrWithdrawAddressParams {
	return &PostDistributionDelegatorsDelegatorAddrWithdrawAddressParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostDistributionDelegatorsDelegatorAddrWithdrawAddressParamsWithTimeout creates a new PostDistributionDelegatorsDelegatorAddrWithdrawAddressParams object
// with the ability to set a timeout on a request.
func NewPostDistributionDelegatorsDelegatorAddrWithdrawAddressParamsWithTimeout(timeout time.Duration) *PostDistributionDelegatorsDelegatorAddrWithdrawAddressParams {
	return &PostDistributionDelegatorsDelegatorAddrWithdrawAddressParams{
		timeout: timeout,
	}
}

// NewPostDistributionDelegatorsDelegatorAddrWithdrawAddressParamsWithContext creates a new PostDistributionDelegatorsDelegatorAddrWithdrawAddressParams object
// with the ability to set a context for a request.
func NewPostDistributionDelegatorsDelegatorAddrWithdrawAddressParamsWithContext(ctx context.Context) *PostDistributionDelegatorsDelegatorAddrWithdrawAddressParams {
	return &PostDistributionDelegatorsDelegatorAddrWithdrawAddressParams{
		Context: ctx,
	}
}

// NewPostDistributionDelegatorsDelegatorAddrWithdrawAddressParamsWithHTTPClient creates a new PostDistributionDelegatorsDelegatorAddrWithdrawAddressParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostDistributionDelegatorsDelegatorAddrWithdrawAddressParamsWithHTTPClient(client *http.Client) *PostDistributionDelegatorsDelegatorAddrWithdrawAddressParams {
	return &PostDistributionDelegatorsDelegatorAddrWithdrawAddressParams{
		HTTPClient: client,
	}
}

/* PostDistributionDelegatorsDelegatorAddrWithdrawAddressParams contains all the parameters to send to the API endpoint
   for the post distribution delegators delegator addr withdraw address operation.

   Typically these are written to a http.Request.
*/
type PostDistributionDelegatorsDelegatorAddrWithdrawAddressParams struct {

	// WithdrawRequestBody.
	WithdrawRequestBody PostDistributionDelegatorsDelegatorAddrWithdrawAddressBody

	/* DelegatorAddr.

	   Bech32 AccAddress of Delegator
	*/
	DelegatorAddr string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post distribution delegators delegator addr withdraw address params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDistributionDelegatorsDelegatorAddrWithdrawAddressParams) WithDefaults() *PostDistributionDelegatorsDelegatorAddrWithdrawAddressParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post distribution delegators delegator addr withdraw address params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostDistributionDelegatorsDelegatorAddrWithdrawAddressParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post distribution delegators delegator addr withdraw address params
func (o *PostDistributionDelegatorsDelegatorAddrWithdrawAddressParams) WithTimeout(timeout time.Duration) *PostDistributionDelegatorsDelegatorAddrWithdrawAddressParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post distribution delegators delegator addr withdraw address params
func (o *PostDistributionDelegatorsDelegatorAddrWithdrawAddressParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post distribution delegators delegator addr withdraw address params
func (o *PostDistributionDelegatorsDelegatorAddrWithdrawAddressParams) WithContext(ctx context.Context) *PostDistributionDelegatorsDelegatorAddrWithdrawAddressParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post distribution delegators delegator addr withdraw address params
func (o *PostDistributionDelegatorsDelegatorAddrWithdrawAddressParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post distribution delegators delegator addr withdraw address params
func (o *PostDistributionDelegatorsDelegatorAddrWithdrawAddressParams) WithHTTPClient(client *http.Client) *PostDistributionDelegatorsDelegatorAddrWithdrawAddressParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post distribution delegators delegator addr withdraw address params
func (o *PostDistributionDelegatorsDelegatorAddrWithdrawAddressParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithWithdrawRequestBody adds the withdrawRequestBody to the post distribution delegators delegator addr withdraw address params
func (o *PostDistributionDelegatorsDelegatorAddrWithdrawAddressParams) WithWithdrawRequestBody(withdrawRequestBody PostDistributionDelegatorsDelegatorAddrWithdrawAddressBody) *PostDistributionDelegatorsDelegatorAddrWithdrawAddressParams {
	o.SetWithdrawRequestBody(withdrawRequestBody)
	return o
}

// SetWithdrawRequestBody adds the withdrawRequestBody to the post distribution delegators delegator addr withdraw address params
func (o *PostDistributionDelegatorsDelegatorAddrWithdrawAddressParams) SetWithdrawRequestBody(withdrawRequestBody PostDistributionDelegatorsDelegatorAddrWithdrawAddressBody) {
	o.WithdrawRequestBody = withdrawRequestBody
}

// WithDelegatorAddr adds the delegatorAddr to the post distribution delegators delegator addr withdraw address params
func (o *PostDistributionDelegatorsDelegatorAddrWithdrawAddressParams) WithDelegatorAddr(delegatorAddr string) *PostDistributionDelegatorsDelegatorAddrWithdrawAddressParams {
	o.SetDelegatorAddr(delegatorAddr)
	return o
}

// SetDelegatorAddr adds the delegatorAddr to the post distribution delegators delegator addr withdraw address params
func (o *PostDistributionDelegatorsDelegatorAddrWithdrawAddressParams) SetDelegatorAddr(delegatorAddr string) {
	o.DelegatorAddr = delegatorAddr
}

// WriteToRequest writes these params to a swagger request
func (o *PostDistributionDelegatorsDelegatorAddrWithdrawAddressParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.WithdrawRequestBody); err != nil {
		return err
	}

	// path param delegatorAddr
	if err := r.SetPathParam("delegatorAddr", o.DelegatorAddr); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

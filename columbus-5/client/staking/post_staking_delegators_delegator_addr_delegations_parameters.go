// Code generated by go-swagger; DO NOT EDIT.

package staking

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

// NewPostStakingDelegatorsDelegatorAddrDelegationsParams creates a new PostStakingDelegatorsDelegatorAddrDelegationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostStakingDelegatorsDelegatorAddrDelegationsParams() *PostStakingDelegatorsDelegatorAddrDelegationsParams {
	return &PostStakingDelegatorsDelegatorAddrDelegationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostStakingDelegatorsDelegatorAddrDelegationsParamsWithTimeout creates a new PostStakingDelegatorsDelegatorAddrDelegationsParams object
// with the ability to set a timeout on a request.
func NewPostStakingDelegatorsDelegatorAddrDelegationsParamsWithTimeout(timeout time.Duration) *PostStakingDelegatorsDelegatorAddrDelegationsParams {
	return &PostStakingDelegatorsDelegatorAddrDelegationsParams{
		timeout: timeout,
	}
}

// NewPostStakingDelegatorsDelegatorAddrDelegationsParamsWithContext creates a new PostStakingDelegatorsDelegatorAddrDelegationsParams object
// with the ability to set a context for a request.
func NewPostStakingDelegatorsDelegatorAddrDelegationsParamsWithContext(ctx context.Context) *PostStakingDelegatorsDelegatorAddrDelegationsParams {
	return &PostStakingDelegatorsDelegatorAddrDelegationsParams{
		Context: ctx,
	}
}

// NewPostStakingDelegatorsDelegatorAddrDelegationsParamsWithHTTPClient creates a new PostStakingDelegatorsDelegatorAddrDelegationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostStakingDelegatorsDelegatorAddrDelegationsParamsWithHTTPClient(client *http.Client) *PostStakingDelegatorsDelegatorAddrDelegationsParams {
	return &PostStakingDelegatorsDelegatorAddrDelegationsParams{
		HTTPClient: client,
	}
}

/* PostStakingDelegatorsDelegatorAddrDelegationsParams contains all the parameters to send to the API endpoint
   for the post staking delegators delegator addr delegations operation.

   Typically these are written to a http.Request.
*/
type PostStakingDelegatorsDelegatorAddrDelegationsParams struct {

	/* Delegation.

	   Delegate an amount of liquid coins to a validator
	*/
	Delegation PostStakingDelegatorsDelegatorAddrDelegationsBody

	/* DelegatorAddr.

	   Bech32 AccAddress of Delegator
	*/
	DelegatorAddr string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post staking delegators delegator addr delegations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostStakingDelegatorsDelegatorAddrDelegationsParams) WithDefaults() *PostStakingDelegatorsDelegatorAddrDelegationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post staking delegators delegator addr delegations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostStakingDelegatorsDelegatorAddrDelegationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post staking delegators delegator addr delegations params
func (o *PostStakingDelegatorsDelegatorAddrDelegationsParams) WithTimeout(timeout time.Duration) *PostStakingDelegatorsDelegatorAddrDelegationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post staking delegators delegator addr delegations params
func (o *PostStakingDelegatorsDelegatorAddrDelegationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post staking delegators delegator addr delegations params
func (o *PostStakingDelegatorsDelegatorAddrDelegationsParams) WithContext(ctx context.Context) *PostStakingDelegatorsDelegatorAddrDelegationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post staking delegators delegator addr delegations params
func (o *PostStakingDelegatorsDelegatorAddrDelegationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post staking delegators delegator addr delegations params
func (o *PostStakingDelegatorsDelegatorAddrDelegationsParams) WithHTTPClient(client *http.Client) *PostStakingDelegatorsDelegatorAddrDelegationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post staking delegators delegator addr delegations params
func (o *PostStakingDelegatorsDelegatorAddrDelegationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDelegation adds the delegation to the post staking delegators delegator addr delegations params
func (o *PostStakingDelegatorsDelegatorAddrDelegationsParams) WithDelegation(delegation PostStakingDelegatorsDelegatorAddrDelegationsBody) *PostStakingDelegatorsDelegatorAddrDelegationsParams {
	o.SetDelegation(delegation)
	return o
}

// SetDelegation adds the delegation to the post staking delegators delegator addr delegations params
func (o *PostStakingDelegatorsDelegatorAddrDelegationsParams) SetDelegation(delegation PostStakingDelegatorsDelegatorAddrDelegationsBody) {
	o.Delegation = delegation
}

// WithDelegatorAddr adds the delegatorAddr to the post staking delegators delegator addr delegations params
func (o *PostStakingDelegatorsDelegatorAddrDelegationsParams) WithDelegatorAddr(delegatorAddr string) *PostStakingDelegatorsDelegatorAddrDelegationsParams {
	o.SetDelegatorAddr(delegatorAddr)
	return o
}

// SetDelegatorAddr adds the delegatorAddr to the post staking delegators delegator addr delegations params
func (o *PostStakingDelegatorsDelegatorAddrDelegationsParams) SetDelegatorAddr(delegatorAddr string) {
	o.DelegatorAddr = delegatorAddr
}

// WriteToRequest writes these params to a swagger request
func (o *PostStakingDelegatorsDelegatorAddrDelegationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Delegation); err != nil {
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

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

// NewGetStakingDelegatorsDelegatorAddrDelegationsParams creates a new GetStakingDelegatorsDelegatorAddrDelegationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetStakingDelegatorsDelegatorAddrDelegationsParams() *GetStakingDelegatorsDelegatorAddrDelegationsParams {
	return &GetStakingDelegatorsDelegatorAddrDelegationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetStakingDelegatorsDelegatorAddrDelegationsParamsWithTimeout creates a new GetStakingDelegatorsDelegatorAddrDelegationsParams object
// with the ability to set a timeout on a request.
func NewGetStakingDelegatorsDelegatorAddrDelegationsParamsWithTimeout(timeout time.Duration) *GetStakingDelegatorsDelegatorAddrDelegationsParams {
	return &GetStakingDelegatorsDelegatorAddrDelegationsParams{
		timeout: timeout,
	}
}

// NewGetStakingDelegatorsDelegatorAddrDelegationsParamsWithContext creates a new GetStakingDelegatorsDelegatorAddrDelegationsParams object
// with the ability to set a context for a request.
func NewGetStakingDelegatorsDelegatorAddrDelegationsParamsWithContext(ctx context.Context) *GetStakingDelegatorsDelegatorAddrDelegationsParams {
	return &GetStakingDelegatorsDelegatorAddrDelegationsParams{
		Context: ctx,
	}
}

// NewGetStakingDelegatorsDelegatorAddrDelegationsParamsWithHTTPClient creates a new GetStakingDelegatorsDelegatorAddrDelegationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetStakingDelegatorsDelegatorAddrDelegationsParamsWithHTTPClient(client *http.Client) *GetStakingDelegatorsDelegatorAddrDelegationsParams {
	return &GetStakingDelegatorsDelegatorAddrDelegationsParams{
		HTTPClient: client,
	}
}

/* GetStakingDelegatorsDelegatorAddrDelegationsParams contains all the parameters to send to the API endpoint
   for the get staking delegators delegator addr delegations operation.

   Typically these are written to a http.Request.
*/
type GetStakingDelegatorsDelegatorAddrDelegationsParams struct {

	/* DelegatorAddr.

	   Bech32 AccAddress of Delegator
	*/
	DelegatorAddr string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get staking delegators delegator addr delegations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStakingDelegatorsDelegatorAddrDelegationsParams) WithDefaults() *GetStakingDelegatorsDelegatorAddrDelegationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get staking delegators delegator addr delegations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetStakingDelegatorsDelegatorAddrDelegationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get staking delegators delegator addr delegations params
func (o *GetStakingDelegatorsDelegatorAddrDelegationsParams) WithTimeout(timeout time.Duration) *GetStakingDelegatorsDelegatorAddrDelegationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get staking delegators delegator addr delegations params
func (o *GetStakingDelegatorsDelegatorAddrDelegationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get staking delegators delegator addr delegations params
func (o *GetStakingDelegatorsDelegatorAddrDelegationsParams) WithContext(ctx context.Context) *GetStakingDelegatorsDelegatorAddrDelegationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get staking delegators delegator addr delegations params
func (o *GetStakingDelegatorsDelegatorAddrDelegationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get staking delegators delegator addr delegations params
func (o *GetStakingDelegatorsDelegatorAddrDelegationsParams) WithHTTPClient(client *http.Client) *GetStakingDelegatorsDelegatorAddrDelegationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get staking delegators delegator addr delegations params
func (o *GetStakingDelegatorsDelegatorAddrDelegationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDelegatorAddr adds the delegatorAddr to the get staking delegators delegator addr delegations params
func (o *GetStakingDelegatorsDelegatorAddrDelegationsParams) WithDelegatorAddr(delegatorAddr string) *GetStakingDelegatorsDelegatorAddrDelegationsParams {
	o.SetDelegatorAddr(delegatorAddr)
	return o
}

// SetDelegatorAddr adds the delegatorAddr to the get staking delegators delegator addr delegations params
func (o *GetStakingDelegatorsDelegatorAddrDelegationsParams) SetDelegatorAddr(delegatorAddr string) {
	o.DelegatorAddr = delegatorAddr
}

// WriteToRequest writes these params to a swagger request
func (o *GetStakingDelegatorsDelegatorAddrDelegationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param delegatorAddr
	if err := r.SetPathParam("delegatorAddr", o.DelegatorAddr); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

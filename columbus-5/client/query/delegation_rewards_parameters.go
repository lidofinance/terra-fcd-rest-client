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

// NewDelegationRewardsParams creates a new DelegationRewardsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDelegationRewardsParams() *DelegationRewardsParams {
	return &DelegationRewardsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDelegationRewardsParamsWithTimeout creates a new DelegationRewardsParams object
// with the ability to set a timeout on a request.
func NewDelegationRewardsParamsWithTimeout(timeout time.Duration) *DelegationRewardsParams {
	return &DelegationRewardsParams{
		timeout: timeout,
	}
}

// NewDelegationRewardsParamsWithContext creates a new DelegationRewardsParams object
// with the ability to set a context for a request.
func NewDelegationRewardsParamsWithContext(ctx context.Context) *DelegationRewardsParams {
	return &DelegationRewardsParams{
		Context: ctx,
	}
}

// NewDelegationRewardsParamsWithHTTPClient creates a new DelegationRewardsParams object
// with the ability to set a custom HTTPClient for a request.
func NewDelegationRewardsParamsWithHTTPClient(client *http.Client) *DelegationRewardsParams {
	return &DelegationRewardsParams{
		HTTPClient: client,
	}
}

/* DelegationRewardsParams contains all the parameters to send to the API endpoint
   for the delegation rewards operation.

   Typically these are written to a http.Request.
*/
type DelegationRewardsParams struct {

	/* DelegatorAddress.

	   delegator_address defines the delegator address to query for.
	*/
	DelegatorAddress string

	/* ValidatorAddress.

	   validator_address defines the validator address to query for.
	*/
	ValidatorAddress string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delegation rewards params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DelegationRewardsParams) WithDefaults() *DelegationRewardsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delegation rewards params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DelegationRewardsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delegation rewards params
func (o *DelegationRewardsParams) WithTimeout(timeout time.Duration) *DelegationRewardsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delegation rewards params
func (o *DelegationRewardsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delegation rewards params
func (o *DelegationRewardsParams) WithContext(ctx context.Context) *DelegationRewardsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delegation rewards params
func (o *DelegationRewardsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delegation rewards params
func (o *DelegationRewardsParams) WithHTTPClient(client *http.Client) *DelegationRewardsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delegation rewards params
func (o *DelegationRewardsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDelegatorAddress adds the delegatorAddress to the delegation rewards params
func (o *DelegationRewardsParams) WithDelegatorAddress(delegatorAddress string) *DelegationRewardsParams {
	o.SetDelegatorAddress(delegatorAddress)
	return o
}

// SetDelegatorAddress adds the delegatorAddress to the delegation rewards params
func (o *DelegationRewardsParams) SetDelegatorAddress(delegatorAddress string) {
	o.DelegatorAddress = delegatorAddress
}

// WithValidatorAddress adds the validatorAddress to the delegation rewards params
func (o *DelegationRewardsParams) WithValidatorAddress(validatorAddress string) *DelegationRewardsParams {
	o.SetValidatorAddress(validatorAddress)
	return o
}

// SetValidatorAddress adds the validatorAddress to the delegation rewards params
func (o *DelegationRewardsParams) SetValidatorAddress(validatorAddress string) {
	o.ValidatorAddress = validatorAddress
}

// WriteToRequest writes these params to a swagger request
func (o *DelegationRewardsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param delegator_address
	if err := r.SetPathParam("delegator_address", o.DelegatorAddress); err != nil {
		return err
	}

	// path param validator_address
	if err := r.SetPathParam("validator_address", o.ValidatorAddress); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

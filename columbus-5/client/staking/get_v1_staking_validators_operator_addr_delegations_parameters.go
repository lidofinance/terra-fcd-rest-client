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
	"github.com/go-openapi/swag"
)

// NewGetV1StakingValidatorsOperatorAddrDelegationsParams creates a new GetV1StakingValidatorsOperatorAddrDelegationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1StakingValidatorsOperatorAddrDelegationsParams() *GetV1StakingValidatorsOperatorAddrDelegationsParams {
	return &GetV1StakingValidatorsOperatorAddrDelegationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1StakingValidatorsOperatorAddrDelegationsParamsWithTimeout creates a new GetV1StakingValidatorsOperatorAddrDelegationsParams object
// with the ability to set a timeout on a request.
func NewGetV1StakingValidatorsOperatorAddrDelegationsParamsWithTimeout(timeout time.Duration) *GetV1StakingValidatorsOperatorAddrDelegationsParams {
	return &GetV1StakingValidatorsOperatorAddrDelegationsParams{
		timeout: timeout,
	}
}

// NewGetV1StakingValidatorsOperatorAddrDelegationsParamsWithContext creates a new GetV1StakingValidatorsOperatorAddrDelegationsParams object
// with the ability to set a context for a request.
func NewGetV1StakingValidatorsOperatorAddrDelegationsParamsWithContext(ctx context.Context) *GetV1StakingValidatorsOperatorAddrDelegationsParams {
	return &GetV1StakingValidatorsOperatorAddrDelegationsParams{
		Context: ctx,
	}
}

// NewGetV1StakingValidatorsOperatorAddrDelegationsParamsWithHTTPClient creates a new GetV1StakingValidatorsOperatorAddrDelegationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1StakingValidatorsOperatorAddrDelegationsParamsWithHTTPClient(client *http.Client) *GetV1StakingValidatorsOperatorAddrDelegationsParams {
	return &GetV1StakingValidatorsOperatorAddrDelegationsParams{
		HTTPClient: client,
	}
}

/* GetV1StakingValidatorsOperatorAddrDelegationsParams contains all the parameters to send to the API endpoint
   for the get v1 staking validators operator addr delegations operation.

   Typically these are written to a http.Request.
*/
type GetV1StakingValidatorsOperatorAddrDelegationsParams struct {

	/* Limit.

	   Page size
	*/
	Limit *float64

	/* OperatorAddr.

	   validators operator address
	*/
	OperatorAddr string

	/* Page.

	   Page number
	*/
	Page *float64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 staking validators operator addr delegations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1StakingValidatorsOperatorAddrDelegationsParams) WithDefaults() *GetV1StakingValidatorsOperatorAddrDelegationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 staking validators operator addr delegations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1StakingValidatorsOperatorAddrDelegationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 staking validators operator addr delegations params
func (o *GetV1StakingValidatorsOperatorAddrDelegationsParams) WithTimeout(timeout time.Duration) *GetV1StakingValidatorsOperatorAddrDelegationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 staking validators operator addr delegations params
func (o *GetV1StakingValidatorsOperatorAddrDelegationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 staking validators operator addr delegations params
func (o *GetV1StakingValidatorsOperatorAddrDelegationsParams) WithContext(ctx context.Context) *GetV1StakingValidatorsOperatorAddrDelegationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 staking validators operator addr delegations params
func (o *GetV1StakingValidatorsOperatorAddrDelegationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 staking validators operator addr delegations params
func (o *GetV1StakingValidatorsOperatorAddrDelegationsParams) WithHTTPClient(client *http.Client) *GetV1StakingValidatorsOperatorAddrDelegationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 staking validators operator addr delegations params
func (o *GetV1StakingValidatorsOperatorAddrDelegationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the get v1 staking validators operator addr delegations params
func (o *GetV1StakingValidatorsOperatorAddrDelegationsParams) WithLimit(limit *float64) *GetV1StakingValidatorsOperatorAddrDelegationsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get v1 staking validators operator addr delegations params
func (o *GetV1StakingValidatorsOperatorAddrDelegationsParams) SetLimit(limit *float64) {
	o.Limit = limit
}

// WithOperatorAddr adds the operatorAddr to the get v1 staking validators operator addr delegations params
func (o *GetV1StakingValidatorsOperatorAddrDelegationsParams) WithOperatorAddr(operatorAddr string) *GetV1StakingValidatorsOperatorAddrDelegationsParams {
	o.SetOperatorAddr(operatorAddr)
	return o
}

// SetOperatorAddr adds the operatorAddr to the get v1 staking validators operator addr delegations params
func (o *GetV1StakingValidatorsOperatorAddrDelegationsParams) SetOperatorAddr(operatorAddr string) {
	o.OperatorAddr = operatorAddr
}

// WithPage adds the page to the get v1 staking validators operator addr delegations params
func (o *GetV1StakingValidatorsOperatorAddrDelegationsParams) WithPage(page *float64) *GetV1StakingValidatorsOperatorAddrDelegationsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get v1 staking validators operator addr delegations params
func (o *GetV1StakingValidatorsOperatorAddrDelegationsParams) SetPage(page *float64) {
	o.Page = page
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1StakingValidatorsOperatorAddrDelegationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit float64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatFloat64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	// path param operatorAddr
	if err := r.SetPathParam("operatorAddr", o.OperatorAddr); err != nil {
		return err
	}

	if o.Page != nil {

		// query param page
		var qrPage float64

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatFloat64(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

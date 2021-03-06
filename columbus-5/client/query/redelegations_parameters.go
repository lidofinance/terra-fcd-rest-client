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
	"github.com/go-openapi/swag"
)

// NewRedelegationsParams creates a new RedelegationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRedelegationsParams() *RedelegationsParams {
	return &RedelegationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRedelegationsParamsWithTimeout creates a new RedelegationsParams object
// with the ability to set a timeout on a request.
func NewRedelegationsParamsWithTimeout(timeout time.Duration) *RedelegationsParams {
	return &RedelegationsParams{
		timeout: timeout,
	}
}

// NewRedelegationsParamsWithContext creates a new RedelegationsParams object
// with the ability to set a context for a request.
func NewRedelegationsParamsWithContext(ctx context.Context) *RedelegationsParams {
	return &RedelegationsParams{
		Context: ctx,
	}
}

// NewRedelegationsParamsWithHTTPClient creates a new RedelegationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewRedelegationsParamsWithHTTPClient(client *http.Client) *RedelegationsParams {
	return &RedelegationsParams{
		HTTPClient: client,
	}
}

/* RedelegationsParams contains all the parameters to send to the API endpoint
   for the redelegations operation.

   Typically these are written to a http.Request.
*/
type RedelegationsParams struct {

	/* DelegatorAddr.

	   delegator_addr defines the delegator address to query for.
	*/
	DelegatorAddr string

	/* DstValidatorAddr.

	   dst_validator_addr defines the validator address to redelegate to.
	*/
	DstValidatorAddr *string

	/* PaginationCountTotal.

	     count_total is set to true  to indicate that the result set should include
	a count of the total number of items available for pagination in UIs.
	count_total is only respected when offset is used. It is ignored when key
	is set.

	     Format: boolean
	*/
	PaginationCountTotal *bool

	/* PaginationKey.

	     key is a value returned in PageResponse.next_key to begin
	querying the next page most efficiently. Only one of offset or key
	should be set.

	     Format: byte
	*/
	PaginationKey *strfmt.Base64

	/* PaginationLimit.

	     limit is the total number of results to be returned in the result page.
	If left empty it will default to a value to be set by each app.

	     Format: uint64
	*/
	PaginationLimit *string

	/* PaginationOffset.

	     offset is a numeric offset that can be used when key is unavailable.
	It is less efficient than using key. Only one of offset or key should
	be set.

	     Format: uint64
	*/
	PaginationOffset *string

	/* PaginationReverse.

	   reverse is set to true if results are to be returned in the descending order.

	   Format: boolean
	*/
	PaginationReverse *bool

	/* SrcValidatorAddr.

	   src_validator_addr defines the validator address to redelegate from.
	*/
	SrcValidatorAddr *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the redelegations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RedelegationsParams) WithDefaults() *RedelegationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the redelegations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RedelegationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the redelegations params
func (o *RedelegationsParams) WithTimeout(timeout time.Duration) *RedelegationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the redelegations params
func (o *RedelegationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the redelegations params
func (o *RedelegationsParams) WithContext(ctx context.Context) *RedelegationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the redelegations params
func (o *RedelegationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the redelegations params
func (o *RedelegationsParams) WithHTTPClient(client *http.Client) *RedelegationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the redelegations params
func (o *RedelegationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDelegatorAddr adds the delegatorAddr to the redelegations params
func (o *RedelegationsParams) WithDelegatorAddr(delegatorAddr string) *RedelegationsParams {
	o.SetDelegatorAddr(delegatorAddr)
	return o
}

// SetDelegatorAddr adds the delegatorAddr to the redelegations params
func (o *RedelegationsParams) SetDelegatorAddr(delegatorAddr string) {
	o.DelegatorAddr = delegatorAddr
}

// WithDstValidatorAddr adds the dstValidatorAddr to the redelegations params
func (o *RedelegationsParams) WithDstValidatorAddr(dstValidatorAddr *string) *RedelegationsParams {
	o.SetDstValidatorAddr(dstValidatorAddr)
	return o
}

// SetDstValidatorAddr adds the dstValidatorAddr to the redelegations params
func (o *RedelegationsParams) SetDstValidatorAddr(dstValidatorAddr *string) {
	o.DstValidatorAddr = dstValidatorAddr
}

// WithPaginationCountTotal adds the paginationCountTotal to the redelegations params
func (o *RedelegationsParams) WithPaginationCountTotal(paginationCountTotal *bool) *RedelegationsParams {
	o.SetPaginationCountTotal(paginationCountTotal)
	return o
}

// SetPaginationCountTotal adds the paginationCountTotal to the redelegations params
func (o *RedelegationsParams) SetPaginationCountTotal(paginationCountTotal *bool) {
	o.PaginationCountTotal = paginationCountTotal
}

// WithPaginationKey adds the paginationKey to the redelegations params
func (o *RedelegationsParams) WithPaginationKey(paginationKey *strfmt.Base64) *RedelegationsParams {
	o.SetPaginationKey(paginationKey)
	return o
}

// SetPaginationKey adds the paginationKey to the redelegations params
func (o *RedelegationsParams) SetPaginationKey(paginationKey *strfmt.Base64) {
	o.PaginationKey = paginationKey
}

// WithPaginationLimit adds the paginationLimit to the redelegations params
func (o *RedelegationsParams) WithPaginationLimit(paginationLimit *string) *RedelegationsParams {
	o.SetPaginationLimit(paginationLimit)
	return o
}

// SetPaginationLimit adds the paginationLimit to the redelegations params
func (o *RedelegationsParams) SetPaginationLimit(paginationLimit *string) {
	o.PaginationLimit = paginationLimit
}

// WithPaginationOffset adds the paginationOffset to the redelegations params
func (o *RedelegationsParams) WithPaginationOffset(paginationOffset *string) *RedelegationsParams {
	o.SetPaginationOffset(paginationOffset)
	return o
}

// SetPaginationOffset adds the paginationOffset to the redelegations params
func (o *RedelegationsParams) SetPaginationOffset(paginationOffset *string) {
	o.PaginationOffset = paginationOffset
}

// WithPaginationReverse adds the paginationReverse to the redelegations params
func (o *RedelegationsParams) WithPaginationReverse(paginationReverse *bool) *RedelegationsParams {
	o.SetPaginationReverse(paginationReverse)
	return o
}

// SetPaginationReverse adds the paginationReverse to the redelegations params
func (o *RedelegationsParams) SetPaginationReverse(paginationReverse *bool) {
	o.PaginationReverse = paginationReverse
}

// WithSrcValidatorAddr adds the srcValidatorAddr to the redelegations params
func (o *RedelegationsParams) WithSrcValidatorAddr(srcValidatorAddr *string) *RedelegationsParams {
	o.SetSrcValidatorAddr(srcValidatorAddr)
	return o
}

// SetSrcValidatorAddr adds the srcValidatorAddr to the redelegations params
func (o *RedelegationsParams) SetSrcValidatorAddr(srcValidatorAddr *string) {
	o.SrcValidatorAddr = srcValidatorAddr
}

// WriteToRequest writes these params to a swagger request
func (o *RedelegationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param delegator_addr
	if err := r.SetPathParam("delegator_addr", o.DelegatorAddr); err != nil {
		return err
	}

	if o.DstValidatorAddr != nil {

		// query param dst_validator_addr
		var qrDstValidatorAddr string

		if o.DstValidatorAddr != nil {
			qrDstValidatorAddr = *o.DstValidatorAddr
		}
		qDstValidatorAddr := qrDstValidatorAddr
		if qDstValidatorAddr != "" {

			if err := r.SetQueryParam("dst_validator_addr", qDstValidatorAddr); err != nil {
				return err
			}
		}
	}

	if o.PaginationCountTotal != nil {

		// query param pagination.count_total
		var qrPaginationCountTotal bool

		if o.PaginationCountTotal != nil {
			qrPaginationCountTotal = *o.PaginationCountTotal
		}
		qPaginationCountTotal := swag.FormatBool(qrPaginationCountTotal)
		if qPaginationCountTotal != "" {

			if err := r.SetQueryParam("pagination.count_total", qPaginationCountTotal); err != nil {
				return err
			}
		}
	}

	if o.PaginationKey != nil {

		// query param pagination.key
		var qrPaginationKey strfmt.Base64

		if o.PaginationKey != nil {
			qrPaginationKey = *o.PaginationKey
		}
		qPaginationKey := qrPaginationKey.String()
		if qPaginationKey != "" {

			if err := r.SetQueryParam("pagination.key", qPaginationKey); err != nil {
				return err
			}
		}
	}

	if o.PaginationLimit != nil {

		// query param pagination.limit
		var qrPaginationLimit string

		if o.PaginationLimit != nil {
			qrPaginationLimit = *o.PaginationLimit
		}
		qPaginationLimit := qrPaginationLimit
		if qPaginationLimit != "" {

			if err := r.SetQueryParam("pagination.limit", qPaginationLimit); err != nil {
				return err
			}
		}
	}

	if o.PaginationOffset != nil {

		// query param pagination.offset
		var qrPaginationOffset string

		if o.PaginationOffset != nil {
			qrPaginationOffset = *o.PaginationOffset
		}
		qPaginationOffset := qrPaginationOffset
		if qPaginationOffset != "" {

			if err := r.SetQueryParam("pagination.offset", qPaginationOffset); err != nil {
				return err
			}
		}
	}

	if o.PaginationReverse != nil {

		// query param pagination.reverse
		var qrPaginationReverse bool

		if o.PaginationReverse != nil {
			qrPaginationReverse = *o.PaginationReverse
		}
		qPaginationReverse := swag.FormatBool(qrPaginationReverse)
		if qPaginationReverse != "" {

			if err := r.SetQueryParam("pagination.reverse", qPaginationReverse); err != nil {
				return err
			}
		}
	}

	if o.SrcValidatorAddr != nil {

		// query param src_validator_addr
		var qrSrcValidatorAddr string

		if o.SrcValidatorAddr != nil {
			qrSrcValidatorAddr = *o.SrcValidatorAddr
		}
		qSrcValidatorAddr := qrSrcValidatorAddr
		if qSrcValidatorAddr != "" {

			if err := r.SetQueryParam("src_validator_addr", qSrcValidatorAddr); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

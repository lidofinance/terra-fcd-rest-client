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

// NewGrantsParams creates a new GrantsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGrantsParams() *GrantsParams {
	return &GrantsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGrantsParamsWithTimeout creates a new GrantsParams object
// with the ability to set a timeout on a request.
func NewGrantsParamsWithTimeout(timeout time.Duration) *GrantsParams {
	return &GrantsParams{
		timeout: timeout,
	}
}

// NewGrantsParamsWithContext creates a new GrantsParams object
// with the ability to set a context for a request.
func NewGrantsParamsWithContext(ctx context.Context) *GrantsParams {
	return &GrantsParams{
		Context: ctx,
	}
}

// NewGrantsParamsWithHTTPClient creates a new GrantsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGrantsParamsWithHTTPClient(client *http.Client) *GrantsParams {
	return &GrantsParams{
		HTTPClient: client,
	}
}

/* GrantsParams contains all the parameters to send to the API endpoint
   for the grants operation.

   Typically these are written to a http.Request.
*/
type GrantsParams struct {

	// Grantee.
	Grantee *string

	// Granter.
	Granter *string

	/* MsgTypeURL.

	   Optional, msg_type_url, when set, will query only grants matching given msg type.
	*/
	MsgTypeURL *string

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the grants params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GrantsParams) WithDefaults() *GrantsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the grants params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GrantsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the grants params
func (o *GrantsParams) WithTimeout(timeout time.Duration) *GrantsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the grants params
func (o *GrantsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the grants params
func (o *GrantsParams) WithContext(ctx context.Context) *GrantsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the grants params
func (o *GrantsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the grants params
func (o *GrantsParams) WithHTTPClient(client *http.Client) *GrantsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the grants params
func (o *GrantsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGrantee adds the grantee to the grants params
func (o *GrantsParams) WithGrantee(grantee *string) *GrantsParams {
	o.SetGrantee(grantee)
	return o
}

// SetGrantee adds the grantee to the grants params
func (o *GrantsParams) SetGrantee(grantee *string) {
	o.Grantee = grantee
}

// WithGranter adds the granter to the grants params
func (o *GrantsParams) WithGranter(granter *string) *GrantsParams {
	o.SetGranter(granter)
	return o
}

// SetGranter adds the granter to the grants params
func (o *GrantsParams) SetGranter(granter *string) {
	o.Granter = granter
}

// WithMsgTypeURL adds the msgTypeURL to the grants params
func (o *GrantsParams) WithMsgTypeURL(msgTypeURL *string) *GrantsParams {
	o.SetMsgTypeURL(msgTypeURL)
	return o
}

// SetMsgTypeURL adds the msgTypeUrl to the grants params
func (o *GrantsParams) SetMsgTypeURL(msgTypeURL *string) {
	o.MsgTypeURL = msgTypeURL
}

// WithPaginationCountTotal adds the paginationCountTotal to the grants params
func (o *GrantsParams) WithPaginationCountTotal(paginationCountTotal *bool) *GrantsParams {
	o.SetPaginationCountTotal(paginationCountTotal)
	return o
}

// SetPaginationCountTotal adds the paginationCountTotal to the grants params
func (o *GrantsParams) SetPaginationCountTotal(paginationCountTotal *bool) {
	o.PaginationCountTotal = paginationCountTotal
}

// WithPaginationKey adds the paginationKey to the grants params
func (o *GrantsParams) WithPaginationKey(paginationKey *strfmt.Base64) *GrantsParams {
	o.SetPaginationKey(paginationKey)
	return o
}

// SetPaginationKey adds the paginationKey to the grants params
func (o *GrantsParams) SetPaginationKey(paginationKey *strfmt.Base64) {
	o.PaginationKey = paginationKey
}

// WithPaginationLimit adds the paginationLimit to the grants params
func (o *GrantsParams) WithPaginationLimit(paginationLimit *string) *GrantsParams {
	o.SetPaginationLimit(paginationLimit)
	return o
}

// SetPaginationLimit adds the paginationLimit to the grants params
func (o *GrantsParams) SetPaginationLimit(paginationLimit *string) {
	o.PaginationLimit = paginationLimit
}

// WithPaginationOffset adds the paginationOffset to the grants params
func (o *GrantsParams) WithPaginationOffset(paginationOffset *string) *GrantsParams {
	o.SetPaginationOffset(paginationOffset)
	return o
}

// SetPaginationOffset adds the paginationOffset to the grants params
func (o *GrantsParams) SetPaginationOffset(paginationOffset *string) {
	o.PaginationOffset = paginationOffset
}

// WithPaginationReverse adds the paginationReverse to the grants params
func (o *GrantsParams) WithPaginationReverse(paginationReverse *bool) *GrantsParams {
	o.SetPaginationReverse(paginationReverse)
	return o
}

// SetPaginationReverse adds the paginationReverse to the grants params
func (o *GrantsParams) SetPaginationReverse(paginationReverse *bool) {
	o.PaginationReverse = paginationReverse
}

// WriteToRequest writes these params to a swagger request
func (o *GrantsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Grantee != nil {

		// query param grantee
		var qrGrantee string

		if o.Grantee != nil {
			qrGrantee = *o.Grantee
		}
		qGrantee := qrGrantee
		if qGrantee != "" {

			if err := r.SetQueryParam("grantee", qGrantee); err != nil {
				return err
			}
		}
	}

	if o.Granter != nil {

		// query param granter
		var qrGranter string

		if o.Granter != nil {
			qrGranter = *o.Granter
		}
		qGranter := qrGranter
		if qGranter != "" {

			if err := r.SetQueryParam("granter", qGranter); err != nil {
				return err
			}
		}
	}

	if o.MsgTypeURL != nil {

		// query param msg_type_url
		var qrMsgTypeURL string

		if o.MsgTypeURL != nil {
			qrMsgTypeURL = *o.MsgTypeURL
		}
		qMsgTypeURL := qrMsgTypeURL
		if qMsgTypeURL != "" {

			if err := r.SetQueryParam("msg_type_url", qMsgTypeURL); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

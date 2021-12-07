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

// NewValidatorSlashesParams creates a new ValidatorSlashesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewValidatorSlashesParams() *ValidatorSlashesParams {
	return &ValidatorSlashesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewValidatorSlashesParamsWithTimeout creates a new ValidatorSlashesParams object
// with the ability to set a timeout on a request.
func NewValidatorSlashesParamsWithTimeout(timeout time.Duration) *ValidatorSlashesParams {
	return &ValidatorSlashesParams{
		timeout: timeout,
	}
}

// NewValidatorSlashesParamsWithContext creates a new ValidatorSlashesParams object
// with the ability to set a context for a request.
func NewValidatorSlashesParamsWithContext(ctx context.Context) *ValidatorSlashesParams {
	return &ValidatorSlashesParams{
		Context: ctx,
	}
}

// NewValidatorSlashesParamsWithHTTPClient creates a new ValidatorSlashesParams object
// with the ability to set a custom HTTPClient for a request.
func NewValidatorSlashesParamsWithHTTPClient(client *http.Client) *ValidatorSlashesParams {
	return &ValidatorSlashesParams{
		HTTPClient: client,
	}
}

/* ValidatorSlashesParams contains all the parameters to send to the API endpoint
   for the validator slashes operation.

   Typically these are written to a http.Request.
*/
type ValidatorSlashesParams struct {

	/* EndingHeight.

	   starting_height defines the optional ending height to query the slashes.

	   Format: uint64
	*/
	EndingHeight *string

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

	/* StartingHeight.

	   starting_height defines the optional starting height to query the slashes.

	   Format: uint64
	*/
	StartingHeight *string

	/* ValidatorAddress.

	   validator_address defines the validator address to query for.
	*/
	ValidatorAddress string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the validator slashes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidatorSlashesParams) WithDefaults() *ValidatorSlashesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the validator slashes params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ValidatorSlashesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the validator slashes params
func (o *ValidatorSlashesParams) WithTimeout(timeout time.Duration) *ValidatorSlashesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the validator slashes params
func (o *ValidatorSlashesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the validator slashes params
func (o *ValidatorSlashesParams) WithContext(ctx context.Context) *ValidatorSlashesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the validator slashes params
func (o *ValidatorSlashesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the validator slashes params
func (o *ValidatorSlashesParams) WithHTTPClient(client *http.Client) *ValidatorSlashesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the validator slashes params
func (o *ValidatorSlashesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEndingHeight adds the endingHeight to the validator slashes params
func (o *ValidatorSlashesParams) WithEndingHeight(endingHeight *string) *ValidatorSlashesParams {
	o.SetEndingHeight(endingHeight)
	return o
}

// SetEndingHeight adds the endingHeight to the validator slashes params
func (o *ValidatorSlashesParams) SetEndingHeight(endingHeight *string) {
	o.EndingHeight = endingHeight
}

// WithPaginationCountTotal adds the paginationCountTotal to the validator slashes params
func (o *ValidatorSlashesParams) WithPaginationCountTotal(paginationCountTotal *bool) *ValidatorSlashesParams {
	o.SetPaginationCountTotal(paginationCountTotal)
	return o
}

// SetPaginationCountTotal adds the paginationCountTotal to the validator slashes params
func (o *ValidatorSlashesParams) SetPaginationCountTotal(paginationCountTotal *bool) {
	o.PaginationCountTotal = paginationCountTotal
}

// WithPaginationKey adds the paginationKey to the validator slashes params
func (o *ValidatorSlashesParams) WithPaginationKey(paginationKey *strfmt.Base64) *ValidatorSlashesParams {
	o.SetPaginationKey(paginationKey)
	return o
}

// SetPaginationKey adds the paginationKey to the validator slashes params
func (o *ValidatorSlashesParams) SetPaginationKey(paginationKey *strfmt.Base64) {
	o.PaginationKey = paginationKey
}

// WithPaginationLimit adds the paginationLimit to the validator slashes params
func (o *ValidatorSlashesParams) WithPaginationLimit(paginationLimit *string) *ValidatorSlashesParams {
	o.SetPaginationLimit(paginationLimit)
	return o
}

// SetPaginationLimit adds the paginationLimit to the validator slashes params
func (o *ValidatorSlashesParams) SetPaginationLimit(paginationLimit *string) {
	o.PaginationLimit = paginationLimit
}

// WithPaginationOffset adds the paginationOffset to the validator slashes params
func (o *ValidatorSlashesParams) WithPaginationOffset(paginationOffset *string) *ValidatorSlashesParams {
	o.SetPaginationOffset(paginationOffset)
	return o
}

// SetPaginationOffset adds the paginationOffset to the validator slashes params
func (o *ValidatorSlashesParams) SetPaginationOffset(paginationOffset *string) {
	o.PaginationOffset = paginationOffset
}

// WithPaginationReverse adds the paginationReverse to the validator slashes params
func (o *ValidatorSlashesParams) WithPaginationReverse(paginationReverse *bool) *ValidatorSlashesParams {
	o.SetPaginationReverse(paginationReverse)
	return o
}

// SetPaginationReverse adds the paginationReverse to the validator slashes params
func (o *ValidatorSlashesParams) SetPaginationReverse(paginationReverse *bool) {
	o.PaginationReverse = paginationReverse
}

// WithStartingHeight adds the startingHeight to the validator slashes params
func (o *ValidatorSlashesParams) WithStartingHeight(startingHeight *string) *ValidatorSlashesParams {
	o.SetStartingHeight(startingHeight)
	return o
}

// SetStartingHeight adds the startingHeight to the validator slashes params
func (o *ValidatorSlashesParams) SetStartingHeight(startingHeight *string) {
	o.StartingHeight = startingHeight
}

// WithValidatorAddress adds the validatorAddress to the validator slashes params
func (o *ValidatorSlashesParams) WithValidatorAddress(validatorAddress string) *ValidatorSlashesParams {
	o.SetValidatorAddress(validatorAddress)
	return o
}

// SetValidatorAddress adds the validatorAddress to the validator slashes params
func (o *ValidatorSlashesParams) SetValidatorAddress(validatorAddress string) {
	o.ValidatorAddress = validatorAddress
}

// WriteToRequest writes these params to a swagger request
func (o *ValidatorSlashesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EndingHeight != nil {

		// query param ending_height
		var qrEndingHeight string

		if o.EndingHeight != nil {
			qrEndingHeight = *o.EndingHeight
		}
		qEndingHeight := qrEndingHeight
		if qEndingHeight != "" {

			if err := r.SetQueryParam("ending_height", qEndingHeight); err != nil {
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

	if o.StartingHeight != nil {

		// query param starting_height
		var qrStartingHeight string

		if o.StartingHeight != nil {
			qrStartingHeight = *o.StartingHeight
		}
		qStartingHeight := qrStartingHeight
		if qStartingHeight != "" {

			if err := r.SetQueryParam("starting_height", qStartingHeight); err != nil {
				return err
			}
		}
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

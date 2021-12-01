// Code generated by go-swagger; DO NOT EDIT.

package governance

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

// NewGetV1GovProposalsParams creates a new GetV1GovProposalsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetV1GovProposalsParams() *GetV1GovProposalsParams {
	return &GetV1GovProposalsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetV1GovProposalsParamsWithTimeout creates a new GetV1GovProposalsParams object
// with the ability to set a timeout on a request.
func NewGetV1GovProposalsParamsWithTimeout(timeout time.Duration) *GetV1GovProposalsParams {
	return &GetV1GovProposalsParams{
		timeout: timeout,
	}
}

// NewGetV1GovProposalsParamsWithContext creates a new GetV1GovProposalsParams object
// with the ability to set a context for a request.
func NewGetV1GovProposalsParamsWithContext(ctx context.Context) *GetV1GovProposalsParams {
	return &GetV1GovProposalsParams{
		Context: ctx,
	}
}

// NewGetV1GovProposalsParamsWithHTTPClient creates a new GetV1GovProposalsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetV1GovProposalsParamsWithHTTPClient(client *http.Client) *GetV1GovProposalsParams {
	return &GetV1GovProposalsParams{
		HTTPClient: client,
	}
}

/* GetV1GovProposalsParams contains all the parameters to send to the API endpoint
   for the get v1 gov proposals operation.

   Typically these are written to a http.Request.
*/
type GetV1GovProposalsParams struct {

	/* Status.

	   'deposit', 'voting', 'passed', 'rejected'
	*/
	Status *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get v1 gov proposals params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1GovProposalsParams) WithDefaults() *GetV1GovProposalsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get v1 gov proposals params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetV1GovProposalsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get v1 gov proposals params
func (o *GetV1GovProposalsParams) WithTimeout(timeout time.Duration) *GetV1GovProposalsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get v1 gov proposals params
func (o *GetV1GovProposalsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get v1 gov proposals params
func (o *GetV1GovProposalsParams) WithContext(ctx context.Context) *GetV1GovProposalsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get v1 gov proposals params
func (o *GetV1GovProposalsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get v1 gov proposals params
func (o *GetV1GovProposalsParams) WithHTTPClient(client *http.Client) *GetV1GovProposalsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get v1 gov proposals params
func (o *GetV1GovProposalsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithStatus adds the status to the get v1 gov proposals params
func (o *GetV1GovProposalsParams) WithStatus(status *string) *GetV1GovProposalsParams {
	o.SetStatus(status)
	return o
}

// SetStatus adds the status to the get v1 gov proposals params
func (o *GetV1GovProposalsParams) SetStatus(status *string) {
	o.Status = status
}

// WriteToRequest writes these params to a swagger request
func (o *GetV1GovProposalsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Status != nil {

		// query param status
		var qrStatus string

		if o.Status != nil {
			qrStatus = *o.Status
		}
		qStatus := qrStatus
		if qStatus != "" {

			if err := r.SetQueryParam("status", qStatus); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
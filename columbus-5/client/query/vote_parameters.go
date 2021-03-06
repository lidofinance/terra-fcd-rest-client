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

// NewVoteParams creates a new VoteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVoteParams() *VoteParams {
	return &VoteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVoteParamsWithTimeout creates a new VoteParams object
// with the ability to set a timeout on a request.
func NewVoteParamsWithTimeout(timeout time.Duration) *VoteParams {
	return &VoteParams{
		timeout: timeout,
	}
}

// NewVoteParamsWithContext creates a new VoteParams object
// with the ability to set a context for a request.
func NewVoteParamsWithContext(ctx context.Context) *VoteParams {
	return &VoteParams{
		Context: ctx,
	}
}

// NewVoteParamsWithHTTPClient creates a new VoteParams object
// with the ability to set a custom HTTPClient for a request.
func NewVoteParamsWithHTTPClient(client *http.Client) *VoteParams {
	return &VoteParams{
		HTTPClient: client,
	}
}

/* VoteParams contains all the parameters to send to the API endpoint
   for the vote operation.

   Typically these are written to a http.Request.
*/
type VoteParams struct {

	/* ProposalID.

	   proposal_id defines the unique id of the proposal.

	   Format: uint64
	*/
	ProposalID string

	/* Voter.

	   voter defines the oter address for the proposals.
	*/
	Voter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the vote params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VoteParams) WithDefaults() *VoteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the vote params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VoteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the vote params
func (o *VoteParams) WithTimeout(timeout time.Duration) *VoteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the vote params
func (o *VoteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the vote params
func (o *VoteParams) WithContext(ctx context.Context) *VoteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the vote params
func (o *VoteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the vote params
func (o *VoteParams) WithHTTPClient(client *http.Client) *VoteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the vote params
func (o *VoteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProposalID adds the proposalID to the vote params
func (o *VoteParams) WithProposalID(proposalID string) *VoteParams {
	o.SetProposalID(proposalID)
	return o
}

// SetProposalID adds the proposalId to the vote params
func (o *VoteParams) SetProposalID(proposalID string) {
	o.ProposalID = proposalID
}

// WithVoter adds the voter to the vote params
func (o *VoteParams) WithVoter(voter string) *VoteParams {
	o.SetVoter(voter)
	return o
}

// SetVoter adds the voter to the vote params
func (o *VoteParams) SetVoter(voter string) {
	o.Voter = voter
}

// WriteToRequest writes these params to a swagger request
func (o *VoteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param proposal_id
	if err := r.SetPathParam("proposal_id", o.ProposalID); err != nil {
		return err
	}

	// path param voter
	if err := r.SetPathParam("voter", o.Voter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

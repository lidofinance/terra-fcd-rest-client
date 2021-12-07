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

// NewChannelParams creates a new ChannelParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewChannelParams() *ChannelParams {
	return &ChannelParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewChannelParamsWithTimeout creates a new ChannelParams object
// with the ability to set a timeout on a request.
func NewChannelParamsWithTimeout(timeout time.Duration) *ChannelParams {
	return &ChannelParams{
		timeout: timeout,
	}
}

// NewChannelParamsWithContext creates a new ChannelParams object
// with the ability to set a context for a request.
func NewChannelParamsWithContext(ctx context.Context) *ChannelParams {
	return &ChannelParams{
		Context: ctx,
	}
}

// NewChannelParamsWithHTTPClient creates a new ChannelParams object
// with the ability to set a custom HTTPClient for a request.
func NewChannelParamsWithHTTPClient(client *http.Client) *ChannelParams {
	return &ChannelParams{
		HTTPClient: client,
	}
}

/* ChannelParams contains all the parameters to send to the API endpoint
   for the channel operation.

   Typically these are written to a http.Request.
*/
type ChannelParams struct {

	/* ChannelID.

	   channel unique identifier
	*/
	ChannelID string

	/* PortID.

	   port unique identifier
	*/
	PortID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the channel params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChannelParams) WithDefaults() *ChannelParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the channel params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChannelParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the channel params
func (o *ChannelParams) WithTimeout(timeout time.Duration) *ChannelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the channel params
func (o *ChannelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the channel params
func (o *ChannelParams) WithContext(ctx context.Context) *ChannelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the channel params
func (o *ChannelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the channel params
func (o *ChannelParams) WithHTTPClient(client *http.Client) *ChannelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the channel params
func (o *ChannelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithChannelID adds the channelID to the channel params
func (o *ChannelParams) WithChannelID(channelID string) *ChannelParams {
	o.SetChannelID(channelID)
	return o
}

// SetChannelID adds the channelId to the channel params
func (o *ChannelParams) SetChannelID(channelID string) {
	o.ChannelID = channelID
}

// WithPortID adds the portID to the channel params
func (o *ChannelParams) WithPortID(portID string) *ChannelParams {
	o.SetPortID(portID)
	return o
}

// SetPortID adds the portId to the channel params
func (o *ChannelParams) SetPortID(portID string) {
	o.PortID = portID
}

// WriteToRequest writes these params to a swagger request
func (o *ChannelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param channel_id
	if err := r.SetPathParam("channel_id", o.ChannelID); err != nil {
		return err
	}

	// path param port_id
	if err := r.SetPathParam("port_id", o.PortID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

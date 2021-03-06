// Code generated by go-swagger; DO NOT EDIT.

package market

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

// NewGetMarketSwapParams creates a new GetMarketSwapParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMarketSwapParams() *GetMarketSwapParams {
	return &GetMarketSwapParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMarketSwapParamsWithTimeout creates a new GetMarketSwapParams object
// with the ability to set a timeout on a request.
func NewGetMarketSwapParamsWithTimeout(timeout time.Duration) *GetMarketSwapParams {
	return &GetMarketSwapParams{
		timeout: timeout,
	}
}

// NewGetMarketSwapParamsWithContext creates a new GetMarketSwapParams object
// with the ability to set a context for a request.
func NewGetMarketSwapParamsWithContext(ctx context.Context) *GetMarketSwapParams {
	return &GetMarketSwapParams{
		Context: ctx,
	}
}

// NewGetMarketSwapParamsWithHTTPClient creates a new GetMarketSwapParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetMarketSwapParamsWithHTTPClient(client *http.Client) *GetMarketSwapParams {
	return &GetMarketSwapParams{
		HTTPClient: client,
	}
}

/* GetMarketSwapParams contains all the parameters to send to the API endpoint
   for the get market swap operation.

   Typically these are written to a http.Request.
*/
type GetMarketSwapParams struct {

	/* AskDenom.

	   Then coin denom want to ask
	*/
	AskDenom string

	/* OfferCoin.

	   coin expression want to swap
	*/
	OfferCoin string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get market swap params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMarketSwapParams) WithDefaults() *GetMarketSwapParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get market swap params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMarketSwapParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get market swap params
func (o *GetMarketSwapParams) WithTimeout(timeout time.Duration) *GetMarketSwapParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get market swap params
func (o *GetMarketSwapParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get market swap params
func (o *GetMarketSwapParams) WithContext(ctx context.Context) *GetMarketSwapParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get market swap params
func (o *GetMarketSwapParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get market swap params
func (o *GetMarketSwapParams) WithHTTPClient(client *http.Client) *GetMarketSwapParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get market swap params
func (o *GetMarketSwapParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAskDenom adds the askDenom to the get market swap params
func (o *GetMarketSwapParams) WithAskDenom(askDenom string) *GetMarketSwapParams {
	o.SetAskDenom(askDenom)
	return o
}

// SetAskDenom adds the askDenom to the get market swap params
func (o *GetMarketSwapParams) SetAskDenom(askDenom string) {
	o.AskDenom = askDenom
}

// WithOfferCoin adds the offerCoin to the get market swap params
func (o *GetMarketSwapParams) WithOfferCoin(offerCoin string) *GetMarketSwapParams {
	o.SetOfferCoin(offerCoin)
	return o
}

// SetOfferCoin adds the offerCoin to the get market swap params
func (o *GetMarketSwapParams) SetOfferCoin(offerCoin string) {
	o.OfferCoin = offerCoin
}

// WriteToRequest writes these params to a swagger request
func (o *GetMarketSwapParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param ask_denom
	qrAskDenom := o.AskDenom
	qAskDenom := qrAskDenom
	if qAskDenom != "" {

		if err := r.SetQueryParam("ask_denom", qAskDenom); err != nil {
			return err
		}
	}

	// query param offer_coin
	qrOfferCoin := o.OfferCoin
	qOfferCoin := qrOfferCoin
	if qOfferCoin != "" {

		if err := r.SetQueryParam("offer_coin", qOfferCoin); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

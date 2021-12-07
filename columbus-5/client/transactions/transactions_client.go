// Code generated by go-swagger; DO NOT EDIT.

package transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new transactions API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for transactions API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetTxs(params *GetTxsParams, opts ...ClientOption) (*GetTxsOK, error)

	GetTxsHash(params *GetTxsHashParams, opts ...ClientOption) (*GetTxsHashOK, error)

	GetV1Mempool(params *GetV1MempoolParams, opts ...ClientOption) (*GetV1MempoolOK, error)

	GetV1MempoolTxhash(params *GetV1MempoolTxhashParams, opts ...ClientOption) (*GetV1MempoolTxhashOK, error)

	GetV1TxTxhash(params *GetV1TxTxhashParams, opts ...ClientOption) (*GetV1TxTxhashOK, error)

	GetV1Txs(params *GetV1TxsParams, opts ...ClientOption) (*GetV1TxsOK, error)

	GetV1TxsGasPrices(params *GetV1TxsGasPricesParams, opts ...ClientOption) (*GetV1TxsGasPricesOK, error)

	PostTxs(params *PostTxsParams, opts ...ClientOption) (*PostTxsOK, error)

	PostTxsDecode(params *PostTxsDecodeParams, opts ...ClientOption) (*PostTxsDecodeOK, error)

	PostTxsEncode(params *PostTxsEncodeParams, opts ...ClientOption) (*PostTxsEncodeOK, error)

	PostTxsEstimateFee(params *PostTxsEstimateFeeParams, opts ...ClientOption) (*PostTxsEstimateFeeOK, error)

	PostV1Txs(params *PostV1TxsParams, opts ...ClientOption) (*PostV1TxsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetTxs searches transactions

  Search transactions by events.
*/
func (a *Client) GetTxs(params *GetTxsParams, opts ...ClientOption) (*GetTxsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTxsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetTxs",
		Method:             "GET",
		PathPattern:        "/txs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTxsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTxsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetTxs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetTxsHash gets a tx by hash

  Retrieve a transaction using its hash.
*/
func (a *Client) GetTxsHash(params *GetTxsHashParams, opts ...ClientOption) (*GetTxsHashOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTxsHashParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetTxsHash",
		Method:             "GET",
		PathPattern:        "/txs/{hash}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTxsHashReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTxsHashOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetTxsHash: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetV1Mempool gets transactions in mempool

  Get transactions in mempool
*/
func (a *Client) GetV1Mempool(params *GetV1MempoolParams, opts ...ClientOption) (*GetV1MempoolOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1MempoolParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1Mempool",
		Method:             "GET",
		PathPattern:        "/v1/mempool",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetV1MempoolReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetV1MempoolOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1Mempool: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetV1MempoolTxhash gets transaction in mempool

  Get transaction in mempool
*/
func (a *Client) GetV1MempoolTxhash(params *GetV1MempoolTxhashParams, opts ...ClientOption) (*GetV1MempoolTxhashOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1MempoolTxhashParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1MempoolTxhash",
		Method:             "GET",
		PathPattern:        "/v1/mempool/{txhash}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetV1MempoolTxhashReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetV1MempoolTxhashOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1MempoolTxhash: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetV1TxTxhash gets tx

  Get Tx
*/
func (a *Client) GetV1TxTxhash(params *GetV1TxTxhashParams, opts ...ClientOption) (*GetV1TxTxhashOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1TxTxhashParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1TxTxhash",
		Method:             "GET",
		PathPattern:        "/v1/tx/{txhash}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetV1TxTxhashReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetV1TxTxhashOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1TxTxhash: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetV1Txs gets tx list

  Get Tx List
*/
func (a *Client) GetV1Txs(params *GetV1TxsParams, opts ...ClientOption) (*GetV1TxsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1TxsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1Txs",
		Method:             "GET",
		PathPattern:        "/v1/txs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetV1TxsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetV1TxsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1Txs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetV1TxsGasPrices gets gas prices

  Get gas prices
*/
func (a *Client) GetV1TxsGasPrices(params *GetV1TxsGasPricesParams, opts ...ClientOption) (*GetV1TxsGasPricesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1TxsGasPricesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1TxsGasPrices",
		Method:             "GET",
		PathPattern:        "/v1/txs/gas_prices",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetV1TxsGasPricesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetV1TxsGasPricesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1TxsGasPrices: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostTxs broadcasts a signed tx

  Broadcast a signed tx to a full node
*/
func (a *Client) PostTxs(params *PostTxsParams, opts ...ClientOption) (*PostTxsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostTxsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostTxs",
		Method:             "POST",
		PathPattern:        "/txs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostTxsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostTxsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostTxs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostTxsDecode decodes a transaction from the amino wire format

  Decode a transaction (signed or not) from base64-encoded Amino serialized bytes to JSON
*/
func (a *Client) PostTxsDecode(params *PostTxsDecodeParams, opts ...ClientOption) (*PostTxsDecodeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostTxsDecodeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostTxsDecode",
		Method:             "POST",
		PathPattern:        "/txs/decode",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostTxsDecodeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostTxsDecodeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostTxsDecode: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostTxsEncode encodes a legacy transaction to the proto wire format

  Encode a legacy transaction (signed or not) from JSON to base64-encoded Proto serialized bytes
*/
func (a *Client) PostTxsEncode(params *PostTxsEncodeParams, opts ...ClientOption) (*PostTxsEncodeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostTxsEncodeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostTxsEncode",
		Method:             "POST",
		PathPattern:        "/txs/encode",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostTxsEncodeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostTxsEncodeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostTxsEncode: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostTxsEstimateFee estimates fee and gas of a transaction

  Estimate fee and gas of a transaction according to given parameters
*/
func (a *Client) PostTxsEstimateFee(params *PostTxsEstimateFeeParams, opts ...ClientOption) (*PostTxsEstimateFeeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostTxsEstimateFeeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostTxsEstimateFee",
		Method:             "POST",
		PathPattern:        "/txs/estimate_fee",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostTxsEstimateFeeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostTxsEstimateFeeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostTxsEstimateFee: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PostV1Txs broadcasts txs

  Broadcast Txs
*/
func (a *Client) PostV1Txs(params *PostV1TxsParams, opts ...ClientOption) (*PostV1TxsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostV1TxsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostV1Txs",
		Method:             "POST",
		PathPattern:        "/v1/txs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostV1TxsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostV1TxsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for PostV1Txs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

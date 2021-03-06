// Code generated by go-swagger; DO NOT EDIT.

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	BroadcastTx(params *BroadcastTxParams, opts ...ClientOption) (*BroadcastTxOK, error)

	ComputeTax(params *ComputeTaxParams, opts ...ClientOption) (*ComputeTaxOK, error)

	GetBlockByHeight(params *GetBlockByHeightParams, opts ...ClientOption) (*GetBlockByHeightOK, error)

	GetLatestBlock(params *GetLatestBlockParams, opts ...ClientOption) (*GetLatestBlockOK, error)

	GetLatestValidatorSet(params *GetLatestValidatorSetParams, opts ...ClientOption) (*GetLatestValidatorSetOK, error)

	GetTx(params *GetTxParams, opts ...ClientOption) (*GetTxOK, error)

	GetTxsEvent(params *GetTxsEventParams, opts ...ClientOption) (*GetTxsEventOK, error)

	GetValidatorSetByHeight(params *GetValidatorSetByHeightParams, opts ...ClientOption) (*GetValidatorSetByHeightOK, error)

	Simulate(params *SimulateParams, opts ...ClientOption) (*SimulateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  BroadcastTx broadcasts tx broadcast transaction
*/
func (a *Client) BroadcastTx(params *BroadcastTxParams, opts ...ClientOption) (*BroadcastTxOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBroadcastTxParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BroadcastTx",
		Method:             "POST",
		PathPattern:        "/cosmos/tx/v1beta1/txs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &BroadcastTxReader{formats: a.formats},
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
	success, ok := result.(*BroadcastTxOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BroadcastTxDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ComputeTax estimates fee simulates executing a transaction for estimating gas usage
*/
func (a *Client) ComputeTax(params *ComputeTaxParams, opts ...ClientOption) (*ComputeTaxOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewComputeTaxParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ComputeTax",
		Method:             "POST",
		PathPattern:        "/terra/tx/v1beta1/compute_tax",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ComputeTaxReader{formats: a.formats},
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
	success, ok := result.(*ComputeTaxOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ComputeTaxDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetBlockByHeight gets block by height queries block for given height
*/
func (a *Client) GetBlockByHeight(params *GetBlockByHeightParams, opts ...ClientOption) (*GetBlockByHeightOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBlockByHeightParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetBlockByHeight",
		Method:             "GET",
		PathPattern:        "/cosmos/base/tendermint/v1beta1/blocks/{height}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBlockByHeightReader{formats: a.formats},
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
	success, ok := result.(*GetBlockByHeightOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetBlockByHeightDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetLatestBlock gets latest block returns the latest block
*/
func (a *Client) GetLatestBlock(params *GetLatestBlockParams, opts ...ClientOption) (*GetLatestBlockOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLatestBlockParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetLatestBlock",
		Method:             "GET",
		PathPattern:        "/cosmos/base/tendermint/v1beta1/blocks/latest",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetLatestBlockReader{formats: a.formats},
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
	success, ok := result.(*GetLatestBlockOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetLatestBlockDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetLatestValidatorSet gets latest validator set queries latest validator set
*/
func (a *Client) GetLatestValidatorSet(params *GetLatestValidatorSetParams, opts ...ClientOption) (*GetLatestValidatorSetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLatestValidatorSetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetLatestValidatorSet",
		Method:             "GET",
		PathPattern:        "/cosmos/base/tendermint/v1beta1/validatorsets/latest",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetLatestValidatorSetReader{formats: a.formats},
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
	success, ok := result.(*GetLatestValidatorSetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetLatestValidatorSetDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetTx gets tx fetches a tx by hash
*/
func (a *Client) GetTx(params *GetTxParams, opts ...ClientOption) (*GetTxOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTxParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetTx",
		Method:             "GET",
		PathPattern:        "/cosmos/tx/v1beta1/txs/{hash}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTxReader{formats: a.formats},
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
	success, ok := result.(*GetTxOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetTxDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetTxsEvent gets txs event fetches txs by event
*/
func (a *Client) GetTxsEvent(params *GetTxsEventParams, opts ...ClientOption) (*GetTxsEventOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTxsEventParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetTxsEvent",
		Method:             "GET",
		PathPattern:        "/cosmos/tx/v1beta1/txs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetTxsEventReader{formats: a.formats},
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
	success, ok := result.(*GetTxsEventOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetTxsEventDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  GetValidatorSetByHeight gets validator set by height queries validator set at a given height
*/
func (a *Client) GetValidatorSetByHeight(params *GetValidatorSetByHeightParams, opts ...ClientOption) (*GetValidatorSetByHeightOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetValidatorSetByHeightParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetValidatorSetByHeight",
		Method:             "GET",
		PathPattern:        "/cosmos/base/tendermint/v1beta1/validatorsets/{height}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetValidatorSetByHeightReader{formats: a.formats},
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
	success, ok := result.(*GetValidatorSetByHeightOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetValidatorSetByHeightDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  Simulate simulates simulates executing a transaction for estimating gas usage
*/
func (a *Client) Simulate(params *SimulateParams, opts ...ClientOption) (*SimulateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSimulateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Simulate",
		Method:             "POST",
		PathPattern:        "/cosmos/tx/v1beta1/simulate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SimulateReader{formats: a.formats},
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
	success, ok := result.(*SimulateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SimulateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

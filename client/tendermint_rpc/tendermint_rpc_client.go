// Code generated by go-swagger; DO NOT EDIT.

package tendermint_rpc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new tendermint rpc API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for tendermint rpc API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetBlocksHeight(params *GetBlocksHeightParams, opts ...ClientOption) (*GetBlocksHeightOK, error)

	GetBlocksLatest(params *GetBlocksLatestParams, opts ...ClientOption) (*GetBlocksLatestOK, error)

	GetSyncing(params *GetSyncingParams, opts ...ClientOption) (*GetSyncingOK, error)

	GetValidatorsetsHeight(params *GetValidatorsetsHeightParams, opts ...ClientOption) (*GetValidatorsetsHeightOK, error)

	GetValidatorsetsLatest(params *GetValidatorsetsLatestParams, opts ...ClientOption) (*GetValidatorsetsLatestOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetBlocksHeight gets a block at a certain height
*/
func (a *Client) GetBlocksHeight(params *GetBlocksHeightParams, opts ...ClientOption) (*GetBlocksHeightOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBlocksHeightParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetBlocksHeight",
		Method:             "GET",
		PathPattern:        "/blocks/{height}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBlocksHeightReader{formats: a.formats},
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
	success, ok := result.(*GetBlocksHeightOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetBlocksHeight: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetBlocksLatest gets the latest block
*/
func (a *Client) GetBlocksLatest(params *GetBlocksLatestParams, opts ...ClientOption) (*GetBlocksLatestOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetBlocksLatestParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetBlocksLatest",
		Method:             "GET",
		PathPattern:        "/blocks/latest",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetBlocksLatestReader{formats: a.formats},
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
	success, ok := result.(*GetBlocksLatestOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetBlocksLatest: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetSyncing syncings state of node

  Get if the node is currently syning with other nodes
*/
func (a *Client) GetSyncing(params *GetSyncingParams, opts ...ClientOption) (*GetSyncingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSyncingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetSyncing",
		Method:             "GET",
		PathPattern:        "/syncing",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSyncingReader{formats: a.formats},
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
	success, ok := result.(*GetSyncingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetSyncing: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetValidatorsetsHeight gets a validator set a certain height
*/
func (a *Client) GetValidatorsetsHeight(params *GetValidatorsetsHeightParams, opts ...ClientOption) (*GetValidatorsetsHeightOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetValidatorsetsHeightParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetValidatorsetsHeight",
		Method:             "GET",
		PathPattern:        "/validatorsets/{height}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetValidatorsetsHeightReader{formats: a.formats},
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
	success, ok := result.(*GetValidatorsetsHeightOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetValidatorsetsHeight: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetValidatorsetsLatest gets the latest validator set
*/
func (a *Client) GetValidatorsetsLatest(params *GetValidatorsetsLatestParams, opts ...ClientOption) (*GetValidatorsetsLatestOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetValidatorsetsLatestParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetValidatorsetsLatest",
		Method:             "GET",
		PathPattern:        "/validatorsets/latest",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetValidatorsetsLatestReader{formats: a.formats},
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
	success, ok := result.(*GetValidatorsetsLatestOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetValidatorsetsLatest: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
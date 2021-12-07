// Code generated by go-swagger; DO NOT EDIT.

package dashboard

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new dashboard API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for dashboard API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GetV1Dashboard(params *GetV1DashboardParams, opts ...ClientOption) (*GetV1DashboardOK, error)

	GetV1DashboardAccountGrowth(params *GetV1DashboardAccountGrowthParams, opts ...ClientOption) (*GetV1DashboardAccountGrowthOK, error)

	GetV1DashboardActiveAccounts(params *GetV1DashboardActiveAccountsParams, opts ...ClientOption) (*GetV1DashboardActiveAccountsOK, error)

	GetV1DashboardBlockRewards(params *GetV1DashboardBlockRewardsParams, opts ...ClientOption) (*GetV1DashboardBlockRewardsOK, error)

	GetV1DashboardLastHourOpsTxsCount(params *GetV1DashboardLastHourOpsTxsCountParams, opts ...ClientOption) (*GetV1DashboardLastHourOpsTxsCountOK, error)

	GetV1DashboardRegisteredAccounts(params *GetV1DashboardRegisteredAccountsParams, opts ...ClientOption) (*GetV1DashboardRegisteredAccountsOK, error)

	GetV1DashboardSeigniorageProceeds(params *GetV1DashboardSeigniorageProceedsParams, opts ...ClientOption) (*GetV1DashboardSeigniorageProceedsOK, error)

	GetV1DashboardStakingRatio(params *GetV1DashboardStakingRatioParams, opts ...ClientOption) (*GetV1DashboardStakingRatioOK, error)

	GetV1DashboardStakingReturn(params *GetV1DashboardStakingReturnParams, opts ...ClientOption) (*GetV1DashboardStakingReturnOK, error)

	GetV1DashboardTxVolume(params *GetV1DashboardTxVolumeParams, opts ...ClientOption) (*GetV1DashboardTxVolumeOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetV1Dashboard gets information to be used on the dashboard

  Get information to be used on the dashboard
*/
func (a *Client) GetV1Dashboard(params *GetV1DashboardParams, opts ...ClientOption) (*GetV1DashboardOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1DashboardParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1Dashboard",
		Method:             "GET",
		PathPattern:        "/v1/dashboard",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetV1DashboardReader{formats: a.formats},
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
	success, ok := result.(*GetV1DashboardOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1Dashboard: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetV1DashboardAccountGrowth gets account growth history

  Get account growth history
*/
func (a *Client) GetV1DashboardAccountGrowth(params *GetV1DashboardAccountGrowthParams, opts ...ClientOption) (*GetV1DashboardAccountGrowthOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1DashboardAccountGrowthParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1DashboardAccountGrowth",
		Method:             "GET",
		PathPattern:        "/v1/dashboard/account_growth",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetV1DashboardAccountGrowthReader{formats: a.formats},
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
	success, ok := result.(*GetV1DashboardAccountGrowthOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1DashboardAccountGrowth: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetV1DashboardActiveAccounts gets active accounts count history

  Get active accounts count history
*/
func (a *Client) GetV1DashboardActiveAccounts(params *GetV1DashboardActiveAccountsParams, opts ...ClientOption) (*GetV1DashboardActiveAccountsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1DashboardActiveAccountsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1DashboardActiveAccounts",
		Method:             "GET",
		PathPattern:        "/v1/dashboard/active_accounts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetV1DashboardActiveAccountsReader{formats: a.formats},
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
	success, ok := result.(*GetV1DashboardActiveAccountsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1DashboardActiveAccounts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetV1DashboardBlockRewards gets block reward history

  Get block reward history
*/
func (a *Client) GetV1DashboardBlockRewards(params *GetV1DashboardBlockRewardsParams, opts ...ClientOption) (*GetV1DashboardBlockRewardsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1DashboardBlockRewardsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1DashboardBlockRewards",
		Method:             "GET",
		PathPattern:        "/v1/dashboard/block_rewards",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetV1DashboardBlockRewardsReader{formats: a.formats},
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
	success, ok := result.(*GetV1DashboardBlockRewardsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1DashboardBlockRewards: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetV1DashboardLastHourOpsTxsCount get v1 dashboard last hour ops txs count API
*/
func (a *Client) GetV1DashboardLastHourOpsTxsCount(params *GetV1DashboardLastHourOpsTxsCountParams, opts ...ClientOption) (*GetV1DashboardLastHourOpsTxsCountOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1DashboardLastHourOpsTxsCountParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1DashboardLastHourOpsTxsCount",
		Method:             "GET",
		PathPattern:        "/v1/dashboard/last_hour_ops_txs_count",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetV1DashboardLastHourOpsTxsCountReader{formats: a.formats},
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
	success, ok := result.(*GetV1DashboardLastHourOpsTxsCountOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1DashboardLastHourOpsTxsCount: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetV1DashboardRegisteredAccounts gets registered accounts count history

  Get registered accounts count history
*/
func (a *Client) GetV1DashboardRegisteredAccounts(params *GetV1DashboardRegisteredAccountsParams, opts ...ClientOption) (*GetV1DashboardRegisteredAccountsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1DashboardRegisteredAccountsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1DashboardRegisteredAccounts",
		Method:             "GET",
		PathPattern:        "/v1/dashboard/registered_accounts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetV1DashboardRegisteredAccountsReader{formats: a.formats},
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
	success, ok := result.(*GetV1DashboardRegisteredAccountsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1DashboardRegisteredAccounts: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetV1DashboardSeigniorageProceeds gets the amount of seigniorage in the start of the day

  Get the amount of seigniorage in the start of the day
*/
func (a *Client) GetV1DashboardSeigniorageProceeds(params *GetV1DashboardSeigniorageProceedsParams, opts ...ClientOption) (*GetV1DashboardSeigniorageProceedsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1DashboardSeigniorageProceedsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1DashboardSeigniorageProceeds",
		Method:             "GET",
		PathPattern:        "/v1/dashboard/seigniorage_proceeds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetV1DashboardSeigniorageProceedsReader{formats: a.formats},
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
	success, ok := result.(*GetV1DashboardSeigniorageProceedsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1DashboardSeigniorageProceeds: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetV1DashboardStakingRatio gets the historical staking ratio

  Get the historical staking ratio
*/
func (a *Client) GetV1DashboardStakingRatio(params *GetV1DashboardStakingRatioParams, opts ...ClientOption) (*GetV1DashboardStakingRatioOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1DashboardStakingRatioParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1DashboardStakingRatio",
		Method:             "GET",
		PathPattern:        "/v1/dashboard/staking_ratio",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetV1DashboardStakingRatioReader{formats: a.formats},
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
	success, ok := result.(*GetV1DashboardStakingRatioOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1DashboardStakingRatio: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetV1DashboardStakingReturn gets staking return history

  Get staking return history
*/
func (a *Client) GetV1DashboardStakingReturn(params *GetV1DashboardStakingReturnParams, opts ...ClientOption) (*GetV1DashboardStakingReturnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1DashboardStakingReturnParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1DashboardStakingReturn",
		Method:             "GET",
		PathPattern:        "/v1/dashboard/staking_return",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetV1DashboardStakingReturnReader{formats: a.formats},
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
	success, ok := result.(*GetV1DashboardStakingReturnOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1DashboardStakingReturn: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetV1DashboardTxVolume gets tx volume history

  Get tx volume history
*/
func (a *Client) GetV1DashboardTxVolume(params *GetV1DashboardTxVolumeParams, opts ...ClientOption) (*GetV1DashboardTxVolumeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetV1DashboardTxVolumeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetV1DashboardTxVolume",
		Method:             "GET",
		PathPattern:        "/v1/dashboard/tx_volume",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetV1DashboardTxVolumeReader{formats: a.formats},
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
	success, ok := result.(*GetV1DashboardTxVolumeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for GetV1DashboardTxVolume: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

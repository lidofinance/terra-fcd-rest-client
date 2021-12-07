// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IbcCoreChannelV1QueryChannelsResponse QueryChannelsResponse is the response type for the Query/Channels RPC method.
//
// swagger:model ibc.core.channel.v1.QueryChannelsResponse
type IbcCoreChannelV1QueryChannelsResponse struct {

	// list of stored channels of the chain.
	Channels []*IbcCoreChannelV1QueryChannelsResponseChannelsItems0 `json:"channels"`

	// height
	Height *IbcCoreChannelV1QueryChannelsResponseHeight `json:"height,omitempty"`

	// pagination
	Pagination *IbcCoreChannelV1QueryChannelsResponsePagination `json:"pagination,omitempty"`
}

// Validate validates this ibc core channel v1 query channels response
func (m *IbcCoreChannelV1QueryChannelsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChannels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHeight(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePagination(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IbcCoreChannelV1QueryChannelsResponse) validateChannels(formats strfmt.Registry) error {
	if swag.IsZero(m.Channels) { // not required
		return nil
	}

	for i := 0; i < len(m.Channels); i++ {
		if swag.IsZero(m.Channels[i]) { // not required
			continue
		}

		if m.Channels[i] != nil {
			if err := m.Channels[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("channels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("channels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IbcCoreChannelV1QueryChannelsResponse) validateHeight(formats strfmt.Registry) error {
	if swag.IsZero(m.Height) { // not required
		return nil
	}

	if m.Height != nil {
		if err := m.Height.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("height")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("height")
			}
			return err
		}
	}

	return nil
}

func (m *IbcCoreChannelV1QueryChannelsResponse) validatePagination(formats strfmt.Registry) error {
	if swag.IsZero(m.Pagination) { // not required
		return nil
	}

	if m.Pagination != nil {
		if err := m.Pagination.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pagination")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this ibc core channel v1 query channels response based on the context it is used
func (m *IbcCoreChannelV1QueryChannelsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateChannels(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHeight(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePagination(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IbcCoreChannelV1QueryChannelsResponse) contextValidateChannels(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Channels); i++ {

		if m.Channels[i] != nil {
			if err := m.Channels[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("channels" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("channels" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IbcCoreChannelV1QueryChannelsResponse) contextValidateHeight(ctx context.Context, formats strfmt.Registry) error {

	if m.Height != nil {
		if err := m.Height.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("height")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("height")
			}
			return err
		}
	}

	return nil
}

func (m *IbcCoreChannelV1QueryChannelsResponse) contextValidatePagination(ctx context.Context, formats strfmt.Registry) error {

	if m.Pagination != nil {
		if err := m.Pagination.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pagination")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pagination")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IbcCoreChannelV1QueryChannelsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IbcCoreChannelV1QueryChannelsResponse) UnmarshalBinary(b []byte) error {
	var res IbcCoreChannelV1QueryChannelsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IbcCoreChannelV1QueryChannelsResponseChannelsItems0 IdentifiedChannel defines a channel with additional port and channel
// identifier fields.
//
// swagger:model IbcCoreChannelV1QueryChannelsResponseChannelsItems0
type IbcCoreChannelV1QueryChannelsResponseChannelsItems0 struct {

	// channel identifier
	ChannelID string `json:"channel_id,omitempty"`

	// list of connection identifiers, in order, along which packets sent on
	// this channel will travel
	ConnectionHops []string `json:"connection_hops"`

	// counterparty
	Counterparty *IbcCoreChannelV1QueryChannelsResponseChannelsItems0Counterparty `json:"counterparty,omitempty"`

	// whether the channel is ordered or unordered
	//
	// - ORDER_NONE_UNSPECIFIED: zero-value for channel ordering
	//  - ORDER_UNORDERED: packets can be delivered in any order, which may differ from the order in
	// which they were sent.
	//  - ORDER_ORDERED: packets are delivered exactly in the order which they were sent
	// Enum: [ORDER_NONE_UNSPECIFIED ORDER_UNORDERED ORDER_ORDERED]
	Ordering *string `json:"ordering,omitempty"`

	// port identifier
	PortID string `json:"port_id,omitempty"`

	// current state of the channel end
	//
	// State defines if a channel is in one of the following states:
	// CLOSED, INIT, TRYOPEN, OPEN or UNINITIALIZED.
	//
	//  - STATE_UNINITIALIZED_UNSPECIFIED: Default State
	//  - STATE_INIT: A channel has just started the opening handshake.
	//  - STATE_TRYOPEN: A channel has acknowledged the handshake step on the counterparty chain.
	//  - STATE_OPEN: A channel has completed the handshake. Open channels are
	// ready to send and receive packets.
	//  - STATE_CLOSED: A channel has been closed and can no longer be used to send or receive
	// packets.
	// Enum: [STATE_UNINITIALIZED_UNSPECIFIED STATE_INIT STATE_TRYOPEN STATE_OPEN STATE_CLOSED]
	State *string `json:"state,omitempty"`

	// opaque channel version, which is agreed upon during the handshake
	Version string `json:"version,omitempty"`
}

// Validate validates this ibc core channel v1 query channels response channels items0
func (m *IbcCoreChannelV1QueryChannelsResponseChannelsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCounterparty(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrdering(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IbcCoreChannelV1QueryChannelsResponseChannelsItems0) validateCounterparty(formats strfmt.Registry) error {
	if swag.IsZero(m.Counterparty) { // not required
		return nil
	}

	if m.Counterparty != nil {
		if err := m.Counterparty.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("counterparty")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("counterparty")
			}
			return err
		}
	}

	return nil
}

var ibcCoreChannelV1QueryChannelsResponseChannelsItems0TypeOrderingPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ORDER_NONE_UNSPECIFIED","ORDER_UNORDERED","ORDER_ORDERED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ibcCoreChannelV1QueryChannelsResponseChannelsItems0TypeOrderingPropEnum = append(ibcCoreChannelV1QueryChannelsResponseChannelsItems0TypeOrderingPropEnum, v)
	}
}

const (

	// IbcCoreChannelV1QueryChannelsResponseChannelsItems0OrderingORDERNONEUNSPECIFIED captures enum value "ORDER_NONE_UNSPECIFIED"
	IbcCoreChannelV1QueryChannelsResponseChannelsItems0OrderingORDERNONEUNSPECIFIED string = "ORDER_NONE_UNSPECIFIED"

	// IbcCoreChannelV1QueryChannelsResponseChannelsItems0OrderingORDERUNORDERED captures enum value "ORDER_UNORDERED"
	IbcCoreChannelV1QueryChannelsResponseChannelsItems0OrderingORDERUNORDERED string = "ORDER_UNORDERED"

	// IbcCoreChannelV1QueryChannelsResponseChannelsItems0OrderingORDERORDERED captures enum value "ORDER_ORDERED"
	IbcCoreChannelV1QueryChannelsResponseChannelsItems0OrderingORDERORDERED string = "ORDER_ORDERED"
)

// prop value enum
func (m *IbcCoreChannelV1QueryChannelsResponseChannelsItems0) validateOrderingEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ibcCoreChannelV1QueryChannelsResponseChannelsItems0TypeOrderingPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IbcCoreChannelV1QueryChannelsResponseChannelsItems0) validateOrdering(formats strfmt.Registry) error {
	if swag.IsZero(m.Ordering) { // not required
		return nil
	}

	// value enum
	if err := m.validateOrderingEnum("ordering", "body", *m.Ordering); err != nil {
		return err
	}

	return nil
}

var ibcCoreChannelV1QueryChannelsResponseChannelsItems0TypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["STATE_UNINITIALIZED_UNSPECIFIED","STATE_INIT","STATE_TRYOPEN","STATE_OPEN","STATE_CLOSED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ibcCoreChannelV1QueryChannelsResponseChannelsItems0TypeStatePropEnum = append(ibcCoreChannelV1QueryChannelsResponseChannelsItems0TypeStatePropEnum, v)
	}
}

const (

	// IbcCoreChannelV1QueryChannelsResponseChannelsItems0StateSTATEUNINITIALIZEDUNSPECIFIED captures enum value "STATE_UNINITIALIZED_UNSPECIFIED"
	IbcCoreChannelV1QueryChannelsResponseChannelsItems0StateSTATEUNINITIALIZEDUNSPECIFIED string = "STATE_UNINITIALIZED_UNSPECIFIED"

	// IbcCoreChannelV1QueryChannelsResponseChannelsItems0StateSTATEINIT captures enum value "STATE_INIT"
	IbcCoreChannelV1QueryChannelsResponseChannelsItems0StateSTATEINIT string = "STATE_INIT"

	// IbcCoreChannelV1QueryChannelsResponseChannelsItems0StateSTATETRYOPEN captures enum value "STATE_TRYOPEN"
	IbcCoreChannelV1QueryChannelsResponseChannelsItems0StateSTATETRYOPEN string = "STATE_TRYOPEN"

	// IbcCoreChannelV1QueryChannelsResponseChannelsItems0StateSTATEOPEN captures enum value "STATE_OPEN"
	IbcCoreChannelV1QueryChannelsResponseChannelsItems0StateSTATEOPEN string = "STATE_OPEN"

	// IbcCoreChannelV1QueryChannelsResponseChannelsItems0StateSTATECLOSED captures enum value "STATE_CLOSED"
	IbcCoreChannelV1QueryChannelsResponseChannelsItems0StateSTATECLOSED string = "STATE_CLOSED"
)

// prop value enum
func (m *IbcCoreChannelV1QueryChannelsResponseChannelsItems0) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ibcCoreChannelV1QueryChannelsResponseChannelsItems0TypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IbcCoreChannelV1QueryChannelsResponseChannelsItems0) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this ibc core channel v1 query channels response channels items0 based on the context it is used
func (m *IbcCoreChannelV1QueryChannelsResponseChannelsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCounterparty(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IbcCoreChannelV1QueryChannelsResponseChannelsItems0) contextValidateCounterparty(ctx context.Context, formats strfmt.Registry) error {

	if m.Counterparty != nil {
		if err := m.Counterparty.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("counterparty")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("counterparty")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IbcCoreChannelV1QueryChannelsResponseChannelsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IbcCoreChannelV1QueryChannelsResponseChannelsItems0) UnmarshalBinary(b []byte) error {
	var res IbcCoreChannelV1QueryChannelsResponseChannelsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IbcCoreChannelV1QueryChannelsResponseChannelsItems0Counterparty counterparty channel end
//
// swagger:model IbcCoreChannelV1QueryChannelsResponseChannelsItems0Counterparty
type IbcCoreChannelV1QueryChannelsResponseChannelsItems0Counterparty struct {

	// channel end on the counterparty chain
	ChannelID string `json:"channel_id,omitempty"`

	// port on the counterparty chain which owns the other end of the channel.
	PortID string `json:"port_id,omitempty"`
}

// Validate validates this ibc core channel v1 query channels response channels items0 counterparty
func (m *IbcCoreChannelV1QueryChannelsResponseChannelsItems0Counterparty) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ibc core channel v1 query channels response channels items0 counterparty based on context it is used
func (m *IbcCoreChannelV1QueryChannelsResponseChannelsItems0Counterparty) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IbcCoreChannelV1QueryChannelsResponseChannelsItems0Counterparty) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IbcCoreChannelV1QueryChannelsResponseChannelsItems0Counterparty) UnmarshalBinary(b []byte) error {
	var res IbcCoreChannelV1QueryChannelsResponseChannelsItems0Counterparty
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IbcCoreChannelV1QueryChannelsResponseHeight query block height
//
// Normally the RevisionHeight is incremented at each height while keeping
// RevisionNumber the same. However some consensus algorithms may choose to
// reset the height in certain conditions e.g. hard forks, state-machine
// breaking changes In these cases, the RevisionNumber is incremented so that
// height continues to be monitonically increasing even as the RevisionHeight
// gets reset
//
// swagger:model IbcCoreChannelV1QueryChannelsResponseHeight
type IbcCoreChannelV1QueryChannelsResponseHeight struct {

	// the height within the given revision
	RevisionHeight string `json:"revision_height,omitempty"`

	// the revision that the client is currently on
	RevisionNumber string `json:"revision_number,omitempty"`
}

// Validate validates this ibc core channel v1 query channels response height
func (m *IbcCoreChannelV1QueryChannelsResponseHeight) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ibc core channel v1 query channels response height based on context it is used
func (m *IbcCoreChannelV1QueryChannelsResponseHeight) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IbcCoreChannelV1QueryChannelsResponseHeight) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IbcCoreChannelV1QueryChannelsResponseHeight) UnmarshalBinary(b []byte) error {
	var res IbcCoreChannelV1QueryChannelsResponseHeight
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IbcCoreChannelV1QueryChannelsResponsePagination pagination response
//
// PageResponse is to be embedded in gRPC response messages where the
// corresponding request message has used PageRequest.
//
//  message SomeResponse {
//          repeated Bar results = 1;
//          PageResponse page = 2;
//  }
//
// swagger:model IbcCoreChannelV1QueryChannelsResponsePagination
type IbcCoreChannelV1QueryChannelsResponsePagination struct {

	// next_key is the key to be passed to PageRequest.key to
	// query the next page most efficiently
	// Format: byte
	NextKey strfmt.Base64 `json:"next_key,omitempty"`

	// total is total number of results available if PageRequest.count_total
	// was set, its value is undefined otherwise
	Total string `json:"total,omitempty"`
}

// Validate validates this ibc core channel v1 query channels response pagination
func (m *IbcCoreChannelV1QueryChannelsResponsePagination) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ibc core channel v1 query channels response pagination based on context it is used
func (m *IbcCoreChannelV1QueryChannelsResponsePagination) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IbcCoreChannelV1QueryChannelsResponsePagination) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IbcCoreChannelV1QueryChannelsResponsePagination) UnmarshalBinary(b []byte) error {
	var res IbcCoreChannelV1QueryChannelsResponsePagination
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

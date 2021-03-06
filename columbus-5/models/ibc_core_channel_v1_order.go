// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// IbcCoreChannelV1Order Order defines if a channel is ORDERED or UNORDERED
//
// - ORDER_NONE_UNSPECIFIED: zero-value for channel ordering
//  - ORDER_UNORDERED: packets can be delivered in any order, which may differ from the order in
// which they were sent.
//  - ORDER_ORDERED: packets are delivered exactly in the order which they were sent
//
// swagger:model ibc.core.channel.v1.Order
type IbcCoreChannelV1Order string

func NewIbcCoreChannelV1Order(value IbcCoreChannelV1Order) *IbcCoreChannelV1Order {
	v := value
	return &v
}

const (

	// IbcCoreChannelV1OrderORDERNONEUNSPECIFIED captures enum value "ORDER_NONE_UNSPECIFIED"
	IbcCoreChannelV1OrderORDERNONEUNSPECIFIED IbcCoreChannelV1Order = "ORDER_NONE_UNSPECIFIED"

	// IbcCoreChannelV1OrderORDERUNORDERED captures enum value "ORDER_UNORDERED"
	IbcCoreChannelV1OrderORDERUNORDERED IbcCoreChannelV1Order = "ORDER_UNORDERED"

	// IbcCoreChannelV1OrderORDERORDERED captures enum value "ORDER_ORDERED"
	IbcCoreChannelV1OrderORDERORDERED IbcCoreChannelV1Order = "ORDER_ORDERED"
)

// for schema
var ibcCoreChannelV1OrderEnum []interface{}

func init() {
	var res []IbcCoreChannelV1Order
	if err := json.Unmarshal([]byte(`["ORDER_NONE_UNSPECIFIED","ORDER_UNORDERED","ORDER_ORDERED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ibcCoreChannelV1OrderEnum = append(ibcCoreChannelV1OrderEnum, v)
	}
}

func (m IbcCoreChannelV1Order) validateIbcCoreChannelV1OrderEnum(path, location string, value IbcCoreChannelV1Order) error {
	if err := validate.EnumCase(path, location, value, ibcCoreChannelV1OrderEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this ibc core channel v1 order
func (m IbcCoreChannelV1Order) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateIbcCoreChannelV1OrderEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this ibc core channel v1 order based on context it is used
func (m IbcCoreChannelV1Order) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

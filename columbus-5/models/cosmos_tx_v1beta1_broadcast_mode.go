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

// CosmosTxV1beta1BroadcastMode BroadcastMode specifies the broadcast mode for the TxService.Broadcast RPC method.
//
//  - BROADCAST_MODE_UNSPECIFIED: zero-value for mode ordering
//  - BROADCAST_MODE_BLOCK: BROADCAST_MODE_BLOCK defines a tx broadcasting mode where the client waits for
// the tx to be committed in a block.
//  - BROADCAST_MODE_SYNC: BROADCAST_MODE_SYNC defines a tx broadcasting mode where the client waits for
// a CheckTx execution response only.
//  - BROADCAST_MODE_ASYNC: BROADCAST_MODE_ASYNC defines a tx broadcasting mode where the client returns
// immediately.
//
// swagger:model cosmos.tx.v1beta1.BroadcastMode
type CosmosTxV1beta1BroadcastMode string

func NewCosmosTxV1beta1BroadcastMode(value CosmosTxV1beta1BroadcastMode) *CosmosTxV1beta1BroadcastMode {
	v := value
	return &v
}

const (

	// CosmosTxV1beta1BroadcastModeBROADCASTMODEUNSPECIFIED captures enum value "BROADCAST_MODE_UNSPECIFIED"
	CosmosTxV1beta1BroadcastModeBROADCASTMODEUNSPECIFIED CosmosTxV1beta1BroadcastMode = "BROADCAST_MODE_UNSPECIFIED"

	// CosmosTxV1beta1BroadcastModeBROADCASTMODEBLOCK captures enum value "BROADCAST_MODE_BLOCK"
	CosmosTxV1beta1BroadcastModeBROADCASTMODEBLOCK CosmosTxV1beta1BroadcastMode = "BROADCAST_MODE_BLOCK"

	// CosmosTxV1beta1BroadcastModeBROADCASTMODESYNC captures enum value "BROADCAST_MODE_SYNC"
	CosmosTxV1beta1BroadcastModeBROADCASTMODESYNC CosmosTxV1beta1BroadcastMode = "BROADCAST_MODE_SYNC"

	// CosmosTxV1beta1BroadcastModeBROADCASTMODEASYNC captures enum value "BROADCAST_MODE_ASYNC"
	CosmosTxV1beta1BroadcastModeBROADCASTMODEASYNC CosmosTxV1beta1BroadcastMode = "BROADCAST_MODE_ASYNC"
)

// for schema
var cosmosTxV1beta1BroadcastModeEnum []interface{}

func init() {
	var res []CosmosTxV1beta1BroadcastMode
	if err := json.Unmarshal([]byte(`["BROADCAST_MODE_UNSPECIFIED","BROADCAST_MODE_BLOCK","BROADCAST_MODE_SYNC","BROADCAST_MODE_ASYNC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cosmosTxV1beta1BroadcastModeEnum = append(cosmosTxV1beta1BroadcastModeEnum, v)
	}
}

func (m CosmosTxV1beta1BroadcastMode) validateCosmosTxV1beta1BroadcastModeEnum(path, location string, value CosmosTxV1beta1BroadcastMode) error {
	if err := validate.EnumCase(path, location, value, cosmosTxV1beta1BroadcastModeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this cosmos tx v1beta1 broadcast mode
func (m CosmosTxV1beta1BroadcastMode) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCosmosTxV1beta1BroadcastModeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this cosmos tx v1beta1 broadcast mode based on context it is used
func (m CosmosTxV1beta1BroadcastMode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

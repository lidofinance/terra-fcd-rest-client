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

// CosmosTxV1beta1OrderBy OrderBy defines the sorting order
//
// - ORDER_BY_UNSPECIFIED: ORDER_BY_UNSPECIFIED specifies an unknown sorting order. OrderBy defaults to ASC in this case.
//  - ORDER_BY_ASC: ORDER_BY_ASC defines ascending order
//  - ORDER_BY_DESC: ORDER_BY_DESC defines descending order
//
// swagger:model cosmos.tx.v1beta1.OrderBy
type CosmosTxV1beta1OrderBy string

func NewCosmosTxV1beta1OrderBy(value CosmosTxV1beta1OrderBy) *CosmosTxV1beta1OrderBy {
	v := value
	return &v
}

const (

	// CosmosTxV1beta1OrderByORDERBYUNSPECIFIED captures enum value "ORDER_BY_UNSPECIFIED"
	CosmosTxV1beta1OrderByORDERBYUNSPECIFIED CosmosTxV1beta1OrderBy = "ORDER_BY_UNSPECIFIED"

	// CosmosTxV1beta1OrderByORDERBYASC captures enum value "ORDER_BY_ASC"
	CosmosTxV1beta1OrderByORDERBYASC CosmosTxV1beta1OrderBy = "ORDER_BY_ASC"

	// CosmosTxV1beta1OrderByORDERBYDESC captures enum value "ORDER_BY_DESC"
	CosmosTxV1beta1OrderByORDERBYDESC CosmosTxV1beta1OrderBy = "ORDER_BY_DESC"
)

// for schema
var cosmosTxV1beta1OrderByEnum []interface{}

func init() {
	var res []CosmosTxV1beta1OrderBy
	if err := json.Unmarshal([]byte(`["ORDER_BY_UNSPECIFIED","ORDER_BY_ASC","ORDER_BY_DESC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cosmosTxV1beta1OrderByEnum = append(cosmosTxV1beta1OrderByEnum, v)
	}
}

func (m CosmosTxV1beta1OrderBy) validateCosmosTxV1beta1OrderByEnum(path, location string, value CosmosTxV1beta1OrderBy) error {
	if err := validate.EnumCase(path, location, value, cosmosTxV1beta1OrderByEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this cosmos tx v1beta1 order by
func (m CosmosTxV1beta1OrderBy) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCosmosTxV1beta1OrderByEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this cosmos tx v1beta1 order by based on context it is used
func (m CosmosTxV1beta1OrderBy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

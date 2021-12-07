// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetIndividualWasmContractResultCodeInfo get individual wasm contract result code info
//
// swagger:model getIndividualWasmContractResult.code.info
type GetIndividualWasmContractResultCodeInfo struct {

	// description
	// Required: true
	Description *string `json:"description"`

	// tx memo
	// Required: true
	Memo *string `json:"memo"`

	// code name
	// Required: true
	Name *string `json:"name"`

	// code repo url
	// Required: true
	RepoURL *string `json:"repo_url"`
}

// Validate validates this get individual wasm contract result code info
func (m *GetIndividualWasmContractResultCodeInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRepoURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetIndividualWasmContractResultCodeInfo) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *GetIndividualWasmContractResultCodeInfo) validateMemo(formats strfmt.Registry) error {

	if err := validate.Required("memo", "body", m.Memo); err != nil {
		return err
	}

	return nil
}

func (m *GetIndividualWasmContractResultCodeInfo) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *GetIndividualWasmContractResultCodeInfo) validateRepoURL(formats strfmt.Registry) error {

	if err := validate.Required("repo_url", "body", m.RepoURL); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get individual wasm contract result code info based on context it is used
func (m *GetIndividualWasmContractResultCodeInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetIndividualWasmContractResultCodeInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetIndividualWasmContractResultCodeInfo) UnmarshalBinary(b []byte) error {
	var res GetIndividualWasmContractResultCodeInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

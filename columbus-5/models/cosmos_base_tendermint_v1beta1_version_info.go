// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CosmosBaseTendermintV1beta1VersionInfo VersionInfo is the type for the GetNodeInfoResponse message.
//
// swagger:model cosmos.base.tendermint.v1beta1.VersionInfo
type CosmosBaseTendermintV1beta1VersionInfo struct {

	// app name
	AppName string `json:"app_name,omitempty"`

	// build deps
	BuildDeps []*CosmosBaseTendermintV1beta1VersionInfoBuildDepsItems0 `json:"build_deps"`

	// build tags
	BuildTags string `json:"build_tags,omitempty"`

	// cosmos sdk version
	CosmosSdkVersion string `json:"cosmos_sdk_version,omitempty"`

	// git commit
	GitCommit string `json:"git_commit,omitempty"`

	// go version
	GoVersion string `json:"go_version,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this cosmos base tendermint v1beta1 version info
func (m *CosmosBaseTendermintV1beta1VersionInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBuildDeps(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosBaseTendermintV1beta1VersionInfo) validateBuildDeps(formats strfmt.Registry) error {
	if swag.IsZero(m.BuildDeps) { // not required
		return nil
	}

	for i := 0; i < len(m.BuildDeps); i++ {
		if swag.IsZero(m.BuildDeps[i]) { // not required
			continue
		}

		if m.BuildDeps[i] != nil {
			if err := m.BuildDeps[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("build_deps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("build_deps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this cosmos base tendermint v1beta1 version info based on the context it is used
func (m *CosmosBaseTendermintV1beta1VersionInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBuildDeps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosBaseTendermintV1beta1VersionInfo) contextValidateBuildDeps(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.BuildDeps); i++ {

		if m.BuildDeps[i] != nil {
			if err := m.BuildDeps[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("build_deps" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("build_deps" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CosmosBaseTendermintV1beta1VersionInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosBaseTendermintV1beta1VersionInfo) UnmarshalBinary(b []byte) error {
	var res CosmosBaseTendermintV1beta1VersionInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CosmosBaseTendermintV1beta1VersionInfoBuildDepsItems0 Module is the type for VersionInfo
//
// swagger:model CosmosBaseTendermintV1beta1VersionInfoBuildDepsItems0
type CosmosBaseTendermintV1beta1VersionInfoBuildDepsItems0 struct {

	// module path
	Path string `json:"path,omitempty"`

	// checksum
	Sum string `json:"sum,omitempty"`

	// module version
	Version string `json:"version,omitempty"`
}

// Validate validates this cosmos base tendermint v1beta1 version info build deps items0
func (m *CosmosBaseTendermintV1beta1VersionInfoBuildDepsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cosmos base tendermint v1beta1 version info build deps items0 based on context it is used
func (m *CosmosBaseTendermintV1beta1VersionInfoBuildDepsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CosmosBaseTendermintV1beta1VersionInfoBuildDepsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosBaseTendermintV1beta1VersionInfoBuildDepsItems0) UnmarshalBinary(b []byte) error {
	var res CosmosBaseTendermintV1beta1VersionInfoBuildDepsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
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

// CosmosUpgradeV1beta1QueryCurrentPlanResponse QueryCurrentPlanResponse is the response type for the Query/CurrentPlan RPC
// method.
//
// swagger:model cosmos.upgrade.v1beta1.QueryCurrentPlanResponse
type CosmosUpgradeV1beta1QueryCurrentPlanResponse struct {

	// plan
	Plan *CosmosUpgradeV1beta1QueryCurrentPlanResponsePlan `json:"plan,omitempty"`
}

// Validate validates this cosmos upgrade v1beta1 query current plan response
func (m *CosmosUpgradeV1beta1QueryCurrentPlanResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePlan(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosUpgradeV1beta1QueryCurrentPlanResponse) validatePlan(formats strfmt.Registry) error {
	if swag.IsZero(m.Plan) { // not required
		return nil
	}

	if m.Plan != nil {
		if err := m.Plan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("plan")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("plan")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cosmos upgrade v1beta1 query current plan response based on the context it is used
func (m *CosmosUpgradeV1beta1QueryCurrentPlanResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePlan(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosUpgradeV1beta1QueryCurrentPlanResponse) contextValidatePlan(ctx context.Context, formats strfmt.Registry) error {

	if m.Plan != nil {
		if err := m.Plan.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("plan")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("plan")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CosmosUpgradeV1beta1QueryCurrentPlanResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosUpgradeV1beta1QueryCurrentPlanResponse) UnmarshalBinary(b []byte) error {
	var res CosmosUpgradeV1beta1QueryCurrentPlanResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CosmosUpgradeV1beta1QueryCurrentPlanResponsePlan plan is the current upgrade plan.
//
// swagger:model CosmosUpgradeV1beta1QueryCurrentPlanResponsePlan
type CosmosUpgradeV1beta1QueryCurrentPlanResponsePlan struct {

	// The height at which the upgrade must be performed.
	// Only used if Time is not set.
	Height string `json:"height,omitempty"`

	// Any application specific upgrade info to be included on-chain
	// such as a git commit that validators could automatically upgrade to
	Info string `json:"info,omitempty"`

	// Sets the name for the upgrade. This name will be used by the upgraded
	// version of the software to apply any special "on-upgrade" commands during
	// the first BeginBlock method after the upgrade is applied. It is also used
	// to detect whether a software version can handle a given upgrade. If no
	// upgrade handler with this name has been set in the software, it will be
	// assumed that the software is out-of-date when the upgrade Time or Height is
	// reached and the software will exit.
	Name string `json:"name,omitempty"`

	// Deprecated: Time based upgrades have been deprecated. Time based upgrade logic
	// has been removed from the SDK.
	// If this field is not empty, an error will be thrown.
	// Format: date-time
	Time strfmt.DateTime `json:"time,omitempty"`

	// upgraded client state
	UpgradedClientState *CosmosUpgradeV1beta1QueryCurrentPlanResponsePlanUpgradedClientState `json:"upgraded_client_state,omitempty"`
}

// Validate validates this cosmos upgrade v1beta1 query current plan response plan
func (m *CosmosUpgradeV1beta1QueryCurrentPlanResponsePlan) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpgradedClientState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosUpgradeV1beta1QueryCurrentPlanResponsePlan) validateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.Time) { // not required
		return nil
	}

	if err := validate.FormatOf("plan"+"."+"time", "body", "date-time", m.Time.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CosmosUpgradeV1beta1QueryCurrentPlanResponsePlan) validateUpgradedClientState(formats strfmt.Registry) error {
	if swag.IsZero(m.UpgradedClientState) { // not required
		return nil
	}

	if m.UpgradedClientState != nil {
		if err := m.UpgradedClientState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("plan" + "." + "upgraded_client_state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("plan" + "." + "upgraded_client_state")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cosmos upgrade v1beta1 query current plan response plan based on the context it is used
func (m *CosmosUpgradeV1beta1QueryCurrentPlanResponsePlan) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUpgradedClientState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CosmosUpgradeV1beta1QueryCurrentPlanResponsePlan) contextValidateUpgradedClientState(ctx context.Context, formats strfmt.Registry) error {

	if m.UpgradedClientState != nil {
		if err := m.UpgradedClientState.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("plan" + "." + "upgraded_client_state")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("plan" + "." + "upgraded_client_state")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CosmosUpgradeV1beta1QueryCurrentPlanResponsePlan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosUpgradeV1beta1QueryCurrentPlanResponsePlan) UnmarshalBinary(b []byte) error {
	var res CosmosUpgradeV1beta1QueryCurrentPlanResponsePlan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CosmosUpgradeV1beta1QueryCurrentPlanResponsePlanUpgradedClientState `Any` contains an arbitrary serialized protocol buffer message along with a
// URL that describes the type of the serialized message.
//
// Protobuf library provides support to pack/unpack Any values in the form
// of utility functions or additional generated methods of the Any type.
//
// Example 1: Pack and unpack a message in C++.
//
//     Foo foo = ...;
//     Any any;
//     any.PackFrom(foo);
//     ...
//     if (any.UnpackTo(&foo)) {
//       ...
//     }
//
// Example 2: Pack and unpack a message in Java.
//
//     Foo foo = ...;
//     Any any = Any.pack(foo);
//     ...
//     if (any.is(Foo.class)) {
//       foo = any.unpack(Foo.class);
//     }
//
//  Example 3: Pack and unpack a message in Python.
//
//     foo = Foo(...)
//     any = Any()
//     any.Pack(foo)
//     ...
//     if any.Is(Foo.DESCRIPTOR):
//       any.Unpack(foo)
//       ...
//
//  Example 4: Pack and unpack a message in Go
//
//      foo := &pb.Foo{...}
//      any, err := ptypes.MarshalAny(foo)
//      ...
//      foo := &pb.Foo{}
//      if err := ptypes.UnmarshalAny(any, foo); err != nil {
//        ...
//      }
//
// The pack methods provided by protobuf library will by default use
// 'type.googleapis.com/full.type.name' as the type URL and the unpack
// methods only use the fully qualified type name after the last '/'
// in the type URL, for example "foo.bar.com/x/y.z" will yield type
// name "y.z".
//
//
// JSON
// ====
// The JSON representation of an `Any` value uses the regular
// representation of the deserialized, embedded message, with an
// additional field `@type` which contains the type URL. Example:
//
//     package google.profile;
//     message Person {
//       string first_name = 1;
//       string last_name = 2;
//     }
//
//     {
//       "@type": "type.googleapis.com/google.profile.Person",
//       "firstName": <string>,
//       "lastName": <string>
//     }
//
// If the embedded message type is well-known and has a custom JSON
// representation, that representation will be embedded adding a field
// `value` which holds the custom JSON in addition to the `@type`
// field. Example (for message [google.protobuf.Duration][]):
//
//     {
//       "@type": "type.googleapis.com/google.protobuf.Duration",
//       "value": "1.212s"
//     }
//
// swagger:model CosmosUpgradeV1beta1QueryCurrentPlanResponsePlanUpgradedClientState
type CosmosUpgradeV1beta1QueryCurrentPlanResponsePlanUpgradedClientState struct {

	// A URL/resource name that uniquely identifies the type of the serialized
	// protocol buffer message. This string must contain at least
	// one "/" character. The last segment of the URL's path must represent
	// the fully qualified name of the type (as in
	// `path/google.protobuf.Duration`). The name should be in a canonical form
	// (e.g., leading "." is not accepted).
	//
	// In practice, teams usually precompile into the binary all types that they
	// expect it to use in the context of Any. However, for URLs which use the
	// scheme `http`, `https`, or no scheme, one can optionally set up a type
	// server that maps type URLs to message definitions as follows:
	//
	// * If no scheme is provided, `https` is assumed.
	// * An HTTP GET on the URL must yield a [google.protobuf.Type][]
	//   value in binary format, or produce an error.
	// * Applications are allowed to cache lookup results based on the
	//   URL, or have them precompiled into a binary to avoid any
	//   lookup. Therefore, binary compatibility needs to be preserved
	//   on changes to types. (Use versioned type names to manage
	//   breaking changes.)
	//
	// Note: this functionality is not currently available in the official
	// protobuf release, and it is not used for type URLs beginning with
	// type.googleapis.com.
	//
	// Schemes other than `http`, `https` (or the empty scheme) might be
	// used with implementation specific semantics.
	TypeURL string `json:"type_url,omitempty"`

	// Must be a valid serialized protocol buffer of the above specified type.
	// Format: byte
	Value strfmt.Base64 `json:"value,omitempty"`
}

// Validate validates this cosmos upgrade v1beta1 query current plan response plan upgraded client state
func (m *CosmosUpgradeV1beta1QueryCurrentPlanResponsePlanUpgradedClientState) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this cosmos upgrade v1beta1 query current plan response plan upgraded client state based on context it is used
func (m *CosmosUpgradeV1beta1QueryCurrentPlanResponsePlanUpgradedClientState) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CosmosUpgradeV1beta1QueryCurrentPlanResponsePlanUpgradedClientState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CosmosUpgradeV1beta1QueryCurrentPlanResponsePlanUpgradedClientState) UnmarshalBinary(b []byte) error {
	var res CosmosUpgradeV1beta1QueryCurrentPlanResponsePlanUpgradedClientState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

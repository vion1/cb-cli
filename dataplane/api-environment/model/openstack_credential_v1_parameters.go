// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// OpenstackCredentialV1Parameters openstack credential v1 parameters
// swagger:model OpenstackCredentialV1Parameters
type OpenstackCredentialV1Parameters struct {

	// endpoint
	// Required: true
	Endpoint *string `json:"endpoint"`

	// facing
	// Required: true
	// Enum: [public admin internal]
	Facing *string `json:"facing"`

	// keystone v2
	KeystoneV2 *KeystoneV2Parameters `json:"keystoneV2,omitempty"`

	// keystone v3
	KeystoneV3 *KeystoneV3Parameters `json:"keystoneV3,omitempty"`

	// password
	// Required: true
	Password *string `json:"password"`

	// user name
	// Required: true
	UserName *string `json:"userName"`
}

// Validate validates this openstack credential v1 parameters
func (m *OpenstackCredentialV1Parameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndpoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFacing(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeystoneV2(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeystoneV3(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenstackCredentialV1Parameters) validateEndpoint(formats strfmt.Registry) error {

	if err := validate.Required("endpoint", "body", m.Endpoint); err != nil {
		return err
	}

	return nil
}

var openstackCredentialV1ParametersTypeFacingPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["public","admin","internal"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		openstackCredentialV1ParametersTypeFacingPropEnum = append(openstackCredentialV1ParametersTypeFacingPropEnum, v)
	}
}

const (

	// OpenstackCredentialV1ParametersFacingPublic captures enum value "public"
	OpenstackCredentialV1ParametersFacingPublic string = "public"

	// OpenstackCredentialV1ParametersFacingAdmin captures enum value "admin"
	OpenstackCredentialV1ParametersFacingAdmin string = "admin"

	// OpenstackCredentialV1ParametersFacingInternal captures enum value "internal"
	OpenstackCredentialV1ParametersFacingInternal string = "internal"
)

// prop value enum
func (m *OpenstackCredentialV1Parameters) validateFacingEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, openstackCredentialV1ParametersTypeFacingPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *OpenstackCredentialV1Parameters) validateFacing(formats strfmt.Registry) error {

	if err := validate.Required("facing", "body", m.Facing); err != nil {
		return err
	}

	// value enum
	if err := m.validateFacingEnum("facing", "body", *m.Facing); err != nil {
		return err
	}

	return nil
}

func (m *OpenstackCredentialV1Parameters) validateKeystoneV2(formats strfmt.Registry) error {

	if swag.IsZero(m.KeystoneV2) { // not required
		return nil
	}

	if m.KeystoneV2 != nil {
		if err := m.KeystoneV2.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("keystoneV2")
			}
			return err
		}
	}

	return nil
}

func (m *OpenstackCredentialV1Parameters) validateKeystoneV3(formats strfmt.Registry) error {

	if swag.IsZero(m.KeystoneV3) { // not required
		return nil
	}

	if m.KeystoneV3 != nil {
		if err := m.KeystoneV3.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("keystoneV3")
			}
			return err
		}
	}

	return nil
}

func (m *OpenstackCredentialV1Parameters) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	return nil
}

func (m *OpenstackCredentialV1Parameters) validateUserName(formats strfmt.Registry) error {

	if err := validate.Required("userName", "body", m.UserName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenstackCredentialV1Parameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenstackCredentialV1Parameters) UnmarshalBinary(b []byte) error {
	var res OpenstackCredentialV1Parameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

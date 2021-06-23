// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// NetworkV1Request network v1 request
// swagger:model NetworkV1Request
type NetworkV1Request struct {

	// provider specific parameters of the specified network
	Aws *AwsNetworkV1Parameters `json:"aws,omitempty"`

	// provider specific parameters of the specified network
	Azure *AzureNetworkV1Parameters `json:"azure,omitempty"`

	// provider specific parameters of the specified network
	Gcp *GcpNetworkV1Parameters `json:"gcp,omitempty"`

	// mock network parameters
	Mock *MockNetworkV1Parameters `json:"mock,omitempty"`

	// provider specific parameters of the specified network
	Openstack *OpenstackNetworkV1Parameters `json:"openstack,omitempty"`

	// provider specific parameters of the specified network
	Yarn YarnNetworkV1Parameters `json:"yarn,omitempty"`
}

// Validate validates this network v1 request
func (m *NetworkV1Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAws(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzure(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGcp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMock(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOpenstack(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkV1Request) validateAws(formats strfmt.Registry) error {

	if swag.IsZero(m.Aws) { // not required
		return nil
	}

	if m.Aws != nil {
		if err := m.Aws.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkV1Request) validateAzure(formats strfmt.Registry) error {

	if swag.IsZero(m.Azure) { // not required
		return nil
	}

	if m.Azure != nil {
		if err := m.Azure.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azure")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkV1Request) validateGcp(formats strfmt.Registry) error {

	if swag.IsZero(m.Gcp) { // not required
		return nil
	}

	if m.Gcp != nil {
		if err := m.Gcp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcp")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkV1Request) validateMock(formats strfmt.Registry) error {

	if swag.IsZero(m.Mock) { // not required
		return nil
	}

	if m.Mock != nil {
		if err := m.Mock.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mock")
			}
			return err
		}
	}

	return nil
}

func (m *NetworkV1Request) validateOpenstack(formats strfmt.Registry) error {

	if swag.IsZero(m.Openstack) { // not required
		return nil
	}

	if m.Openstack != nil {
		if err := m.Openstack.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("openstack")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkV1Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkV1Request) UnmarshalBinary(b []byte) error {
	var res NetworkV1Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

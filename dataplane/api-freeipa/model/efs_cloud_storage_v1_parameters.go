// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EfsCloudStorageV1Parameters efs cloud storage v1 parameters
// swagger:model EfsCloudStorageV1Parameters
type EfsCloudStorageV1Parameters struct {

	// instance profile
	// Required: true
	InstanceProfile *string `json:"instanceProfile"`
}

// Validate validates this efs cloud storage v1 parameters
func (m *EfsCloudStorageV1Parameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstanceProfile(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EfsCloudStorageV1Parameters) validateInstanceProfile(formats strfmt.Registry) error {

	if err := validate.Required("instanceProfile", "body", m.InstanceProfile); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EfsCloudStorageV1Parameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EfsCloudStorageV1Parameters) UnmarshalBinary(b []byte) error {
	var res EfsCloudStorageV1Parameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

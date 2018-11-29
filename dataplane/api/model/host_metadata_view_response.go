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

// HostMetadataViewResponse host metadata view response
// swagger:model HostMetadataViewResponse
type HostMetadataViewResponse struct {

	// id of the resource
	ID int64 `json:"id,omitempty"`

	// name of the resource
	// Required: true
	Name *string `json:"name"`

	// state of the host
	State string `json:"state,omitempty"`
}

// Validate validates this host metadata view response
func (m *HostMetadataViewResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HostMetadataViewResponse) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HostMetadataViewResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HostMetadataViewResponse) UnmarshalBinary(b []byte) error {
	var res HostMetadataViewResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

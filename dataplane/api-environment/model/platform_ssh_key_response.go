// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PlatformSSHKeyResponse platform Ssh key response
// swagger:model PlatformSshKeyResponse
type PlatformSSHKeyResponse struct {

	// name
	Name string `json:"name,omitempty"`

	// properties
	Properties map[string]interface{} `json:"properties,omitempty"`
}

// Validate validates this platform Ssh key response
func (m *PlatformSSHKeyResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PlatformSSHKeyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlatformSSHKeyResponse) UnmarshalBinary(b []byte) error {
	var res PlatformSSHKeyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

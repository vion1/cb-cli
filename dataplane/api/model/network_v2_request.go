// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// NetworkV2Request network v2 request
// swagger:model NetworkV2Request
type NetworkV2Request struct {

	// provider specific parameters of the specified network
	Parameters map[string]interface{} `json:"parameters,omitempty"`

	// the subnet definition of the network in CIDR format
	SubnetCIDR string `json:"subnetCIDR,omitempty"`
}

// Validate validates this network v2 request
func (m *NetworkV2Request) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NetworkV2Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkV2Request) UnmarshalBinary(b []byte) error {
	var res NetworkV2Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

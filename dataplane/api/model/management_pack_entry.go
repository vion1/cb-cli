// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ManagementPackEntry management pack entry
// swagger:model ManagementPackEntry
type ManagementPackEntry struct {

	// mpack Url
	MpackURL string `json:"mpackUrl,omitempty"`
}

// Validate validates this management pack entry
func (m *ManagementPackEntry) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ManagementPackEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ManagementPackEntry) UnmarshalBinary(b []byte) error {
	var res ManagementPackEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

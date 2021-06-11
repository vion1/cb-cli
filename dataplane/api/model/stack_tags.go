// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// StackTags stack tags
// swagger:model StackTags
type StackTags struct {

	// application tags
	ApplicationTags map[string]string `json:"applicationTags,omitempty"`

	// default tags
	DefaultTags map[string]string `json:"defaultTags,omitempty"`

	// user defined tags
	UserDefinedTags map[string]string `json:"userDefinedTags,omitempty"`
}

// Validate validates this stack tags
func (m *StackTags) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StackTags) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StackTags) UnmarshalBinary(b []byte) error {
	var res StackTags
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

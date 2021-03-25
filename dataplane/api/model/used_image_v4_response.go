// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// UsedImageV4Response used image v4 response
// swagger:model UsedImageV4Response
type UsedImageV4Response struct {

	// image Id
	ImageID string `json:"imageId,omitempty"`

	// stack version
	StackVersion string `json:"stackVersion,omitempty"`
}

// Validate validates this used image v4 response
func (m *UsedImageV4Response) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UsedImageV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UsedImageV4Response) UnmarshalBinary(b []byte) error {
	var res UsedImageV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

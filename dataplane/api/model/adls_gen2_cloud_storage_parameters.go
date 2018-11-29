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

// AdlsGen2CloudStorageParameters adls gen2 cloud storage parameters
// swagger:model AdlsGen2CloudStorageParameters
type AdlsGen2CloudStorageParameters struct {

	// account key
	// Required: true
	AccountKey *string `json:"accountKey"`

	// account name
	// Required: true
	AccountName *string `json:"accountName"`
}

// Validate validates this adls gen2 cloud storage parameters
func (m *AdlsGen2CloudStorageParameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccountName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AdlsGen2CloudStorageParameters) validateAccountKey(formats strfmt.Registry) error {

	if err := validate.Required("accountKey", "body", m.AccountKey); err != nil {
		return err
	}

	return nil
}

func (m *AdlsGen2CloudStorageParameters) validateAccountName(formats strfmt.Registry) error {

	if err := validate.Required("accountName", "body", m.AccountName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AdlsGen2CloudStorageParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdlsGen2CloudStorageParameters) UnmarshalBinary(b []byte) error {
	var res AdlsGen2CloudStorageParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

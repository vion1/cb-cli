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

// StorageLocationRequest storage location request
// swagger:model StorageLocationRequest
type StorageLocationRequest struct {

	// property file
	// Required: true
	PropertyFile *string `json:"propertyFile"`

	// property name
	// Required: true
	PropertyName *string `json:"propertyName"`

	// value
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this storage location request
func (m *StorageLocationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePropertyFile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePropertyName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StorageLocationRequest) validatePropertyFile(formats strfmt.Registry) error {

	if err := validate.Required("propertyFile", "body", m.PropertyFile); err != nil {
		return err
	}

	return nil
}

func (m *StorageLocationRequest) validatePropertyName(formats strfmt.Registry) error {

	if err := validate.Required("propertyName", "body", m.PropertyName); err != nil {
		return err
	}

	return nil
}

func (m *StorageLocationRequest) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StorageLocationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorageLocationRequest) UnmarshalBinary(b []byte) error {
	var res StorageLocationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

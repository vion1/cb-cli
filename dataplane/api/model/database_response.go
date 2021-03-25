// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DatabaseResponse database response
// swagger:model DatabaseResponse
type DatabaseResponse struct {

	// availability type
	// Required: true
	// Enum: [NONE NON_HA HA ON_ROOT_VOLUME]
	AvailabilityType *string `json:"availabilityType"`
}

// Validate validates this database response
func (m *DatabaseResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailabilityType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var databaseResponseTypeAvailabilityTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NONE","NON_HA","HA","ON_ROOT_VOLUME"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		databaseResponseTypeAvailabilityTypePropEnum = append(databaseResponseTypeAvailabilityTypePropEnum, v)
	}
}

const (

	// DatabaseResponseAvailabilityTypeNONE captures enum value "NONE"
	DatabaseResponseAvailabilityTypeNONE string = "NONE"

	// DatabaseResponseAvailabilityTypeNONHA captures enum value "NON_HA"
	DatabaseResponseAvailabilityTypeNONHA string = "NON_HA"

	// DatabaseResponseAvailabilityTypeHA captures enum value "HA"
	DatabaseResponseAvailabilityTypeHA string = "HA"

	// DatabaseResponseAvailabilityTypeONROOTVOLUME captures enum value "ON_ROOT_VOLUME"
	DatabaseResponseAvailabilityTypeONROOTVOLUME string = "ON_ROOT_VOLUME"
)

// prop value enum
func (m *DatabaseResponse) validateAvailabilityTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, databaseResponseTypeAvailabilityTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DatabaseResponse) validateAvailabilityType(formats strfmt.Registry) error {

	if err := validate.Required("availabilityType", "body", m.AvailabilityType); err != nil {
		return err
	}

	// value enum
	if err := m.validateAvailabilityTypeEnum("availabilityType", "body", *m.AvailabilityType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DatabaseResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatabaseResponse) UnmarshalBinary(b []byte) error {
	var res DatabaseResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

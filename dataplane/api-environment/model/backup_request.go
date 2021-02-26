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

// BackupRequest backup request
// swagger:model BackupRequest
type BackupRequest struct {

	// backup adls gen2 attributes
	AdlsGen2 *AdlsGen2CloudStorageV1Parameters `json:"adlsGen2,omitempty"`

	// backup cloudwatch attributes
	Cloudwatch *BackupCloudwatchParams `json:"cloudwatch,omitempty"`

	// backup gcs attributes
	Gcs *GcsCloudStorageV1Parameters `json:"gcs,omitempty"`

	// backup s3 attributes
	S3 *S3CloudStorageV1Parameters `json:"s3,omitempty"`

	// backup storage location / container
	// Required: true
	StorageLocation *string `json:"storageLocation"`
}

// Validate validates this backup request
func (m *BackupRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdlsGen2(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudwatch(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGcs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateS3(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStorageLocation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BackupRequest) validateAdlsGen2(formats strfmt.Registry) error {

	if swag.IsZero(m.AdlsGen2) { // not required
		return nil
	}

	if m.AdlsGen2 != nil {
		if err := m.AdlsGen2.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("adlsGen2")
			}
			return err
		}
	}

	return nil
}

func (m *BackupRequest) validateCloudwatch(formats strfmt.Registry) error {

	if swag.IsZero(m.Cloudwatch) { // not required
		return nil
	}

	if m.Cloudwatch != nil {
		if err := m.Cloudwatch.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloudwatch")
			}
			return err
		}
	}

	return nil
}

func (m *BackupRequest) validateGcs(formats strfmt.Registry) error {

	if swag.IsZero(m.Gcs) { // not required
		return nil
	}

	if m.Gcs != nil {
		if err := m.Gcs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcs")
			}
			return err
		}
	}

	return nil
}

func (m *BackupRequest) validateS3(formats strfmt.Registry) error {

	if swag.IsZero(m.S3) { // not required
		return nil
	}

	if m.S3 != nil {
		if err := m.S3.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("s3")
			}
			return err
		}
	}

	return nil
}

func (m *BackupRequest) validateStorageLocation(formats strfmt.Registry) error {

	if err := validate.Required("storageLocation", "body", m.StorageLocation); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BackupRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BackupRequest) UnmarshalBinary(b []byte) error {
	var res BackupRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

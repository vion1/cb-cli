// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DistroXV1Request distro x v1 request
// swagger:model DistroXV1Request
type DistroXV1Request struct {

	// aws specific parameters for stack
	Aws AwsDistroXV1Parameters `json:"aws,omitempty"`

	// azure specific parameters for stack
	Azure *AzureDistroXV1Parameters `json:"azure,omitempty"`

	// cloud platform
	// Enum: [AWS GCP AZURE OPENSTACK YARN MOCK]
	CloudPlatform string `json:"cloudPlatform,omitempty"`

	// cluster
	Cluster *DistroXClusterV1Request `json:"cluster,omitempty"`

	// environment name
	EnvironmentName string `json:"environmentName,omitempty"`

	// external database
	ExternalDatabase *DistroXDatabaseRequest `json:"externalDatabase,omitempty"`

	// gcp
	Gcp GcpDistroXV1Parameters `json:"gcp,omitempty"`

	// image
	Image *DistroXImageV1Request `json:"image,omitempty"`

	// inputs
	Inputs map[string]interface{} `json:"inputs,omitempty"`

	// instance groups
	// Unique: true
	InstanceGroups []*InstanceGroupV1Request `json:"instanceGroups"`

	// name of the stack
	// Required: true
	Name *string `json:"name"`

	// network
	Network *NetworkV1Request `json:"network,omitempty"`

	// sdx
	Sdx *SdxV1Request `json:"sdx,omitempty"`

	// tags
	Tags *TagsV1Request `json:"tags,omitempty"`

	// time to live
	TimeToLive int64 `json:"timeToLive,omitempty"`

	// yarn
	Yarn *YarnDistroXV1Parameters `json:"yarn,omitempty"`
}

// Validate validates this distro x v1 request
func (m *DistroXV1Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAzure(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudPlatform(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalDatabase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInstanceGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetwork(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSdx(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateYarn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DistroXV1Request) validateAzure(formats strfmt.Registry) error {

	if swag.IsZero(m.Azure) { // not required
		return nil
	}

	if m.Azure != nil {
		if err := m.Azure.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azure")
			}
			return err
		}
	}

	return nil
}

var distroXV1RequestTypeCloudPlatformPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AWS","GCP","AZURE","OPENSTACK","YARN","MOCK"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		distroXV1RequestTypeCloudPlatformPropEnum = append(distroXV1RequestTypeCloudPlatformPropEnum, v)
	}
}

const (

	// DistroXV1RequestCloudPlatformAWS captures enum value "AWS"
	DistroXV1RequestCloudPlatformAWS string = "AWS"

	// DistroXV1RequestCloudPlatformGCP captures enum value "GCP"
	DistroXV1RequestCloudPlatformGCP string = "GCP"

	// DistroXV1RequestCloudPlatformAZURE captures enum value "AZURE"
	DistroXV1RequestCloudPlatformAZURE string = "AZURE"

	// DistroXV1RequestCloudPlatformOPENSTACK captures enum value "OPENSTACK"
	DistroXV1RequestCloudPlatformOPENSTACK string = "OPENSTACK"

	// DistroXV1RequestCloudPlatformYARN captures enum value "YARN"
	DistroXV1RequestCloudPlatformYARN string = "YARN"

	// DistroXV1RequestCloudPlatformMOCK captures enum value "MOCK"
	DistroXV1RequestCloudPlatformMOCK string = "MOCK"
)

// prop value enum
func (m *DistroXV1Request) validateCloudPlatformEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, distroXV1RequestTypeCloudPlatformPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *DistroXV1Request) validateCloudPlatform(formats strfmt.Registry) error {

	if swag.IsZero(m.CloudPlatform) { // not required
		return nil
	}

	// value enum
	if err := m.validateCloudPlatformEnum("cloudPlatform", "body", m.CloudPlatform); err != nil {
		return err
	}

	return nil
}

func (m *DistroXV1Request) validateCluster(formats strfmt.Registry) error {

	if swag.IsZero(m.Cluster) { // not required
		return nil
	}

	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *DistroXV1Request) validateExternalDatabase(formats strfmt.Registry) error {

	if swag.IsZero(m.ExternalDatabase) { // not required
		return nil
	}

	if m.ExternalDatabase != nil {
		if err := m.ExternalDatabase.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("externalDatabase")
			}
			return err
		}
	}

	return nil
}

func (m *DistroXV1Request) validateImage(formats strfmt.Registry) error {

	if swag.IsZero(m.Image) { // not required
		return nil
	}

	if m.Image != nil {
		if err := m.Image.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("image")
			}
			return err
		}
	}

	return nil
}

func (m *DistroXV1Request) validateInstanceGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.InstanceGroups) { // not required
		return nil
	}

	if err := validate.UniqueItems("instanceGroups", "body", m.InstanceGroups); err != nil {
		return err
	}

	for i := 0; i < len(m.InstanceGroups); i++ {
		if swag.IsZero(m.InstanceGroups[i]) { // not required
			continue
		}

		if m.InstanceGroups[i] != nil {
			if err := m.InstanceGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("instanceGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DistroXV1Request) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *DistroXV1Request) validateNetwork(formats strfmt.Registry) error {

	if swag.IsZero(m.Network) { // not required
		return nil
	}

	if m.Network != nil {
		if err := m.Network.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network")
			}
			return err
		}
	}

	return nil
}

func (m *DistroXV1Request) validateSdx(formats strfmt.Registry) error {

	if swag.IsZero(m.Sdx) { // not required
		return nil
	}

	if m.Sdx != nil {
		if err := m.Sdx.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sdx")
			}
			return err
		}
	}

	return nil
}

func (m *DistroXV1Request) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	if m.Tags != nil {
		if err := m.Tags.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tags")
			}
			return err
		}
	}

	return nil
}

func (m *DistroXV1Request) validateYarn(formats strfmt.Registry) error {

	if swag.IsZero(m.Yarn) { // not required
		return nil
	}

	if m.Yarn != nil {
		if err := m.Yarn.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("yarn")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DistroXV1Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DistroXV1Request) UnmarshalBinary(b []byte) error {
	var res DistroXV1Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

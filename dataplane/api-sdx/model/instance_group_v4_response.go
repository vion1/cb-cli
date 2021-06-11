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

// InstanceGroupV4Response instance group v4 response
// swagger:model InstanceGroupV4Response
type InstanceGroupV4Response struct {

	// aws specific parameters for instance group
	Aws AwsInstanceGroupV4Parameters `json:"aws,omitempty"`

	// azure specific parameters for instance group
	Azure *AzureInstanceGroupV4Parameters `json:"azure,omitempty"`

	// gcp specific parameters for instance group
	Gcp *GcpInstanceGroupV4Parameters `json:"gcp,omitempty"`

	// id of the resource
	ID int64 `json:"id,omitempty"`

	// metadata of instances
	// Unique: true
	Metadata []*InstanceMetaDataV4Response `json:"metadata"`

	// minimum nodecount in an instancegroup
	MinimumNodeCount int32 `json:"minimumNodeCount,omitempty"`

	// mock
	Mock MockInstanceGroupV4Parameters `json:"mock,omitempty"`

	// name of the instance group
	// Required: true
	Name *string `json:"name"`

	// number of nodes
	// Required: true
	// Maximum: 100000
	// Minimum: 0
	NodeCount *int32 `json:"nodeCount"`

	// openstack specific parameters for instance group
	Openstack *OpenStackInstanceGroupV4Parameters `json:"openstack,omitempty"`

	// recipes
	Recipes []*RecipeV4Response `json:"recipes"`

	// recovery mode of the hostgroup's nodes
	// Enum: [MANUAL AUTO]
	RecoveryMode string `json:"recoveryMode,omitempty"`

	// type of the instance group scalability, default value is ALLOWED
	// Enum: [ALLOWED FORBIDDEN ONLY_UPSCALE ONLY_DOWNSCALE]
	ScalabilityOption string `json:"scalabilityOption,omitempty"`

	// instancegroup related securitygroup
	SecurityGroup *SecurityGroupV4Response `json:"securityGroup,omitempty"`

	// instancegroup related template
	Template *InstanceTemplateV4Response `json:"template,omitempty"`

	// type of the instance group, default value is CORE
	// Enum: [CORE GATEWAY]
	Type string `json:"type,omitempty"`
}

// Validate validates this instance group v4 response
func (m *InstanceGroupV4Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAzure(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGcp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNodeCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOpenstack(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecipes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecoveryMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScalabilityOption(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstanceGroupV4Response) validateAzure(formats strfmt.Registry) error {

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

func (m *InstanceGroupV4Response) validateGcp(formats strfmt.Registry) error {

	if swag.IsZero(m.Gcp) { // not required
		return nil
	}

	if m.Gcp != nil {
		if err := m.Gcp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcp")
			}
			return err
		}
	}

	return nil
}

func (m *InstanceGroupV4Response) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if err := validate.UniqueItems("metadata", "body", m.Metadata); err != nil {
		return err
	}

	for i := 0; i < len(m.Metadata); i++ {
		if swag.IsZero(m.Metadata[i]) { // not required
			continue
		}

		if m.Metadata[i] != nil {
			if err := m.Metadata[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("metadata" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *InstanceGroupV4Response) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *InstanceGroupV4Response) validateNodeCount(formats strfmt.Registry) error {

	if err := validate.Required("nodeCount", "body", m.NodeCount); err != nil {
		return err
	}

	if err := validate.MinimumInt("nodeCount", "body", int64(*m.NodeCount), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("nodeCount", "body", int64(*m.NodeCount), 100000, false); err != nil {
		return err
	}

	return nil
}

func (m *InstanceGroupV4Response) validateOpenstack(formats strfmt.Registry) error {

	if swag.IsZero(m.Openstack) { // not required
		return nil
	}

	if m.Openstack != nil {
		if err := m.Openstack.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("openstack")
			}
			return err
		}
	}

	return nil
}

func (m *InstanceGroupV4Response) validateRecipes(formats strfmt.Registry) error {

	if swag.IsZero(m.Recipes) { // not required
		return nil
	}

	for i := 0; i < len(m.Recipes); i++ {
		if swag.IsZero(m.Recipes[i]) { // not required
			continue
		}

		if m.Recipes[i] != nil {
			if err := m.Recipes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("recipes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var instanceGroupV4ResponseTypeRecoveryModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["MANUAL","AUTO"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		instanceGroupV4ResponseTypeRecoveryModePropEnum = append(instanceGroupV4ResponseTypeRecoveryModePropEnum, v)
	}
}

const (

	// InstanceGroupV4ResponseRecoveryModeMANUAL captures enum value "MANUAL"
	InstanceGroupV4ResponseRecoveryModeMANUAL string = "MANUAL"

	// InstanceGroupV4ResponseRecoveryModeAUTO captures enum value "AUTO"
	InstanceGroupV4ResponseRecoveryModeAUTO string = "AUTO"
)

// prop value enum
func (m *InstanceGroupV4Response) validateRecoveryModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, instanceGroupV4ResponseTypeRecoveryModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *InstanceGroupV4Response) validateRecoveryMode(formats strfmt.Registry) error {

	if swag.IsZero(m.RecoveryMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateRecoveryModeEnum("recoveryMode", "body", m.RecoveryMode); err != nil {
		return err
	}

	return nil
}

var instanceGroupV4ResponseTypeScalabilityOptionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ALLOWED","FORBIDDEN","ONLY_UPSCALE","ONLY_DOWNSCALE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		instanceGroupV4ResponseTypeScalabilityOptionPropEnum = append(instanceGroupV4ResponseTypeScalabilityOptionPropEnum, v)
	}
}

const (

	// InstanceGroupV4ResponseScalabilityOptionALLOWED captures enum value "ALLOWED"
	InstanceGroupV4ResponseScalabilityOptionALLOWED string = "ALLOWED"

	// InstanceGroupV4ResponseScalabilityOptionFORBIDDEN captures enum value "FORBIDDEN"
	InstanceGroupV4ResponseScalabilityOptionFORBIDDEN string = "FORBIDDEN"

	// InstanceGroupV4ResponseScalabilityOptionONLYUPSCALE captures enum value "ONLY_UPSCALE"
	InstanceGroupV4ResponseScalabilityOptionONLYUPSCALE string = "ONLY_UPSCALE"

	// InstanceGroupV4ResponseScalabilityOptionONLYDOWNSCALE captures enum value "ONLY_DOWNSCALE"
	InstanceGroupV4ResponseScalabilityOptionONLYDOWNSCALE string = "ONLY_DOWNSCALE"
)

// prop value enum
func (m *InstanceGroupV4Response) validateScalabilityOptionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, instanceGroupV4ResponseTypeScalabilityOptionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *InstanceGroupV4Response) validateScalabilityOption(formats strfmt.Registry) error {

	if swag.IsZero(m.ScalabilityOption) { // not required
		return nil
	}

	// value enum
	if err := m.validateScalabilityOptionEnum("scalabilityOption", "body", m.ScalabilityOption); err != nil {
		return err
	}

	return nil
}

func (m *InstanceGroupV4Response) validateSecurityGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.SecurityGroup) { // not required
		return nil
	}

	if m.SecurityGroup != nil {
		if err := m.SecurityGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("securityGroup")
			}
			return err
		}
	}

	return nil
}

func (m *InstanceGroupV4Response) validateTemplate(formats strfmt.Registry) error {

	if swag.IsZero(m.Template) { // not required
		return nil
	}

	if m.Template != nil {
		if err := m.Template.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("template")
			}
			return err
		}
	}

	return nil
}

var instanceGroupV4ResponseTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CORE","GATEWAY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		instanceGroupV4ResponseTypeTypePropEnum = append(instanceGroupV4ResponseTypeTypePropEnum, v)
	}
}

const (

	// InstanceGroupV4ResponseTypeCORE captures enum value "CORE"
	InstanceGroupV4ResponseTypeCORE string = "CORE"

	// InstanceGroupV4ResponseTypeGATEWAY captures enum value "GATEWAY"
	InstanceGroupV4ResponseTypeGATEWAY string = "GATEWAY"
)

// prop value enum
func (m *InstanceGroupV4Response) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, instanceGroupV4ResponseTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *InstanceGroupV4Response) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InstanceGroupV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstanceGroupV4Response) UnmarshalBinary(b []byte) error {
	var res InstanceGroupV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

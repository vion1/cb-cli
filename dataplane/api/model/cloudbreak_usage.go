// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CloudbreakUsage cloudbreak usage
// swagger:model CloudbreakUsage
type CloudbreakUsage struct {

	// availability zone of the stack
	AvailabilityZone string `json:"availabilityZone,omitempty"`

	// id of the blueprint
	BlueprintID int64 `json:"blueprintId,omitempty"`

	// name of the blueprint
	BlueprintName string `json:"blueprintName,omitempty"`

	// computed costs of instance usage
	Costs float64 `json:"costs,omitempty"`

	// the day the usage of resources happened
	Day string `json:"day,omitempty"`

	// time since the instances are running in millisec
	Duration string `json:"duration,omitempty"`

	// flex subscription id
	FlexID string `json:"flexId,omitempty"`

	// group name of instance
	InstanceGroup string `json:"instanceGroup,omitempty"`

	// hours since the instance is running
	InstanceHours int64 `json:"instanceHours,omitempty"`

	// number of instances running
	InstanceNum int32 `json:"instanceNum,omitempty"`

	// type of instance
	InstanceType string `json:"instanceType,omitempty"`

	// maximum number of instances running
	Peak int32 `json:"peak,omitempty"`

	// cloud provider of the stack
	Provider string `json:"provider,omitempty"`

	// region of the stack
	Region string `json:"region,omitempty"`

	// id of the stack
	StackID int64 `json:"stackId,omitempty"`

	// name of the stack
	StackName string `json:"stackName,omitempty"`

	// unique id of the cluster
	StackUUID string `json:"stackUuid,omitempty"`

	// ambari username
	Username string `json:"username,omitempty"`
}

// Validate validates this cloudbreak usage
func (m *CloudbreakUsage) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CloudbreakUsage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudbreakUsage) UnmarshalBinary(b []byte) error {
	var res CloudbreakUsage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

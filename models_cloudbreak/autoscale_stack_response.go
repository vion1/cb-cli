// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AutoscaleStackResponse autoscale stack response
// swagger:model AutoscaleStackResponse

type AutoscaleStackResponse struct {

	// account id of the resource owner that is provided by OAuth provider
	Account string `json:"account,omitempty"`

	// public ambari ip of the stack
	AmbariServerIP string `json:"ambariServerIp,omitempty"`

	// status of the cluster
	ClusterStatus string `json:"clusterStatus,omitempty"`

	// creation time of the stack in long
	Created int64 `json:"created,omitempty"`

	// port of the gateway secured proxy
	GatewayPort int32 `json:"gatewayPort,omitempty"`

	// name of the stack
	// Required: true
	// Max Length: 40
	// Min Length: 5
	// Pattern: (^[a-z][-a-z0-9]*[a-z0-9]$)
	Name *string `json:"name"`

	// id of the resource owner that is provided by OAuth provider
	Owner string `json:"owner,omitempty"`

	// ambari password
	Password string `json:"password,omitempty"`

	// id of the stack
	StackID int64 `json:"stackId,omitempty"`

	// status of the stack
	Status string `json:"status,omitempty"`

	// ambari username
	UserName string `json:"userName,omitempty"`
}

/* polymorph AutoscaleStackResponse account false */

/* polymorph AutoscaleStackResponse ambariServerIp false */

/* polymorph AutoscaleStackResponse clusterStatus false */

/* polymorph AutoscaleStackResponse created false */

/* polymorph AutoscaleStackResponse gatewayPort false */

/* polymorph AutoscaleStackResponse name false */

/* polymorph AutoscaleStackResponse owner false */

/* polymorph AutoscaleStackResponse password false */

/* polymorph AutoscaleStackResponse stackId false */

/* polymorph AutoscaleStackResponse status false */

/* polymorph AutoscaleStackResponse userName false */

// Validate validates this autoscale stack response
func (m *AutoscaleStackResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var autoscaleStackResponseTypeClusterStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["REQUESTED","CREATE_IN_PROGRESS","AVAILABLE","UPDATE_IN_PROGRESS","UPDATE_REQUESTED","UPDATE_FAILED","CREATE_FAILED","ENABLE_SECURITY_FAILED","PRE_DELETE_IN_PROGRESS","DELETE_IN_PROGRESS","DELETE_FAILED","DELETE_COMPLETED","STOPPED","STOP_REQUESTED","START_REQUESTED","STOP_IN_PROGRESS","START_IN_PROGRESS","START_FAILED","STOP_FAILED","WAIT_FOR_SYNC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		autoscaleStackResponseTypeClusterStatusPropEnum = append(autoscaleStackResponseTypeClusterStatusPropEnum, v)
	}
}

const (
	// AutoscaleStackResponseClusterStatusREQUESTED captures enum value "REQUESTED"
	AutoscaleStackResponseClusterStatusREQUESTED string = "REQUESTED"
	// AutoscaleStackResponseClusterStatusCREATEINPROGRESS captures enum value "CREATE_IN_PROGRESS"
	AutoscaleStackResponseClusterStatusCREATEINPROGRESS string = "CREATE_IN_PROGRESS"
	// AutoscaleStackResponseClusterStatusAVAILABLE captures enum value "AVAILABLE"
	AutoscaleStackResponseClusterStatusAVAILABLE string = "AVAILABLE"
	// AutoscaleStackResponseClusterStatusUPDATEINPROGRESS captures enum value "UPDATE_IN_PROGRESS"
	AutoscaleStackResponseClusterStatusUPDATEINPROGRESS string = "UPDATE_IN_PROGRESS"
	// AutoscaleStackResponseClusterStatusUPDATEREQUESTED captures enum value "UPDATE_REQUESTED"
	AutoscaleStackResponseClusterStatusUPDATEREQUESTED string = "UPDATE_REQUESTED"
	// AutoscaleStackResponseClusterStatusUPDATEFAILED captures enum value "UPDATE_FAILED"
	AutoscaleStackResponseClusterStatusUPDATEFAILED string = "UPDATE_FAILED"
	// AutoscaleStackResponseClusterStatusCREATEFAILED captures enum value "CREATE_FAILED"
	AutoscaleStackResponseClusterStatusCREATEFAILED string = "CREATE_FAILED"
	// AutoscaleStackResponseClusterStatusENABLESECURITYFAILED captures enum value "ENABLE_SECURITY_FAILED"
	AutoscaleStackResponseClusterStatusENABLESECURITYFAILED string = "ENABLE_SECURITY_FAILED"
	// AutoscaleStackResponseClusterStatusPREDELETEINPROGRESS captures enum value "PRE_DELETE_IN_PROGRESS"
	AutoscaleStackResponseClusterStatusPREDELETEINPROGRESS string = "PRE_DELETE_IN_PROGRESS"
	// AutoscaleStackResponseClusterStatusDELETEINPROGRESS captures enum value "DELETE_IN_PROGRESS"
	AutoscaleStackResponseClusterStatusDELETEINPROGRESS string = "DELETE_IN_PROGRESS"
	// AutoscaleStackResponseClusterStatusDELETEFAILED captures enum value "DELETE_FAILED"
	AutoscaleStackResponseClusterStatusDELETEFAILED string = "DELETE_FAILED"
	// AutoscaleStackResponseClusterStatusDELETECOMPLETED captures enum value "DELETE_COMPLETED"
	AutoscaleStackResponseClusterStatusDELETECOMPLETED string = "DELETE_COMPLETED"
	// AutoscaleStackResponseClusterStatusSTOPPED captures enum value "STOPPED"
	AutoscaleStackResponseClusterStatusSTOPPED string = "STOPPED"
	// AutoscaleStackResponseClusterStatusSTOPREQUESTED captures enum value "STOP_REQUESTED"
	AutoscaleStackResponseClusterStatusSTOPREQUESTED string = "STOP_REQUESTED"
	// AutoscaleStackResponseClusterStatusSTARTREQUESTED captures enum value "START_REQUESTED"
	AutoscaleStackResponseClusterStatusSTARTREQUESTED string = "START_REQUESTED"
	// AutoscaleStackResponseClusterStatusSTOPINPROGRESS captures enum value "STOP_IN_PROGRESS"
	AutoscaleStackResponseClusterStatusSTOPINPROGRESS string = "STOP_IN_PROGRESS"
	// AutoscaleStackResponseClusterStatusSTARTINPROGRESS captures enum value "START_IN_PROGRESS"
	AutoscaleStackResponseClusterStatusSTARTINPROGRESS string = "START_IN_PROGRESS"
	// AutoscaleStackResponseClusterStatusSTARTFAILED captures enum value "START_FAILED"
	AutoscaleStackResponseClusterStatusSTARTFAILED string = "START_FAILED"
	// AutoscaleStackResponseClusterStatusSTOPFAILED captures enum value "STOP_FAILED"
	AutoscaleStackResponseClusterStatusSTOPFAILED string = "STOP_FAILED"
	// AutoscaleStackResponseClusterStatusWAITFORSYNC captures enum value "WAIT_FOR_SYNC"
	AutoscaleStackResponseClusterStatusWAITFORSYNC string = "WAIT_FOR_SYNC"
)

// prop value enum
func (m *AutoscaleStackResponse) validateClusterStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, autoscaleStackResponseTypeClusterStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AutoscaleStackResponse) validateClusterStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateClusterStatusEnum("clusterStatus", "body", m.ClusterStatus); err != nil {
		return err
	}

	return nil
}

func (m *AutoscaleStackResponse) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 5); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 40); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(*m.Name), `(^[a-z][-a-z0-9]*[a-z0-9]$)`); err != nil {
		return err
	}

	return nil
}

var autoscaleStackResponseTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["REQUESTED","CREATE_IN_PROGRESS","AVAILABLE","UPDATE_IN_PROGRESS","UPDATE_REQUESTED","UPDATE_FAILED","CREATE_FAILED","ENABLE_SECURITY_FAILED","PRE_DELETE_IN_PROGRESS","DELETE_IN_PROGRESS","DELETE_FAILED","DELETE_COMPLETED","STOPPED","STOP_REQUESTED","START_REQUESTED","STOP_IN_PROGRESS","START_IN_PROGRESS","START_FAILED","STOP_FAILED","WAIT_FOR_SYNC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		autoscaleStackResponseTypeStatusPropEnum = append(autoscaleStackResponseTypeStatusPropEnum, v)
	}
}

const (
	// AutoscaleStackResponseStatusREQUESTED captures enum value "REQUESTED"
	AutoscaleStackResponseStatusREQUESTED string = "REQUESTED"
	// AutoscaleStackResponseStatusCREATEINPROGRESS captures enum value "CREATE_IN_PROGRESS"
	AutoscaleStackResponseStatusCREATEINPROGRESS string = "CREATE_IN_PROGRESS"
	// AutoscaleStackResponseStatusAVAILABLE captures enum value "AVAILABLE"
	AutoscaleStackResponseStatusAVAILABLE string = "AVAILABLE"
	// AutoscaleStackResponseStatusUPDATEINPROGRESS captures enum value "UPDATE_IN_PROGRESS"
	AutoscaleStackResponseStatusUPDATEINPROGRESS string = "UPDATE_IN_PROGRESS"
	// AutoscaleStackResponseStatusUPDATEREQUESTED captures enum value "UPDATE_REQUESTED"
	AutoscaleStackResponseStatusUPDATEREQUESTED string = "UPDATE_REQUESTED"
	// AutoscaleStackResponseStatusUPDATEFAILED captures enum value "UPDATE_FAILED"
	AutoscaleStackResponseStatusUPDATEFAILED string = "UPDATE_FAILED"
	// AutoscaleStackResponseStatusCREATEFAILED captures enum value "CREATE_FAILED"
	AutoscaleStackResponseStatusCREATEFAILED string = "CREATE_FAILED"
	// AutoscaleStackResponseStatusENABLESECURITYFAILED captures enum value "ENABLE_SECURITY_FAILED"
	AutoscaleStackResponseStatusENABLESECURITYFAILED string = "ENABLE_SECURITY_FAILED"
	// AutoscaleStackResponseStatusPREDELETEINPROGRESS captures enum value "PRE_DELETE_IN_PROGRESS"
	AutoscaleStackResponseStatusPREDELETEINPROGRESS string = "PRE_DELETE_IN_PROGRESS"
	// AutoscaleStackResponseStatusDELETEINPROGRESS captures enum value "DELETE_IN_PROGRESS"
	AutoscaleStackResponseStatusDELETEINPROGRESS string = "DELETE_IN_PROGRESS"
	// AutoscaleStackResponseStatusDELETEFAILED captures enum value "DELETE_FAILED"
	AutoscaleStackResponseStatusDELETEFAILED string = "DELETE_FAILED"
	// AutoscaleStackResponseStatusDELETECOMPLETED captures enum value "DELETE_COMPLETED"
	AutoscaleStackResponseStatusDELETECOMPLETED string = "DELETE_COMPLETED"
	// AutoscaleStackResponseStatusSTOPPED captures enum value "STOPPED"
	AutoscaleStackResponseStatusSTOPPED string = "STOPPED"
	// AutoscaleStackResponseStatusSTOPREQUESTED captures enum value "STOP_REQUESTED"
	AutoscaleStackResponseStatusSTOPREQUESTED string = "STOP_REQUESTED"
	// AutoscaleStackResponseStatusSTARTREQUESTED captures enum value "START_REQUESTED"
	AutoscaleStackResponseStatusSTARTREQUESTED string = "START_REQUESTED"
	// AutoscaleStackResponseStatusSTOPINPROGRESS captures enum value "STOP_IN_PROGRESS"
	AutoscaleStackResponseStatusSTOPINPROGRESS string = "STOP_IN_PROGRESS"
	// AutoscaleStackResponseStatusSTARTINPROGRESS captures enum value "START_IN_PROGRESS"
	AutoscaleStackResponseStatusSTARTINPROGRESS string = "START_IN_PROGRESS"
	// AutoscaleStackResponseStatusSTARTFAILED captures enum value "START_FAILED"
	AutoscaleStackResponseStatusSTARTFAILED string = "START_FAILED"
	// AutoscaleStackResponseStatusSTOPFAILED captures enum value "STOP_FAILED"
	AutoscaleStackResponseStatusSTOPFAILED string = "STOP_FAILED"
	// AutoscaleStackResponseStatusWAITFORSYNC captures enum value "WAIT_FOR_SYNC"
	AutoscaleStackResponseStatusWAITFORSYNC string = "WAIT_FOR_SYNC"
)

// prop value enum
func (m *AutoscaleStackResponse) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, autoscaleStackResponseTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AutoscaleStackResponse) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AutoscaleStackResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AutoscaleStackResponse) UnmarshalBinary(b []byte) error {
	var res AutoscaleStackResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

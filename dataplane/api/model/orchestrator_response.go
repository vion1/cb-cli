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

// OrchestratorResponse orchestrator response
// swagger:model OrchestratorResponse
type OrchestratorResponse struct {

	// endpoint for the container orchestration api
	APIEndpoint string `json:"apiEndpoint,omitempty"`

	// orchestrator specific parameters, like authentication details
	Parameters map[string]interface{} `json:"parameters,omitempty"`

	// type of the orchestrator
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this orchestrator response
func (m *OrchestratorResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OrchestratorResponse) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OrchestratorResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OrchestratorResponse) UnmarshalBinary(b []byte) error {
	var res OrchestratorResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

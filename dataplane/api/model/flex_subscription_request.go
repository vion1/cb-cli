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

// FlexSubscriptionRequest flex subscription request
// swagger:model FlexSubscriptionRequest
type FlexSubscriptionRequest struct {

	// name of the resource
	// Required: true
	Name *string `json:"name"`

	// Identifier of SmartSense subscription Cloudbreak domain object json representation.
	// Read Only: true
	SmartSenseSubscriptionID int64 `json:"smartSenseSubscriptionId,omitempty"`

	// Identifier of Flex subscription.
	// Read Only: true
	// Pattern: ^(FLEX-[0-9]{10}$)
	SubscriptionID string `json:"subscriptionId,omitempty"`

	// true if the flex subscription is the default one
	UsedAsDefault *bool `json:"usedAsDefault,omitempty"`

	// true if the flex subscription was used for the controller
	UsedForController *bool `json:"usedForController,omitempty"`
}

// Validate validates this flex subscription request
func (m *FlexSubscriptionRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriptionID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FlexSubscriptionRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *FlexSubscriptionRequest) validateSubscriptionID(formats strfmt.Registry) error {

	if swag.IsZero(m.SubscriptionID) { // not required
		return nil
	}

	if err := validate.Pattern("subscriptionId", "body", string(m.SubscriptionID), `^(FLEX-[0-9]{10}$)`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FlexSubscriptionRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FlexSubscriptionRequest) UnmarshalBinary(b []byte) error {
	var res FlexSubscriptionRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

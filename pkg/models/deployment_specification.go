// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DeploymentSpecification deployment specification
// swagger:model DeploymentSpecification
type DeploymentSpecification struct {

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Required: true
	Name *string `json:"name"`

	// The id of the project the current user belongs to.
	// Required: true
	ProjectID *string `json:"projectId"`
}

// Validate validates this deployment specification
func (m *DeploymentSpecification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DeploymentSpecification) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *DeploymentSpecification) validateProjectID(formats strfmt.Registry) error {

	if err := validate.Required("projectId", "body", m.ProjectID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DeploymentSpecification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeploymentSpecification) UnmarshalBinary(b []byte) error {
	var res DeploymentSpecification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

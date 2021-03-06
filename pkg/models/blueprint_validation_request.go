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

// BlueprintValidationRequest BlueprintValidationRequest
// swagger:model BlueprintValidationRequest
type BlueprintValidationRequest struct {

	// Blueprint Id
	// Format: uuid
	BlueprintID strfmt.UUID `json:"blueprintId,omitempty"`

	// Blueprint YAML content
	Content string `json:"content,omitempty"`

	// Blueprint request inputs
	Inputs interface{} `json:"inputs,omitempty"`

	// Project Id
	ProjectID string `json:"projectId,omitempty"`
}

// Validate validates this blueprint validation request
func (m *BlueprintValidationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBlueprintID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BlueprintValidationRequest) validateBlueprintID(formats strfmt.Registry) error {

	if swag.IsZero(m.BlueprintID) { // not required
		return nil
	}

	if err := validate.FormatOf("blueprintId", "body", "uuid", m.BlueprintID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BlueprintValidationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BlueprintValidationRequest) UnmarshalBinary(b []byte) error {
	var res BlueprintValidationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

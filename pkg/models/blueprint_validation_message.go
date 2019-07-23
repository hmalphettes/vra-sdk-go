// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BlueprintValidationMessage BlueprintValidationMessage
// swagger:model BlueprintValidationMessage
type BlueprintValidationMessage struct {

	// Validation message
	// Read Only: true
	Message string `json:"message,omitempty"`

	// Metadata
	// Read Only: true
	Metadata map[string]string `json:"metadata,omitempty"`

	// Validation path
	// Read Only: true
	Path string `json:"path,omitempty"`

	// Resource name
	// Read Only: true
	ResourceName string `json:"resourceName,omitempty"`

	// Message type
	// Read Only: true
	// Enum: [INFO WARNING ERROR]
	Type string `json:"type,omitempty"`
}

// Validate validates this blueprint validation message
func (m *BlueprintValidationMessage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var blueprintValidationMessageTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["INFO","WARNING","ERROR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		blueprintValidationMessageTypeTypePropEnum = append(blueprintValidationMessageTypeTypePropEnum, v)
	}
}

const (

	// BlueprintValidationMessageTypeINFO captures enum value "INFO"
	BlueprintValidationMessageTypeINFO string = "INFO"

	// BlueprintValidationMessageTypeWARNING captures enum value "WARNING"
	BlueprintValidationMessageTypeWARNING string = "WARNING"

	// BlueprintValidationMessageTypeERROR captures enum value "ERROR"
	BlueprintValidationMessageTypeERROR string = "ERROR"
)

// prop value enum
func (m *BlueprintValidationMessage) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, blueprintValidationMessageTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *BlueprintValidationMessage) validateType(formats strfmt.Registry) error {

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
func (m *BlueprintValidationMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BlueprintValidationMessage) UnmarshalBinary(b []byte) error {
	var res BlueprintValidationMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

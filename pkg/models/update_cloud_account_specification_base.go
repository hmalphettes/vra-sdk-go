// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// UpdateCloudAccountSpecificationBase update cloud account specification base
// swagger:model UpdateCloudAccountSpecificationBase
type UpdateCloudAccountSpecificationBase struct {

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	Name string `json:"name,omitempty"`

	// A set of tag keys and optional values to set on the Cloud Account
	Tags []*Tag `json:"tags"`
}

// Validate validates this update cloud account specification base
func (m *UpdateCloudAccountSpecificationBase) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateCloudAccountSpecificationBase) validateTags(formats strfmt.Registry) error {

	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateCloudAccountSpecificationBase) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateCloudAccountSpecificationBase) UnmarshalBinary(b []byte) error {
	var res UpdateCloudAccountSpecificationBase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

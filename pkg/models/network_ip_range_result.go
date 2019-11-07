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

// NetworkIPRangeResult State object representing a query result of network IP range.
// swagger:model NetworkIPRangeResult
type NetworkIPRangeResult struct {

	// List of content items
	// Read Only: true
	Content []*NetworkIPRange `json:"content"`

	// Number of elements in the current page
	// Read Only: true
	NumberOfElements int64 `json:"numberOfElements,omitempty"`

	// Total number of elements. In some cases the field may not be populated
	// Read Only: true
	TotalElements int64 `json:"totalElements,omitempty"`
}

// Validate validates this network IP range result
func (m *NetworkIPRangeResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContent(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkIPRangeResult) validateContent(formats strfmt.Registry) error {

	if swag.IsZero(m.Content) { // not required
		return nil
	}

	for i := 0; i < len(m.Content); i++ {
		if swag.IsZero(m.Content[i]) { // not required
			continue
		}

		if m.Content[i] != nil {
			if err := m.Content[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("content" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkIPRangeResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkIPRangeResult) UnmarshalBinary(b []byte) error {
	var res NetworkIPRangeResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
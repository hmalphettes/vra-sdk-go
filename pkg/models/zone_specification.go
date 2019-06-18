// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ZoneSpecification Specification for a zone.
// swagger:model ZoneSpecification
type ZoneSpecification struct {

	// A list of key value pair of properties that will  be used
	CustomProperties map[string]string `json:"customProperties,omitempty"`

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Required: true
	Name *string `json:"name"`

	// Placement policy for the zone. One of DEFAULT, SPREAD or BINPACK.
	PlacementPolicy string `json:"placementPolicy,omitempty"`

	// The id of the region for which this profile is created
	// Required: true
	RegionID *string `json:"regionId"`

	// A set of tag keys and optional values that are effectively applied to all compute resources in this zone, but only in the context of this zone.
	Tags []*Tag `json:"tags"`

	// A set of tag keys and optional values that will be used
	TagsToMatch []*Tag `json:"tagsToMatch"`
}

// Validate validates this zone specification
func (m *ZoneSpecification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTagsToMatch(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ZoneSpecification) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ZoneSpecification) validateRegionID(formats strfmt.Registry) error {

	if err := validate.Required("regionId", "body", m.RegionID); err != nil {
		return err
	}

	return nil
}

func (m *ZoneSpecification) validateTags(formats strfmt.Registry) error {

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

func (m *ZoneSpecification) validateTagsToMatch(formats strfmt.Registry) error {

	if swag.IsZero(m.TagsToMatch) { // not required
		return nil
	}

	for i := 0; i < len(m.TagsToMatch); i++ {
		if swag.IsZero(m.TagsToMatch[i]) { // not required
			continue
		}

		if m.TagsToMatch[i] != nil {
			if err := m.TagsToMatch[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tagsToMatch" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ZoneSpecification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ZoneSpecification) UnmarshalBinary(b []byte) error {
	var res ZoneSpecification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

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

// MachineSpecification Specification for a cloud agnostic machine.
// swagger:model MachineSpecification
type MachineSpecification struct {

	// A valid cloud config data in json-escaped yaml syntax
	BootConfig *MachineBootConfig `json:"bootConfig,omitempty"`

	// Constraints that are used to drive placement policies for the virtual machine that is produced from this specification. Constraint expressions are matched against tags on existing placement targets.
	Constraints []*Constraint `json:"constraints"`

	// Additional custom properties that may be used toextend the machine.
	CustomProperties map[string]string `json:"customProperties,omitempty"`

	// deployment Id
	DeploymentID string `json:"deploymentId,omitempty"`

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// A set of disk specifications for this machine.
	Disks []*DiskAttachmentSpecification `json:"disks"`

	// Flavor of machine instance.
	// Required: true
	Flavor *string `json:"flavor"`

	// Type of image used for this machine.
	// Required: true
	Image *string `json:"image"`

	// Constraints that are used to drive placement policies for the image disk. Constraint expressions are matched against tags on existing placement targets.
	ImageDiskConstraints []*Constraint `json:"imageDiskConstraints"`

	// Direct image reference used for this machine (name, path, location, uri, etc.). Valid if no image property is provided
	// Required: true
	ImageRef *string `json:"imageRef"`

	// Number of machines to provision - default 1.
	MachineCount int32 `json:"machineCount,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Required: true
	Name *string `json:"name"`

	// A set of network interface controller specifications for this machine. If not specified, then a default network connection will be created.
	Nics []*NetworkInterfaceSpecification `json:"nics"`

	// The id of the project the current user belongs to.
	// Required: true
	ProjectID *string `json:"projectId"`

	// A set of tag keys and optional values that should be set on any resource that is produced from this specification.
	Tags []*Tag `json:"tags"`
}

// Validate validates this machine specification
func (m *MachineSpecification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBootConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConstraints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlavor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageDiskConstraints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImageRef(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MachineSpecification) validateBootConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.BootConfig) { // not required
		return nil
	}

	if m.BootConfig != nil {
		if err := m.BootConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bootConfig")
			}
			return err
		}
	}

	return nil
}

func (m *MachineSpecification) validateConstraints(formats strfmt.Registry) error {

	if swag.IsZero(m.Constraints) { // not required
		return nil
	}

	for i := 0; i < len(m.Constraints); i++ {
		if swag.IsZero(m.Constraints[i]) { // not required
			continue
		}

		if m.Constraints[i] != nil {
			if err := m.Constraints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("constraints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MachineSpecification) validateDisks(formats strfmt.Registry) error {

	if swag.IsZero(m.Disks) { // not required
		return nil
	}

	for i := 0; i < len(m.Disks); i++ {
		if swag.IsZero(m.Disks[i]) { // not required
			continue
		}

		if m.Disks[i] != nil {
			if err := m.Disks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MachineSpecification) validateFlavor(formats strfmt.Registry) error {

	if err := validate.Required("flavor", "body", m.Flavor); err != nil {
		return err
	}

	return nil
}

func (m *MachineSpecification) validateImage(formats strfmt.Registry) error {

	if err := validate.Required("image", "body", m.Image); err != nil {
		return err
	}

	return nil
}

func (m *MachineSpecification) validateImageDiskConstraints(formats strfmt.Registry) error {

	if swag.IsZero(m.ImageDiskConstraints) { // not required
		return nil
	}

	for i := 0; i < len(m.ImageDiskConstraints); i++ {
		if swag.IsZero(m.ImageDiskConstraints[i]) { // not required
			continue
		}

		if m.ImageDiskConstraints[i] != nil {
			if err := m.ImageDiskConstraints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("imageDiskConstraints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MachineSpecification) validateImageRef(formats strfmt.Registry) error {

	if err := validate.Required("imageRef", "body", m.ImageRef); err != nil {
		return err
	}

	return nil
}

func (m *MachineSpecification) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *MachineSpecification) validateNics(formats strfmt.Registry) error {

	if swag.IsZero(m.Nics) { // not required
		return nil
	}

	for i := 0; i < len(m.Nics); i++ {
		if swag.IsZero(m.Nics[i]) { // not required
			continue
		}

		if m.Nics[i] != nil {
			if err := m.Nics[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MachineSpecification) validateProjectID(formats strfmt.Registry) error {

	if err := validate.Required("projectId", "body", m.ProjectID); err != nil {
		return err
	}

	return nil
}

func (m *MachineSpecification) validateTags(formats strfmt.Registry) error {

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
func (m *MachineSpecification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MachineSpecification) UnmarshalBinary(b []byte) error {
	var res MachineSpecification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

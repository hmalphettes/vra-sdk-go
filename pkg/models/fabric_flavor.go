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

// FabricFlavor Represents a fabric flavor from the corresponding cloud end-point
// swagger:model FabricFlavor
type FabricFlavor struct {

	// Size of the boot disk (in megabytes). Not populated when inapplicable.
	BootDiskSizeInMB int32 `json:"bootDiskSizeInMB,omitempty"`

	// Number of CPU cores. Not populated when inapplicable.
	CPUCount int32 `json:"cpuCount,omitempty"`

	// Number of data disks. Not populated when inapplicable.
	DataDiskMaxCount int32 `json:"dataDiskMaxCount,omitempty"`

	// Size of the data disks (in megabytes). Not populated when inapplicable.
	DataDiskSizeInMB int32 `json:"dataDiskSizeInMB,omitempty"`

	// The internal identification used by the corresponding cloud end-point
	ID string `json:"id,omitempty"`

	// Total amount of memory (in megabytes). Not populated when inapplicable.
	MemoryInMB int64 `json:"memoryInMB,omitempty"`

	// The value of the instance type in the corresponding cloud.
	// Required: true
	Name *string `json:"name"`

	// The type of network supported by this instance type. Not populated when inapplicable.
	NetworkType string `json:"networkType,omitempty"`

	// The type of storage supported by this instance type. Not populated when inapplicable.
	StorageType string `json:"storageType,omitempty"`
}

// Validate validates this fabric flavor
func (m *FabricFlavor) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FabricFlavor) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FabricFlavor) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FabricFlavor) UnmarshalBinary(b []byte) error {
	var res FabricFlavor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

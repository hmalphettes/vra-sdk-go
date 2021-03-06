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

// DataCollector State object representing a data collector.<br>The data collector is an OVA tool that contains the credentials and protocols needed to create a connection between a data collector appliance on a host vCenter and a vCenter-based cloud account.
// swagger:model DataCollector
type DataCollector struct {

	// Data collector identifier
	// Required: true
	DcID *string `json:"dcId"`

	// Data collector host name
	// Required: true
	HostName *string `json:"hostName"`

	// Ip Address of the data collector VM
	// Required: true
	IPAddress *string `json:"ipAddress"`

	// Data collector name
	// Required: true
	Name *string `json:"name"`

	// Current status of the data collector
	// Required: true
	Status *string `json:"status"`
}

// Validate validates this data collector
func (m *DataCollector) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDcID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIPAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataCollector) validateDcID(formats strfmt.Registry) error {

	if err := validate.Required("dcId", "body", m.DcID); err != nil {
		return err
	}

	return nil
}

func (m *DataCollector) validateHostName(formats strfmt.Registry) error {

	if err := validate.Required("hostName", "body", m.HostName); err != nil {
		return err
	}

	return nil
}

func (m *DataCollector) validateIPAddress(formats strfmt.Registry) error {

	if err := validate.Required("ipAddress", "body", m.IPAddress); err != nil {
		return err
	}

	return nil
}

func (m *DataCollector) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *DataCollector) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DataCollector) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataCollector) UnmarshalBinary(b []byte) error {
	var res DataCollector
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

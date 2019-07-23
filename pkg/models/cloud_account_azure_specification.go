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

// CloudAccountAzureSpecification Specification for a Azure cloud account.<br><br>A cloud account identifies a cloud account type and an account-specific deployment region or data center where the associated cloud account resources are hosted.
// swagger:model CloudAccountAzureSpecification
type CloudAccountAzureSpecification struct {

	// Azure Client Application ID
	// Required: true
	ClientApplicationID *string `json:"clientApplicationId"`

	// Azure Client Application Secret Key
	// Required: true
	ClientApplicationSecretKey *string `json:"clientApplicationSecretKey"`

	// Create default cloud zones for the enabled regions.
	CreateDefaultZones bool `json:"createDefaultZones,omitempty"`

	// A human-friendly description.
	Description string `json:"description,omitempty"`

	// A human-friendly name used as an identifier in APIs that support this option.
	// Required: true
	Name *string `json:"name"`

	// A set of Region names to enable provisioning on. Refer to /iaas/cloud-accounts-azure/region-enumeration..
	// Required: true
	RegionIds []string `json:"regionIds"`

	// Azure Subscribtion ID
	// Required: true
	SubscriptionID *string `json:"subscriptionId"`

	// A set of tag keys and optional values to set on the Cloud Account
	Tags []*Tag `json:"tags"`

	// Azure Tenant ID
	// Required: true
	TenantID *string `json:"tenantId"`
}

// Validate validates this cloud account azure specification
func (m *CloudAccountAzureSpecification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClientApplicationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientApplicationSecretKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegionIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriptionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenantID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudAccountAzureSpecification) validateClientApplicationID(formats strfmt.Registry) error {

	if err := validate.Required("clientApplicationId", "body", m.ClientApplicationID); err != nil {
		return err
	}

	return nil
}

func (m *CloudAccountAzureSpecification) validateClientApplicationSecretKey(formats strfmt.Registry) error {

	if err := validate.Required("clientApplicationSecretKey", "body", m.ClientApplicationSecretKey); err != nil {
		return err
	}

	return nil
}

func (m *CloudAccountAzureSpecification) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CloudAccountAzureSpecification) validateRegionIds(formats strfmt.Registry) error {

	if err := validate.Required("regionIds", "body", m.RegionIds); err != nil {
		return err
	}

	return nil
}

func (m *CloudAccountAzureSpecification) validateSubscriptionID(formats strfmt.Registry) error {

	if err := validate.Required("subscriptionId", "body", m.SubscriptionID); err != nil {
		return err
	}

	return nil
}

func (m *CloudAccountAzureSpecification) validateTags(formats strfmt.Registry) error {

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

func (m *CloudAccountAzureSpecification) validateTenantID(formats strfmt.Registry) error {

	if err := validate.Required("tenantId", "body", m.TenantID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudAccountAzureSpecification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudAccountAzureSpecification) UnmarshalBinary(b []byte) error {
	var res CloudAccountAzureSpecification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

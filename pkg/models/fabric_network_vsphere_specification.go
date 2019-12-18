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

// FabricNetworkVsphereSpecification Specification for updating a Vsphere FabricNetwork
// swagger:model FabricNetworkVsphereSpecification
type FabricNetworkVsphereSpecification struct {

	// Network CIDR to be used.
	Cidr string `json:"cidr,omitempty"`

	// IPv4 default gateway to be used.
	DefaultGateway string `json:"defaultGateway,omitempty"`

	// IPv6 default gateway to be used.
	DefaultIPV6Gateway string `json:"defaultIpv6Gateway,omitempty"`

	// A list of DNS search domains that were set on this resource instance.
	DNSSearchDomains []string `json:"dnsSearchDomains"`

	// A list of DNS server addresses that were set on this resource instance.
	DNSServerAddresses []string `json:"dnsServerAddresses"`

	// Domain value.
	Domain string `json:"domain,omitempty"`

	// Network IPv6 CIDR to be used.
	IPV6Cidr string `json:"ipv6Cidr,omitempty"`

	// Indicates whether this is the default subnet for the zone.
	IsDefault bool `json:"isDefault,omitempty"`

	// Indicates whether the sub-network supports public IP assignment.
	IsPublic bool `json:"isPublic,omitempty"`

	// A set of tag keys and optional values that were set on this resource instance.
	Tags []*Tag `json:"tags"`
}

// Validate validates this fabric network vsphere specification
func (m *FabricNetworkVsphereSpecification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FabricNetworkVsphereSpecification) validateTags(formats strfmt.Registry) error {

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
func (m *FabricNetworkVsphereSpecification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FabricNetworkVsphereSpecification) UnmarshalBinary(b []byte) error {
	var res FabricNetworkVsphereSpecification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// URI URI
// swagger:model URI
type URI struct {

	// absolute
	Absolute bool `json:"absolute,omitempty"`

	// authority
	Authority string `json:"authority,omitempty"`

	// fragment
	Fragment string `json:"fragment,omitempty"`

	// host
	Host string `json:"host,omitempty"`

	// opaque
	Opaque bool `json:"opaque,omitempty"`

	// path
	Path string `json:"path,omitempty"`

	// port
	Port int32 `json:"port,omitempty"`

	// query
	Query string `json:"query,omitempty"`

	// raw authority
	RawAuthority string `json:"rawAuthority,omitempty"`

	// raw fragment
	RawFragment string `json:"rawFragment,omitempty"`

	// raw path
	RawPath string `json:"rawPath,omitempty"`

	// raw query
	RawQuery string `json:"rawQuery,omitempty"`

	// raw scheme specific part
	RawSchemeSpecificPart string `json:"rawSchemeSpecificPart,omitempty"`

	// raw user info
	RawUserInfo string `json:"rawUserInfo,omitempty"`

	// scheme
	Scheme string `json:"scheme,omitempty"`

	// scheme specific part
	SchemeSpecificPart string `json:"schemeSpecificPart,omitempty"`

	// user info
	UserInfo string `json:"userInfo,omitempty"`
}

// Validate validates this URI
func (m *URI) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *URI) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *URI) UnmarshalBinary(b []byte) error {
	var res URI
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
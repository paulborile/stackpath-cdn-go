// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// SchemacdnOrigin schemacdn origin
// swagger:model schemacdnOrigin
type SchemacdnOrigin struct {

	// Whether or not the origin is dedicated for a site. Dedicated origins cannot be used by any site other than that which it is dedicated for.
	Dedicated bool `json:"dedicated"`

	// Origin hostname or IP address
	Hostname string `json:"hostname,omitempty"`

	// Origin ID
	ID string `json:"id,omitempty"`

	// Path on the origin, defaults to /
	Path string `json:"path,omitempty"`

	// The port to connect to for non-encrypted connections
	Port int32 `json:"port,omitempty"`

	// The port to connect  to for encrypted connections
	SecurePort int32 `json:"securePort,omitempty"`

	// The site the origin is dedicated for
	SiteID string `json:"siteId,omitempty"`
}

// Validate validates this schemacdn origin
func (m *SchemacdnOrigin) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SchemacdnOrigin) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SchemacdnOrigin) UnmarshalBinary(b []byte) error {
	var res SchemacdnOrigin
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

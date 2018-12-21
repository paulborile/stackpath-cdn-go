// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CustconfDynamicOrigin custconf dynamic origin
// swagger:model custconfDynamicOrigin
type CustconfDynamicOrigin struct {

	// String of values deliminated by a ',' character.
	AllowedDomains string `json:"allowedDomains,omitempty"`

	// enabled
	Enabled bool `json:"enabled"`

	// This is used by the API to perform conflict checking.
	ID string `json:"id,omitempty"`

	// query param
	QueryParam string `json:"queryParam,omitempty"`
}

// Validate validates this custconf dynamic origin
func (m *CustconfDynamicOrigin) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustconfDynamicOrigin) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustconfDynamicOrigin) UnmarshalBinary(b []byte) error {
	var res CustconfDynamicOrigin
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CustconfCrossDomain This allows you to override the default cross domain policy file
// (crossdomain.xml) delivered by the CDN caching servers.
// swagger:model custconfCrossDomain
type CustconfCrossDomain struct {

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// The contents of the cross domain file you want delivered instead of the
	// default
	File string `json:"file,omitempty"`

	// This is used by the API to perform conflict checking.
	ID string `json:"id,omitempty"`
}

// Validate validates this custconf cross domain
func (m *CustconfCrossDomain) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustconfCrossDomain) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustconfCrossDomain) UnmarshalBinary(b []byte) error {
	var res CustconfCrossDomain
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
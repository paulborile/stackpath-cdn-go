// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CustconfFlv The flash initial bytes policy allows you to force the CDN to send the
// initial bytes of a FLV file which contains the header information that is
// used when jumping to different offsets in the file.
// swagger:model custconfFlv
type CustconfFlv struct {

	// enabled
	Enabled bool `json:"enabled"`

	// This is used by the API to perform conflict checking.
	ID string `json:"id,omitempty"`

	// This setting is typically set to 13 bytes.
	InitByteSize string `json:"initByteSize,omitempty"`
}

// Validate validates this custconf flv
func (m *CustconfFlv) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustconfFlv) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustconfFlv) UnmarshalBinary(b []byte) error {
	var res CustconfFlv
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

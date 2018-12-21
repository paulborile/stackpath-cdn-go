// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CustconfNrtReporting custconf nrt reporting
// swagger:model custconfNrtReporting
type CustconfNrtReporting struct {

	// enabled
	Enabled bool `json:"enabled"`

	// This is used by the API to perform conflict checking.
	ID string `json:"id,omitempty"`

	// report v host
	ReportVHost bool `json:"reportVHost"`
}

// Validate validates this custconf nrt reporting
func (m *CustconfNrtReporting) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustconfNrtReporting) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustconfNrtReporting) UnmarshalBinary(b []byte) error {
	var res CustconfNrtReporting
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

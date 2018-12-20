// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CustconfCloseHalfOpenConnections Instructs the CDN caching server to fully close the connection immediately
// after receiving a TCP FIN from the client.
// swagger:model custconfCloseHalfOpenConnections
type CustconfCloseHalfOpenConnections struct {

	// Force the close of client connections upon receiving TCP FIN from clients.
	Enabled bool `json:"enabled,omitempty"`

	// This is used by the API to perform conflict checking.
	ID string `json:"id,omitempty"`
}

// Validate validates this custconf close half open connections
func (m *CustconfCloseHalfOpenConnections) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustconfCloseHalfOpenConnections) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustconfCloseHalfOpenConnections) UnmarshalBinary(b []byte) error {
	var res CustconfCloseHalfOpenConnections
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CustconfAuthURLSignAKv1 The Akamai URL Signing v1 policy allows you to create a signed URL that
// implements the same signing method used by Akamai; therefore, published URLs
// from an Akamai CDN network can be transitioned to the Highwinds network
// without you having to change your signing methods.
// swagger:model custconfAuthUrlSignAKv1
type CustconfAuthURLSignAKv1 struct {

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// extract
	Extract string `json:"extract,omitempty"`

	// String of values deliminated by a ',' character.
	HeaderFilter string `json:"headerFilter,omitempty"`

	// This is used by the API to perform conflict checking.
	ID string `json:"id,omitempty"`

	// String of values deliminated by a ',' character.
	MethodFilter string `json:"methodFilter,omitempty"`

	// The authentication parameter defines the query string parameter in the
	// request URL that contains the authentication information.
	Param string `json:"param,omitempty"`

	// String of values deliminated by a ',' character.
	PathFilter string `json:"pathFilter,omitempty"`

	// The salt is used as a shared secret in the signing process. This value
	// should only be known by Highwinds and by systems authorized to sign your
	// content.
	Salt string `json:"salt,omitempty"`
}

// Validate validates this custconf auth Url sign a kv1
func (m *CustconfAuthURLSignAKv1) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustconfAuthURLSignAKv1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustconfAuthURLSignAKv1) UnmarshalBinary(b []byte) error {
	var res CustconfAuthURLSignAKv1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
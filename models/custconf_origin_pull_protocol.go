// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CustconfOriginPullProtocol The Origin Pull Protocol allows you to configure the CDN caching servers to
// use secured or non-secured connection to Origin.
// swagger:model custconfOriginPullProtocol
type CustconfOriginPullProtocol struct {

	// This key allows you to configure the CDN caching servers to use SNI while
	// making Secured Connection to Origin.
	EnableSNI bool `json:"enableSNI"`

	// This is used by the API to perform conflict checking.
	ID string `json:"id,omitempty"`

	// protocol
	Protocol CustconfOriginPullProtocolProtocolEnumWrapperValue `json:"protocol,omitempty"`
}

// Validate validates this custconf origin pull protocol
func (m *CustconfOriginPullProtocol) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustconfOriginPullProtocol) validateProtocol(formats strfmt.Registry) error {

	if swag.IsZero(m.Protocol) { // not required
		return nil
	}

	if err := m.Protocol.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("protocol")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CustconfOriginPullProtocol) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustconfOriginPullProtocol) UnmarshalBinary(b []byte) error {
	var res CustconfOriginPullProtocol
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

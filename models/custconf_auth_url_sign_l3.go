// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CustconfAuthURLSignL3 The Level 3 URL Signing policy allows you to create a signed URL that
// implements the same signing method used by Level 3; therefore, published URLs
// from an Level 3 CDN network can be transitioned to the Highwinds network
// without you having to change your signing methods.
// swagger:model custconfAuthUrlSignL3
type CustconfAuthURLSignL3 struct {

	// Provides the capability to rename the query string parameter that is used
	// to inject the Client's IP address into the hashing algorithm. This
	// configuration is only applicable when injectClientIPAddress is set to true.
	ClientIPAddressField string `json:"clientIPAddressField,omitempty"`

	// enabled
	Enabled bool `json:"enabled"`

	// String of values deliminated by a ',' character. A list of patterns that
	// are used to describe query string parameters that should be omitted from
	// the hashing algorithm if contained in the URL. A asterisk '*' by itself
	// indicates to exclude all query string parameters from the hashing
	// algorithm. The tokenField is always excluded. On the other hand, the
	// startField and/or expireField are always included in the hashing algorithm
	// if present in the request even if listed here. Users may explicitly specify
	// parameters to keep (not exclude) by preceding the glob with an exclamation
	// "!". This may be useful if a User wants to exclude all query string
	// parameters except one ore more known parameters.  For example, a value of
	// '*,!version' means exclude all parameters except "version".
	ExcludedParameters string `json:"excludedParameters,omitempty"`

	// This is the name of the query string parameter that contains the time after
	// which the URL is considered invalid. If defined, requests must contain the
	// parameter, and its value must be in the future.
	ExpireField string `json:"expireField,omitempty"`

	// String of values deliminated by a ',' character.
	HeaderFilter string `json:"headerFilter,omitempty"`

	// This is used by the API to perform conflict checking.
	ID string `json:"id,omitempty"`

	// Indicates whether or not to include the Host without the request Protocol
	// when calculating the signature.
	IncludeHostOnly bool `json:"includeHostOnly"`

	// Indicates whether or not to include both the Protocol and Host when
	// calculating the signature.
	IncludeProtocolAndHost bool `json:"includeProtocolAndHost"`

	// Indicates whether or not to include the Client's IP address when
	// calculating the signature.
	InjectClientIPAddress bool `json:"injectClientIPAddress"`

	// String of values deliminated by a ',' character.
	MethodFilter string `json:"methodFilter,omitempty"`

	// String of values deliminated by a ',' character.
	PathFilter string `json:"pathFilter,omitempty"`

	// String of values deliminated by a ',' character. An ordered list of shared
	// secrets. The order is CRITICAL and it MUST be identical to the ordered
	// table used by the Client.
	SharedSecretTable string `json:"sharedSecretTable,omitempty"`

	// The name of the query string parameter that contains the start time when
	// the request is considered valid.
	StartField string `json:"startField,omitempty"`

	// time format
	TimeFormat AuthURLSignL3TimeFormatEnumWrapperValue `json:"timeFormat,omitempty"`

	// This is the name of the query string parameter that will be used by the
	// publisher to specify the signature for the URL.
	TokenField string `json:"tokenField,omitempty"`
}

// Validate validates this custconf auth Url sign l3
func (m *CustconfAuthURLSignL3) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTimeFormat(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustconfAuthURLSignL3) validateTimeFormat(formats strfmt.Registry) error {

	if swag.IsZero(m.TimeFormat) { // not required
		return nil
	}

	if err := m.TimeFormat.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("timeFormat")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CustconfAuthURLSignL3) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustconfAuthURLSignL3) UnmarshalBinary(b []byte) error {
	var res CustconfAuthURLSignL3
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

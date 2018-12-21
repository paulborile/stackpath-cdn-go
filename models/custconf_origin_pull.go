// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CustconfOriginPull The origin pull settings policy contains a list of settings that control the
// behavior of origin pull requests.
// swagger:model custconfOriginPull
type CustconfOriginPull struct {

	// default behavior
	DefaultBehavior OriginPullDefaultBehaviorEnumWrapperValue `json:"defaultBehavior,omitempty"`

	// enabled
	Enabled bool `json:"enabled"`

	// This is used by the API to perform conflict checking.
	ID string `json:"id,omitempty"`

	// GFS sends a path without any query string parameters when making external
	// origin requests regardless if any parameters were sent by the User-Agent.
	NoQSParams bool `json:"noQSParams"`

	// pass all headers on dedup
	PassAllHeadersOnDedup bool `json:"passAllHeadersOnDedup"`

	// redirect action
	RedirectAction OriginPullRedirectActionEnumWrapperValue `json:"redirectAction,omitempty"`

	// String of values deliminated by a ',' character. List of HTTP Methods that
	// define types of origin pull requests that can be retried if a failure
	// occurs after sending a previous request.
	RetryMethods string `json:"retryMethods,omitempty"`

	// The number of seconds for an Edge server to may wait for a Shield to
	// respond.  This value overrides the value CDN dynamically calculates per
	// shield request. This policy does not affect how long to wait for connecting
	// to an origin.
	ShieldResponseTimeoutOverride int64 `json:"shieldResponseTimeoutOverride,omitempty"`

	// This value instructs the CDN to use the original Host header from the
	// client request when revalidating "no-cache" assets against your origin.
	TransparentMode bool `json:"transparentMode"`
}

// Validate validates this custconf origin pull
func (m *CustconfOriginPull) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefaultBehavior(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRedirectAction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustconfOriginPull) validateDefaultBehavior(formats strfmt.Registry) error {

	if swag.IsZero(m.DefaultBehavior) { // not required
		return nil
	}

	if err := m.DefaultBehavior.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("defaultBehavior")
		}
		return err
	}

	return nil
}

func (m *CustconfOriginPull) validateRedirectAction(formats strfmt.Registry) error {

	if swag.IsZero(m.RedirectAction) { // not required
		return nil
	}

	if err := m.RedirectAction.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("redirectAction")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CustconfOriginPull) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustconfOriginPull) UnmarshalBinary(b []byte) error {
	var res CustconfOriginPull
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

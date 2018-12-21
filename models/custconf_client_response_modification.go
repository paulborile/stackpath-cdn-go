// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CustconfClientResponseModification custconf client response modification
// swagger:model custconfClientResponseModification
type CustconfClientResponseModification struct {

	// String of values deliminated by a '|' character. This is the static HTTP
	// header you want inserted into the CDN response. Start value with "append:"
	// or "replace:" which defines if Header will be replace or added. Default is
	// append.
	AddHeaders string `json:"addHeaders,omitempty"`

	// client request filter
	ClientRequestFilter []*CustconfRequestFilter `json:"clientRequestFilter"`

	// client response filter
	ClientResponseFilter []*CustconfResponseFilter `json:"clientResponseFilter"`

	// enabled
	Enabled bool `json:"enabled"`

	// flow control
	FlowControl CustconfClientResponseModificationFlowControlEnumWrapperValue `json:"flowControl,omitempty"`

	// String of values deliminated by a ',' character.
	HeaderFilter string `json:"headerFilter,omitempty"`

	// header pattern
	HeaderPattern string `json:"headerPattern,omitempty"`

	// header rewrite
	HeaderRewrite string `json:"headerRewrite,omitempty"`

	// This is used by the API to perform conflict checking.
	ID string `json:"id,omitempty"`

	// String of values deliminated by a ',' character.
	MethodFilter string `json:"methodFilter,omitempty"`

	// String of values deliminated by a ',' character.
	PathFilter string `json:"pathFilter,omitempty"`

	// status code rewrite
	StatusCodeRewrite int64 `json:"statusCodeRewrite,omitempty"`
}

// Validate validates this custconf client response modification
func (m *CustconfClientResponseModification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClientRequestFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClientResponseFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlowControl(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustconfClientResponseModification) validateClientRequestFilter(formats strfmt.Registry) error {

	if swag.IsZero(m.ClientRequestFilter) { // not required
		return nil
	}

	for i := 0; i < len(m.ClientRequestFilter); i++ {
		if swag.IsZero(m.ClientRequestFilter[i]) { // not required
			continue
		}

		if m.ClientRequestFilter[i] != nil {
			if err := m.ClientRequestFilter[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clientRequestFilter" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CustconfClientResponseModification) validateClientResponseFilter(formats strfmt.Registry) error {

	if swag.IsZero(m.ClientResponseFilter) { // not required
		return nil
	}

	for i := 0; i < len(m.ClientResponseFilter); i++ {
		if swag.IsZero(m.ClientResponseFilter[i]) { // not required
			continue
		}

		if m.ClientResponseFilter[i] != nil {
			if err := m.ClientResponseFilter[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("clientResponseFilter" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CustconfClientResponseModification) validateFlowControl(formats strfmt.Registry) error {

	if swag.IsZero(m.FlowControl) { // not required
		return nil
	}

	if err := m.FlowControl.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("flowControl")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CustconfClientResponseModification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustconfClientResponseModification) UnmarshalBinary(b []byte) error {
	var res CustconfClientResponseModification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

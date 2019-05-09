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

// CdnGetMetricsResponse The response from a request to retrieve CDN metrics from a stack
// swagger:model cdnGetMetricsResponse
type CdnGetMetricsResponse struct {

	// The requested CDN metrics
	Series []*GetMetricsResponseMetricSeries `json:"series,omitempty"`
}

// Validate validates this cdn get metrics response
func (m *CdnGetMetricsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSeries(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CdnGetMetricsResponse) validateSeries(formats strfmt.Registry) error {

	if swag.IsZero(m.Series) { // not required
		return nil
	}

	for i := 0; i < len(m.Series); i++ {
		if swag.IsZero(m.Series[i]) { // not required
			continue
		}

		if m.Series[i] != nil {
			if err := m.Series[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("series" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CdnGetMetricsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CdnGetMetricsResponse) UnmarshalBinary(b []byte) error {
	var res CdnGetMetricsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

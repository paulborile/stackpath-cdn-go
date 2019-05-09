// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// GetMetricsResponseMetricSample The data points associated with a series of metrics
// swagger:model GetMetricsResponseMetricSample
type GetMetricsResponseMetricSample struct {

	// An individual data point
	Values []float64 `json:"values,omitempty"`
}

// Validate validates this get metrics response metric sample
func (m *GetMetricsResponseMetricSample) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetMetricsResponseMetricSample) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetMetricsResponseMetricSample) UnmarshalBinary(b []byte) error {
	var res GetMetricsResponseMetricSample
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

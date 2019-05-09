// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

// VaryHeaderFieldProxyBehaviorEnumWrapperValue vary header field proxy behavior enum wrapper value
// swagger:model VaryHeaderFieldProxyBehaviorEnumWrapperValue
type VaryHeaderFieldProxyBehaviorEnumWrapperValue string

const (

	// VaryHeaderFieldProxyBehaviorEnumWrapperValueUNKNOWN captures enum value "UNKNOWN"
	VaryHeaderFieldProxyBehaviorEnumWrapperValueUNKNOWN VaryHeaderFieldProxyBehaviorEnumWrapperValue = "UNKNOWN"

	// VaryHeaderFieldProxyBehaviorEnumWrapperValueWhole captures enum value "whole"
	VaryHeaderFieldProxyBehaviorEnumWrapperValueWhole VaryHeaderFieldProxyBehaviorEnumWrapperValue = "whole"

	// VaryHeaderFieldProxyBehaviorEnumWrapperValueFiltered captures enum value "filtered"
	VaryHeaderFieldProxyBehaviorEnumWrapperValueFiltered VaryHeaderFieldProxyBehaviorEnumWrapperValue = "filtered"

	// VaryHeaderFieldProxyBehaviorEnumWrapperValueNone captures enum value "none"
	VaryHeaderFieldProxyBehaviorEnumWrapperValueNone VaryHeaderFieldProxyBehaviorEnumWrapperValue = "none"
)

// for schema
var varyHeaderFieldProxyBehaviorEnumWrapperValueEnum []interface{}

func init() {
	var res []VaryHeaderFieldProxyBehaviorEnumWrapperValue
	if err := json.Unmarshal([]byte(`["UNKNOWN","whole","filtered","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		varyHeaderFieldProxyBehaviorEnumWrapperValueEnum = append(varyHeaderFieldProxyBehaviorEnumWrapperValueEnum, v)
	}
}

func (m VaryHeaderFieldProxyBehaviorEnumWrapperValue) validateVaryHeaderFieldProxyBehaviorEnumWrapperValueEnum(path, location string, value VaryHeaderFieldProxyBehaviorEnumWrapperValue) error {
	if err := validate.Enum(path, location, value, varyHeaderFieldProxyBehaviorEnumWrapperValueEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this vary header field proxy behavior enum wrapper value
func (m VaryHeaderFieldProxyBehaviorEnumWrapperValue) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateVaryHeaderFieldProxyBehaviorEnumWrapperValueEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
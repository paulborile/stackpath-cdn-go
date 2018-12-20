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

// DNSOverrideTypeEnumWrapperValue Dns override type enum wrapper value
// swagger:model DnsOverrideTypeEnumWrapperValue
type DNSOverrideTypeEnumWrapperValue string

const (

	// DNSOverrideTypeEnumWrapperValueUNKNOWN captures enum value "UNKNOWN"
	DNSOverrideTypeEnumWrapperValueUNKNOWN DNSOverrideTypeEnumWrapperValue = "UNKNOWN"

	// DNSOverrideTypeEnumWrapperValueA captures enum value "A"
	DNSOverrideTypeEnumWrapperValueA DNSOverrideTypeEnumWrapperValue = "A"

	// DNSOverrideTypeEnumWrapperValueAAAA captures enum value "AAAA"
	DNSOverrideTypeEnumWrapperValueAAAA DNSOverrideTypeEnumWrapperValue = "AAAA"

	// DNSOverrideTypeEnumWrapperValueCNAME captures enum value "CNAME"
	DNSOverrideTypeEnumWrapperValueCNAME DNSOverrideTypeEnumWrapperValue = "CNAME"
)

// for schema
var dnsOverrideTypeEnumWrapperValueEnum []interface{}

func init() {
	var res []DNSOverrideTypeEnumWrapperValue
	if err := json.Unmarshal([]byte(`["UNKNOWN","A","AAAA","CNAME"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dnsOverrideTypeEnumWrapperValueEnum = append(dnsOverrideTypeEnumWrapperValueEnum, v)
	}
}

func (m DNSOverrideTypeEnumWrapperValue) validateDNSOverrideTypeEnumWrapperValueEnum(path, location string, value DNSOverrideTypeEnumWrapperValue) error {
	if err := validate.Enum(path, location, value, dnsOverrideTypeEnumWrapperValueEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this Dns override type enum wrapper value
func (m DNSOverrideTypeEnumWrapperValue) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDNSOverrideTypeEnumWrapperValueEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
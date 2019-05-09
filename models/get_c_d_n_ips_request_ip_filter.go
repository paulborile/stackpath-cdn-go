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

// GetCDNIpsRequestIPFilter The kinds of IP addresses that can be retrieved
//
// - ALL: Retrieve all IP address types
//  - IPV4: Retrieve only IPv4 addresses
//  - IPV6: Retrieve only IPv6 addresses
// swagger:model GetCDNIPsRequestIPFilter
type GetCDNIpsRequestIPFilter string

const (

	// GetCDNIpsRequestIPFilterALL captures enum value "ALL"
	GetCDNIpsRequestIPFilterALL GetCDNIpsRequestIPFilter = "ALL"

	// GetCDNIpsRequestIPFilterIPV4 captures enum value "IPV4"
	GetCDNIpsRequestIPFilterIPV4 GetCDNIpsRequestIPFilter = "IPV4"

	// GetCDNIpsRequestIPFilterIPV6 captures enum value "IPV6"
	GetCDNIpsRequestIPFilterIPV6 GetCDNIpsRequestIPFilter = "IPV6"
)

// for schema
var getCDNIpsRequestIpFilterEnum []interface{}

func init() {
	var res []GetCDNIpsRequestIPFilter
	if err := json.Unmarshal([]byte(`["ALL","IPV4","IPV6"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCDNIpsRequestIpFilterEnum = append(getCDNIpsRequestIpFilterEnum, v)
	}
}

func (m GetCDNIpsRequestIPFilter) validateGetCDNIpsRequestIPFilterEnum(path, location string, value GetCDNIpsRequestIPFilter) error {
	if err := validate.Enum(path, location, value, getCDNIpsRequestIpFilterEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this get c d n ips request IP filter
func (m GetCDNIpsRequestIPFilter) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateGetCDNIpsRequestIPFilterEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
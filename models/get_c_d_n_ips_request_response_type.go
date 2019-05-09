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

// GetCDNIpsRequestResponseType The response formats that can be retrieved
//
// - JSON: Return a JSON serialized response
//  - PLAIN_TEXT: Return a plain text response with one IP address on each line
// swagger:model GetCDNIPsRequestResponseType
type GetCDNIpsRequestResponseType string

const (

	// GetCDNIpsRequestResponseTypeJSON captures enum value "JSON"
	GetCDNIpsRequestResponseTypeJSON GetCDNIpsRequestResponseType = "JSON"

	// GetCDNIpsRequestResponseTypePLAINTEXT captures enum value "PLAIN_TEXT"
	GetCDNIpsRequestResponseTypePLAINTEXT GetCDNIpsRequestResponseType = "PLAIN_TEXT"
)

// for schema
var getCDNIpsRequestResponseTypeEnum []interface{}

func init() {
	var res []GetCDNIpsRequestResponseType
	if err := json.Unmarshal([]byte(`["JSON","PLAIN_TEXT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCDNIpsRequestResponseTypeEnum = append(getCDNIpsRequestResponseTypeEnum, v)
	}
}

func (m GetCDNIpsRequestResponseType) validateGetCDNIpsRequestResponseTypeEnum(path, location string, value GetCDNIpsRequestResponseType) error {
	if err := validate.Enum(path, location, value, getCDNIpsRequestResponseTypeEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this get c d n ips request response type
func (m GetCDNIpsRequestResponseType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateGetCDNIpsRequestResponseTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
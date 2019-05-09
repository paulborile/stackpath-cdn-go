// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// StackpathRPCBadRequest stackpath rpc bad request
// swagger:model stackpath.rpc.BadRequest
type StackpathRPCBadRequest struct {

	// field violations
	FieldViolations []*StackpathRPCBadRequestFieldViolation `json:"fieldViolations,omitempty"`
}

// AtType gets the at type of this subtype
func (m *StackpathRPCBadRequest) AtType() string {
	return "stackpath.rpc.BadRequest"
}

// SetAtType sets the at type of this subtype
func (m *StackpathRPCBadRequest) SetAtType(val string) {

}

// FieldViolations gets the field violations of this subtype

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *StackpathRPCBadRequest) UnmarshalJSON(raw []byte) error {
	var data struct {

		// field violations
		FieldViolations []*StackpathRPCBadRequestFieldViolation `json:"fieldViolations,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var base struct {
		/* Just the base type fields. Used for unmashalling polymorphic types.*/

		AtType string `json:"@type"`
	}
	buf = bytes.NewBuffer(raw)
	dec = json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&base); err != nil {
		return err
	}

	var result StackpathRPCBadRequest

	if base.AtType != result.AtType() {
		/* Not the type we're looking for. */
		return errors.New(422, "invalid @type value: %q", base.AtType)
	}

	result.FieldViolations = data.FieldViolations

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m StackpathRPCBadRequest) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {

		// field violations
		FieldViolations []*StackpathRPCBadRequestFieldViolation `json:"fieldViolations,omitempty"`
	}{

		FieldViolations: m.FieldViolations,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		AtType string `json:"@type"`
	}{

		AtType: m.AtType(),
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this stackpath rpc bad request
func (m *StackpathRPCBadRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFieldViolations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StackpathRPCBadRequest) validateFieldViolations(formats strfmt.Registry) error {

	if swag.IsZero(m.FieldViolations) { // not required
		return nil
	}

	for i := 0; i < len(m.FieldViolations); i++ {
		if swag.IsZero(m.FieldViolations[i]) { // not required
			continue
		}

		if m.FieldViolations[i] != nil {
			if err := m.FieldViolations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fieldViolations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *StackpathRPCBadRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StackpathRPCBadRequest) UnmarshalBinary(b []byte) error {
	var res StackpathRPCBadRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

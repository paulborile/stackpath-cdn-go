// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CustconfContentDispositionByURL The content disposition by URL policy allows you to control
// Content-Disposition HTTP header delivered by the CDN caching servers.  The
// policy gives you control over each of the header directives and allows you to
// specify a URL pattern match for determining when to apply the policy. Please
// note this policy takes precedence over all the other content disposition
// policies.
// swagger:model custconfContentDispositionByURL
type CustconfContentDispositionByURL struct {

	// This is the name of the query string parameter which contains the file name
	// to use in the Content-Disposition header.  This setting is typically used
	// by customers to configure a "friendly name" for URLs that have obfuscated
	// file names.
	DispositionNameQSParam string `json:"dispositionNameQSParam,omitempty"`

	// This setting allows you to define a query string parameter that when
	// present in the URL contains a string that should be used for the
	// Content-Disposition header.  The string specified in the URL will
	// completely replace the value the CDN would have used based on other
	// policies defined for the Content-Disposition header.
	DispositionOverrideQSParam string `json:"dispositionOverrideQSParam,omitempty"`

	// This is the name of the query string parameter which contains the
	// disposition type to use in the Content-Disposition header. Typically, this
	// value is set to attachment if you want the browser to present the user with
	// a "File Download" dialog box and set to inline if you want the browser to
	// render the content inline (play an audio/video file instead of downloading
	// it).
	DispositionTypeQSParam string `json:"dispositionTypeQSParam,omitempty"`

	// enabled
	Enabled bool `json:"enabled"`

	// This is used by the API to perform conflict checking.
	ID string `json:"id,omitempty"`

	// This setting allows you to force the Content-Disposition generated by the
	// CDN for this policy to override the Content-Disposition header cached from
	// your origin.
	OverrideOriginHeader bool `json:"overrideOriginHeader"`
}

// Validate validates this custconf content disposition by URL
func (m *CustconfContentDispositionByURL) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustconfContentDispositionByURL) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustconfContentDispositionByURL) UnmarshalBinary(b []byte) error {
	var res CustconfContentDispositionByURL
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

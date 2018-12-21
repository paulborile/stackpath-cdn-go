// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CustconfOriginPullCacheExtension The cache extension policy allows you to define cache revalidation exceptions
// on expired content.  This policy is applied by the CDN caching servers when
// they are are unable to revalidate an expired asset with your origin due to
// network connectivity issues or unresponsiveness from your origin.
// swagger:model custconfOriginPullCacheExtension
type CustconfOriginPullCacheExtension struct {

	// enabled
	Enabled bool `json:"enabled"`

	// This is the number of seconds to extend an asset's TTL when the origin is
	// unavailable. The CDN will continue to retry the origin up to the Origin
	// Unavailable Max TTL.
	ExpiredCacheExtension int32 `json:"expiredCacheExtension,omitempty"`

	// This is used by the API to perform conflict checking.
	ID string `json:"id,omitempty"`

	// The origin unavailable max TTL value is used by the caching server when
	// your origin is unresponsive or the CDN cannot establish a connection to
	// your origin.  Under these conditions, the CDN can continue to serve expired
	// assets from the cache. The value specified in this field establishes a
	// maximum allowable TTL for your expired assets.  If your origin connectivity
	// or responsiveness is not corrected within your maximum allowable TTL, the
	// CDN no longer serves your expired assets.
	OriginUnreachableCacheExtension int32 `json:"originUnreachableCacheExtension,omitempty"`
}

// Validate validates this custconf origin pull cache extension
func (m *CustconfOriginPullCacheExtension) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustconfOriginPullCacheExtension) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustconfOriginPullCacheExtension) UnmarshalBinary(b []byte) error {
	var res CustconfOriginPullCacheExtension
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

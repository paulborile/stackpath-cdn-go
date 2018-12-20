// Code generated by go-swagger; DO NOT EDIT.

package cdn

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewEnableSiteParams creates a new EnableSiteParams object
// with the default values initialized.
func NewEnableSiteParams() *EnableSiteParams {
	var ()
	return &EnableSiteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEnableSiteParamsWithTimeout creates a new EnableSiteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEnableSiteParamsWithTimeout(timeout time.Duration) *EnableSiteParams {
	var ()
	return &EnableSiteParams{

		timeout: timeout,
	}
}

// NewEnableSiteParamsWithContext creates a new EnableSiteParams object
// with the default values initialized, and the ability to set a context for a request
func NewEnableSiteParamsWithContext(ctx context.Context) *EnableSiteParams {
	var ()
	return &EnableSiteParams{

		Context: ctx,
	}
}

// NewEnableSiteParamsWithHTTPClient creates a new EnableSiteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEnableSiteParamsWithHTTPClient(client *http.Client) *EnableSiteParams {
	var ()
	return &EnableSiteParams{
		HTTPClient: client,
	}
}

/*EnableSiteParams contains all the parameters to send to the API endpoint
for the enable site operation typically these are written to a http.Request
*/
type EnableSiteParams struct {

	/*SiteID*/
	SiteID string
	/*StackID*/
	StackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the enable site params
func (o *EnableSiteParams) WithTimeout(timeout time.Duration) *EnableSiteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the enable site params
func (o *EnableSiteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the enable site params
func (o *EnableSiteParams) WithContext(ctx context.Context) *EnableSiteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the enable site params
func (o *EnableSiteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the enable site params
func (o *EnableSiteParams) WithHTTPClient(client *http.Client) *EnableSiteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the enable site params
func (o *EnableSiteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSiteID adds the siteID to the enable site params
func (o *EnableSiteParams) WithSiteID(siteID string) *EnableSiteParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the enable site params
func (o *EnableSiteParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WithStackID adds the stackID to the enable site params
func (o *EnableSiteParams) WithStackID(stackID string) *EnableSiteParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the enable site params
func (o *EnableSiteParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WriteToRequest writes these params to a swagger request
func (o *EnableSiteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param site_id
	if err := r.SetPathParam("site_id", o.SiteID); err != nil {
		return err
	}

	// path param stack_id
	if err := r.SetPathParam("stack_id", o.StackID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
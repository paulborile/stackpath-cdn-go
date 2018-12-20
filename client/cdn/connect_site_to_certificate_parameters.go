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

// NewConnectSiteToCertificateParams creates a new ConnectSiteToCertificateParams object
// with the default values initialized.
func NewConnectSiteToCertificateParams() *ConnectSiteToCertificateParams {
	var ()
	return &ConnectSiteToCertificateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewConnectSiteToCertificateParamsWithTimeout creates a new ConnectSiteToCertificateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewConnectSiteToCertificateParamsWithTimeout(timeout time.Duration) *ConnectSiteToCertificateParams {
	var ()
	return &ConnectSiteToCertificateParams{

		timeout: timeout,
	}
}

// NewConnectSiteToCertificateParamsWithContext creates a new ConnectSiteToCertificateParams object
// with the default values initialized, and the ability to set a context for a request
func NewConnectSiteToCertificateParamsWithContext(ctx context.Context) *ConnectSiteToCertificateParams {
	var ()
	return &ConnectSiteToCertificateParams{

		Context: ctx,
	}
}

// NewConnectSiteToCertificateParamsWithHTTPClient creates a new ConnectSiteToCertificateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewConnectSiteToCertificateParamsWithHTTPClient(client *http.Client) *ConnectSiteToCertificateParams {
	var ()
	return &ConnectSiteToCertificateParams{
		HTTPClient: client,
	}
}

/*ConnectSiteToCertificateParams contains all the parameters to send to the API endpoint
for the connect site to certificate operation typically these are written to a http.Request
*/
type ConnectSiteToCertificateParams struct {

	/*CertificateID*/
	CertificateID string
	/*SiteID*/
	SiteID string
	/*StackID*/
	StackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the connect site to certificate params
func (o *ConnectSiteToCertificateParams) WithTimeout(timeout time.Duration) *ConnectSiteToCertificateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the connect site to certificate params
func (o *ConnectSiteToCertificateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the connect site to certificate params
func (o *ConnectSiteToCertificateParams) WithContext(ctx context.Context) *ConnectSiteToCertificateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the connect site to certificate params
func (o *ConnectSiteToCertificateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the connect site to certificate params
func (o *ConnectSiteToCertificateParams) WithHTTPClient(client *http.Client) *ConnectSiteToCertificateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the connect site to certificate params
func (o *ConnectSiteToCertificateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCertificateID adds the certificateID to the connect site to certificate params
func (o *ConnectSiteToCertificateParams) WithCertificateID(certificateID string) *ConnectSiteToCertificateParams {
	o.SetCertificateID(certificateID)
	return o
}

// SetCertificateID adds the certificateId to the connect site to certificate params
func (o *ConnectSiteToCertificateParams) SetCertificateID(certificateID string) {
	o.CertificateID = certificateID
}

// WithSiteID adds the siteID to the connect site to certificate params
func (o *ConnectSiteToCertificateParams) WithSiteID(siteID string) *ConnectSiteToCertificateParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the connect site to certificate params
func (o *ConnectSiteToCertificateParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WithStackID adds the stackID to the connect site to certificate params
func (o *ConnectSiteToCertificateParams) WithStackID(stackID string) *ConnectSiteToCertificateParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the connect site to certificate params
func (o *ConnectSiteToCertificateParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WriteToRequest writes these params to a swagger request
func (o *ConnectSiteToCertificateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param certificate_id
	if err := r.SetPathParam("certificate_id", o.CertificateID); err != nil {
		return err
	}

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
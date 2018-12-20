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

// NewGetScopeOriginsParams creates a new GetScopeOriginsParams object
// with the default values initialized.
func NewGetScopeOriginsParams() *GetScopeOriginsParams {
	var ()
	return &GetScopeOriginsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetScopeOriginsParamsWithTimeout creates a new GetScopeOriginsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetScopeOriginsParamsWithTimeout(timeout time.Duration) *GetScopeOriginsParams {
	var ()
	return &GetScopeOriginsParams{

		timeout: timeout,
	}
}

// NewGetScopeOriginsParamsWithContext creates a new GetScopeOriginsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetScopeOriginsParamsWithContext(ctx context.Context) *GetScopeOriginsParams {
	var ()
	return &GetScopeOriginsParams{

		Context: ctx,
	}
}

// NewGetScopeOriginsParamsWithHTTPClient creates a new GetScopeOriginsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetScopeOriginsParamsWithHTTPClient(client *http.Client) *GetScopeOriginsParams {
	var ()
	return &GetScopeOriginsParams{
		HTTPClient: client,
	}
}

/*GetScopeOriginsParams contains all the parameters to send to the API endpoint
for the get scope origins operation typically these are written to a http.Request
*/
type GetScopeOriginsParams struct {

	/*PageRequestAfter
	  after is the cursor value after which data will be returned.

	*/
	PageRequestAfter *string
	/*PageRequestFilter
	  filter will accept sql style constraints.

	*/
	PageRequestFilter *string
	/*PageRequestFirst
	  first is the number of items desired.

	*/
	PageRequestFirst *string
	/*PageRequestSortBy
	  sort_by will sort the response by the given field.

	*/
	PageRequestSortBy *string
	/*ScopeID*/
	ScopeID string
	/*SiteID*/
	SiteID string
	/*StackID*/
	StackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get scope origins params
func (o *GetScopeOriginsParams) WithTimeout(timeout time.Duration) *GetScopeOriginsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get scope origins params
func (o *GetScopeOriginsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get scope origins params
func (o *GetScopeOriginsParams) WithContext(ctx context.Context) *GetScopeOriginsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get scope origins params
func (o *GetScopeOriginsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get scope origins params
func (o *GetScopeOriginsParams) WithHTTPClient(client *http.Client) *GetScopeOriginsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get scope origins params
func (o *GetScopeOriginsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageRequestAfter adds the pageRequestAfter to the get scope origins params
func (o *GetScopeOriginsParams) WithPageRequestAfter(pageRequestAfter *string) *GetScopeOriginsParams {
	o.SetPageRequestAfter(pageRequestAfter)
	return o
}

// SetPageRequestAfter adds the pageRequestAfter to the get scope origins params
func (o *GetScopeOriginsParams) SetPageRequestAfter(pageRequestAfter *string) {
	o.PageRequestAfter = pageRequestAfter
}

// WithPageRequestFilter adds the pageRequestFilter to the get scope origins params
func (o *GetScopeOriginsParams) WithPageRequestFilter(pageRequestFilter *string) *GetScopeOriginsParams {
	o.SetPageRequestFilter(pageRequestFilter)
	return o
}

// SetPageRequestFilter adds the pageRequestFilter to the get scope origins params
func (o *GetScopeOriginsParams) SetPageRequestFilter(pageRequestFilter *string) {
	o.PageRequestFilter = pageRequestFilter
}

// WithPageRequestFirst adds the pageRequestFirst to the get scope origins params
func (o *GetScopeOriginsParams) WithPageRequestFirst(pageRequestFirst *string) *GetScopeOriginsParams {
	o.SetPageRequestFirst(pageRequestFirst)
	return o
}

// SetPageRequestFirst adds the pageRequestFirst to the get scope origins params
func (o *GetScopeOriginsParams) SetPageRequestFirst(pageRequestFirst *string) {
	o.PageRequestFirst = pageRequestFirst
}

// WithPageRequestSortBy adds the pageRequestSortBy to the get scope origins params
func (o *GetScopeOriginsParams) WithPageRequestSortBy(pageRequestSortBy *string) *GetScopeOriginsParams {
	o.SetPageRequestSortBy(pageRequestSortBy)
	return o
}

// SetPageRequestSortBy adds the pageRequestSortBy to the get scope origins params
func (o *GetScopeOriginsParams) SetPageRequestSortBy(pageRequestSortBy *string) {
	o.PageRequestSortBy = pageRequestSortBy
}

// WithScopeID adds the scopeID to the get scope origins params
func (o *GetScopeOriginsParams) WithScopeID(scopeID string) *GetScopeOriginsParams {
	o.SetScopeID(scopeID)
	return o
}

// SetScopeID adds the scopeId to the get scope origins params
func (o *GetScopeOriginsParams) SetScopeID(scopeID string) {
	o.ScopeID = scopeID
}

// WithSiteID adds the siteID to the get scope origins params
func (o *GetScopeOriginsParams) WithSiteID(siteID string) *GetScopeOriginsParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the get scope origins params
func (o *GetScopeOriginsParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WithStackID adds the stackID to the get scope origins params
func (o *GetScopeOriginsParams) WithStackID(stackID string) *GetScopeOriginsParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the get scope origins params
func (o *GetScopeOriginsParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WriteToRequest writes these params to a swagger request
func (o *GetScopeOriginsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PageRequestAfter != nil {

		// query param page_request.after
		var qrPageRequestAfter string
		if o.PageRequestAfter != nil {
			qrPageRequestAfter = *o.PageRequestAfter
		}
		qPageRequestAfter := qrPageRequestAfter
		if qPageRequestAfter != "" {
			if err := r.SetQueryParam("page_request.after", qPageRequestAfter); err != nil {
				return err
			}
		}

	}

	if o.PageRequestFilter != nil {

		// query param page_request.filter
		var qrPageRequestFilter string
		if o.PageRequestFilter != nil {
			qrPageRequestFilter = *o.PageRequestFilter
		}
		qPageRequestFilter := qrPageRequestFilter
		if qPageRequestFilter != "" {
			if err := r.SetQueryParam("page_request.filter", qPageRequestFilter); err != nil {
				return err
			}
		}

	}

	if o.PageRequestFirst != nil {

		// query param page_request.first
		var qrPageRequestFirst string
		if o.PageRequestFirst != nil {
			qrPageRequestFirst = *o.PageRequestFirst
		}
		qPageRequestFirst := qrPageRequestFirst
		if qPageRequestFirst != "" {
			if err := r.SetQueryParam("page_request.first", qPageRequestFirst); err != nil {
				return err
			}
		}

	}

	if o.PageRequestSortBy != nil {

		// query param page_request.sort_by
		var qrPageRequestSortBy string
		if o.PageRequestSortBy != nil {
			qrPageRequestSortBy = *o.PageRequestSortBy
		}
		qPageRequestSortBy := qrPageRequestSortBy
		if qPageRequestSortBy != "" {
			if err := r.SetQueryParam("page_request.sort_by", qPageRequestSortBy); err != nil {
				return err
			}
		}

	}

	// path param scope_id
	if err := r.SetPathParam("scope_id", o.ScopeID); err != nil {
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
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

// NewGetSitesParams creates a new GetSitesParams object
// with the default values initialized.
func NewGetSitesParams() *GetSitesParams {
	var ()
	return &GetSitesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSitesParamsWithTimeout creates a new GetSitesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSitesParamsWithTimeout(timeout time.Duration) *GetSitesParams {
	var ()
	return &GetSitesParams{

		timeout: timeout,
	}
}

// NewGetSitesParamsWithContext creates a new GetSitesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSitesParamsWithContext(ctx context.Context) *GetSitesParams {
	var ()
	return &GetSitesParams{

		Context: ctx,
	}
}

// NewGetSitesParamsWithHTTPClient creates a new GetSitesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSitesParamsWithHTTPClient(client *http.Client) *GetSitesParams {
	var ()
	return &GetSitesParams{
		HTTPClient: client,
	}
}

/*GetSitesParams contains all the parameters to send to the API endpoint
for the get sites operation typically these are written to a http.Request
*/
type GetSitesParams struct {

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
	/*StackID*/
	StackID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get sites params
func (o *GetSitesParams) WithTimeout(timeout time.Duration) *GetSitesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get sites params
func (o *GetSitesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get sites params
func (o *GetSitesParams) WithContext(ctx context.Context) *GetSitesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get sites params
func (o *GetSitesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get sites params
func (o *GetSitesParams) WithHTTPClient(client *http.Client) *GetSitesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get sites params
func (o *GetSitesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageRequestAfter adds the pageRequestAfter to the get sites params
func (o *GetSitesParams) WithPageRequestAfter(pageRequestAfter *string) *GetSitesParams {
	o.SetPageRequestAfter(pageRequestAfter)
	return o
}

// SetPageRequestAfter adds the pageRequestAfter to the get sites params
func (o *GetSitesParams) SetPageRequestAfter(pageRequestAfter *string) {
	o.PageRequestAfter = pageRequestAfter
}

// WithPageRequestFilter adds the pageRequestFilter to the get sites params
func (o *GetSitesParams) WithPageRequestFilter(pageRequestFilter *string) *GetSitesParams {
	o.SetPageRequestFilter(pageRequestFilter)
	return o
}

// SetPageRequestFilter adds the pageRequestFilter to the get sites params
func (o *GetSitesParams) SetPageRequestFilter(pageRequestFilter *string) {
	o.PageRequestFilter = pageRequestFilter
}

// WithPageRequestFirst adds the pageRequestFirst to the get sites params
func (o *GetSitesParams) WithPageRequestFirst(pageRequestFirst *string) *GetSitesParams {
	o.SetPageRequestFirst(pageRequestFirst)
	return o
}

// SetPageRequestFirst adds the pageRequestFirst to the get sites params
func (o *GetSitesParams) SetPageRequestFirst(pageRequestFirst *string) {
	o.PageRequestFirst = pageRequestFirst
}

// WithPageRequestSortBy adds the pageRequestSortBy to the get sites params
func (o *GetSitesParams) WithPageRequestSortBy(pageRequestSortBy *string) *GetSitesParams {
	o.SetPageRequestSortBy(pageRequestSortBy)
	return o
}

// SetPageRequestSortBy adds the pageRequestSortBy to the get sites params
func (o *GetSitesParams) SetPageRequestSortBy(pageRequestSortBy *string) {
	o.PageRequestSortBy = pageRequestSortBy
}

// WithStackID adds the stackID to the get sites params
func (o *GetSitesParams) WithStackID(stackID string) *GetSitesParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the get sites params
func (o *GetSitesParams) SetStackID(stackID string) {
	o.StackID = stackID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSitesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param stack_id
	if err := r.SetPathParam("stack_id", o.StackID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
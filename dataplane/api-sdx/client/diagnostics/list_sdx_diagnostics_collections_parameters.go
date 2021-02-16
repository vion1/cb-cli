// Code generated by go-swagger; DO NOT EDIT.

package diagnostics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListSdxDiagnosticsCollectionsParams creates a new ListSdxDiagnosticsCollectionsParams object
// with the default values initialized.
func NewListSdxDiagnosticsCollectionsParams() *ListSdxDiagnosticsCollectionsParams {
	var ()
	return &ListSdxDiagnosticsCollectionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListSdxDiagnosticsCollectionsParamsWithTimeout creates a new ListSdxDiagnosticsCollectionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListSdxDiagnosticsCollectionsParamsWithTimeout(timeout time.Duration) *ListSdxDiagnosticsCollectionsParams {
	var ()
	return &ListSdxDiagnosticsCollectionsParams{

		timeout: timeout,
	}
}

// NewListSdxDiagnosticsCollectionsParamsWithContext creates a new ListSdxDiagnosticsCollectionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewListSdxDiagnosticsCollectionsParamsWithContext(ctx context.Context) *ListSdxDiagnosticsCollectionsParams {
	var ()
	return &ListSdxDiagnosticsCollectionsParams{

		Context: ctx,
	}
}

// NewListSdxDiagnosticsCollectionsParamsWithHTTPClient creates a new ListSdxDiagnosticsCollectionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListSdxDiagnosticsCollectionsParamsWithHTTPClient(client *http.Client) *ListSdxDiagnosticsCollectionsParams {
	var ()
	return &ListSdxDiagnosticsCollectionsParams{
		HTTPClient: client,
	}
}

/*ListSdxDiagnosticsCollectionsParams contains all the parameters to send to the API endpoint
for the list sdx diagnostics collections operation typically these are written to a http.Request
*/
type ListSdxDiagnosticsCollectionsParams struct {

	/*Crn*/
	Crn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list sdx diagnostics collections params
func (o *ListSdxDiagnosticsCollectionsParams) WithTimeout(timeout time.Duration) *ListSdxDiagnosticsCollectionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list sdx diagnostics collections params
func (o *ListSdxDiagnosticsCollectionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list sdx diagnostics collections params
func (o *ListSdxDiagnosticsCollectionsParams) WithContext(ctx context.Context) *ListSdxDiagnosticsCollectionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list sdx diagnostics collections params
func (o *ListSdxDiagnosticsCollectionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list sdx diagnostics collections params
func (o *ListSdxDiagnosticsCollectionsParams) WithHTTPClient(client *http.Client) *ListSdxDiagnosticsCollectionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list sdx diagnostics collections params
func (o *ListSdxDiagnosticsCollectionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCrn adds the crn to the list sdx diagnostics collections params
func (o *ListSdxDiagnosticsCollectionsParams) WithCrn(crn string) *ListSdxDiagnosticsCollectionsParams {
	o.SetCrn(crn)
	return o
}

// SetCrn adds the crn to the list sdx diagnostics collections params
func (o *ListSdxDiagnosticsCollectionsParams) SetCrn(crn string) {
	o.Crn = crn
}

// WriteToRequest writes these params to a swagger request
func (o *ListSdxDiagnosticsCollectionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param crn
	if err := r.SetPathParam("crn", o.Crn); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

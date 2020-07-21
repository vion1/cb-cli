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

// NewGetSdxCmVMLogsParams creates a new GetSdxCmVMLogsParams object
// with the default values initialized.
func NewGetSdxCmVMLogsParams() *GetSdxCmVMLogsParams {

	return &GetSdxCmVMLogsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSdxCmVMLogsParamsWithTimeout creates a new GetSdxCmVMLogsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSdxCmVMLogsParamsWithTimeout(timeout time.Duration) *GetSdxCmVMLogsParams {

	return &GetSdxCmVMLogsParams{

		timeout: timeout,
	}
}

// NewGetSdxCmVMLogsParamsWithContext creates a new GetSdxCmVMLogsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSdxCmVMLogsParamsWithContext(ctx context.Context) *GetSdxCmVMLogsParams {

	return &GetSdxCmVMLogsParams{

		Context: ctx,
	}
}

// NewGetSdxCmVMLogsParamsWithHTTPClient creates a new GetSdxCmVMLogsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSdxCmVMLogsParamsWithHTTPClient(client *http.Client) *GetSdxCmVMLogsParams {

	return &GetSdxCmVMLogsParams{
		HTTPClient: client,
	}
}

/*GetSdxCmVMLogsParams contains all the parameters to send to the API endpoint
for the get sdx cm Vm logs operation typically these are written to a http.Request
*/
type GetSdxCmVMLogsParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get sdx cm Vm logs params
func (o *GetSdxCmVMLogsParams) WithTimeout(timeout time.Duration) *GetSdxCmVMLogsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get sdx cm Vm logs params
func (o *GetSdxCmVMLogsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get sdx cm Vm logs params
func (o *GetSdxCmVMLogsParams) WithContext(ctx context.Context) *GetSdxCmVMLogsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get sdx cm Vm logs params
func (o *GetSdxCmVMLogsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get sdx cm Vm logs params
func (o *GetSdxCmVMLogsParams) WithHTTPClient(client *http.Client) *GetSdxCmVMLogsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get sdx cm Vm logs params
func (o *GetSdxCmVMLogsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetSdxCmVMLogsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

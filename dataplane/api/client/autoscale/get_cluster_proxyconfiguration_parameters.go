// Code generated by go-swagger; DO NOT EDIT.

package autoscale

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

// NewGetClusterProxyconfigurationParams creates a new GetClusterProxyconfigurationParams object
// with the default values initialized.
func NewGetClusterProxyconfigurationParams() *GetClusterProxyconfigurationParams {

	return &GetClusterProxyconfigurationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetClusterProxyconfigurationParamsWithTimeout creates a new GetClusterProxyconfigurationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetClusterProxyconfigurationParamsWithTimeout(timeout time.Duration) *GetClusterProxyconfigurationParams {

	return &GetClusterProxyconfigurationParams{

		timeout: timeout,
	}
}

// NewGetClusterProxyconfigurationParamsWithContext creates a new GetClusterProxyconfigurationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetClusterProxyconfigurationParamsWithContext(ctx context.Context) *GetClusterProxyconfigurationParams {

	return &GetClusterProxyconfigurationParams{

		Context: ctx,
	}
}

// NewGetClusterProxyconfigurationParamsWithHTTPClient creates a new GetClusterProxyconfigurationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetClusterProxyconfigurationParamsWithHTTPClient(client *http.Client) *GetClusterProxyconfigurationParams {

	return &GetClusterProxyconfigurationParams{
		HTTPClient: client,
	}
}

/*GetClusterProxyconfigurationParams contains all the parameters to send to the API endpoint
for the get cluster proxyconfiguration operation typically these are written to a http.Request
*/
type GetClusterProxyconfigurationParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get cluster proxyconfiguration params
func (o *GetClusterProxyconfigurationParams) WithTimeout(timeout time.Duration) *GetClusterProxyconfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cluster proxyconfiguration params
func (o *GetClusterProxyconfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cluster proxyconfiguration params
func (o *GetClusterProxyconfigurationParams) WithContext(ctx context.Context) *GetClusterProxyconfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cluster proxyconfiguration params
func (o *GetClusterProxyconfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cluster proxyconfiguration params
func (o *GetClusterProxyconfigurationParams) WithHTTPClient(client *http.Client) *GetClusterProxyconfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cluster proxyconfiguration params
func (o *GetClusterProxyconfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetClusterProxyconfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

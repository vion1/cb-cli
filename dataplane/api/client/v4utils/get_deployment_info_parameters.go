// Code generated by go-swagger; DO NOT EDIT.

package v4utils

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

// NewGetDeploymentInfoParams creates a new GetDeploymentInfoParams object
// with the default values initialized.
func NewGetDeploymentInfoParams() *GetDeploymentInfoParams {

	return &GetDeploymentInfoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeploymentInfoParamsWithTimeout creates a new GetDeploymentInfoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDeploymentInfoParamsWithTimeout(timeout time.Duration) *GetDeploymentInfoParams {

	return &GetDeploymentInfoParams{

		timeout: timeout,
	}
}

// NewGetDeploymentInfoParamsWithContext creates a new GetDeploymentInfoParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDeploymentInfoParamsWithContext(ctx context.Context) *GetDeploymentInfoParams {

	return &GetDeploymentInfoParams{

		Context: ctx,
	}
}

// NewGetDeploymentInfoParamsWithHTTPClient creates a new GetDeploymentInfoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDeploymentInfoParamsWithHTTPClient(client *http.Client) *GetDeploymentInfoParams {

	return &GetDeploymentInfoParams{
		HTTPClient: client,
	}
}

/*GetDeploymentInfoParams contains all the parameters to send to the API endpoint
for the get deployment info operation typically these are written to a http.Request
*/
type GetDeploymentInfoParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get deployment info params
func (o *GetDeploymentInfoParams) WithTimeout(timeout time.Duration) *GetDeploymentInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get deployment info params
func (o *GetDeploymentInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get deployment info params
func (o *GetDeploymentInfoParams) WithContext(ctx context.Context) *GetDeploymentInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get deployment info params
func (o *GetDeploymentInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get deployment info params
func (o *GetDeploymentInfoParams) WithHTTPClient(client *http.Client) *GetDeploymentInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get deployment info params
func (o *GetDeploymentInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeploymentInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package v1env

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

	model "github.com/hortonworks/cb-cli/dataplane/api-environment/model"
)

// NewGetCreateEnvironmentForCliParams creates a new GetCreateEnvironmentForCliParams object
// with the default values initialized.
func NewGetCreateEnvironmentForCliParams() *GetCreateEnvironmentForCliParams {
	var ()
	return &GetCreateEnvironmentForCliParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCreateEnvironmentForCliParamsWithTimeout creates a new GetCreateEnvironmentForCliParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCreateEnvironmentForCliParamsWithTimeout(timeout time.Duration) *GetCreateEnvironmentForCliParams {
	var ()
	return &GetCreateEnvironmentForCliParams{

		timeout: timeout,
	}
}

// NewGetCreateEnvironmentForCliParamsWithContext creates a new GetCreateEnvironmentForCliParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCreateEnvironmentForCliParamsWithContext(ctx context.Context) *GetCreateEnvironmentForCliParams {
	var ()
	return &GetCreateEnvironmentForCliParams{

		Context: ctx,
	}
}

// NewGetCreateEnvironmentForCliParamsWithHTTPClient creates a new GetCreateEnvironmentForCliParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCreateEnvironmentForCliParamsWithHTTPClient(client *http.Client) *GetCreateEnvironmentForCliParams {
	var ()
	return &GetCreateEnvironmentForCliParams{
		HTTPClient: client,
	}
}

/*GetCreateEnvironmentForCliParams contains all the parameters to send to the API endpoint
for the get create environment for cli operation typically these are written to a http.Request
*/
type GetCreateEnvironmentForCliParams struct {

	/*Body*/
	Body *model.EnvironmentV1Request

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get create environment for cli params
func (o *GetCreateEnvironmentForCliParams) WithTimeout(timeout time.Duration) *GetCreateEnvironmentForCliParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get create environment for cli params
func (o *GetCreateEnvironmentForCliParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get create environment for cli params
func (o *GetCreateEnvironmentForCliParams) WithContext(ctx context.Context) *GetCreateEnvironmentForCliParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get create environment for cli params
func (o *GetCreateEnvironmentForCliParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get create environment for cli params
func (o *GetCreateEnvironmentForCliParams) WithHTTPClient(client *http.Client) *GetCreateEnvironmentForCliParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get create environment for cli params
func (o *GetCreateEnvironmentForCliParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get create environment for cli params
func (o *GetCreateEnvironmentForCliParams) WithBody(body *model.EnvironmentV1Request) *GetCreateEnvironmentForCliParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get create environment for cli params
func (o *GetCreateEnvironmentForCliParams) SetBody(body *model.EnvironmentV1Request) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetCreateEnvironmentForCliParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package v1env

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

	model "github.com/hortonworks/cb-cli/dataplane/api-environment/model"
)

// NewEditEnvironmentV1ByCrnParams creates a new EditEnvironmentV1ByCrnParams object
// with the default values initialized.
func NewEditEnvironmentV1ByCrnParams() *EditEnvironmentV1ByCrnParams {
	var ()
	return &EditEnvironmentV1ByCrnParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewEditEnvironmentV1ByCrnParamsWithTimeout creates a new EditEnvironmentV1ByCrnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewEditEnvironmentV1ByCrnParamsWithTimeout(timeout time.Duration) *EditEnvironmentV1ByCrnParams {
	var ()
	return &EditEnvironmentV1ByCrnParams{

		timeout: timeout,
	}
}

// NewEditEnvironmentV1ByCrnParamsWithContext creates a new EditEnvironmentV1ByCrnParams object
// with the default values initialized, and the ability to set a context for a request
func NewEditEnvironmentV1ByCrnParamsWithContext(ctx context.Context) *EditEnvironmentV1ByCrnParams {
	var ()
	return &EditEnvironmentV1ByCrnParams{

		Context: ctx,
	}
}

// NewEditEnvironmentV1ByCrnParamsWithHTTPClient creates a new EditEnvironmentV1ByCrnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewEditEnvironmentV1ByCrnParamsWithHTTPClient(client *http.Client) *EditEnvironmentV1ByCrnParams {
	var ()
	return &EditEnvironmentV1ByCrnParams{
		HTTPClient: client,
	}
}

/*EditEnvironmentV1ByCrnParams contains all the parameters to send to the API endpoint
for the edit environment v1 by crn operation typically these are written to a http.Request
*/
type EditEnvironmentV1ByCrnParams struct {

	/*Body*/
	Body *model.EnvironmentEditV1Request
	/*Crn*/
	Crn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the edit environment v1 by crn params
func (o *EditEnvironmentV1ByCrnParams) WithTimeout(timeout time.Duration) *EditEnvironmentV1ByCrnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edit environment v1 by crn params
func (o *EditEnvironmentV1ByCrnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edit environment v1 by crn params
func (o *EditEnvironmentV1ByCrnParams) WithContext(ctx context.Context) *EditEnvironmentV1ByCrnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edit environment v1 by crn params
func (o *EditEnvironmentV1ByCrnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edit environment v1 by crn params
func (o *EditEnvironmentV1ByCrnParams) WithHTTPClient(client *http.Client) *EditEnvironmentV1ByCrnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edit environment v1 by crn params
func (o *EditEnvironmentV1ByCrnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the edit environment v1 by crn params
func (o *EditEnvironmentV1ByCrnParams) WithBody(body *model.EnvironmentEditV1Request) *EditEnvironmentV1ByCrnParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the edit environment v1 by crn params
func (o *EditEnvironmentV1ByCrnParams) SetBody(body *model.EnvironmentEditV1Request) {
	o.Body = body
}

// WithCrn adds the crn to the edit environment v1 by crn params
func (o *EditEnvironmentV1ByCrnParams) WithCrn(crn string) *EditEnvironmentV1ByCrnParams {
	o.SetCrn(crn)
	return o
}

// SetCrn adds the crn to the edit environment v1 by crn params
func (o *EditEnvironmentV1ByCrnParams) SetCrn(crn string) {
	o.Crn = crn
}

// WriteToRequest writes these params to a swagger request
func (o *EditEnvironmentV1ByCrnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param crn
	if err := r.SetPathParam("crn", o.Crn); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

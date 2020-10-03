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
)

// NewGetCreateEnvironmentForCliByCrnParams creates a new GetCreateEnvironmentForCliByCrnParams object
// with the default values initialized.
func NewGetCreateEnvironmentForCliByCrnParams() *GetCreateEnvironmentForCliByCrnParams {
	var ()
	return &GetCreateEnvironmentForCliByCrnParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCreateEnvironmentForCliByCrnParamsWithTimeout creates a new GetCreateEnvironmentForCliByCrnParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCreateEnvironmentForCliByCrnParamsWithTimeout(timeout time.Duration) *GetCreateEnvironmentForCliByCrnParams {
	var ()
	return &GetCreateEnvironmentForCliByCrnParams{

		timeout: timeout,
	}
}

// NewGetCreateEnvironmentForCliByCrnParamsWithContext creates a new GetCreateEnvironmentForCliByCrnParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCreateEnvironmentForCliByCrnParamsWithContext(ctx context.Context) *GetCreateEnvironmentForCliByCrnParams {
	var ()
	return &GetCreateEnvironmentForCliByCrnParams{

		Context: ctx,
	}
}

// NewGetCreateEnvironmentForCliByCrnParamsWithHTTPClient creates a new GetCreateEnvironmentForCliByCrnParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCreateEnvironmentForCliByCrnParamsWithHTTPClient(client *http.Client) *GetCreateEnvironmentForCliByCrnParams {
	var ()
	return &GetCreateEnvironmentForCliByCrnParams{
		HTTPClient: client,
	}
}

/*GetCreateEnvironmentForCliByCrnParams contains all the parameters to send to the API endpoint
for the get create environment for cli by crn operation typically these are written to a http.Request
*/
type GetCreateEnvironmentForCliByCrnParams struct {

	/*Crn*/
	Crn string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get create environment for cli by crn params
func (o *GetCreateEnvironmentForCliByCrnParams) WithTimeout(timeout time.Duration) *GetCreateEnvironmentForCliByCrnParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get create environment for cli by crn params
func (o *GetCreateEnvironmentForCliByCrnParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get create environment for cli by crn params
func (o *GetCreateEnvironmentForCliByCrnParams) WithContext(ctx context.Context) *GetCreateEnvironmentForCliByCrnParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get create environment for cli by crn params
func (o *GetCreateEnvironmentForCliByCrnParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get create environment for cli by crn params
func (o *GetCreateEnvironmentForCliByCrnParams) WithHTTPClient(client *http.Client) *GetCreateEnvironmentForCliByCrnParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get create environment for cli by crn params
func (o *GetCreateEnvironmentForCliByCrnParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCrn adds the crn to the get create environment for cli by crn params
func (o *GetCreateEnvironmentForCliByCrnParams) WithCrn(crn string) *GetCreateEnvironmentForCliByCrnParams {
	o.SetCrn(crn)
	return o
}

// SetCrn adds the crn to the get create environment for cli by crn params
func (o *GetCreateEnvironmentForCliByCrnParams) SetCrn(crn string) {
	o.Crn = crn
}

// WriteToRequest writes these params to a swagger request
func (o *GetCreateEnvironmentForCliByCrnParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

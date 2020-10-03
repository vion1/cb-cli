// Code generated by go-swagger; DO NOT EDIT.

package v1freeipatest

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewUserShowV1Params creates a new UserShowV1Params object
// with the default values initialized.
func NewUserShowV1Params() *UserShowV1Params {
	var ()
	return &UserShowV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserShowV1ParamsWithTimeout creates a new UserShowV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserShowV1ParamsWithTimeout(timeout time.Duration) *UserShowV1Params {
	var ()
	return &UserShowV1Params{

		timeout: timeout,
	}
}

// NewUserShowV1ParamsWithContext creates a new UserShowV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewUserShowV1ParamsWithContext(ctx context.Context) *UserShowV1Params {
	var ()
	return &UserShowV1Params{

		Context: ctx,
	}
}

// NewUserShowV1ParamsWithHTTPClient creates a new UserShowV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserShowV1ParamsWithHTTPClient(client *http.Client) *UserShowV1Params {
	var ()
	return &UserShowV1Params{
		HTTPClient: client,
	}
}

/*UserShowV1Params contains all the parameters to send to the API endpoint
for the user show v1 operation typically these are written to a http.Request
*/
type UserShowV1Params struct {

	/*ID*/
	ID int64
	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user show v1 params
func (o *UserShowV1Params) WithTimeout(timeout time.Duration) *UserShowV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user show v1 params
func (o *UserShowV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user show v1 params
func (o *UserShowV1Params) WithContext(ctx context.Context) *UserShowV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user show v1 params
func (o *UserShowV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user show v1 params
func (o *UserShowV1Params) WithHTTPClient(client *http.Client) *UserShowV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user show v1 params
func (o *UserShowV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the user show v1 params
func (o *UserShowV1Params) WithID(id int64) *UserShowV1Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the user show v1 params
func (o *UserShowV1Params) SetID(id int64) {
	o.ID = id
}

// WithName adds the name to the user show v1 params
func (o *UserShowV1Params) WithName(name string) *UserShowV1Params {
	o.SetName(name)
	return o
}

// SetName adds the name to the user show v1 params
func (o *UserShowV1Params) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *UserShowV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

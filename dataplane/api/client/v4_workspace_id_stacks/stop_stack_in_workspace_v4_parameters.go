// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_stacks

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

// NewStopStackInWorkspaceV4Params creates a new StopStackInWorkspaceV4Params object
// with the default values initialized.
func NewStopStackInWorkspaceV4Params() *StopStackInWorkspaceV4Params {
	var ()
	return &StopStackInWorkspaceV4Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewStopStackInWorkspaceV4ParamsWithTimeout creates a new StopStackInWorkspaceV4Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewStopStackInWorkspaceV4ParamsWithTimeout(timeout time.Duration) *StopStackInWorkspaceV4Params {
	var ()
	return &StopStackInWorkspaceV4Params{

		timeout: timeout,
	}
}

// NewStopStackInWorkspaceV4ParamsWithContext creates a new StopStackInWorkspaceV4Params object
// with the default values initialized, and the ability to set a context for a request
func NewStopStackInWorkspaceV4ParamsWithContext(ctx context.Context) *StopStackInWorkspaceV4Params {
	var ()
	return &StopStackInWorkspaceV4Params{

		Context: ctx,
	}
}

// NewStopStackInWorkspaceV4ParamsWithHTTPClient creates a new StopStackInWorkspaceV4Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStopStackInWorkspaceV4ParamsWithHTTPClient(client *http.Client) *StopStackInWorkspaceV4Params {
	var ()
	return &StopStackInWorkspaceV4Params{
		HTTPClient: client,
	}
}

/*StopStackInWorkspaceV4Params contains all the parameters to send to the API endpoint
for the stop stack in workspace v4 operation typically these are written to a http.Request
*/
type StopStackInWorkspaceV4Params struct {

	/*Name*/
	Name string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the stop stack in workspace v4 params
func (o *StopStackInWorkspaceV4Params) WithTimeout(timeout time.Duration) *StopStackInWorkspaceV4Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stop stack in workspace v4 params
func (o *StopStackInWorkspaceV4Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stop stack in workspace v4 params
func (o *StopStackInWorkspaceV4Params) WithContext(ctx context.Context) *StopStackInWorkspaceV4Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stop stack in workspace v4 params
func (o *StopStackInWorkspaceV4Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stop stack in workspace v4 params
func (o *StopStackInWorkspaceV4Params) WithHTTPClient(client *http.Client) *StopStackInWorkspaceV4Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stop stack in workspace v4 params
func (o *StopStackInWorkspaceV4Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the stop stack in workspace v4 params
func (o *StopStackInWorkspaceV4Params) WithName(name string) *StopStackInWorkspaceV4Params {
	o.SetName(name)
	return o
}

// SetName adds the name to the stop stack in workspace v4 params
func (o *StopStackInWorkspaceV4Params) SetName(name string) {
	o.Name = name
}

// WithWorkspaceID adds the workspaceID to the stop stack in workspace v4 params
func (o *StopStackInWorkspaceV4Params) WithWorkspaceID(workspaceID int64) *StopStackInWorkspaceV4Params {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the stop stack in workspace v4 params
func (o *StopStackInWorkspaceV4Params) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *StopStackInWorkspaceV4Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param workspaceId
	if err := r.SetPathParam("workspaceId", swag.FormatInt64(o.WorkspaceID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

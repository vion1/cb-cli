// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_blueprints_util

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

// NewGetServiceListParams creates a new GetServiceListParams object
// with the default values initialized.
func NewGetServiceListParams() *GetServiceListParams {
	var ()
	return &GetServiceListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetServiceListParamsWithTimeout creates a new GetServiceListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetServiceListParamsWithTimeout(timeout time.Duration) *GetServiceListParams {
	var ()
	return &GetServiceListParams{

		timeout: timeout,
	}
}

// NewGetServiceListParamsWithContext creates a new GetServiceListParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetServiceListParamsWithContext(ctx context.Context) *GetServiceListParams {
	var ()
	return &GetServiceListParams{

		Context: ctx,
	}
}

// NewGetServiceListParamsWithHTTPClient creates a new GetServiceListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetServiceListParamsWithHTTPClient(client *http.Client) *GetServiceListParams {
	var ()
	return &GetServiceListParams{
		HTTPClient: client,
	}
}

/*GetServiceListParams contains all the parameters to send to the API endpoint
for the get service list operation typically these are written to a http.Request
*/
type GetServiceListParams struct {

	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get service list params
func (o *GetServiceListParams) WithTimeout(timeout time.Duration) *GetServiceListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get service list params
func (o *GetServiceListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get service list params
func (o *GetServiceListParams) WithContext(ctx context.Context) *GetServiceListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get service list params
func (o *GetServiceListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get service list params
func (o *GetServiceListParams) WithHTTPClient(client *http.Client) *GetServiceListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get service list params
func (o *GetServiceListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithWorkspaceID adds the workspaceID to the get service list params
func (o *GetServiceListParams) WithWorkspaceID(workspaceID int64) *GetServiceListParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the get service list params
func (o *GetServiceListParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetServiceListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param workspaceId
	if err := r.SetPathParam("workspaceId", swag.FormatInt64(o.WorkspaceID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

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

// NewGetServiceVersionsByBlueprintNameParams creates a new GetServiceVersionsByBlueprintNameParams object
// with the default values initialized.
func NewGetServiceVersionsByBlueprintNameParams() *GetServiceVersionsByBlueprintNameParams {
	var ()
	return &GetServiceVersionsByBlueprintNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetServiceVersionsByBlueprintNameParamsWithTimeout creates a new GetServiceVersionsByBlueprintNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetServiceVersionsByBlueprintNameParamsWithTimeout(timeout time.Duration) *GetServiceVersionsByBlueprintNameParams {
	var ()
	return &GetServiceVersionsByBlueprintNameParams{

		timeout: timeout,
	}
}

// NewGetServiceVersionsByBlueprintNameParamsWithContext creates a new GetServiceVersionsByBlueprintNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetServiceVersionsByBlueprintNameParamsWithContext(ctx context.Context) *GetServiceVersionsByBlueprintNameParams {
	var ()
	return &GetServiceVersionsByBlueprintNameParams{

		Context: ctx,
	}
}

// NewGetServiceVersionsByBlueprintNameParamsWithHTTPClient creates a new GetServiceVersionsByBlueprintNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetServiceVersionsByBlueprintNameParamsWithHTTPClient(client *http.Client) *GetServiceVersionsByBlueprintNameParams {
	var ()
	return &GetServiceVersionsByBlueprintNameParams{
		HTTPClient: client,
	}
}

/*GetServiceVersionsByBlueprintNameParams contains all the parameters to send to the API endpoint
for the get service versions by blueprint name operation typically these are written to a http.Request
*/
type GetServiceVersionsByBlueprintNameParams struct {

	/*BlueprintName*/
	BlueprintName *string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get service versions by blueprint name params
func (o *GetServiceVersionsByBlueprintNameParams) WithTimeout(timeout time.Duration) *GetServiceVersionsByBlueprintNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get service versions by blueprint name params
func (o *GetServiceVersionsByBlueprintNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get service versions by blueprint name params
func (o *GetServiceVersionsByBlueprintNameParams) WithContext(ctx context.Context) *GetServiceVersionsByBlueprintNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get service versions by blueprint name params
func (o *GetServiceVersionsByBlueprintNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get service versions by blueprint name params
func (o *GetServiceVersionsByBlueprintNameParams) WithHTTPClient(client *http.Client) *GetServiceVersionsByBlueprintNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get service versions by blueprint name params
func (o *GetServiceVersionsByBlueprintNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBlueprintName adds the blueprintName to the get service versions by blueprint name params
func (o *GetServiceVersionsByBlueprintNameParams) WithBlueprintName(blueprintName *string) *GetServiceVersionsByBlueprintNameParams {
	o.SetBlueprintName(blueprintName)
	return o
}

// SetBlueprintName adds the blueprintName to the get service versions by blueprint name params
func (o *GetServiceVersionsByBlueprintNameParams) SetBlueprintName(blueprintName *string) {
	o.BlueprintName = blueprintName
}

// WithWorkspaceID adds the workspaceID to the get service versions by blueprint name params
func (o *GetServiceVersionsByBlueprintNameParams) WithWorkspaceID(workspaceID int64) *GetServiceVersionsByBlueprintNameParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the get service versions by blueprint name params
func (o *GetServiceVersionsByBlueprintNameParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetServiceVersionsByBlueprintNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BlueprintName != nil {

		// query param blueprintName
		var qrBlueprintName string
		if o.BlueprintName != nil {
			qrBlueprintName = *o.BlueprintName
		}
		qBlueprintName := qrBlueprintName
		if qBlueprintName != "" {
			if err := r.SetQueryParam("blueprintName", qBlueprintName); err != nil {
				return err
			}
		}

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

// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id

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

// NewGetKnoxServicesParams creates a new GetKnoxServicesParams object
// with the default values initialized.
func NewGetKnoxServicesParams() *GetKnoxServicesParams {
	var ()
	return &GetKnoxServicesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetKnoxServicesParamsWithTimeout creates a new GetKnoxServicesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetKnoxServicesParamsWithTimeout(timeout time.Duration) *GetKnoxServicesParams {
	var ()
	return &GetKnoxServicesParams{

		timeout: timeout,
	}
}

// NewGetKnoxServicesParamsWithContext creates a new GetKnoxServicesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetKnoxServicesParamsWithContext(ctx context.Context) *GetKnoxServicesParams {
	var ()
	return &GetKnoxServicesParams{

		Context: ctx,
	}
}

// NewGetKnoxServicesParamsWithHTTPClient creates a new GetKnoxServicesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetKnoxServicesParamsWithHTTPClient(client *http.Client) *GetKnoxServicesParams {
	var ()
	return &GetKnoxServicesParams{
		HTTPClient: client,
	}
}

/*GetKnoxServicesParams contains all the parameters to send to the API endpoint
for the get knox services operation typically these are written to a http.Request
*/
type GetKnoxServicesParams struct {

	/*BlueprintName*/
	BlueprintName *string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get knox services params
func (o *GetKnoxServicesParams) WithTimeout(timeout time.Duration) *GetKnoxServicesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get knox services params
func (o *GetKnoxServicesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get knox services params
func (o *GetKnoxServicesParams) WithContext(ctx context.Context) *GetKnoxServicesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get knox services params
func (o *GetKnoxServicesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get knox services params
func (o *GetKnoxServicesParams) WithHTTPClient(client *http.Client) *GetKnoxServicesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get knox services params
func (o *GetKnoxServicesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBlueprintName adds the blueprintName to the get knox services params
func (o *GetKnoxServicesParams) WithBlueprintName(blueprintName *string) *GetKnoxServicesParams {
	o.SetBlueprintName(blueprintName)
	return o
}

// SetBlueprintName adds the blueprintName to the get knox services params
func (o *GetKnoxServicesParams) SetBlueprintName(blueprintName *string) {
	o.BlueprintName = blueprintName
}

// WithWorkspaceID adds the workspaceID to the get knox services params
func (o *GetKnoxServicesParams) WithWorkspaceID(workspaceID int64) *GetKnoxServicesParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the get knox services params
func (o *GetKnoxServicesParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetKnoxServicesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_cluster_templates

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

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// NewCreateClusterTemplateInWorkspaceParams creates a new CreateClusterTemplateInWorkspaceParams object
// with the default values initialized.
func NewCreateClusterTemplateInWorkspaceParams() *CreateClusterTemplateInWorkspaceParams {
	var ()
	return &CreateClusterTemplateInWorkspaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateClusterTemplateInWorkspaceParamsWithTimeout creates a new CreateClusterTemplateInWorkspaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateClusterTemplateInWorkspaceParamsWithTimeout(timeout time.Duration) *CreateClusterTemplateInWorkspaceParams {
	var ()
	return &CreateClusterTemplateInWorkspaceParams{

		timeout: timeout,
	}
}

// NewCreateClusterTemplateInWorkspaceParamsWithContext creates a new CreateClusterTemplateInWorkspaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateClusterTemplateInWorkspaceParamsWithContext(ctx context.Context) *CreateClusterTemplateInWorkspaceParams {
	var ()
	return &CreateClusterTemplateInWorkspaceParams{

		Context: ctx,
	}
}

// NewCreateClusterTemplateInWorkspaceParamsWithHTTPClient creates a new CreateClusterTemplateInWorkspaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateClusterTemplateInWorkspaceParamsWithHTTPClient(client *http.Client) *CreateClusterTemplateInWorkspaceParams {
	var ()
	return &CreateClusterTemplateInWorkspaceParams{
		HTTPClient: client,
	}
}

/*CreateClusterTemplateInWorkspaceParams contains all the parameters to send to the API endpoint
for the create cluster template in workspace operation typically these are written to a http.Request
*/
type CreateClusterTemplateInWorkspaceParams struct {

	/*Body*/
	Body *model.ClusterTemplateV4Request
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create cluster template in workspace params
func (o *CreateClusterTemplateInWorkspaceParams) WithTimeout(timeout time.Duration) *CreateClusterTemplateInWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create cluster template in workspace params
func (o *CreateClusterTemplateInWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create cluster template in workspace params
func (o *CreateClusterTemplateInWorkspaceParams) WithContext(ctx context.Context) *CreateClusterTemplateInWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create cluster template in workspace params
func (o *CreateClusterTemplateInWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create cluster template in workspace params
func (o *CreateClusterTemplateInWorkspaceParams) WithHTTPClient(client *http.Client) *CreateClusterTemplateInWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create cluster template in workspace params
func (o *CreateClusterTemplateInWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create cluster template in workspace params
func (o *CreateClusterTemplateInWorkspaceParams) WithBody(body *model.ClusterTemplateV4Request) *CreateClusterTemplateInWorkspaceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create cluster template in workspace params
func (o *CreateClusterTemplateInWorkspaceParams) SetBody(body *model.ClusterTemplateV4Request) {
	o.Body = body
}

// WithWorkspaceID adds the workspaceID to the create cluster template in workspace params
func (o *CreateClusterTemplateInWorkspaceParams) WithWorkspaceID(workspaceID int64) *CreateClusterTemplateInWorkspaceParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the create cluster template in workspace params
func (o *CreateClusterTemplateInWorkspaceParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateClusterTemplateInWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
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

// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_recipes

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

// NewGetCreateRecipeRequestForCliParams creates a new GetCreateRecipeRequestForCliParams object
// with the default values initialized.
func NewGetCreateRecipeRequestForCliParams() *GetCreateRecipeRequestForCliParams {
	var ()
	return &GetCreateRecipeRequestForCliParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCreateRecipeRequestForCliParamsWithTimeout creates a new GetCreateRecipeRequestForCliParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCreateRecipeRequestForCliParamsWithTimeout(timeout time.Duration) *GetCreateRecipeRequestForCliParams {
	var ()
	return &GetCreateRecipeRequestForCliParams{

		timeout: timeout,
	}
}

// NewGetCreateRecipeRequestForCliParamsWithContext creates a new GetCreateRecipeRequestForCliParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCreateRecipeRequestForCliParamsWithContext(ctx context.Context) *GetCreateRecipeRequestForCliParams {
	var ()
	return &GetCreateRecipeRequestForCliParams{

		Context: ctx,
	}
}

// NewGetCreateRecipeRequestForCliParamsWithHTTPClient creates a new GetCreateRecipeRequestForCliParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCreateRecipeRequestForCliParamsWithHTTPClient(client *http.Client) *GetCreateRecipeRequestForCliParams {
	var ()
	return &GetCreateRecipeRequestForCliParams{
		HTTPClient: client,
	}
}

/*GetCreateRecipeRequestForCliParams contains all the parameters to send to the API endpoint
for the get create recipe request for cli operation typically these are written to a http.Request
*/
type GetCreateRecipeRequestForCliParams struct {

	/*Body*/
	Body *model.RecipeV4Request
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get create recipe request for cli params
func (o *GetCreateRecipeRequestForCliParams) WithTimeout(timeout time.Duration) *GetCreateRecipeRequestForCliParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get create recipe request for cli params
func (o *GetCreateRecipeRequestForCliParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get create recipe request for cli params
func (o *GetCreateRecipeRequestForCliParams) WithContext(ctx context.Context) *GetCreateRecipeRequestForCliParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get create recipe request for cli params
func (o *GetCreateRecipeRequestForCliParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get create recipe request for cli params
func (o *GetCreateRecipeRequestForCliParams) WithHTTPClient(client *http.Client) *GetCreateRecipeRequestForCliParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get create recipe request for cli params
func (o *GetCreateRecipeRequestForCliParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get create recipe request for cli params
func (o *GetCreateRecipeRequestForCliParams) WithBody(body *model.RecipeV4Request) *GetCreateRecipeRequestForCliParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get create recipe request for cli params
func (o *GetCreateRecipeRequestForCliParams) SetBody(body *model.RecipeV4Request) {
	o.Body = body
}

// WithWorkspaceID adds the workspaceID to the get create recipe request for cli params
func (o *GetCreateRecipeRequestForCliParams) WithWorkspaceID(workspaceID int64) *GetCreateRecipeRequestForCliParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the get create recipe request for cli params
func (o *GetCreateRecipeRequestForCliParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetCreateRecipeRequestForCliParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

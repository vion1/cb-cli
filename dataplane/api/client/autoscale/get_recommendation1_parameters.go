// Code generated by go-swagger; DO NOT EDIT.

package autoscale

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetRecommendation1Params creates a new GetRecommendation1Params object
// with the default values initialized.
func NewGetRecommendation1Params() *GetRecommendation1Params {
	var ()
	return &GetRecommendation1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRecommendation1ParamsWithTimeout creates a new GetRecommendation1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRecommendation1ParamsWithTimeout(timeout time.Duration) *GetRecommendation1Params {
	var ()
	return &GetRecommendation1Params{

		timeout: timeout,
	}
}

// NewGetRecommendation1ParamsWithContext creates a new GetRecommendation1Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetRecommendation1ParamsWithContext(ctx context.Context) *GetRecommendation1Params {
	var ()
	return &GetRecommendation1Params{

		Context: ctx,
	}
}

// NewGetRecommendation1ParamsWithHTTPClient creates a new GetRecommendation1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRecommendation1ParamsWithHTTPClient(client *http.Client) *GetRecommendation1Params {
	var ()
	return &GetRecommendation1Params{
		HTTPClient: client,
	}
}

/*GetRecommendation1Params contains all the parameters to send to the API endpoint
for the get recommendation 1 operation typically these are written to a http.Request
*/
type GetRecommendation1Params struct {

	/*BlueprintName*/
	BlueprintName *string
	/*WorkspaceID*/
	WorkspaceID *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get recommendation 1 params
func (o *GetRecommendation1Params) WithTimeout(timeout time.Duration) *GetRecommendation1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get recommendation 1 params
func (o *GetRecommendation1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get recommendation 1 params
func (o *GetRecommendation1Params) WithContext(ctx context.Context) *GetRecommendation1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get recommendation 1 params
func (o *GetRecommendation1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get recommendation 1 params
func (o *GetRecommendation1Params) WithHTTPClient(client *http.Client) *GetRecommendation1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get recommendation 1 params
func (o *GetRecommendation1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBlueprintName adds the blueprintName to the get recommendation 1 params
func (o *GetRecommendation1Params) WithBlueprintName(blueprintName *string) *GetRecommendation1Params {
	o.SetBlueprintName(blueprintName)
	return o
}

// SetBlueprintName adds the blueprintName to the get recommendation 1 params
func (o *GetRecommendation1Params) SetBlueprintName(blueprintName *string) {
	o.BlueprintName = blueprintName
}

// WithWorkspaceID adds the workspaceID to the get recommendation 1 params
func (o *GetRecommendation1Params) WithWorkspaceID(workspaceID *int64) *GetRecommendation1Params {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the get recommendation 1 params
func (o *GetRecommendation1Params) SetWorkspaceID(workspaceID *int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRecommendation1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.WorkspaceID != nil {

		// query param workspaceId
		var qrWorkspaceID int64
		if o.WorkspaceID != nil {
			qrWorkspaceID = *o.WorkspaceID
		}
		qWorkspaceID := swag.FormatInt64(qrWorkspaceID)
		if qWorkspaceID != "" {
			if err := r.SetQueryParam("workspaceId", qWorkspaceID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

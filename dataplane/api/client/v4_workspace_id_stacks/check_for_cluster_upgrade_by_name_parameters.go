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

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// NewCheckForClusterUpgradeByNameParams creates a new CheckForClusterUpgradeByNameParams object
// with the default values initialized.
func NewCheckForClusterUpgradeByNameParams() *CheckForClusterUpgradeByNameParams {
	var ()
	return &CheckForClusterUpgradeByNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCheckForClusterUpgradeByNameParamsWithTimeout creates a new CheckForClusterUpgradeByNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCheckForClusterUpgradeByNameParamsWithTimeout(timeout time.Duration) *CheckForClusterUpgradeByNameParams {
	var ()
	return &CheckForClusterUpgradeByNameParams{

		timeout: timeout,
	}
}

// NewCheckForClusterUpgradeByNameParamsWithContext creates a new CheckForClusterUpgradeByNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewCheckForClusterUpgradeByNameParamsWithContext(ctx context.Context) *CheckForClusterUpgradeByNameParams {
	var ()
	return &CheckForClusterUpgradeByNameParams{

		Context: ctx,
	}
}

// NewCheckForClusterUpgradeByNameParamsWithHTTPClient creates a new CheckForClusterUpgradeByNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCheckForClusterUpgradeByNameParamsWithHTTPClient(client *http.Client) *CheckForClusterUpgradeByNameParams {
	var ()
	return &CheckForClusterUpgradeByNameParams{
		HTTPClient: client,
	}
}

/*CheckForClusterUpgradeByNameParams contains all the parameters to send to the API endpoint
for the check for cluster upgrade by name operation typically these are written to a http.Request
*/
type CheckForClusterUpgradeByNameParams struct {

	/*AccountID*/
	AccountID *string
	/*Body*/
	Body *model.UpgradeV4Request
	/*Name*/
	Name string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the check for cluster upgrade by name params
func (o *CheckForClusterUpgradeByNameParams) WithTimeout(timeout time.Duration) *CheckForClusterUpgradeByNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the check for cluster upgrade by name params
func (o *CheckForClusterUpgradeByNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the check for cluster upgrade by name params
func (o *CheckForClusterUpgradeByNameParams) WithContext(ctx context.Context) *CheckForClusterUpgradeByNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the check for cluster upgrade by name params
func (o *CheckForClusterUpgradeByNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the check for cluster upgrade by name params
func (o *CheckForClusterUpgradeByNameParams) WithHTTPClient(client *http.Client) *CheckForClusterUpgradeByNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the check for cluster upgrade by name params
func (o *CheckForClusterUpgradeByNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the check for cluster upgrade by name params
func (o *CheckForClusterUpgradeByNameParams) WithAccountID(accountID *string) *CheckForClusterUpgradeByNameParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the check for cluster upgrade by name params
func (o *CheckForClusterUpgradeByNameParams) SetAccountID(accountID *string) {
	o.AccountID = accountID
}

// WithBody adds the body to the check for cluster upgrade by name params
func (o *CheckForClusterUpgradeByNameParams) WithBody(body *model.UpgradeV4Request) *CheckForClusterUpgradeByNameParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the check for cluster upgrade by name params
func (o *CheckForClusterUpgradeByNameParams) SetBody(body *model.UpgradeV4Request) {
	o.Body = body
}

// WithName adds the name to the check for cluster upgrade by name params
func (o *CheckForClusterUpgradeByNameParams) WithName(name string) *CheckForClusterUpgradeByNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the check for cluster upgrade by name params
func (o *CheckForClusterUpgradeByNameParams) SetName(name string) {
	o.Name = name
}

// WithWorkspaceID adds the workspaceID to the check for cluster upgrade by name params
func (o *CheckForClusterUpgradeByNameParams) WithWorkspaceID(workspaceID int64) *CheckForClusterUpgradeByNameParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the check for cluster upgrade by name params
func (o *CheckForClusterUpgradeByNameParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *CheckForClusterUpgradeByNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccountID != nil {

		// query param accountId
		var qrAccountID string
		if o.AccountID != nil {
			qrAccountID = *o.AccountID
		}
		qAccountID := qrAccountID
		if qAccountID != "" {
			if err := r.SetQueryParam("accountId", qAccountID); err != nil {
				return err
			}
		}

	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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

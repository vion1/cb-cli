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

// NewDatabaseRestoreParams creates a new DatabaseRestoreParams object
// with the default values initialized.
func NewDatabaseRestoreParams() *DatabaseRestoreParams {
	var ()
	return &DatabaseRestoreParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDatabaseRestoreParamsWithTimeout creates a new DatabaseRestoreParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDatabaseRestoreParamsWithTimeout(timeout time.Duration) *DatabaseRestoreParams {
	var ()
	return &DatabaseRestoreParams{

		timeout: timeout,
	}
}

// NewDatabaseRestoreParamsWithContext creates a new DatabaseRestoreParams object
// with the default values initialized, and the ability to set a context for a request
func NewDatabaseRestoreParamsWithContext(ctx context.Context) *DatabaseRestoreParams {
	var ()
	return &DatabaseRestoreParams{

		Context: ctx,
	}
}

// NewDatabaseRestoreParamsWithHTTPClient creates a new DatabaseRestoreParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDatabaseRestoreParamsWithHTTPClient(client *http.Client) *DatabaseRestoreParams {
	var ()
	return &DatabaseRestoreParams{
		HTTPClient: client,
	}
}

/*DatabaseRestoreParams contains all the parameters to send to the API endpoint
for the database restore operation typically these are written to a http.Request
*/
type DatabaseRestoreParams struct {

	/*AccountID*/
	AccountID *string
	/*BackupID*/
	BackupID *string
	/*BackupLocation*/
	BackupLocation *string
	/*Name*/
	Name string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the database restore params
func (o *DatabaseRestoreParams) WithTimeout(timeout time.Duration) *DatabaseRestoreParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the database restore params
func (o *DatabaseRestoreParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the database restore params
func (o *DatabaseRestoreParams) WithContext(ctx context.Context) *DatabaseRestoreParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the database restore params
func (o *DatabaseRestoreParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the database restore params
func (o *DatabaseRestoreParams) WithHTTPClient(client *http.Client) *DatabaseRestoreParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the database restore params
func (o *DatabaseRestoreParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the database restore params
func (o *DatabaseRestoreParams) WithAccountID(accountID *string) *DatabaseRestoreParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the database restore params
func (o *DatabaseRestoreParams) SetAccountID(accountID *string) {
	o.AccountID = accountID
}

// WithBackupID adds the backupID to the database restore params
func (o *DatabaseRestoreParams) WithBackupID(backupID *string) *DatabaseRestoreParams {
	o.SetBackupID(backupID)
	return o
}

// SetBackupID adds the backupId to the database restore params
func (o *DatabaseRestoreParams) SetBackupID(backupID *string) {
	o.BackupID = backupID
}

// WithBackupLocation adds the backupLocation to the database restore params
func (o *DatabaseRestoreParams) WithBackupLocation(backupLocation *string) *DatabaseRestoreParams {
	o.SetBackupLocation(backupLocation)
	return o
}

// SetBackupLocation adds the backupLocation to the database restore params
func (o *DatabaseRestoreParams) SetBackupLocation(backupLocation *string) {
	o.BackupLocation = backupLocation
}

// WithName adds the name to the database restore params
func (o *DatabaseRestoreParams) WithName(name string) *DatabaseRestoreParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the database restore params
func (o *DatabaseRestoreParams) SetName(name string) {
	o.Name = name
}

// WithWorkspaceID adds the workspaceID to the database restore params
func (o *DatabaseRestoreParams) WithWorkspaceID(workspaceID int64) *DatabaseRestoreParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the database restore params
func (o *DatabaseRestoreParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *DatabaseRestoreParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.BackupID != nil {

		// query param backupId
		var qrBackupID string
		if o.BackupID != nil {
			qrBackupID = *o.BackupID
		}
		qBackupID := qrBackupID
		if qBackupID != "" {
			if err := r.SetQueryParam("backupId", qBackupID); err != nil {
				return err
			}
		}

	}

	if o.BackupLocation != nil {

		// query param backupLocation
		var qrBackupLocation string
		if o.BackupLocation != nil {
			qrBackupLocation = *o.BackupLocation
		}
		qBackupLocation := qrBackupLocation
		if qBackupLocation != "" {
			if err := r.SetQueryParam("backupLocation", qBackupLocation); err != nil {
				return err
			}
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

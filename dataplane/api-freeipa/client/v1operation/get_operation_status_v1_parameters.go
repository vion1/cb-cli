// Code generated by go-swagger; DO NOT EDIT.

package v1operation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetOperationStatusV1Params creates a new GetOperationStatusV1Params object
// with the default values initialized.
func NewGetOperationStatusV1Params() *GetOperationStatusV1Params {
	var ()
	return &GetOperationStatusV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOperationStatusV1ParamsWithTimeout creates a new GetOperationStatusV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOperationStatusV1ParamsWithTimeout(timeout time.Duration) *GetOperationStatusV1Params {
	var ()
	return &GetOperationStatusV1Params{

		timeout: timeout,
	}
}

// NewGetOperationStatusV1ParamsWithContext creates a new GetOperationStatusV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetOperationStatusV1ParamsWithContext(ctx context.Context) *GetOperationStatusV1Params {
	var ()
	return &GetOperationStatusV1Params{

		Context: ctx,
	}
}

// NewGetOperationStatusV1ParamsWithHTTPClient creates a new GetOperationStatusV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOperationStatusV1ParamsWithHTTPClient(client *http.Client) *GetOperationStatusV1Params {
	var ()
	return &GetOperationStatusV1Params{
		HTTPClient: client,
	}
}

/*GetOperationStatusV1Params contains all the parameters to send to the API endpoint
for the get operation status v1 operation typically these are written to a http.Request
*/
type GetOperationStatusV1Params struct {

	/*AccountID*/
	AccountID *string
	/*OperationID*/
	OperationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get operation status v1 params
func (o *GetOperationStatusV1Params) WithTimeout(timeout time.Duration) *GetOperationStatusV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get operation status v1 params
func (o *GetOperationStatusV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get operation status v1 params
func (o *GetOperationStatusV1Params) WithContext(ctx context.Context) *GetOperationStatusV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get operation status v1 params
func (o *GetOperationStatusV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get operation status v1 params
func (o *GetOperationStatusV1Params) WithHTTPClient(client *http.Client) *GetOperationStatusV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get operation status v1 params
func (o *GetOperationStatusV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the get operation status v1 params
func (o *GetOperationStatusV1Params) WithAccountID(accountID *string) *GetOperationStatusV1Params {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get operation status v1 params
func (o *GetOperationStatusV1Params) SetAccountID(accountID *string) {
	o.AccountID = accountID
}

// WithOperationID adds the operationID to the get operation status v1 params
func (o *GetOperationStatusV1Params) WithOperationID(operationID string) *GetOperationStatusV1Params {
	o.SetOperationID(operationID)
	return o
}

// SetOperationID adds the operationId to the get operation status v1 params
func (o *GetOperationStatusV1Params) SetOperationID(operationID string) {
	o.OperationID = operationID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOperationStatusV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param operationId
	qrOperationID := o.OperationID
	qOperationID := qrOperationID
	if qOperationID != "" {
		if err := r.SetQueryParam("operationId", qOperationID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

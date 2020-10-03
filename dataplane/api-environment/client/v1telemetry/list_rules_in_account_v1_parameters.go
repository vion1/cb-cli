// Code generated by go-swagger; DO NOT EDIT.

package v1telemetry

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

// NewListRulesInAccountV1Params creates a new ListRulesInAccountV1Params object
// with the default values initialized.
func NewListRulesInAccountV1Params() *ListRulesInAccountV1Params {
	var ()
	return &ListRulesInAccountV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewListRulesInAccountV1ParamsWithTimeout creates a new ListRulesInAccountV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewListRulesInAccountV1ParamsWithTimeout(timeout time.Duration) *ListRulesInAccountV1Params {
	var ()
	return &ListRulesInAccountV1Params{

		timeout: timeout,
	}
}

// NewListRulesInAccountV1ParamsWithContext creates a new ListRulesInAccountV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewListRulesInAccountV1ParamsWithContext(ctx context.Context) *ListRulesInAccountV1Params {
	var ()
	return &ListRulesInAccountV1Params{

		Context: ctx,
	}
}

// NewListRulesInAccountV1ParamsWithHTTPClient creates a new ListRulesInAccountV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListRulesInAccountV1ParamsWithHTTPClient(client *http.Client) *ListRulesInAccountV1Params {
	var ()
	return &ListRulesInAccountV1Params{
		HTTPClient: client,
	}
}

/*ListRulesInAccountV1Params contains all the parameters to send to the API endpoint
for the list rules in account v1 operation typically these are written to a http.Request
*/
type ListRulesInAccountV1Params struct {

	/*AccountID*/
	AccountID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list rules in account v1 params
func (o *ListRulesInAccountV1Params) WithTimeout(timeout time.Duration) *ListRulesInAccountV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list rules in account v1 params
func (o *ListRulesInAccountV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list rules in account v1 params
func (o *ListRulesInAccountV1Params) WithContext(ctx context.Context) *ListRulesInAccountV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list rules in account v1 params
func (o *ListRulesInAccountV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list rules in account v1 params
func (o *ListRulesInAccountV1Params) WithHTTPClient(client *http.Client) *ListRulesInAccountV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list rules in account v1 params
func (o *ListRulesInAccountV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountID adds the accountID to the list rules in account v1 params
func (o *ListRulesInAccountV1Params) WithAccountID(accountID string) *ListRulesInAccountV1Params {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the list rules in account v1 params
func (o *ListRulesInAccountV1Params) SetAccountID(accountID string) {
	o.AccountID = accountID
}

// WriteToRequest writes these params to a swagger request
func (o *ListRulesInAccountV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param accountId
	if err := r.SetPathParam("accountId", o.AccountID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

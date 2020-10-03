// Code generated by go-swagger; DO NOT EDIT.

package v1distrox

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

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// NewCollectDistroxCmDiagnosticsV4Params creates a new CollectDistroxCmDiagnosticsV4Params object
// with the default values initialized.
func NewCollectDistroxCmDiagnosticsV4Params() *CollectDistroxCmDiagnosticsV4Params {
	var ()
	return &CollectDistroxCmDiagnosticsV4Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewCollectDistroxCmDiagnosticsV4ParamsWithTimeout creates a new CollectDistroxCmDiagnosticsV4Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewCollectDistroxCmDiagnosticsV4ParamsWithTimeout(timeout time.Duration) *CollectDistroxCmDiagnosticsV4Params {
	var ()
	return &CollectDistroxCmDiagnosticsV4Params{

		timeout: timeout,
	}
}

// NewCollectDistroxCmDiagnosticsV4ParamsWithContext creates a new CollectDistroxCmDiagnosticsV4Params object
// with the default values initialized, and the ability to set a context for a request
func NewCollectDistroxCmDiagnosticsV4ParamsWithContext(ctx context.Context) *CollectDistroxCmDiagnosticsV4Params {
	var ()
	return &CollectDistroxCmDiagnosticsV4Params{

		Context: ctx,
	}
}

// NewCollectDistroxCmDiagnosticsV4ParamsWithHTTPClient creates a new CollectDistroxCmDiagnosticsV4Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCollectDistroxCmDiagnosticsV4ParamsWithHTTPClient(client *http.Client) *CollectDistroxCmDiagnosticsV4Params {
	var ()
	return &CollectDistroxCmDiagnosticsV4Params{
		HTTPClient: client,
	}
}

/*CollectDistroxCmDiagnosticsV4Params contains all the parameters to send to the API endpoint
for the collect distrox cm diagnostics v4 operation typically these are written to a http.Request
*/
type CollectDistroxCmDiagnosticsV4Params struct {

	/*Body*/
	Body *model.DiagnosticsCollectionV1Request

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the collect distrox cm diagnostics v4 params
func (o *CollectDistroxCmDiagnosticsV4Params) WithTimeout(timeout time.Duration) *CollectDistroxCmDiagnosticsV4Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the collect distrox cm diagnostics v4 params
func (o *CollectDistroxCmDiagnosticsV4Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the collect distrox cm diagnostics v4 params
func (o *CollectDistroxCmDiagnosticsV4Params) WithContext(ctx context.Context) *CollectDistroxCmDiagnosticsV4Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the collect distrox cm diagnostics v4 params
func (o *CollectDistroxCmDiagnosticsV4Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the collect distrox cm diagnostics v4 params
func (o *CollectDistroxCmDiagnosticsV4Params) WithHTTPClient(client *http.Client) *CollectDistroxCmDiagnosticsV4Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the collect distrox cm diagnostics v4 params
func (o *CollectDistroxCmDiagnosticsV4Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the collect distrox cm diagnostics v4 params
func (o *CollectDistroxCmDiagnosticsV4Params) WithBody(body *model.DiagnosticsCollectionV1Request) *CollectDistroxCmDiagnosticsV4Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the collect distrox cm diagnostics v4 params
func (o *CollectDistroxCmDiagnosticsV4Params) SetBody(body *model.DiagnosticsCollectionV1Request) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CollectDistroxCmDiagnosticsV4Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

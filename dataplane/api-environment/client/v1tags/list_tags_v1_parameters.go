// Code generated by go-swagger; DO NOT EDIT.

package v1tags

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListTagsV1Params creates a new ListTagsV1Params object
// with the default values initialized.
func NewListTagsV1Params() *ListTagsV1Params {

	return &ListTagsV1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewListTagsV1ParamsWithTimeout creates a new ListTagsV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewListTagsV1ParamsWithTimeout(timeout time.Duration) *ListTagsV1Params {

	return &ListTagsV1Params{

		timeout: timeout,
	}
}

// NewListTagsV1ParamsWithContext creates a new ListTagsV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewListTagsV1ParamsWithContext(ctx context.Context) *ListTagsV1Params {

	return &ListTagsV1Params{

		Context: ctx,
	}
}

// NewListTagsV1ParamsWithHTTPClient creates a new ListTagsV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListTagsV1ParamsWithHTTPClient(client *http.Client) *ListTagsV1Params {

	return &ListTagsV1Params{
		HTTPClient: client,
	}
}

/*ListTagsV1Params contains all the parameters to send to the API endpoint
for the list tags v1 operation typically these are written to a http.Request
*/
type ListTagsV1Params struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list tags v1 params
func (o *ListTagsV1Params) WithTimeout(timeout time.Duration) *ListTagsV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list tags v1 params
func (o *ListTagsV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list tags v1 params
func (o *ListTagsV1Params) WithContext(ctx context.Context) *ListTagsV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list tags v1 params
func (o *ListTagsV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list tags v1 params
func (o *ListTagsV1Params) WithHTTPClient(client *http.Client) *ListTagsV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list tags v1 params
func (o *ListTagsV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListTagsV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package v1flexsubscriptions

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

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// NewPostPrivateFlexSubscriptionParams creates a new PostPrivateFlexSubscriptionParams object
// with the default values initialized.
func NewPostPrivateFlexSubscriptionParams() *PostPrivateFlexSubscriptionParams {
	var ()
	return &PostPrivateFlexSubscriptionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostPrivateFlexSubscriptionParamsWithTimeout creates a new PostPrivateFlexSubscriptionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostPrivateFlexSubscriptionParamsWithTimeout(timeout time.Duration) *PostPrivateFlexSubscriptionParams {
	var ()
	return &PostPrivateFlexSubscriptionParams{

		timeout: timeout,
	}
}

// NewPostPrivateFlexSubscriptionParamsWithContext creates a new PostPrivateFlexSubscriptionParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostPrivateFlexSubscriptionParamsWithContext(ctx context.Context) *PostPrivateFlexSubscriptionParams {
	var ()
	return &PostPrivateFlexSubscriptionParams{

		Context: ctx,
	}
}

// NewPostPrivateFlexSubscriptionParamsWithHTTPClient creates a new PostPrivateFlexSubscriptionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostPrivateFlexSubscriptionParamsWithHTTPClient(client *http.Client) *PostPrivateFlexSubscriptionParams {
	var ()
	return &PostPrivateFlexSubscriptionParams{
		HTTPClient: client,
	}
}

/*PostPrivateFlexSubscriptionParams contains all the parameters to send to the API endpoint
for the post private flex subscription operation typically these are written to a http.Request
*/
type PostPrivateFlexSubscriptionParams struct {

	/*Body*/
	Body *model.FlexSubscriptionRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post private flex subscription params
func (o *PostPrivateFlexSubscriptionParams) WithTimeout(timeout time.Duration) *PostPrivateFlexSubscriptionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post private flex subscription params
func (o *PostPrivateFlexSubscriptionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post private flex subscription params
func (o *PostPrivateFlexSubscriptionParams) WithContext(ctx context.Context) *PostPrivateFlexSubscriptionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post private flex subscription params
func (o *PostPrivateFlexSubscriptionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post private flex subscription params
func (o *PostPrivateFlexSubscriptionParams) WithHTTPClient(client *http.Client) *PostPrivateFlexSubscriptionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post private flex subscription params
func (o *PostPrivateFlexSubscriptionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post private flex subscription params
func (o *PostPrivateFlexSubscriptionParams) WithBody(body *model.FlexSubscriptionRequest) *PostPrivateFlexSubscriptionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post private flex subscription params
func (o *PostPrivateFlexSubscriptionParams) SetBody(body *model.FlexSubscriptionRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostPrivateFlexSubscriptionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

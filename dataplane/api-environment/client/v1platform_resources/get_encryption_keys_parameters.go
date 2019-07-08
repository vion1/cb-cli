// Code generated by go-swagger; DO NOT EDIT.

package v1platform_resources

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

// NewGetEncryptionKeysParams creates a new GetEncryptionKeysParams object
// with the default values initialized.
func NewGetEncryptionKeysParams() *GetEncryptionKeysParams {
	var ()
	return &GetEncryptionKeysParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEncryptionKeysParamsWithTimeout creates a new GetEncryptionKeysParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEncryptionKeysParamsWithTimeout(timeout time.Duration) *GetEncryptionKeysParams {
	var ()
	return &GetEncryptionKeysParams{

		timeout: timeout,
	}
}

// NewGetEncryptionKeysParamsWithContext creates a new GetEncryptionKeysParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEncryptionKeysParamsWithContext(ctx context.Context) *GetEncryptionKeysParams {
	var ()
	return &GetEncryptionKeysParams{

		Context: ctx,
	}
}

// NewGetEncryptionKeysParamsWithHTTPClient creates a new GetEncryptionKeysParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEncryptionKeysParamsWithHTTPClient(client *http.Client) *GetEncryptionKeysParams {
	var ()
	return &GetEncryptionKeysParams{
		HTTPClient: client,
	}
}

/*GetEncryptionKeysParams contains all the parameters to send to the API endpoint
for the get encryption keys operation typically these are written to a http.Request
*/
type GetEncryptionKeysParams struct {

	/*AvailabilityZone*/
	AvailabilityZone *string
	/*CredentialCrn*/
	CredentialCrn *string
	/*CredentialName*/
	CredentialName *string
	/*PlatformVariant*/
	PlatformVariant *string
	/*Region*/
	Region *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get encryption keys params
func (o *GetEncryptionKeysParams) WithTimeout(timeout time.Duration) *GetEncryptionKeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get encryption keys params
func (o *GetEncryptionKeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get encryption keys params
func (o *GetEncryptionKeysParams) WithContext(ctx context.Context) *GetEncryptionKeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get encryption keys params
func (o *GetEncryptionKeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get encryption keys params
func (o *GetEncryptionKeysParams) WithHTTPClient(client *http.Client) *GetEncryptionKeysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get encryption keys params
func (o *GetEncryptionKeysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAvailabilityZone adds the availabilityZone to the get encryption keys params
func (o *GetEncryptionKeysParams) WithAvailabilityZone(availabilityZone *string) *GetEncryptionKeysParams {
	o.SetAvailabilityZone(availabilityZone)
	return o
}

// SetAvailabilityZone adds the availabilityZone to the get encryption keys params
func (o *GetEncryptionKeysParams) SetAvailabilityZone(availabilityZone *string) {
	o.AvailabilityZone = availabilityZone
}

// WithCredentialCrn adds the credentialCrn to the get encryption keys params
func (o *GetEncryptionKeysParams) WithCredentialCrn(credentialCrn *string) *GetEncryptionKeysParams {
	o.SetCredentialCrn(credentialCrn)
	return o
}

// SetCredentialCrn adds the credentialCrn to the get encryption keys params
func (o *GetEncryptionKeysParams) SetCredentialCrn(credentialCrn *string) {
	o.CredentialCrn = credentialCrn
}

// WithCredentialName adds the credentialName to the get encryption keys params
func (o *GetEncryptionKeysParams) WithCredentialName(credentialName *string) *GetEncryptionKeysParams {
	o.SetCredentialName(credentialName)
	return o
}

// SetCredentialName adds the credentialName to the get encryption keys params
func (o *GetEncryptionKeysParams) SetCredentialName(credentialName *string) {
	o.CredentialName = credentialName
}

// WithPlatformVariant adds the platformVariant to the get encryption keys params
func (o *GetEncryptionKeysParams) WithPlatformVariant(platformVariant *string) *GetEncryptionKeysParams {
	o.SetPlatformVariant(platformVariant)
	return o
}

// SetPlatformVariant adds the platformVariant to the get encryption keys params
func (o *GetEncryptionKeysParams) SetPlatformVariant(platformVariant *string) {
	o.PlatformVariant = platformVariant
}

// WithRegion adds the region to the get encryption keys params
func (o *GetEncryptionKeysParams) WithRegion(region *string) *GetEncryptionKeysParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the get encryption keys params
func (o *GetEncryptionKeysParams) SetRegion(region *string) {
	o.Region = region
}

// WriteToRequest writes these params to a swagger request
func (o *GetEncryptionKeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AvailabilityZone != nil {

		// query param availabilityZone
		var qrAvailabilityZone string
		if o.AvailabilityZone != nil {
			qrAvailabilityZone = *o.AvailabilityZone
		}
		qAvailabilityZone := qrAvailabilityZone
		if qAvailabilityZone != "" {
			if err := r.SetQueryParam("availabilityZone", qAvailabilityZone); err != nil {
				return err
			}
		}

	}

	if o.CredentialCrn != nil {

		// query param credentialCrn
		var qrCredentialCrn string
		if o.CredentialCrn != nil {
			qrCredentialCrn = *o.CredentialCrn
		}
		qCredentialCrn := qrCredentialCrn
		if qCredentialCrn != "" {
			if err := r.SetQueryParam("credentialCrn", qCredentialCrn); err != nil {
				return err
			}
		}

	}

	if o.CredentialName != nil {

		// query param credentialName
		var qrCredentialName string
		if o.CredentialName != nil {
			qrCredentialName = *o.CredentialName
		}
		qCredentialName := qrCredentialName
		if qCredentialName != "" {
			if err := r.SetQueryParam("credentialName", qCredentialName); err != nil {
				return err
			}
		}

	}

	if o.PlatformVariant != nil {

		// query param platformVariant
		var qrPlatformVariant string
		if o.PlatformVariant != nil {
			qrPlatformVariant = *o.PlatformVariant
		}
		qPlatformVariant := qrPlatformVariant
		if qPlatformVariant != "" {
			if err := r.SetQueryParam("platformVariant", qPlatformVariant); err != nil {
				return err
			}
		}

	}

	if o.Region != nil {

		// query param region
		var qrRegion string
		if o.Region != nil {
			qrRegion = *o.Region
		}
		qRegion := qrRegion
		if qRegion != "" {
			if err := r.SetQueryParam("region", qRegion); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

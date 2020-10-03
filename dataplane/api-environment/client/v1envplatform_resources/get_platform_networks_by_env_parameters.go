// Code generated by go-swagger; DO NOT EDIT.

package v1envplatform_resources

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

// NewGetPlatformNetworksByEnvParams creates a new GetPlatformNetworksByEnvParams object
// with the default values initialized.
func NewGetPlatformNetworksByEnvParams() *GetPlatformNetworksByEnvParams {
	var ()
	return &GetPlatformNetworksByEnvParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPlatformNetworksByEnvParamsWithTimeout creates a new GetPlatformNetworksByEnvParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPlatformNetworksByEnvParamsWithTimeout(timeout time.Duration) *GetPlatformNetworksByEnvParams {
	var ()
	return &GetPlatformNetworksByEnvParams{

		timeout: timeout,
	}
}

// NewGetPlatformNetworksByEnvParamsWithContext creates a new GetPlatformNetworksByEnvParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPlatformNetworksByEnvParamsWithContext(ctx context.Context) *GetPlatformNetworksByEnvParams {
	var ()
	return &GetPlatformNetworksByEnvParams{

		Context: ctx,
	}
}

// NewGetPlatformNetworksByEnvParamsWithHTTPClient creates a new GetPlatformNetworksByEnvParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPlatformNetworksByEnvParamsWithHTTPClient(client *http.Client) *GetPlatformNetworksByEnvParams {
	var ()
	return &GetPlatformNetworksByEnvParams{
		HTTPClient: client,
	}
}

/*GetPlatformNetworksByEnvParams contains all the parameters to send to the API endpoint
for the get platform networks by env operation typically these are written to a http.Request
*/
type GetPlatformNetworksByEnvParams struct {

	/*AvailabilityZone*/
	AvailabilityZone *string
	/*EnvironmentCrn*/
	EnvironmentCrn *string
	/*PlatformVariant*/
	PlatformVariant *string
	/*Region*/
	Region *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get platform networks by env params
func (o *GetPlatformNetworksByEnvParams) WithTimeout(timeout time.Duration) *GetPlatformNetworksByEnvParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get platform networks by env params
func (o *GetPlatformNetworksByEnvParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get platform networks by env params
func (o *GetPlatformNetworksByEnvParams) WithContext(ctx context.Context) *GetPlatformNetworksByEnvParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get platform networks by env params
func (o *GetPlatformNetworksByEnvParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get platform networks by env params
func (o *GetPlatformNetworksByEnvParams) WithHTTPClient(client *http.Client) *GetPlatformNetworksByEnvParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get platform networks by env params
func (o *GetPlatformNetworksByEnvParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAvailabilityZone adds the availabilityZone to the get platform networks by env params
func (o *GetPlatformNetworksByEnvParams) WithAvailabilityZone(availabilityZone *string) *GetPlatformNetworksByEnvParams {
	o.SetAvailabilityZone(availabilityZone)
	return o
}

// SetAvailabilityZone adds the availabilityZone to the get platform networks by env params
func (o *GetPlatformNetworksByEnvParams) SetAvailabilityZone(availabilityZone *string) {
	o.AvailabilityZone = availabilityZone
}

// WithEnvironmentCrn adds the environmentCrn to the get platform networks by env params
func (o *GetPlatformNetworksByEnvParams) WithEnvironmentCrn(environmentCrn *string) *GetPlatformNetworksByEnvParams {
	o.SetEnvironmentCrn(environmentCrn)
	return o
}

// SetEnvironmentCrn adds the environmentCrn to the get platform networks by env params
func (o *GetPlatformNetworksByEnvParams) SetEnvironmentCrn(environmentCrn *string) {
	o.EnvironmentCrn = environmentCrn
}

// WithPlatformVariant adds the platformVariant to the get platform networks by env params
func (o *GetPlatformNetworksByEnvParams) WithPlatformVariant(platformVariant *string) *GetPlatformNetworksByEnvParams {
	o.SetPlatformVariant(platformVariant)
	return o
}

// SetPlatformVariant adds the platformVariant to the get platform networks by env params
func (o *GetPlatformNetworksByEnvParams) SetPlatformVariant(platformVariant *string) {
	o.PlatformVariant = platformVariant
}

// WithRegion adds the region to the get platform networks by env params
func (o *GetPlatformNetworksByEnvParams) WithRegion(region *string) *GetPlatformNetworksByEnvParams {
	o.SetRegion(region)
	return o
}

// SetRegion adds the region to the get platform networks by env params
func (o *GetPlatformNetworksByEnvParams) SetRegion(region *string) {
	o.Region = region
}

// WriteToRequest writes these params to a swagger request
func (o *GetPlatformNetworksByEnvParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.EnvironmentCrn != nil {

		// query param environmentCrn
		var qrEnvironmentCrn string
		if o.EnvironmentCrn != nil {
			qrEnvironmentCrn = *o.EnvironmentCrn
		}
		qEnvironmentCrn := qrEnvironmentCrn
		if qEnvironmentCrn != "" {
			if err := r.SetQueryParam("environmentCrn", qEnvironmentCrn); err != nil {
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

package smartsensesubscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewGetPrivateSmartSenseSubscriptionsParams creates a new GetPrivateSmartSenseSubscriptionsParams object
// with the default values initialized.
func NewGetPrivateSmartSenseSubscriptionsParams() *GetPrivateSmartSenseSubscriptionsParams {

	return &GetPrivateSmartSenseSubscriptionsParams{}
}

/*GetPrivateSmartSenseSubscriptionsParams contains all the parameters to send to the API endpoint
for the get private smart sense subscriptions operation typically these are written to a http.Request
*/
type GetPrivateSmartSenseSubscriptionsParams struct {
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrivateSmartSenseSubscriptionsParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

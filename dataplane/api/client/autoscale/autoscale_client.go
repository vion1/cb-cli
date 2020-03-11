// Code generated by go-swagger; DO NOT EDIT.

package autoscale

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new autoscale API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for autoscale API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AuthorizeForAutoscale authorize for autoscale API
*/
func (a *Client) AuthorizeForAutoscale(params *AuthorizeForAutoscaleParams) (*AuthorizeForAutoscaleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAuthorizeForAutoscaleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "authorizeForAutoscale",
		Method:             "GET",
		PathPattern:        "/autoscale/stack/crn/{crn}/authorize/{userId}/{tenant}/{permission}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AuthorizeForAutoscaleReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AuthorizeForAutoscaleOK), nil

}

/*
GetAllStackForAutoscale retrieves all stacks

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) GetAllStackForAutoscale(params *GetAllStackForAutoscaleParams) (*GetAllStackForAutoscaleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllStackForAutoscaleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAllStackForAutoscale",
		Method:             "GET",
		PathPattern:        "/autoscale/stack/all",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAllStackForAutoscaleReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAllStackForAutoscaleOK), nil

}

/*
GetCertificateStackForAutoscale retrieves the TLS certificate used by the gateway

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) GetCertificateStackForAutoscale(params *GetCertificateStackForAutoscaleParams) (*GetCertificateStackForAutoscaleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCertificateStackForAutoscaleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCertificateStackForAutoscale",
		Method:             "GET",
		PathPattern:        "/autoscale/stack/crn/{crn}/certificate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCertificateStackForAutoscaleReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCertificateStackForAutoscaleOK), nil

}

/*
GetClusterProxyconfiguration get cluster proxyconfiguration API
*/
func (a *Client) GetClusterProxyconfiguration(params *GetClusterProxyconfigurationParams) (*GetClusterProxyconfigurationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterProxyconfigurationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getClusterProxyconfiguration",
		Method:             "GET",
		PathPattern:        "/autoscale/clusterproxy",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetClusterProxyconfigurationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClusterProxyconfigurationOK), nil

}

/*
GetStackForAmbariForAutoscale retrieves stack by ambari address

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) GetStackForAmbariForAutoscale(params *GetStackForAmbariForAutoscaleParams) (*GetStackForAmbariForAutoscaleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStackForAmbariForAutoscaleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getStackForAmbariForAutoscale",
		Method:             "POST",
		PathPattern:        "/autoscale/ambari",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetStackForAmbariForAutoscaleReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetStackForAmbariForAutoscaleOK), nil

}

/*
GetStackForAutoscale retrieves stack by crn

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) GetStackForAutoscale(params *GetStackForAutoscaleParams) (*GetStackForAutoscaleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStackForAutoscaleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getStackForAutoscale",
		Method:             "GET",
		PathPattern:        "/autoscale/stack/crn/{crn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetStackForAutoscaleReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetStackForAutoscaleOK), nil

}

/*
GetStackStatusForAutoscale retrieves stack by crn

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) GetStackStatusForAutoscale(params *GetStackStatusForAutoscaleParams) (*GetStackStatusForAutoscaleOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetStackStatusForAutoscaleParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getStackStatusForAutoscale",
		Method:             "GET",
		PathPattern:        "/autoscale/stack/crn/{crn}/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetStackStatusForAutoscaleReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetStackStatusForAutoscaleOK), nil

}

/*
PutClusterForAutoscale updates stack by id

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) PutClusterForAutoscale(params *PutClusterForAutoscaleParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutClusterForAutoscaleParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putClusterForAutoscale",
		Method:             "PUT",
		PathPattern:        "/autoscale/stack/crn/{crn}/{userId}/cluster",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutClusterForAutoscaleReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
PutStackForAutoscale updates stack by id

Stacks are template instances - a running cloud infrastructure created based on a template. Stacks are always launched on behalf of a cloud user account. Stacks support a wide range of resources, allowing you to build a highly available, reliable, and scalable infrastructure for your application needs.
*/
func (a *Client) PutStackForAutoscale(params *PutStackForAutoscaleParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutStackForAutoscaleParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putStackForAutoscale",
		Method:             "PUT",
		PathPattern:        "/autoscale/stack/crn/{crn}/{userId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PutStackForAutoscaleReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

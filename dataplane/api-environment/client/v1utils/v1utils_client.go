// Code generated by go-swagger; DO NOT EDIT.

package v1utils

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new v1utils API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v1utils API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CheckClientVersionOfEnvironmentV1 checks the client version
*/
func (a *Client) CheckClientVersionOfEnvironmentV1(params *CheckClientVersionOfEnvironmentV1Params) (*CheckClientVersionOfEnvironmentV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCheckClientVersionOfEnvironmentV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "checkClientVersionOfEnvironmentV1",
		Method:             "GET",
		PathPattern:        "/v1/utils/client",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CheckClientVersionOfEnvironmentV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CheckClientVersionOfEnvironmentV1OK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

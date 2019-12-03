// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/dataplane/api-environment/client/v1credentials"
	"github.com/hortonworks/cb-cli/dataplane/api-environment/client/v1env"
	"github.com/hortonworks/cb-cli/dataplane/api-environment/client/v1platform_resources"
	"github.com/hortonworks/cb-cli/dataplane/api-environment/client/v1proxies"
	"github.com/hortonworks/cb-cli/dataplane/api-environment/client/v1utils"
)

// Default environment HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/api"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http", "https"}

// NewHTTPClient creates a new environment HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Environment {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new environment HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Environment {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new environment client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Environment {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Environment)
	cli.Transport = transport

	cli.V1credentials = v1credentials.New(transport, formats)

	cli.V1env = v1env.New(transport, formats)

	cli.V1platformResources = v1platform_resources.New(transport, formats)

	cli.V1proxies = v1proxies.New(transport, formats)

	cli.V1utils = v1utils.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Environment is a client for environment
type Environment struct {
	V1credentials *v1credentials.Client

	V1env *v1env.Client

	V1platformResources *v1platform_resources.Client

	V1proxies *v1proxies.Client

	V1utils *v1utils.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Environment) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.V1credentials.SetTransport(transport)

	c.V1env.SetTransport(transport)

	c.V1platformResources.SetTransport(transport)

	c.V1proxies.SetTransport(transport)

	c.V1utils.SetTransport(transport)

}

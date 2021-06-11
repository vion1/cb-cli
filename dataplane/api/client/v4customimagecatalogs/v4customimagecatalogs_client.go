// Code generated by go-swagger; DO NOT EDIT.

package v4customimagecatalogs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new v4customimagecatalogs API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v4customimagecatalogs API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateCustomImage creates a new image in an already existing custom image catalog

Provides an interface for custom image management.
*/
func (a *Client) CreateCustomImage(params *CreateCustomImageParams) (*CreateCustomImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateCustomImageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createCustomImage",
		Method:             "POST",
		PathPattern:        "/v4/custom_image_catalogs/{name}/image",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateCustomImageReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateCustomImageOK), nil

}

/*
CreateCustomImageCatalog creates custom image catalog

Provides an interface for custom image management.
*/
func (a *Client) CreateCustomImageCatalog(params *CreateCustomImageCatalogParams) (*CreateCustomImageCatalogOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateCustomImageCatalogParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createCustomImageCatalog",
		Method:             "POST",
		PathPattern:        "/v4/custom_image_catalogs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateCustomImageCatalogReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateCustomImageCatalogOK), nil

}

/*
DeleteCustomImage deletes an existing image from a custom image catalog

Provides an interface to determine available Virtual Machine images for the given version of Cloudbreak.
*/
func (a *Client) DeleteCustomImage(params *DeleteCustomImageParams) (*DeleteCustomImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCustomImageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteCustomImage",
		Method:             "DELETE",
		PathPattern:        "/v4/custom_image_catalogs/{name}/image/{imageId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteCustomImageReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteCustomImageOK), nil

}

/*
DeleteCustomImageCatalogByName deletes custom image catalog by name

Provides an interface for custom image management.
*/
func (a *Client) DeleteCustomImageCatalogByName(params *DeleteCustomImageCatalogByNameParams) (*DeleteCustomImageCatalogByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteCustomImageCatalogByNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteCustomImageCatalogByName",
		Method:             "DELETE",
		PathPattern:        "/v4/custom_image_catalogs/name/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteCustomImageCatalogByNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteCustomImageCatalogByNameOK), nil

}

/*
GetCustomImage gets custom image by name in custom image catalog

Provides an interface to determine available Virtual Machine images for the given version of Cloudbreak.
*/
func (a *Client) GetCustomImage(params *GetCustomImageParams) (*GetCustomImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCustomImageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCustomImage",
		Method:             "GET",
		PathPattern:        "/v4/custom_image_catalogs/{name}/image/{imageId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCustomImageReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCustomImageOK), nil

}

/*
GetCustomImageCatalogByName gets custom image catalog by name

Provides an interface for custom image management.
*/
func (a *Client) GetCustomImageCatalogByName(params *GetCustomImageCatalogByNameParams) (*GetCustomImageCatalogByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCustomImageCatalogByNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCustomImageCatalogByName",
		Method:             "GET",
		PathPattern:        "/v4/custom_image_catalogs/name/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCustomImageCatalogByNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCustomImageCatalogByNameOK), nil

}

/*
ListCustomImageCatalogs lists image catalogs

Provides an interface for custom image management.
*/
func (a *Client) ListCustomImageCatalogs(params *ListCustomImageCatalogsParams) (*ListCustomImageCatalogsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListCustomImageCatalogsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listCustomImageCatalogs",
		Method:             "GET",
		PathPattern:        "/v4/custom_image_catalogs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListCustomImageCatalogsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListCustomImageCatalogsOK), nil

}

/*
UpdateCustomImage updates an existing image in a custom image catalog

Provides an interface to determine available Virtual Machine images for the given version of Cloudbreak.
*/
func (a *Client) UpdateCustomImage(params *UpdateCustomImageParams) (*UpdateCustomImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateCustomImageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateCustomImage",
		Method:             "PUT",
		PathPattern:        "/v4/custom_image_catalogs/{name}/image/{imageId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateCustomImageReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateCustomImageOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

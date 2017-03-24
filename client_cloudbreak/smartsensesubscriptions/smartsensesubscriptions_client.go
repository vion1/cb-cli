package smartsensesubscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// New creates a new smartsensesubscriptions API client.
func New(transport client.Transport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for smartsensesubscriptions API
*/
type Client struct {
	transport client.Transport
	formats   strfmt.Registry
}

/*
DeleteSmartSenseSubscriptionByID deletes smart sense subscription by id

SmartSense subscriptions could be configured.
*/
func (a *Client) DeleteSmartSenseSubscriptionByID(params *DeleteSmartSenseSubscriptionByIDParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSmartSenseSubscriptionByIDParams()
	}

	_, err := a.transport.Submit(&client.Operation{
		ID:                 "deleteSmartSenseSubscriptionById",
		Method:             "DELETE",
		PathPattern:        "/smartsensesubscriptions/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteSmartSenseSubscriptionByIDReader{formats: a.formats},
	})
	if err != nil {
		return err
	}
	return nil
}

/*
GetPrivateSmartSenseSubscriptions retrieves private smart sense subscriptions

SmartSense subscriptions could be configured.
*/
func (a *Client) GetPrivateSmartSenseSubscriptions(params *GetPrivateSmartSenseSubscriptionsParams) (*GetPrivateSmartSenseSubscriptionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrivateSmartSenseSubscriptionsParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getPrivateSmartSenseSubscriptions",
		Method:             "GET",
		PathPattern:        "/smartsensesubscriptions/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPrivateSmartSenseSubscriptionsReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPrivateSmartSenseSubscriptionsOK), nil
}

/*
GetPublicSmartSenseSubscriptions retrieves public and private owned smart sense subscriptions

SmartSense subscriptions could be configured.
*/
func (a *Client) GetPublicSmartSenseSubscriptions(params *GetPublicSmartSenseSubscriptionsParams) (*GetPublicSmartSenseSubscriptionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPublicSmartSenseSubscriptionsParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getPublicSmartSenseSubscriptions",
		Method:             "GET",
		PathPattern:        "/smartsensesubscriptions/account",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPublicSmartSenseSubscriptionsReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPublicSmartSenseSubscriptionsOK), nil
}

/*
GetSmartSenseSubscriptionByID retrieves smart sense subscription by id

SmartSense subscriptions could be configured.
*/
func (a *Client) GetSmartSenseSubscriptionByID(params *GetSmartSenseSubscriptionByIDParams) (*GetSmartSenseSubscriptionByIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSmartSenseSubscriptionByIDParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "getSmartSenseSubscriptionById",
		Method:             "GET",
		PathPattern:        "/smartsensesubscriptions/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSmartSenseSubscriptionByIDReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSmartSenseSubscriptionByIDOK), nil
}

/*
PostPrivateSmartSenseSubscription creates smart sense subscription as private resource

SmartSense subscriptions could be configured.
*/
func (a *Client) PostPrivateSmartSenseSubscription(params *PostPrivateSmartSenseSubscriptionParams) (*PostPrivateSmartSenseSubscriptionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostPrivateSmartSenseSubscriptionParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "postPrivateSmartSenseSubscription",
		Method:             "POST",
		PathPattern:        "/smartsensesubscriptions/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostPrivateSmartSenseSubscriptionReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostPrivateSmartSenseSubscriptionOK), nil
}

/*
PostPublicSmartSenseSubscription creates smart sense subscription as public resource

SmartSense subscriptions could be configured.
*/
func (a *Client) PostPublicSmartSenseSubscription(params *PostPublicSmartSenseSubscriptionParams) (*PostPublicSmartSenseSubscriptionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostPublicSmartSenseSubscriptionParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:                 "postPublicSmartSenseSubscription",
		Method:             "POST",
		PathPattern:        "/smartsensesubscriptions/account",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostPublicSmartSenseSubscriptionReader{formats: a.formats},
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostPublicSmartSenseSubscriptionOK), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport client.Transport) {
	a.transport = transport
}

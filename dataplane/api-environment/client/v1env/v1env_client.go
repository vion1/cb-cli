// Code generated by go-swagger; DO NOT EDIT.

package v1env

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new v1env API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v1env API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ChangeCredentialInEnvironmentV1 changes the credential of the environment and the clusters in the environment of a given name

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) ChangeCredentialInEnvironmentV1(params *ChangeCredentialInEnvironmentV1Params) (*ChangeCredentialInEnvironmentV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeCredentialInEnvironmentV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "changeCredentialInEnvironmentV1",
		Method:             "PUT",
		PathPattern:        "/v1/env/name/{name}/change_credential",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ChangeCredentialInEnvironmentV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ChangeCredentialInEnvironmentV1OK), nil

}

/*
ChangeCredentialInEnvironmentV1ByCrn changes the credential of the environment and the clusters in the environment of a given c r n

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) ChangeCredentialInEnvironmentV1ByCrn(params *ChangeCredentialInEnvironmentV1ByCrnParams) (*ChangeCredentialInEnvironmentV1ByCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeCredentialInEnvironmentV1ByCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "changeCredentialInEnvironmentV1ByCrn",
		Method:             "PUT",
		PathPattern:        "/v1/env/crn/{crn}/change_credential",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ChangeCredentialInEnvironmentV1ByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ChangeCredentialInEnvironmentV1ByCrnOK), nil

}

/*
ChangeTelemetryFeaturesInEnvironmentV1ByCrn changes telemetry features s of the environment in the environment of a given c r n

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) ChangeTelemetryFeaturesInEnvironmentV1ByCrn(params *ChangeTelemetryFeaturesInEnvironmentV1ByCrnParams) (*ChangeTelemetryFeaturesInEnvironmentV1ByCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeTelemetryFeaturesInEnvironmentV1ByCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "changeTelemetryFeaturesInEnvironmentV1ByCrn",
		Method:             "PUT",
		PathPattern:        "/v1/env/crn/{crn}/change_telemetry_features",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ChangeTelemetryFeaturesInEnvironmentV1ByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ChangeTelemetryFeaturesInEnvironmentV1ByCrnOK), nil

}

/*
ChangeTelemetryFeaturesInEnvironmentV1ByName changes telemetry features s of the environment in the environment of a given name

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) ChangeTelemetryFeaturesInEnvironmentV1ByName(params *ChangeTelemetryFeaturesInEnvironmentV1ByNameParams) (*ChangeTelemetryFeaturesInEnvironmentV1ByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewChangeTelemetryFeaturesInEnvironmentV1ByNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "changeTelemetryFeaturesInEnvironmentV1ByName",
		Method:             "PUT",
		PathPattern:        "/v1/env/name/{name}/change_telemetry_features",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ChangeTelemetryFeaturesInEnvironmentV1ByNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ChangeTelemetryFeaturesInEnvironmentV1ByNameOK), nil

}

/*
CreateEnvironmentV1 creates an environment

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) CreateEnvironmentV1(params *CreateEnvironmentV1Params) (*CreateEnvironmentV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateEnvironmentV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createEnvironmentV1",
		Method:             "POST",
		PathPattern:        "/v1/env",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateEnvironmentV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateEnvironmentV1OK), nil

}

/*
DeleteEnvironmentV1ByCrn deletes an environment by c r n only possible if no cluster is running in the environment

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) DeleteEnvironmentV1ByCrn(params *DeleteEnvironmentV1ByCrnParams) (*DeleteEnvironmentV1ByCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteEnvironmentV1ByCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteEnvironmentV1ByCrn",
		Method:             "DELETE",
		PathPattern:        "/v1/env/crn/{crn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteEnvironmentV1ByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteEnvironmentV1ByCrnOK), nil

}

/*
DeleteEnvironmentV1ByName deletes an environment by name only possible if no cluster is running in the environment

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) DeleteEnvironmentV1ByName(params *DeleteEnvironmentV1ByNameParams) (*DeleteEnvironmentV1ByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteEnvironmentV1ByNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteEnvironmentV1ByName",
		Method:             "DELETE",
		PathPattern:        "/v1/env/name/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteEnvironmentV1ByNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteEnvironmentV1ByNameOK), nil

}

/*
DeleteEnvironmentsByCrn deletes multiple environment by c r ns only possible if no cluster is running in the environments

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) DeleteEnvironmentsByCrn(params *DeleteEnvironmentsByCrnParams) (*DeleteEnvironmentsByCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteEnvironmentsByCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteEnvironmentsByCrn",
		Method:             "DELETE",
		PathPattern:        "/v1/env/crn",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteEnvironmentsByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteEnvironmentsByCrnOK), nil

}

/*
DeleteEnvironmentsByName deletes multiple environment by names only possible if no cluster is running in the environments

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) DeleteEnvironmentsByName(params *DeleteEnvironmentsByNameParams) (*DeleteEnvironmentsByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteEnvironmentsByNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteEnvironmentsByName",
		Method:             "DELETE",
		PathPattern:        "/v1/env/name",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteEnvironmentsByNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteEnvironmentsByNameOK), nil

}

/*
EditEnvironmentV1 edits an environment by name location regions and description can be changed

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) EditEnvironmentV1(params *EditEnvironmentV1Params) (*EditEnvironmentV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEditEnvironmentV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "editEnvironmentV1",
		Method:             "PUT",
		PathPattern:        "/v1/env/name/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EditEnvironmentV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*EditEnvironmentV1OK), nil

}

/*
EditEnvironmentV1ByCrn edits an environment by c r n location regions and description can be changed

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) EditEnvironmentV1ByCrn(params *EditEnvironmentV1ByCrnParams) (*EditEnvironmentV1ByCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEditEnvironmentV1ByCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "editEnvironmentV1ByCrn",
		Method:             "PUT",
		PathPattern:        "/v1/env/crn/{crn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &EditEnvironmentV1ByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*EditEnvironmentV1ByCrnOK), nil

}

/*
GetCreateEnvironmentForCli produces cli command input for environment creation

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) GetCreateEnvironmentForCli(params *GetCreateEnvironmentForCliParams) (*GetCreateEnvironmentForCliOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCreateEnvironmentForCliParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCreateEnvironmentForCli",
		Method:             "POST",
		PathPattern:        "/v1/env/cli_create",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCreateEnvironmentForCliReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCreateEnvironmentForCliOK), nil

}

/*
GetCreateEnvironmentForCliByCrn produces cli command input for environment creation

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) GetCreateEnvironmentForCliByCrn(params *GetCreateEnvironmentForCliByCrnParams) (*GetCreateEnvironmentForCliByCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCreateEnvironmentForCliByCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCreateEnvironmentForCliByCrn",
		Method:             "POST",
		PathPattern:        "/v1/env/crn/{crn}/cli_create",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCreateEnvironmentForCliByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCreateEnvironmentForCliByCrnOK), nil

}

/*
GetCreateEnvironmentForCliByName produces cli command input for environment creation

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) GetCreateEnvironmentForCliByName(params *GetCreateEnvironmentForCliByNameParams) (*GetCreateEnvironmentForCliByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCreateEnvironmentForCliByNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCreateEnvironmentForCliByName",
		Method:             "POST",
		PathPattern:        "/v1/env/name/{name}/cli_create",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCreateEnvironmentForCliByNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCreateEnvironmentForCliByNameOK), nil

}

/*
GetCrnByNameV1 gets the crn of an environment by name

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) GetCrnByNameV1(params *GetCrnByNameV1Params) (*GetCrnByNameV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCrnByNameV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCrnByNameV1",
		Method:             "GET",
		PathPattern:        "/v1/env/crnByName/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCrnByNameV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCrnByNameV1OK), nil

}

/*
GetEnvironmentFlowLogsProgressByResourceCrn lists recent flow operations progress details for resource by resource crn

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) GetEnvironmentFlowLogsProgressByResourceCrn(params *GetEnvironmentFlowLogsProgressByResourceCrnParams) (*GetEnvironmentFlowLogsProgressByResourceCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEnvironmentFlowLogsProgressByResourceCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getEnvironmentFlowLogsProgressByResourceCrn",
		Method:             "GET",
		PathPattern:        "/v1/env/progress/resource/crn/{resourceCrn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetEnvironmentFlowLogsProgressByResourceCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetEnvironmentFlowLogsProgressByResourceCrnOK), nil

}

/*
GetEnvironmentLastFlowLogProgressByResourceCrn gets last flow operation progress details for resource by resource crn

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) GetEnvironmentLastFlowLogProgressByResourceCrn(params *GetEnvironmentLastFlowLogProgressByResourceCrnParams) (*GetEnvironmentLastFlowLogProgressByResourceCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEnvironmentLastFlowLogProgressByResourceCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getEnvironmentLastFlowLogProgressByResourceCrn",
		Method:             "GET",
		PathPattern:        "/v1/env/progress/resource/crn/{resourceCrn}/last",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetEnvironmentLastFlowLogProgressByResourceCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetEnvironmentLastFlowLogProgressByResourceCrnOK), nil

}

/*
GetEnvironmentV1ByCrn gets an environment by c r n

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) GetEnvironmentV1ByCrn(params *GetEnvironmentV1ByCrnParams) (*GetEnvironmentV1ByCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEnvironmentV1ByCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getEnvironmentV1ByCrn",
		Method:             "GET",
		PathPattern:        "/v1/env/crn/{crn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetEnvironmentV1ByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetEnvironmentV1ByCrnOK), nil

}

/*
GetEnvironmentV1ByName gets an environment by name

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) GetEnvironmentV1ByName(params *GetEnvironmentV1ByNameParams) (*GetEnvironmentV1ByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetEnvironmentV1ByNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getEnvironmentV1ByName",
		Method:             "GET",
		PathPattern:        "/v1/env/name/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetEnvironmentV1ByNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetEnvironmentV1ByNameOK), nil

}

/*
InternalListEnvironmentV1 lists all environments by account ID using the internal actor

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) InternalListEnvironmentV1(params *InternalListEnvironmentV1Params) (*InternalListEnvironmentV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInternalListEnvironmentV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "internalListEnvironmentV1",
		Method:             "GET",
		PathPattern:        "/v1/env/internal",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &InternalListEnvironmentV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*InternalListEnvironmentV1OK), nil

}

/*
ListEnvironmentV1 lists all environments

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) ListEnvironmentV1(params *ListEnvironmentV1Params) (*ListEnvironmentV1OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListEnvironmentV1Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listEnvironmentV1",
		Method:             "GET",
		PathPattern:        "/v1/env",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListEnvironmentV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListEnvironmentV1OK), nil

}

/*
StartEnvironmentByCrnV1 starts an environment by c r n the freeipa datalake and datahubs will be started in this order

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) StartEnvironmentByCrnV1(params *StartEnvironmentByCrnV1Params) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartEnvironmentByCrnV1Params()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "startEnvironmentByCrnV1",
		Method:             "POST",
		PathPattern:        "/v1/env/crn/{crn}/start",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StartEnvironmentByCrnV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
StartEnvironmentByNameV1 starts an environment by name the freeipa datalake and datahubs will be started in this order

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) StartEnvironmentByNameV1(params *StartEnvironmentByNameV1Params) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartEnvironmentByNameV1Params()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "startEnvironmentByNameV1",
		Method:             "POST",
		PathPattern:        "/v1/env/name/{name}/start",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StartEnvironmentByNameV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
StopEnvironmentByCrnV1 stops an environment by c r n the datahubs datalake and freeipa will be stopped in this order

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) StopEnvironmentByCrnV1(params *StopEnvironmentByCrnV1Params) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStopEnvironmentByCrnV1Params()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "stopEnvironmentByCrnV1",
		Method:             "POST",
		PathPattern:        "/v1/env/crn/{crn}/stop",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StopEnvironmentByCrnV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
StopEnvironmentByNameV1 stops an environment by name the datahubs datalake and freeipa will be stopped in this order

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) StopEnvironmentByNameV1(params *StopEnvironmentByNameV1Params) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStopEnvironmentByNameV1Params()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "stopEnvironmentByNameV1",
		Method:             "POST",
		PathPattern:        "/v1/env/name/{name}/stop",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StopEnvironmentByNameV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
UpdateConfigsInEnvironmentByCrnV1 updates the configuration for all stacks in the environment by the environment c r n

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) UpdateConfigsInEnvironmentByCrnV1(params *UpdateConfigsInEnvironmentByCrnV1Params) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateConfigsInEnvironmentByCrnV1Params()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateConfigsInEnvironmentByCrnV1",
		Method:             "POST",
		PathPattern:        "/v1/env/crn/{crn}/update_config",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateConfigsInEnvironmentByCrnV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
UpdateEnvironmentLoadBalancersByCrnV1 updates all cluster in an environment with load balancers including adding the endpoint access gateway if it s enabled environment is specified by c r n

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) UpdateEnvironmentLoadBalancersByCrnV1(params *UpdateEnvironmentLoadBalancersByCrnV1Params) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateEnvironmentLoadBalancersByCrnV1Params()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateEnvironmentLoadBalancersByCrnV1",
		Method:             "PUT",
		PathPattern:        "/v1/env/crn/{crn}/update_load_balancers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateEnvironmentLoadBalancersByCrnV1Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
UpdateEnvironmentLoadBalancersByNameV11 updates all cluster in an environment with load balancers including adding the endpoint access gateway if it s enabled environment is specified by name

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) UpdateEnvironmentLoadBalancersByNameV11(params *UpdateEnvironmentLoadBalancersByNameV11Params) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateEnvironmentLoadBalancersByNameV11Params()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateEnvironmentLoadBalancersByNameV11",
		Method:             "PUT",
		PathPattern:        "/v1/env/name/{name}/update_load_balancers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateEnvironmentLoadBalancersByNameV11Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
ValidateCloudStorage validate cloud storage API
*/
func (a *Client) ValidateCloudStorage(params *ValidateCloudStorageParams) (*ValidateCloudStorageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateCloudStorageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "validateCloudStorage",
		Method:             "POST",
		PathPattern:        "/v1/env/validate_cloud_storage",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ValidateCloudStorageReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ValidateCloudStorageOK), nil

}

/*
VerifyCredentialByEnvCrn verifies the credential used by the given environment

Environment consists of a credential and various other resources and enables users to quickly create clusters in given regions in a given cloud provider.
*/
func (a *Client) VerifyCredentialByEnvCrn(params *VerifyCredentialByEnvCrnParams) (*VerifyCredentialByEnvCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVerifyCredentialByEnvCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "verifyCredentialByEnvCrn",
		Method:             "GET",
		PathPattern:        "/v1/env/crn/{crn}/verify_credential",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &VerifyCredentialByEnvCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*VerifyCredentialByEnvCrnOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

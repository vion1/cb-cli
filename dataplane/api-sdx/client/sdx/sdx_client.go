// Code generated by go-swagger; DO NOT EDIT.

package sdx

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new sdx API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for sdx API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CheckForUpgrade checks for upgrade options by name
*/
func (a *Client) CheckForUpgrade(params *CheckForUpgradeParams) (*CheckForUpgradeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCheckForUpgradeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "checkForUpgrade",
		Method:             "GET",
		PathPattern:        "/sdx/{name}/check_for_upgrade",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CheckForUpgradeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CheckForUpgradeOK), nil

}

/*
CheckForUpgradeByCrn checks for upgrade options by crn
*/
func (a *Client) CheckForUpgradeByCrn(params *CheckForUpgradeByCrnParams) (*CheckForUpgradeByCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCheckForUpgradeByCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "checkForUpgradeByCrn",
		Method:             "GET",
		PathPattern:        "/sdx/crn/{crn}/check_for_upgrade",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CheckForUpgradeByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CheckForUpgradeByCrnOK), nil

}

/*
CreateSdx creates s d x cluster
*/
func (a *Client) CreateSdx(params *CreateSdxParams) (*CreateSdxOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSdxParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createSdx",
		Method:             "POST",
		PathPattern:        "/sdx/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateSdxReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateSdxOK), nil

}

/*
DeleteSdx deletes s d x cluster
*/
func (a *Client) DeleteSdx(params *DeleteSdxParams) (*DeleteSdxOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSdxParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteSdx",
		Method:             "DELETE",
		PathPattern:        "/sdx/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteSdxReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteSdxOK), nil

}

/*
DeleteSdxByCrn deletes s d x cluster by crn
*/
func (a *Client) DeleteSdxByCrn(params *DeleteSdxByCrnParams) (*DeleteSdxByCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSdxByCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteSdxByCrn",
		Method:             "DELETE",
		PathPattern:        "/sdx/crn/{clusterCrn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteSdxByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteSdxByCrnOK), nil

}

/*
GetSdx gets s d x cluster
*/
func (a *Client) GetSdx(params *GetSdxParams) (*GetSdxOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSdxParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSdx",
		Method:             "GET",
		PathPattern:        "/sdx/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSdxReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSdxOK), nil

}

/*
GetSdxByCrn gets s d x cluster by crn
*/
func (a *Client) GetSdxByCrn(params *GetSdxByCrnParams) (*GetSdxByCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSdxByCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSdxByCrn",
		Method:             "GET",
		PathPattern:        "/sdx/crn/{clusterCrn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSdxByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSdxByCrnOK), nil

}

/*
GetSdxByEnvCrn gets s d x cluster by environment crn
*/
func (a *Client) GetSdxByEnvCrn(params *GetSdxByEnvCrnParams) (*GetSdxByEnvCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSdxByEnvCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSdxByEnvCrn",
		Method:             "GET",
		PathPattern:        "/sdx/envcrn/{envCrn}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSdxByEnvCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSdxByEnvCrnOK), nil

}

/*
GetSdxDetail gets s d x cluster detail
*/
func (a *Client) GetSdxDetail(params *GetSdxDetailParams) (*GetSdxDetailOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSdxDetailParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSdxDetail",
		Method:             "GET",
		PathPattern:        "/sdx/{name}/detail",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSdxDetailReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSdxDetailOK), nil

}

/*
GetSdxDetailByCrn gets s d x cluster detail by crn
*/
func (a *Client) GetSdxDetailByCrn(params *GetSdxDetailByCrnParams) (*GetSdxDetailByCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSdxDetailByCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSdxDetailByCrn",
		Method:             "GET",
		PathPattern:        "/sdx/crn/{clusterCrn}/detail",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetSdxDetailByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSdxDetailByCrnOK), nil

}

/*
ListSdx lists s d x clusters
*/
func (a *Client) ListSdx(params *ListSdxParams) (*ListSdxOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSdxParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listSdx",
		Method:             "GET",
		PathPattern:        "/sdx/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListSdxReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListSdxOK), nil

}

/*
RepairSdxNode repairs an sdx node in the specified hostgroup
*/
func (a *Client) RepairSdxNode(params *RepairSdxNodeParams) (*RepairSdxNodeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRepairSdxNodeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "repairSdxNode",
		Method:             "POST",
		PathPattern:        "/sdx/{name}/manual_repair",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RepairSdxNodeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RepairSdxNodeOK), nil

}

/*
RepairSdxNodeByCrn repairs an sdx node in the specified hostgroup by crn
*/
func (a *Client) RepairSdxNodeByCrn(params *RepairSdxNodeByCrnParams) (*RepairSdxNodeByCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRepairSdxNodeByCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "repairSdxNodeByCrn",
		Method:             "POST",
		PathPattern:        "/sdx/crn/{crn}/manual_repair",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RepairSdxNodeByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RepairSdxNodeByCrnOK), nil

}

/*
RetrySdx retries sdx
*/
func (a *Client) RetrySdx(params *RetrySdxParams) (*RetrySdxOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrySdxParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "retrySdx",
		Method:             "POST",
		PathPattern:        "/sdx/{name}/retry",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RetrySdxReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RetrySdxOK), nil

}

/*
RetrySdxByCrn retries sdx by crn
*/
func (a *Client) RetrySdxByCrn(params *RetrySdxByCrnParams) (*RetrySdxByCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrySdxByCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "retrySdxByCrn",
		Method:             "POST",
		PathPattern:        "/sdx/crn/{crn}/retry",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RetrySdxByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RetrySdxByCrnOK), nil

}

/*
StartSdxByCrn starts sdx by crn
*/
func (a *Client) StartSdxByCrn(params *StartSdxByCrnParams) (*StartSdxByCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartSdxByCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "startSdxByCrn",
		Method:             "POST",
		PathPattern:        "/sdx/crn/{crn}/start",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StartSdxByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StartSdxByCrnOK), nil

}

/*
StartSdxByName starts sdx
*/
func (a *Client) StartSdxByName(params *StartSdxByNameParams) (*StartSdxByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartSdxByNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "startSdxByName",
		Method:             "POST",
		PathPattern:        "/sdx/{name}/start",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StartSdxByNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StartSdxByNameOK), nil

}

/*
StopSdxByCrn stops sdx by crn
*/
func (a *Client) StopSdxByCrn(params *StopSdxByCrnParams) (*StopSdxByCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStopSdxByCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "stopSdxByCrn",
		Method:             "POST",
		PathPattern:        "/sdx/crn/{crn}/stop",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StopSdxByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StopSdxByCrnOK), nil

}

/*
StopSdxByName stops sdx
*/
func (a *Client) StopSdxByName(params *StopSdxByNameParams) (*StopSdxByNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStopSdxByNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "stopSdxByName",
		Method:             "POST",
		PathPattern:        "/sdx/{name}/stop",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StopSdxByNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StopSdxByNameOK), nil

}

/*
SyncSdx syncs s d x cluster by name
*/
func (a *Client) SyncSdx(params *SyncSdxParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSyncSdxParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "syncSdx",
		Method:             "POST",
		PathPattern:        "/sdx/{name}/sync",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SyncSdxReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
SyncSdxByCrn syncs s d x cluster by crn
*/
func (a *Client) SyncSdxByCrn(params *SyncSdxByCrnParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSyncSdxByCrnParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "syncSdxByCrn",
		Method:             "POST",
		PathPattern:        "/sdx/crn/{crn}/sync",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SyncSdxByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
UpgradeDatalakeCluster upgrades the datalake cluster to the latest images
*/
func (a *Client) UpgradeDatalakeCluster(params *UpgradeDatalakeClusterParams) (*UpgradeDatalakeClusterOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpgradeDatalakeClusterParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "upgradeDatalakeCluster",
		Method:             "POST",
		PathPattern:        "/sdx/{name}/upgrade",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpgradeDatalakeClusterReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpgradeDatalakeClusterOK), nil

}

/*
UpgradeDatalakeClusterByCrn upgrades the datalake cluster to the latest images
*/
func (a *Client) UpgradeDatalakeClusterByCrn(params *UpgradeDatalakeClusterByCrnParams) (*UpgradeDatalakeClusterByCrnOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpgradeDatalakeClusterByCrnParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "upgradeDatalakeClusterByCrn",
		Method:             "POST",
		PathPattern:        "/sdx/crn/{crn}/upgrade",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpgradeDatalakeClusterByCrnReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpgradeDatalakeClusterByCrnOK), nil

}

/*
Versions lists datalake versions
*/
func (a *Client) Versions(params *VersionsParams) (*VersionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVersionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "versions",
		Method:             "GET",
		PathPattern:        "/sdx/versions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &VersionsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*VersionsOK), nil

}

/*
Versions lists datalake versions
*/
func (a *Client) Versions(params *VersionsParams) (*VersionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVersionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "versions",
		Method:             "GET",
		PathPattern:        "/sdx/versions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &VersionsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*VersionsOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

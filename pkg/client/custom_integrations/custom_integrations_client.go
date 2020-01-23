// Code generated by go-swagger; DO NOT EDIT.

package custom_integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new custom integrations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for custom integrations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateUsingPOST creates a custom integration

Create a Custom Integration to be consumed in pipelines as custom tasks
*/
func (a *Client) CreateUsingPOST(params *CreateUsingPOSTParams) (*CreateUsingPOSTOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateUsingPOSTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createUsingPOST",
		Method:             "POST",
		PathPattern:        "/codestream/api/custom-integrations",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateUsingPOSTReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateUsingPOSTOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createUsingPOST: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteDraftByIDUsingDELETE deletes a custom integration and its versions

Delete a Custom Integration with the given id and all its versions
*/
func (a *Client) DeleteDraftByIDUsingDELETE(params *DeleteDraftByIDUsingDELETEParams) (*DeleteDraftByIDUsingDELETEOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDraftByIDUsingDELETEParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteDraftByIdUsingDELETE",
		Method:             "DELETE",
		PathPattern:        "/codestream/api/custom-integrations/{id}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteDraftByIDUsingDELETEReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteDraftByIDUsingDELETEOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteDraftByIdUsingDELETE: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteVersionByIDUsingDELETE deletes a custom integration version

Delete a Custom Integration version with the given id and version
*/
func (a *Client) DeleteVersionByIDUsingDELETE(params *DeleteVersionByIDUsingDELETEParams) (*DeleteVersionByIDUsingDELETEOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteVersionByIDUsingDELETEParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteVersionByIdUsingDELETE",
		Method:             "DELETE",
		PathPattern:        "/codestream/api/custom-integrations/{id}/versions/{version}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteVersionByIDUsingDELETEReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteVersionByIDUsingDELETEOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteVersionByIdUsingDELETE: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeprecateByIDAndVersionUsingPOST deprecates a custom integration version

Deprecate a Custom Integration version
*/
func (a *Client) DeprecateByIDAndVersionUsingPOST(params *DeprecateByIDAndVersionUsingPOSTParams) (*DeprecateByIDAndVersionUsingPOSTOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeprecateByIDAndVersionUsingPOSTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deprecateByIdAndVersionUsingPOST",
		Method:             "POST",
		PathPattern:        "/codestream/api/custom-integrations/{id}/versions/{version}/deprecate",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeprecateByIDAndVersionUsingPOSTReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeprecateByIDAndVersionUsingPOSTOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deprecateByIdAndVersionUsingPOST: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetCustomIntegrationVersionsByIDUsingGET gets all versions of a custom integration by id

Get all versions of a Custom Integration with specified id, paging and filter parameters
*/
func (a *Client) GetCustomIntegrationVersionsByIDUsingGET(params *GetCustomIntegrationVersionsByIDUsingGETParams) (*GetCustomIntegrationVersionsByIDUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCustomIntegrationVersionsByIDUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCustomIntegrationVersionsByIdUsingGET",
		Method:             "GET",
		PathPattern:        "/codestream/api/custom-integrations/{id}/versions",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCustomIntegrationVersionsByIDUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCustomIntegrationVersionsByIDUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCustomIntegrationVersionsByIdUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetCustomIntegrationsUsingGET gets all custom integrations

Get all Custom Integrations with specified paging and filter parameters.
*/
func (a *Client) GetCustomIntegrationsUsingGET(params *GetCustomIntegrationsUsingGETParams) (*GetCustomIntegrationsUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCustomIntegrationsUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCustomIntegrationsUsingGET",
		Method:             "GET",
		PathPattern:        "/codestream/api/custom-integrations",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetCustomIntegrationsUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetCustomIntegrationsUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getCustomIntegrationsUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetDraftByIDUsingGET gets a custom integration by id

Get details of a Custom Integration with the given id
*/
func (a *Client) GetDraftByIDUsingGET(params *GetDraftByIDUsingGETParams) (*GetDraftByIDUsingGETOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDraftByIDUsingGETParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDraftByIdUsingGET",
		Method:             "GET",
		PathPattern:        "/codestream/api/custom-integrations/{id}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetDraftByIDUsingGETReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDraftByIDUsingGETOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDraftByIdUsingGET: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetVersionByIDUsingGETMixin3 gets a custom integration by version

Get a Custom Integration with the given id and version
*/
func (a *Client) GetVersionByIDUsingGETMixin3(params *GetVersionByIDUsingGETMixin3Params) (*GetVersionByIDUsingGETMixin3OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVersionByIDUsingGETMixin3Params()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getVersionByIdUsingGETMixin3",
		Method:             "GET",
		PathPattern:        "/codestream/api/custom-integrations/{id}/versions/{version}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetVersionByIDUsingGETMixin3Reader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetVersionByIDUsingGETMixin3OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getVersionByIdUsingGETMixin3: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ReleaseByIDAndVersionUsingPOST releases a custom integration version

Release a Custom Integration version to be consumable in pipelines
*/
func (a *Client) ReleaseByIDAndVersionUsingPOST(params *ReleaseByIDAndVersionUsingPOSTParams) (*ReleaseByIDAndVersionUsingPOSTOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReleaseByIDAndVersionUsingPOSTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "releaseByIdAndVersionUsingPOST",
		Method:             "POST",
		PathPattern:        "/codestream/api/custom-integrations/{id}/versions/{version}/release",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ReleaseByIDAndVersionUsingPOSTReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReleaseByIDAndVersionUsingPOSTOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for releaseByIdAndVersionUsingPOST: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RestoreByIDAndVersionUsingPOST restores a custom integration from a version to current draft

Restore a Custom Integration from the given version to the current draft
*/
func (a *Client) RestoreByIDAndVersionUsingPOST(params *RestoreByIDAndVersionUsingPOSTParams) (*RestoreByIDAndVersionUsingPOSTOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRestoreByIDAndVersionUsingPOSTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "restoreByIdAndVersionUsingPOST",
		Method:             "POST",
		PathPattern:        "/codestream/api/custom-integrations/{id}/versions/{version}/restore",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RestoreByIDAndVersionUsingPOSTReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RestoreByIDAndVersionUsingPOSTOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for restoreByIdAndVersionUsingPOST: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateByIDUsingPUT updates a custom integration by id

Update a Custom Integration with the given id
*/
func (a *Client) UpdateByIDUsingPUT(params *UpdateByIDUsingPUTParams) (*UpdateByIDUsingPUTOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateByIDUsingPUTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateByIdUsingPUT",
		Method:             "PUT",
		PathPattern:        "/codestream/api/custom-integrations/{id}",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateByIDUsingPUTReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateByIDUsingPUTOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateByIdUsingPUT: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
VersionByIDUsingPOST creates a custom integration version

Create a Custom Integration version from the current draft
*/
func (a *Client) VersionByIDUsingPOST(params *VersionByIDUsingPOSTParams) (*VersionByIDUsingPOSTOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewVersionByIDUsingPOSTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "versionByIdUsingPOST",
		Method:             "POST",
		PathPattern:        "/codestream/api/custom-integrations/{id}/versions",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &VersionByIDUsingPOSTReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*VersionByIDUsingPOSTOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for versionByIdUsingPOST: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
WithdrawByIDAndVersionUsingPOST withdraws a custom integration version

Withdraw a released/deprecated Custom Integration version to make it un-consumable in pipelines
*/
func (a *Client) WithdrawByIDAndVersionUsingPOST(params *WithdrawByIDAndVersionUsingPOSTParams) (*WithdrawByIDAndVersionUsingPOSTOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewWithdrawByIDAndVersionUsingPOSTParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "withdrawByIdAndVersionUsingPOST",
		Method:             "POST",
		PathPattern:        "/codestream/api/custom-integrations/{id}/versions/{version}/withdraw",
		ProducesMediaTypes: []string{"*/*"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &WithdrawByIDAndVersionUsingPOSTReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*WithdrawByIDAndVersionUsingPOSTOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for withdrawByIdAndVersionUsingPOST: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
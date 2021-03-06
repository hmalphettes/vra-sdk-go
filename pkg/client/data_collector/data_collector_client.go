// Code generated by go-swagger; DO NOT EDIT.

package data_collector

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new data collector API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for data collector API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateDataCollector creates data collector

Create a new Data Collector.

Note: Data collector endpoints are not available in vRA on-prem release.
*/
func (a *Client) CreateDataCollector(params *CreateDataCollectorParams) (*CreateDataCollectorCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDataCollectorParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createDataCollector",
		Method:             "POST",
		PathPattern:        "/iaas/api/data-collectors",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateDataCollectorReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateDataCollectorCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createDataCollector: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteDataCollector deletes data collector

Delete Data Collector with a given id.

Note: Data collector endpoints are not available in vRA on-prem release.
*/
func (a *Client) DeleteDataCollector(params *DeleteDataCollectorParams) (*DeleteDataCollectorNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteDataCollectorParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteDataCollector",
		Method:             "DELETE",
		PathPattern:        "/iaas/api/data-collectors/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteDataCollectorReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteDataCollectorNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteDataCollector: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetDataCollector gets data collector

Get Data Collector with a given id.

Note: Data collector endpoints are not available in vRA on-prem release.
*/
func (a *Client) GetDataCollector(params *GetDataCollectorParams) (*GetDataCollectorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDataCollectorParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDataCollector",
		Method:             "GET",
		PathPattern:        "/iaas/api/data-collectors/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDataCollectorReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDataCollectorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDataCollector: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetDataCollectors gets data collectors

Get all Data Collectors.

Note: Data collector endpoints are not available in vRA on-prem release.
*/
func (a *Client) GetDataCollectors(params *GetDataCollectorsParams) (*GetDataCollectorsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDataCollectorsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDataCollectors",
		Method:             "GET",
		PathPattern:        "/iaas/api/data-collectors",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDataCollectorsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDataCollectorsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDataCollectors: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

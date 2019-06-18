// Code generated by go-swagger; DO NOT EDIT.

package load_balancer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new load balancer API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for load balancer API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateLoadBalancer creates load balancer

Create load balancer
*/
func (a *Client) CreateLoadBalancer(params *CreateLoadBalancerParams) (*CreateLoadBalancerAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateLoadBalancerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createLoadBalancer",
		Method:             "POST",
		PathPattern:        "/iaas/api/load-balancers",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateLoadBalancerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateLoadBalancerAccepted), nil

}

/*
DeleteLoadBalancer deletes load balancer

Delete load balancer with a given id
*/
func (a *Client) DeleteLoadBalancer(params *DeleteLoadBalancerParams) (*DeleteLoadBalancerAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteLoadBalancerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteLoadBalancer",
		Method:             "DELETE",
		PathPattern:        "/iaas/api/load-balancers/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteLoadBalancerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteLoadBalancerAccepted), nil

}

/*
DeleteLoadBalancerOperation deletes operation for load balancer

Second day delete operation for load balancer
*/
func (a *Client) DeleteLoadBalancerOperation(params *DeleteLoadBalancerOperationParams) (*DeleteLoadBalancerOperationAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteLoadBalancerOperationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteLoadBalancerOperation",
		Method:             "POST",
		PathPattern:        "/iaas/api/load-balancers/{id}/operations/delete",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteLoadBalancerOperationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteLoadBalancerOperationAccepted), nil

}

/*
GetLoadBalancer gets load balancer

Get load balancer with a given id
*/
func (a *Client) GetLoadBalancer(params *GetLoadBalancerParams) (*GetLoadBalancerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLoadBalancerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLoadBalancer",
		Method:             "GET",
		PathPattern:        "/iaas/api/load-balancers/{id}",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLoadBalancerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLoadBalancerOK), nil

}

/*
GetLoadBalancers gets load balancers

Get all load balancers
*/
func (a *Client) GetLoadBalancers(params *GetLoadBalancersParams) (*GetLoadBalancersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetLoadBalancersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLoadBalancers",
		Method:             "GET",
		PathPattern:        "/iaas/api/load-balancers",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLoadBalancersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLoadBalancersOK), nil

}

/*
ScaleLoadBalancer scales operation for load balancer

Second day scale operation for load balancer
*/
func (a *Client) ScaleLoadBalancer(params *ScaleLoadBalancerParams) (*ScaleLoadBalancerAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewScaleLoadBalancerParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "scaleLoadBalancer",
		Method:             "POST",
		PathPattern:        "/iaas/api/load-balancers/{id}/operations/scale",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ScaleLoadBalancerReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ScaleLoadBalancerAccepted), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

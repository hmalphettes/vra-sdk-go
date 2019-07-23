// Code generated by go-swagger; DO NOT EDIT.

package about

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new about API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for about API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetAboutPage gets about page

The page contains information about the supported API versions and the latest API version. The version parameter is optional but highly recommended.
If you do not specify explicitly an exact version, you will be calling the latest supported API version.
Here is an example of a call which specifies the exact version you are using: `GET /iaas/api/network-profiles?apiVersion=2019-01-15`
*/
func (a *Client) GetAboutPage(params *GetAboutPageParams) (*GetAboutPageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAboutPageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAboutPage",
		Method:             "GET",
		PathPattern:        "/iaas/api/about",
		ProducesMediaTypes: []string{"app/json", "application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAboutPageReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAboutPageOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

// Code generated by go-swagger; DO NOT EDIT.

package network_ip_range

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetNetworkIPRangeParams creates a new GetNetworkIPRangeParams object
// with the default values initialized.
func NewGetNetworkIPRangeParams() *GetNetworkIPRangeParams {
	var ()
	return &GetNetworkIPRangeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkIPRangeParamsWithTimeout creates a new GetNetworkIPRangeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNetworkIPRangeParamsWithTimeout(timeout time.Duration) *GetNetworkIPRangeParams {
	var ()
	return &GetNetworkIPRangeParams{

		timeout: timeout,
	}
}

// NewGetNetworkIPRangeParamsWithContext creates a new GetNetworkIPRangeParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNetworkIPRangeParamsWithContext(ctx context.Context) *GetNetworkIPRangeParams {
	var ()
	return &GetNetworkIPRangeParams{

		Context: ctx,
	}
}

// NewGetNetworkIPRangeParamsWithHTTPClient creates a new GetNetworkIPRangeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNetworkIPRangeParamsWithHTTPClient(client *http.Client) *GetNetworkIPRangeParams {
	var ()
	return &GetNetworkIPRangeParams{
		HTTPClient: client,
	}
}

/*GetNetworkIPRangeParams contains all the parameters to send to the API endpoint
for the get network IP range operation typically these are written to a http.Request
*/
type GetNetworkIPRangeParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /iaas/api/about

	*/
	APIVersion *string
	/*ID
	  The ID of the network IP range.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get network IP range params
func (o *GetNetworkIPRangeParams) WithTimeout(timeout time.Duration) *GetNetworkIPRangeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network IP range params
func (o *GetNetworkIPRangeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network IP range params
func (o *GetNetworkIPRangeParams) WithContext(ctx context.Context) *GetNetworkIPRangeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network IP range params
func (o *GetNetworkIPRangeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network IP range params
func (o *GetNetworkIPRangeParams) WithHTTPClient(client *http.Client) *GetNetworkIPRangeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network IP range params
func (o *GetNetworkIPRangeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get network IP range params
func (o *GetNetworkIPRangeParams) WithAPIVersion(aPIVersion *string) *GetNetworkIPRangeParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get network IP range params
func (o *GetNetworkIPRangeParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get network IP range params
func (o *GetNetworkIPRangeParams) WithID(id string) *GetNetworkIPRangeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get network IP range params
func (o *GetNetworkIPRangeParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkIPRangeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.APIVersion != nil {

		// query param apiVersion
		var qrAPIVersion string
		if o.APIVersion != nil {
			qrAPIVersion = *o.APIVersion
		}
		qAPIVersion := qrAPIVersion
		if qAPIVersion != "" {
			if err := r.SetQueryParam("apiVersion", qAPIVersion); err != nil {
				return err
			}
		}

	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
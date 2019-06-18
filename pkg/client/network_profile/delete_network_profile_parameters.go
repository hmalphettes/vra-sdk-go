// Code generated by go-swagger; DO NOT EDIT.

package network_profile

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

// NewDeleteNetworkProfileParams creates a new DeleteNetworkProfileParams object
// with the default values initialized.
func NewDeleteNetworkProfileParams() *DeleteNetworkProfileParams {
	var ()
	return &DeleteNetworkProfileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNetworkProfileParamsWithTimeout creates a new DeleteNetworkProfileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteNetworkProfileParamsWithTimeout(timeout time.Duration) *DeleteNetworkProfileParams {
	var ()
	return &DeleteNetworkProfileParams{

		timeout: timeout,
	}
}

// NewDeleteNetworkProfileParamsWithContext creates a new DeleteNetworkProfileParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteNetworkProfileParamsWithContext(ctx context.Context) *DeleteNetworkProfileParams {
	var ()
	return &DeleteNetworkProfileParams{

		Context: ctx,
	}
}

// NewDeleteNetworkProfileParamsWithHTTPClient creates a new DeleteNetworkProfileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteNetworkProfileParamsWithHTTPClient(client *http.Client) *DeleteNetworkProfileParams {
	var ()
	return &DeleteNetworkProfileParams{
		HTTPClient: client,
	}
}

/*DeleteNetworkProfileParams contains all the parameters to send to the API endpoint
for the delete network profile operation typically these are written to a http.Request
*/
type DeleteNetworkProfileParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /iaas/api/about

	*/
	APIVersion *string
	/*ID
	  The ID of the network profile.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete network profile params
func (o *DeleteNetworkProfileParams) WithTimeout(timeout time.Duration) *DeleteNetworkProfileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete network profile params
func (o *DeleteNetworkProfileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete network profile params
func (o *DeleteNetworkProfileParams) WithContext(ctx context.Context) *DeleteNetworkProfileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete network profile params
func (o *DeleteNetworkProfileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete network profile params
func (o *DeleteNetworkProfileParams) WithHTTPClient(client *http.Client) *DeleteNetworkProfileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete network profile params
func (o *DeleteNetworkProfileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the delete network profile params
func (o *DeleteNetworkProfileParams) WithAPIVersion(aPIVersion *string) *DeleteNetworkProfileParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the delete network profile params
func (o *DeleteNetworkProfileParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the delete network profile params
func (o *DeleteNetworkProfileParams) WithID(id string) *DeleteNetworkProfileParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete network profile params
func (o *DeleteNetworkProfileParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNetworkProfileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// Code generated by go-swagger; DO NOT EDIT.

package deployment_actions

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

// NewGetResourceActionsUsingGET1Params creates a new GetResourceActionsUsingGET1Params object
// with the default values initialized.
func NewGetResourceActionsUsingGET1Params() *GetResourceActionsUsingGET1Params {
	var ()
	return &GetResourceActionsUsingGET1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetResourceActionsUsingGET1ParamsWithTimeout creates a new GetResourceActionsUsingGET1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetResourceActionsUsingGET1ParamsWithTimeout(timeout time.Duration) *GetResourceActionsUsingGET1Params {
	var ()
	return &GetResourceActionsUsingGET1Params{

		timeout: timeout,
	}
}

// NewGetResourceActionsUsingGET1ParamsWithContext creates a new GetResourceActionsUsingGET1Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetResourceActionsUsingGET1ParamsWithContext(ctx context.Context) *GetResourceActionsUsingGET1Params {
	var ()
	return &GetResourceActionsUsingGET1Params{

		Context: ctx,
	}
}

// NewGetResourceActionsUsingGET1ParamsWithHTTPClient creates a new GetResourceActionsUsingGET1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetResourceActionsUsingGET1ParamsWithHTTPClient(client *http.Client) *GetResourceActionsUsingGET1Params {
	var ()
	return &GetResourceActionsUsingGET1Params{
		HTTPClient: client,
	}
}

/*GetResourceActionsUsingGET1Params contains all the parameters to send to the API endpoint
for the get resource actions using get1 operation typically these are written to a http.Request
*/
type GetResourceActionsUsingGET1Params struct {

	/*DepID
	  Deployment ID

	*/
	DepID strfmt.UUID
	/*ResourceID
	  Resource ID

	*/
	ResourceID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get resource actions using get1 params
func (o *GetResourceActionsUsingGET1Params) WithTimeout(timeout time.Duration) *GetResourceActionsUsingGET1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get resource actions using get1 params
func (o *GetResourceActionsUsingGET1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get resource actions using get1 params
func (o *GetResourceActionsUsingGET1Params) WithContext(ctx context.Context) *GetResourceActionsUsingGET1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get resource actions using get1 params
func (o *GetResourceActionsUsingGET1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get resource actions using get1 params
func (o *GetResourceActionsUsingGET1Params) WithHTTPClient(client *http.Client) *GetResourceActionsUsingGET1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get resource actions using get1 params
func (o *GetResourceActionsUsingGET1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDepID adds the depID to the get resource actions using get1 params
func (o *GetResourceActionsUsingGET1Params) WithDepID(depID strfmt.UUID) *GetResourceActionsUsingGET1Params {
	o.SetDepID(depID)
	return o
}

// SetDepID adds the depId to the get resource actions using get1 params
func (o *GetResourceActionsUsingGET1Params) SetDepID(depID strfmt.UUID) {
	o.DepID = depID
}

// WithResourceID adds the resourceID to the get resource actions using get1 params
func (o *GetResourceActionsUsingGET1Params) WithResourceID(resourceID strfmt.UUID) *GetResourceActionsUsingGET1Params {
	o.SetResourceID(resourceID)
	return o
}

// SetResourceID adds the resourceId to the get resource actions using get1 params
func (o *GetResourceActionsUsingGET1Params) SetResourceID(resourceID strfmt.UUID) {
	o.ResourceID = resourceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetResourceActionsUsingGET1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param depId
	if err := r.SetPathParam("depId", o.DepID.String()); err != nil {
		return err
	}

	// path param resourceId
	if err := r.SetPathParam("resourceId", o.ResourceID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
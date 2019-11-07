// Code generated by go-swagger; DO NOT EDIT.

package disk

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

// NewGetMachineDisksParams creates a new GetMachineDisksParams object
// with the default values initialized.
func NewGetMachineDisksParams() *GetMachineDisksParams {
	var ()
	return &GetMachineDisksParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMachineDisksParamsWithTimeout creates a new GetMachineDisksParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMachineDisksParamsWithTimeout(timeout time.Duration) *GetMachineDisksParams {
	var ()
	return &GetMachineDisksParams{

		timeout: timeout,
	}
}

// NewGetMachineDisksParamsWithContext creates a new GetMachineDisksParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMachineDisksParamsWithContext(ctx context.Context) *GetMachineDisksParams {
	var ()
	return &GetMachineDisksParams{

		Context: ctx,
	}
}

// NewGetMachineDisksParamsWithHTTPClient creates a new GetMachineDisksParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMachineDisksParamsWithHTTPClient(client *http.Client) *GetMachineDisksParams {
	var ()
	return &GetMachineDisksParams{
		HTTPClient: client,
	}
}

/*GetMachineDisksParams contains all the parameters to send to the API endpoint
for the get machine disks operation typically these are written to a http.Request
*/
type GetMachineDisksParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /iaas/api/about

	*/
	APIVersion *string
	/*ID
	  The ID of the machine.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get machine disks params
func (o *GetMachineDisksParams) WithTimeout(timeout time.Duration) *GetMachineDisksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get machine disks params
func (o *GetMachineDisksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get machine disks params
func (o *GetMachineDisksParams) WithContext(ctx context.Context) *GetMachineDisksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get machine disks params
func (o *GetMachineDisksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get machine disks params
func (o *GetMachineDisksParams) WithHTTPClient(client *http.Client) *GetMachineDisksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get machine disks params
func (o *GetMachineDisksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get machine disks params
func (o *GetMachineDisksParams) WithAPIVersion(aPIVersion *string) *GetMachineDisksParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get machine disks params
func (o *GetMachineDisksParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get machine disks params
func (o *GetMachineDisksParams) WithID(id string) *GetMachineDisksParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get machine disks params
func (o *GetMachineDisksParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetMachineDisksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewGetMachineDiskParams creates a new GetMachineDiskParams object
// with the default values initialized.
func NewGetMachineDiskParams() *GetMachineDiskParams {
	var ()
	return &GetMachineDiskParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMachineDiskParamsWithTimeout creates a new GetMachineDiskParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMachineDiskParamsWithTimeout(timeout time.Duration) *GetMachineDiskParams {
	var ()
	return &GetMachineDiskParams{

		timeout: timeout,
	}
}

// NewGetMachineDiskParamsWithContext creates a new GetMachineDiskParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMachineDiskParamsWithContext(ctx context.Context) *GetMachineDiskParams {
	var ()
	return &GetMachineDiskParams{

		Context: ctx,
	}
}

// NewGetMachineDiskParamsWithHTTPClient creates a new GetMachineDiskParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMachineDiskParamsWithHTTPClient(client *http.Client) *GetMachineDiskParams {
	var ()
	return &GetMachineDiskParams{
		HTTPClient: client,
	}
}

/*GetMachineDiskParams contains all the parameters to send to the API endpoint
for the get machine disk operation typically these are written to a http.Request
*/
type GetMachineDiskParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /iaas/api/about

	*/
	APIVersion *string
	/*ID
	  The ID of the machine.

	*/
	ID string
	/*Id1
	  The ID of the disk.

	*/
	Id1 string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get machine disk params
func (o *GetMachineDiskParams) WithTimeout(timeout time.Duration) *GetMachineDiskParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get machine disk params
func (o *GetMachineDiskParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get machine disk params
func (o *GetMachineDiskParams) WithContext(ctx context.Context) *GetMachineDiskParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get machine disk params
func (o *GetMachineDiskParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get machine disk params
func (o *GetMachineDiskParams) WithHTTPClient(client *http.Client) *GetMachineDiskParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get machine disk params
func (o *GetMachineDiskParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get machine disk params
func (o *GetMachineDiskParams) WithAPIVersion(aPIVersion *string) *GetMachineDiskParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get machine disk params
func (o *GetMachineDiskParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get machine disk params
func (o *GetMachineDiskParams) WithID(id string) *GetMachineDiskParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get machine disk params
func (o *GetMachineDiskParams) SetID(id string) {
	o.ID = id
}

// WithId1 adds the id1 to the get machine disk params
func (o *GetMachineDiskParams) WithId1(id1 string) *GetMachineDiskParams {
	o.SetId1(id1)
	return o
}

// SetId1 adds the id1 to the get machine disk params
func (o *GetMachineDiskParams) SetId1(id1 string) {
	o.Id1 = id1
}

// WriteToRequest writes these params to a swagger request
func (o *GetMachineDiskParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id1
	if err := r.SetPathParam("id1", o.Id1); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

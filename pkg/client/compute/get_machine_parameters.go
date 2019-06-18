// Code generated by go-swagger; DO NOT EDIT.

package compute

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

// NewGetMachineParams creates a new GetMachineParams object
// with the default values initialized.
func NewGetMachineParams() *GetMachineParams {
	var ()
	return &GetMachineParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetMachineParamsWithTimeout creates a new GetMachineParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetMachineParamsWithTimeout(timeout time.Duration) *GetMachineParams {
	var ()
	return &GetMachineParams{

		timeout: timeout,
	}
}

// NewGetMachineParamsWithContext creates a new GetMachineParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetMachineParamsWithContext(ctx context.Context) *GetMachineParams {
	var ()
	return &GetMachineParams{

		Context: ctx,
	}
}

// NewGetMachineParamsWithHTTPClient creates a new GetMachineParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetMachineParamsWithHTTPClient(client *http.Client) *GetMachineParams {
	var ()
	return &GetMachineParams{
		HTTPClient: client,
	}
}

/*GetMachineParams contains all the parameters to send to the API endpoint
for the get machine operation typically these are written to a http.Request
*/
type GetMachineParams struct {

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

// WithTimeout adds the timeout to the get machine params
func (o *GetMachineParams) WithTimeout(timeout time.Duration) *GetMachineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get machine params
func (o *GetMachineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get machine params
func (o *GetMachineParams) WithContext(ctx context.Context) *GetMachineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get machine params
func (o *GetMachineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get machine params
func (o *GetMachineParams) WithHTTPClient(client *http.Client) *GetMachineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get machine params
func (o *GetMachineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get machine params
func (o *GetMachineParams) WithAPIVersion(aPIVersion *string) *GetMachineParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get machine params
func (o *GetMachineParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get machine params
func (o *GetMachineParams) WithID(id string) *GetMachineParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get machine params
func (o *GetMachineParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetMachineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

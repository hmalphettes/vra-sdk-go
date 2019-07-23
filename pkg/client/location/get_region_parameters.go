// Code generated by go-swagger; DO NOT EDIT.

package location

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

// NewGetRegionParams creates a new GetRegionParams object
// with the default values initialized.
func NewGetRegionParams() *GetRegionParams {
	var ()
	return &GetRegionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRegionParamsWithTimeout creates a new GetRegionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRegionParamsWithTimeout(timeout time.Duration) *GetRegionParams {
	var ()
	return &GetRegionParams{

		timeout: timeout,
	}
}

// NewGetRegionParamsWithContext creates a new GetRegionParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRegionParamsWithContext(ctx context.Context) *GetRegionParams {
	var ()
	return &GetRegionParams{

		Context: ctx,
	}
}

// NewGetRegionParamsWithHTTPClient creates a new GetRegionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRegionParamsWithHTTPClient(client *http.Client) *GetRegionParams {
	var ()
	return &GetRegionParams{
		HTTPClient: client,
	}
}

/*GetRegionParams contains all the parameters to send to the API endpoint
for the get region operation typically these are written to a http.Request
*/
type GetRegionParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /iaas/api/about

	*/
	APIVersion *string
	/*ID
	  The ID of the region.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get region params
func (o *GetRegionParams) WithTimeout(timeout time.Duration) *GetRegionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get region params
func (o *GetRegionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get region params
func (o *GetRegionParams) WithContext(ctx context.Context) *GetRegionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get region params
func (o *GetRegionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get region params
func (o *GetRegionParams) WithHTTPClient(client *http.Client) *GetRegionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get region params
func (o *GetRegionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get region params
func (o *GetRegionParams) WithAPIVersion(aPIVersion *string) *GetRegionParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get region params
func (o *GetRegionParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get region params
func (o *GetRegionParams) WithID(id string) *GetRegionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get region params
func (o *GetRegionParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetRegionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// Code generated by go-swagger; DO NOT EDIT.

package executions

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

// NewGetUsingGET4Params creates a new GetUsingGET4Params object
// with the default values initialized.
func NewGetUsingGET4Params() *GetUsingGET4Params {
	var ()
	return &GetUsingGET4Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUsingGET4ParamsWithTimeout creates a new GetUsingGET4Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUsingGET4ParamsWithTimeout(timeout time.Duration) *GetUsingGET4Params {
	var ()
	return &GetUsingGET4Params{

		timeout: timeout,
	}
}

// NewGetUsingGET4ParamsWithContext creates a new GetUsingGET4Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetUsingGET4ParamsWithContext(ctx context.Context) *GetUsingGET4Params {
	var ()
	return &GetUsingGET4Params{

		Context: ctx,
	}
}

// NewGetUsingGET4ParamsWithHTTPClient creates a new GetUsingGET4Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUsingGET4ParamsWithHTTPClient(client *http.Client) *GetUsingGET4Params {
	var ()
	return &GetUsingGET4Params{
		HTTPClient: client,
	}
}

/*GetUsingGET4Params contains all the parameters to send to the API endpoint
for the get using g e t 4 operation typically these are written to a http.Request
*/
type GetUsingGET4Params struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about

	*/
	APIVersion string
	/*ID
	  The ID of the User Operation

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get using g e t 4 params
func (o *GetUsingGET4Params) WithTimeout(timeout time.Duration) *GetUsingGET4Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get using g e t 4 params
func (o *GetUsingGET4Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get using g e t 4 params
func (o *GetUsingGET4Params) WithContext(ctx context.Context) *GetUsingGET4Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get using g e t 4 params
func (o *GetUsingGET4Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get using g e t 4 params
func (o *GetUsingGET4Params) WithHTTPClient(client *http.Client) *GetUsingGET4Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get using g e t 4 params
func (o *GetUsingGET4Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get using g e t 4 params
func (o *GetUsingGET4Params) WithAPIVersion(aPIVersion string) *GetUsingGET4Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get using g e t 4 params
func (o *GetUsingGET4Params) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get using g e t 4 params
func (o *GetUsingGET4Params) WithID(id string) *GetUsingGET4Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the get using g e t 4 params
func (o *GetUsingGET4Params) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetUsingGET4Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param apiVersion
	qrAPIVersion := o.APIVersion
	qAPIVersion := qrAPIVersion
	if err := r.SetQueryParam("apiVersion", qAPIVersion); err != nil {
		return err
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
// Code generated by go-swagger; DO NOT EDIT.

package pipelines

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

// NewGetUsingGET6Params creates a new GetUsingGET6Params object
// with the default values initialized.
func NewGetUsingGET6Params() *GetUsingGET6Params {
	var ()
	return &GetUsingGET6Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUsingGET6ParamsWithTimeout creates a new GetUsingGET6Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUsingGET6ParamsWithTimeout(timeout time.Duration) *GetUsingGET6Params {
	var ()
	return &GetUsingGET6Params{

		timeout: timeout,
	}
}

// NewGetUsingGET6ParamsWithContext creates a new GetUsingGET6Params object
// with the default values initialized, and the ability to set a context for a request
func NewGetUsingGET6ParamsWithContext(ctx context.Context) *GetUsingGET6Params {
	var ()
	return &GetUsingGET6Params{

		Context: ctx,
	}
}

// NewGetUsingGET6ParamsWithHTTPClient creates a new GetUsingGET6Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUsingGET6ParamsWithHTTPClient(client *http.Client) *GetUsingGET6Params {
	var ()
	return &GetUsingGET6Params{
		HTTPClient: client,
	}
}

/*GetUsingGET6Params contains all the parameters to send to the API endpoint
for the get using g e t 6 operation typically these are written to a http.Request
*/
type GetUsingGET6Params struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about

	*/
	APIVersion string
	/*Name
	  The name of the Variable

	*/
	Name string
	/*Project
	  The project the Variable belongs to

	*/
	Project string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get using g e t 6 params
func (o *GetUsingGET6Params) WithTimeout(timeout time.Duration) *GetUsingGET6Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get using g e t 6 params
func (o *GetUsingGET6Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get using g e t 6 params
func (o *GetUsingGET6Params) WithContext(ctx context.Context) *GetUsingGET6Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get using g e t 6 params
func (o *GetUsingGET6Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get using g e t 6 params
func (o *GetUsingGET6Params) WithHTTPClient(client *http.Client) *GetUsingGET6Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get using g e t 6 params
func (o *GetUsingGET6Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get using g e t 6 params
func (o *GetUsingGET6Params) WithAPIVersion(aPIVersion string) *GetUsingGET6Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get using g e t 6 params
func (o *GetUsingGET6Params) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithName adds the name to the get using g e t 6 params
func (o *GetUsingGET6Params) WithName(name string) *GetUsingGET6Params {
	o.SetName(name)
	return o
}

// SetName adds the name to the get using g e t 6 params
func (o *GetUsingGET6Params) SetName(name string) {
	o.Name = name
}

// WithProject adds the project to the get using g e t 6 params
func (o *GetUsingGET6Params) WithProject(project string) *GetUsingGET6Params {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the get using g e t 6 params
func (o *GetUsingGET6Params) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *GetUsingGET6Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param project
	if err := r.SetPathParam("project", o.Project); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
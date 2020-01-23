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

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// NewUpdateUsingPUT3Params creates a new UpdateUsingPUT3Params object
// with the default values initialized.
func NewUpdateUsingPUT3Params() *UpdateUsingPUT3Params {
	var ()
	return &UpdateUsingPUT3Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateUsingPUT3ParamsWithTimeout creates a new UpdateUsingPUT3Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateUsingPUT3ParamsWithTimeout(timeout time.Duration) *UpdateUsingPUT3Params {
	var ()
	return &UpdateUsingPUT3Params{

		timeout: timeout,
	}
}

// NewUpdateUsingPUT3ParamsWithContext creates a new UpdateUsingPUT3Params object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateUsingPUT3ParamsWithContext(ctx context.Context) *UpdateUsingPUT3Params {
	var ()
	return &UpdateUsingPUT3Params{

		Context: ctx,
	}
}

// NewUpdateUsingPUT3ParamsWithHTTPClient creates a new UpdateUsingPUT3Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateUsingPUT3ParamsWithHTTPClient(client *http.Client) *UpdateUsingPUT3Params {
	var ()
	return &UpdateUsingPUT3Params{
		HTTPClient: client,
	}
}

/*UpdateUsingPUT3Params contains all the parameters to send to the API endpoint
for the update using p u t 3 operation typically these are written to a http.Request
*/
type UpdateUsingPUT3Params struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about

	*/
	APIVersion string
	/*Body
	  Pipeline specification

	*/
	Body models.PipelineSpec
	/*Name
	  The name of the Pipeline

	*/
	Name string
	/*Project
	  The project the Pipeline belongs to

	*/
	Project string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update using p u t 3 params
func (o *UpdateUsingPUT3Params) WithTimeout(timeout time.Duration) *UpdateUsingPUT3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update using p u t 3 params
func (o *UpdateUsingPUT3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update using p u t 3 params
func (o *UpdateUsingPUT3Params) WithContext(ctx context.Context) *UpdateUsingPUT3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update using p u t 3 params
func (o *UpdateUsingPUT3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update using p u t 3 params
func (o *UpdateUsingPUT3Params) WithHTTPClient(client *http.Client) *UpdateUsingPUT3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update using p u t 3 params
func (o *UpdateUsingPUT3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the update using p u t 3 params
func (o *UpdateUsingPUT3Params) WithAPIVersion(aPIVersion string) *UpdateUsingPUT3Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the update using p u t 3 params
func (o *UpdateUsingPUT3Params) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the update using p u t 3 params
func (o *UpdateUsingPUT3Params) WithBody(body models.PipelineSpec) *UpdateUsingPUT3Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update using p u t 3 params
func (o *UpdateUsingPUT3Params) SetBody(body models.PipelineSpec) {
	o.Body = body
}

// WithName adds the name to the update using p u t 3 params
func (o *UpdateUsingPUT3Params) WithName(name string) *UpdateUsingPUT3Params {
	o.SetName(name)
	return o
}

// SetName adds the name to the update using p u t 3 params
func (o *UpdateUsingPUT3Params) SetName(name string) {
	o.Name = name
}

// WithProject adds the project to the update using p u t 3 params
func (o *UpdateUsingPUT3Params) WithProject(project string) *UpdateUsingPUT3Params {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the update using p u t 3 params
func (o *UpdateUsingPUT3Params) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateUsingPUT3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if err := r.SetBodyParam(o.Body); err != nil {
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
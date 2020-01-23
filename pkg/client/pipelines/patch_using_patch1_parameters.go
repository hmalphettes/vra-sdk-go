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

// NewPatchUsingPATCH1Params creates a new PatchUsingPATCH1Params object
// with the default values initialized.
func NewPatchUsingPATCH1Params() *PatchUsingPATCH1Params {
	var ()
	return &PatchUsingPATCH1Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchUsingPATCH1ParamsWithTimeout creates a new PatchUsingPATCH1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchUsingPATCH1ParamsWithTimeout(timeout time.Duration) *PatchUsingPATCH1Params {
	var ()
	return &PatchUsingPATCH1Params{

		timeout: timeout,
	}
}

// NewPatchUsingPATCH1ParamsWithContext creates a new PatchUsingPATCH1Params object
// with the default values initialized, and the ability to set a context for a request
func NewPatchUsingPATCH1ParamsWithContext(ctx context.Context) *PatchUsingPATCH1Params {
	var ()
	return &PatchUsingPATCH1Params{

		Context: ctx,
	}
}

// NewPatchUsingPATCH1ParamsWithHTTPClient creates a new PatchUsingPATCH1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchUsingPATCH1ParamsWithHTTPClient(client *http.Client) *PatchUsingPATCH1Params {
	var ()
	return &PatchUsingPATCH1Params{
		HTTPClient: client,
	}
}

/*PatchUsingPATCH1Params contains all the parameters to send to the API endpoint
for the patch using p a t c h 1 operation typically these are written to a http.Request
*/
type PatchUsingPATCH1Params struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about

	*/
	APIVersion string
	/*Body
	  Patch Request for a pipeline

	*/
	Body models.PipelinePatchRequest
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

// WithTimeout adds the timeout to the patch using p a t c h 1 params
func (o *PatchUsingPATCH1Params) WithTimeout(timeout time.Duration) *PatchUsingPATCH1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch using p a t c h 1 params
func (o *PatchUsingPATCH1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch using p a t c h 1 params
func (o *PatchUsingPATCH1Params) WithContext(ctx context.Context) *PatchUsingPATCH1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch using p a t c h 1 params
func (o *PatchUsingPATCH1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch using p a t c h 1 params
func (o *PatchUsingPATCH1Params) WithHTTPClient(client *http.Client) *PatchUsingPATCH1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch using p a t c h 1 params
func (o *PatchUsingPATCH1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the patch using p a t c h 1 params
func (o *PatchUsingPATCH1Params) WithAPIVersion(aPIVersion string) *PatchUsingPATCH1Params {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the patch using p a t c h 1 params
func (o *PatchUsingPATCH1Params) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the patch using p a t c h 1 params
func (o *PatchUsingPATCH1Params) WithBody(body models.PipelinePatchRequest) *PatchUsingPATCH1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch using p a t c h 1 params
func (o *PatchUsingPATCH1Params) SetBody(body models.PipelinePatchRequest) {
	o.Body = body
}

// WithName adds the name to the patch using p a t c h 1 params
func (o *PatchUsingPATCH1Params) WithName(name string) *PatchUsingPATCH1Params {
	o.SetName(name)
	return o
}

// SetName adds the name to the patch using p a t c h 1 params
func (o *PatchUsingPATCH1Params) SetName(name string) {
	o.Name = name
}

// WithProject adds the project to the patch using p a t c h 1 params
func (o *PatchUsingPATCH1Params) WithProject(project string) *PatchUsingPATCH1Params {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the patch using p a t c h 1 params
func (o *PatchUsingPATCH1Params) SetProject(project string) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *PatchUsingPATCH1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
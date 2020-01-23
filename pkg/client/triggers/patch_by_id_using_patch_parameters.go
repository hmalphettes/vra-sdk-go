// Code generated by go-swagger; DO NOT EDIT.

package triggers

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

// NewPatchByIDUsingPATCHParams creates a new PatchByIDUsingPATCHParams object
// with the default values initialized.
func NewPatchByIDUsingPATCHParams() *PatchByIDUsingPATCHParams {
	var ()
	return &PatchByIDUsingPATCHParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchByIDUsingPATCHParamsWithTimeout creates a new PatchByIDUsingPATCHParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchByIDUsingPATCHParamsWithTimeout(timeout time.Duration) *PatchByIDUsingPATCHParams {
	var ()
	return &PatchByIDUsingPATCHParams{

		timeout: timeout,
	}
}

// NewPatchByIDUsingPATCHParamsWithContext creates a new PatchByIDUsingPATCHParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchByIDUsingPATCHParamsWithContext(ctx context.Context) *PatchByIDUsingPATCHParams {
	var ()
	return &PatchByIDUsingPATCHParams{

		Context: ctx,
	}
}

// NewPatchByIDUsingPATCHParamsWithHTTPClient creates a new PatchByIDUsingPATCHParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchByIDUsingPATCHParamsWithHTTPClient(client *http.Client) *PatchByIDUsingPATCHParams {
	var ()
	return &PatchByIDUsingPATCHParams{
		HTTPClient: client,
	}
}

/*PatchByIDUsingPATCHParams contains all the parameters to send to the API endpoint
for the patch by Id using p a t c h operation typically these are written to a http.Request
*/
type PatchByIDUsingPATCHParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about

	*/
	APIVersion string
	/*DockerRegistryWebhookPatch
	  dockerRegistryWebhookPatch

	*/
	DockerRegistryWebhookPatch models.DockerRegistryWebhookPatch
	/*ID
	  id

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch by Id using p a t c h params
func (o *PatchByIDUsingPATCHParams) WithTimeout(timeout time.Duration) *PatchByIDUsingPATCHParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch by Id using p a t c h params
func (o *PatchByIDUsingPATCHParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch by Id using p a t c h params
func (o *PatchByIDUsingPATCHParams) WithContext(ctx context.Context) *PatchByIDUsingPATCHParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch by Id using p a t c h params
func (o *PatchByIDUsingPATCHParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch by Id using p a t c h params
func (o *PatchByIDUsingPATCHParams) WithHTTPClient(client *http.Client) *PatchByIDUsingPATCHParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch by Id using p a t c h params
func (o *PatchByIDUsingPATCHParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the patch by Id using p a t c h params
func (o *PatchByIDUsingPATCHParams) WithAPIVersion(aPIVersion string) *PatchByIDUsingPATCHParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the patch by Id using p a t c h params
func (o *PatchByIDUsingPATCHParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithDockerRegistryWebhookPatch adds the dockerRegistryWebhookPatch to the patch by Id using p a t c h params
func (o *PatchByIDUsingPATCHParams) WithDockerRegistryWebhookPatch(dockerRegistryWebhookPatch models.DockerRegistryWebhookPatch) *PatchByIDUsingPATCHParams {
	o.SetDockerRegistryWebhookPatch(dockerRegistryWebhookPatch)
	return o
}

// SetDockerRegistryWebhookPatch adds the dockerRegistryWebhookPatch to the patch by Id using p a t c h params
func (o *PatchByIDUsingPATCHParams) SetDockerRegistryWebhookPatch(dockerRegistryWebhookPatch models.DockerRegistryWebhookPatch) {
	o.DockerRegistryWebhookPatch = dockerRegistryWebhookPatch
}

// WithID adds the id to the patch by Id using p a t c h params
func (o *PatchByIDUsingPATCHParams) WithID(id string) *PatchByIDUsingPATCHParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch by Id using p a t c h params
func (o *PatchByIDUsingPATCHParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchByIDUsingPATCHParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if err := r.SetBodyParam(o.DockerRegistryWebhookPatch); err != nil {
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
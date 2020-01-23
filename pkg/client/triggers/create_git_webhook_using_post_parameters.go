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

// NewCreateGitWebhookUsingPOSTParams creates a new CreateGitWebhookUsingPOSTParams object
// with the default values initialized.
func NewCreateGitWebhookUsingPOSTParams() *CreateGitWebhookUsingPOSTParams {
	var ()
	return &CreateGitWebhookUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateGitWebhookUsingPOSTParamsWithTimeout creates a new CreateGitWebhookUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateGitWebhookUsingPOSTParamsWithTimeout(timeout time.Duration) *CreateGitWebhookUsingPOSTParams {
	var ()
	return &CreateGitWebhookUsingPOSTParams{

		timeout: timeout,
	}
}

// NewCreateGitWebhookUsingPOSTParamsWithContext creates a new CreateGitWebhookUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateGitWebhookUsingPOSTParamsWithContext(ctx context.Context) *CreateGitWebhookUsingPOSTParams {
	var ()
	return &CreateGitWebhookUsingPOSTParams{

		Context: ctx,
	}
}

// NewCreateGitWebhookUsingPOSTParamsWithHTTPClient creates a new CreateGitWebhookUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateGitWebhookUsingPOSTParamsWithHTTPClient(client *http.Client) *CreateGitWebhookUsingPOSTParams {
	var ()
	return &CreateGitWebhookUsingPOSTParams{
		HTTPClient: client,
	}
}

/*CreateGitWebhookUsingPOSTParams contains all the parameters to send to the API endpoint
for the create git webhook using p o s t operation typically these are written to a http.Request
*/
type CreateGitWebhookUsingPOSTParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about

	*/
	APIVersion string
	/*GitWebhookSpec
	  gitWebhookSpec

	*/
	GitWebhookSpec models.GitWebhookSpec

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create git webhook using p o s t params
func (o *CreateGitWebhookUsingPOSTParams) WithTimeout(timeout time.Duration) *CreateGitWebhookUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create git webhook using p o s t params
func (o *CreateGitWebhookUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create git webhook using p o s t params
func (o *CreateGitWebhookUsingPOSTParams) WithContext(ctx context.Context) *CreateGitWebhookUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create git webhook using p o s t params
func (o *CreateGitWebhookUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create git webhook using p o s t params
func (o *CreateGitWebhookUsingPOSTParams) WithHTTPClient(client *http.Client) *CreateGitWebhookUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create git webhook using p o s t params
func (o *CreateGitWebhookUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the create git webhook using p o s t params
func (o *CreateGitWebhookUsingPOSTParams) WithAPIVersion(aPIVersion string) *CreateGitWebhookUsingPOSTParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the create git webhook using p o s t params
func (o *CreateGitWebhookUsingPOSTParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithGitWebhookSpec adds the gitWebhookSpec to the create git webhook using p o s t params
func (o *CreateGitWebhookUsingPOSTParams) WithGitWebhookSpec(gitWebhookSpec models.GitWebhookSpec) *CreateGitWebhookUsingPOSTParams {
	o.SetGitWebhookSpec(gitWebhookSpec)
	return o
}

// SetGitWebhookSpec adds the gitWebhookSpec to the create git webhook using p o s t params
func (o *CreateGitWebhookUsingPOSTParams) SetGitWebhookSpec(gitWebhookSpec models.GitWebhookSpec) {
	o.GitWebhookSpec = gitWebhookSpec
}

// WriteToRequest writes these params to a swagger request
func (o *CreateGitWebhookUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if err := r.SetBodyParam(o.GitWebhookSpec); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
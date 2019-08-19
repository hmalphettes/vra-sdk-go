// Code generated by go-swagger; DO NOT EDIT.

package cloud_account

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

// NewCreateGcpCloudAccountParams creates a new CreateGcpCloudAccountParams object
// with the default values initialized.
func NewCreateGcpCloudAccountParams() *CreateGcpCloudAccountParams {
	var ()
	return &CreateGcpCloudAccountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateGcpCloudAccountParamsWithTimeout creates a new CreateGcpCloudAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateGcpCloudAccountParamsWithTimeout(timeout time.Duration) *CreateGcpCloudAccountParams {
	var ()
	return &CreateGcpCloudAccountParams{

		timeout: timeout,
	}
}

// NewCreateGcpCloudAccountParamsWithContext creates a new CreateGcpCloudAccountParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateGcpCloudAccountParamsWithContext(ctx context.Context) *CreateGcpCloudAccountParams {
	var ()
	return &CreateGcpCloudAccountParams{

		Context: ctx,
	}
}

// NewCreateGcpCloudAccountParamsWithHTTPClient creates a new CreateGcpCloudAccountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateGcpCloudAccountParamsWithHTTPClient(client *http.Client) *CreateGcpCloudAccountParams {
	var ()
	return &CreateGcpCloudAccountParams{
		HTTPClient: client,
	}
}

/*CreateGcpCloudAccountParams contains all the parameters to send to the API endpoint
for the create gcp cloud account operation typically these are written to a http.Request
*/
type CreateGcpCloudAccountParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /iaas/api/about

	*/
	APIVersion *string
	/*Body
	  CloudAccountGcp specification

	*/
	Body *models.CloudAccountGcpSpecification

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create gcp cloud account params
func (o *CreateGcpCloudAccountParams) WithTimeout(timeout time.Duration) *CreateGcpCloudAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create gcp cloud account params
func (o *CreateGcpCloudAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create gcp cloud account params
func (o *CreateGcpCloudAccountParams) WithContext(ctx context.Context) *CreateGcpCloudAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create gcp cloud account params
func (o *CreateGcpCloudAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create gcp cloud account params
func (o *CreateGcpCloudAccountParams) WithHTTPClient(client *http.Client) *CreateGcpCloudAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create gcp cloud account params
func (o *CreateGcpCloudAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the create gcp cloud account params
func (o *CreateGcpCloudAccountParams) WithAPIVersion(aPIVersion *string) *CreateGcpCloudAccountParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the create gcp cloud account params
func (o *CreateGcpCloudAccountParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the create gcp cloud account params
func (o *CreateGcpCloudAccountParams) WithBody(body *models.CloudAccountGcpSpecification) *CreateGcpCloudAccountParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create gcp cloud account params
func (o *CreateGcpCloudAccountParams) SetBody(body *models.CloudAccountGcpSpecification) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateGcpCloudAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
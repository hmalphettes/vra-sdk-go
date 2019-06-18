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
)

// NewDeleteAwsCloudAccountParams creates a new DeleteAwsCloudAccountParams object
// with the default values initialized.
func NewDeleteAwsCloudAccountParams() *DeleteAwsCloudAccountParams {
	var ()
	return &DeleteAwsCloudAccountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAwsCloudAccountParamsWithTimeout creates a new DeleteAwsCloudAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAwsCloudAccountParamsWithTimeout(timeout time.Duration) *DeleteAwsCloudAccountParams {
	var ()
	return &DeleteAwsCloudAccountParams{

		timeout: timeout,
	}
}

// NewDeleteAwsCloudAccountParamsWithContext creates a new DeleteAwsCloudAccountParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAwsCloudAccountParamsWithContext(ctx context.Context) *DeleteAwsCloudAccountParams {
	var ()
	return &DeleteAwsCloudAccountParams{

		Context: ctx,
	}
}

// NewDeleteAwsCloudAccountParamsWithHTTPClient creates a new DeleteAwsCloudAccountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAwsCloudAccountParamsWithHTTPClient(client *http.Client) *DeleteAwsCloudAccountParams {
	var ()
	return &DeleteAwsCloudAccountParams{
		HTTPClient: client,
	}
}

/*DeleteAwsCloudAccountParams contains all the parameters to send to the API endpoint
for the delete aws cloud account operation typically these are written to a http.Request
*/
type DeleteAwsCloudAccountParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /iaas/api/about

	*/
	APIVersion *string
	/*ID
	  The ID of the Cloud Account

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete aws cloud account params
func (o *DeleteAwsCloudAccountParams) WithTimeout(timeout time.Duration) *DeleteAwsCloudAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete aws cloud account params
func (o *DeleteAwsCloudAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete aws cloud account params
func (o *DeleteAwsCloudAccountParams) WithContext(ctx context.Context) *DeleteAwsCloudAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete aws cloud account params
func (o *DeleteAwsCloudAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete aws cloud account params
func (o *DeleteAwsCloudAccountParams) WithHTTPClient(client *http.Client) *DeleteAwsCloudAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete aws cloud account params
func (o *DeleteAwsCloudAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the delete aws cloud account params
func (o *DeleteAwsCloudAccountParams) WithAPIVersion(aPIVersion *string) *DeleteAwsCloudAccountParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the delete aws cloud account params
func (o *DeleteAwsCloudAccountParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the delete aws cloud account params
func (o *DeleteAwsCloudAccountParams) WithID(id string) *DeleteAwsCloudAccountParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete aws cloud account params
func (o *DeleteAwsCloudAccountParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAwsCloudAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

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
)

// NewGetGitWebhookByIDUsingGETParams creates a new GetGitWebhookByIDUsingGETParams object
// with the default values initialized.
func NewGetGitWebhookByIDUsingGETParams() *GetGitWebhookByIDUsingGETParams {
	var ()
	return &GetGitWebhookByIDUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetGitWebhookByIDUsingGETParamsWithTimeout creates a new GetGitWebhookByIDUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetGitWebhookByIDUsingGETParamsWithTimeout(timeout time.Duration) *GetGitWebhookByIDUsingGETParams {
	var ()
	return &GetGitWebhookByIDUsingGETParams{

		timeout: timeout,
	}
}

// NewGetGitWebhookByIDUsingGETParamsWithContext creates a new GetGitWebhookByIDUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetGitWebhookByIDUsingGETParamsWithContext(ctx context.Context) *GetGitWebhookByIDUsingGETParams {
	var ()
	return &GetGitWebhookByIDUsingGETParams{

		Context: ctx,
	}
}

// NewGetGitWebhookByIDUsingGETParamsWithHTTPClient creates a new GetGitWebhookByIDUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetGitWebhookByIDUsingGETParamsWithHTTPClient(client *http.Client) *GetGitWebhookByIDUsingGETParams {
	var ()
	return &GetGitWebhookByIDUsingGETParams{
		HTTPClient: client,
	}
}

/*GetGitWebhookByIDUsingGETParams contains all the parameters to send to the API endpoint
for the get git webhook by ID using g e t operation typically these are written to a http.Request
*/
type GetGitWebhookByIDUsingGETParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about

	*/
	APIVersion string
	/*ID
	  id

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get git webhook by ID using get params
func (o *GetGitWebhookByIDUsingGETParams) WithTimeout(timeout time.Duration) *GetGitWebhookByIDUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get git webhook by ID using get params
func (o *GetGitWebhookByIDUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get git webhook by ID using get params
func (o *GetGitWebhookByIDUsingGETParams) WithContext(ctx context.Context) *GetGitWebhookByIDUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get git webhook by ID using get params
func (o *GetGitWebhookByIDUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get git webhook by ID using get params
func (o *GetGitWebhookByIDUsingGETParams) WithHTTPClient(client *http.Client) *GetGitWebhookByIDUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get git webhook by ID using get params
func (o *GetGitWebhookByIDUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the get git webhook by ID using get params
func (o *GetGitWebhookByIDUsingGETParams) WithAPIVersion(aPIVersion string) *GetGitWebhookByIDUsingGETParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the get git webhook by ID using get params
func (o *GetGitWebhookByIDUsingGETParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the get git webhook by ID using get params
func (o *GetGitWebhookByIDUsingGETParams) WithID(id string) *GetGitWebhookByIDUsingGETParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get git webhook by ID using get params
func (o *GetGitWebhookByIDUsingGETParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetGitWebhookByIDUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
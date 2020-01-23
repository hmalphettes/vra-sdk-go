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

// NewManualTriggerUsingPOSTParams creates a new ManualTriggerUsingPOSTParams object
// with the default values initialized.
func NewManualTriggerUsingPOSTParams() *ManualTriggerUsingPOSTParams {
	var ()
	return &ManualTriggerUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewManualTriggerUsingPOSTParamsWithTimeout creates a new ManualTriggerUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewManualTriggerUsingPOSTParamsWithTimeout(timeout time.Duration) *ManualTriggerUsingPOSTParams {
	var ()
	return &ManualTriggerUsingPOSTParams{

		timeout: timeout,
	}
}

// NewManualTriggerUsingPOSTParamsWithContext creates a new ManualTriggerUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewManualTriggerUsingPOSTParamsWithContext(ctx context.Context) *ManualTriggerUsingPOSTParams {
	var ()
	return &ManualTriggerUsingPOSTParams{

		Context: ctx,
	}
}

// NewManualTriggerUsingPOSTParamsWithHTTPClient creates a new ManualTriggerUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewManualTriggerUsingPOSTParamsWithHTTPClient(client *http.Client) *ManualTriggerUsingPOSTParams {
	var ()
	return &ManualTriggerUsingPOSTParams{
		HTTPClient: client,
	}
}

/*ManualTriggerUsingPOSTParams contains all the parameters to send to the API endpoint
for the manual trigger using p o s t operation typically these are written to a http.Request
*/
type ManualTriggerUsingPOSTParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about

	*/
	APIVersion string
	/*GerritManualTrigger
	  gerritManualTrigger

	*/
	GerritManualTrigger models.GerritManualTrigger

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the manual trigger using p o s t params
func (o *ManualTriggerUsingPOSTParams) WithTimeout(timeout time.Duration) *ManualTriggerUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the manual trigger using p o s t params
func (o *ManualTriggerUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the manual trigger using p o s t params
func (o *ManualTriggerUsingPOSTParams) WithContext(ctx context.Context) *ManualTriggerUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the manual trigger using p o s t params
func (o *ManualTriggerUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the manual trigger using p o s t params
func (o *ManualTriggerUsingPOSTParams) WithHTTPClient(client *http.Client) *ManualTriggerUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the manual trigger using p o s t params
func (o *ManualTriggerUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the manual trigger using p o s t params
func (o *ManualTriggerUsingPOSTParams) WithAPIVersion(aPIVersion string) *ManualTriggerUsingPOSTParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the manual trigger using p o s t params
func (o *ManualTriggerUsingPOSTParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithGerritManualTrigger adds the gerritManualTrigger to the manual trigger using p o s t params
func (o *ManualTriggerUsingPOSTParams) WithGerritManualTrigger(gerritManualTrigger models.GerritManualTrigger) *ManualTriggerUsingPOSTParams {
	o.SetGerritManualTrigger(gerritManualTrigger)
	return o
}

// SetGerritManualTrigger adds the gerritManualTrigger to the manual trigger using p o s t params
func (o *ManualTriggerUsingPOSTParams) SetGerritManualTrigger(gerritManualTrigger models.GerritManualTrigger) {
	o.GerritManualTrigger = gerritManualTrigger
}

// WriteToRequest writes these params to a swagger request
func (o *ManualTriggerUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if err := r.SetBodyParam(o.GerritManualTrigger); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
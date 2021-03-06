// Code generated by go-swagger; DO NOT EDIT.

package deployment_actions

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

// NewSubmitResourceActionRequestUsingPOSTParams creates a new SubmitResourceActionRequestUsingPOSTParams object
// with the default values initialized.
func NewSubmitResourceActionRequestUsingPOSTParams() *SubmitResourceActionRequestUsingPOSTParams {
	var ()
	return &SubmitResourceActionRequestUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSubmitResourceActionRequestUsingPOSTParamsWithTimeout creates a new SubmitResourceActionRequestUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSubmitResourceActionRequestUsingPOSTParamsWithTimeout(timeout time.Duration) *SubmitResourceActionRequestUsingPOSTParams {
	var ()
	return &SubmitResourceActionRequestUsingPOSTParams{

		timeout: timeout,
	}
}

// NewSubmitResourceActionRequestUsingPOSTParamsWithContext creates a new SubmitResourceActionRequestUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewSubmitResourceActionRequestUsingPOSTParamsWithContext(ctx context.Context) *SubmitResourceActionRequestUsingPOSTParams {
	var ()
	return &SubmitResourceActionRequestUsingPOSTParams{

		Context: ctx,
	}
}

// NewSubmitResourceActionRequestUsingPOSTParamsWithHTTPClient creates a new SubmitResourceActionRequestUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSubmitResourceActionRequestUsingPOSTParamsWithHTTPClient(client *http.Client) *SubmitResourceActionRequestUsingPOSTParams {
	var ()
	return &SubmitResourceActionRequestUsingPOSTParams{
		HTTPClient: client,
	}
}

/*SubmitResourceActionRequestUsingPOSTParams contains all the parameters to send to the API endpoint
for the submit resource action request using p o s t operation typically these are written to a http.Request
*/
type SubmitResourceActionRequestUsingPOSTParams struct {

	/*ActionRequest
	  actionRequest

	*/
	ActionRequest *models.ResourceActionRequest
	/*DepID
	  Deployment ID

	*/
	DepID strfmt.UUID
	/*ResourceID
	  Resource ID

	*/
	ResourceID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the submit resource action request using p o s t params
func (o *SubmitResourceActionRequestUsingPOSTParams) WithTimeout(timeout time.Duration) *SubmitResourceActionRequestUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the submit resource action request using p o s t params
func (o *SubmitResourceActionRequestUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the submit resource action request using p o s t params
func (o *SubmitResourceActionRequestUsingPOSTParams) WithContext(ctx context.Context) *SubmitResourceActionRequestUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the submit resource action request using p o s t params
func (o *SubmitResourceActionRequestUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the submit resource action request using p o s t params
func (o *SubmitResourceActionRequestUsingPOSTParams) WithHTTPClient(client *http.Client) *SubmitResourceActionRequestUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the submit resource action request using p o s t params
func (o *SubmitResourceActionRequestUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithActionRequest adds the actionRequest to the submit resource action request using p o s t params
func (o *SubmitResourceActionRequestUsingPOSTParams) WithActionRequest(actionRequest *models.ResourceActionRequest) *SubmitResourceActionRequestUsingPOSTParams {
	o.SetActionRequest(actionRequest)
	return o
}

// SetActionRequest adds the actionRequest to the submit resource action request using p o s t params
func (o *SubmitResourceActionRequestUsingPOSTParams) SetActionRequest(actionRequest *models.ResourceActionRequest) {
	o.ActionRequest = actionRequest
}

// WithDepID adds the depID to the submit resource action request using p o s t params
func (o *SubmitResourceActionRequestUsingPOSTParams) WithDepID(depID strfmt.UUID) *SubmitResourceActionRequestUsingPOSTParams {
	o.SetDepID(depID)
	return o
}

// SetDepID adds the depId to the submit resource action request using p o s t params
func (o *SubmitResourceActionRequestUsingPOSTParams) SetDepID(depID strfmt.UUID) {
	o.DepID = depID
}

// WithResourceID adds the resourceID to the submit resource action request using p o s t params
func (o *SubmitResourceActionRequestUsingPOSTParams) WithResourceID(resourceID strfmt.UUID) *SubmitResourceActionRequestUsingPOSTParams {
	o.SetResourceID(resourceID)
	return o
}

// SetResourceID adds the resourceId to the submit resource action request using p o s t params
func (o *SubmitResourceActionRequestUsingPOSTParams) SetResourceID(resourceID strfmt.UUID) {
	o.ResourceID = resourceID
}

// WriteToRequest writes these params to a swagger request
func (o *SubmitResourceActionRequestUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ActionRequest != nil {
		if err := r.SetBodyParam(o.ActionRequest); err != nil {
			return err
		}
	}

	// path param depId
	if err := r.SetPathParam("depId", o.DepID.String()); err != nil {
		return err
	}

	// path param resourceId
	if err := r.SetPathParam("resourceId", o.ResourceID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

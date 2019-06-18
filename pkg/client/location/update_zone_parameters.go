// Code generated by go-swagger; DO NOT EDIT.

package location

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

	models "github.com/vmware/cas-sdk-go/pkg/models"
)

// NewUpdateZoneParams creates a new UpdateZoneParams object
// with the default values initialized.
func NewUpdateZoneParams() *UpdateZoneParams {
	var ()
	return &UpdateZoneParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateZoneParamsWithTimeout creates a new UpdateZoneParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateZoneParamsWithTimeout(timeout time.Duration) *UpdateZoneParams {
	var ()
	return &UpdateZoneParams{

		timeout: timeout,
	}
}

// NewUpdateZoneParamsWithContext creates a new UpdateZoneParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateZoneParamsWithContext(ctx context.Context) *UpdateZoneParams {
	var ()
	return &UpdateZoneParams{

		Context: ctx,
	}
}

// NewUpdateZoneParamsWithHTTPClient creates a new UpdateZoneParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateZoneParamsWithHTTPClient(client *http.Client) *UpdateZoneParams {
	var ()
	return &UpdateZoneParams{
		HTTPClient: client,
	}
}

/*UpdateZoneParams contains all the parameters to send to the API endpoint
for the update zone operation typically these are written to a http.Request
*/
type UpdateZoneParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /iaas/api/about

	*/
	APIVersion *string
	/*Body
	  Zone specification

	*/
	Body *models.ZoneSpecification
	/*ID
	  The ID of the zone.

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update zone params
func (o *UpdateZoneParams) WithTimeout(timeout time.Duration) *UpdateZoneParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update zone params
func (o *UpdateZoneParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update zone params
func (o *UpdateZoneParams) WithContext(ctx context.Context) *UpdateZoneParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update zone params
func (o *UpdateZoneParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update zone params
func (o *UpdateZoneParams) WithHTTPClient(client *http.Client) *UpdateZoneParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update zone params
func (o *UpdateZoneParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the update zone params
func (o *UpdateZoneParams) WithAPIVersion(aPIVersion *string) *UpdateZoneParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the update zone params
func (o *UpdateZoneParams) SetAPIVersion(aPIVersion *string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the update zone params
func (o *UpdateZoneParams) WithBody(body *models.ZoneSpecification) *UpdateZoneParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update zone params
func (o *UpdateZoneParams) SetBody(body *models.ZoneSpecification) {
	o.Body = body
}

// WithID adds the id to the update zone params
func (o *UpdateZoneParams) WithID(id string) *UpdateZoneParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update zone params
func (o *UpdateZoneParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateZoneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

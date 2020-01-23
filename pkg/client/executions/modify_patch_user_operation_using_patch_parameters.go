// Code generated by go-swagger; DO NOT EDIT.

package executions

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

// NewModifyPatchUserOperationUsingPATCHParams creates a new ModifyPatchUserOperationUsingPATCHParams object
// with the default values initialized.
func NewModifyPatchUserOperationUsingPATCHParams() *ModifyPatchUserOperationUsingPATCHParams {
	var ()
	return &ModifyPatchUserOperationUsingPATCHParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewModifyPatchUserOperationUsingPATCHParamsWithTimeout creates a new ModifyPatchUserOperationUsingPATCHParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewModifyPatchUserOperationUsingPATCHParamsWithTimeout(timeout time.Duration) *ModifyPatchUserOperationUsingPATCHParams {
	var ()
	return &ModifyPatchUserOperationUsingPATCHParams{

		timeout: timeout,
	}
}

// NewModifyPatchUserOperationUsingPATCHParamsWithContext creates a new ModifyPatchUserOperationUsingPATCHParams object
// with the default values initialized, and the ability to set a context for a request
func NewModifyPatchUserOperationUsingPATCHParamsWithContext(ctx context.Context) *ModifyPatchUserOperationUsingPATCHParams {
	var ()
	return &ModifyPatchUserOperationUsingPATCHParams{

		Context: ctx,
	}
}

// NewModifyPatchUserOperationUsingPATCHParamsWithHTTPClient creates a new ModifyPatchUserOperationUsingPATCHParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewModifyPatchUserOperationUsingPATCHParamsWithHTTPClient(client *http.Client) *ModifyPatchUserOperationUsingPATCHParams {
	var ()
	return &ModifyPatchUserOperationUsingPATCHParams{
		HTTPClient: client,
	}
}

/*ModifyPatchUserOperationUsingPATCHParams contains all the parameters to send to the API endpoint
for the modify patch user operation using p a t c h operation typically these are written to a http.Request
*/
type ModifyPatchUserOperationUsingPATCHParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about

	*/
	APIVersion string
	/*Body
	  User's response to the User Operation request

	*/
	Body models.UserOpResponse
	/*ID
	  The ID of the User Operation

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the modify patch user operation using p a t c h params
func (o *ModifyPatchUserOperationUsingPATCHParams) WithTimeout(timeout time.Duration) *ModifyPatchUserOperationUsingPATCHParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the modify patch user operation using p a t c h params
func (o *ModifyPatchUserOperationUsingPATCHParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the modify patch user operation using p a t c h params
func (o *ModifyPatchUserOperationUsingPATCHParams) WithContext(ctx context.Context) *ModifyPatchUserOperationUsingPATCHParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the modify patch user operation using p a t c h params
func (o *ModifyPatchUserOperationUsingPATCHParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the modify patch user operation using p a t c h params
func (o *ModifyPatchUserOperationUsingPATCHParams) WithHTTPClient(client *http.Client) *ModifyPatchUserOperationUsingPATCHParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the modify patch user operation using p a t c h params
func (o *ModifyPatchUserOperationUsingPATCHParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the modify patch user operation using p a t c h params
func (o *ModifyPatchUserOperationUsingPATCHParams) WithAPIVersion(aPIVersion string) *ModifyPatchUserOperationUsingPATCHParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the modify patch user operation using p a t c h params
func (o *ModifyPatchUserOperationUsingPATCHParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithBody adds the body to the modify patch user operation using p a t c h params
func (o *ModifyPatchUserOperationUsingPATCHParams) WithBody(body models.UserOpResponse) *ModifyPatchUserOperationUsingPATCHParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the modify patch user operation using p a t c h params
func (o *ModifyPatchUserOperationUsingPATCHParams) SetBody(body models.UserOpResponse) {
	o.Body = body
}

// WithID adds the id to the modify patch user operation using p a t c h params
func (o *ModifyPatchUserOperationUsingPATCHParams) WithID(id string) *ModifyPatchUserOperationUsingPATCHParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the modify patch user operation using p a t c h params
func (o *ModifyPatchUserOperationUsingPATCHParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ModifyPatchUserOperationUsingPATCHParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
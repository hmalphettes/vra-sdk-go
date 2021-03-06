// Code generated by go-swagger; DO NOT EDIT.

package catalog_admin_items

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

// NewSetItemIconUsingPATCHParams creates a new SetItemIconUsingPATCHParams object
// with the default values initialized.
func NewSetItemIconUsingPATCHParams() *SetItemIconUsingPATCHParams {
	var ()
	return &SetItemIconUsingPATCHParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetItemIconUsingPATCHParamsWithTimeout creates a new SetItemIconUsingPATCHParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetItemIconUsingPATCHParamsWithTimeout(timeout time.Duration) *SetItemIconUsingPATCHParams {
	var ()
	return &SetItemIconUsingPATCHParams{

		timeout: timeout,
	}
}

// NewSetItemIconUsingPATCHParamsWithContext creates a new SetItemIconUsingPATCHParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetItemIconUsingPATCHParamsWithContext(ctx context.Context) *SetItemIconUsingPATCHParams {
	var ()
	return &SetItemIconUsingPATCHParams{

		Context: ctx,
	}
}

// NewSetItemIconUsingPATCHParamsWithHTTPClient creates a new SetItemIconUsingPATCHParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetItemIconUsingPATCHParamsWithHTTPClient(client *http.Client) *SetItemIconUsingPATCHParams {
	var ()
	return &SetItemIconUsingPATCHParams{
		HTTPClient: client,
	}
}

/*SetItemIconUsingPATCHParams contains all the parameters to send to the API endpoint
for the set item icon using p a t c h operation typically these are written to a http.Request
*/
type SetItemIconUsingPATCHParams struct {

	/*ID
	  The unique id of item to update.

	*/
	ID strfmt.UUID
	/*Patch
	  The patch that apply to the item

	*/
	Patch *models.AdminCatalogItemPatch

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set item icon using p a t c h params
func (o *SetItemIconUsingPATCHParams) WithTimeout(timeout time.Duration) *SetItemIconUsingPATCHParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set item icon using p a t c h params
func (o *SetItemIconUsingPATCHParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set item icon using p a t c h params
func (o *SetItemIconUsingPATCHParams) WithContext(ctx context.Context) *SetItemIconUsingPATCHParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set item icon using p a t c h params
func (o *SetItemIconUsingPATCHParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set item icon using p a t c h params
func (o *SetItemIconUsingPATCHParams) WithHTTPClient(client *http.Client) *SetItemIconUsingPATCHParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set item icon using p a t c h params
func (o *SetItemIconUsingPATCHParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the set item icon using p a t c h params
func (o *SetItemIconUsingPATCHParams) WithID(id strfmt.UUID) *SetItemIconUsingPATCHParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the set item icon using p a t c h params
func (o *SetItemIconUsingPATCHParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithPatch adds the patch to the set item icon using p a t c h params
func (o *SetItemIconUsingPATCHParams) WithPatch(patch *models.AdminCatalogItemPatch) *SetItemIconUsingPATCHParams {
	o.SetPatch(patch)
	return o
}

// SetPatch adds the patch to the set item icon using p a t c h params
func (o *SetItemIconUsingPATCHParams) SetPatch(patch *models.AdminCatalogItemPatch) {
	o.Patch = patch
}

// WriteToRequest writes these params to a swagger request
func (o *SetItemIconUsingPATCHParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if o.Patch != nil {
		if err := r.SetBodyParam(o.Patch); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

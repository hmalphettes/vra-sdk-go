// Code generated by go-swagger; DO NOT EDIT.

package catalog_items

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetVersionsUsingGETParams creates a new GetVersionsUsingGETParams object
// with the default values initialized.
func NewGetVersionsUsingGETParams() *GetVersionsUsingGETParams {
	var ()
	return &GetVersionsUsingGETParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVersionsUsingGETParamsWithTimeout creates a new GetVersionsUsingGETParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVersionsUsingGETParamsWithTimeout(timeout time.Duration) *GetVersionsUsingGETParams {
	var ()
	return &GetVersionsUsingGETParams{

		timeout: timeout,
	}
}

// NewGetVersionsUsingGETParamsWithContext creates a new GetVersionsUsingGETParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVersionsUsingGETParamsWithContext(ctx context.Context) *GetVersionsUsingGETParams {
	var ()
	return &GetVersionsUsingGETParams{

		Context: ctx,
	}
}

// NewGetVersionsUsingGETParamsWithHTTPClient creates a new GetVersionsUsingGETParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVersionsUsingGETParamsWithHTTPClient(client *http.Client) *GetVersionsUsingGETParams {
	var ()
	return &GetVersionsUsingGETParams{
		HTTPClient: client,
	}
}

/*GetVersionsUsingGETParams contains all the parameters to send to the API endpoint
for the get versions using g e t operation typically these are written to a http.Request
*/
type GetVersionsUsingGETParams struct {

	/*ID
	  Catalog Item ID

	*/
	ID strfmt.UUID
	/*Page
	  Results page you want to retrieve (0..N)

	*/
	Page *int32
	/*Size
	  Number of records per page.

	*/
	Size *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get versions using get params
func (o *GetVersionsUsingGETParams) WithTimeout(timeout time.Duration) *GetVersionsUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get versions using get params
func (o *GetVersionsUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get versions using get params
func (o *GetVersionsUsingGETParams) WithContext(ctx context.Context) *GetVersionsUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get versions using get params
func (o *GetVersionsUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get versions using get params
func (o *GetVersionsUsingGETParams) WithHTTPClient(client *http.Client) *GetVersionsUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get versions using get params
func (o *GetVersionsUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get versions using get params
func (o *GetVersionsUsingGETParams) WithID(id strfmt.UUID) *GetVersionsUsingGETParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get versions using get params
func (o *GetVersionsUsingGETParams) SetID(id strfmt.UUID) {
	o.ID = id
}

// WithPage adds the page to the get versions using get params
func (o *GetVersionsUsingGETParams) WithPage(page *int32) *GetVersionsUsingGETParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get versions using get params
func (o *GetVersionsUsingGETParams) SetPage(page *int32) {
	o.Page = page
}

// WithSize adds the size to the get versions using get params
func (o *GetVersionsUsingGETParams) WithSize(size *int32) *GetVersionsUsingGETParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get versions using get params
func (o *GetVersionsUsingGETParams) SetSize(size *int32) {
	o.Size = size
}

// WriteToRequest writes these params to a swagger request
func (o *GetVersionsUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID.String()); err != nil {
		return err
	}

	if o.Page != nil {

		// query param page
		var qrPage int32
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if o.Size != nil {

		// query param size
		var qrSize int32
		if o.Size != nil {
			qrSize = *o.Size
		}
		qSize := swag.FormatInt32(qrSize)
		if qSize != "" {
			if err := r.SetQueryParam("size", qSize); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

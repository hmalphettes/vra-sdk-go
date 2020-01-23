// Code generated by go-swagger; DO NOT EDIT.

package custom_integrations

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

// NewWithdrawByIDAndVersionUsingPOSTParams creates a new WithdrawByIDAndVersionUsingPOSTParams object
// with the default values initialized.
func NewWithdrawByIDAndVersionUsingPOSTParams() *WithdrawByIDAndVersionUsingPOSTParams {
	var ()
	return &WithdrawByIDAndVersionUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWithdrawByIDAndVersionUsingPOSTParamsWithTimeout creates a new WithdrawByIDAndVersionUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWithdrawByIDAndVersionUsingPOSTParamsWithTimeout(timeout time.Duration) *WithdrawByIDAndVersionUsingPOSTParams {
	var ()
	return &WithdrawByIDAndVersionUsingPOSTParams{

		timeout: timeout,
	}
}

// NewWithdrawByIDAndVersionUsingPOSTParamsWithContext creates a new WithdrawByIDAndVersionUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewWithdrawByIDAndVersionUsingPOSTParamsWithContext(ctx context.Context) *WithdrawByIDAndVersionUsingPOSTParams {
	var ()
	return &WithdrawByIDAndVersionUsingPOSTParams{

		Context: ctx,
	}
}

// NewWithdrawByIDAndVersionUsingPOSTParamsWithHTTPClient creates a new WithdrawByIDAndVersionUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWithdrawByIDAndVersionUsingPOSTParamsWithHTTPClient(client *http.Client) *WithdrawByIDAndVersionUsingPOSTParams {
	var ()
	return &WithdrawByIDAndVersionUsingPOSTParams{
		HTTPClient: client,
	}
}

/*WithdrawByIDAndVersionUsingPOSTParams contains all the parameters to send to the API endpoint
for the withdraw by Id and version using p o s t operation typically these are written to a http.Request
*/
type WithdrawByIDAndVersionUsingPOSTParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /codestream/api/about

	*/
	APIVersion string
	/*ID
	  The ID of the Custom Integration

	*/
	ID string
	/*Version
	  The version of the Custom Integration

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the withdraw by Id and version using p o s t params
func (o *WithdrawByIDAndVersionUsingPOSTParams) WithTimeout(timeout time.Duration) *WithdrawByIDAndVersionUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the withdraw by Id and version using p o s t params
func (o *WithdrawByIDAndVersionUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the withdraw by Id and version using p o s t params
func (o *WithdrawByIDAndVersionUsingPOSTParams) WithContext(ctx context.Context) *WithdrawByIDAndVersionUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the withdraw by Id and version using p o s t params
func (o *WithdrawByIDAndVersionUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the withdraw by Id and version using p o s t params
func (o *WithdrawByIDAndVersionUsingPOSTParams) WithHTTPClient(client *http.Client) *WithdrawByIDAndVersionUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the withdraw by Id and version using p o s t params
func (o *WithdrawByIDAndVersionUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the withdraw by Id and version using p o s t params
func (o *WithdrawByIDAndVersionUsingPOSTParams) WithAPIVersion(aPIVersion string) *WithdrawByIDAndVersionUsingPOSTParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the withdraw by Id and version using p o s t params
func (o *WithdrawByIDAndVersionUsingPOSTParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the withdraw by Id and version using p o s t params
func (o *WithdrawByIDAndVersionUsingPOSTParams) WithID(id string) *WithdrawByIDAndVersionUsingPOSTParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the withdraw by Id and version using p o s t params
func (o *WithdrawByIDAndVersionUsingPOSTParams) SetID(id string) {
	o.ID = id
}

// WithVersion adds the version to the withdraw by Id and version using p o s t params
func (o *WithdrawByIDAndVersionUsingPOSTParams) WithVersion(version string) *WithdrawByIDAndVersionUsingPOSTParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the withdraw by Id and version using p o s t params
func (o *WithdrawByIDAndVersionUsingPOSTParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *WithdrawByIDAndVersionUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param version
	if err := r.SetPathParam("version", o.Version); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
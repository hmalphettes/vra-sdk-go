// Code generated by go-swagger; DO NOT EDIT.

package compute

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

// NewRevertMachineSnapshotParams creates a new RevertMachineSnapshotParams object
// with the default values initialized.
func NewRevertMachineSnapshotParams() *RevertMachineSnapshotParams {
	var ()
	return &RevertMachineSnapshotParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRevertMachineSnapshotParamsWithTimeout creates a new RevertMachineSnapshotParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRevertMachineSnapshotParamsWithTimeout(timeout time.Duration) *RevertMachineSnapshotParams {
	var ()
	return &RevertMachineSnapshotParams{

		timeout: timeout,
	}
}

// NewRevertMachineSnapshotParamsWithContext creates a new RevertMachineSnapshotParams object
// with the default values initialized, and the ability to set a context for a request
func NewRevertMachineSnapshotParamsWithContext(ctx context.Context) *RevertMachineSnapshotParams {
	var ()
	return &RevertMachineSnapshotParams{

		Context: ctx,
	}
}

// NewRevertMachineSnapshotParamsWithHTTPClient creates a new RevertMachineSnapshotParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRevertMachineSnapshotParamsWithHTTPClient(client *http.Client) *RevertMachineSnapshotParams {
	var ()
	return &RevertMachineSnapshotParams{
		HTTPClient: client,
	}
}

/*RevertMachineSnapshotParams contains all the parameters to send to the API endpoint
for the revert machine snapshot operation typically these are written to a http.Request
*/
type RevertMachineSnapshotParams struct {

	/*APIVersion
	  The version of the API in yyyy-MM-dd format (UTC). For versioning information please refer to /iaas/api/about

	*/
	APIVersion string
	/*ID
	  The id of the Machine.

	*/
	ID string
	/*ID
	  Snapshot id to revert.

	*/
	QueryID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the revert machine snapshot params
func (o *RevertMachineSnapshotParams) WithTimeout(timeout time.Duration) *RevertMachineSnapshotParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the revert machine snapshot params
func (o *RevertMachineSnapshotParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the revert machine snapshot params
func (o *RevertMachineSnapshotParams) WithContext(ctx context.Context) *RevertMachineSnapshotParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the revert machine snapshot params
func (o *RevertMachineSnapshotParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the revert machine snapshot params
func (o *RevertMachineSnapshotParams) WithHTTPClient(client *http.Client) *RevertMachineSnapshotParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the revert machine snapshot params
func (o *RevertMachineSnapshotParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPIVersion adds the aPIVersion to the revert machine snapshot params
func (o *RevertMachineSnapshotParams) WithAPIVersion(aPIVersion string) *RevertMachineSnapshotParams {
	o.SetAPIVersion(aPIVersion)
	return o
}

// SetAPIVersion adds the apiVersion to the revert machine snapshot params
func (o *RevertMachineSnapshotParams) SetAPIVersion(aPIVersion string) {
	o.APIVersion = aPIVersion
}

// WithID adds the id to the revert machine snapshot params
func (o *RevertMachineSnapshotParams) WithID(id string) *RevertMachineSnapshotParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the revert machine snapshot params
func (o *RevertMachineSnapshotParams) SetID(id string) {
	o.ID = id
}

// WithQueryID adds the id to the revert machine snapshot params
func (o *RevertMachineSnapshotParams) WithQueryID(id string) *RevertMachineSnapshotParams {
	o.SetQueryID(id)
	return o
}

// SetQueryID adds the id to the revert machine snapshot params
func (o *RevertMachineSnapshotParams) SetQueryID(id string) {
	o.QueryID = id
}

// WriteToRequest writes these params to a swagger request
func (o *RevertMachineSnapshotParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param apiVersion
	qrAPIVersion := o.APIVersion
	qAPIVersion := qrAPIVersion
	if qAPIVersion != "" {
		if err := r.SetQueryParam("apiVersion", qAPIVersion); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// query param id
	qrID := o.QueryID
	qID := qrID
	if qID != "" {
		if err := r.SetQueryParam("id", qID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

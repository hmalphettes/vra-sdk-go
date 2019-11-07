// Code generated by go-swagger; DO NOT EDIT.

package policies

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

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// NewCreatePolicyUsingPOSTParams creates a new CreatePolicyUsingPOSTParams object
// with the default values initialized.
func NewCreatePolicyUsingPOSTParams() *CreatePolicyUsingPOSTParams {
	var ()
	return &CreatePolicyUsingPOSTParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePolicyUsingPOSTParamsWithTimeout creates a new CreatePolicyUsingPOSTParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreatePolicyUsingPOSTParamsWithTimeout(timeout time.Duration) *CreatePolicyUsingPOSTParams {
	var ()
	return &CreatePolicyUsingPOSTParams{

		timeout: timeout,
	}
}

// NewCreatePolicyUsingPOSTParamsWithContext creates a new CreatePolicyUsingPOSTParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreatePolicyUsingPOSTParamsWithContext(ctx context.Context) *CreatePolicyUsingPOSTParams {
	var ()
	return &CreatePolicyUsingPOSTParams{

		Context: ctx,
	}
}

// NewCreatePolicyUsingPOSTParamsWithHTTPClient creates a new CreatePolicyUsingPOSTParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreatePolicyUsingPOSTParamsWithHTTPClient(client *http.Client) *CreatePolicyUsingPOSTParams {
	var ()
	return &CreatePolicyUsingPOSTParams{
		HTTPClient: client,
	}
}

/*CreatePolicyUsingPOSTParams contains all the parameters to send to the API endpoint
for the create policy using p o s t operation typically these are written to a http.Request
*/
type CreatePolicyUsingPOSTParams struct {

	/*Policy
	  The policy to be created

	*/
	Policy *models.Policy
	/*ValidationOnly
	  For a dry run that will do policy validation only instead of create a policy

	*/
	ValidationOnly *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create policy using p o s t params
func (o *CreatePolicyUsingPOSTParams) WithTimeout(timeout time.Duration) *CreatePolicyUsingPOSTParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create policy using p o s t params
func (o *CreatePolicyUsingPOSTParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create policy using p o s t params
func (o *CreatePolicyUsingPOSTParams) WithContext(ctx context.Context) *CreatePolicyUsingPOSTParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create policy using p o s t params
func (o *CreatePolicyUsingPOSTParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create policy using p o s t params
func (o *CreatePolicyUsingPOSTParams) WithHTTPClient(client *http.Client) *CreatePolicyUsingPOSTParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create policy using p o s t params
func (o *CreatePolicyUsingPOSTParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPolicy adds the policy to the create policy using p o s t params
func (o *CreatePolicyUsingPOSTParams) WithPolicy(policy *models.Policy) *CreatePolicyUsingPOSTParams {
	o.SetPolicy(policy)
	return o
}

// SetPolicy adds the policy to the create policy using p o s t params
func (o *CreatePolicyUsingPOSTParams) SetPolicy(policy *models.Policy) {
	o.Policy = policy
}

// WithValidationOnly adds the validationOnly to the create policy using p o s t params
func (o *CreatePolicyUsingPOSTParams) WithValidationOnly(validationOnly *bool) *CreatePolicyUsingPOSTParams {
	o.SetValidationOnly(validationOnly)
	return o
}

// SetValidationOnly adds the validationOnly to the create policy using p o s t params
func (o *CreatePolicyUsingPOSTParams) SetValidationOnly(validationOnly *bool) {
	o.ValidationOnly = validationOnly
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePolicyUsingPOSTParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Policy != nil {
		if err := r.SetBodyParam(o.Policy); err != nil {
			return err
		}
	}

	if o.ValidationOnly != nil {

		// query param validationOnly
		var qrValidationOnly bool
		if o.ValidationOnly != nil {
			qrValidationOnly = *o.ValidationOnly
		}
		qValidationOnly := swag.FormatBool(qrValidationOnly)
		if qValidationOnly != "" {
			if err := r.SetQueryParam("validationOnly", qValidationOnly); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
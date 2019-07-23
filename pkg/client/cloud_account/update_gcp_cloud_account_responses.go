// Code generated by go-swagger; DO NOT EDIT.

package cloud_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/cas-sdk-go/pkg/models"
)

// UpdateGcpCloudAccountReader is a Reader for the UpdateGcpCloudAccount structure.
type UpdateGcpCloudAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateGcpCloudAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateGcpCloudAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewUpdateGcpCloudAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewUpdateGcpCloudAccountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateGcpCloudAccountOK creates a UpdateGcpCloudAccountOK with default headers values
func NewUpdateGcpCloudAccountOK() *UpdateGcpCloudAccountOK {
	return &UpdateGcpCloudAccountOK{}
}

/*UpdateGcpCloudAccountOK handles this case with default header values.

successful operation
*/
type UpdateGcpCloudAccountOK struct {
	Payload *models.CloudAccountGcp
}

func (o *UpdateGcpCloudAccountOK) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/cloud-accounts-gcp/{id}][%d] updateGcpCloudAccountOK  %+v", 200, o.Payload)
}

func (o *UpdateGcpCloudAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudAccountGcp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateGcpCloudAccountForbidden creates a UpdateGcpCloudAccountForbidden with default headers values
func NewUpdateGcpCloudAccountForbidden() *UpdateGcpCloudAccountForbidden {
	return &UpdateGcpCloudAccountForbidden{}
}

/*UpdateGcpCloudAccountForbidden handles this case with default header values.

Forbidden
*/
type UpdateGcpCloudAccountForbidden struct {
}

func (o *UpdateGcpCloudAccountForbidden) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/cloud-accounts-gcp/{id}][%d] updateGcpCloudAccountForbidden ", 403)
}

func (o *UpdateGcpCloudAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateGcpCloudAccountNotFound creates a UpdateGcpCloudAccountNotFound with default headers values
func NewUpdateGcpCloudAccountNotFound() *UpdateGcpCloudAccountNotFound {
	return &UpdateGcpCloudAccountNotFound{}
}

/*UpdateGcpCloudAccountNotFound handles this case with default header values.

Not Found
*/
type UpdateGcpCloudAccountNotFound struct {
}

func (o *UpdateGcpCloudAccountNotFound) Error() string {
	return fmt.Sprintf("[PATCH /iaas/api/cloud-accounts-gcp/{id}][%d] updateGcpCloudAccountNotFound ", 404)
}

func (o *UpdateGcpCloudAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

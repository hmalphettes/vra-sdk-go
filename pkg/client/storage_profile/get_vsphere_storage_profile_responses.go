// Code generated by go-swagger; DO NOT EDIT.

package storage_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// GetVSphereStorageProfileReader is a Reader for the GetVSphereStorageProfile structure.
type GetVSphereStorageProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVSphereStorageProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVSphereStorageProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetVSphereStorageProfileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetVSphereStorageProfileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetVSphereStorageProfileOK creates a GetVSphereStorageProfileOK with default headers values
func NewGetVSphereStorageProfileOK() *GetVSphereStorageProfileOK {
	return &GetVSphereStorageProfileOK{}
}

/*GetVSphereStorageProfileOK handles this case with default header values.

successful operation
*/
type GetVSphereStorageProfileOK struct {
	Payload *models.VsphereStorageProfile
}

func (o *GetVSphereStorageProfileOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/storage-profiles-vsphere/{id}][%d] getVSphereStorageProfileOK  %+v", 200, o.Payload)
}

func (o *GetVSphereStorageProfileOK) GetPayload() *models.VsphereStorageProfile {
	return o.Payload
}

func (o *GetVSphereStorageProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VsphereStorageProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVSphereStorageProfileForbidden creates a GetVSphereStorageProfileForbidden with default headers values
func NewGetVSphereStorageProfileForbidden() *GetVSphereStorageProfileForbidden {
	return &GetVSphereStorageProfileForbidden{}
}

/*GetVSphereStorageProfileForbidden handles this case with default header values.

Forbidden
*/
type GetVSphereStorageProfileForbidden struct {
}

func (o *GetVSphereStorageProfileForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/storage-profiles-vsphere/{id}][%d] getVSphereStorageProfileForbidden ", 403)
}

func (o *GetVSphereStorageProfileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetVSphereStorageProfileNotFound creates a GetVSphereStorageProfileNotFound with default headers values
func NewGetVSphereStorageProfileNotFound() *GetVSphereStorageProfileNotFound {
	return &GetVSphereStorageProfileNotFound{}
}

/*GetVSphereStorageProfileNotFound handles this case with default header values.

Not Found
*/
type GetVSphereStorageProfileNotFound struct {
}

func (o *GetVSphereStorageProfileNotFound) Error() string {
	return fmt.Sprintf("[GET /iaas/api/storage-profiles-vsphere/{id}][%d] getVSphereStorageProfileNotFound ", 404)
}

func (o *GetVSphereStorageProfileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

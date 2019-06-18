// Code generated by go-swagger; DO NOT EDIT.

package disk

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/cas-sdk-go/pkg/models"
)

// CreateBlockDeviceReader is a Reader for the CreateBlockDevice structure.
type CreateBlockDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateBlockDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewCreateBlockDeviceAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateBlockDeviceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewCreateBlockDeviceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateBlockDeviceAccepted creates a CreateBlockDeviceAccepted with default headers values
func NewCreateBlockDeviceAccepted() *CreateBlockDeviceAccepted {
	return &CreateBlockDeviceAccepted{}
}

/*CreateBlockDeviceAccepted handles this case with default header values.

successful operation
*/
type CreateBlockDeviceAccepted struct {
	Payload *models.RequestTracker
}

func (o *CreateBlockDeviceAccepted) Error() string {
	return fmt.Sprintf("[POST /iaas/api/block-devices][%d] createBlockDeviceAccepted  %+v", 202, o.Payload)
}

func (o *CreateBlockDeviceAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestTracker)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateBlockDeviceBadRequest creates a CreateBlockDeviceBadRequest with default headers values
func NewCreateBlockDeviceBadRequest() *CreateBlockDeviceBadRequest {
	return &CreateBlockDeviceBadRequest{}
}

/*CreateBlockDeviceBadRequest handles this case with default header values.

Invalid Request - bad data
*/
type CreateBlockDeviceBadRequest struct {
}

func (o *CreateBlockDeviceBadRequest) Error() string {
	return fmt.Sprintf("[POST /iaas/api/block-devices][%d] createBlockDeviceBadRequest ", 400)
}

func (o *CreateBlockDeviceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateBlockDeviceForbidden creates a CreateBlockDeviceForbidden with default headers values
func NewCreateBlockDeviceForbidden() *CreateBlockDeviceForbidden {
	return &CreateBlockDeviceForbidden{}
}

/*CreateBlockDeviceForbidden handles this case with default header values.

Forbidden
*/
type CreateBlockDeviceForbidden struct {
}

func (o *CreateBlockDeviceForbidden) Error() string {
	return fmt.Sprintf("[POST /iaas/api/block-devices][%d] createBlockDeviceForbidden ", 403)
}

func (o *CreateBlockDeviceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

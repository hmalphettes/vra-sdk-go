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

// DeleteBlockDeviceReader is a Reader for the DeleteBlockDevice structure.
type DeleteBlockDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteBlockDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewDeleteBlockDeviceAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewDeleteBlockDeviceForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteBlockDeviceNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteBlockDeviceAccepted creates a DeleteBlockDeviceAccepted with default headers values
func NewDeleteBlockDeviceAccepted() *DeleteBlockDeviceAccepted {
	return &DeleteBlockDeviceAccepted{}
}

/*DeleteBlockDeviceAccepted handles this case with default header values.

successful operation
*/
type DeleteBlockDeviceAccepted struct {
	Payload *models.RequestTracker
}

func (o *DeleteBlockDeviceAccepted) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/block-devices/{id}][%d] deleteBlockDeviceAccepted  %+v", 202, o.Payload)
}

func (o *DeleteBlockDeviceAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestTracker)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteBlockDeviceForbidden creates a DeleteBlockDeviceForbidden with default headers values
func NewDeleteBlockDeviceForbidden() *DeleteBlockDeviceForbidden {
	return &DeleteBlockDeviceForbidden{}
}

/*DeleteBlockDeviceForbidden handles this case with default header values.

Forbidden
*/
type DeleteBlockDeviceForbidden struct {
}

func (o *DeleteBlockDeviceForbidden) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/block-devices/{id}][%d] deleteBlockDeviceForbidden ", 403)
}

func (o *DeleteBlockDeviceForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteBlockDeviceNotFound creates a DeleteBlockDeviceNotFound with default headers values
func NewDeleteBlockDeviceNotFound() *DeleteBlockDeviceNotFound {
	return &DeleteBlockDeviceNotFound{}
}

/*DeleteBlockDeviceNotFound handles this case with default header values.

Not Found
*/
type DeleteBlockDeviceNotFound struct {
}

func (o *DeleteBlockDeviceNotFound) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/block-devices/{id}][%d] deleteBlockDeviceNotFound ", 404)
}

func (o *DeleteBlockDeviceNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

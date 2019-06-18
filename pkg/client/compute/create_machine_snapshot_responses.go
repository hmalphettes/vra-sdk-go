// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/cas-sdk-go/pkg/models"
)

// CreateMachineSnapshotReader is a Reader for the CreateMachineSnapshot structure.
type CreateMachineSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateMachineSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewCreateMachineSnapshotAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewCreateMachineSnapshotForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreateMachineSnapshotNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateMachineSnapshotAccepted creates a CreateMachineSnapshotAccepted with default headers values
func NewCreateMachineSnapshotAccepted() *CreateMachineSnapshotAccepted {
	return &CreateMachineSnapshotAccepted{}
}

/*CreateMachineSnapshotAccepted handles this case with default header values.

successful operation
*/
type CreateMachineSnapshotAccepted struct {
	Payload *models.RequestTracker
}

func (o *CreateMachineSnapshotAccepted) Error() string {
	return fmt.Sprintf("[POST /iaas/api/machines/{id}/operations/snapshots][%d] createMachineSnapshotAccepted  %+v", 202, o.Payload)
}

func (o *CreateMachineSnapshotAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestTracker)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateMachineSnapshotForbidden creates a CreateMachineSnapshotForbidden with default headers values
func NewCreateMachineSnapshotForbidden() *CreateMachineSnapshotForbidden {
	return &CreateMachineSnapshotForbidden{}
}

/*CreateMachineSnapshotForbidden handles this case with default header values.

Forbidden
*/
type CreateMachineSnapshotForbidden struct {
}

func (o *CreateMachineSnapshotForbidden) Error() string {
	return fmt.Sprintf("[POST /iaas/api/machines/{id}/operations/snapshots][%d] createMachineSnapshotForbidden ", 403)
}

func (o *CreateMachineSnapshotForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateMachineSnapshotNotFound creates a CreateMachineSnapshotNotFound with default headers values
func NewCreateMachineSnapshotNotFound() *CreateMachineSnapshotNotFound {
	return &CreateMachineSnapshotNotFound{}
}

/*CreateMachineSnapshotNotFound handles this case with default header values.

Not Found
*/
type CreateMachineSnapshotNotFound struct {
}

func (o *CreateMachineSnapshotNotFound) Error() string {
	return fmt.Sprintf("[POST /iaas/api/machines/{id}/operations/snapshots][%d] createMachineSnapshotNotFound ", 404)
}

func (o *CreateMachineSnapshotNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

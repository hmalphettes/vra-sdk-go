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

// DeleteMachineSnapshotReader is a Reader for the DeleteMachineSnapshot structure.
type DeleteMachineSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMachineSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewDeleteMachineSnapshotAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewDeleteMachineSnapshotForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteMachineSnapshotNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteMachineSnapshotAccepted creates a DeleteMachineSnapshotAccepted with default headers values
func NewDeleteMachineSnapshotAccepted() *DeleteMachineSnapshotAccepted {
	return &DeleteMachineSnapshotAccepted{}
}

/*DeleteMachineSnapshotAccepted handles this case with default header values.

successful operation
*/
type DeleteMachineSnapshotAccepted struct {
	Payload *models.RequestTracker
}

func (o *DeleteMachineSnapshotAccepted) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/machines/{id}/snapshots/{id1}][%d] deleteMachineSnapshotAccepted  %+v", 202, o.Payload)
}

func (o *DeleteMachineSnapshotAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestTracker)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMachineSnapshotForbidden creates a DeleteMachineSnapshotForbidden with default headers values
func NewDeleteMachineSnapshotForbidden() *DeleteMachineSnapshotForbidden {
	return &DeleteMachineSnapshotForbidden{}
}

/*DeleteMachineSnapshotForbidden handles this case with default header values.

Forbidden
*/
type DeleteMachineSnapshotForbidden struct {
}

func (o *DeleteMachineSnapshotForbidden) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/machines/{id}/snapshots/{id1}][%d] deleteMachineSnapshotForbidden ", 403)
}

func (o *DeleteMachineSnapshotForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteMachineSnapshotNotFound creates a DeleteMachineSnapshotNotFound with default headers values
func NewDeleteMachineSnapshotNotFound() *DeleteMachineSnapshotNotFound {
	return &DeleteMachineSnapshotNotFound{}
}

/*DeleteMachineSnapshotNotFound handles this case with default header values.

Not Found
*/
type DeleteMachineSnapshotNotFound struct {
}

func (o *DeleteMachineSnapshotNotFound) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/machines/{id}/snapshots/{id1}][%d] deleteMachineSnapshotNotFound ", 404)
}

func (o *DeleteMachineSnapshotNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

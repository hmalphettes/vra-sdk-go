// Code generated by go-swagger; DO NOT EDIT.

package triggers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// DeleteUsingDELETEMixin3Reader is a Reader for the DeleteUsingDELETEMixin3 structure.
type DeleteUsingDELETEMixin3Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUsingDELETEMixin3Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteUsingDELETEMixin3OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteUsingDELETEMixin3Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteUsingDELETEMixin3Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteUsingDELETEMixin3NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteUsingDELETEMixin3InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteUsingDELETEMixin3OK creates a DeleteUsingDELETEMixin3OK with default headers values
func NewDeleteUsingDELETEMixin3OK() *DeleteUsingDELETEMixin3OK {
	return &DeleteUsingDELETEMixin3OK{}
}

/*DeleteUsingDELETEMixin3OK handles this case with default header values.

'Success' with Delete a Docker Registry Event
*/
type DeleteUsingDELETEMixin3OK struct {
	Payload models.DockerRegistryEvent
}

func (o *DeleteUsingDELETEMixin3OK) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/registry-events/{id}][%d] deleteUsingDELETEMixin3OK  %+v", 200, o.Payload)
}

func (o *DeleteUsingDELETEMixin3OK) GetPayload() models.DockerRegistryEvent {
	return o.Payload
}

func (o *DeleteUsingDELETEMixin3OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalDockerRegistryEvent(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewDeleteUsingDELETEMixin3Unauthorized creates a DeleteUsingDELETEMixin3Unauthorized with default headers values
func NewDeleteUsingDELETEMixin3Unauthorized() *DeleteUsingDELETEMixin3Unauthorized {
	return &DeleteUsingDELETEMixin3Unauthorized{}
}

/*DeleteUsingDELETEMixin3Unauthorized handles this case with default header values.

Unauthorized Request
*/
type DeleteUsingDELETEMixin3Unauthorized struct {
}

func (o *DeleteUsingDELETEMixin3Unauthorized) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/registry-events/{id}][%d] deleteUsingDELETEMixin3Unauthorized ", 401)
}

func (o *DeleteUsingDELETEMixin3Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUsingDELETEMixin3Forbidden creates a DeleteUsingDELETEMixin3Forbidden with default headers values
func NewDeleteUsingDELETEMixin3Forbidden() *DeleteUsingDELETEMixin3Forbidden {
	return &DeleteUsingDELETEMixin3Forbidden{}
}

/*DeleteUsingDELETEMixin3Forbidden handles this case with default header values.

Forbidden
*/
type DeleteUsingDELETEMixin3Forbidden struct {
}

func (o *DeleteUsingDELETEMixin3Forbidden) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/registry-events/{id}][%d] deleteUsingDELETEMixin3Forbidden ", 403)
}

func (o *DeleteUsingDELETEMixin3Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUsingDELETEMixin3NotFound creates a DeleteUsingDELETEMixin3NotFound with default headers values
func NewDeleteUsingDELETEMixin3NotFound() *DeleteUsingDELETEMixin3NotFound {
	return &DeleteUsingDELETEMixin3NotFound{}
}

/*DeleteUsingDELETEMixin3NotFound handles this case with default header values.

Not Found
*/
type DeleteUsingDELETEMixin3NotFound struct {
}

func (o *DeleteUsingDELETEMixin3NotFound) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/registry-events/{id}][%d] deleteUsingDELETEMixin3NotFound ", 404)
}

func (o *DeleteUsingDELETEMixin3NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUsingDELETEMixin3InternalServerError creates a DeleteUsingDELETEMixin3InternalServerError with default headers values
func NewDeleteUsingDELETEMixin3InternalServerError() *DeleteUsingDELETEMixin3InternalServerError {
	return &DeleteUsingDELETEMixin3InternalServerError{}
}

/*DeleteUsingDELETEMixin3InternalServerError handles this case with default header values.

Server Error
*/
type DeleteUsingDELETEMixin3InternalServerError struct {
}

func (o *DeleteUsingDELETEMixin3InternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/registry-events/{id}][%d] deleteUsingDELETEMixin3InternalServerError ", 500)
}

func (o *DeleteUsingDELETEMixin3InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
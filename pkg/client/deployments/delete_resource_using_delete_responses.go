// Code generated by go-swagger; DO NOT EDIT.

package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// DeleteResourceUsingDELETEReader is a Reader for the DeleteResourceUsingDELETE structure.
type DeleteResourceUsingDELETEReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteResourceUsingDELETEReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteResourceUsingDELETEOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteResourceUsingDELETEUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteResourceUsingDELETEOK creates a DeleteResourceUsingDELETEOK with default headers values
func NewDeleteResourceUsingDELETEOK() *DeleteResourceUsingDELETEOK {
	return &DeleteResourceUsingDELETEOK{}
}

/*DeleteResourceUsingDELETEOK handles this case with default header values.

OK
*/
type DeleteResourceUsingDELETEOK struct {
	Payload *models.DeploymentRequest
}

func (o *DeleteResourceUsingDELETEOK) Error() string {
	return fmt.Sprintf("[DELETE /deployment/api/deployments/{depId}/resources/{resourceId}][%d] deleteResourceUsingDELETEOK  %+v", 200, o.Payload)
}

func (o *DeleteResourceUsingDELETEOK) GetPayload() *models.DeploymentRequest {
	return o.Payload
}

func (o *DeleteResourceUsingDELETEOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeploymentRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteResourceUsingDELETEUnauthorized creates a DeleteResourceUsingDELETEUnauthorized with default headers values
func NewDeleteResourceUsingDELETEUnauthorized() *DeleteResourceUsingDELETEUnauthorized {
	return &DeleteResourceUsingDELETEUnauthorized{}
}

/*DeleteResourceUsingDELETEUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteResourceUsingDELETEUnauthorized struct {
}

func (o *DeleteResourceUsingDELETEUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /deployment/api/deployments/{depId}/resources/{resourceId}][%d] deleteResourceUsingDELETEUnauthorized ", 401)
}

func (o *DeleteResourceUsingDELETEUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

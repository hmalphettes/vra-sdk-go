// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/cas-sdk-go/pkg/models"
)

// DeleteNetworkReader is a Reader for the DeleteNetwork structure.
type DeleteNetworkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNetworkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteNetworkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 202:
		result := NewDeleteNetworkAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewDeleteNetworkForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteNetworkNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteNetworkOK creates a DeleteNetworkOK with default headers values
func NewDeleteNetworkOK() *DeleteNetworkOK {
	return &DeleteNetworkOK{}
}

/*DeleteNetworkOK handles this case with default header values.

OK
*/
type DeleteNetworkOK struct {
}

func (o *DeleteNetworkOK) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/networks/{id}][%d] deleteNetworkOK ", 200)
}

func (o *DeleteNetworkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNetworkAccepted creates a DeleteNetworkAccepted with default headers values
func NewDeleteNetworkAccepted() *DeleteNetworkAccepted {
	return &DeleteNetworkAccepted{}
}

/*DeleteNetworkAccepted handles this case with default header values.

successful operation
*/
type DeleteNetworkAccepted struct {
	Payload *models.RequestTracker
}

func (o *DeleteNetworkAccepted) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/networks/{id}][%d] deleteNetworkAccepted  %+v", 202, o.Payload)
}

func (o *DeleteNetworkAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RequestTracker)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNetworkForbidden creates a DeleteNetworkForbidden with default headers values
func NewDeleteNetworkForbidden() *DeleteNetworkForbidden {
	return &DeleteNetworkForbidden{}
}

/*DeleteNetworkForbidden handles this case with default header values.

Forbidden
*/
type DeleteNetworkForbidden struct {
}

func (o *DeleteNetworkForbidden) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/networks/{id}][%d] deleteNetworkForbidden ", 403)
}

func (o *DeleteNetworkForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteNetworkNotFound creates a DeleteNetworkNotFound with default headers values
func NewDeleteNetworkNotFound() *DeleteNetworkNotFound {
	return &DeleteNetworkNotFound{}
}

/*DeleteNetworkNotFound handles this case with default header values.

Not Found
*/
type DeleteNetworkNotFound struct {
}

func (o *DeleteNetworkNotFound) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/networks/{id}][%d] deleteNetworkNotFound ", 404)
}

func (o *DeleteNetworkNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

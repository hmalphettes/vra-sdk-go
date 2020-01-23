// Code generated by go-swagger; DO NOT EDIT.

package endpoints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// UpdateByNameUsingPUT1Reader is a Reader for the UpdateByNameUsingPUT1 structure.
type UpdateByNameUsingPUT1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateByNameUsingPUT1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateByNameUsingPUT1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateByNameUsingPUT1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateByNameUsingPUT1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateByNameUsingPUT1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateByNameUsingPUT1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateByNameUsingPUT1OK creates a UpdateByNameUsingPUT1OK with default headers values
func NewUpdateByNameUsingPUT1OK() *UpdateByNameUsingPUT1OK {
	return &UpdateByNameUsingPUT1OK{}
}

/*UpdateByNameUsingPUT1OK handles this case with default header values.

'Success' with the updated Endpoint
*/
type UpdateByNameUsingPUT1OK struct {
	Payload models.Endpoint
}

func (o *UpdateByNameUsingPUT1OK) Error() string {
	return fmt.Sprintf("[PUT /codestream/api/endpoints/{project}/{name}][%d] updateByNameUsingPUT1OK  %+v", 200, o.Payload)
}

func (o *UpdateByNameUsingPUT1OK) GetPayload() models.Endpoint {
	return o.Payload
}

func (o *UpdateByNameUsingPUT1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalEndpoint(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewUpdateByNameUsingPUT1Unauthorized creates a UpdateByNameUsingPUT1Unauthorized with default headers values
func NewUpdateByNameUsingPUT1Unauthorized() *UpdateByNameUsingPUT1Unauthorized {
	return &UpdateByNameUsingPUT1Unauthorized{}
}

/*UpdateByNameUsingPUT1Unauthorized handles this case with default header values.

Unauthorized Request
*/
type UpdateByNameUsingPUT1Unauthorized struct {
}

func (o *UpdateByNameUsingPUT1Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /codestream/api/endpoints/{project}/{name}][%d] updateByNameUsingPUT1Unauthorized ", 401)
}

func (o *UpdateByNameUsingPUT1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateByNameUsingPUT1Forbidden creates a UpdateByNameUsingPUT1Forbidden with default headers values
func NewUpdateByNameUsingPUT1Forbidden() *UpdateByNameUsingPUT1Forbidden {
	return &UpdateByNameUsingPUT1Forbidden{}
}

/*UpdateByNameUsingPUT1Forbidden handles this case with default header values.

Forbidden
*/
type UpdateByNameUsingPUT1Forbidden struct {
}

func (o *UpdateByNameUsingPUT1Forbidden) Error() string {
	return fmt.Sprintf("[PUT /codestream/api/endpoints/{project}/{name}][%d] updateByNameUsingPUT1Forbidden ", 403)
}

func (o *UpdateByNameUsingPUT1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateByNameUsingPUT1NotFound creates a UpdateByNameUsingPUT1NotFound with default headers values
func NewUpdateByNameUsingPUT1NotFound() *UpdateByNameUsingPUT1NotFound {
	return &UpdateByNameUsingPUT1NotFound{}
}

/*UpdateByNameUsingPUT1NotFound handles this case with default header values.

Not Found
*/
type UpdateByNameUsingPUT1NotFound struct {
}

func (o *UpdateByNameUsingPUT1NotFound) Error() string {
	return fmt.Sprintf("[PUT /codestream/api/endpoints/{project}/{name}][%d] updateByNameUsingPUT1NotFound ", 404)
}

func (o *UpdateByNameUsingPUT1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateByNameUsingPUT1InternalServerError creates a UpdateByNameUsingPUT1InternalServerError with default headers values
func NewUpdateByNameUsingPUT1InternalServerError() *UpdateByNameUsingPUT1InternalServerError {
	return &UpdateByNameUsingPUT1InternalServerError{}
}

/*UpdateByNameUsingPUT1InternalServerError handles this case with default header values.

Server Error
*/
type UpdateByNameUsingPUT1InternalServerError struct {
}

func (o *UpdateByNameUsingPUT1InternalServerError) Error() string {
	return fmt.Sprintf("[PUT /codestream/api/endpoints/{project}/{name}][%d] updateByNameUsingPUT1InternalServerError ", 500)
}

func (o *UpdateByNameUsingPUT1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
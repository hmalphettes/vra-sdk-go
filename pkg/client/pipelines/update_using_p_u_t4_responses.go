// Code generated by go-swagger; DO NOT EDIT.

package pipelines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// UpdateUsingPUT4Reader is a Reader for the UpdateUsingPUT4 structure.
type UpdateUsingPUT4Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUsingPUT4Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateUsingPUT4OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateUsingPUT4Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateUsingPUT4Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateUsingPUT4NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateUsingPUT4InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateUsingPUT4OK creates a UpdateUsingPUT4OK with default headers values
func NewUpdateUsingPUT4OK() *UpdateUsingPUT4OK {
	return &UpdateUsingPUT4OK{}
}

/*UpdateUsingPUT4OK handles this case with default header values.

'Success' with the updated Variable
*/
type UpdateUsingPUT4OK struct {
	Payload models.Variable
}

func (o *UpdateUsingPUT4OK) Error() string {
	return fmt.Sprintf("[PUT /codestream/api/variables/{id}][%d] updateUsingPUT4OK  %+v", 200, o.Payload)
}

func (o *UpdateUsingPUT4OK) GetPayload() models.Variable {
	return o.Payload
}

func (o *UpdateUsingPUT4OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalVariable(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewUpdateUsingPUT4Unauthorized creates a UpdateUsingPUT4Unauthorized with default headers values
func NewUpdateUsingPUT4Unauthorized() *UpdateUsingPUT4Unauthorized {
	return &UpdateUsingPUT4Unauthorized{}
}

/*UpdateUsingPUT4Unauthorized handles this case with default header values.

Unauthorized Request
*/
type UpdateUsingPUT4Unauthorized struct {
}

func (o *UpdateUsingPUT4Unauthorized) Error() string {
	return fmt.Sprintf("[PUT /codestream/api/variables/{id}][%d] updateUsingPUT4Unauthorized ", 401)
}

func (o *UpdateUsingPUT4Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUsingPUT4Forbidden creates a UpdateUsingPUT4Forbidden with default headers values
func NewUpdateUsingPUT4Forbidden() *UpdateUsingPUT4Forbidden {
	return &UpdateUsingPUT4Forbidden{}
}

/*UpdateUsingPUT4Forbidden handles this case with default header values.

Forbidden
*/
type UpdateUsingPUT4Forbidden struct {
}

func (o *UpdateUsingPUT4Forbidden) Error() string {
	return fmt.Sprintf("[PUT /codestream/api/variables/{id}][%d] updateUsingPUT4Forbidden ", 403)
}

func (o *UpdateUsingPUT4Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUsingPUT4NotFound creates a UpdateUsingPUT4NotFound with default headers values
func NewUpdateUsingPUT4NotFound() *UpdateUsingPUT4NotFound {
	return &UpdateUsingPUT4NotFound{}
}

/*UpdateUsingPUT4NotFound handles this case with default header values.

Not Found
*/
type UpdateUsingPUT4NotFound struct {
}

func (o *UpdateUsingPUT4NotFound) Error() string {
	return fmt.Sprintf("[PUT /codestream/api/variables/{id}][%d] updateUsingPUT4NotFound ", 404)
}

func (o *UpdateUsingPUT4NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateUsingPUT4InternalServerError creates a UpdateUsingPUT4InternalServerError with default headers values
func NewUpdateUsingPUT4InternalServerError() *UpdateUsingPUT4InternalServerError {
	return &UpdateUsingPUT4InternalServerError{}
}

/*UpdateUsingPUT4InternalServerError handles this case with default header values.

Server Error
*/
type UpdateUsingPUT4InternalServerError struct {
}

func (o *UpdateUsingPUT4InternalServerError) Error() string {
	return fmt.Sprintf("[PUT /codestream/api/variables/{id}][%d] updateUsingPUT4InternalServerError ", 500)
}

func (o *UpdateUsingPUT4InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
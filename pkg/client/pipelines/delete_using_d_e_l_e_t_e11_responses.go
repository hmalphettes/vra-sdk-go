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

// DeleteUsingDELETE11Reader is a Reader for the DeleteUsingDELETE11 structure.
type DeleteUsingDELETE11Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUsingDELETE11Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteUsingDELETE11OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteUsingDELETE11Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteUsingDELETE11Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteUsingDELETE11NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteUsingDELETE11InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteUsingDELETE11OK creates a DeleteUsingDELETE11OK with default headers values
func NewDeleteUsingDELETE11OK() *DeleteUsingDELETE11OK {
	return &DeleteUsingDELETE11OK{}
}

/*DeleteUsingDELETE11OK handles this case with default header values.

'Success' with deleted Variable
*/
type DeleteUsingDELETE11OK struct {
	Payload models.Variable
}

func (o *DeleteUsingDELETE11OK) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/variables/{id}][%d] deleteUsingDELETE11OK  %+v", 200, o.Payload)
}

func (o *DeleteUsingDELETE11OK) GetPayload() models.Variable {
	return o.Payload
}

func (o *DeleteUsingDELETE11OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalVariable(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewDeleteUsingDELETE11Unauthorized creates a DeleteUsingDELETE11Unauthorized with default headers values
func NewDeleteUsingDELETE11Unauthorized() *DeleteUsingDELETE11Unauthorized {
	return &DeleteUsingDELETE11Unauthorized{}
}

/*DeleteUsingDELETE11Unauthorized handles this case with default header values.

Unauthorized Request
*/
type DeleteUsingDELETE11Unauthorized struct {
}

func (o *DeleteUsingDELETE11Unauthorized) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/variables/{id}][%d] deleteUsingDELETE11Unauthorized ", 401)
}

func (o *DeleteUsingDELETE11Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUsingDELETE11Forbidden creates a DeleteUsingDELETE11Forbidden with default headers values
func NewDeleteUsingDELETE11Forbidden() *DeleteUsingDELETE11Forbidden {
	return &DeleteUsingDELETE11Forbidden{}
}

/*DeleteUsingDELETE11Forbidden handles this case with default header values.

Forbidden
*/
type DeleteUsingDELETE11Forbidden struct {
}

func (o *DeleteUsingDELETE11Forbidden) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/variables/{id}][%d] deleteUsingDELETE11Forbidden ", 403)
}

func (o *DeleteUsingDELETE11Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUsingDELETE11NotFound creates a DeleteUsingDELETE11NotFound with default headers values
func NewDeleteUsingDELETE11NotFound() *DeleteUsingDELETE11NotFound {
	return &DeleteUsingDELETE11NotFound{}
}

/*DeleteUsingDELETE11NotFound handles this case with default header values.

Not Found
*/
type DeleteUsingDELETE11NotFound struct {
}

func (o *DeleteUsingDELETE11NotFound) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/variables/{id}][%d] deleteUsingDELETE11NotFound ", 404)
}

func (o *DeleteUsingDELETE11NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUsingDELETE11InternalServerError creates a DeleteUsingDELETE11InternalServerError with default headers values
func NewDeleteUsingDELETE11InternalServerError() *DeleteUsingDELETE11InternalServerError {
	return &DeleteUsingDELETE11InternalServerError{}
}

/*DeleteUsingDELETE11InternalServerError handles this case with default header values.

Server Error
*/
type DeleteUsingDELETE11InternalServerError struct {
}

func (o *DeleteUsingDELETE11InternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/variables/{id}][%d] deleteUsingDELETE11InternalServerError ", 500)
}

func (o *DeleteUsingDELETE11InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
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

// DeleteByIDUsingDELETE1Reader is a Reader for the DeleteByIDUsingDELETE1 structure.
type DeleteByIDUsingDELETE1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteByIDUsingDELETE1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteByIDUsingDELETE1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteByIDUsingDELETE1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteByIDUsingDELETE1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteByIDUsingDELETE1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteByIDUsingDELETE1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteByIDUsingDELETE1OK creates a DeleteByIDUsingDELETE1OK with default headers values
func NewDeleteByIDUsingDELETE1OK() *DeleteByIDUsingDELETE1OK {
	return &DeleteByIDUsingDELETE1OK{}
}

/*DeleteByIDUsingDELETE1OK handles this case with default header values.

'Success' with Gerrit Listener Delete
*/
type DeleteByIDUsingDELETE1OK struct {
	Payload models.GerritListener
}

func (o *DeleteByIDUsingDELETE1OK) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/gerrit-listeners/{id}][%d] deleteByIdUsingDELETE1OK  %+v", 200, o.Payload)
}

func (o *DeleteByIDUsingDELETE1OK) GetPayload() models.GerritListener {
	return o.Payload
}

func (o *DeleteByIDUsingDELETE1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalGerritListener(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewDeleteByIDUsingDELETE1Unauthorized creates a DeleteByIDUsingDELETE1Unauthorized with default headers values
func NewDeleteByIDUsingDELETE1Unauthorized() *DeleteByIDUsingDELETE1Unauthorized {
	return &DeleteByIDUsingDELETE1Unauthorized{}
}

/*DeleteByIDUsingDELETE1Unauthorized handles this case with default header values.

Unauthorized Request
*/
type DeleteByIDUsingDELETE1Unauthorized struct {
}

func (o *DeleteByIDUsingDELETE1Unauthorized) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/gerrit-listeners/{id}][%d] deleteByIdUsingDELETE1Unauthorized ", 401)
}

func (o *DeleteByIDUsingDELETE1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteByIDUsingDELETE1Forbidden creates a DeleteByIDUsingDELETE1Forbidden with default headers values
func NewDeleteByIDUsingDELETE1Forbidden() *DeleteByIDUsingDELETE1Forbidden {
	return &DeleteByIDUsingDELETE1Forbidden{}
}

/*DeleteByIDUsingDELETE1Forbidden handles this case with default header values.

Forbidden
*/
type DeleteByIDUsingDELETE1Forbidden struct {
}

func (o *DeleteByIDUsingDELETE1Forbidden) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/gerrit-listeners/{id}][%d] deleteByIdUsingDELETE1Forbidden ", 403)
}

func (o *DeleteByIDUsingDELETE1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteByIDUsingDELETE1NotFound creates a DeleteByIDUsingDELETE1NotFound with default headers values
func NewDeleteByIDUsingDELETE1NotFound() *DeleteByIDUsingDELETE1NotFound {
	return &DeleteByIDUsingDELETE1NotFound{}
}

/*DeleteByIDUsingDELETE1NotFound handles this case with default header values.

Not Found
*/
type DeleteByIDUsingDELETE1NotFound struct {
}

func (o *DeleteByIDUsingDELETE1NotFound) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/gerrit-listeners/{id}][%d] deleteByIdUsingDELETE1NotFound ", 404)
}

func (o *DeleteByIDUsingDELETE1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteByIDUsingDELETE1InternalServerError creates a DeleteByIDUsingDELETE1InternalServerError with default headers values
func NewDeleteByIDUsingDELETE1InternalServerError() *DeleteByIDUsingDELETE1InternalServerError {
	return &DeleteByIDUsingDELETE1InternalServerError{}
}

/*DeleteByIDUsingDELETE1InternalServerError handles this case with default header values.

Server Error
*/
type DeleteByIDUsingDELETE1InternalServerError struct {
}

func (o *DeleteByIDUsingDELETE1InternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /codestream/api/gerrit-listeners/{id}][%d] deleteByIdUsingDELETE1InternalServerError ", 500)
}

func (o *DeleteByIDUsingDELETE1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
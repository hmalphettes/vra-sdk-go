// Code generated by go-swagger; DO NOT EDIT.

package executions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// GetUsingGET4Reader is a Reader for the GetUsingGET4 structure.
type GetUsingGET4Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsingGET4Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsingGET4OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetUsingGET4Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUsingGET4Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUsingGET4NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUsingGET4InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUsingGET4OK creates a GetUsingGET4OK with default headers values
func NewGetUsingGET4OK() *GetUsingGET4OK {
	return &GetUsingGET4OK{}
}

/*GetUsingGET4OK handles this case with default header values.

'Success' with the requested User Operation
*/
type GetUsingGET4OK struct {
	Payload models.UserOperation
}

func (o *GetUsingGET4OK) Error() string {
	return fmt.Sprintf("[GET /codestream/api/user-operations/{id}][%d] getUsingGET4OK  %+v", 200, o.Payload)
}

func (o *GetUsingGET4OK) GetPayload() models.UserOperation {
	return o.Payload
}

func (o *GetUsingGET4OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalUserOperation(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewGetUsingGET4Unauthorized creates a GetUsingGET4Unauthorized with default headers values
func NewGetUsingGET4Unauthorized() *GetUsingGET4Unauthorized {
	return &GetUsingGET4Unauthorized{}
}

/*GetUsingGET4Unauthorized handles this case with default header values.

Unauthorized Request
*/
type GetUsingGET4Unauthorized struct {
}

func (o *GetUsingGET4Unauthorized) Error() string {
	return fmt.Sprintf("[GET /codestream/api/user-operations/{id}][%d] getUsingGET4Unauthorized ", 401)
}

func (o *GetUsingGET4Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsingGET4Forbidden creates a GetUsingGET4Forbidden with default headers values
func NewGetUsingGET4Forbidden() *GetUsingGET4Forbidden {
	return &GetUsingGET4Forbidden{}
}

/*GetUsingGET4Forbidden handles this case with default header values.

Forbidden
*/
type GetUsingGET4Forbidden struct {
}

func (o *GetUsingGET4Forbidden) Error() string {
	return fmt.Sprintf("[GET /codestream/api/user-operations/{id}][%d] getUsingGET4Forbidden ", 403)
}

func (o *GetUsingGET4Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsingGET4NotFound creates a GetUsingGET4NotFound with default headers values
func NewGetUsingGET4NotFound() *GetUsingGET4NotFound {
	return &GetUsingGET4NotFound{}
}

/*GetUsingGET4NotFound handles this case with default header values.

Not Found
*/
type GetUsingGET4NotFound struct {
}

func (o *GetUsingGET4NotFound) Error() string {
	return fmt.Sprintf("[GET /codestream/api/user-operations/{id}][%d] getUsingGET4NotFound ", 404)
}

func (o *GetUsingGET4NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsingGET4InternalServerError creates a GetUsingGET4InternalServerError with default headers values
func NewGetUsingGET4InternalServerError() *GetUsingGET4InternalServerError {
	return &GetUsingGET4InternalServerError{}
}

/*GetUsingGET4InternalServerError handles this case with default header values.

Server Error
*/
type GetUsingGET4InternalServerError struct {
}

func (o *GetUsingGET4InternalServerError) Error() string {
	return fmt.Sprintf("[GET /codestream/api/user-operations/{id}][%d] getUsingGET4InternalServerError ", 500)
}

func (o *GetUsingGET4InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
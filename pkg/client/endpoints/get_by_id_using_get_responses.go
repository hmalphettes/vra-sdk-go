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

// GetByIDUsingGETReader is a Reader for the GetByIDUsingGET structure.
type GetByIDUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetByIDUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetByIDUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetByIDUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetByIDUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetByIDUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetByIDUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetByIDUsingGETOK creates a GetByIDUsingGETOK with default headers values
func NewGetByIDUsingGETOK() *GetByIDUsingGETOK {
	return &GetByIDUsingGETOK{}
}

/*GetByIDUsingGETOK handles this case with default header values.

'Success' with the requested Endpoint
*/
type GetByIDUsingGETOK struct {
	Payload models.Endpoint
}

func (o *GetByIDUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /codestream/api/endpoints/{id}][%d] getByIdUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetByIDUsingGETOK) GetPayload() models.Endpoint {
	return o.Payload
}

func (o *GetByIDUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalEndpoint(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewGetByIDUsingGETUnauthorized creates a GetByIDUsingGETUnauthorized with default headers values
func NewGetByIDUsingGETUnauthorized() *GetByIDUsingGETUnauthorized {
	return &GetByIDUsingGETUnauthorized{}
}

/*GetByIDUsingGETUnauthorized handles this case with default header values.

Unauthorized Request
*/
type GetByIDUsingGETUnauthorized struct {
}

func (o *GetByIDUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /codestream/api/endpoints/{id}][%d] getByIdUsingGETUnauthorized ", 401)
}

func (o *GetByIDUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetByIDUsingGETForbidden creates a GetByIDUsingGETForbidden with default headers values
func NewGetByIDUsingGETForbidden() *GetByIDUsingGETForbidden {
	return &GetByIDUsingGETForbidden{}
}

/*GetByIDUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetByIDUsingGETForbidden struct {
}

func (o *GetByIDUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /codestream/api/endpoints/{id}][%d] getByIdUsingGETForbidden ", 403)
}

func (o *GetByIDUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetByIDUsingGETNotFound creates a GetByIDUsingGETNotFound with default headers values
func NewGetByIDUsingGETNotFound() *GetByIDUsingGETNotFound {
	return &GetByIDUsingGETNotFound{}
}

/*GetByIDUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetByIDUsingGETNotFound struct {
}

func (o *GetByIDUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /codestream/api/endpoints/{id}][%d] getByIdUsingGETNotFound ", 404)
}

func (o *GetByIDUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetByIDUsingGETInternalServerError creates a GetByIDUsingGETInternalServerError with default headers values
func NewGetByIDUsingGETInternalServerError() *GetByIDUsingGETInternalServerError {
	return &GetByIDUsingGETInternalServerError{}
}

/*GetByIDUsingGETInternalServerError handles this case with default header values.

Server Error
*/
type GetByIDUsingGETInternalServerError struct {
}

func (o *GetByIDUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /codestream/api/endpoints/{id}][%d] getByIdUsingGETInternalServerError ", 500)
}

func (o *GetByIDUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
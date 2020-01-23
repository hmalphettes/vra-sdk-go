// Code generated by go-swagger; DO NOT EDIT.

package endpoints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// GetAllUsingGETReader is a Reader for the GetAllUsingGET structure.
type GetAllUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAllUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetAllUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAllUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAllUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAllUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAllUsingGETOK creates a GetAllUsingGETOK with default headers values
func NewGetAllUsingGETOK() *GetAllUsingGETOK {
	return &GetAllUsingGETOK{}
}

/*GetAllUsingGETOK handles this case with default header values.

'Success' with the requested Endpoints
*/
type GetAllUsingGETOK struct {
	Payload *models.Endpoints
}

func (o *GetAllUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /codestream/api/endpoints][%d] getAllUsingGETOK  %+v", 200, o.Payload)
}

func (o *GetAllUsingGETOK) GetPayload() *models.Endpoints {
	return o.Payload
}

func (o *GetAllUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Endpoints)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllUsingGETUnauthorized creates a GetAllUsingGETUnauthorized with default headers values
func NewGetAllUsingGETUnauthorized() *GetAllUsingGETUnauthorized {
	return &GetAllUsingGETUnauthorized{}
}

/*GetAllUsingGETUnauthorized handles this case with default header values.

Unauthorized Request
*/
type GetAllUsingGETUnauthorized struct {
}

func (o *GetAllUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /codestream/api/endpoints][%d] getAllUsingGETUnauthorized ", 401)
}

func (o *GetAllUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllUsingGETForbidden creates a GetAllUsingGETForbidden with default headers values
func NewGetAllUsingGETForbidden() *GetAllUsingGETForbidden {
	return &GetAllUsingGETForbidden{}
}

/*GetAllUsingGETForbidden handles this case with default header values.

Forbidden
*/
type GetAllUsingGETForbidden struct {
}

func (o *GetAllUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /codestream/api/endpoints][%d] getAllUsingGETForbidden ", 403)
}

func (o *GetAllUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllUsingGETNotFound creates a GetAllUsingGETNotFound with default headers values
func NewGetAllUsingGETNotFound() *GetAllUsingGETNotFound {
	return &GetAllUsingGETNotFound{}
}

/*GetAllUsingGETNotFound handles this case with default header values.

Not Found
*/
type GetAllUsingGETNotFound struct {
}

func (o *GetAllUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /codestream/api/endpoints][%d] getAllUsingGETNotFound ", 404)
}

func (o *GetAllUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllUsingGETInternalServerError creates a GetAllUsingGETInternalServerError with default headers values
func NewGetAllUsingGETInternalServerError() *GetAllUsingGETInternalServerError {
	return &GetAllUsingGETInternalServerError{}
}

/*GetAllUsingGETInternalServerError handles this case with default header values.

Server Error
*/
type GetAllUsingGETInternalServerError struct {
}

func (o *GetAllUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /codestream/api/endpoints][%d] getAllUsingGETInternalServerError ", 500)
}

func (o *GetAllUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
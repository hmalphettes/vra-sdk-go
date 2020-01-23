// Code generated by go-swagger; DO NOT EDIT.

package about

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// HandleGetAboutUsingGETReader is a Reader for the HandleGetAboutUsingGET structure.
type HandleGetAboutUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HandleGetAboutUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewHandleGetAboutUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewHandleGetAboutUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewHandleGetAboutUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewHandleGetAboutUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewHandleGetAboutUsingGETInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewHandleGetAboutUsingGETOK creates a HandleGetAboutUsingGETOK with default headers values
func NewHandleGetAboutUsingGETOK() *HandleGetAboutUsingGETOK {
	return &HandleGetAboutUsingGETOK{}
}

/*HandleGetAboutUsingGETOK handles this case with default header values.

'Success' with the requested Endpoint
*/
type HandleGetAboutUsingGETOK struct {
	Payload *models.About
}

func (o *HandleGetAboutUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /codestream/api/about][%d] handleGetAboutUsingGETOK  %+v", 200, o.Payload)
}

func (o *HandleGetAboutUsingGETOK) GetPayload() *models.About {
	return o.Payload
}

func (o *HandleGetAboutUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.About)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHandleGetAboutUsingGETUnauthorized creates a HandleGetAboutUsingGETUnauthorized with default headers values
func NewHandleGetAboutUsingGETUnauthorized() *HandleGetAboutUsingGETUnauthorized {
	return &HandleGetAboutUsingGETUnauthorized{}
}

/*HandleGetAboutUsingGETUnauthorized handles this case with default header values.

Unauthorized Request
*/
type HandleGetAboutUsingGETUnauthorized struct {
}

func (o *HandleGetAboutUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /codestream/api/about][%d] handleGetAboutUsingGETUnauthorized ", 401)
}

func (o *HandleGetAboutUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHandleGetAboutUsingGETForbidden creates a HandleGetAboutUsingGETForbidden with default headers values
func NewHandleGetAboutUsingGETForbidden() *HandleGetAboutUsingGETForbidden {
	return &HandleGetAboutUsingGETForbidden{}
}

/*HandleGetAboutUsingGETForbidden handles this case with default header values.

Forbidden
*/
type HandleGetAboutUsingGETForbidden struct {
}

func (o *HandleGetAboutUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /codestream/api/about][%d] handleGetAboutUsingGETForbidden ", 403)
}

func (o *HandleGetAboutUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHandleGetAboutUsingGETNotFound creates a HandleGetAboutUsingGETNotFound with default headers values
func NewHandleGetAboutUsingGETNotFound() *HandleGetAboutUsingGETNotFound {
	return &HandleGetAboutUsingGETNotFound{}
}

/*HandleGetAboutUsingGETNotFound handles this case with default header values.

Not Found
*/
type HandleGetAboutUsingGETNotFound struct {
}

func (o *HandleGetAboutUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /codestream/api/about][%d] handleGetAboutUsingGETNotFound ", 404)
}

func (o *HandleGetAboutUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewHandleGetAboutUsingGETInternalServerError creates a HandleGetAboutUsingGETInternalServerError with default headers values
func NewHandleGetAboutUsingGETInternalServerError() *HandleGetAboutUsingGETInternalServerError {
	return &HandleGetAboutUsingGETInternalServerError{}
}

/*HandleGetAboutUsingGETInternalServerError handles this case with default header values.

Server Error
*/
type HandleGetAboutUsingGETInternalServerError struct {
}

func (o *HandleGetAboutUsingGETInternalServerError) Error() string {
	return fmt.Sprintf("[GET /codestream/api/about][%d] handleGetAboutUsingGETInternalServerError ", 500)
}

func (o *HandleGetAboutUsingGETInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
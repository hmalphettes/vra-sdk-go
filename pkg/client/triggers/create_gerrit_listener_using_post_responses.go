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

// CreateGerritListenerUsingPOSTReader is a Reader for the CreateGerritListenerUsingPOST structure.
type CreateGerritListenerUsingPOSTReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateGerritListenerUsingPOSTReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateGerritListenerUsingPOSTOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateGerritListenerUsingPOSTUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateGerritListenerUsingPOSTForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateGerritListenerUsingPOSTNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateGerritListenerUsingPOSTInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateGerritListenerUsingPOSTOK creates a CreateGerritListenerUsingPOSTOK with default headers values
func NewCreateGerritListenerUsingPOSTOK() *CreateGerritListenerUsingPOSTOK {
	return &CreateGerritListenerUsingPOSTOK{}
}

/*CreateGerritListenerUsingPOSTOK handles this case with default header values.

'Success' with Gerrit Listener Creation
*/
type CreateGerritListenerUsingPOSTOK struct {
	Payload models.GerritListener
}

func (o *CreateGerritListenerUsingPOSTOK) Error() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-listeners][%d] createGerritListenerUsingPOSTOK  %+v", 200, o.Payload)
}

func (o *CreateGerritListenerUsingPOSTOK) GetPayload() models.GerritListener {
	return o.Payload
}

func (o *CreateGerritListenerUsingPOSTOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalGerritListener(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewCreateGerritListenerUsingPOSTUnauthorized creates a CreateGerritListenerUsingPOSTUnauthorized with default headers values
func NewCreateGerritListenerUsingPOSTUnauthorized() *CreateGerritListenerUsingPOSTUnauthorized {
	return &CreateGerritListenerUsingPOSTUnauthorized{}
}

/*CreateGerritListenerUsingPOSTUnauthorized handles this case with default header values.

Unauthorized Request
*/
type CreateGerritListenerUsingPOSTUnauthorized struct {
}

func (o *CreateGerritListenerUsingPOSTUnauthorized) Error() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-listeners][%d] createGerritListenerUsingPOSTUnauthorized ", 401)
}

func (o *CreateGerritListenerUsingPOSTUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateGerritListenerUsingPOSTForbidden creates a CreateGerritListenerUsingPOSTForbidden with default headers values
func NewCreateGerritListenerUsingPOSTForbidden() *CreateGerritListenerUsingPOSTForbidden {
	return &CreateGerritListenerUsingPOSTForbidden{}
}

/*CreateGerritListenerUsingPOSTForbidden handles this case with default header values.

Forbidden
*/
type CreateGerritListenerUsingPOSTForbidden struct {
}

func (o *CreateGerritListenerUsingPOSTForbidden) Error() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-listeners][%d] createGerritListenerUsingPOSTForbidden ", 403)
}

func (o *CreateGerritListenerUsingPOSTForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateGerritListenerUsingPOSTNotFound creates a CreateGerritListenerUsingPOSTNotFound with default headers values
func NewCreateGerritListenerUsingPOSTNotFound() *CreateGerritListenerUsingPOSTNotFound {
	return &CreateGerritListenerUsingPOSTNotFound{}
}

/*CreateGerritListenerUsingPOSTNotFound handles this case with default header values.

Not Found
*/
type CreateGerritListenerUsingPOSTNotFound struct {
}

func (o *CreateGerritListenerUsingPOSTNotFound) Error() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-listeners][%d] createGerritListenerUsingPOSTNotFound ", 404)
}

func (o *CreateGerritListenerUsingPOSTNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateGerritListenerUsingPOSTInternalServerError creates a CreateGerritListenerUsingPOSTInternalServerError with default headers values
func NewCreateGerritListenerUsingPOSTInternalServerError() *CreateGerritListenerUsingPOSTInternalServerError {
	return &CreateGerritListenerUsingPOSTInternalServerError{}
}

/*CreateGerritListenerUsingPOSTInternalServerError handles this case with default header values.

Server Error
*/
type CreateGerritListenerUsingPOSTInternalServerError struct {
}

func (o *CreateGerritListenerUsingPOSTInternalServerError) Error() string {
	return fmt.Sprintf("[POST /codestream/api/gerrit-listeners][%d] createGerritListenerUsingPOSTInternalServerError ", 500)
}

func (o *CreateGerritListenerUsingPOSTInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
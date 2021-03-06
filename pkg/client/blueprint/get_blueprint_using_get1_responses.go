// Code generated by go-swagger; DO NOT EDIT.

package blueprint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// GetBlueprintUsingGET1Reader is a Reader for the GetBlueprintUsingGET1 structure.
type GetBlueprintUsingGET1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBlueprintUsingGET1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetBlueprintUsingGET1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetBlueprintUsingGET1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetBlueprintUsingGET1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetBlueprintUsingGET1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBlueprintUsingGET1OK creates a GetBlueprintUsingGET1OK with default headers values
func NewGetBlueprintUsingGET1OK() *GetBlueprintUsingGET1OK {
	return &GetBlueprintUsingGET1OK{}
}

/*GetBlueprintUsingGET1OK handles this case with default header values.

OK
*/
type GetBlueprintUsingGET1OK struct {
	Payload *models.Blueprint
}

func (o *GetBlueprintUsingGET1OK) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprints/{blueprintId}][%d] getBlueprintUsingGET1OK  %+v", 200, o.Payload)
}

func (o *GetBlueprintUsingGET1OK) GetPayload() *models.Blueprint {
	return o.Payload
}

func (o *GetBlueprintUsingGET1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Blueprint)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBlueprintUsingGET1Unauthorized creates a GetBlueprintUsingGET1Unauthorized with default headers values
func NewGetBlueprintUsingGET1Unauthorized() *GetBlueprintUsingGET1Unauthorized {
	return &GetBlueprintUsingGET1Unauthorized{}
}

/*GetBlueprintUsingGET1Unauthorized handles this case with default header values.

Unauthorized
*/
type GetBlueprintUsingGET1Unauthorized struct {
}

func (o *GetBlueprintUsingGET1Unauthorized) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprints/{blueprintId}][%d] getBlueprintUsingGET1Unauthorized ", 401)
}

func (o *GetBlueprintUsingGET1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBlueprintUsingGET1Forbidden creates a GetBlueprintUsingGET1Forbidden with default headers values
func NewGetBlueprintUsingGET1Forbidden() *GetBlueprintUsingGET1Forbidden {
	return &GetBlueprintUsingGET1Forbidden{}
}

/*GetBlueprintUsingGET1Forbidden handles this case with default header values.

Forbidden
*/
type GetBlueprintUsingGET1Forbidden struct {
}

func (o *GetBlueprintUsingGET1Forbidden) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprints/{blueprintId}][%d] getBlueprintUsingGET1Forbidden ", 403)
}

func (o *GetBlueprintUsingGET1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetBlueprintUsingGET1NotFound creates a GetBlueprintUsingGET1NotFound with default headers values
func NewGetBlueprintUsingGET1NotFound() *GetBlueprintUsingGET1NotFound {
	return &GetBlueprintUsingGET1NotFound{}
}

/*GetBlueprintUsingGET1NotFound handles this case with default header values.

Not Found
*/
type GetBlueprintUsingGET1NotFound struct {
}

func (o *GetBlueprintUsingGET1NotFound) Error() string {
	return fmt.Sprintf("[GET /blueprint/api/blueprints/{blueprintId}][%d] getBlueprintUsingGET1NotFound ", 404)
}

func (o *GetBlueprintUsingGET1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

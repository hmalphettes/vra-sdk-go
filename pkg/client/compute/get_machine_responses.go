// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/vra-sdk-go/pkg/models"
)

// GetMachineReader is a Reader for the GetMachine structure.
type GetMachineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMachineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMachineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetMachineForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetMachineNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetMachineOK creates a GetMachineOK with default headers values
func NewGetMachineOK() *GetMachineOK {
	return &GetMachineOK{}
}

/*GetMachineOK handles this case with default header values.

successful operation
*/
type GetMachineOK struct {
	Payload *models.Machine
}

func (o *GetMachineOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/machines/{id}][%d] getMachineOK  %+v", 200, o.Payload)
}

func (o *GetMachineOK) GetPayload() *models.Machine {
	return o.Payload
}

func (o *GetMachineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Machine)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMachineForbidden creates a GetMachineForbidden with default headers values
func NewGetMachineForbidden() *GetMachineForbidden {
	return &GetMachineForbidden{}
}

/*GetMachineForbidden handles this case with default header values.

Forbidden
*/
type GetMachineForbidden struct {
}

func (o *GetMachineForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/machines/{id}][%d] getMachineForbidden ", 403)
}

func (o *GetMachineForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetMachineNotFound creates a GetMachineNotFound with default headers values
func NewGetMachineNotFound() *GetMachineNotFound {
	return &GetMachineNotFound{}
}

/*GetMachineNotFound handles this case with default header values.

Not Found
*/
type GetMachineNotFound struct {
}

func (o *GetMachineNotFound) Error() string {
	return fmt.Sprintf("[GET /iaas/api/machines/{id}][%d] getMachineNotFound ", 404)
}

func (o *GetMachineNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

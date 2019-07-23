// Code generated by go-swagger; DO NOT EDIT.

package fabric_azure_storage_account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vmware/cas-sdk-go/pkg/models"
)

// GetFabricAzureStorageAccountReader is a Reader for the GetFabricAzureStorageAccount structure.
type GetFabricAzureStorageAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFabricAzureStorageAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFabricAzureStorageAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 403:
		result := NewGetFabricAzureStorageAccountForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetFabricAzureStorageAccountNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFabricAzureStorageAccountOK creates a GetFabricAzureStorageAccountOK with default headers values
func NewGetFabricAzureStorageAccountOK() *GetFabricAzureStorageAccountOK {
	return &GetFabricAzureStorageAccountOK{}
}

/*GetFabricAzureStorageAccountOK handles this case with default header values.

successful operation
*/
type GetFabricAzureStorageAccountOK struct {
	Payload *models.FabricAzureStorageAccount
}

func (o *GetFabricAzureStorageAccountOK) Error() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-azure-storage-accounts/{id}][%d] getFabricAzureStorageAccountOK  %+v", 200, o.Payload)
}

func (o *GetFabricAzureStorageAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FabricAzureStorageAccount)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFabricAzureStorageAccountForbidden creates a GetFabricAzureStorageAccountForbidden with default headers values
func NewGetFabricAzureStorageAccountForbidden() *GetFabricAzureStorageAccountForbidden {
	return &GetFabricAzureStorageAccountForbidden{}
}

/*GetFabricAzureStorageAccountForbidden handles this case with default header values.

Forbidden
*/
type GetFabricAzureStorageAccountForbidden struct {
}

func (o *GetFabricAzureStorageAccountForbidden) Error() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-azure-storage-accounts/{id}][%d] getFabricAzureStorageAccountForbidden ", 403)
}

func (o *GetFabricAzureStorageAccountForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFabricAzureStorageAccountNotFound creates a GetFabricAzureStorageAccountNotFound with default headers values
func NewGetFabricAzureStorageAccountNotFound() *GetFabricAzureStorageAccountNotFound {
	return &GetFabricAzureStorageAccountNotFound{}
}

/*GetFabricAzureStorageAccountNotFound handles this case with default header values.

Not Found
*/
type GetFabricAzureStorageAccountNotFound struct {
}

func (o *GetFabricAzureStorageAccountNotFound) Error() string {
	return fmt.Sprintf("[GET /iaas/api/fabric-azure-storage-accounts/{id}][%d] getFabricAzureStorageAccountNotFound ", 404)
}

func (o *GetFabricAzureStorageAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

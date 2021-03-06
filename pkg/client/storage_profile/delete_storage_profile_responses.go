// Code generated by go-swagger; DO NOT EDIT.

package storage_profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteStorageProfileReader is a Reader for the DeleteStorageProfile structure.
type DeleteStorageProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteStorageProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteStorageProfileNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteStorageProfileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteStorageProfileNoContent creates a DeleteStorageProfileNoContent with default headers values
func NewDeleteStorageProfileNoContent() *DeleteStorageProfileNoContent {
	return &DeleteStorageProfileNoContent{}
}

/*DeleteStorageProfileNoContent handles this case with default header values.

No Content
*/
type DeleteStorageProfileNoContent struct {
}

func (o *DeleteStorageProfileNoContent) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/storage-profiles/{id}][%d] deleteStorageProfileNoContent ", 204)
}

func (o *DeleteStorageProfileNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteStorageProfileForbidden creates a DeleteStorageProfileForbidden with default headers values
func NewDeleteStorageProfileForbidden() *DeleteStorageProfileForbidden {
	return &DeleteStorageProfileForbidden{}
}

/*DeleteStorageProfileForbidden handles this case with default header values.

Forbidden
*/
type DeleteStorageProfileForbidden struct {
}

func (o *DeleteStorageProfileForbidden) Error() string {
	return fmt.Sprintf("[DELETE /iaas/api/storage-profiles/{id}][%d] deleteStorageProfileForbidden ", 403)
}

func (o *DeleteStorageProfileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
